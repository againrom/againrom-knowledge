# Claim registry — MAGIC (spells: the book, the cast, the damage and the resistance)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/magic/format.md`](../formats/magic/format.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

Not a file format — the **simulation area** that turns a learned spell into a wound or a buff. It
sits on `claims/hero.md` (the skills, Mind, mana and the five protections are that ledger's fields)
and on `claims/databin.md` (the `Data.bin` **Spells** table is the parameter source), and it owns
the two things neither of those does: the `Spell` / `Spellbook` / `Effect_DirectDamage` classes,
and the one routine that resolves a hit of any kind. Opened by
[EXP-0065](../experiments/EXP-0065-magic/).

**Why its own ledger and not more `HERO-*` rows.** The consumer question is different — "cast a
spell" rather than "build a hero" — and half of what is below is reached by actors that are not
heroes at all: a monster with a `knownSpells` column, an enchanted weapon, the AI. Two rows this
round went the other way for the same test and are in `claims/hero.md`, because their subject is
the actor's own fields: `HERO-REGEN-021` and `HERO-DAMAGE-022`.

Vocabulary is the **engine's own**: the 28 spell names are the array at `L05034` in `rom.exe`
and `Data.bin`'s own `Spell Name` column, which agree 28/28; the column names are `Data.bin`'s
title arrays; `permanent`/`singleuse`/`charges`/`duration`/`continuous` are the image's literals.
Runtime column *i* is title *i+1* throughout (`DAT-GRAM-003`, `HERO-EQUIP-017`). `__ftol`
(`R0279`) truncates toward zero; `IDIV` truncates toward zero.

**EXP-0243 correction controlling the older rows below.** `MAGIC-POWER-004`'s and
`MAGIC-SING-019`'s `min(power/20+2,7)` quantity is Prismatic Spray's ranked-path final victim-list
count, not a radius: the primary victim is first and consumes one place. Empty A returns the primary
before reading the cap. `MAGIC-TARGET-017`'s diplomacy
absence applies to the spell-effect dispatcher, not this selector; the selector directly invokes
the group candidate builder, whose ordinary secondary candidates have already passed its
first-member diplomacy filter. `MAGIC-BOLTLIST-071`'s open fill is closed: simulation sends those
selected ids in list order, and client opcode `0x8a` or `0x8c` copies them into the caster or
temporary projectile source at `+0xac/+0xb0`. The former `L05035` attribution is unrelated.
Those changed-meaning clauses are recorded in `retracted.md`; the new `MAGIC-SPRAY-*` rows are the
current account.

| ID | Claim | Confidence | Status | Evidence |
|----|-------|------------|--------|----------|
| MAGIC-REACH-178 | Fire Ball's calculated range and the order's range are separate stored values. | High / Unknown | ● active | [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/measurements.json`, `verification.md`) |
| MAGIC-REACH-179 | Unit-target cast order8 has a facing gate followed by a footprint-gap comparison, with a self-target bypass. | High / Medium | ● active | [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/orders.tsv`, `evidence/instructions.tsv`) |
| MAGIC-REACH-180 | Point-target cast order9 uses whole-cell maximum-axis distance, and neither selected predicate directly reads altitude. | High | ● active | [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/`, `EXP-0369.md`) |
| MAGIC-REACH-181 | Ordinary client Fire Ball selects a point even over an actor, and the selected cursor arm has a separate aggregate visibility-field gate. | High / Medium | ● active | [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/measurements.json`, `verification.md`) |
| MAGIC-CONSUME-142 | Shipped non-spell Potions have three effects on state, not one universal timer. | High / Medium | ✔ promoted | [EXP-0275](../experiments/EXP-0275-consumable-use/) |
| MAGIC-CONSUME-143 | Timed potion identity is zero, shared across characteristics, and repeat use preserves the old kind. | High | ✔ promoted | [EXP-0275](../experiments/EXP-0275-consumable-use/) |
| MAGIC-CONSUME-144 | Timed potion expiry and save/load preserve the existing attachment, not a fresh potion duration. | High / Medium | ✔ promoted | [EXP-0275](../experiments/EXP-0275-consumable-use/) |
| MAGIC-SPRAY-134 | On the ranked path, Prismatic Spray's capped argument is the final victim-list count and the primary consumes one place; an empty candidate list returns the primary without reading the cap. | High | ● active | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |
| MAGIC-SPRAY-135 | Book and caster-item Prismatic Spray converge on the same selector; a fighter rider does not, and shipped script instants contain no Prismatic Spray. | High / Medium / Unknown | ● active (amended) | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |
| MAGIC-SPRAY-136 | The selected victim order is the visible Prismatic Spray branch order. | High | ● active | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |
| MAGIC-SPRAY-137 | The selector's secondary rank is deterministic, uses unchecked fixed pointer and winner storage, and is independent of cast range. | High | ● active | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |
| MAGIC-SPELL-001 | A spell is a 0x14-byte handle onto a `Data.bin` Spells row, and the id space is 1..28 with the game's own names. | High | ● active | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-BOOK-002 | The spellbook is a sparse array subscripted by spell id, and "knowing" a spell is a non-null slot. | High / Unknown | ● active | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-CAST-003 | `R0268` is the cast, and the whole cost is one flat column subtracted once. | High | ● active (partially retracted) | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-POWER-004 | What a magic school skill buys — `power = clamp(skill[school] + Mind − 30, 0, 100)`, and it is not the weapon skill's term. | High | ● active (partially retracted) | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-DMG-005 | A damage spell becomes an `Effect_DirectDamage` whose elemental kind is its own school, and the second damage byte is a spread. | High | ● active | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-RESIST-006 | A resistance is a straight percentage, it applies only to the elemental component, and a spell never rolls to hit. | High / Unknown | ● active (amended) | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-ITEM-007 | `castSpell` is a dead arm of the effect dispatch because a castSpell effect is never applied — it is read as data. | High | ● active (amended, partially retracted) | [EXP-0065](../experiments/EXP-0065-magic/), [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-SHAPE-008 | Two `Data.bin` columns decide a spell's shape, and they are read at four sites in the apply's tail. | High / Unknown | ● active (partially retracted) | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-EFFMODE-009 | The four gate bits of `effect+0x3d` are the game's own five duration words. | High | ● active | [EXP-0065](../experiments/EXP-0065-magic/) |
| MAGIC-MIND-010 | Everything Mind does to magic: the power term, the Prismatic victim cap through the one unclamped helper, the AI's decision to cast at all, and one spell arm that reads the *target's* Mind. Nothing else. | High / Unknown | ● active | [EXP-0066](../experiments/EXP-0066-mind-spirit/) |
| MAGIC-SPIRIT-011 | Spirit reaches magic through exactly two derived fields and one spell arm — and never directly. | High / Unknown | ● active | [EXP-0066](../experiments/EXP-0066-mind-spirit/) |
| MAGIC-AI-012 | What decides whether a monster casts — and Mind above 59 makes it cast *less*. | High / Medium | ● active | [EXP-0066](../experiments/EXP-0066-mind-spirit/), [EXP-0387](../experiments/EXP-0387-creature-spellbook/EXP-0387.md) |
| MAGIC-221 | The spell a creature picked — from either the per-slot draw or the Mind-gated walk — reaches a cast order through `R0209`'s own 28-entry jump table, which sends 16 ids to the shared writer `R0018`, 10 to a kind-9 order ... | High / Medium | ● active | [EXP-0387](../experiments/EXP-0387-creature-spellbook/EXP-0387.md) |
| MAGIC-CEIL-013 | The reachable domain of the power, the one consumer that reads it unclamped, and the twelve expressions that consume it. | High / Medium | ● active (amended, superseded) | [EXP-0066](../experiments/EXP-0066-mind-spirit/) |
| MAGIC-ARM-014 | 28 spells are 17 rules, and a spell with damage never reaches its own. | High | ● active | [EXP-0067](../experiments/EXP-0067-spell-catalogue/) |
| MAGIC-EFFECT-015 | What an arm builds: the `Effects` column supplies the kind, the arm supplies the number, and only the first record of the column is ever parsed. | High / Medium | ● active (amended) | [EXP-0067](../experiments/EXP-0067-spell-catalogue/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |
| MAGIC-ATTACH-016 | `actor+0x144` is a bitmask of timed-effect ids (including non-spell Potion id 0), same-id timed effects do not stack, and Bless and Curse annihilate each other. | High | ● active (amended, partially retracted, superseded) | [EXP-0067](../experiments/EXP-0067-spell-catalogue/), [EXP-0179](../experiments/EXP-0179-effect-action/)  · [EXP-0275](../experiments/EXP-0275-consumable-use/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |
| MAGIC-TARGET-017 | Two different columns decide what a spell is aimed at, they disagree on one shipped row, and Prismatic Spray has a separate selector. | High | ● active (partially retracted) | [EXP-0067](../experiments/EXP-0067-spell-catalogue/) |
| MAGIC-TRAIN-018 | Casting trains the spell's own school by half its mana cost, and only for a hero. | High | ● active (partially retracted) | **[EXP-0067](../experiments/EXP-0067-spell-catalogue/)**, **[EXP-0135](../experiments/EXP-0135-skill-raise/)**, **[EXP-0191](../experiments/EXP-0191-itemcast-training/)** |
| MAGIC-SING-019 | Eight spells do something no other spell does, and each is a named instruction rather than a shape. | High / Medium | ● active (amended, partially retracted) | [EXP-0067](../experiments/EXP-0067-spell-catalogue/), [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-AUTOCAST-020 | A weapon-borne spell has two triggers, not one, and they are exact complements — the caster's attack is *replaced* by the cast while the fighter's carries it as a rider. | High | ● active | [EXP-0137](../experiments/EXP-0137-staff-autocast/) |
| MAGIC-AUTOCAST-021 | A caster attacking with a spell-carrying weapon cannot miss and cannot be absorbed, and the reason is structural: the routine that rolls is never entered. | High | ● active | [EXP-0137](../experiments/EXP-0137-staff-autocast/) |
| MAGIC-STAFF-022 | The character generator hands a caster a staff and only a staff, and that is a fact about distribution — nothing in the image refuses an item on account of class. | High | ● active | [EXP-0137](../experiments/EXP-0137-staff-autocast/) |
| MAGIC-SPELLHOP-023 | `MAGIC-ITEM-007`'s open hop is dead code, and the question it framed has a false premise: nothing in the image resolves a weapon's spell by NAME. | High | ● active | [EXP-0137](../experiments/EXP-0137-staff-autocast/) |
| MAGIC-ICON-024 | A spell's icon is a POSITION in one pre-composited strip, not an indexed cell — and four of the twenty-eight spells have none. | High | ● active | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| MAGIC-ICON-025 | The panel's other per-slot decisions, and the two shipped text tables cut 24 and 28. | Medium | ● active | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| MAGIC-PIC-026 | The picture a cast puts on the map is COMPUTED, not looked up: `2*spellId + 8`, or `+9` for the burst variant — and seven spells compute an id no art defines. | High / Medium | ● active | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| MAGIC-PIC-027 | The wire, and the parity fork: an EVEN picture id puts nothing on the map. | High / Unknown | ● active (amended, superseded) | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |

### MAGIC-REACH-178

**Fire Ball's calculated range and the order's range are separate stored values.** Both preserved definitions have ID2, Sphere1 and Max Range10. Constructor resolver `R0624` stores parameter6's low byte to Spell+9 at `L03040`. `R0904` clamps signed skill[Sphere]+Mind-30 to 0..100; `R0625` reloads parameter6 and adds integer power/30 when its byte is nonzero, producing10..13 for Fire Ball. Player handler stores `L05036/L05037` and `L05038/L05039`, and unit constructor `L05040/L00075`, copy Spell+9 into order+0x14, replacing weapon reach. Nineteen original producer and 12 terminal-copy executions distinguish clamp/division boundaries, zero/changed base and copied ranges. These are snapshots: the selected player handlers initially set parent state0x0d/0x0e and order kind0.

**Confidence.** High for the two-root scalar data and named conditional producer/copy operations. Unknown whether current power is refreshed before every order decision, and the intervening parent state to child kind8/9 path.

**Original status.** ● active

**Evidence.** [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/measurements.json`, `verification.md`)

### MAGIC-REACH-179

**Unit-target cast order8 has a facing gate followed by a footprint-gap comparison, with a self-target bypass.** `L00108` calls `R0041` unless caster equals target at `L05041`. The predicate requires mover+0 to equal `R0051`'s direction, then compares the low byte from `R0036` unsigned <= order+0x14. For each axis, c=u16(((size+2*cell+511)<<7)+fraction); gap=max(abs(cCaster-cTarget)-(sizeCaster+sizeTarget)*128,0); distance=1+floor(max(gapX,gapY)/256). Cell/fraction come from Position+0/+1 and+4/+5, and typed virtual+0x1c reaches `R0256`, actor+0x49 tokenSize. Success installs action0x0d/Spell/target; refusal reaches approach `R0042`. Original fragments discriminate maximum-axis from Euclidean/Manhattan distance, both footprint contributions, wrong facing and self at range0.

**Confidence.** High for these typed conditional branches, arithmetic and rejecting mutation. Medium for gameplay inference: execution begins after parent gates, stops at approach or success stores, and does not witness movement or release. Other virtual receiver classes remain Unknown.

**Original status.** ● active

**Evidence.** [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/orders.tsv`, `evidence/instructions.tsv`)

### MAGIC-REACH-180

**Point-target cast order9 uses whole-cell maximum-axis distance, and neither selected predicate directly reads altitude.** `L05042` calls `R0086`; it requires mover+0 to equal `R0089`'s direction from current Position+0/+1, then measures max(abs(dx),abs(dy)) from Position+2/+3 to the packed point. Size and fraction do not enter this distance. Success installs action0x0e/Spell/cell; refusal reaches approach `R0087`. The 388 unit/point axis/diagonal cases at range10 admit10 and refuse11 for centred1x1 actors with matched facing, at caster height64 and distinct target heights32/64/96. Coincident cells have one height. Both complete typed predicates and heading helpers produce zero world reads; range11/13 controls move the cutoff accordingly.

**Confidence.** High for the local instruction contract, finite original-fragment sweep and JA-to-JBE controls. The altitude negative is confined to these predicates and selected child-order arms. Parent AI-state, upstream sight, native release, and naturally reachable non-centre/stored-coordinate diagnostic states remain Unknown.

**Original status.** ● active

**Evidence.** [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/`, `EXP-0369.md`)

### MAGIC-REACH-181

**Ordinary client Fire Ball selects a point even over an actor, and the selected cursor arm has a separate aggregate visibility-field gate.** Ordinary selector index1 reads 0 from `L00200+4` through `R0094`; click adds 1 and command remap `L00198+8` remains 2. The `R0211` click arm reaches point builder `R0090` with no hit, a non-actor hit or an actor-class hit; changing only the target-type table cell to 1 exposes unit builder `R0091`. The ordinary point opcode0x1f remains the existing AI-CLICK-050/AI-CMD-032 contract. `R0218` sets capability0x400 when four tile words OR/mask differently from0xc000; selected-spell cursor arm `L05043..L01421` then changes cast to move. An aggregate0xc000 passes, including mixed0x8000/0x4000 corners; no individual-visible-corner equivalence is claimed. Neither fragment compares caster distance.

**Confidence.** High for the pinned table identity, conditional original click/cursor execution and mutation. Medium for whole client/gameplay inference: prior cursor selection, later modifiers, builder/network reception, altitude-to-visibility updates and projected screen-plane production are not executed. No global absence of a visibility gate follows.

**Original status.** ● active

**Evidence.** [EXP-0369](../experiments/EXP-0369-cast-reach/) (`evidence/measurements.json`, `verification.md`)

### MAGIC-CONSUME-142

**Shipped non-spell Potions have three effects on state, not one universal timer.** Both roots contain four mode 0 +1 attribute potions (Body/Reaction/Mind/Spirit), four mode 0 current-health/current-mana potions (+30/+100), and five mode 1 temporary rows. `R0935` converts mode 0 to 8 before `R0612`; mode 8 applies once and stores no attachment. Attribute cases2..5 in `R0843` add to live words +84/+86/+88/+8a and skip modifier bytes +d4..d7 when bit 8 is set; the effect arm first caps at 100, then its common tail `L04870..L06977` invokes actor vt+50, the Human `R0280` derivation that caps live values at `50 + signed modifier`. A potion at the effective cap can therefore be consumed without raising the stat. Health/mana cases6/9 clamp to maxima. Timed rows are Antipoison (actual kind 16 absorption, +50 for 480 ticks), Health/Mana Regeneration (kind 8/11, +100 for 960), and Fighter/Mage Bonus (kind 8/11, +250 for 1920). They modify regeneration percentages or absorption, not per-tick healing directly. Mana-regeneration kind 11 has a mage-bit gate in the apply consumer, not a Potion input refusal.

**Confidence.** High for connected original mode/apply/derive branches; Medium for matching two-root authored population.

**Original status.** ✔ promoted

**Evidence.** [EXP-0275](../experiments/EXP-0275-consumable-use/)

### MAGIC-CONSUME-143

**Timed potion identity is zero, shared across characteristics, and repeat use preserves the old kind.** `R1049` initializes Effect Token+0c=0; `R1000` sets kind+3c, mode+3d and operands+40/+42 but not that id, and `R0891 → R0992 → R0999` is the authored Potion path. `R0612` considers only mode bits 1/2 timed, then searches actor+20 for equal Token+0c. Missing id copies/applies/appends the effect. Present id with incoming continuous bit 2 refreshes only old duration; otherwise it un-applies old, copies only new low magnitude and duration, and re-applies old. Kind, mode and id are not replaced. Thus Health Regeneration followed by Mana Regeneration before expiry keeps health regeneration, and reverse order keeps mana regeneration; Fighter Bonus then Health Regeneration lowers the existing health bonus to 100 and resets duration to 960. All five shipped timed potion rows collide at id 0 and set mask bit 0; no per-kind stack or magnitude sum exists on this path.

**Confidence.** High; constructor/parser identity and exact old-object stores refute per-kind refresh. Runtime sequence remains an unwitnessed, falsifiable prediction.

**Original status.** ✔ promoted

**Evidence.** [EXP-0275](../experiments/EXP-0275-consumable-use/)

### MAGIC-CONSUME-144

**Timed potion expiry and save/load preserve the existing attachment, not a fresh potion duration.** Actor tick `R0037`, except terminal act `0x10`, iterates actor+20 and invokes Effect vt+38=`R0673` before its health branch. Duration/continuous are bits 1/2; charges bit 4 alone is not timed. Continuous reapplies when remaining counter is divisible by 8; counters >9600 skip decrement. Others decrement; on zero, non-continuous un-applies, all clear the id bit and set mode bit 128, which the actor loop removes. Unit Serialize `R0210` passes actor+20 to `R0951`, actor+d4 to the 0x40-byte block serializer `R1570`, and stores live attributes, health/mana and actor+144. Effect `L08329` saves/loads kind, mode, packed magnitude/remaining duration and Token id after the Token head; no attach/restart runs in those readers. Human `L08212 → R0954` delegates Unit. The static record thus carries remaining duration plus already-applied modifiers and live persistent attribute changes; only transient Effect+44 is omitted. The 27 existing first-Effect save witnesses include 18 kind 8/mode 1/id 0 records: 17 hold duration 960 and EN game0007 holds776. Nine others are kind 41/mode 0/id 0. These first-class-body measurements do not identify the attachment owner or prove a newly executed drink-save-reload sequence.

**Confidence.** High for writer/reader and expiry; Medium for archived witnesses; runtime reload and town-to-mission transfer effects Unknown.

**Original status.** ✔ promoted

**Evidence.** [EXP-0275](../experiments/EXP-0275-consumable-use/)

### MAGIC-SPRAY-134

**On the ranked path, Prismatic Spray's capped argument is the final victim-list count and the primary consumes one place; an empty candidate list returns the primary without reading the cap.** `R0269` passes `(u8)min(power/20+2,7)` as the fourth stack argument to `R0109(session,caster,primary,out,cap)`. The selector clears `out`, selects secondaries into ten fixed winner slots, then appends `primary` first, appends winners in selection order while skipping a winner equal to `primary`, and removes the tail while `out.count > cap`. The cap is neither a radius nor an attempt budget: each of `cap` winner scans examines every admitted candidate before the final list is truncated. With nonempty A, a zero cap removes even the primary; with empty A, the early return preserves it. The ordinary authored book/item domain supplies 2..7, but the item expression has no lower clamp and is narrowed to a byte before the call.

**Confidence.** High. The complete caller and selector fix the argument order, clear, empty-A early return, fixed winner storage, append order, duplicate test and ranked-path tail removal. `selector-model.csv` is a fixed illustration of those instructions, not separate evidence.

**Original status.** ● active

**Evidence.** [EXP-0243](../experiments/EXP-0243-prismatic-victims/)

### MAGIC-SPRAY-135

**Book and caster-item Prismatic Spray converge on the same selector; a fighter rider does not, and shipped script instants contain no Prismatic Spray.** A book cast derives power from the Sphere-2 actor expression. With `caster+0x68 != 0`, `R0269` instead finds the item's kind-`0x29` effect and uses signed `effect+0x42`; both divide by 20, add 2 and clamp only above at 7. Actor states `0x0d` and `0x0e` both reach `R0268`; its id-14 arm calls this selector directly. A caster item's admission path reaches that arm while item context is live, after which `R0002` refuses id 14 to prevent a duplicate; a fighter rider reaches only that refusal. Script instant 24 structurally creates state `0x0d` with a unit target and would converge; instant 21 creates state `0x0e` with a null target, while the selector dereferences the primary before any null gate. Both preserved campaign roots contain zero spell-14 instances among 35 instant-21 and 20 instant-24 records. `Spell+0x09`, derived separately from Max Range plus the power bonus, is copied into order `+0x14` and governs cast approach/admission; it is not the selector cap and does not bound the secondary group-sight population.

**Confidence.** High for the routes, formulas and separate order-range field, each asserted as executable bytes on both roots and discriminated by the item/book branch. Medium for the dual-root shipped-script absence. The instant-21 null-target fault is a static prediction, not a witnessed result; temporary-caster group ownership is narrowed by `MAGIC-235` to Medium for casters built by the wrappers (no assignment found in the named population) and stays Unknown for a caster created any other way, and for the link between those wrappers and the script instants.

**Original status.** ● active

**Evidence.** [EXP-0243](../experiments/EXP-0243-prismatic-victims/)

**Amended.** `MAGIC-235` narrows the Unknown on temporary-caster group ownership: a caster built by the `R0914` wrappers has no owner and no group at construction, and no assignment was found in the named population (Medium). The link from those wrappers to the script instants above is by state value `0x0d` or `0x0e` only and remains Unknown. The route, formulas and shipped-script absence above are unchanged.

### MAGIC-SPRAY-136

**The selected victim order is the visible Prismatic Spray branch order.** `R0269` sends opcode `0x8a` for an ordinary caster or `0x8c` for a temporary script caster, with message count `selectedCount+1`, source identity in word 0 and the selected runtime victim ids in words 1 onward. The client `0x8a` arm resolves the caster; the `0x8c` arm creates picture 36 at the packed source cell. Both clear and refill the source object's word array at `+0xac`, set `+0xb0` to the selected count and set picture 36's primary field. The projectile spawner copies that array, and `MAGIC-BOLTLIST-071` draws one generated link per id in its stored order. The old proposed `L05035` `SetSize` site does not fill this array.

**Confidence.** High. Producer and both client consumers are complete instruction streams, the client dispatch is read through the next opcode boundary, and the raw anchors agree on both preserved roots. `MAGIC-BOLTLIST-071` independently fixes the downstream one-link-per-id consumer.

**Original status.** ● active

**Evidence.** [EXP-0243](../experiments/EXP-0243-prismatic-victims/)

### MAGIC-SPRAY-137

**The selector's secondary rank is deterministic, uses unchecked fixed pointer and winner storage, and is independent of cast range.** The pointer array at `stack+0x60` has room for 100 group candidates, but both source-list loops lack a bound check. Their parallel scores go to `[session+0xd74+4*i]`; this round did not establish that scratch region's capacity. Ten `u32` winner-index slots and ten `u16` score-threshold slots are fixed in the stack frame, and a custom cap above ten is not guarded. List A score is `((edgeDistance<<8)+turnCost)&0xffff`; list B score is `((((edgeDistance<<8)+turnCost)&0xffff)<<8)`. Each winner scan uses strict `<`, so equal stored scores retain candidate-source order. Both winner indices and thresholds begin at 65530. A winner index below 60000 gates only the write of 65500 into that candidate's score scratch; append separately requires both winner index and saved threshold below 65000. Ordinary caps stop at 7. Above 100 candidates the pointer writes leave their 100-entry stack region; whether the score write leaves its owning scratch region is Unknown. A cap above ten leaves the fixed winner/threshold storage.

**Confidence.** High for the pointer and winner capacities, both score formulas, strict tie rule and sentinel control flow: each is an instruction-level read asserted on both roots. The score-scratch capacity and the concrete result of either overflow remain Unknown.

**Original status.** ● active

**Evidence.** [EXP-0243](../experiments/EXP-0243-prismatic-victims/)

### MAGIC-SPELL-001

**A spell is a 0x14-byte handle onto a `Data.bin` Spells row, and the id space is 1..28 with the game's own names.** `R0463(id)` and `R1040(name)` construct it (vtable `L05044`); `R0624` resolves — `this+0x04 = Spells[id]` (`L05045`, the collection at `L05046`), `this+0x0c` (u16) `= getParam(row, 1)` = title 2 **`Mana Cost`** (`L05047`), `this+0x09 = getParam(row, 6)` = title 7 `Max Range`, `this+0x0a = (getParam(row, 0x12) == 1)` = title 19 `Defensive`. Id 0 is refused with `"Invalid spell #0 - can't cast."`, an unknown name with `"Invalind spell <s> - no such ID"`. `Spell::Serialize` (`R1013`) stores exactly `+0x08`/`+0x09`/`+0x0a` as bytes and `+0x0c` as a word, restoring `+0x04` from the id, so **`+0x0e`, `+0x0f`, `+0x10` are per-cast scratch** (`MAGIC-POWER-004`). Corpus: `evidence/spell-names.txt` reads the 29-entry name array at `L05034` out of `.data` by virtual address and matches it against the shipped Spells rows — **28/28 by name**, entry 0 `unused_spell_0` having no row, which is also the id the resolver refuses

**Confidence.** High (every field is a named instruction with the column index as its own argument; the id↔name identity is a corpus match on the image's own array, and the four `Protection from …` rows sharing one `switch` arm re-derive it from the code)

**Original status.** ● active

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

### MAGIC-BOOK-002

*(amended by EXP-0066: **nothing about the book is gated on a stat.** `EnumRefs disp:88` and `disp:8a`, whole image, place no hit in `R0464`, `R0017`, `R1041` or the `teachSpell` arm, so neither Mind nor Spirit bounds the capacity, decides what may be learned, or gates learning — `MAGIC-MIND-010`.)* **The spellbook is a sparse array subscripted by spell id, and "knowing" a spell is a non-null slot.** `actor+0x140` holds a `Spellbook` (`0x1c` B) = `[vtable][CObArray 0x14 B at +0x04][int lastFound at +0x18]`. `R0464(book, id, spell)` grows the array to `id+1` with `-1` fill when it is short (`L05048`), deletes any existing occupant, then `array[id] = spell` (`L05049`). `R0017(book, id)` returns 0 when `id >= GetSize()` or the slot is null, else stamps `book+0x18 = id` and returns it. The destructor `R1041` deletes every non-null element. **Learning is only the `teachSpell` effect arm** (`HERO-EFFECT-019` kind 42, `L05050`): no book → nothing; `Get(id) != 0` → nothing; else `new Spell(id)` (`0x14` bytes at `L05051`) and `SetAt`, then `actor+0x150 \|= 0x400000`. `R0842` allocates the book on demand and sets `actor+0x4c` bit 1, plus **bit 2 — `HERO-CLASS-013`'s mage bit — when `manaMax != 0`** (`L05052`). **Unknown:** `EnumRefs callto:R0842` on the repaired table returns **0 hits / 0 owners / 0 orphan**, so no enumerated route reaches the allocator and how an actor first acquires a book is not established; 32 of the 33 `SetAt` sites are in the spawn `R0184` and the human constructor `R0656`

**Confidence.** High (the three container routines are read whole and the sizes are their own allocation arguments) / **Unknown** (the allocator's reachability, with the instrument stated)

**Original status.** ● active

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

### MAGIC-CAST-003

**Correction:** The scheduling clause is partially retracted by MAGIC-CASTCLOCK-171: the delay is a presentation-message field, not a queued simulation apply; the mana arithmetic stands.  **`R0268` is the cast, and the whole cost is one flat column subtracted once.** Eleven instructions decide the economy: `if (!R0442(caster))` skip (`L01998` — `HERO-CLASS-020` fixes that predicate as the **mage** test), `if (caster+0x68 != 0)` skip (`L05053` — a cast from an item), `if ((i16)spell+0x0c > (i16)caster+0x9a) return 0` (`L05054` — refused, and *nothing else happens*), else `caster+0x9a -= spell+0x0c` (`L05055`). **No skill, stat, level or spell-power term enters**: there is no multiply between `L05056` and `L05057`. So a **fighter is never charged**, an **item cast is never charged**, and the shipped costs are the raw `Mana Cost` column, 3…100 with `Fire Sacrifice` at 0 (`evidence/spell-table.txt`). Three further acts of the same routine: casting at anyone but yourself sets any active **invisibility** effect's remaining duration to 1 (`L05058`…`L05059`, spell 15); a `Delivery System == 2` spell is queued with a delay of `distance / SpellEffectSpeed` — a fixed **5** for Lightning and Prismatic Spray (`L05060`, `L05061`); everything else is queued at once, on `L00522`. Instrument: `EnumRefs callto:R0268` — 2 hits, 1 owner (`R0037`, `L01994`/`L01995`), 0 orphan

**Confidence.** High (the routine is read whole; the refusal, the subtraction and both exemptions are the `CMP`/`SUB` operands themselves)

**Original status.** ● partially retracted

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-POWER-004

*(**the "exactly four expressions" clause is RETRACTED by EXP-0066** — the power is read 21 times inside `R0003` alone, in twelve distinct expressions; `MAGIC-CEIL-013` carries the corrected enumeration. Amended by EXP-0066 twice more: `R0253` has **no upper clamp**, so the `clamp(…, 0, 100)` in this headline is `R0003`'s and `R0904`'s and not the Prismatic victim cap's, and the sum can reach **170**, so the clamp binds in play.)* **What a magic school skill buys — `power = clamp(skill[school] + Mind − 30, 0, 100)`, and it is not the weapon skill's term.** `R0253(actor, school)` is `max(skill[school] + Mind − 30, 0)` in six instructions: `L05062` reads the signed 16-bit skill of the school from the array at `+0xa8` of the actor, `L05063` reads the signed 16-bit Mind at `+0x88`, `L05064` forms `skill + Mind − 0x1e`, and a branch-free clamp at zero follows, returned as a byte. `R0003` repeats it at `L05065..L05066` with a `[0,100]` clamp. The school is the spell's own `Sphere` column (`getParam(row, 2)`, `L05067`) and indexes the very array the five weapon skills live in (`HERO-SKILL-009`). The power is consumed in **twelve** distinct expressions, all outside `R0280` (`MAGIC-CEIL-013`), of which four are the shared cast quantities: damage `f = power/30 + 1` (`L05068`, `[L05069] = 30`, `[L04115] = 1`); duration `ftol(1.025^power × durationColumn × 16)` (`L05070`, `L05071`, `[L05072] = 16`); cast range `+ power/30`, or `+ power/3` for Teleport (`L05073`, `L05074`); Prismatic victim cap `min(power/20 + 2, 7)` (`L05075`). **So Mind is worth exactly as much as the school skill, one for one, and there is no `+3 × skill` to-hit term and no `skill/5` damage floor for a spell** — a spell's damage does not pass the to-hit roll at all (`MAGIC-RESIST-006`). Instrument for the absence: `EnumRefs disp:a8`, whole image — a `disp:` sweep, which is the mode a structure field needs and the only one that sees the indexed form `word ptr [reg + i*2 + 0xa8]` these two readers use — **260 hits / 118 distinct owners / 1 in orphan code**, that one being `L05076`, a `dword` store in the unrelated orphan run `L05077..L05078`. Inside the actor family the readers are `R0253`, `R0003`, `R0904`, `R0280` (`HERO-COMBAT-011`'s weapon terms), the four experience routines and `Weapon::Equip`'s ranged arm; blind only to bytes never disassembled. For an **item** cast the power is not this expression at all: it is the item's `castSpell` effect's `+0x42` (`MAGIC-ITEM-007`)

**Confidence.** High (the formula is six named instructions read twice over; the "no symmetric term" clause states its instrument, its mode and its one orphan hit) / **the withdrawn clause carried the same High**, and it was an enumeration argued from the same cell as the formula — the shape `REG-GMAP-066` records

**Original status.** ● partially retracted

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-DMG-005

**A damage spell becomes an `Effect_DirectDamage` whose elemental kind is its own school, and the second damage byte is a spread.** `R0625(spell, power)` fills the scratch pair (`L05079`…`L05080`): `spell+0x0e = ftol(damageMin × f)`, and `spell+0x0f = ftol(damageMax × f) − spell+0x0e` — the subtraction is the `FSUBP` at `L04069`, against the byte just written. `R0003` then allocates `0x60` bytes, runs `R1042`, sets `+0x0c = the spell id` and `+0x44 = the caster`, and dispatches on the school through the jump table at `L05081` into five instruction-for-identical setters differing in one immediate: `R1043/e90/ec0/ef0/f20` write `effect+0x5b = spell+0x0e`, `effect+0x5c = spell+0x0f`, `effect+0x5d = 1/2/3/4/5`. Since `Effect_DirectDamage+0x48` is a `0x16`-byte combat block of the `actor+0xa6` layout (`R1044`, `L05082` offsetting the base by 0x48), those are block `+0x13/+0x14/+0x15` — the three fields `HERO-EFFECT-019`'s `damageFire…damageAstral` arms write. **School *i* becomes damage kind *i*.** The whole arm is skipped when the pair sums to 0 or the spell is 6 (`Heal`) or 11 (`Drain Life`), both of which carry damage columns and are not damage (`L05083`, `L05084`). Corpus: re-executing the arithmetic at power 0 (`f = 1`) makes the rolled range `base … base+spread` reproduce `Data.bin`'s own `damageMin … damageMax` columns **exactly on 9 of 9 damage spells** — 4–8, 7–13, 1–3, 8–16, 10–14, 3–5, 5–15, 5–15, 5–25 (`evidence/spell-table.txt`); under the rival "the second byte is a maximum" reading, 0 of 9

**Confidence.** High (the fill and all five setters are read at instruction level; the corpus test at power 0 discriminates against the only live rival)

**Original status.** ● active

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

### MAGIC-RESIST-006

**A resistance is a straight percentage, it applies only to the elemental component, and a spell never rolls to hit.** In `R0265` (`HERO-DAMAGE-022`) the elemental component is `v = block+0x13 + U[0, block+0x14]`, admitted iff the physical swing landed **or the attacker has no physical damage at all** (`L05085`…`L05086`), then reduced by the protection word of its own kind — `A+0x15 − 1` bounded `0..4` through the jump table at `L04134`, kind 1→`target+0xc4` Fire, 2→`+0xc6` Water, 3→`+0xc8` Air, 4→`+0xca` Earth, 5→`+0xcc` Astral, default printing `"Unknown magic damage type"`. Each arm is the same seven-instruction FPU block: `ftol(v × (100 − p)/100 + 0.75)` — `[L02972] = 100`, `[L05087] = 0.75` — then `max(v, 0)`. **Absorption (`target+0xc0`) is subtracted from the physical component only** (`L04133`), so nothing flat reduces a spell; and because an `Effect_DirectDamage`'s block carries no physical pair, a spell's elemental damage always applies. `HERO-RESIST-012` clamps the protection to `[0,100]`, so **100 is immunity and no value can invert the sign** (`evidence/resistance.txt` re-executes the curve over `[1,1000] × [0,100]`). Consequence in a player's terms, on the sheet `EXP-0063/evidence/owner-testimony.txt` T5 records: Spirit 42 → protection 21 → a 100-point fireball lands for **79**. A second component sits between the two — `block+0x11 + U[0, block+0x12]`, run whenever that pair is nonzero and reduced by `target+0xc6` (protection index 2, `protectionWater`) — and it is **not gated on the hit**: `L05088`, a jump on less-or-equal to `L05089`, jumps a missed swing straight into it. Its filler was not located

**Confidence.** High (seven identical FPU blocks read at instruction level, the jump table resolved, and the flat rival excluded by the presence of the `FDIV`) / **Unknown** (what writes the `block+0x11`/`+0x12` pair)

**Original status.** ● active (amended)

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### MAGIC-ITEM-007

*(amended by EXP-0191: the fighter rider's admission/order and shipped-power clauses are retracted; `MAGIC-ITEMTRAIN-116` carries the full branch, award order, authored corpus and shop producer.)* *(amended by EXP-0191: the item context also suppresses the immediate cast award; later hit, damage, pseudo-damage and kill producers are separate and are enumerated by `MAGIC-ITEMTRAIN-116`.)* *(amended by EXP-0066: the item branch **jumps past the `[0,100]` clamp** — `L05090` reads the signed 16-bit field `effect+0x42` and `L05091` jumps unconditionally to `L05092`, so `effect+0x42` reaches `f = power/30 + 1` as a raw signed 16-bit value and reaches `R0625` truncated to a byte. Shipped items carry 10 and 20, so nothing in the shipped data exercises it; it is the **only** unclamped power in the image that is not `R0253`'s, and it reads no stat.)* *(amended by EXP-0137 on the exclusivity clause and on the Unknown, neither of them a correction of the instructions this row reads. **"only a fighter fires it" is withdrawn** — `L05093` is the fighter predicate and the rider is a fighter's, but this row never read the melee strike's own caller, where the mage takes the same weapon's spell down a second path that replaces the attack (`MAGIC-AUTOCAST-020`, retracted.md). And the Unknown's premise is false: `R1045` is **dead code**, so no hop resolves a spell **name** — the live builder `R1012` takes an id byte (`MAGIC-SPELLHOP-023`))* **`castSpell` is a dead arm of the effect dispatch because a castSpell effect is never applied — it is read as data.** The shipped item name `Wood Staff {castSpell=Fire_Arrow:20}` (`L05094`) becomes an effect of kind `0x29` = 41 on the item's list plus a `Spell` object at `weapon+0x80`. `R0858(item)` walks `item+0x20` for `effect+0x3c == 0x29` (`L05095`) and returns it; both the cast (`L05096`) and the area applier (`L05097`) take **`effect+0x42` as the power** in place of `MAGIC-POWER-004`'s expression, and `MAGIC-CAST-003`'s mana gate exempts the whole path. The trigger is the melee strike `R0246`: after the damage lands, `L05098` requires the weapon to carry a spell and **`L05093` requires the wielder to be a fighter** (`R0859`, `HERO-CLASS-020`'s exact negation of the mage test); the strike then stamps `caster+0x64 = weapon+0x80`, `caster+0x68 = the weapon` and calls `Spell::Apply` (`R0002` → `R0003`), clearing both afterwards. A dead target still triggers it when the weapon's spell is id 2, `Fire Ball` (`L05099`). ~~**Unknown:** `R1045`, which builds the `Spell` into `weapon+0x80` from a name, has **0 hits / 0 owners / 0 orphan** on `EnumRefs callto:R1045`, so the hop that parses the `{castSpell=…:n}` syntax is not located~~ — closed by `MAGIC-SPELLHOP-023`

**Confidence.** High (each hop is a named instruction) for what the strike itself does; the **exclusivity** clause is withdrawn at the High it carried (retracted.md)

**Original status.** ● active (amended; exclusivity, fighter-rider admission/order and shipped-power clauses retracted)

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/), [EXP-0191](../experiments/EXP-0191-itemcast-training/)

**Amended.** The table ledger carried the status "● active (amended; exclusivity, fighter-rider admission/order and shipped-power clauses retracted)". `retracted.md` records a correction against this claim.

### MAGIC-SHAPE-008

*(**the `effect+0x4c` clause is RETRACTED by EXP-0067** — the first term is the `Area Effect Duaration` column, `getParam(row, 0x0b) << 4` at `L05100`…`L05101`, not the distribution value, and the corrected expression is in the headline below. `MAGIC-CEIL-013` quoted the same error and was corrected with it.)* **Two `Data.bin` columns decide a spell's shape, and they are read at four sites in the apply's tail.** `getParam(row, 8)` — title 9 `Distribution system` — at `L05102`: value **1** builds a `PointEffect` (`R1046`) and **requires a target unit**, printing `"Spell, oops - can't cast point effect of x,y"` when there is none; anything else builds an `AreaEffect` (`R0652`) sized by `getParam(row, 9)` = title 10 `Radius, Length/2` (`L05103`) and `getParam(row, 0x0b)` = title 12 `Area Effect Duaration` (`L05100`), with `effect+0x4c = (AreaEffectDuaration << 4) + (power << 4)/10` — the column, shifted left by 4 at `L05104`, into `+0x4c` at `L05101`, and the power term added at `L05105` only when that is nonzero, which also sets `effect+0x08 = 1` — and a second arm for distribution **5** (`L05106`), which instead writes `+0x08 = 2` and `+0x4c = 0`. `getParam(row, 5)` — title 6 `Delivery System` — at `L05107`/`L05108`: **1** attaches the effect immediately, **2** wraps it in a `SpellTransport` (`R0633`) whose flight uses `getParam(row, 7)` = title 8 `Spell Effect Speed` (`L05109`), the same column `MAGIC-CAST-003`'s delay divides by. Corpus, EN `Data.bin`: `Delivery System` is 2 on exactly the four missile spells `Fire Arrow`, `Fire Ball`, `Lightning`, `Prismatic Spray` and 1 on the other 24; `Distribution system` takes `{1, 3, 4, 5}` over 18/5/2/3 rows and never 2. **Unknown:** the five words `Point` / `Round` / `Long` / `Phase` / `Hang On Unit` at `L05110` are the obvious enum for that column and the shipped values do not fall in an order that reading explains; the mapping is left unread

**Confidence.** High (the four call sites and their column indices are the call arguments themselves and the two constructors are named) / **Unknown** (the value→word mapping) / **the withdrawn `+0x4c` clause carried the same High**, and it was a fifth statement — about what is *done* with one of those values — argued from the cell that names the call sites

**Original status.** ● partially retracted

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-EFFMODE-009

**The four gate bits of `effect+0x3d` are the game's own five duration words.** `R0857(word, &out)` is the only referrer of the literals at `L05111` (`refto:` 1 hit each, 1 owner) and is called once, by `R1000` at `L05112`, which is what `Effect::Effect(magicRow)` (`R1047`) calls to build its template. It returns `permanent → 0`, `singleuse → 8`, `charges → 4`, `duration → 1`, `continuous → 2`, the last three also parsing a trailing number into `*out`. Those are exactly the four `.data` bytes `R0843` tests: `[L04012] = 1`, `[L04013] = 2`, `[L04014] = 4`, `[L03991] = 8` (`evidence/spell-names.txt` reads all four out of the image). So `HERO-EFFECT-019`'s unnamed magnitude gate is **"is this effect timed"** — a timed effect keeps a 16-bit magnitude at `+0x40` and its remaining ticks at `+0x42`, an untimed one a 32-bit magnitude at `+0x40` — and `HERO-CAP-015`'s "unless `effect+0x3d & [L03991]`" is **"unless it is single-use"**. Every lasting effect the apply builds ORs in bit 0 (`L05113` and siblings) and writes its ticks to `+0x42`

**Confidence.** High (one parser, one call site, four constants read out of the image, and the mapping is the parser's own return values beside its own literals)

**Original status.** ● active

**Evidence.** [EXP-0065](../experiments/EXP-0065-magic/)

### MAGIC-MIND-010

**Everything Mind does to magic: the power term, the Prismatic victim cap through the one unclamped helper, the AI's decision to cast at all, and one spell arm that reads the *target's* Mind. Nothing else.** Instrument: `EnumRefs disp:88`, whole image, on the repaired table — **461 hits / 235 distinct owners / 4 in orphan code**, those four being 32-bit stores to the stack slot `+0x88` at `L05114`, `L05115`, `L05116`, `L05117` (stack frames, named rather than dropped). The field is a `u16`, so the 461 were reduced by a stated filter and both halves of it were read: **every hit of byte or 16-bit width** (31 — the complete non-dword population, classified one by one in `evidence/readers.csv`) and **every dword hit not based on the stack pointer inside the address ranges where actor methods live** (`L05118`, `L05119`, `L05120`, `L05121`, `L01769`, plus the `L05122..L05123` `CUnit` block), none of which is on an actor pointer — a dword at `+0x88` would span Mind *and* Spirit. There is **no indexed form**: grepping the sweep for `*0x` returns nothing, and `disp:84`/`disp:86` carry no `[reg + i*2 + 0x84]` loop, so the four stats are never treated as an array. On the spell path the readers are exactly four: `R0253` (`L05063`), `R0003` (`L05124`), `R0904` (`L05125`) — the three `power` sites — and `R0003` again at `L05126`, the **`Control Spirit` arm**, which copies the *victim's* `+0x88` **verbatim** into the `"Ghost"` it raises (`L05127`) while halving its Reaction, health and healthMax; the arm is `[L05128, L05129)`, bounded on both sides off the image's own 29-entry jump table at `L03242` (`evidence/apply-arms.txt`). Off the spell path: `R0280` (the cap, sight), `R0810` (`HERO-XP-010`), `R0843` (as the target of effect kind 3), two serialisers, five default-stat writers, and two AI routines (`MAGIC-AI-012`). **Absences, over the same enumeration:** the cast `R0268`, the fill `R0625`, the hit resolver `R0265`, the regeneration tick `R0654` and all three spellbook routines contain **no** hit — so Mind does not touch admission, the mana gate, the cost, the delay, the to-hit question, a target's resistance, regeneration, the book's capacity, what may be learned, or whether learning is gated

**Confidence.** High (the four spell-path readers are named instructions; the classification filter, its two halves and the instrument's mode are all stated, and the arm attribution is bounded on both sides by the image's own table rather than by listing order) / **Unknown** (a wholesale block copy (string move or `memcpy`) of an actor would copy both stats carrying **no displacement at all** and no `disp:` sweep can see it; the field-by-field shape of the `Control Spirit` copy is evidence against one existing on this path, not proof)

**Original status.** ● active

**Evidence.** [EXP-0066](../experiments/EXP-0066-mind-spirit/)

### MAGIC-SPIRIT-011

**Spirit reaches magic through exactly two derived fields and one spell arm — and never directly.** Instrument: `EnumRefs disp:8a`, whole image — **36 hits / 13 distinct owners / 0 in orphan or undisassembled code**, small enough that every hit was read; the table is `evidence/readers.csv`. Of the 13 owners one is not the actor at all (`R0043`, four hits, writing `[actor+0x154] + 0x8a` — the mover). Inside the actor family: `R0280` holds **13 of the 36** — the cap (`L04190`/`L04191`/`L04192`), `manaMax` twice (`L04193`, a doubling, and `L04194` into `R0874`'s `pow(1.1, spirit)`, `HERO-MP-006`), `prot[i] = spirit/2` (`L04195`, one loop body `i = 1..5`) and the eight reads of `spirit/2 + 70` that make up the clamp loop `L05130..L03966` (`HERO-RESIST-012`; MSVC evaluates each `min` arm twice and the body carries two nested `min`s). The remainder are chargen, the two serialisers, five default-stat writers, and `R0843` as the target of effect kind 4. **On the spell path there is exactly one hit and it is not a formula:** `L05131`/`L05132`, the `Control Spirit` arm copying the victim's Spirit verbatim into the `"Ghost"` (`MAGIC-MIND-010`). So a caster's Spirit buys **mana** and a target's Spirit buys **resistance**, each in one hop through `R0280`, and nothing else in magic reads `+0x8a`: not the mana gate (`R0268` reads `caster+0x9a`, the pool), not regeneration (`R0654`'s mana term is `manaMax × (actor+0xe2 + 100) × rate / actor+0x9e`, `HERO-REGEN-021`), not the power, not damage, not duration, not the resistance arithmetic itself (`R0265` reads the derived `+0xc4…+0xcc`)

**Confidence.** High (the enumeration is complete at 36 hits and every one was read; the two hops are `HERO-MP-006`'s and `HERO-RESIST-012`'s own instructions and the absence carries the same instrument as `MAGIC-MIND-010`) / **Unknown** (the same `memcpy` blind spot)

**Original status.** ● active

**Evidence.** [EXP-0066](../experiments/EXP-0066-mind-spirit/)

### MAGIC-AI-012

**What decides whether a monster casts — and Mind above 59 makes it cast *less*.** `R0393(this, actor, target)` is the routine `claims/magic.md`'s open item 1 named. Its **only** caller is `R0009` at `L05133` (`EnumRefs callto:R0393`: 1 hit / 1 owner / 0 orphan), reached only when `actor+0x4c & 4` — the mage bit (`HERO-CLASS-020`) — and `[actor+0x14]+0x28 != 0` (`L01740`, `L01774`), so **a fighter never reaches it**. The body: `if ((i16)actor+0x88 > 59)` (`L01768` compares the 16-bit field `actor+0x88` with 0x3b) **and** `rand()*100 / [this+0] < 30` (`L05134`…`L05135`), call `R0009` again and **return 0 — no spell is ordered this pass**; otherwise walk ids `1..0x1c` through `Spellbook::Get(actor+0x140, id)` (`L05136`, bound `L05137`: the loop index compared with 0x1d), keep a spell iff `spell+0x0c` (Mana Cost) `<= actor+0x9a` (current mana, `L05138`) **and** `spell+0x0a == 0` (`Defensive`, `L05139`), then pick `k = rand()*count / [this+0]` uniformly (`L05140`), set `[actor+0x158]+0x60 = 1` and call `R0018`. `[this+0]` is **`0x8000`**, written by the class's own base constructor at `L00371` immediately after `srand(time(0))`, and `R0179` is the CRT `rand` (`seed = seed×0x343fd + 0x269ec3`, `return (seed>>16) & 0x7fff`) — so the first roll is exactly **30 %** and the second is exactly uniform. Consequences a consumer must reproduce: **the AI never casts a `Defensive` spell**, never picks by school, skill or power, and its only stat input is the Mind gate. A second AI reader exists and is not the spell path: `R0225` (3 callers / 3 owners) scales a score by 3/2 in one mode and 3/4 in another when `A->vt+0x1c() < 2` and `A->Mind >= 15` (`L01766`, `L01767`)

**Confidence.** High (the routine is read whole; the two roll denominators are the constructor's own immediate and the CRT `rand`'s own mask, so the 30 % is arithmetic rather than an inference from the constant 0x1e) / **Medium** on the second prior Unknown, **High** on the first — *narrowed by [EXP-0387-correction]: `[actor+0x14]+0x28` is `Player+0x28` at High. Reading its zero value as "not human" is Medium at most: `UNIT-OWNER-009` is partially retracted on exactly this point and no longer supports reading zero as proof of human authorship on an arbitrary map. `R0009`'s own `L01738`…`L01764` (the per-slot spellbook draw, `AI-341`) runs **before** the call into this routine and only falls through to it when that loop already found no match this tick (`L05141`, a zero test with a jump on zero); so reaching `R0393` at all already means the slot draw failed this call. The hand-back branch (`L05142`…`L05143`, Mind > 59 and roll < 30 %) re-runs `R0009` in full, not merely a bookkeeping re-entry: the call at `L01992` to R0009 passes the same actor and target, and reaching `R0393` at all already required owner `+0x28 != 0`, so that unchanged value keeps the re-entered call's own `L05144`…`L05145` prologue from diverting early — the loop at `L01761` runs again with one to three fresh `rand()` calls, and if it fails again the re-entered call can itself reach `R0393` a second time, so the path can recurse. `R0393` itself then returns 0 unconditionally on this branch, so the outer `R0009` call's own default write (`L01735`…`L00159`) fires regardless of what the re-entered call wrote, overwriting `ord+0x08` (to `5`), `ord+0x0c` (the candidate target) and `ord+0x14` (the actor's own weapon-reach byte); a re-entered call's own `ord+0x28`, `+0x30`, `+0x3c` and `ord+0x60 = 1` are not among those three fields and survive. The other branch — Mind ≤ 59, or Mind > 59 with the roll ≥ 30 %, both reaching `L05146` — is a second, independent cast attempt, not a fallback: it walks the actor's full spellbook (ids `1..0x1c`, not the three slots) and, given a non-empty affordable-and-non-Defensive candidate list, calls `R0018` directly (`L05147`) and returns 1, issuing a cast order outside the `ord+0x78..0x8c` mechanism entirely. An empty candidate list (`L05148`, a jump on zero to `L05149`, list head null) skips the pick and instead tests whatever `Spellbook::Get` last returned inside the `1..0x1c` filter loop — its own final iteration, id `28` (Slow, `MAGIC-ITEMTRAIN-116`): an actor whose book holds id 28 therefore still casts it through `L05147` and returns 1 even though Slow was never filtered for mana or `Defensive`, and only an actor whose book also lacks id 28 returns 0 and reaches the same default-engage fallthrough. So `R0393` reaching `ord+0x08 = 5` is conditional on the hand-back branch or an id-28-less empty spellbook, not a blanket property of "failure here". A fresh call on a later tick re-runs the slot loop with fresh draws independently. Class reachability of this routine at all is narrower than "any mage": a fresh read of `R0184` finds exactly one bit-set touching `actor+0x4c`'s bit 2 (the mage bit) in the whole function, `L01773` (setting bits 0x6) gated on `actor+0x0e ∈ {0x47,0x48}` (Dragon, Daemon; `UNIT-SPELL-007`) — every other spellbook-bearing class (10 of the 12 `UNIT-SPELL-007` names) gets only bit 1 (set at `L01772` as bit 0x2, gating spellbook allocation, not the mage bit). `R0184` is one of three writers of the mage bit, not the sole one: `HERO-CLASS-013`'s own Human placement-path read-modify-write (`L04438`/`L04440` setting bits 0x6/`L04439`) sets the identical bit whenever streamed `ManaMax` is positive, and `R0842` sets it too for a book with positive mana (`MAGIC-BOOK-002`), though `MAGIC-BOOK-002` finds no enumerated caller of `R0842` at all, so that third writer's own practical reachability is itself Unknown. Within `R0184` alone, only Dragon and Daemon receive the mage bit; the other ten classes it hands only bit 1. For an actor without the mage bit, `R0009`'s per-slot draw is unconditional inside that routine; a mage-bit actor with reach `< 2` and owner `+0x28 == 0` skips it, and owner `+0x28 == 0` by itself blocks only `R0393` (`L01774`), not the per-slot draw (`AI-341`), never through this routine (`AI-341`, `MAGIC-221`)*

**Original status.** ● active

**Evidence.** [EXP-0066](../experiments/EXP-0066-mind-spirit/), [EXP-0387](../experiments/EXP-0387-creature-spellbook/EXP-0387.md)

### MAGIC-221

**The spell a creature picked — from either the per-slot draw or the Mind-gated walk — reaches a cast order through `R0209`'s own 28-entry jump table, which sends 16 ids to the shared writer `R0018`, 10 to a kind-9 order `R0209` writes itself, and 2 to no order at all; the Mind-gated walk bypasses `R0209` and calls `R0018` directly.** `R0209(this, actor, target, spellId)`: looks the id up via `Spellbook::Get(actor+0x140, spellId)` (`L05150`…`L05151`, `EnumRefs callto:R0017` places this call among Spellbook::Get's 10 image-wide callers); on a null result (`L05152`, a jump on zero to `L05153`), jumps past every write in the routine, including `ord+0x60 = 1` — **a null lookup writes nothing**; otherwise range-checks `spellId − 1 ≤ 0x1b` (`L05154`…`L05155`) and dispatches through the 28-entry jump table at `L05156` keyed by `spellId − 1` (`L05157`), read directly off both lawful roots by `tools/spellarms` (`evidence/dispatch-table.tsv`; the roots agree byte-for-byte over the table's own span) into six distinct targets: `L05158` (ids 1, 11, 13, 14, 20, 27, 28) and `L05159` (ids 5, 6, 10, 15, 16, 18, 22, 23, 24) both call `R0018` — 16 ids total; `L05160` (ids 2, 3, 7, 8, 12, 17, 19, 21) writes `ord+0x08 = 9` and the target's own cell directly, bypassing `R0018`; `L05161` (id 9, Acid Stream) and `L05162` (id 26, Teleport) also write kind 9 themselves — 10 ids total reach a kind-9 order without `R0018`; `L05163` (id 4, Fire Sacrifice; id 25, Control Spirit) writes no order at all, only `ord+0x60 = 1`. The Teleport arm (`L05162`…`L05164`) contains no call to `R0179` (`AI-RAND-058`'s `rand()`): when `R0167` returns more than 2 it runs a 3×3 nested search over up to nine `R0144` calls (`L05165`…`L05166`) for an empty adjacent cell, and either way it always falls through `L05164` into the same `L05161` tail Acid Stream uses, which overwrites `ord+0x3c` with the cell adjacent to the actor in the target's own direction (`L05167`…`L05168`) regardless of what the search found — out of scope as spell-effect mechanics, noted only for the record. Every arm that does not already return converges on `ord+0x60 = 1` before return (`L05163`…`L05153`). `R0018(this, actor, target, spell)`: when `target` is non-null, or a null target's own occupancy-block/diplomacy search (`R0020`, `R0021`, `AI-336`/`AI-FILTER-001`) finds one, writes `ord+0x08 = 8` (cast-at-actor), `ord+0x28 = target`, `ord+0x30 = spell`, and `ord+0x14` from the actor's own weapon-reach byte (`actor+0x12c`, `L05169`…`L00074`); when the target stays null (a null argument and a failed search), it instead calls `R0022` (`L05170`…`L00076`) and writes no target/kind/spell fields. **Both branches then converge and write `ord+0x14` a second time, from the spell's own Range byte** (`spell+0x9`, `L05040`/`L00075`) — on the found-target branch this replaces the weapon-reach byte just written; on the no-target branch it is the only write to that field inside this routine. The instructions at `L05040`/`L00075` are the same pair `MAGIC-REACH-178` already cites (there, as "unit constructor... copy `Spell+9` into `order+0x14`, replacing weapon reach") — this is a re-read of the same call site, not an independent corroboration from an unrelated one. `R0393` and `R1048` (`MAGIC-AIBIT-082`'s own fixed-id-list picker) both also terminate by calling `R0018` directly, without passing through `R0209` at all (`EnumRefs callto:R0209`: 1 hit / 1 owner, `R0009` only) — confirming `R0018` itself, not `R0209`, as the shared cast-order writer for at least three distinct spell-selection entry points (`EnumRefs callto:R0018`: 13 hits / 8 owners). `R1048` has zero direct or `.rdata`-slot callers image-wide (`EnumRefs callto:R1048`: 0 hits / 0 owners) — noted, not itself a claim of this row, since it does not touch the six spell-slot dwords this experiment traces

**Confidence.** High for `R0209`'s dispatch structure, the spell-id identity of its fourth argument (proven by its own body's `Spellbook::Get` call, not inferred from the caller), the 28-entry table's own id-to-target mapping (read directly off both lawful roots, not decompiler-resolved, and every value checked against the six branch-target addresses already visible in `R0209`'s own disassembly), the null-lookup no-write finding, and the reach-byte/Range-byte write order on the found-target branch (each a complete-body read; the `MAGIC-REACH-178` overlap is the same instructions re-read, not cross-corroboration) / Medium for the two non-`R0018` arms' own further behavior — read only far enough to confirm they write a cast-at-cell order and reach `ord+0x60 = 1`, not disassembled past that for their own random-cell-offset math, which is spell-effect territory out of this experiment's scope

**Original status.** ● active

**Evidence.** [EXP-0387](../experiments/EXP-0387-creature-spellbook/EXP-0387.md)

### MAGIC-CEIL-013

**The reachable domain of the power, the one consumer that reads it unclamped, and the twelve expressions that consume it.** Both terms are bounded by named instructions: the school skill is clamped to `[0,100]` twice in the derive (`L03933`, `L03934`, `HERO-SKILL-009`) and every stat arm of the effect dispatch clamps the live stat with a comparison against 0x64 — Mind's at `L03981`, Spirit's at `L03982` — while the school is the `Sphere` column, `1..5` on all 28 shipped rows, so skill slot 0 is never it. **Therefore `skill + Mind − 30` ranges over `[−30, 170]` and the `[0,100]` clamp binds from `skill + Mind = 130` upward: it is reachable, not defensive.** Three sites compute the expression and **they do not agree**: `R0003` (`L05171`…`L05066`) and `R0904` (`L05172`…`L05173`) clamp to `[0,100]`; **`R0253` does not** — `L05174`…`L05175` is `max(x, 0)` returned as a byte, and its single caller `R0269` (`EnumRefs callto:R0253`: 1 hit / 1 owner / 0 orphan) takes it with a mask of 0xff at `L05176` for the **Prismatic victim cap** and for nothing else. That unclamped read is nonetheless **inert**: `victims = min(power/20 + 2, 7)` reaches 7 at power 100, exactly where the clamp begins, and re-executing both forms over **all 10 201 reachable `(skill, Mind)` pairs** gives a different victim count on **0** of them (`evidence/power-domain.txt`). Ceilings, as integers the engine stores: `f = power/30 + 1` tops at **4.3333**, so damage is 4.25×–4.33× the `Data.bin` columns (Meteor Storm 5–25 → 21–108); duration multiplies by `1.025^100 = 11.81` (Invisibility by `1.05^100 = 131.5` on its own arm, capped at 65 000 ticks from power 148 — unreachable); range's `/30` adds at most **3** cells and Teleport's `/3` adds **33**; the Prismatic victim count saturates at **7**. And the power is consumed in **thirteen** distinct expressions, not four (twelve as first published; EXP-0067 re-counted the same 21 reads and separated Stone Curse's duration from the four this row grouped): 21 reads of the local plus one of `f` inside `R0003`, each attributed to its arm off the image's own jump table (`evidence/apply-arms.txt`) — `Freezing Cloud −(p/15 + 1)`, `Light p/30 + 1`, `Darkness −1 − p/30`, the four `Protection` spells `p/2`, `Shield p/10 + 3`, `Haste`/`Slow p/15 + 1`, `Bless`/`Curse (p×4)/5 + 20`, `Poison Cloud ftol(effect+0x40 × f)`, `Fire Sacrifice`'s two totals capped at `0x200`, **five** `pow(1.025, p)` durations — `L05177` the four Protections, `L05178` Shield, `L05179` Haste/Slow, **`L05180` Stone Curse** and `L05181` Bless/Curse, the Stone Curse one added by EXP-0067 because its arm writes no magnitude and the enumeration reached it through the magnitudes — and two `pow(1.05, p)`, and the area builder's `effect+0x4c = (AreaEffectDuaration << 4) + (p << 4)/10` past the switch, whose first term this row originally quoted from `MAGIC-SHAPE-008` as `distribution << 4` (retracted.md)

**Confidence.** High (every bound, every divisor and the clamp/no-clamp split are named instructions read at instruction level, and the "the unclamped read never matters" clause is not an argument from shape but an exhaustive re-execution over the whole reachable domain, which discriminates) / **Medium** for the store audit alone — over all 28 shipped spells × power 0..100 the maxima are base 43 and spread 87 against byte fields and duration 6 312 against a `u16`, so **nothing overflows and the clamp rather than a truncation is the binding ceiling**; that is a statement about the shipped table and moves with it

**Original status.** ● active (amended)

**Evidence.** [EXP-0066](../experiments/EXP-0066-mind-spirit/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-ARM-014

**28 spells are 17 rules, and a spell with damage never reaches its own.** `R0003`'s per-spell dispatch is the 29-dword table at `L03242` (`L05182` bounds the id against 0x1c, `L03243` jumps through the table entry at `L03242 + id*4`), indexed by the spell id itself: **18 distinct arm addresses, 17 of them reachable by a shipped id**, with `protection_from_fire/water/air/earth` on one address (`L05183`), `bless`/`curse` on one (`L05184`), `haste`/`slow` on one (`L05185`) and the seven pure damage spells on the default (`L03046`). **The switch is reached only when the spell has no damage:** the damage half runs first, and `L05186`, a jump on less-or-equal (scratch base + spread ≤ 0), `L05083`, a compare with 0x6, and `L05084`, a compare with 0xb, are the only ways in — so `fire_arrow`, `fire_ball`, `wall_of_fire`, `acid_stream`, `lightning`, `prismatic_spray` and `meteor_storm` never execute an arm at all and `L03046` is dead code, while `heal` and `drain_life` reach theirs *because* their ids are excluded by name. What makes that safe is a third guard nobody had read: `R0625` writes a literal 0 rather than computing when a damage column is ≤ 0 (`L05187`/`L05188`, `L05189`/`L05190`), so the `−1` sentinel scores 0 — re-executed over the shipped rows, **19 of the 19 spells with no damage column reach their arm; under the unguarded rival, which stores `ftol(−1 × f)` as the byte `0xff`, 0 of 19 do** (`evidence/discrimination.txt` D2). The same routine's `L05191`, a zero test with a jump on zero, means a **`Max Range` column of 0 gets no range bonus**, which bites `fire_sacrifice` and `shield` (D4). Independent corroboration of the table itself: `EXP-0066/evidence/apply-arms.txt` extracted it for another purpose and the two agree **29/29**

**Confidence.** High (the table is read out of `.text` by virtual address, not inferred from similar code, and all 17 arms are read whole; the two guard clauses are their own compare-and-jump operands and the 19/19-vs-0/19 test discriminates against the only live rival)

**Original status.** ● active

**Evidence.** [EXP-0067](../experiments/EXP-0067-spell-catalogue/)

### MAGIC-EFFECT-015

**What an arm builds: the `Effects` column supplies the kind, the arm supplies the number, and only the first record of the column is ever parsed.** Eleven arms allocate `0x48` bytes and call `R1047(row+0x1c)`, which runs `R1000` over the row's own `Effects` string — `Find('=')`, then a linear search of the 50 kind names at `L05192` (`L05193`…`L05194`, **the only back edge in the routine**), then `effect+0x3c = the index`, then `R0857` for the mode word (`MAGIC-EFFMODE-009`). Index 0 means "not found" and the build returns 0, leaving the `Effect` all-zero. **Five arms use `R1049` instead** — `bless`, `curse`, `invisibility`, `wall_of_earth` and the dead default — which writes `+0x3c = 0`, `+0x3d = 0`, `+0x40 = 0`, `+0x0c = 0`, `+0x44 = 0`, so those spells' `Effects` columns (`toHit=+1`, `toHit=-1`, `Invisibility = ON`) are **never parsed**. Where the arm writes `+0x40` the column's own number is dead — the four `Protection` rows ship `+30` and get `power/2` (`L05195`) — and where it does not, the column's number is the magnitude: `poison_cloud`'s `health=-2` scaled by `f` (`L05196`, the only arm that scales a template magnitude) and `stone_curse`'s `absorbtion=+5`. **The parser has one construction site and no loop over `,`**, so `stone_curse`'s second record, `defence=-20:duration 8`, is never built. Per-arm magnitudes, each a named instruction: `p/2` protections, `p/10 + 3` shield, `p/15 + 1` haste, its negation slow (negation at `L05197`) and freezing cloud (negation at `L05198`), `p/30 + 1` light and `−1 − p/30` darkness, `(p×4)/5 + 20` bless and curse; durations are `ftol(1.025^p × SpellDuration × 16)` recomputed inside the arm, except `freezing_cloud`, `light`, `darkness` and `poison_cloud`, which write no `+0x42` and keep the mode and count their own column parsed. `spell+0x10`, which `R0625` fills with exactly that duration, is read by none of the twelve routines that take a `Spell*` (the spellbook fold reads it from a record, `TEXT-096`): in those routines the only `word ptr [reg + 0x10]` is a stack argument

**Confidence.** High (every arm is read at instruction level and every magnitude is its own divide, add or negation; the "one record" clause is the routine's own single back edge and single construction site) / **Medium** for the two magnitudes that come from the string — `poison_cloud`'s `−2` and `stone_curse`'s `+5` — in the original experiment because numeric conversion was unread. `MAGIC-POISONINPUT-157` now reads the Poison numeric/parser path and conditionally executes its valid-text input to kind 6, mode 2, magnitude −2 and counter 128; CString/decimal-library substitution keeps that joined parse Medium. Stone Curse's numeric conversion remains outside that new population

**Original status.** ● active (Poison parser frontier amended)

**Evidence.** [EXP-0067](../experiments/EXP-0067-spell-catalogue/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/)

**Amended.** The table ledger carried the status "● active (Poison parser frontier amended)". `retracted.md` records a correction against this claim.

**Amended.** The statement that `spell+0x10` is read by nothing holds among the twelve routines that take a `Spell*`; the spellbook fold `R0082` reads the field from the record by displacement for the duration line (`TEXT-096`).

### MAGIC-ATTACH-016

**`actor+0x144` is a bitmask of timed-effect ids (including non-spell Potion id 0), same-id timed effects do not stack, and Bless and Curse annihilate each other.** `R0612(actor)` is what an attached effect runs: `L05199` builds the one-bit value and sets `1 << effect+0x0c` — which is why every arm stamps the spell id, and which is the mask `R0254` tests (`L05200` builds the one-bit value and tests it against the field `+0x144`). Before that, `L05201`…`L05202` walks the actor's effect list for a `+0x0c` of **23 while the new one is 27, or 27 while the new one is 23**: the old one is un-applied and its bit cleared, and the new effect is then **dropped** (`L05203`, an unconditional jump past the attach) — so casting `curse` on a blessed actor removes the bless and leaves nothing, and vice versa. Otherwise an effect of the **same** id already present is not stacked: `continuous` refreshes its `+0x42` only (`L05204`), anything else is un-applied, given the new `+0x40` and `+0x42`, and re-applied (`L05205`…`L05206`); only when none is present is a copy stored on `actor+0x20`. An **untimed** effect (`+0x3d & 3 == 0`, `R0674`; charges bit 4 alone does not enter this timer) is applied and never stored. `R0673` is the tick: a `continuous` effect re-runs `vt+0x40` when its **old remaining counter is divisible by 8**, before decrement (`L05207`, `L05208`'s sign-preserving mask with 7). This gives eight-tick spacing while unrefreshed; incoming continuous refresh can change that phase and cause application in the later actor phase of the same world tick (`MAGIC-POISONPHASE-159`); **`+0x42 > 9600` skips the countdown entirely** (`L05209`, a compare with 0x2580), which no shipped spell can reach — the longest is `1.025^100 × 30 × 16 = 5670`; on expiry it un-applies unless continuous, clears the id bit, and ORs `[L05210] = 128` into `+0x3d`, a **fifth** flag byte that is not one of `MAGIC-EFFMODE-009`'s five words. **The mask is self-healing**: `R0265` pairs every `R0254(actor, id)` bit test with `R1050(actor, id)`, which walks the effect list for that `+0x0c`, and when the bit is set with no effect behind it the bit is cleared in place — `L05211` masks with 0xff7fffff for id 23 and `L05212` with 0xf7ffffff for id 27. Instrument: `EnumRefs disp:144` — 125 hits / 63 owners / 2 orphan, the actor-family ones being `R0185` (zeroes it), `R0612` (sets), `R0673` (clears on expiry), `R0265` (clears a stale bit), `R0210` (serializes it), `R0254`, and the AI's `R0225`/`R1048` testing `0x100000`, i.e. bit 20, `stone_curse` *(this enumeration is INCOMPLETE — [EXP-0179] adds three actor-family readers it omits: `R0016` at `L05213` and `L00543`, `R0394` at `L05214` and `R0293` at `L05215`. `MAGIC-MASKREAD-081` carries the complete set; see [`retracted.md`](retracted.md))*

**Confidence.** High (the set, the clear, the annihilation and the no-stack replacement are read at instruction level, and the enumeration states its mode and its two orphan hits)

**Original status.** ● active (amended)

**Evidence.** [EXP-0067](../experiments/EXP-0067-spell-catalogue/), [EXP-0179](../experiments/EXP-0179-effect-action/)  · [EXP-0275](../experiments/EXP-0275-consumable-use/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-TARGET-017

*(EXP-0243 retracts this row's Prismatic no-diplomacy and own-party clauses: the selector reaches the first-member diplomacy filter through `R0110`; other spell-effect arms stand.)* **Two different columns decide what a spell is aimed at, they disagree on one shipped row, and Prismatic Spray has a separate selector.** `R0621` is `getParam(row, 4) == 1` — title 5 **`Spell Target`** — and it is the cast's own routing: 1 queues the cast against the **target unit** (`R0617`), anything else against the **point** (`R0618`, `L05216`…`L03025`). At apply time a *different* column decides the shape, `Distribution system` (`MAGIC-SHAPE-008`), and value 1 **requires a unit**: with none, the tail prints `"Spell, oops - can't cast point effect of x,y"` and jumps to the epilogue, so the effect is built, orphaned and dropped (`L05217`, `L05218`). The two columns agree on 27 of 28 rows; `shield` alone ships `Spell Target 2` with `Distribution 1`. **No direct diplomacy-matrix operand occurs in the selector itself:** `EnumRefs disp:a9c4` — 32 hits / 14 owners / 0 orphan, and `imm:a9c4` returns **0**, the mode that would have missed the table entirely — puts no hit in `R0109`, the area collector, and exactly **one** in the whole 6 KB apply, inside `heal`'s arm at `L05219`, where bit 0 of `[L00004] + 0xa9c4 + caster.player×0x32 + target.player` **refuses** the heal. So `heal` is the only spell-effect arm with a direct matrix load; Prismatic secondaries separately reach diplomacy through `R0110`, while the other AreaEffect arms remain unchanged. Dead targets: `heal` refuses `health <= -10` or `target+0x98 == 0` (`L05220`, `L05221`); `drain_life`, `stone_curse`, `haste`/`slow` and `bless`/`curse` require `target->vt+0x2c()`, which the base actor returns 1 for; `control_spirit` requires a corpse at stage 2 and nothing else. And `MAGIC-RESIST-006`'s "a spell never rolls to hit" **holds for all 17 arms** — no arm calls `R0265` — with one arm that reaches a resistance by another route (`MAGIC-SING-019`)

**Confidence.** High (both routing tests are the `getParam` index and the comparison that reads it; the diplomacy clause is an enumeration with its mode, its blind spot and its one in-apply hit named in place)

**Original status.** ● partially retracted

**Evidence.** [EXP-0067](../experiments/EXP-0067-spell-catalogue/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-TRAIN-018

*(partially retracted by EXP-0191: the immediate book-cast award stands. The later claim that every nonzero-id damage or kill award reaches a school and none can reach a weapon skill is false: bit 4 selects the school arm, the clear-bit item rider substitutes the current weapon skill, `MAGIC-ITEMTRAIN-116` replaces that clause; `MAGIC-ITEMKILL-117` follows the separate kill feed through its post-apply repair.)* **Casting trains the spell's own school by half its mana cost, and only for a hero.** `L02776`, a virtual call of the slot at offset 0x68, runs on **every** apply from a book — gated on `R1051(caster)`, i.e. `caster+0x4c & 2`, the spellbook bit, and on `caster+0x68 == 0`, not an item cast (`L05222`, `L05223`, `L05224`). On the plain actor vtable `L00001` that slot is `R1052`, a two-instruction stub, so **a monster's casting trains nothing**; on both hero vtables (`L00002`, `L00003`) it is `R1053`, which requires the mage predicate `R0442`, reads the spell's `Sphere` (`getParam(row,2)`) as the skill slot and its `Mana Cost` (`getParam(row,1)`) as the amount, and calls `vt+0x5c` with `ftol(manaCost × 0.5 + 0.5)` — `[L02973] = 0.5`, so **round(manaCost/2), 2…50 over the shipped table, and 0 for `fire_sacrifice`, which is free**. Its two siblings in the same vtable block take a spell id the same way and are the same three-line shape: `vt+0x60` (`R1054`, a kill) awards `ftol(victim+0x1c × 0.5)`, `vt+0x64` (`R0870`, damage) awards `ftol(victim+0x1c × 0.5 × amount / victim.healthMax + 1)`, with a nonzero identifier, but the sink's class-bit arm still decides between a school and the current weapon skill — which is how `drain_life` (its real amount), `slow`, `curse` and `stone_curse` (`ftol(target.healthMax × 0.03)`, `[L05225] = 0.03`) pay experience for doing no damage. Two consequences: an **area spell trains once per unit hit**, because `R0269` runs a whole `Spell::Apply` per collected actor; and the whole of `claims/hero.md`'s open item 5 and `EXP-0065`'s open item 6 — "the school's training" — is these three routines. *(amended by EXP-0135: the sink `R0810` is read end to end and its two class arms are `HERO-SKILLUP-073` / `HERO-SKILLGATE-074`. Two clauses of this row gain a bound the round could not state. **this row's call-site gate and the callee's predicate are two DIFFERENT bits of `actor+0x4c`**, and both must hold for a cast to train: `R1051` is a mask with 0x2 at `L05226` and `R0442` — which `R1053` requires, and on which the sink itself branches at `L04498` — is a mask with 0x4 at `L05227`. Since the sink's `& 4` arm is the one that takes a school, **that conclusion applies only while bit 4 is set; a clear-bit item rider can route the same damage award to the current weapon skill**: the sibling arm of the sink refuses any slot but 0 and then substitutes `actor+0xb6`. The reverse follows as well — a book carrier's melee kill arrives with spell id 0, hence slot 0, and the book arm drops it. And the amount is capped on arrival at one level's worth of the target school, `S(n+1) - S(n)`, or a fifth of that with more than one participant, so `acid_stream`'s 50 is not always 50.)*

**Confidence.** High (the gate, the two column reads and the rounding are named instructions; the fighter/hero split is the vtable slot itself, not a naming argument) / ~~**Unknown** (what `R0810` does with the three arguments)~~ — **closed by [EXP-0135]** (`HERO-SKILLUP-073`, `HERO-SKILLGATE-074`); `HERO-XP-010`, which this row pointed at for the per-skill accounting, is itself **retracted** and replaced by `HERO-XP-077`

**Original status.** ● active (partially retracted)

**Evidence.** **[EXP-0067](../experiments/EXP-0067-spell-catalogue/)**, **[EXP-0135](../experiments/EXP-0135-skill-raise/)**, **[EXP-0191](../experiments/EXP-0191-itemcast-training/)**

**Amended.** The table ledger carried the status "● active (partially retracted)". `retracted.md` records a correction against this claim.

### MAGIC-SING-019

**Correction:** The universal never-queued/never-flies shorthand for Prismatic Spray is partially retracted by MAGIC-DELIVERY-170 and MAGIC-CASTCLOCK-171: the admission fan prepares a transport per victim; selection and the duplicate-prevention wrapper stand.  *(EXP-0243 retracts the Prismatic radius clause: `min(power/20+2,7)` is the final selected-victim count, primary included.)* **Eight spells do something no other spell does, and each is a named instruction rather than a shape.** (a) **`fire_sacrifice`** computes its own damage pair from the caster: `base = min(health + mana, 255)`, `spread = min(min(health + mana + power, 512) − base, 255)`, then `health := 1`, `mana := 0` (`L05228`, `L05229`) and stamps the pair with the **hard-coded fire setter** `R1043` rather than the `Sphere` table; the `protectionFire 100 / 32 ticks` effect it also builds is **never attached to anybody**. (b) **`teleport`** moves the caster with `R1055` and is the only spell whose range divisor is `power/3` (`L05074`). (c) **`control_spirit`** consumes a stage-2 corpse (health `:= −10001`, stage `:= 5`) and constructs a **new `0x198`-byte actor from the `Data.bin` template named by the literal at `L05127` — `Ghost`** — at that cell, owned by the caster's Player, copying `reaction/2 + 1`, Mind, Spirit, `healthMax/2`, `+0xa6` and `+0xbe` from the corpse; the power enters no expression in the arm. (d) **`stone_curse`** multiplies its **duration** by `(100 − target+0xca)/100` — protectionEarth — with a floor of one tick (`L05230`…`L05231`): the only place a resistance shortens an effect instead of reducing damage. (e) **`bless`/`curse`**: the magnitude is a **probability**, `(power×4)/5 + 20`, that `R0265` tests as `magnitude > R0861(100)` — an inclusive `U[0,100]` — to take the maximum (`L05232`, `L04068`) or the minimum (`L05233`) of a damage roll without rolling: `20/101` at power 0, `100/101` at power 100. (f) **`invisibility`** takes its duration from the literal `[L05234] = 3` on a base of `1.05`, ignoring its own `Spell Duration` 20 except as a `> 0` gate, and is cancelled by its owner's next cast at another actor (`L05235`) and by his own melee approach (`R0245`). (g) **`prismatic_spray`** is never queued and never flies: the cast's own `id == 14` arm re-fills the scratch and calls the multi-target applier, giving each entry of the selected list, capped at `min(power/20 + 2, 7)` total victims, a **separate, complete `Spell::Apply`** — and `R0002`, the common item wrapper, **refuses id 14 outright** (`L05236`, a compare with 0xe). EXP-0191 retracts the universal item conclusion: the caster item route has already run the fan during admission while the item context is live, and the wrapper prevents a duplicate; only the fighter rider reaches the wrapper without that admission fan and does nothing. (h) **`wall_of_earth`** builds an `Effect` that carries nothing at all — not even `+0x0c`, so it is the one spell that sets no `actor+0x144` bit

**Confidence.** High for (a), (b), (d), (e), (f), (g), (h) — each is read at instruction level and each singularity is an immediate or an absent store / **High / Medium** for (c)'s `Ghost`: the by-name constructor and the Units row it resolves are read by `MAGIC-249` (row and field provenance High, row streaming Medium) and the difficulty question by `MAGIC-250` (Medium)

**Original status.** ● active (item and delivery clauses partially retracted)

**Evidence.** [EXP-0067](../experiments/EXP-0067-spell-catalogue/), [EXP-0191](../experiments/EXP-0191-itemcast-training/)

**Amended.** The table ledger carried the status "● active (item and delivery clauses partially retracted)". `retracted.md` records a correction against this claim. EXP-0464 narrows the clause (c) Medium: the by-name constructor and the Units row are now read (`MAGIC-249`, `MAGIC-250`).

### MAGIC-AUTOCAST-020

**A weapon-borne spell has two triggers, not one, and they are exact complements — the caster's attack is *replaced* by the cast while the fighter's carries it as a rider.** `MAGIC-ITEM-007` read `R0246` and not its caller. `R0001`, the strike wrapper, has **exactly one** reacher — `L00006`, inside the actor tick `R0037` — and so does the attack start `R0245` (`L05237`); neither address occurs as a dword anywhere in the file, so neither can be dispatched through a vtable, a jump table or a callback array. That one call site is the **else** of a four-part test the tick carries twice, at phase 0 (`L05238`) and at the phase-5 countdown expiry (`L05239`): `actor+0x54 == 3` **and** `actor+0x74 != 0` **and** `R0408(weapon)` = `weapon+0x80 != 0` **and** `R0442(actor)` = `actor+0x4c & 4`. When all four hold, the arm writes `actor+0x64 = weapon+0x80`, `actor+0x68 = weapon` and `actor+0x54 = 0xd`, and the strike is never called; `R0268` then validates, the wind-up runs `actor+0x134` frames, and the call at `L00007` to `R0002` applies the spell — the **same** routine `L03045` reaches from inside the strike, and `R0002` has exactly those two callers. State `0xe` is the same arm aimed at a point (`R0003`, `actor+0x60`/`+0x61`), and after either release an item whose kind `R1056` returns `0xe` is destroyed along with its `Spell` (`L05240`…`L04547`), which is why a staff survives its own cast and something else does not. **The bit's polarity is fixed at its write, not by naming:** `L04440`, a bit-set of 0x6 on `actor+0x4c`, sits behind `L04441`/`L04442`, a test with a jump on less-or-equal, over `actor+0x9c`, the streamed `ManaMax` column (`HERO-HP-071`), and the same arm then allocates the spellbook — so set is the caster, `R0442 != 0` is the caster, and it is the caster whose attack becomes a cast

**Confidence.** High (every hop is a named instruction; all four `rel8` displacements in the phase-5 arm resolve to `L05241` and the arm's exit `eb13` to `L05242`, so the block is internally consistent without a second decode). The single-caller enumeration is **complete by three instruments with disjoint blind spots** — Ghidra's reference manager, a whole-file LE-dword scan that returns 3 `.rdata` slots for `R0037` in the same run, and a byte-by-byte `E8`/`E9` `rel32` walk of `.text` that reproduces `EnumRefs` on three controls (24 / 7 / 1). **Blind spot:** a call whose target address is held in a register loaded from neither an immediate nor a table is invisible to all three. The polarity is discriminated by three independent legs — the `ManaMax` gate, the starting-kit assignment (`MAGIC-STAFF-022`) and `HERO-HP-005`'s multiplier — of which two are new here

**Original status.** ● active

**Evidence.** [EXP-0137](../experiments/EXP-0137-staff-autocast/)

### MAGIC-AUTOCAST-021

**A caster attacking with a spell-carrying weapon cannot miss and cannot be absorbed, and the reason is structural: the routine that rolls is never entered.** `HERO-AUTOHIT-031`'s auto-hit bit `actor+0x4c & 0x10` is not involved and could not be — its only located writer is the **units** streamer's `attackKind == 3` arm, whose shipped population is 8 creature rows, a different collection from any weapon. What happens instead is that the to-hit roll, the damage roll, the flat absorption and the per-class reduction all live inside `R0246` and its `vt+0x4c` damage call (`L02047`, `HERO-DAMAGE-022`), and `MAGIC-AUTOCAST-020`'s diversion means a caster with `weapon+0x80 != 0` never reaches `R0001` at all. What a consumer must build is therefore **not** an auto-hit flag on an item: it is that such an actor's attack **is** a cast and delivers no weapon damage. The reach test changes hands with it — `L05243`…`L00736` compares `actor+0x12c` against `R0247(attacker, target)` on the melee path only, and the cast path's admission is `R0268`'s. The predicate is the item's `Spell`, not its shape: a caster holding a weapon with `weapon+0x80 == 0`, or holding nothing, falls to the melee path and rolls like anyone else

**Confidence.** High for the **bypass**, which is `MAGIC-AUTOCAST-020` plus the single-caller enumeration and adds no new reading. The clause that nothing else grants a weapon auto-hit rests on `HERO-AUTOHIT-031`'s own **Medium** writer-set clause and inherits it — the byte-wide bit-set of 0x10 at `+0x4c` through a base register and a wholesale copy of the containing block are unswept. What a player sees is a **prediction**, not a measurement: the write-up states five, each with what would refute it

**Original status.** ● active

**Evidence.** [EXP-0137](../experiments/EXP-0137-staff-autocast/)

### MAGIC-STAFF-022

**The character generator hands a caster a staff and only a staff, and that is a fact about distribution — nothing in the image refuses an item on account of class.** `R0837` splits on `frame+0xc & 0xff` against 10 and each half is partitioned by `R0442` (`L05244`, `L05245`). `StrDump fnstr:R0837` returns **fourteen** string operands and no others, so the table is complete for that routine: caster, tier `> 10` — `Wood Staff {castSpell=Fire_Arrow:20}` and `Uncommon Wood Staff {castSpell=Fire_Arrow:20}`; caster, tier `<= 10` — the same two at `:10`; non-caster, tier `> 10` — `Uncommon Steel Two Handed Sword`, `Uncommon Steel Axe`, `Uncommon Steel Mace`, `Uncommon Steel Pike`, `Uncommon Magic Wood Short Bow`; non-caster, tier `<= 10` — `Iron Short Sword`, `Uncommon Bronze Axe`, `Uncommon Bronze Mace`, `Bronze Pike`, `Uncommon Wood Short Bow`. The second gate `R0841` is **not** a class test — it compares `actor+0x0e`, the typeID (`DAT-ACT-006`), against `0x22` or `0x24`, which under `HERO-CLASS-013`'s `gender + 0x21`/`0x23` overwrite selects a gender — so it picks the plain or `Uncommon` variant, not the weapon. Each name goes to `R0665`, the item-from-name constructor, on a fresh `0x84`-byte allocation. **What this does and does not say:** the code gives a generated caster exactly one weapon and it is a staff carrying a levelled `Fire_Arrow`; it does not refuse anything, and `ITEM-SUIT-035` (two display flags with no routine that refuses on them) and `ITEM-HUMEQ-030` (no class restriction on the equip path) stand unchanged. *A caster only has staves* is therefore about what the game **gives**, not about what it **allows**

**Confidence.** High for the enumeration — `StrDump fnstr:` prints every string operand of the named function, so the fourteen are a population rather than a selection, and both gates are short bodies read as raw listings. **Blind spot:** an item built from a name assembled at runtime rather than from a literal would carry no string operand and would not appear; nothing in the listing does that, but the instrument could not see it. Says nothing about what a *map*, a *shop* or a `Data.bin` template gives an actor

**Original status.** ● active

**Evidence.** [EXP-0137](../experiments/EXP-0137-staff-autocast/)

### MAGIC-SPELLHOP-023

**`MAGIC-ITEM-007`'s open hop is dead code, and the question it framed has a false premise: nothing in the image resolves a weapon's spell by NAME.** `R1045` — which takes a `CString`, builds a `Spell` through `R1040` and stores it at `item+0x80` — has **no reacher by any instrument**: 0 on `EnumRefs callto:`, **0 occurrences of `R1045` as an LE dword anywhere in the 1 977 344-byte file** (so it is in no vtable, static-init table, jump table or callback array in any section), and **0** `E8`/`E9` `rel32` in the whole of `.text` including bytes never disassembled. The live builder is its near-twin **`R1012`**, which differs in exactly one thing — its argument is a **byte** and it calls `R0463(spell, id)` — reached only from `L05246` inside `R1011`, reached only from `L05247` inside `R0850`, the item's `vt+0x38`. `R1011` is three steps: `R0858(item)` finds the effect with `effect+0x3c == 0x29`, `L05248` reads the spell id out of `effect+0x40`, and `R1012` allocates `0x14` and hangs the `Spell` on `item+0x80`. So the chain from the shipped text is: `R0665(item, name)` zeroes `item+0x80` (`L05249`), resolves the text **outside** the braces to the kind `item+0xc`, and passes the text **inside** them to `R0992` only when non-empty (`L05250`, a jump on less-or-equal); `R0992` is two instructions — `R0999(text, item+0x20)`, which trims and splits on `,` (`L05251`/`L05252` supplying 0x2c) and hands each token to `R1000`, then `vt+0x4c` = `R0962`, which is the **price** walk and adds a castSpell effect's own `vt+0x4c` to `item+0x1c`, clamped at 9 999 999. The four occurrences of the ASCII `castSpell` in the image are the four shipped item names themselves; there is no keyword string and no xref to one, consistent with the token dispatch living in `R1000` against the `Data.bin` Magic-row machinery (`MAGIC-EFFECT-015`) rather than against a literal

**Confidence.** High. The deadness rests on three instruments with disjoint blind spots, each shown able to return non-zero on a control in the same run — the `rel32` walk returns 24 / 7 / 1 for `R0442` / `R0859` / `R0246`, matching `EnumRefs` hit for hit, and the dword scan returns 3 `.rdata` slots for `R0037`. **Blind spot, stated:** an address computed at runtime by arithmetic that never materialises the constant is invisible to all three. **Not established:** which of `effect+0x40` (id) and `effect+0x42` (`MAGIC-ITEM-007`'s power) receives the `:n` suffix and which the name — `R1000` was not read

**Original status.** ● active

**Evidence.** [EXP-0137](../experiments/EXP-0137-staff-autocast/)

### MAGIC-ICON-024

**A spell's icon is a POSITION in one pre-composited strip, not an indexed cell — and four of the twenty-eight spells have none.** `graphics\interface\SpellBook.bmp` (480x85, 24bpp, identical bytes on both roots) is loaded into the global at `L05253` by `R1057` — `L05254` pushes the path, `L05255` stores the object, and the store **follows** the constructor, so a path pairs with the NEXT store; pairing it with the preceding one shifts every interface surface by one slot and makes this atlas look unread. `EnumRefs refto:L05253` on the repaired table returns **7 hits / 3 owners / 0 orphan**: the load, `R1058`'s release sweep, and five in `R1059`, the spell panel, which blits the atlas **whole and once** (`L05256`, a virtual call of the slot at offset 0x38) with no source rectangle and no cell index anywhere in the draw. What is per-spell is the *masking*: `L05257`…`L05258` is `for (i = 0; i < 0x18; i++)`, testing `1 << i` against a 24-bit mask at `+0x148` (`L05259`, a test of the 32-bit field `+0x148` against the bit) — read here as *the spells this panel may offer*, because the identically shaped loop in `R0082` tests a 24-bit mask and constructs a `Spell` for each set bit, not because the field itself has a named consumer and, where the bit is **clear**, pasting `SpellBack.bmp` — 36x36, exactly the cell — at `x = 6 + 38*(i % 12)`, `y = 6 + 38*(i / 12)`. The arithmetic is its own instructions: `L05260` multiplies by the magic constant 0x2aaaaaab and shifts right by 1, the signed divide by 12, `L05261` is a signed divide by 0xc whose remainder is used, and each coordinate is built by address arithmetic (9*d, then 19*d, then 38*d + 6). The grid's bounding box is 6 + 12*38 = 462 <= 480 and 6 + 2*38 = 82 <= 85, i.e. the atlas itself. The **only** slot-to-spell mapping the engine computes is a 24-dword table at **`L00199`** (file `0x1c032c`; the RU `ROM.EXE` carries the same bytes at the same offset) reading `1 2 3 4 5 23 24 16 15 14 13 12 6 7 8 9 10 25 26 22 21 20 19 18`; `EnumRefs refto:L00199` returns **1 hit / 1 owner / 0 orphan** — `L05262` reads a byte of the table at `L00199` with stride 4 in `R0082`, whose loop is `L05263` (the frame-local index compared with 0x18), `L05264` (a read of the 32-bit field `+0x18`, the mask), `L05265` (a call of R0463) — `MAGIC-SPELL-001`'s own `Spell(id)` constructor. **Ids 11, 17, 27, 28 — `drain_life`, `darkness`, `curse`, `slow` — are in no slot**, so 24 cells serve 28 spells and those four have no icon of their own. Corpus: cutting `spellbook.bmp` by that grid yields 24 cells each carrying 782..1194 distinct colours over 1296 pixels, modal share 0..3 %, so no cell is blank (`evidence/icon-cells-*.csv`). **G2:** 24 is an **engine** limit stated twice — the loop bound `0x18` and the table's length — and lifting it changes no shipped file's bytes, but needs both plus a wider atlas; the atlas geometry itself is a plain 24bpp BMP and is a file-format-free limit

**Confidence.** High (the single blit, the grid arithmetic and the table's one reader are each their own instructions; the table was located by a byte-exact scan of the image and then confirmed by what its sole reader constructs, so the shipped text is corroboration rather than the evidence)

**Original status.** ● active

**Evidence.** [EXP-0139](../experiments/EXP-0139-spell-pictures/)

### MAGIC-ICON-025

**The panel's other per-slot decisions, and the two shipped text tables cut 24 and 28.** In `R1059` the selected slot is the panel object's `+0x60` and receives a 36x36 outline — `R1060(x, y, x + 0x24, y + 0x24, 4)` — but only when its known-bit is set (`L05266` reads the 32-bit field `+0x60`, `L05267` tests the known-bit); four further slots at `+0x64`..`+0x70` are matched against the loop index and, on a match, a numeral `index + 5` is drawn two pixels further in when that slot is also the selected one; and the side panels fork on the global at `L01259` — below `0x1e1` none is drawn, `0x1e1`..`0x258` uses `spb800l/r` (the global at `L05268`/the global at `L05269`), `0x259` and above `spb1024l/r` (the global at `L05270`/the global at `L05271`) — the six shipped bitmaps being 80, 96, 192 and 208 wide by 85 high, the atlas's own height. Shipped text, both roots: `main.res::text/spell.txt` is **28** non-empty lines keyed by spell id (line *i* is id *i+1*) and `main.res::text/spells.txt` is **24** lines keyed by bar slot, each `Title#line#line…` — the same 24-against-28 split as the icons, reached from the data before the code was read. Three of the 24 titles disagree with the image's own 29-entry name array (`MAGIC-SPELL-001`, `L05034`): slot 9 reads `Chain Lightning` where the array reads `prismatic_spray`, slot 17 `Raise Spirit` for `control_spirit`, slot 22 `Stone Wall` for `wall_of_earth`. The RU root differs only in bytes — 421 and 1611 against 360 and 1216 — never in record count. **G2:** the two record counts are file-format-free, both files being line-delimited text, but a 25th line is unreachable without the engine change `MAGIC-ICON-024` prices

**Confidence.** Medium (the geometry and the resolution fork are instruction-level, but reading `+0x64`..`+0x70` as hot-keys is inferred from the numeral drawn beside them and from the save section `[SpellBook] Shortcuts` — `L05272` supplying L05273 with `L05274` as the section, read by the array getter `R1061` — not from a consumer; the slot-to-**name** pairing is a corpus match on shipped text and caps at Medium by construction — the slot-to-**id** pairing does not and is `MAGIC-ICON-024`'s)

**Original status.** ● active

**Evidence.** [EXP-0139](../experiments/EXP-0139-spell-pictures/)

### MAGIC-PIC-026

**The picture a cast puts on the map is COMPUTED, not looked up: `2*spellId + 8`, or `+9` for the burst variant — and seven spells compute an id no art defines.** No `Data.bin` `Spells` column carries it: over 28 rows on both roots `Delivery System` takes two values and `Distribution system` four, neither with the cardinality of a 31-valued id nor order-matching one. The arithmetic is in the image, seventeen times. Instrument, chosen because `docs/INSTRUMENT.md` rules out a function-table sweep for a completeness claim: a **raw opcode scan of the whole `.text`**, no disassembler in the path, for the address form `b + b*1 + disp8` with `disp8` in 8 or 9 (the address-form opcode, modrm mod=01 rm=100, sib ss=0 index==base) followed within 28 bytes by a 16-bit store to displacement 0x0e of a register (its prefix and opcode, a modrm and the displacement byte 0x0e). **28 sites match the first pattern image-wide; exactly 17 match both, and all 17 lie in `0x4fce`..`0x5006`** — the eleven discarded sit in `0x429`, `0x48f`, `0x490` and `0x4e9` and store nothing at `+0x0e`, which is what shows the filter discriminates. `+8`: `L05275 L05276 L05277 L05278 L05279 L05280 L05281 L05282 L05283 L05284 L05285`. `+9`: `L05286 L05287 L05288 L05289 L03050 L05290`. The source byte is the spell id at every site — `L05291` reads the byte at `+0x8` off the `Spell` (`MAGIC-SPELL-001` fixes `+0x08` as the id) or `L05292` reads the byte at `+0xc` off an already-built `Effect` (`MAGIC-ATTACH-016` fixes `+0x0c` as the id). **Completeness:** the same scan enumerated every 16-bit store to `[r + 0x0e]` anywhere in `.text` — 51 instructions, 21 of them inside `L05293`..`L05294`, 17 the above and the remaining four a literal `0` (`L05295 L05296 L05297 L05298`). No other value ever reaches the field. Against the shipped `projectiles.reg` (`REG-PROJ-086`), `2*id + 8` is a defined row for spells 1 2 5 6 8 10 11 13 14 16 18 22 23 26 27 and `2*id + 9` for 2 3 4 7 8 9 19 21 — so `fire_ball` and `poison_cloud` own **two** pictures each, and **`light`(12), `invisibility`(15), `darkness`(17), `stone_curse`(20), `haste`(24), `control_spirit`(25) and `slow`(28) compute an id with no row at either parity**. **G2:** the formula is an engine limit and a hard one — it pins the picture id to the spell id, so re-pointing one spell's art means editing `projectiles.reg`'s `ID`, which changes a shipped file's bytes; and it bounds the reachable id at `2*28 + 9 = 65`, one past the `ID + 1` array the loader grows

**Confidence.** High (a complete opcode-level enumeration over the whole section, immune to the function-coverage blind spot, with a control set of 11 near-misses; the two source fields are each fixed by an existing claim) / Medium (the receiving object is the one the arm is filling at sixteen sites and that object's `+0x44` at `L05299`; only the `L05300` site was traced to its 0x48-byte allocation)

**Original status.** ● active

**Evidence.** [EXP-0139](../experiments/EXP-0139-spell-pictures/)

### MAGIC-PIC-027

**The wire, and the parity fork: an EVEN picture id puts nothing on the map.** Four senders write the message opcode the client dispatches on (`ANIM-MSG-005`: `msg+0x9`), located by scanning `.text` for `C6 /r 09 <op>` — `0x86` at `L05301 L05302 L05303 L03055 L05304`, `0x8a` at `L05305`, `0x8b` at `L03021 L03022`, `0x8c` at `L03023`. `R0617(src, spell, target, dir)`, read whole: `L05301` writes opcode `0x86` and `L05306` the source's `+0x04` runtime id; if `word ptr [src + 0x0e]` is zero (`L05307`, `L05308`) the opcode is rewritten to **`0x8b`** and the two cell helpers `R0299`/`R0300` fill `msg+0xa`/`msg+0xb` instead; then `L05309` reads the spell id byte at `+0x8` of the `Spell`, `L05310` forms `2*id + 8` and `L05311` stores it at `msg+0xc` — **`msg+0xc` is `2*spellId + 8`, taken straight off the `Spell`**. `R0618` is the same routine aimed at a place rather than a unit (same store at `L05312`). `R0638` is the **deferred** path and the one consumer of `MAGIC-PIC-026`'s stamp: `L05313` reads the byte at `+0xe` and `L05314` stores it at `msg+0xc`. `R0619` sends `0x8a`/`0x8c` and writes no picture byte at all. On the client, `tools/animdrv -mode dispatch` puts opcode `0x86` on arm `L03098`, and that arm opens with `L05315` reads the byte `msg+0xc`, `L03011` isolates its bit 0 and `L05316` jumps to `L03013` on zero — **an even picture creates no map object**; control goes to `L03013`, which looks the caster up in the world hash at `+0x9b8` and gives it the cast action. An odd picture must additionally index a **loaded** slot (`L05317`, `L05318`, `L03012`) or nothing happens, and only then is a `0x14c`-byte `CProjectile` allocated (`L05319` supplies the size 0x14c, then the call of R0609) with `L03226` storing the 32-bit value at `+0x20` — `picture = msg+0xc` — and its cell from `msg+0xd`/`msg+0xe` scaled by `0x100` plus `0x80`. Arms `0x8b` (`L05320`) and `0x8c` (`L05321`) allocate the same `0x14c` with **no parity guard**. This resolves part of `ANIM-MSG-005`'s open item 1: `0x86` and `0x8b` are one sender's unit-addressed and cell-addressed forms, `0x8a` and `0x8c` another's, and the pair that carries a picture is `0x86`/`0x8b`

**Confidence.** High (every cited address is an instruction in a listing dumped over an address range, and the opcode set comes from a fixed-length byte scan rather than from a sweep that could miss orphan code) / Unknown (which routine calls each sender, hence which spells take the deferred path rather than the direct one)

**Original status.** ● active (amended)

**Evidence.** [EXP-0139](../experiments/EXP-0139-spell-pictures/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

## Open items

1. ~~**The AI's spell choice.**~~ Closed by `MAGIC-AI-012`, narrowed and completed by `MAGIC-221`/
   `AI-341` (`EXP-0387`): `R0209` resolves a slot value as a spell id via `Spellbook::Get`
   and dispatches through its own 28-entry table to six targets — 16 of 28 ids reach the shared
   writer `R0018` (`ord+0x08 = 8`, target and spell fields, reach overwritten by the spell's
   own Range byte), 10 reach a kind-9 order `R0209` writes itself, and 2 write no order; the
   Mind-gated walk bypasses `R0209` and calls `R0018` directly. The three alternative
   orders at
   `[actor+0x158] + 0x78..0x84` are the same class spellbook slots `UNIT-SPELL-007` writes, read by
   a per-slot `rand()` draw in `R0009` that is unconditional for an actor without the mage
   bit; a mage-bit actor with reach `< 2` and owner `+0x28 == 0` skips it instead
   (`AI-341`, narrows `AI-THREAT-044`); `[actor+0x14] + 0x28` is `Player+0x28` (`UNIT-OWNER-009`,
   partially retracted on reading zero as proof of human authorship).
   The other four routines EXP-0065 listed are still unread.
2. **`Data.bin` Magic row → `Effect`.** `R1000` builds the template `R1047` copies;
   only its mode parse is read (`MAGIC-EFFMODE-009`). `HERO-EFFECT-019`'s open item 1 stands.
3. **The `Distribution system` enum**, `MAGIC-SHAPE-008`.
4. **The `block+0x11`/`+0x12` damage pair**, `MAGIC-RESIST-006` — always applied, always resisted
   by protection index 2, and no writer located.
5. **The projectile.** `SpellTransport` (`R0633`, `R1062`, `R0634`) counts
   `+0x4c` down and hands its payload over at zero; its art and its per-tick step are unread.
6. **`R0904`**, the Prismatic Spray fan, which reads a school skill of its own.
7. **How an actor first gets a book** (`MAGIC-BOOK-002`) — one 0-caller routine, not two.
   ~~who authors a weapon's spell (`MAGIC-ITEM-007`)~~ closed by `MAGIC-SPELLHOP-023`: the 0 was
   real and meant *dead code*, and the live builder is one displacement away. What remains of that
   half is narrower — **`R1000`**, which turns one comma-separated brace token into an
   effect and is where the spell **name** becomes the id byte in `effect+0x40` and the `:n` suffix
   becomes a number. `MAGIC-EFFECT-015` reads only its first-record clause.
8. **What `actor+0x13c` is** — the byte `Control Spirit` selects its victim by (`== 2`) and then
   sets to 5 beside a health of `0xd8ef`; and `R0225`'s mode table at
   `[EBX + kind*4 + 0xb94]` and its `vt+0x1c()` (`MAGIC-AI-012`).
9. **The `memcpy` blind spot named in `MAGIC-MIND-010`.** A displacement sweep cannot see a
   wholesale actor copy. Closing it needs a different instrument — a census of `REP MOVSD` sites
   whose count matches an actor's size, which no experiment has run.
10. **Whether a cast object with `actionsegments = 0` is ever drawn.** `MAGIC-CASTSPAWN-033`
   establishes that 44 of the 51 switched picture ids get a zero flight length and that the driver
   then returns finished before its own arm runs. Whether the draw pass sees the object once before
   the driver removes it depends on the order of the two passes inside one world tick, which was not
   read. It decides whether the eight defined but zero-length cast sheets — `p_fire`, `poison_d`,
   `p_water`, `p_air`, `shield`, `p_earth`, `bless`, `Curse` — flash for one frame or never appear.
11. **What fills a unit's own effect lists at `+0x124` and `+0x128`.** Eleven even picture ids take a
   driver arm that inserts `picture << 16 \| 0xffff` into `+0x124`, and two more write
   `picture << 16 \| 0x20` into a slot of `+0x128`. Under `MAGIC-CASTSPAWN-033` those eleven ids are
   unreachable from the cast spawner, so either another producer fills those lists or the arms are
   reachable only through client arm `0x8b`. Neither was established.
12. **`Data.bin` parameter 7**, the divisor `MAGIC-DELIVER-035` uses for the simulation's own flight
   time. Its column title was not read.
| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-CAST-028 | (rom.exe) Four `projectiles.reg` sheets are also drawn as per-unit status overlays, by flag bits and not by a cast. | High / Unknown | ● active (partially retracted) | [EXP-0140](../experiments/EXP-0140-cast-art-drawn/), [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-CASTANIM-029 | A cast always animates the caster, and the even picture is what makes it do so. | High / Medium | ● active | [EXP-0163](../experiments/EXP-0163-cast-picture/) |
| MAGIC-CASTTICK-030 | The tick a spell is applied, in units a consumer can implement: tick `actor+0x134` of the action, counting the tick that started it as 0. | High / Medium | ● active (partially retracted) | [EXP-0163](../experiments/EXP-0163-cast-picture/) |
| MAGIC-BURST-031 | `MAGIC-PIC-027`'s Unknown is answerable: the `Distribution system` column decides the burst picture, and 10 of 28 shipped spells reach it. | High / Medium | ● active | [EXP-0163](../experiments/EXP-0163-cast-picture/) |
| MAGIC-SENDER-032 | Two corrections to `MAGIC-PIC-027`'s sender enumeration, one of which changes what a consumer must build. | Medium | ● active | [EXP-0163](../experiments/EXP-0163-cast-picture/) |
| MAGIC-CASTSPAWN-033 | (rom.exe) The object a cast puts on the map is built by the CASTER, not by the message, and its flight length comes from a hard-coded switch on the picture id -- non-zero for 7 picture ids and zero for the other 44. | High / Medium | ● active (amended, partially retracted) | [EXP-0167](../experiments/EXP-0167-spell-art/) |
| MAGIC-BURSTLIFE-034 | (rom.exe) A burst is stationary, and its sender chooses its lifetime. | High | ● active (amended, partially retracted) | [EXP-0167](../experiments/EXP-0167-spell-art/), [EXP-0175](../experiments/EXP-0175-ring-geometry/) |
| MAGIC-DELIVER-035 | (rom.exe) The simulation and the client each compute a cast projectile's flight length, by different rules, and on the normal path the client's wins. | High / Medium | ● active | [EXP-0167](../experiments/EXP-0167-spell-art/) |
| MAGIC-AREATICK-036 | (rom.exe) The `AreaEffect` per-tick entry is vtable slot `+0x18`, and `effect+0x08` selects one of three exclusive per-tick modes. | High / Medium | ● active | [EXP-0172](../experiments/EXP-0172-area-tick/) |
| MAGIC-AREAPULSE-037 | (rom.exe) A cloud pulses once every 16 ticks, a staged effect runs its first stage immediately and later stages every 3 ticks, and a blast runs once. `effect+0x4c` is a counter with mode-specific meaning. | High | ● active (amended, partially retracted) | [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0175](../experiments/EXP-0175-ring-geometry/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |
| MAGIC-AREAAPPLY-038 | (rom.exe) An area effect invokes its inner effect's per-target virtual entry; the dynamic class determines direct damage versus ordinary attachment. | High / Medium | ● active (amended, partially retracted) | [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/), [EXP-0271](../experiments/EXP-0271-multicell-area/) |
| MAGIC-AREACELL-039 | (rom.exe) Three different cell sets, and the simulated set is not the mask sent on the wire. | High / Medium | ● active (amended, partially retracted) | [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/) |
| MAGIC-MAPLAYER-040 | (rom.exe) The map carries exactly six area-effect layers, one per spell id, and five conflict rules delete one layer when another is laid. | High / Medium | ● active | [EXP-0172](../experiments/EXP-0172-area-tick/) |
| MAGIC-AREAEND-041 | (rom.exe) All three modes end at the same byte; cloud teardown removes its owned cell layers, whose helper can restore saved terrain/flags. | High | ● active (amended, partially retracted) | [EXP-0172](../experiments/EXP-0172-area-tick/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |
| MAGIC-WALLEARTH-042 | (rom.exe) `wall_of_earth` applies nothing to any unit and refuses occupied cells, and no passability write was found for it. | High / Medium | ● active (amended) | [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/) |
| MAGIC-AREATARGET-043 | (rom.exe) No per-tick routine of an area effect consults diplomacy, so an area effect pulses on the caster's own party. | High | ● active (amended) | [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-LIGHTDARK-044 | (rom.exe) `light` and `darkness` are the same effect with opposite sign, and only `darkness` provokes its victims. | High / Medium / Unknown | ● active | [EXP-0172](../experiments/EXP-0172-area-tick/) |
| MAGIC-WALLBLOCK-045 | (rom.exe) `wall_of_earth` blocks movement through the ordinary terrain-block bits, and a routine outside the area module derives them from the cell record. | High | ● active | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| MAGIC-AREACOST-046 | (rom.exe) Every occupied area layer multiplies its cell's movement cost byte by four, and the step-duration path divides it back by four once and stores the result. | High / Unknown | ● active | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| MAGIC-FIREDIV-047 | (rom.exe) `fire_ball`'s footprint-squared divide is a per-cell normalisation, not a size penalty — because a unit is stored in every cell record it covers. | High | ● active (amended, partially retracted) | [EXP-0174](../experiments/EXP-0174-area-movement/), [EXP-0271](../experiments/EXP-0271-multicell-area/) |
| MAGIC-RING-048 | (rom.exe) The three staged AreaEffect spells use three distinct cell generators. | High / Medium | ✔ promoted | [EXP-0175](../experiments/EXP-0175-ring-geometry/) |
| MAGIC-BOLTGATE-069 | (rom.exe) Exactly two picture ids produce a drawn path, and the routine that produces it is one virtual slot whose whole body sits behind a two-comparison gate. | High | ✔ promoted | [EXP-0178](../experiments/EXP-0178-bolt-path/) |
| MAGIC-BOLTSHAPE-070 | The bolt figure starts with a bounded random walk, inserts midpoint knots, samples quadratic triples and rotates the resulting points onto the projectile-to-target segment. | High | ✔ promoted (partially retracted, superseded) | [EXP-0178](../experiments/EXP-0178-bolt-path/); [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-BOLTLIST-071 | (rom.exe) The point list is rebuilt on every driver tick and it does not accumulate: picture 34 replaces it, picture 36 clears it and concatenates one generated set per victim id. | High | ✔ promoted (superseded) | [EXP-0178](../experiments/EXP-0178-bolt-path/) |
| MAGIC-BOLTSTILL-072 | On the inspected action-1 driver path, Lightning and Prismatic Spray keep their raw position; normal caster construction gives a 13-call countdown. | High | ✔ promoted (amended) | [EXP-0178](../experiments/EXP-0178-bolt-path/); [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-TRAIL-073 | (rom.exe) The trail behind a Fire Arrow or a Fire Ball is a queue of at most six PAST positions, and its oldest entry is dropped one at a time. | High | ✔ promoted | [EXP-0178](../experiments/EXP-0178-bolt-path/) |
| MAGIC-BOLTEND-074 | (rom.exe) Nothing is created when a Lightning or Prismatic Spray figure ends, and the whole figure ceases with the object. | High / Unknown | ✔ promoted | [EXP-0178](../experiments/EXP-0178-bolt-path/) |
| MAGIC-AREADRAW-049 | (rom.exe) The three `AreaEffect` tick modes create three different kinds of drawable, and only one of them creates one object per covered cell. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-OVERLAY-050 | (rom.exe) What a cloud puts on the client is a per-cell bitmask, not a set of objects. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-OVERLAYART-051 | (rom.exe) The retained overlay's art is four immediates, and `R0607` is the per-cell area-effect draw, not a per-unit status overlay. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-CLOUDLIFE-052 | (rom.exe) A cloud's overlay stands for the whole life of the effect, and the damage pulse draws nothing. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-CELLSET-053 | (rom.exe) The cells drawn are the cells the map layer actually holds, re-read from the map, not the nominal geometry. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-AREARADIUS-054 | (rom.exe) An `AreaEffect`'s radius is the shipped `Radius, Length/2` column read once at construction, not a power term, and the overlay bitmap is sized for the shipped values and no more. | High / Medium | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-LIGHTDRAW-055 | (rom.exe) `light` and `darkness` create no sprite. They write the terrain brightness plane directly, four vertices per covered cell, and the removal recomputes those vertices. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-LIGHTLEVEL-056 | (rom.exe) A higher brightness byte is darker: `light`'s 0 is the brightest row of the shading table and `darkness`'s 0x50 is a half-brightness row. | High / Medium | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-UNITLIGHT-057 | (rom.exe) The same per-cell mask drives a second, per-frame lighting input: the light level units are drawn at. | High | ● active (amended, partially retracted) | [EXP-0176](../experiments/EXP-0176-area-draw/), [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| MAGIC-WALLFIRE-058 | (rom.exe) `wall_of_fire` is the singular overlay spell, in four separate ways. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| MAGIC-MARK-059 | (rom.exe) A lasting effect draws on its actor through an array of 8-byte mark records the client unit owns, and the unit draw walks that array twice, once before its own sprite and once after. | High | ✔ promoted | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-MARK-060 | (rom.exe) The mark set is not stored: it is re-derived every rebuild from a list of kind-and-countdown pairs the simulation opens and closes with two dedicated messages. | High / Medium | ✔ promoted | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-MARK-061 | (rom.exe) Ten spells put a mark on an actor and eighteen put none; the split is a 45-byte table in `.text`, not a data column. | High | ✔ promoted | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-PROT-062 | (rom.exe) The four Protections share one builder and differ only by a placement table; together they form a diamond of half-width 6 pixels. | High | ✔ promoted | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-SHIELD-063 | (rom.exe) `shield` is the one effect that draws a figure enclosing the actor rather than a mark beside it. | Medium | ● active (superseded) | [EXP-0177](../experiments/EXP-0177-effect-marks/), [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-BLESS-064 | (rom.exe) `bless` and `curse` each place twenty marks on a rotating circle, and the two rotate in opposite senses. | Medium | ✔ promoted | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-CLOUD-065 | (rom.exe) `poison_cloud` puts one mark above the actor; `heal` and `drain_life` append no mark and instead carry the previous rebuild's records forward. | Medium | ● active (partially retracted) | [EXP-0177](../experiments/EXP-0177-effect-marks/), [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-ACTOR-066 | (rom.exe) `stone_curse` and `invisibility` change the actor's own sprite instead of adding a mark, and the unit draw tests for them by kind directly. | High / Medium | ✔ promoted (amended, superseded) | [EXP-0177](../experiments/EXP-0177-effect-marks/) |
| MAGIC-ACTGATE-079 | (rom.exe) An actor carrying spell 20 runs no arm of the order switch, and the refusal is one byte above the choice of action; it has one bounded escape. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MAGIC-ACTKEY-080 | (rom.exe) The refusal is keyed to the spell id, and to nothing about the effect. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MAGIC-MASKREAD-081 | (rom.exe) The reader set of `actor+0x144` is nine routines, and `MAGIC-ATTACH-016`'s enumeration names six of them. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MAGIC-AIBIT-082 | (rom.exe) Two AI decisions read bit 20, neither of them gates an action, and the second is one case of a general no-restack rule. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MAGIC-EFFLOOK-083 | (rom.exe) The effect list at the drawable has one reader, `R0615`, and its complete call-site set is fourteen addresses over six routines, all in the presentation layer. | High | ● active | [EXP-0182](../experiments/EXP-0182-effect-at-actor/) |
| MAGIC-STONEDRAW-084 | (rom.exe) The `stone_curse` arm of the unit draw has two parts: it replaces the frame index with the facing, and it replaces the shade table with a greyscale one. | High / Medium | ● active | [EXP-0182](../experiments/EXP-0182-effect-at-actor/) |
| MAGIC-INVISOWN-085 | (rom.exe) The per-player bit that gates the invisibility sprite is an ownership test: an `invisibility` element hides the unit from every client except its owner's. | High / Medium | ● active | [EXP-0182](../experiments/EXP-0182-effect-at-actor/) |
| MAGIC-089 | (rom.exe) Heal and Drain Life carry an eight-output particle cohort with a constant, tile-scaled vertical step. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-090 | (rom.exe) Heal and Drain Life append a 3–5-particle elliptical cohort on each rebuild whose unsigned countdown is at least 8. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-091 | (rom.exe) Shield Component A is a sampled midpoint circle transformed into two ordered kind-44 records per source point. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-092 | (rom.exe) Shield Component B is a second, independently rasterised pair of circles appended after all Component A records. | High | ● active (amended, partially retracted) | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-093 | (rom.exe) Meteor Storm owns placement on the simulation side: 32 stages at effect ticks `0,3,...,93`, two ordered bounded-RNG calls per stage, and one stationary projectile per accepted cell. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| MAGIC-ITEMTRAIN-116 | A weapon-borne spell receives no immediate cast award; training comes from later event producers, and cadence follows accepted targets and damage ticks rather than releases. | High | ● active | [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-ITEMKILL-117 | Item-spell delayed-kill credit is not an award for release; it consumes the victim's surviving attribution state. | High / Medium | ● active (amended, partially retracted) | [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-ATTRGATE-118 | PointEffect and AreaEffect attribution are two different post-payload programmes; definition loss clears prior credit, and their null-owner outcomes differ. | High / Medium | ● active | [EXP-0191](../experiments/EXP-0191-itemcast-training/) |
| MAGIC-CADENCE-126 | A cast adds the Spell row's `Complication Level` only when state is `0x0d` or `0x0e` and `actor+0x64` still holds the Spell at recovery; retained order casts do, weapon-diverted caster casts do not. | High / Medium | ● active | [EXP-0234](../experiments/EXP-0234-action-cadence/) |
| MAGIC-CADENCE-127 | Insufficient mana writes no failed-cast recovery; after a completed cast leaves `actor+0x136 = 1`, a retained refusal attempts admission on three actor ticks, skips one, then repeats. | High / Unknown | ● active | [EXP-0234](../experiments/EXP-0234-action-cadence/) |

### MAGIC-CAST-028

**(rom.exe) Four `projectiles.reg` sheets are also drawn as per-unit status overlays, by flag bits and not by a cast.** `R0607` takes a flag word and, for four bits, looks a projectile record up by a **constant byte offset** into the global at `L02822` and blits it through `vt+0x18` at the unit's position less the registry's `Width`/`Height` halves: bit `0x8` -> `+0x3c` = id **15** (`firewall`), bit `0x80` -> `+0x5c` = id **23** (`freeze`), bit `0x100` -> `+0x64` = id **25** (`poison`), bit `0x80000` -> `+0xbc` = id **47** (`wall`). Each takes its phase from `R0376(unit+0xa70/2 + x*y, 0, 0)` reduced modulo the record's `Phases` -- a hash of the unit's own position, so two afflicted units in the same frame are at different phases and one unit's overlay does not flicker as it moves. The `firewall` arm is the odd one: its phase is `hash % 5 + 3`, restricted to five of that sheet's eleven frames. `0x80` and `0x100` are mutually exclusive in the routine's structure (`if`/`else`), and both are gated on a second argument alongside the bit. This is a **second consumer** of the same table that `MAGIC-PIC-026`'s `2*spellId + 8` addresses, reached without any message and without a projectile object, and it is why four of the effect sheets have to be readable as a loop rather than as a one-shot animation. **G2:** the four ids are immediates in `.text`, so the sheet a status effect shows is an **engine** limit -- re-pointing it means editing `rom.exe`, not `projectiles.reg` -- while the art behind each id is a data file and free. **RETRACTED IN PART by [EXP-0176] (`MAGIC-OVERLAYART-051`): the subject is wrong. The routine's two call sites are both in the map draw and the flag word is the client's per-cell AREA-EFFECT mask at `view+0x9f0`, whose bits are `1 << spellId` for spell ids 3, 7, 8 and 19. Nothing about a unit reaches the routine, and the phase multiplies the cell's SCREEN position within the visible window, not a unit's position, so the two consequences the row drew -- that two afflicted units are at different phases and that one unit's overlay does not flicker as it moves -- do not follow. Everything else in this row stands: the four bits, the four record ids, the constant-offset lookup, the phase expression, the `firewall` exception and the `0x80`/`0x100` exclusivity. The row's own Unknown, which flag word this is and which effect sets each bit, is closed.**

**Confidence.** High (one routine read whole, every id an immediate displacement in a listed instruction, and the phase expression read from the same listing) / Unknown (which simulation flag word this is and which effect sets each bit: the routine's caller was not read, so the mapping bit -> named spell effect is not established here)

**Original status.** ● partially retracted

**Evidence.** [EXP-0140](../experiments/EXP-0140-cast-art-drawn/), [EXP-0176](../experiments/EXP-0176-area-draw/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-CASTANIM-029

**A cast always animates the caster, and the even picture is what makes it do so.** Every cast reaches one of two senders from `MAGIC-CAST-003`'s cast routine `R0268`, and `EnumRefs callto:` gives each exactly one call site: `L03024` to `R0617` (aimed at a unit) and `L03025` to `R0618` (aimed at a place), the `MAGIC-TARGET-017` fork. Both write `msg+0xc = 2*spellId + 8`, which is **even for every spell id**, so `MAGIC-PIC-027`'s parity guard `L03011`, the bit-0 isolation, takes the even branch on every cast. That branch is not a no-op. `L05322` reads `msg+0xa`, the **caster's** runtime id (`L05306`, `L05323`), finds it in the world hash at `+0x9b8`, and then: `L05324` stores `classRecord+0x68` at the 32-bit field `+0xa0`, `L05325` stores the byte 0x8 at `+0x84` and `L02566` stores the 32-bit zero at `+0x94`. `ANIM-STATE-002` fixes action code **8** as the cast arm of `R0548`'s 8-entry table; `ANIM-CLOCK-001` fixes `+0x94` as the action clock and `+0xa0` as the ticks remaining. The class record is `units.reg`: the arm reaches it as `R1063(L05326, unit+0x20)`, and `R0548` reads the same array directly at `L05327` reads the global `[L02113]` and `L02637` its entry at the index, so `+0x68` is `REG-UNITS-049`'s **expanded Attack timeline length**. The whole branch is skipped when `unit+0xa0 != 0` (`L05328`) and when the class record's `+0x1c` `AttackPhases` is 0 (`L05329`). **A cast therefore plays the caster's own Attack run, once, one frame per game tick, for `len(Attack)` ticks** — the cast arm sets `+0x70 = +0x94` with no modulus (`L02500`) and `+0x94` increments once per tick at `L05330`, while the common tail decrements `+0xa0` at `L02481` and the action ends when it reaches 0 (`L05331`). Corpus, 34 classes: expanded Attack lengths 8 to 28 (`REG-UNITS-049`). The melee attack takes the same two fields through a different message: `R0245` to `R0549` to `R1064` sends opcode `0x72`, whose arm at `L02511`/`L05332` writes the same `+0xa0` and `+0x84 = 7`, shoot. **G2:** the run length is a registry field and the action code an immediate, so lengthening a cast animation is a `units.reg` edit and re-pointing which run a cast plays is a `rom.exe` edit

**Confidence.** High (every hop is a named instruction in a listing dumped over an address range; the registry's identity is settled by a second, independent reader of the same array that `REG-UNITS-049` and `ANIM-STATE-002` already interpret, and the two senders' call-site counts come from `EnumRefs callto:`, which also enumerates `.rdata` dwords holding the address) / Medium (that no other message can reach the same fields first: the guard `unit+0xa0 != 0` makes the outcome order-dependent between the `0x72` and `0x86` arms, and no probe here established which arrives first when both are sent in one tick)

**Original status.** ● active

**Evidence.** [EXP-0163](../experiments/EXP-0163-cast-picture/)

### MAGIC-CASTTICK-030

**Correction:** Application-at-windup wording is narrowed by MAGIC-DELIVERY-170 and MAGIC-CASTCLOCK-171: the call prepares payloads; a delivery2 transport must expire before its child enters application.  **The tick a spell is applied, in units a consumer can implement: tick `actor+0x134` of the action, counting the tick that started it as 0.** The actor tick `R0037` dispatches on `actor+0x54` through an 8-bit index table at `L01775` into `L01776` (`L05333`, `L01797`) and then on the sub-phase `actor+0x58`. Sub-phase **0** starts the action: either the attack start (`L05237`, a call of `R0245`), or `MAGIC-AUTOCAST-020`'s diversion to state `0xd`; in both cases the arm ends at `L05334`, which stores 0x5 at the 32-bit field `+0x58`, and loads the countdown — `L05335`/`L04379` `actor+0x6c = actor+0x134` for a cast, `L02726`..`L04380` `actor+0x6c = actor+0x134 + <out-param of R0245>` for a melee attack. Sub-phase **5** is one decrement per tick: `L05336`/`L05337` read the byte `actor+0x6c` and decrement it by 1, and `L05338` jumps to `L05339` while it is non-zero. On the tick it reaches **zero** the arm falls through to exactly one application — `L00006` (a call of `R0001`) for a weapon strike, or `L00007` (a call of `R0002`, unit) / `L00008` (a call of `R0003`, point) for a cast — and then sets `L05340 actor+0x58 = 7`, the recovery. **A melee swing and a spell application land on the same tick of the same counter; they are the two branches of one `if`.** The wind-up `actor+0x134` and the recovery `actor+0x135` are bytes with defaults **8** and **4**: `L05341`/`L05342`, `L05343`/`L05344` and `L05345`/`L05346` are the only immediate writes, and `L04639`/`L04640` overwrite them from elements **12** and **13** of an equipped item's own array (`R0476(item+0x3c + 8, 12` \| `13)`, `-1` meaning leave the default). The sum is what the simulation tells the client: `L05347`..`L05348` puts `actor+0x134 + actor+0x135` in `msg+0xd` of the `0x72` message and `L02469` stores `globalClock + that` in `actor+0x138`, `ANIM-CLOCK-001`'s end-of-run tick. **The client does not use it for the animation.** Neither the `0x72` arm nor the `0x86` arm reads `msg+0xd`; both take the run length from the class record instead (`MAGIC-CASTANIM-029`). Inside the client's own arms the visible moments are separate numbers again: the attack arm calls `vt+0x64` on the tick `+0x94 == classRecord+0x100` `AttackDelay` (`L05349`, `L02693`) and the cast arm calls `vt+0x5c` on the tick `+0x94 == classRecord+0xfc` `ShootDelay` (`L05350`, `L02723`), except when `+0xa4 == 0x3c`, which fires it at tick 0 (`L02720`..`L02722`). **The original therefore runs two clocks: the simulation applies the effect after `actor+0x134` ticks, and the client spawns the visible projectile after `ShootDelay` ticks, and nothing in the image equates the two.** A consumer that wants swing and application to coincide has to choose which one to follow; the simulation's is the one that changes hashed state

**Confidence.** High for the simulation clock (the sub-phase machine is read whole from one listing, every branch displacement in the block resolves inside it, and the `+0x134`/`+0x135` write census is complete under the fields' own byte width from `EnumRefs disp:134` and `disp:135`) / High for the client's two delay reads (named instructions in `R0548`, with `REG-UNITS-049` supplying the column names) / Medium for the negative half, that `msg+0xd` is unread on the client: both arms were read whole, but only those two arms, so a third consumer of the field elsewhere in the dispatcher is not excluded

**Original status.** ● partially retracted

**Evidence.** [EXP-0163](../experiments/EXP-0163-cast-picture/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### MAGIC-BURST-031

**`MAGIC-PIC-027`'s Unknown is answerable: the `Distribution system` column decides the burst picture, and 10 of 28 shipped spells reach it.** The deferred sender `R0638` has exactly one caller, `L03060` inside `R0639` (`EnumRefs callto:R0638` — 1 hit, 1 owner, 0 orphan), and `R0639` has exactly one caller, `L05351` inside `R0629`. `R0629` is slot `+0x18` of vtable **`L05352`**, which `R0652` installs (`L05353` stores the vtable L05352 into the object's first field) — `MAGIC-SHAPE-008`'s **AreaEffect**, allocated with the 0x50-byte size supplied at `L05354`. The sibling **PointEffect** (`R1046`, `L05355` storing the vtable L05356 into the first field, the 0x4c-byte size supplied at `L05357`) carries `R1065` in the same slot. `MAGIC-SHAPE-008` already fixes the fork: `getParam(row, 8)`, title 9 `Distribution system`, at `L05102` — value **1** builds the PointEffect, anything else the AreaEffect. **The spells that put a burst picture on the map are exactly those whose `Distribution system` is not 1.** Shipped EN `Data.bin`, `evidence/spell-effect-class.csv`: **18 PointEffect, 10 AreaEffect** — `fire_ball`(2), `wall_of_fire`(3), `fire_sacrifice`(4), `freezing_cloud`(7), `poison_cloud`(8), `acid_stream`(9), `light`(12), `darkness`(17), `wall_of_earth`(19), `meteor_storm`(21). Independent agreement: `MAGIC-PIC-026` reports `2*id + 9` is a defined `projectiles.reg` row for spells **2 3 4 7 8 9 19 21** — eight of exactly these ten, the two without art being `light` and `darkness`, whose visible effect is a lighting change. `R0639` paints a **pattern of cells**, not one: it switches on `this+0x0c`, the spell id (`MAGIC-ATTACH-016`), with its own arms for **4** `fire_sacrifice`, **9** `acid_stream` and **0x15** `meteor_storm` and a default for the rest, and for each cell it recomputes `L05286` (`2*id + 9`) and stores it before sending, so the odd id is written once per cell per ring. Rings advance a stage counter `this+0x4b` and end at `L05358`, which stores the byte 1 at `+0x40`. **The 11 `+8` stamps `MAGIC-PIC-026` found inside `R0003` never reach the wire on this path**: the only reader of that field is `R0638` at `L05313`, and its only caller overwrites the field with `2*id + 9` immediately before every call

**Confidence.** High for the routing (three single-caller enumerations by `EnumRefs callto:`, which also lists the `.rdata` dwords holding each address, plus a vtable dump in which the two classes differ in the decisive slot, plus each constructor's own vtable store; the column and its two branches are `MAGIC-SHAPE-008`'s already-High reading) / Medium for the 18/10 split, a corpus count over one root's `Data.bin` — what raises the AreaEffect set above corpus agreement is that a second, disjoint instrument, `MAGIC-PIC-026`'s scan of the shipped `projectiles.reg`, independently selects 8 of the same 10 / Medium for the last clause: the `+0x0e` reader census is a byte-, word- and `MOVZX`-form scan of the whole `.text` (34 hits, 18 owners, 0 orphan) whose hits outside these two classes were not each attributed to their object

**Original status.** ● active

**Evidence.** [EXP-0163](../experiments/EXP-0163-cast-picture/)

### MAGIC-SENDER-032

**Two corrections to `MAGIC-PIC-027`'s sender enumeration, one of which changes what a consumer must build.** (1) `L05304` is not a sender. It is the message class's **constructor** `R0626`, which installs vtable `L05359` and default-initialises `+0x9 = 0x86`, `+0xc = 0`, `+0xd = 0`, `+0xe = 0` (`L05360`..`L05361`); the opcode byte is a default, not a send. (2) `L03055` is not inside `R0638`. It is inside a **separate routine `R0635`**, and that routine is reached far more than the four `MAGIC-PIC-027` read: `EnumRefs callto:R0635` returns **11 call sites over 6 owners**, all but one in the effect module (`R1066`, `R1067`, `R0637`, `R0630`, `R0636` six times, `R0131`), against 1 caller for `R0638`. `R0635(effect, flag)` branches on `effect+0x0c`: **== 2** sends opcode `0x86` at the effect's own cell with `L05362`/`L05363`, a copy of the byte at `+0xe` to `msg+0xc` — the same picture field, forwarded the same way — and `msg+0xf = 0x16`; anything else sends opcode **`0x87`**, a message `MAGIC-PIC-027` did not enumerate, carrying a `(2r+1)` square cell mask built from `effect+0x49` into `msg+0x10` onward (`L05364`, `L05365`, `L05366`, `L05367`). **Consequence:** a consumer that builds only the four senders of `MAGIC-PIC-027` will draw a cast's start and an area effect's rings and still miss whatever these eleven sites emit. This row states what was measured and does not name `R0635`'s argument class

**Confidence.** Medium. The two corrections are each a named instruction in a dumped listing and the call-site counts come from `EnumRefs callto:`, so the **negative** half — that `L05304` sends nothing and that `MAGIC-PIC-027` mis-attributed `L03055` — is settled. The **positive** half is not: the class of `R0635`'s argument was not identified, so `effect+0x0c == 2` is not resolved to a spell property, and opcode `0x87`'s client arm was not read. Both are named as open in the experiment

**Original status.** ● active

**Evidence.** [EXP-0163](../experiments/EXP-0163-cast-picture/)

### MAGIC-CASTSPAWN-033

**(rom.exe) The object a cast puts on the map is built by the CASTER, not by the message, and its flight length comes from a hard-coded switch on the picture id -- non-zero for 7 picture ids and zero for the other 44.** `EnumRefs callto:R0609` plus `imm:14c` enumerate every construction of the `0x14c`-byte `CProjectile`: 6 call hits over 4 owners, 0 orphan -- the cast spawner `R0620`, the unit shot `R0603`, three arms of the client dispatcher `R0509` (`L03002`, `L03003`, `L03004`) and the savegame loader `R0099`. `R0620` is `CUnit` `vt+0x5c` and `CAirUnit` `vt+0x5c` (`.rdata` slots `L05368`, `L05369`; both vtables dumped, 0 of 33 slots without a function), and inside the unit action driver that slot has exactly two call sites, both in the **cast arm**: the 8-entry jump table at `L02482` sends action 8 to `L02712`, and `L02723`, a virtual call of the slot at offset 0x5c, fires on the tick `actionphase == classRecord+0xfc` `ShootDelay` (`L05370`, `L05350`) while `L02722` fires on tick 0 when `actionspell == 0x3c` (`L02720`). A whole-image `EnumRefs re:` sweep for a call through a register-based `+0x5c` slot returns **32 hits over 26 owners, 1 orphan**, of which these are the only two in `R0548`; the other 30 sit in the runtime and MFC address families and their object class was not established here. The spawner: `L05371`/`L05372`, a copy of the 32-bit field `+0xa4` of the source to `+0x20` of the new object -- **`picture = actionspell`, which `MAGIC-CASTANIM-029` fixes as `2*spellId + 8`**; the spawn point is the class's own muzzle table `classRecord+0xec` indexed by `(casterDir - 8) & 0xe` minus `classRecord+0x34`, times 8, added to caster `+0x58` (`L05373`..`L05374`), used only when `classRecord+0xf0` is non-zero and the picture is not `0x3c`, otherwise the class bounding-box centre `(class+0x8c - class+0x84)/2` (`L05375`); `actionx`/`actiony`/`actionz` are the target's `+0x58`/`+0x5c`/`+0x10` when `actiontarget` is set and the caster's own `+0x88`/`+0x8c`/`+0x90` otherwise; `action = 1` (`L05376`), `actionphase = 0` (`L05377`). The flight length is `L05378` bounding `picture - 10` at 0x32, `L05379` reading the index byte from the table at `L05380` and `L05381` jumping through the table at `L05382` with stride 4 -- **51 index bytes, an 8-entry jump table, six distinct arms**, read out of the PE by virtual address (`tools/castflight -mode tables`): `dist/200` for picture 10, `dist/384` for 12, `1` for 20 and 30, `13` for 34 and 36, `21` for 60, and `0` for the remaining 44. **A zero means the driver `R0558` returns finished at `L02817` before reaching its own switch at `L02834`, so those cast objects never execute a driver arm.** `picture == 0x3c` alone allocates a **second** `CProjectile`, copy-constructed at `L05383`/`L03009` and placed at the caster's own bounding-box centre -- teleport draws two sprites. **G2:** the 51-byte table and its arms are `.text`, so which spells throw something and for how long is an **engine** limit; the muzzle table, `ShootDelay` and the sheet are `units.reg` / `projectiles.reg` data and free

**Confidence.** High (every quantity is an immediate, a displacement or a compare bound in a listing dumped over an address range; the two switch tables are read from the PE by section walk with no disassembler in the path; the construction census and the `vt+0x5c` call-site sweep are both `EnumRefs`, which attributes orphan hits rather than dropping them; and the competing model "the switch is also reached by a unit's shot" is excluded by the vtable dump giving `vt+0x58` a different routine that computes `dist/200` inline) / Medium (the muzzle-table reading, whose `classRecord+0xec` / `+0xf0` field names are not established here)

**Amended.** MAGIC-261 names the class fields and exact fallback arithmetic.
The second Teleport object's caster-centre placement clause is partially
retracted in `retracted.md`: MAGIC-265 establishes destination-derived raw
`+08/+0c` and source-derived cached `+28/+2c` at construction. Native first
draw remains Unknown. Construction, timing and lifetime clauses stand.

**Original status.** ● active

**Evidence.** [EXP-0167](../experiments/EXP-0167-spell-art/)

### MAGIC-BURSTLIFE-034

**(rom.exe) A burst is stationary, and its sender chooses its lifetime.** The odd-picture client arm copies its own position into `actionx` and `actiony`, clears `actiontarget`, and loads `msg+0xf` into `actionsegments`; the projectile driver therefore divides zero on all three axes. The staged sender uses **16** ticks, raised to **18** for `acid_stream`; the other sender uses **22**. `acid` has 9 phases and gets 18 ticks, and `fireexpl` has 11 phases and gets 22, so those two bespoke lifetimes equal `2 * Phases`. The staged sender also runs on stage ticks **0, 3, 6, ...**: `L05384` reloads the stage timer with 2, but `L05385`…`L05386` consumes old values 2 and 1 on separate return ticks before old value 0 permits the next stage. The prior two-tick clause is corrected by EXP-0175. The client plays sound id `500 + picture`, at volume `(10000 - screenDistance)/100`, except for picture 51. **G2:** lifetimes 16, 18 and 22 and the stage-timer reload are `.text` immediates. Changing them changes engine code, not a burst sheet or `Data.bin`.

**Confidence.** High for stationarity, the three lifetimes and the corrected stage sequence. The object fields and immediates are named instructions; `tools/ringgeom` asserts the complete counter sequence on both preserved roots.

**Original status.** ● active (cadence clause amended)

**Evidence.** [EXP-0167](../experiments/EXP-0167-spell-art/), [EXP-0175](../experiments/EXP-0175-ring-geometry/)

**Amended.** The table ledger carried the status "● active (cadence clause amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-DELIVER-035

**(rom.exe) The simulation and the client each compute a cast projectile's flight length, by different rules, and on the normal path the client's wins.** The 4th argument of both cast senders `R0617` and `R0618` reaches `msg+0xf`, which every object-building client arm reads as `actionsegments` (`MAGIC-BURSTLIFE-034`), so it is a **flight time**; `MAGIC-PIC-027` named it `dir`, and no arm read here reads it as a direction. Its value is computed in `MAGIC-CAST-003`'s cast routine `R0268`: `L05387` initialises it to **0**; `L05388` (a call of R0476 with the argument 0x5) reads `Data.bin` parameter **5**, `Delivery System`, and only when it is **2** does `L05060`, a signed divide by a 32-bit value, sets it to the caster-to-target distance divided by parameter **7**; `L05061` stores 0x5 at the frame local −0x10 and so overrides it to **5** for spell ids **13** and **14**. On a normal cast this number is never used: the senders write an even picture, the client's parity guard takes the caster-animation branch, and the projectile is built later by the caster's own spawner, which computes its own length from the `L05380` switch (`MAGIC-CASTSPAWN-033`). The simulation's number is used only when the source has no client-side runtime id, which rewrites the opcode to `0x8b` (`L05308`, `L03021`) -- an arm that builds the object directly, with an even picture and no parity guard, and takes `actionsegments` from `msg+0xf` (`L05389`). **Two disjoint instruments select the same four spells.** `Delivery System == 2` holds for exactly `fire_arrow`(1), `fire_ball`(2), `lightning`(13) and `prismatic_spray`(14) over the shipped 28 rows on both roots; the client's own 51-byte table gives a travelling or ramped arm to exactly the four cast pictures 10, 12, 34, 36 -- those same four spells -- and `ANIM-CAST-027` independently gives those four picture ids the only four special draw arms in `R0556` (a smoke trail for 10 and 12, a polyline for 34 and 36). **G2:** the column is data and the table is `.text`, so making a fifth spell throw something is a `Data.bin` edit **and** a `rom.exe` edit, and doing only the first changes what the simulation sends without changing what the client draws

**Confidence.** High for the field's role and the two computations (named instructions in dumped listings, and the negative half -- that the even branch never reads `msg+0xf` -- rests on that branch being read whole rather than sampled) / Medium for the four-spell agreement: the `Delivery System` half is a corpus census over two roots, and what lifts it is that a second, disjoint instrument in `.text` selects the same four

**Original status.** ● active

**Evidence.** [EXP-0167](../experiments/EXP-0167-spell-art/)

### MAGIC-AREATICK-036

**(rom.exe) The `AreaEffect` per-tick entry is vtable slot `+0x18`, and `effect+0x08` selects one of three exclusive per-tick modes.** `EnumRefs callto:R0629` returns 1 hit / 0 owners / 0 orphan: the `.rdata` dword at `L05390`, which is `L05352 + 0x18`, `MAGIC-BURST-031`'s AreaEffect vtable. The driver is `R0641`: it walks the effect list, clears a dead caster reference at `L05391` (`effect+0x3c` non-null and `caster+0x14 == 0` gives `effect+0x3c = 0`), calls `effect->vtable[+0x18]()` at `L03082`, and removes the effect from the list at `L05392` when `effect+0x40` is non-zero. It runs once per simulation tick: `R0193` increments the tick counter `session+0x4` at `L05393`…`L05394` and calls `R0426(session+0x14)` at `L01866`, which calls `R0641(that+0x4)` at `L03066`. The effect is added to that list at `L05395` and `L03053`. `R0629` reads two one-instruction predicates: `R0631` is `effect+0x08 & 2` (`L05396`) and `R0632` is `effect+0x08 & 1` (`L05397`). Bit 1 set selects `R0639`, the ring walker; bit 0 set selects the cloud branch; neither selects `R0630`, a single-pass blast. The builder `R0003` writes `+0x08` three times: `L05398` sets 0 in the constructor, `L05399` (storing 1 at `+0x8`) runs when `+0x4c > 0` after `Area Effect Duaration` is loaded, and `L05400` (storing 2 at `+0x8`) runs when `Distribution system == 5` together with `L05401` (storing the 16-bit zero at `+0x4c`). The distribution-5 arm runs second and overwrites both fields, so **`meteor_storm`'s shipped `Area Effect Duaration = 10` is overwritten to 0 and has no effect**; the same holds for any authored row combining the two columns. Over both roots the ten AreaEffect spells split **1 blast** (`fire_ball`), **3 ring** (`fire_sacrifice`, `acid_stream`, `meteor_storm`), **6 cloud** (`wall_of_fire`, `freezing_cloud`, `poison_cloud`, `light`, `darkness`, `wall_of_earth`)

**Confidence.** High for the routing and the mode selection: the `callto:` enumeration names the single `.rdata` slot that reaches `R0629`, the tick chain is three named call instructions, and each of the three `+0x08` writes and both predicates are named instructions in dumped listings. The competing model, that the six unread `R0635` owners of `MAGIC-SENDER-032` are the tick slots of six distinct `AreaEffect` subclasses, is excluded: `callto:` on each returns a single caller inside this one module / Medium for the 1/3/6 split, which is a corpus count of two `Data.bin` columns over both roots

**Original status.** ● active

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/)

### MAGIC-AREAPULSE-037

**(rom.exe) A cloud pulses once every 16 ticks, a staged effect runs its first stage immediately and later stages every 3 ticks, and a blast runs once. `effect+0x4c` is a counter with mode-specific meaning.** In cloud mode `+0x4c` starts at `(AreaEffectDuaration << 4) + (power << 4)/10`. The first driver call paints without decrement or pulse. On a later call, a positive old counter decrements and the new value pulses when divisible by 16, **including zero**; cleanup occurs on the next call. For unchanged initial positive `V0`, there are `ceil(V0/16)` pulse-branch entries and `V0 + 2` calls including registration and teardown. Thus `V0=240` has 15 pulse entries, last at relative tick 240 and teardown at 241; these are not per-target damage counts. `MAGIC-CLOUDCLOCK-154` corrects the former positive-only endpoint and counts; the historical cloud examples are retracted separately. In staged mode the builder sets `+0x4c = 0`. `R0639` reads the old word, decrements it, returns when the old value was positive, and otherwise reloads **2** (`L05385`…`L05402`). The resulting sequence is stage ticks **0, 3, 6, ...**, not every 2 ticks. The earlier reading treated the reload value as the interval and omitted the tick on which the counter reaches 0. `effect+0x4b` increments after the complete cell loop, and `+0x40 = 1` when the incremented stage reaches the arm limit. Blast mode does not read `+0x4c` and sets `+0x40 = 1` in its only pass.

**Confidence.** High. The counter load, decrement, old-value test, reload, stage increment and completion comparison are named instructions in the complete routine listing. `tools/ringgeom` independently asserts their byte strings on both preserved roots and re-executes the counter sequence. The cloud endpoint is re-read from complete original instructions and conditionally executed in `MAGIC-CLOUDCLOCK-154`; the first-paint and zero-pulse calls are distinguished explicitly.

**Original status.** ● active (staged cadence and cloud endpoint amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0175](../experiments/EXP-0175-ring-geometry/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/)

**Amended.** The table ledger carried the status "● active (staged cadence and cloud endpoint amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-AREAAPPLY-038

**(rom.exe) An area effect invokes its inner effect's per-target virtual entry; the dynamic class determines direct damage versus ordinary attachment.** `R0267(unit)` is the single per-unit apply of all three modes: `EnumRefs callto:R0267` returns **7 hits / 3 owners / 0 orphan** — `R0629` at `L05403` (cloud pulse), `R0639` at `L05404`, `L05405` and `L05406` (ring), `R0630` at `L05407`, `L05408` and `L05409` (blast). Its body returns on a null unit (`L05410`), returns for spell id `0x13` (`L05411`, a compare with 0x13), and otherwise pushes the unit and calls `[areaEffect+0x44]->vtable[+0x3c]` at `L05412`…`L05413`. `areaEffect+0x44` is the inner `Effect` the spell's own apply arm built and passed to the AreaEffect constructor (`L05414` pushes it, `L03047` calls `R0652`); the builder then writes `innerEffect+0x44 = caster` at `L05415`. The two **base Effect** constructors install vtable **`L02981`** — `L05416` in the plain `R1049` and `L05417` in the template `R1047`, both storing the vtable L02981 into the first field — and slot `+0x3c` of that vtable is **`R0612`**, the attach routine `MAGIC-ATTACH-016` decodes. `EnumRefs callto:R0612` returns **1 hit / 0 owners / 0 orphan**, the `.rdata` dword at `L05418 = L02981 + 0x3c`, This identifies the base-Effect attach slot, not every inner effect's virtual target. For the base-Effect route, the applied magnitude is the inner `Effect`'s `+0x40`, written by the spell's own apply arm (`MAGIC-CEIL-013`), and no per-tick magnitude exists anywhere in the area module; whether an application damages once or is stored and re-run is the inner effect's mode byte `+0x3d` (`MAGIC-EFFMODE-009` — an untimed effect is applied and never stored, a `continuous` one is stored and re-runs `vt+0x40` every 8th tick from `R0673`); and `MAGIC-ATTACH-016` governs timed base-Effect re-application, not every area damage object. **The broader virtual-target and universal non-stacking clauses are REFUTED by `UNIT-AREADIRECT-072`: `Effect_DirectDamage` overwrites its vtable with `L02980`, whose `+0x3c` points directly to `R0564`; the read damage-ring producers carry that object in area `+0x44`. Untimed base effects are also applied without being stored.** Two per-spell exceptions in the same routine: for spell id **2** (`fire_ball`), `L05419`…`L05420` copy-construct the inner effect (`R1068`), call `unit->vtable[+0x1c]` twice, multiply the two results, divide the copy's byte fields at `+0x5b` and `+0x5c` by that square, apply the copy through `R0564` and destroy it at `R1069`; for spell id **0xc** (`light`), `L05421`, a compare with 0xc, jumps to the epilogue before the kill-credit block

**Confidence.** High for the routing: three `EnumRefs` enumerations reported with hit, owner and orphan counts, both constructor vtable stores, and a vtable dump in which `+0x3c` of `L02981` is `R0612` while `+0x3c` of both `L05356` and `L05352` is null, **This does not establish an exclusive target class for `L05413`; the derived DirectDamage table was omitted (amended by `UNIT-AREADIRECT-072`).** The magnitude clause is a negative scoped to what was read: the whole per-tick path was dumped and contains no arithmetic on a damage value / **Medium** for the purpose of the `fire_ball` divide. The divisor is the square of one `vt+0x1c` result and all three modes iterate per cell, which reads as a normalisation for a multi-cell unit being found in several cells, but nothing read here establishes that a large unit's pointer is stored in every cell record it covers. *(Discriminated 2026-08-15 by [EXP-0174]: it is stored in every one of its `n × n` cells (`TERR-FOOTPRINT-147`), the blast reads all three occupant slots per cell, and `R0564` applies without a stacking test — so the divide is a **per-cell normalisation**, `MAGIC-FIREDIV-047`.)*

**Original status.** ● active (amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/), [EXP-0271](../experiments/EXP-0271-multicell-area/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-AREACELL-039

**(rom.exe) Three different cell sets, and the simulated set is not the mask sent on the wire.** `Distribution system == 3` (`fire_ball`, `freezing_cloud`, `poison_cloud`, `light`, `darkness`) paints through `R1067`, which walks `dx, dy` over `[-r, +r]` with `r = effect+0x49` and filters on `abs(dx) + abs(dy) <= r + 1` (`L05422`, `L05423`, `L05424`): a **diamond of Manhattan radius `r + 1`**, not the `(2r+1)` square `R0635` transmits on opcode `0x87` (`MAGIC-SENDER-032`). `Distribution system == 4` (`wall_of_fire`, `wall_of_earth`) paints through `R1066`, which builds its pattern with `R1070(L05425, L05426, effect+0x4a)`; both tables are `struct { int count; int dx[20] at +0x04; int dy[20] at +0x54; }`, table A holding **10** cells as a 5x2 block and table B **9** as a doubled anti-diagonal, and the 8-way switch at `L05427` picks a table, optionally swaps the `dx`/`dy` arrays and negates one or both multipliers. **A wall occupies 10 cells on an axis and 9 diagonally whatever its `Radius, Length/2` column**, while its radius still bounds the pulse scan and the removal loop. `effect+0x4a` is written at `L05428` as `(map->R0089(caster, targetCell) & 0xff) >> 5`, one of 8 directions. The blast mode paints no cell at all: `R0630` walks the `(2r+1)` square once and ends at `L03056`. The pulse scan at `L05429`…`L05430` walks the `(2r+1)` square regardless of which shape was painted, and per cell reads the effect stored in this spell's own map layer (`MAGIC-MAPLAYER-040`), so a cloud pulses only where it painted. **Which occupants are reached differs by mode.** The map cell record carries three occupant dwords in its payload, read by four accessors that share one hash lookup and differ only in the field returned: `R0035` and `R1071` return `+0x4` (`L05431`, `L05432`), `R1072` returns `+0x8` (`L05433`), `R1073` returns `+0xc` (`L05434`). The cloud pulse reads **only `+0x4`** (`L05435`); ring (`L05436`, `L05437`, `L05438`) and blast (`L05439`, `L05440`, `L05441`) read **all three**

**Confidence.** High for the shapes and the occupant readings: every loop bound, every filter and every accessor is a named instruction in a dumped listing, the two wall tables were read out of the shipped `ROM.EXE` bytes, and the four accessors were disassembled and differ only in the returned displacement. The diamond-versus-square conclusion is discriminating because both shapes exist in the same effect and the transmitted one is a different routine / Medium for the per-spell attribution, which rests on the `Distribution system` column read from `Data.bin` over both roots. **The occupant count is corrected**: the record carries **four** occupant dwords, not three — the fourth is `+0x10`, the sack slot, with its own three accessors ([EXP-0174], `TERR-CELLREC-146`, `claims/retracted.md`). The three named here are read correctly and are correctly attributed; what they are *keyed by* is the actor's movement domain

**Original status.** ● active (amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-MAPLAYER-040

**(rom.exe) The map carries exactly six area-effect layers, one per spell id, and five conflict rules delete one layer when another is laid.** `R1074` maps a spell id to a layer index as `idx = spellId - 3` over the valid range `0..0x10`, with a byte table at `L05442` and a jump table at `L05443`; it returns **0..5 for spell ids 3, 7, 8, 19, 12, 17** and for every other id logs `"MapLayer() call - Invalid Area Effect"` (`L05444`) and returns 0. The slots live in the map's per-cell record at payload dwords 5..10: `R1075` writes `scratch[0x14 + idx*4] = effect` at `L05445` and then counts occupied slots in a 6-iteration loop (`L05446`, a loop count of 0x6) into the payload byte at `+0x2`; `R1076` walks the same 6 slots (`L05447`, a count of 0x6). **The six spells with a layer are exactly the six cloud-mode spells** of `MAGIC-AREATICK-036`; the four AreaEffect spells without one — `fire_ball`, `fire_sacrifice`, `acid_stream`, `meteor_storm` — are exactly the four that never register a cell, so the log arm is unreachable on shipped data. `R0636` is the conflict rule, dispatched on `spellId - 2` over `0..0xf`: `fire_ball` and `wall_of_fire` remove the poison layer and then the freezing layer, `freezing_cloud` removes the fire layer, `poison_cloud` removes itself when a fire layer is present, and `light` and `darkness` each remove both the light and the darkness layer. Each removal calls `map->R1077(other, x, y)` and, for one of the pair, `R0635(other, 0)`

**Confidence.** High for the index function and the slot count: the byte and jump tables were read out of the shipped `ROM.EXE` bytes, and the 6 is an immediate in two independent loops in different modules. The claim that the layered set equals the cloud-mode set is corroborated by two disjoint instruments, the mode coming from two `Data.bin` columns and the layer from a `.text` jump table / Medium for the five conflict rules, which are read from one dispatch table and its arms and were not exercised

**Original status.** ● active

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/)

### MAGIC-AREAEND-041

**(rom.exe) All three modes end at the same byte; cloud teardown removes its owned cell layers, whose helper can restore saved terrain/flags.** `R0641` reaps an effect whose `effect+0x40` is non-zero (`L05392`). Cloud mode reaches it through `R0637`, which walks the `(2r+1)` square calling `map->R1077(effect, x, y)` per cell, then `R0635(effect, 0)` and `+0x48 = 0` at `L05448`, after which the caller sets `+0x40 = 1` at `L05449`. Blast mode sets it at `L03056` at the end of its single pass. Ring mode sets it at `L05358` when the stage counter `+0x4b` reaches the per-arm limit (`MAGIC-BURST-031`). The complete removal helper can restore the cell record's saved terrain/flag bytes when that record becomes empty (`MAGIC-CLOUDEND-163`), so the former blanket no-terrain-write clause is retracted. This cloud cleanup does not revoke effects already attached to units by `MAGIC-AREAAPPLY-038`; the tested Poison attachment expires on its separate `+0x42` countdown (`MAGIC-POISONREFRESH-158`, `MAGIC-FIREPOISON-160`). `R1075`'s record-creation tail sets only bit `0x20` of the plane at `map + 0x10000` (`L05450`) and snapshots the existing terrain and flag bytes

**Confidence.** High for the three end sites and the removal walk, each a named instruction in a dumped listing. The prior negative omitted the complete helper `R1078`; its empty-record restoration is now read through both returns. Poison non-revocation is conditionally exercised, while native target transitions remain outside that instrument

**Original status.** ● active (terrain-cleanup clause amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/)

**Amended.** The table ledger carried the status "● active (terrain-cleanup clause amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-WALLEARTH-042

**(rom.exe) `wall_of_earth` applies nothing to any unit and refuses occupied cells, and no passability write was found for it.** Three things are unique to spell id `0x13` in this module. (1) `R0267`, the per-unit apply, returns immediately for it at `L05411`…`L05451`, so the spell touches no unit at any power. This agrees with its inner effect: `MAGIC-EFFECT-015` records that `wall_of_earth` uses the plain constructor, which is `R1049`, and that constructor zeroes `+0x0c`, `+0x3c`, `+0x3d`, `+0x40` and `+0x44` at `L05297`…`L05452`, so the effect carries no spell id, no mode and no magnitude. (2) `R1079`, the per-cell add, tests `map->R0035(x + (y << 8))` at `L05453` for id `0x13` only and **skips the cell when occupant slot `+0x4` is non-null**; it also skips `R0636` and `R1080` on that arm (`L05454`…`L05455`). (3) Its apply arm at `L05129` allocates the effect and calls `R1049` with no arguments, writing no magnitude, no spell id and no picture. **No passability write was found.** What was searched: every caller of `R1074` (7 hits, 4 owners — the add `R1075`, the remove `R1078`, and the queries `R1081` and `R1082`) and every caller of each layer accessor (`callto:R1081` 1 owner, `callto:R1083` 2, `callto:R1075` 1, `callto:R1078` 1). The four read sites are the damage pulse, the sender `R0635`, the conflict rule `R0636` and the client duration sync `R1076`; none is a mover or a path search. The area module makes exactly **two** terrain-plane writes and `wall_of_earth` reaches only one of them: `R1075` sets bit `0x20` of the plane at `map + 0x10000` at `L05450`, the "this cell has a spell record" bit that all four cell accessors gate on, and `R1080` sets bit `0x10` of **both** planes, `map + 0x10000` at `L05456` and `map + 0x20000` at `L05457`. `EnumRefs callto:R1080` returns **2 hits / 2 owners / 0 orphan**, `L05458` inside `R1079` and `L05459` inside `R0630`, and the first is on the arm `wall_of_earth` skips. **What makes `wall_of_earth` block movement is not established.** The census covers the map-layer slots and the area module's own plane writes, not every path a mover could take

**Confidence.** High for the three uniquenesses, each a named comparison or call in a dumped listing, with the zeroing constructor read instruction by instruction and independently agreeing with `MAGIC-EFFECT-015` / Medium for the negative: the enumeration of layer readers is complete and reported with hit and owner counts, but it bounds only the layer slots, so absence of a passability effect is **not** established and is named open. *(Answered 2026-08-15 by [EXP-0174], `MAGIC-WALLBLOCK-045`: every statement in this row stands — the passability write is not in the area module and no layer reader is a mover. `R1075` stores the effect into cell-record payload `+0x20` and then calls **`R0453`**, which derives plane bits 0 and 2 from that slot on both planes; the search blocks on its own mask. The one map call this row identifies, `R1084`, is the whole path. **The recompute and its `+0x20` arm were already promoted** in `formats/terrain/format.md` from EXP-0070 and EXP-0085, under the open item *"what each of the six holds is still open"* — so this row's question was one hop from an answer this repository had published, in a file the round did not cross-read.)*

**Original status.** ● active (amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0174](../experiments/EXP-0174-area-movement/)

**Amended.** The table ledger carried the status "● active (amended)".

### MAGIC-AREATARGET-043

*(amended by EXP-0191: the `vt+0x20` call is movement domain, not liveness; the branch and its attribution-only scope stand, and `MAGIC-ATTRGATE-118` replaces the label.)* **(rom.exe) No per-tick routine of an area effect consults diplomacy, so an area effect pulses on the caster's own party.** `EnumRefs disp:a9c4`, the diplomacy table displacement `MAGIC-TARGET-017` identifies, returns **32 hits / 14 distinct owners / 0 orphan**: `R0131` (7), `R0430`, `R0061` (2), `R0132` (6), `R0128` (3), `R0127` (2), `R0810`, `R0003`, `R0021`, `R0124`, `R0262` (4), `R1085`, `R0400`, `R0293`. **None of the per-tick routines is among them**: `R0629`, `R1086`, `R0639`, `R1066`, `R1067`, `R1079`, `R0637`, `R0267`, `R0630`, `R0636`. The single hit inside the 6 KB apply `R0003` is `L05219`, `heal`'s arm, which `MAGIC-TARGET-017` already reports. `MAGIC-TARGET-017` also reports `imm:a9c4` returning 0 image-wide, which covers the addressing form a displacement sweep would miss. `MAGIC-TARGET-017`'s conclusion therefore extends from the cast path to the tick path: every unit standing in an area effect's cells is applied to, whatever its player. The attribution block's other gates are the null check and `unit->vtable[+0x20]() <= 0` at `L05460`…`L05461`; the virtual returns movement domain `actor+0x4a`, so it does not suppress credit merely because the payload was lethal

**Confidence.** High. The enumeration is reported with its mode, its hit, owner and orphan counts, and its blind spot named, and it reproduces `MAGIC-TARGET-017`'s own 32/14 figure on the same query. The negative is a complete owner list checked against a complete list of the routines on the per-tick path, every one of which was dumped. The movement-domain correction is fixed independently by `TERR-MOVE-055` and EXP-0191's raw call-site assertions

**Original status.** ● active (amended)

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/), [EXP-0191](../experiments/EXP-0191-itemcast-training/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### MAGIC-LIGHTDARK-044

**(rom.exe) `light` and `darkness` are the same effect with opposite sign, and only `darkness` provokes its victims.** Both are cloud-mode spells with `Radius, Length/2 = 4` and `Area Effect Duaration = 20` on both roots. Their apply arms at `L05462` and `L05463` allocate the same effect, call the same template constructor `R1047(row + 0x1c)`, and differ in one field: `light` writes `effect+0x40 = power/30 + 1` at `L05464`…`L05465` and `darkness` writes `effect+0x40 = -1 - power/30` at `L05466`…`L05467`. `MAGIC-CEIL-013` names the same two expressions and the same field. At pulse time both reach `R0267` and apply through `[effect+0x44]->vtable[+0x3c]` like every other cloud spell (`MAGIC-AREAAPPLY-038`). They then diverge: `light` returns at `L05468` before the kill-credit block, while `darkness` continues to write `unit+0x40 = caster` and `unit+0x48 = innerEffect+0x0c` (`L05469`, `L05470`), the two fields `R0427` reads back as a unit's kill-credit fields at `L05471` and `L05472`, and then calls `[L00004]->R0226(caster, unit)` at `L00877`, which sets the victim's AI alert state at `victim+0x158`. **Nothing in either arm writes a lighting or visibility value**, and the two are the only AreaEffect spells with no `projectiles.reg` row (`MAGIC-BURST-031`), so what the client draws from map layers 4 and 5 is where their visible difference must live; that was not read

**Confidence.** High for the two magnitude expressions and the `light` early return, each a named instruction in a dumped listing, and for the shared constructor. The kill-credit field identity rests on `R0427` reading the same two displacements back / Medium for the alert clause, which names `R0226`'s writes but does not establish that the state it sets is what makes a unit attack. **Unknown**: the client rendering of layers 4 and 5

**Original status.** ● active

**Evidence.** [EXP-0172](../experiments/EXP-0172-area-tick/)

### MAGIC-WALLBLOCK-045

**(rom.exe) `wall_of_earth` blocks movement through the ordinary terrain-block bits, and a routine outside the area module derives them from the cell record.** The wall's own arm of the per-cell add makes exactly one map call: `L05473`/`L05474` (a compare with 0x13 and a jump to `L05475` when not equal) select it, and it ends at `L05455`, an unconditional jump to `L05476`, after `L05477` calls R1084. `R1084(effect, x, y)` packs the coordinates and tail-calls `R1075(effect, (y<<8)+x)`, which resolves the area layer with `L05478` (a call of R1074), stores the effect at `L05445` (the payload slot `payload+0x14+4*layer`), and calls `R0453` on the same cell — `L05479` on the arm where the record existed, `L05480` on the arm that creates one. **`R0453` recomputes both plane bytes for one cell from that record.** That routine was already read end to end by EXP-0070 and EXP-0085 and its pseudocode is in `formats/terrain/format.md`, including the line `if payload+0x20: static \|= 5 ; dynamic \|= 5` and the note *"what each of the six holds is still open"*. **What is new here is which spell owns `+0x20`**, which is what turns that arm into an answer. Its last conditional arm is `L05481`..`L05482` (a read of the payload dword `+0x20`, a zero test and a jump on zero): when it is non-null the arm sets the bits of the constant 5 in the static plane (`L05483`, `L05484`, `L05485`) and into the dynamic plane (`L05486`, `L05487`, `L05488`) — **bits 0 and 2, the terrain-block and object-occupies bits `TERR-PASS-073` names**. Payload `+0x20` is area layer index 3, and `tools/areamove` reads `R1074`'s byte table at `L05442` and jump table at `L05443` out of the shipped image: index 3 is reached only from byte-table entry 16, i.e. **spell id 19, `wall_of_earth`**, and no other of the six layers has a passability arm. The search then blocks on the mover's own mask (`MOVE-DOM-025`): `0x41` carries bit 0, `0x44` carries bit 2, `0x82` carries neither — so **movement domains 1 and 2 are blocked and domain 3 is not**. Removal restores passability by the same route: `R1078` clears the slot at `L05489` (storing zero at `payload+0x14+4*layer`) and calls `R0453` at `L05490`. **None of the three candidate mechanisms is the mechanism**: the search reads neither the layer index nor any occupant slot, and the two bits the area module writes itself — bit 5 at `L05450` and bit 4 at `L05456`/`L05457` — are in no mover mask

**Confidence.** High (every instruction above is transcribed from full listings of five routines, `evidence/listing-plane-derive.txt` and `evidence/listing-layer-add-remove.txt`; `tools/areamove` re-asserts 95 of them as literal byte strings read through the PE section table with no disassembler, 0 mismatches on both roots, and derives the spell-to-slot mapping from the image's own tables. The reading is discriminating rather than merely consistent: of the six area layers exactly one has a passability arm, and it is the one the layer-index tables assign to spell id 19. **Attribution:** the recompute, its `+0x20` arm and the `OR 5` polarity are EXP-0070's and EXP-0085's, promoted into `formats/terrain/format.md`; this row adds the spell-to-slot link, the removal path and the mask consequence. `MAGIC-WALLEARTH-042` was published one hop from an answer this repository had already promoted, and the reason nobody made the hop is that the spec calls `+0x20` a payload displacement while the area module calls it a layer index — `docs/INSTRUMENT.md` rule 8's corollary, grep the offset rather than the id, applies to `formats/` as well as to the ledgers)
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Original status.** ● active

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

### MAGIC-AREACOST-046

**(rom.exe) Every occupied area layer multiplies its cell's movement cost byte by four, and the step-duration path divides it back by four once and stores the result.** In `R0453` the cost byte is first reset from the record's saved copy (`L05491`/`L05492` copy the first payload byte to `&cost[cell]`), so the multiplication is not cumulative across calls; then `L05493`/`L05494` (the slots at `payload+0x14` and a count of 6) walk the six layer slots and `L05495`, a left shift of the cost byte by 2, runs once per non-null slot — **that loop is EXP-0085's**, promoted into `formats/terrain/format.md`; what this row adds is the truncation it implies and the routine that divides it back. The shift is **8-bit**, so bits leave the byte: with the shipped cost range 6…16 (`TERR-COST-052`) one layer gives 24…64, two layers give 96…256 where 256 truncates to 0, and three layers truncate every shipped value to 0 or 128. On the consuming side `R1087(map, cell)` returns `cost[cell]` when static bit 5 is clear (`L05496` tests bit 0x20), and otherwise looks the record up and, when the layer-count byte `payload+0x2` is non-zero (`L05497` reads the byte at stack offset +0xe), executes `L05498` (a right shift by 2) and **`L05499` (a byte store), writing the divided value back into the cost plane**. `R1088` (`TERR-MOVE-056`) is its only caller, twice — `L05500` and `L05501` — and averages the two results (`L05502`/`L05503`: the two results added and halved), substituting 8 for a mean of 0 at `L05504`. So the route search (`MOVE-COST-002`, which reads the cost plane inline for a `movementType == 1` mover) sees the multiplied byte while the step-duration sees it divided once

**Confidence.** High for the instructions: both loops, both shifts, the store-back and the mean are named instructions in full listings, and `tools/areamove` asserts them as byte strings / **Unknown** for the net effect in a running session: whether repeated `R1087` calls can drive a layered cell's cost down before `R0453` restores it depends on a call ordering that was not established (the order is now read: `MOVE-084`, `MOVE-085`)

**Original status.** ● active

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

### MAGIC-FIREDIV-047

**(rom.exe) `fire_ball`'s footprint-squared divide is a per-cell normalisation, not a size penalty — because a unit is stored in every cell record it covers.** `MAGIC-AREAAPPLY-038` graded this Medium and named the missing input. `R0050(map, actor)` reads the footprint side once (`L05505`, a virtual call of the slot at offset 0x1c), then runs two nested loops bounded by that side (`L05506`, `L05507`) and calls `R0458` once per covered cell (`L05508`) with `x0 + inner` and `y0 + outer`, so a unit of side `n` occupies `n²` cell records (`TERR-FOOTPRINT-147`). `fire_ball` is blast mode: `R0630` walks the `(2r+1)` square once and reads all three occupant slots per cell (`L05439`, `L05440`, `L05441`), passing each to `R0267`, whose spell-id-2 arm divides the copied inner effect's `+0x5b` and `+0x5c` by the square of the target's own `vt+0x1c` and applies through `R0564`. **`R0564` is a direct application with no stacking test**: `L05509` tests the target through `vt+0x2c`, `L05510` computes damage through `vt+0x4c`, `L05511` (a 16-bit store at `+0x94`) writes back the reduced health. **The old claim that every other area application uses `R0612` and repeated visits are idempotent is REFUTED (`UNIT-AREADIRECT-072`).** Non-fireball DirectDamage objects also reach `R0564` through their own `+0x3c` slot; only the spell-2 copy/divide branch is exclusive in this local dispatcher. The unit footprint-squared division and partial-coverage result stand. They do not extend to rectangular Buildings, whose `vt+0x1c` returns 1 (`UNIT-AREAHP-073`). Two exact consequences: the division is integer on byte fields, so `n²` applications of `floor(m/n²)` total **at most** `m` and a magnitude below `n²` divides to 0; and a unit only partly inside the square is found fewer than `n²` times and takes proportionally less

**Confidence.** High (the `n²` storage is a two-loop walk with the side as both bounds, read end to end; `R0564` was read end to end and contains no stacking test; the discriminator is a present instruction pair, not an absence)

**Original status.** ● active (amended)

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/), [EXP-0271](../experiments/EXP-0271-multicell-area/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-RING-048

**(rom.exe) The three staged AreaEffect spells use three distinct cell generators.** All offsets are relative to the target cell and are visited in the written order. `fire_sacrifice` has two fixed stages: stage 0 is `(-1,1) (-1,0) (-1,-1) (0,1) (0,-1) (1,1) (1,0) (1,-1)`; stage 1 is `(-2,1) (-2,0) (-2,-1) (-1,2) (0,2) (1,2) (-1,-2) (0,-2) (1,-2) (2,1) (2,0) (2,-1)`. It ignores `effect+0x4a`. `acid_stream` has six stages and an orientation `o = (directionByte & 0xff) >> 5`. Define `A_s = [(x,s) for x=-s..s]` for `s=0..4`, `A_5=[]`, and `B_s = [(s-i,i) for i=0..s]` for `s=0..5`. Even orientations use A: `o=0:(+dx,-dy)`, `2:(+dy,+dx)`, `4:(-dx,+dy)`, `6:(-dy,+dx)`. Odd orientations use B: `o=1:(+dx,-dy)`, `3:(+dx,+dy)`, `5:(-dx,+dy)`, `7:(-dx,-dy)`. Thus the sixth stage is empty on even orientations and has six cells on odd orientations. `meteor_storm` has 32 stages of one cell each; every stage calls `R0861(5)` twice in x-then-y order and uses `(first-2, second-2)`, so each coordinate lies in `[-2,3]` and stages may repeat. The two fixed families use `0xa4`-byte records `{count:i32, dx[20]:i32, dy[20]:i32}`; terminal counts are 2 and 6, while Meteor uses immediate 32 and one mutable record of count 1. Construction sets `stage=0`, `timer=0`, mode 2 and the orientation. Stages run on ticks `0,3,6,...`. Each target-plus-offset coordinate is truncated to bytes. The position helper accepts it exactly when `8 <= x <= mapWidth-9` and `8 <= y <= mapHeight-9`; on rejection it clamps the stored position to that range but the walker neither sends nor applies the cell. An accepted cell sends picture `2*spellId+9` once and applies to occupant slots `+0x4`, `+0x8`, `+0xc` in that order. The walker increments the byte stage counter after the full list and sets `+0x40=1` when `stage >= limit`; the common driver reaps it on the same tick. Staged effects register no map layer and perform no separate cleanup. Shipped `Data.bin` selects exactly ids 4, 9 and 21 through `Distribution system = 5` on both roots. **G2:** the 20-cell table capacity, 2/6/32 stage bounds, 8-orientation transform, `[-2,3]` random domain, byte stage and map coordinates, and timer width are engine tables, immediates or fields. Changing the shipped fixed geometry changes `ROM.EXE`, not `Data.bin`; an externalized replacement table can preserve every shipped file byte. More than 20 cells, 255 stages, 8 directions or single-byte coordinates requires a wider engine representation.

**Confidence.** High for the consumer, tables, transforms, random order, cell gate, terminal conditions and cleanup: `R0639`, `R0290`, both table helpers and the common driver were read end to end; 49 byte assertions and every table value reproduce identically from both preserved `ROM.EXE` roots. Medium for the three-row shipped attribution, which is a `Data.bin` corpus measurement over both roots. The competing shared-generator model is excluded by two separate fixed table families and a random arm with no stage table.

**Original status.** ✔ promoted

**Evidence.** [EXP-0175](../experiments/EXP-0175-ring-geometry/)

### MAGIC-BOLTGATE-069

**(rom.exe) Exactly two picture ids produce a drawn path, and the routine that produces it is one virtual slot whose whole body sits behind a two-comparison gate.** `R1089` is `CProjectile` vtable slot `+0x50`; `EnumRefs callto:R1089` returns 1 hit, the `.rdata` slot `L05512`, 0 orphan. The driver `R0558` calls it on itself once per tick, on the path taken when `actionsegments != 0` (`L05513`, a jump to `L05514` on non-zero) and the object's `action` byte is 1 (`L05515`/`L05516`: a decrement of the action byte and a jump to `L05517` on non-zero, and the spawner sets `action = 1` at `L05376`): `L05518`..`L05519` (a virtual call of the slot at offset 0x50 on the object). Its entry is `L05520` (a read of the 32-bit field `+0x20`, the picture), then `L05521`/`L05522` (the picture biased by −0x22, jumping to `L05523` on zero) and `L05524`/`L05525` (a further bias of −2, jumping to `L05526` on non-zero) — so only pictures **34** and **36** have a body, and every other picture returns having touched nothing. The geometry generator `R1090` has `EnumRefs callto:R1090` = 2 hits, 1 owner, 0 orphan, both inside that routine (`L05527`, `L05528`), so nothing else in the image can reach a generated path. The draw switch agrees independently: the byte table at `L02868` and the jump table at `L05529`, indexed by `picture - 7` and bounded at 0x35 by `L02866`, give the two polyline arms `L05530` and `L05531` to exactly one picture each, 34 and 36, while 48 of its 54 indices take the default arm. By `MAGIC-PIC-026`'s `2*spellId + 8` those are spell **13 `Lightning`** and spell **14 `Prismatic Spray`**, two of `MAGIC-DELIVER-035`'s four `Delivery System == 2` spells and the two whose flight-length arm is the constant 13 (`MAGIC-CASTSPAWN-033`). **G2:** both comparisons are immediates in `.text` and both draw arms are hard-coded ids, so giving a third spell a drawn path is a `rom.exe` edit; re-pointing either spell's art at another `projectiles.reg` row moves the art and loses the path

**Confidence.** High. The gate is five named instructions re-read from the raw image bytes at their own addresses (`tools/boltpath -mode asserts`, 74 of 74 holding on three roots); the reachability rests on `EnumRefs callto:`, which counts `.rdata` slots and orphan hits rather than dropping them, and returns a single slot for each of the driver, the producer and the draw, all three in the `CProjectile` vtable `L02587`; the draw switch is read out of the PE by section walk with no disassembler in the path. The competing model — that the two fields are a base-class pair whose projectile use is incidental — is excluded by the generator's two-caller census

**Original status.** ✔ promoted

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/)

### MAGIC-BOLTSHAPE-070
The picture gate remains MAGIC-BOLTGATE-069.
The complete builder first creates the bounded random source walk, inserts
midpoint knots, evaluates quadratic samples, rejects an excessive sampled
ordinate, then rotates/truncates accepted samples. MAGIC-275 through
MAGIC-278 and MAGIC-282 define those local passes and their exact arithmetic.
The final point list is therefore interpolated from random knots; the former
unqualified "not an interpolation" clause is withdrawn. The stored -0.5 is
used for midpoint insertion before the quadratic helper, not inside its fit.

**Confidence.** High for the random source structure and the enumerated local
pass order. The former Medium parametric reading is superseded by the exact
operand/control-word recipe in MAGIC-275, MAGIC-277 and MAGIC-282. The
endpoint-order Unknown is closed by MAGIC-275's two producer captures.
Native entry CW and cache/frame scheduling remain Unknown in those claims.

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/),
[EXP-0504](../experiments/EXP-0504-bolt-figure/).

**Amended.** The unqualified no-interpolation clause and the -0.5 smoothing
attribution are partially retracted in claims/retracted.md. The random-walk
source, gate, rotation and whole-figure retry stand; the former parametric
reading and endpoint Unknown are superseded by MAGIC-275 through MAGIC-278
and MAGIC-282.

### MAGIC-BOLTLIST-071

*(EXP-0243 closes and retracts the fill Unknown below: opcode `0x8a/0x8c` refills `+0xac/+0xb0` with selected ids; `L05035` is unrelated.)* **(rom.exe) The point list is rebuilt on every driver tick and it does not accumulate: picture 34 replaces it, picture 36 clears it and concatenates one generated set per victim id.** The list is a `CArray` of 8-byte records at object `+0x110`: `R1093(newSize, growBy)` is its `SetSize`, and with `newSize == 0` it frees the buffer and zeroes `[this+0x4]`, `[this+0x8]` and `[this+0xc]` (`L05554` to `L05555`), which fixes `+0x114` as the data pointer, `+0x118` as the count and `+0x11c` as the capacity. **Picture 34** reads `actiontarget` (`L05556` reads the 16-bit field `+0x86`), resolves it through the world map at `+0x9b8` (`L05557`, a call of L05558), passes the constant tag 0x22 (`L05559`), generates at `L05528`, then calls `SetSize(generatedCount, -1)` at `L05560` and copies the new records over the array; when the target does not resolve it jumps to `L05561` and leaves the array as it was. **Picture 36** clears first — `L05562`/`L05563` (a call of R1093 with the argument 0x0) — then loops over `frame+0xb0` entries of the word array at `frame+0xac` (`L05564` reads the 16-bit entry at the loop index), resolving each id through the same hash, tagging each generated set with the loop index modulo 7 (`L05565`..`L05566`: a signed divide by 0x7 whose remainder is passed on), and appending: `L05567` reads the current count at `+0x8`, `L05568` adds the new one, `L05569` calls R1093 to grow the array. That victim-id array reaches the projectile from the caster, copied word by word by the spawner: `L05570` reads the count `+0xb0`, `L05571` forms the address `+0xa8` of the destination array, `L05572` calls L05573, `L05574` reads the source array pointer `+0xac` and `L05575` reads one 16-bit entry. Since the generator always spans the whole source-to-target segment and no argument carries a partial length, every point of the figure is present from the first tick

**Confidence.** High for the array identity, the clear, the replace, the concatenate and the spawner copy, each a named instruction in a dumped listing and re-read from the raw image bytes by `tools/boltpath`; `R1093` was read end to end, which is what makes `newSize == 0` a clear rather than a resize. **EXP-0243 closure:** opcode `0x8a/0x8c` refills caster or temporary-source `+0xac/+0xb0` with selected victim ids in order, so the selected-victim count is the link count; the former `L05035` lead is unrelated

**Original status.** ✔ promoted

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/)

**Amended.** The table ledger carried the status "✔ promoted". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-BOLTSTILL-072

**(rom.exe) A Lightning or Prismatic Spray projectile keeps its raw position; normal caster construction supplies a 13-call countdown rather than distance.** The driver computes a per-axis step before its switch — `L05576`, `L05577` and `L02830`, three signed divides by `actionsegments` — into stack slots, and **only the default arm applies it**: `L05578`..`L05579` read the three step values and store them as the 32-bit fields `+0x8`, `+0x10` and `+0xc`, the value read from stack slot `0x20` going to `+0x10` and the other two going to `+0x8` and `+0xc` respectively. Pictures 34 and 36 take arm `L05580` instead — a bias of −1, `L02882` (a bound at 0xc) and `L05581` (a jump through the table at `L02881` with stride 4) — whose thirteen entries write only `+0x70`, the sheet frame. So on this action-1 path the object's `x`, `y` and `z` keep the spawner's values, while normal caster construction's `MAGIC-CASTSPAWN-033` flight-length arm `L05582` gives it `actionsegments = 13` and `L05583`/`L02836` (a decrement stored to the 32-bit field `+0xa0`) counts that down once per tick. The target is still re-read every tick before the switch (`L02824`, `L02825`, `L02826`), and the producer resolves it again itself, so the figure follows a moving target by being regenerated onto it rather than by travelling toward it. **G2:** the arm assignment is the `.text` switch table at `L02879` and the length is the `.text` switch table at `L05380`, so making these two spells travel is an engine edit

**Confidence.** High. The three signed divides, the four position writes of the default arm, the arm's own three instructions and the decrement are each their own instruction in a dumped listing and re-read from the raw image bytes; the arm assignment is read out of the PE by section walk (`tools/castflight -mode tables`), which excludes the competing model that the 34/36 arm falls through into the default one

**Original status.** ✔ promoted

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/)

**Amended.** The universal 13-tick clause is narrowed to normal caster
construction in claims/retracted.md. MAGIC-281 supplies direct-message,
source-cell and saved-state lifetimes. Stationary raw-coordinate stores stand.

### MAGIC-TRAIL-073

**(rom.exe) The trail behind a Fire Arrow or a Fire Ball is a queue of at most six PAST positions, and its oldest entry is dropped one at a time.** The array is the `CObArray` the constructor builds at object `+0x138` — `L05584`/`L05585` (the address `+0x138` passed to a call of L05586), and `0x138 + 0x14 = 0x14c`, the object size — so `+0x13c` is its data pointer and `+0x140` its count, the pair `ANIM-CAST-027` reports the default draw arm walking. The driver saves the object's position at entry, before any arm has moved it: `L05514` reads the 32-bit field `+0xc` and `L05587` the field `+0x8`. After the switch, and only for pictures 10 and 12 (`L05588`..`L05589`: the picture `+0x20` biased by −0xa and a further −0x2), it runs `L05590` (the 32-bit count `+0x140` compared with 0x6), and when the count has reached six calls `RemoveAt(0, 1)` — `L05591`..`L05592` (a call of L05593 with the arguments 0x1 and 0x0) — before appending the saved position packed as `(y shl 16) or x` through `SetAtGrow` at `L05594` (a call of L05595). The two arms at `L05590` and `L05596` are the same code twice, one per picture. This trail therefore accumulates and its entries expire individually, oldest first, which is the opposite of `MAGIC-BOLTLIST-071`'s wholesale rebuild, and six is a hard bound in `.text`. **G2:** the bound, the packing and the per-picture selection are immediates in `.text`

**Confidence.** High. Every quantity is an immediate or a displacement in a dumped listing, all re-read from the raw image bytes by `tools/boltpath`; the saved-before-moved ordering is fixed by the two loads sitting above the switch and the append sitting below it, in one linear routine

**Original status.** ✔ promoted

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/)

### MAGIC-BOLTEND-074

**(rom.exe) Nothing is created when a Lightning or Prismatic Spray figure ends, and the whole figure ceases with the object.** On the tick `actionsegments` is already 0 the driver returns before its switch and before the producer call: `L02817`..`L05513` (the 32-bit field `+0xa0` compared with zero, jumping to `L05514` on non-zero) and `L02818` (the return value set to 0) — the finished return `ANIM-PROJ-025` identifies. That path allocates nothing, sends nothing and writes nothing but `+0x74 = 1` for picture 13, which is neither of these two. The point list is a member array of the object, so it goes with it, and nothing removes points before that. Neither spell has a second picture to leave behind: `2*spellId + 9` gives 35 for Lightning and 37 for Prismatic Spray, and `tools/castflight -mode map` finds no `projectiles.reg` row defined for either id on either preserved root, which extends `MAGIC-PIC-027`'s undefined-picture list by measurement. The damage is on the simulation's own clock and is not tied to this figure (`MAGIC-CASTTICK-030`)

**Confidence.** High for the finished path and for the absence of a burst row, the first from a dumped listing re-read as raw bytes and the second a corpus census over both roots. **Unknown:** which routine deletes the object once the driver has returned 0. The driver's own return is the only removal instruction read here; the reaper was not located, and an `EnumRefs re:` sweep for a call through a register-based `+0x3c` slot returns 209 hits over 61 owners, none of which was classified

**Original status.** ✔ promoted

**Evidence.** [EXP-0178](../experiments/EXP-0178-bolt-path/)

### MAGIC-AREADRAW-049

**(rom.exe) The three `AreaEffect` tick modes create three different kinds of drawable, and only one of them creates one object per covered cell.** `R0629` forks to the staged arm (`L05351`), the first-tick paint arm (`L05597` to `L05598`), the end arm (`L05599` to `L03058`) or the blast arm (`L03079`). A **cloud** creates no object at all: its paint arm registers each cell in the map layer and then sends exactly one message. A **blast** sends exactly one message and the client builds one transient object at the effect's own position (`L05600`) with a 22-tick life (`L02927`). A **staged** effect sends one message per accepted cell inside its cell loop (`L03060`), each building one transient object with a 16-tick life, 18 for `acid_stream` (`L02925`, `L02926`). Two senders are involved and they carry different opcodes: `R0638` writes `0x86` (`L05303`) and `R0635` writes `0x86` for spell id 2 and `0x87` for every other spell (`L05601`, `L03055`, `L05364`). Shipped `Data.bin` puts exactly one spell in the blast mode (`fire_ball`), three in the staged mode and six in the cloud mode, identically on both roots. **G2:** the mode fork, the two opcodes and the three lifetimes are `.text` immediates; which mode a spell takes is the `Distribution system` and `Area Effect Duaration` columns and is data

**Confidence.** High: the fork, both senders and all three arms were read end to end, and the send counts are a byte census rather than a reading. `tools/areadraw` scans each routine's range for direct `rel32` calls and finds 0 sends in the tick routine, 1 in each cloud generator, 1 in the end arm, 1 in the blast and 1 in the staged arm; 165 byte assertions and 11 censuses reproduce identically on both roots

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-OVERLAY-050

**(rom.exe) What a cloud puts on the client is a per-cell bitmask, not a set of objects.** Opcode `0x87` carries the bounding box `(centre-radius, 2*radius+1)` at `msg+0x0b..0x0e`, an add/remove flag at `msg+0x0f`, the picture at `msg+0x0a`, and a cell bitmap from `msg+0x10` whose bit index is `rowY*(2*radius+1) + colX` (`L05602`, `L05603`). The client arm at `L05604` converts the picture back to the spell id with `picture/2 - 4` (`L05605`) and calls `R0483`, which walks the box, tests bit `y*w + x` (`L05606`, `L05607`) and keys a hash at `view+0x9f0` by `x` OR `y << 8` (`L05608`, `L05609`). An add ORs `1 << spellId` into the cell's dword (`L05610`); a remove ANDs it out (`L05611`) and drops the entry when the dword reaches zero (`L05612`). Nothing else is stored per cell: no object, no counter, no phase. `EnumRefs callto:R0483` returns exactly one call site, this arm, with 0 hits in orphan or undisassembled code. **G2:** the mask is a dword, so a cell can carry at most 32 concurrent area layers, and the spell id **is** the bit index, an engine limit a spell id above 31 would break silently

**Confidence.** High: the sender, the client arm and the consumer were read end to end; the one-caller statement is a whole-image `callto:` sweep; the rival reading, a sheet index rather than a spell id, is excluded because the six values `picture/2-4` produces are exactly the six spell ids that reach this path and are the constants both consumers test

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-OVERLAYART-051

**(rom.exe) The retained overlay's art is four immediates, and `R0607` is the per-cell area-effect draw, not a per-unit status overlay.** `EnumRefs callto:R0607` returns exactly two call sites, both inside the map draw `R0379` (`L02991`, `L02995`), each reached after looking the drawn cell up in the `view+0x9f0` mask (`L05613`, `L05614`) and each gated on that cell's fog word being exactly `0xc000` (`L05615`, `L05616`). The routine's four arms test mask bit 3, 7, 8 and 19 and index projectile records **15, 23, 25 and 47** as literal immediates (`L05617`, `L05618`, `L05619`, `L05620`) into the array at `L02822`. Those are `2*spellId + 9` for `wall_of_fire`, `freezing_cloud`, `poison_cloud` and `wall_of_earth`, but the value is baked, not computed: the picture id the message carried was already discarded by `MAGIC-OVERLAY-050`'s conversion. The fourth argument splits the arms into two passes. Bits 3 and 19 draw when it is non-zero, bits 7 and 8 when it is zero (`L05621`, `L05622`, `L05623`, `L05624`). Arms 2 and 3 exit the routine after drawing while arm 1 falls through, so `freezing_cloud` and `poison_cloud` cannot both draw at one cell but `wall_of_fire` and `wall_of_earth` can. This **corrects `MAGIC-CAST-028`**, which read the same routine as a per-unit status overlay phased by a unit's position. **G2:** the four sheet ids and the four mask bits are `.text` immediates, so which four of the six cloud spells have a sprite at all is an engine limit; the art behind each id is a data file and free

**Confidence.** High: the routine was read whole, every id and bit is an immediate in a listed instruction re-asserted as bytes, and the two-caller enumeration is a whole-image `callto:` sweep with 0 hits in orphan code

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-CLOUDLIFE-052

**(rom.exe) A cloud's overlay stands for the whole life of the effect, and the damage pulse draws nothing.** The paint arm runs once, on the first tick, guarded by the byte `effect+0x48` which it then sets to 1 (`L05625`, `L05626`, `L05627`). The damage pulse at `abs(timer) mod 16 == 0` (`L05628`) walks the cells and calls only the per-unit apply; a byte census of `R0629..R1094` finds **0** direct calls to either sender. Teardown is the end arm `R0637`, reached when the duration word `effect+0x4c` falls to 0: it clears the map layer cell by cell (`L05629`) and then sends **one** message with flag 0 (`L05630`, `L05631`), whose bitmap is the cells at which that spell's layer slot is now empty (`L05632`). The overlay is therefore added by one message and removed by one message, per effect and not per cell. Nothing else clears it: `imm:9f0` returns ten owners over the whole image, of which the only writer is `R0483` and the readers are the two map-draw sites, the per-frame light-grid rebuild `R0608` and the overview redraw `R0484`

**Confidence.** High: the guard, the pulse, the end arm and the sender polarity are byte-asserted, and the absence of a send in the pulse is a call census over the routine's own byte range rather than a reading

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-CELLSET-053

**(rom.exe) The cells drawn are the cells the map layer actually holds, re-read from the map, not the nominal geometry.** `R0635` does not transmit the generator's cell list. It walks the `(2*radius+1)` square and, per cell, calls `R1082`, which returns the effect registered in that spell's own layer slot at that cell, and sets the bit on that answer (`L05633` under flag 1, `L05632` under flag 0). The registration `R1079` skips a cell when the map's structure-slot accessor returns non-zero (`L05634`) and the first byte of the record the cell accessor returns has bit 2 set (`L05635`), and for `wall_of_earth` also any cell already holding a ground actor (`L05636`). A skipped cell is therefore absent from the bitmap and is never drawn. This amends `MAGIC-AREACELL-039`'s parenthesis: the sender does transmit a `(2r+1)` square, but as a bounding box carrying a per-cell mask, so the drawn set is the diamond or the fixed table minus its skips, not the square

**Confidence.** High: the per-cell query and both polarities are byte-asserted, and the registration's two skip conditions are each their own instruction in a listing read end to end

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-AREARADIUS-054

**(rom.exe) An `AreaEffect`'s radius is the shipped `Radius, Length/2` column read once at construction, not a power term, and the overlay bitmap is sized for the shipped values and no more.** `R0003` reads Spells column element 9 (`L05103`) and passes it to the constructor, which stores it at `effect+0x49` (`L05637`); `EnumRefs callto:R0652` returns exactly one construction site. `MAGIC-POWER-004`'s `min(power/20 + 2, 7)` is at `L05075`, inside the different routine `R0269`, and feeds Prismatic Spray's selected-victim count and sender. The shipped `Radius, Length/2` over the ten `AreaEffect` rows is 1, 2, 3 or 4 on both roots. The `0x87` sender zeroes exactly 12 bitmap bytes (`L05367`), which is 96 bits, and writes bit `y*(2r+1)+x` with no bound; `(2r+1)^2 <= 96` holds for `r <= 4` and fails from `r = 5`, and the largest shipped radius among the six cloud spells is 4. **G2:** raising a cloud spell's radius column above 4 writes bits outside the region the sender zeroes. What reaches the wire past 96 bits is not established: the message class's transmitted length is taken from its own `vt+0x10` on a `.bss` instance whose vtable is installed at runtime and was not read

**Confidence.** High for the radius source and the one construction site, both byte-asserted with a whole-image `callto:` sweep; High for the 12-byte zero and the unbounded bit write, which are their own instructions; Medium for the shipped radius range, which is a `Data.bin` corpus measurement over both roots

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-LIGHTDRAW-055

**(rom.exe) `light` and `darkness` create no sprite. They write the terrain brightness plane directly, four vertices per covered cell, and the removal recomputes those vertices.** Both are cloud-mode and lay the same per-cell mask as the other four, but no arm of the draw `R0607` tests bit 12 or bit 17, and the retained path never touches the projectile record array, which is why `MAGIC-PIC-026`'s observation that pictures 33 and 43 have no `projectiles.reg` row costs them nothing. Instead the add path of `R0483` has an arm for spell id 12 (`L05638`) and one for spell id 17 (`L05639`), both writing four bytes into `[[view+0x80] + 0x18]` at `p`, `p+1`, `p+stride`, `p+stride+1` with stride `view+0x84`: **light writes 0** (`L05640`, `L05641`, `L05642`, `L05643`) and **darkness writes 0x50** (`L05644`, `L05645`, `L05646`, `L05647`). That plane is `TERR-LIGHT-012`'s per-vertex brightness grid, the one both terrain blitters sample. On the remove the two ids share one arm that calls `R0468(x, y, 2, 2)` (`L05648`, `L05649`), the routine that recomputes the same plane from the terrain slope and the sky intensity bytes `L05650` and `L02293`. **G2:** the two values and the 2x2 vertex footprint are `.text` immediates; there is no data column behind either

**Confidence.** High: both arms and the removal call are byte-asserted, and the plane is identified by two independent routes, the field `[[view+0x80]+0x18]` and the routine the removal calls writing the same plane, with `TERR-LIGHT-012` having established it as the blitters' source at High

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-LIGHTLEVEL-056

**(rom.exe) A higher brightness byte is darker: `light`'s 0 is the brightest row of the shading table and `darkness`'s 0x50 is a half-brightness row.** `TERR-LIGHT-018` gives the terrain shading table 96 rows of 256 entries and `TERR-LIGHT-019` its per-entry transform `out = clamp(((palette + tint) * m)/32, 0, 255)` with `m` counting **down** from 96 to 1 as the row index rises. Level 0 is therefore `m = 96`, a multiplier of 3.0 that saturates most palette entries, and level `0x50` = 80 is `m = 16`, a multiplier of 0.5. `R0468` clamps its own output to `[0, 95]` (`L05651 = 95.0`, `L05652 = 0.0`), so both written values are inside the table. `TERR-LIGHT-029` measures the natural terrain level over the 38 shipped height grids as 30 to 70 at the default sun angle, so `light` is brighter than any vertex the shipped maps produce and `darkness` darker than any of them

**Confidence.** High for the polarity and the two multipliers, which follow from `TERR-LIGHT-019`'s transform read at High and the clamp constants read from the image. Medium for the clause that `darkness` is darker than any natural vertex: the shipped-corpus census is 30 to 70 but `TERR-LIGHT-029`'s proved per-axis bound is `[30, 89]`, so 80 is inside the range the engine could compute for some height field

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-UNITLIGHT-057

**(rom.exe) The same per-cell mask drives a second, per-frame lighting input: the light level units are drawn at.** `R0608` rebuilds `CMapView+0xb0`, which is `TERR-SPR-065`'s light level handed to the unit body pass `vt+0x28`, by first filling it with `[L05650] / 4` (`L05653`, `L05654`), the sky intensity `TERR-LIGHT-119` puts at 14 by day and 32 at night, giving 3 and 8; then walking the same `view+0x9f0` hash (`L05655`) and, per cell in the visible window, testing three bits: `0x1000` (`light`) sets the level to **0** (`L05656`), `0x20000` (`darkness`) sets it to **12** (`L05657`), and `0x8` (`wall_of_fire`) calls `R1095(x, y, 1, level)` (`L05658`), which splats a level over a radius-1 square. `light` and `darkness` therefore change how units standing inside them are drawn as well as the ground, on a different scale from the terrain plane, and `wall_of_fire` is a light source. The splat is gated on `[L05659] == 0`, which was not read

**Confidence.** High: the seed, the three bit tests, the two literal levels and the splat call are each byte-asserted, and `TERR-SPR-065` identifies `CMapView+0xb0` at High

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

**Amended.** The radius-1 square and L05659 stamp-gate clauses are partially retracted. MAGIC-271 gives the clipped distance-table footprint; MAGIC-273 separates the Lighting stamp gate from the Animation invalidation gate. The mask bits, literal unit levels and wall_of_fire source call stand. See claims/retracted.md.

### MAGIC-WALLFIRE-058

**(rom.exe) `wall_of_fire` is the singular overlay spell, in four separate ways.** First, it is the only spell id with an add arm in `R0483` that is not a lighting write: it skips a cell whose tile word already carries `0x2000` (`L05660`), skips one whose terrain type byte is 0 (`L05661`), requires that type's entry in the table at `L02099` to have `+0x44` other than `-1` (`L02100`, `L05662`), records the cell in a **second** hash at `view+0xa7c` (`L05663`) and then ORs `0x2000` into the tile word (`L05664`), so a cell only burns where the terrain type permits it. Second, it is the only draw arm whose frame law is not `mod Phases` (`ANIM-WALLFIREFRAME-033`). Third, it is the only arm of `R0607` that falls through instead of exiting, so it can be drawn together with `wall_of_earth` at one cell. Fourth, it is the only mask bit that appears in both the sprite draw and the per-unit light grid, that is, the only area effect that is also a light source. Its removal arm does none of this: it increments `view+0x3f6c` and nothing else (`L05665`), so what was written to the tile word and to the `view+0xa7c` hash is not undone on the path this experiment read

**Confidence.** High for the four singularities, each byte-asserted in a routine read end to end. The terrain type table `L02099` and the `view+0xa7c` hash's consumer were not read, so what the terrain-type test selects, and what un-does the tile-word bit, are open

**Original status.** ● active

**Evidence.** [EXP-0176](../experiments/EXP-0176-area-draw/)

### MAGIC-MARK-059

**(rom.exe) A lasting effect draws on its actor through an array of 8-byte mark records the client unit owns, and the unit draw walks that array twice, once before its own sprite and once after.** The array object is at `unit+0x110`: data pointer `unit+0x114`, count `unit+0x118`, capacity `unit+0x11c`. A record is `+0x00 i16 dx`, `+0x02 i16 dy`, `+0x04 i16 depth`, `+0x06 u8 record index`, `+0x07 u8 phase`. `R0552` walks it at `L05666` before dispatching the actor's own sprite and again at `L05667` after. The first pass draws a record only when `depth > 0` (`L05668`, a jump on less-or-equal), the second only when `depth <= 0` (`L05669`, a jump on greater); both otherwise compute the same position, `x = unit+0x60 - Width/2 + dx` and `y = unit+0x64 - unit+0x68 - unit+0x10 - Height/2 + dy - depth`, where `Width` and `Height` are `record+0x1c` and `record+0x20`. Art resolves through the same record array at `L02822` that `MAGIC-PIC-026` addresses, by the byte at `record+0x06` (`L05670`), with `R0604` loading the sheet when `record+0x38` is zero. Both passes blit through the sprite's `vt+0x18` with `(x, y, phase, 0, 0)`, taking `phase` from `record+0x07` (`L05671`, `L05672`). The vertical base every builder starts from is `R0586`, which returns the actor class record field `+0xd0`, loaded from the `units.reg` key `TileSize` at `L05673` and defaulting to the parent record's value or to 1 -- so every mark offset scales with the actor's footprint. **G2:** the record width, the field offsets, the two-pass split and the `TileSize` multiplier are engine layout and immediates; changing any of them changes `ROM.EXE`. `TileSize` itself and the sheet behind each record index are data and are free.

**Confidence.** High. `R0552` was read end to end, both passes and the anchor arithmetic are single instruction sequences, and the depth polarity is fixed by two opposite conditional branches on the same word. 159 byte assertions reproduce identically from both preserved `ROM.EXE` roots.

**Original status.** ✔ promoted

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

### MAGIC-MARK-060

**(rom.exe) The mark set is not stored: it is re-derived every rebuild from a list of kind-and-countdown pairs the simulation opens and closes with two dedicated messages.** The client unit carries a second array at `unit+0x124`, data `unit+0x128`, count `unit+0x12c`, whose elements are one dword each: high 16 bits a kind, low 16 bits a countdown. `R0612`, `MAGIC-ATTACH-016`'s attach routine, sets `1 << effect+0x0c` into `actor+0x144` (`L05674` through `L05675`) and then, when `effect+0x0c` is non-zero (`L05676`, a zero test), calls `R0627`, which writes opcode `0x88` (`L05677`) and carries the target runtime id at `msg+0x0a` and the `Effect` picture field `effect+0x0e` at `msg+0x0c` (`L05678`). `R0673`, the expiry routine, clears the same bit (`L05679`, a bit inversion of the mask) and calls `R0628`, which passes opcode `0x89` (`L05680`). Attach also evicts a same-spell effect first, so a recast sends `0x89` then `0x88`. On the client, `R0509` case `0x88` builds `(kind << 16) or 0xffff` (`L05681`, `L05682`) and either overwrites the element of that kind, located by `R0615`, or appends one; case `0x89` removes it (`L05683`). There is a **third** writer: the projectile driver `R0558` dispatches the arriving projectile's picture (`projectile+0x20`) over the closed range 13..64 through the tables at `L02879` and `L03000`, and two of its arms write the same list. The arm at `L05684`, selected only by pictures **20 `healing`** and **30 `Drain`**, replaces or appends an element with countdown **`0x20` = 32** (`L05685`, `L05686`); the arm at `L05687`, selected by pictures 18, 24, 28, 40, 44, 48, 52, 54, 56, 62 and 64, appends one with countdown `0xffff` and no by-kind check (`L05688`). Under `MAGIC-CASTSPAWN-033` the cast spawner gives flight length 0 to all eleven of those and 1 to the two, so on the cast path only the countdown-32 arm is reached. `R0610` decrements every element's low half once per rebuild and removes it at zero (`L05689`, `L05690`), then discards and rewrites the whole mark array. There are therefore **two lifetimes**: an element opened by `0x88` at `0xffff` is bounded by the `0x89` message and its countdown is only the phase clock, while an element the projectile driver opens at 32 expires by itself after 32 rebuilds. **G2:** the two opcodes, the `0xffff` initial value, the 16-bit split of the element and the one-dword element are engine constants; a longer countdown or a wider kind changes `ROM.EXE` and the wire format together.

**Confidence.** High. Both simulation routines, both message senders and both client arms were read end to end; the pairing is fixed by the two opcode immediates and by the same `effect+0x0c` gate on each side. Medium for the reachability clause about the countdown-`0xffff` driver arm, which rests on `MAGIC-CASTSPAWN-033`'s flight-length table rather than on a fresh enumeration of every projectile creator. The three writers were found by `EnumRefs disp:128` and `callto:R0615` together, and either sweep alone misses one.

**Original status.** ✔ promoted

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

### MAGIC-MARK-061

**(rom.exe) Ten spells put a mark on an actor and eighteen put none; the split is a 45-byte table in `.text`, not a data column.** `R0610` computes `kind = element >> 16`, rejects anything outside the closed range `0x12..0x3e` (`L05691` biases the kind by −0x12, `L05692` bounds it at 0x2c), and dispatches through the index table at `L05693` and the 11-entry jump table at `L05694`. A kind is `2*spellId + 8`, the even half of `MAGIC-PIC-026`'s pair. Ten kinds reach seven builders: `0x12` `0x1c` `0x28` `0x34`, the four Protections, share `L05695`; `0x14` `heal` takes `L05696`; `0x18` `poison_cloud` takes `L05697`; `0x1e` `drain_life` takes `L05698`; `0x2c` `shield` takes `L05699`; `0x36` `bless` takes `L05700`; `0x3e` `curse` takes `L05701`. The other 35 kinds in range reach `L05702`, which appends nothing and only runs the countdown. Consequently `freezing_cloud` (0x16), `invisibility` (0x26), `stone_curse` (0x30), `haste` (0x38), `light` (0x20), `darkness` (0x2a), `lightning` (0x22), `prismatic_spray` (0x24), `acid_stream` (0x1a), `wall_of_earth` (0x2e), `meteor_storm` (0x32), `control_spirit` (0x3a) and `teleport` (0x3c) draw no mark, and `slow` (0x40), `fire_arrow` (0x0a), `fire_ball` (0x0c), `wall_of_fire` (0x0e) and `fire_sacrifice` (0x10) fall outside the dispatch range entirely. Each spell's arm is independent of every other, so an actor carrying several effects shows the union of their record sets. **G2:** which spells can be marked is an engine table -- adding one means editing `ROM.EXE`, not `Data.bin` or a registry.

**Confidence.** High. The dispatch bound, both tables and every arm entry are extracted from the shipped image by `tools/effectmark` and cross-checked against the image's own 29-entry spell name array at `L05034`, which the probe reads rather than transcribing; identical on both roots.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Original status.** ✔ promoted

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

### MAGIC-PROT-062

**(rom.exe) The four Protections share one builder and differ only by a placement table; together they form a diamond of half-width 6 pixels.** `R1096` receives the vertical base `TileSize*32`, the kind, and `countdown mod 6` as the phase (`L05703` through `L05704`). It dispatches the kind over `0x12..0x34` through the index table at `L05705` and the 5-entry jump table at `L05706`, setting one of four terms, then writes exactly one 8-byte record: `dx = -h`, `dy = -(v + base)`, `depth = 0`, `record index = kind`, `phase = the caller's phase`. Four of the 35 index bytes select a non-centre arm, and the assignment is `protection_from_fire` (kind `0x12`) `dx = -6`; `protection_from_water` (`0x1c`) `dx = +6`; `protection_from_air` (`0x28`) `dy = -6`; `protection_from_earth` (`0x34`) `dy = +6`. All four therefore sit `TileSize*32` pixels above the actor anchor, 6 pixels left, right, up and down of one another, and all four carry `depth = 0`, so all four draw after the actor sprite. Nothing in the arm consults the other effects, so the diamond is what four simultaneous Protections produce rather than a case the code recognises. **G2:** the 6-pixel radius, the four-way assignment and the `*32` base are `.text` immediates and tables; a different arrangement means editing `ROM.EXE`. The art at each record index and the `TileSize` scale are data.

**Confidence.** High. The builder was read end to end, the single append is one instruction pair at `L05707`, and both placement tables are extracted from the shipped image; the competing shared-anchor reading is excluded by the four distinct jump targets the table selects.

**Original status.** ✔ promoted

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

### MAGIC-SHIELD-063

**(rom.exe) `shield` is the one effect that draws a figure enclosing the actor rather than a mark beside it.** Its arm calls `R1097` with `TileSize*11`, `float TileSize*28`, `float TileSize*16` and `countdown mod 90` (`L05708` through `L05709`). That builder calls `R1098`, which calls the midpoint circle rasteriser `R1099` -- its initial term is `2 - 2r` at `L05710` and `L05711` -- and then rotates each generated point through `FSIN` and `FCOS` and appends **two** records per point, both with record index `0x2c` (`L05712`, the constant 0x2c, stored at two record offsets 8 bytes apart). The vertical envelope is `ftol(abs(phase/45 - 1) * TileSize*28)`, using `1/45` at `L05713` and `1.0` at `L05714` with an absolute value at `L05715`, and the two records of a point are placed at `-(envelope + a)` and `envelope - a`. The envelope contracts to zero at phase 45 and expands again, on a 90-step cycle. The per-record frame index is `4` minus a quotient by 9 (`L05716` loads 0x4, `L05717` subtracts the phase term from it). Because a signed vertical term is written into the record's `depth` field, `MAGIC-MARK-059`'s two passes split the figure: the half above the actor draws behind its sprite and the half below draws in front. **G2:** the rasteriser, the 90-step cycle, the `*11`, `*28` and `*16` multipliers and the two-records-per-point rule are engine code; the sheet at record index `0x2c` is data. **↪ SUPERSEDED IN PRECISION by `MAGIC-091` and `MAGIC-092` (EXP-0189).** The original instructions and 90-step envelope observation stand, but the wording combines two separately rasterised components and leaves their coordinate assignment open. Component A is a fixed-radius rotated point set; Component B is the changing pair of circles.

**Confidence.** Medium. Every cited immediate, constant and call is asserted byte for byte, and the two-records-per-point and rasteriser facts are read from the instructions. The assignment of each computed term to each record was read from the store order rather than proved, and the rotation inside `R1098` was not fully decoded, so the point count and the closed form of the horizontal extent are not established.

**Original status.** ● active (superseded by EXP-0189)

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/), [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

**Amended.** The table ledger carried the status "● active (superseded by EXP-0189)". `retracted.md` records a correction against this claim.

### MAGIC-BLESS-064

**(rom.exe) `bless` and `curse` each place twenty marks on a rotating circle, and the two rotate in opposite senses.** Both arms pass `TileSize*32`, the float `20.0` and `countdown mod 5`; the radius immediate is `0x41a00000` at `L05718` and `L05719`. `R1100` and `R1101` each run a five-step loop appending four records per step, so 20 records, with record index `0x36` and `0x3e` written at four offsets 8 bytes apart at the top of each routine. The angle is in degrees, converted through `pi` at `L05720` and `1/180` at `L05721`. `bless` starts its step counter at 89 and steps down by 18 while it is positive (`L05722`, `L05723`, `L05724`), giving 89, 71, 53, 35, 17; `curse` starts at 0 and steps up by 18 while it is below 90 (`L05725`, `L05726`, `L05727`), giving 0, 18, 36, 54, 72. The step direction is opposite, so the two trails sweep opposite ways as the countdown falls. The per-record frame index is `4 - step/18`, which runs 0,1,2,3,4 for `bless` and 4,3,2,1,0 for `curse` -- five different frames visible at once, one per trail step. The four records of a step take `+cos`, `-cos`, `+sin` and `-sin` terms, and a sine term goes into the record's `depth` field rather than into `dy`, so `MAGIC-MARK-059`'s two passes carry the upper half of the ring behind the actor and the lower half in front. **G2:** the 20-mark count, the five steps, the 18-degree increment, the 20.0 radius and the 89-degree offset are `.text` immediates and constants.

**Confidence.** Medium. The loop bounds, step increments, angle constants, record indices and frame arithmetic are each asserted byte for byte and reproduce on both roots. The assignment of the four sign combinations to the four records per step was read from the store order rather than proved, and the stack frame shifts under the array-growth calls between those stores.

**Original status.** ✔ promoted

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

### MAGIC-CLOUD-065

**(rom.exe) `poison_cloud` puts one mark above the actor; `heal` and `drain_life` append no mark and instead carry the previous rebuild's records forward.** `R1102` receives `TileSize*32` and `countdown mod 6` and appends exactly one record: `dx = 0` (`L05728`), `dy = -base` (`L05729`, a negation), `depth = 0` (`L05730`), record index `0x18` (`L05731`), phase from the caller. `R1103` and `R1104` instead receive a pointer to `unit+0x110` together with `TileSize*16`, `TileSize*8`, `TileSize*32` and the countdown, and walk the array `MAGIC-MARK-059` left at `unit+0x114` on the previous rebuild -- which is still intact, because `R0610` writes the new array only after every arm has run. For each record whose index equals their own, `0x14` and `0x1e` (`L05732`, `L05733`), and whose frame at `+0x07` is below 7 (`L05734`, `L05735`), they increment that frame and subtract `ftol(1 - c*(-1/7))` from the record's `dy`, using `1.0` at `L05714` and `-1/7` at `L05736`, where `c` is the element's countdown read **sign-extended**. A particle therefore lives seven rebuilds and moves on each one. This is the only arm family whose output depends on the previous frame, so `MAGIC-MARK-060`'s re-derivation holds for eight of the ten kinds and not for these two. Their countdown is also the only one that is a real lifetime: `MAGIC-MARK-060`'s projectile-driver arm opens a `healing` or `Drain` element at 32 rather than at `0xffff`, so `c` runs 32 down to 0 and the subtracted rise falls from 5 pixels to 1 over the element's life. **G2:** the seven-frame particle life, the `-1/7` step and the three `TileSize` multipliers are engine constants. **✖ HEAL/DRAIN CLAUSES REFUTED by `MAGIC-089` and `MAGIC-090` (EXP-0189).** The `poison_cloud` clause stands. Heal and Drain do append fresh 3–5-particle cohorts; the motion term reads constant `32T`, not countdown; a cohort is visible at phases 0–7; and Drain adds the step while Heal subtracts it.

**Confidence.** Medium. `R1102` was read end to end and is High on its own; the carry-forward branch of the other two was read and asserted, but their spawn branch -- what they append when no carried record of their index exists -- was not read, so the number of particles alive at once is not established.

**Original status.** ✖ partially retracted (EXP-0189)

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/), [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

**Amended.** The table ledger carried the status "✖ partially retracted (EXP-0189)". `retracted.md` records a correction against this claim.

### MAGIC-ACTOR-066

**(rom.exe) `stone_curse` and `invisibility` change the actor's own sprite instead of adding a mark, and the unit draw tests for them by kind directly.** `R0552` calls `R0615`, the by-kind lookup on the effect list, four times on its own account outside the mark passes. Two calls push kind `0x30`, which is `2*20 + 8`, `stone_curse` (`L05737`, `L05738`), and gate an arm that holds the actor's animation frame. Two push kind `0x26`, which is `2*15 + 8`, `invisibility` (`L05739`, `L05740`), and gate the whole actor sprite on a per-player bit, `L05741` (a test of bit 0x8 in the byte of the tile word at the cell), skipping the sprite when it is clear. `R0379` performs the same `0x26` test at `R1105` before drawing a unit at all. Neither kind reaches a mark arm under `MAGIC-MARK-061`. Two models of what a lasting effect looks like are therefore both present in the image, over disjoint spell sets: eight spells add records to the mark array and these two modify the actor's own draw. This is the client-side half of the long-open question about what makes an invisible unit unseeable. **G2:** both are `.text` arms keyed on hard-coded kinds; neither is reachable from data.

**Confidence.** High for the two kinds, their three call sites and the two arms they gate, each a single instruction sequence asserted byte for byte. Medium for the description of the `stone_curse` arm as an animation hold: the gate and the frame arithmetic were read but the arm's full effect on sprite selection was not traced. The per-player bit's owning table was not identified.

**Original status.** ✔ promoted (amended)

**Evidence.** [EXP-0177](../experiments/EXP-0177-effect-marks/)

**Amended.** The table ledger carried the status "✔ promoted (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-ACTGATE-079

**(rom.exe) An actor carrying spell 20 runs no arm of the order switch, and the refusal is one byte above the choice of action; it has one bounded escape.** `R0016`, the per-actor order machine, reads `actor+0x144` at `L05213` — eleven instructions before it reads the order kind at `L05742` — and when the test of 0x100000 at `L05743` passes and the order object's progress byte `ord+0x09` is 0 (`L05744`…`L05745`) it writes, at `L00546`, the byte 4 loaded at `L05746` to `ord+0x09`. `L00547`, a jump on zero to `L00548` (rel32 `0xd8`), is the only entry to the fifteen-arm order switch and it is taken only while that byte is 0, so kinds 1 walk, 2 attack, 8 cast-at-actor and 9 cast-at-cell (`AI-ORDER-039`) are refused by one test rather than by four. Value 4 routes through the 255 index bytes at `L00090` into the six arms at `L00091`, both read out of the PE by virtual address, and lands on `L00543`: `L02932` (storing 0x1a at the 32-bit field `+0x54`), `L05747` (a test against 0x100000), then either `L05748`, a jump on non-zero to `L00010` (rel32 `0x41e`), or `L05749` (storing the byte 0 at `ord+0x9`). The arm calls nothing. **The refusal has one bounded escape and it is not in the switch:** both exits of the arm reach the machine's common tail at `L00010`, which is not gated on the mask. The tail runs when `mover+0x98` is non-zero, clears that flag (`L05750`), and for any command state but 1, `0xa` and `0x17` calls `L00011` (a call of R0004), target acquisition. When that finds a target the reach test `L00741` (a call of R0041) accepts, it calls `R0005` — one caller image-wide — whose whole body is `ord+0x09 = 1`, `ord+0x15 = 0`, `actor+0x54 = 3`, `actor+0x5c = ord+0x0c`: order kind 2's own attack install, **overwriting the parked 4**. The gate does not re-park while the byte reads 1, so progress arm 1 at `L00092` runs, holding `actor+0x54 = 3` and incrementing `ord+0x15` until it exceeds 2 with `actor+0x136` set, then returns the byte to 0 and the gate parks the actor again. **The escape cannot repeat within one refusal:** `EnumRefs disp:98` (215 hits / 140 owners / 4 orphan) and `disp:99` (0 hits) give `mover+0x98` exactly three setters — `L01720` in `R0178`, `L00161` in `R0043`, `L00162` in `R0055`, whose own `callto:` is those same two — and all three are inside movement executors that `MOVE-GATE-039` shows are reachable only from an order arm. Progress arm 4 calls none of them, and the tail clears the flag itself, so the tail can admit at most one re-decision per refusal, on the first parked tick. Instrument for the writer set: `EnumRefs "text:+ 0x9],"`, whole image — **157 hits / 111 distinct owners / 0 in orphan or undisassembled code**, of which 18 have a base loaded from `actor+0x158`; their value set is `{0,1,2,3,4,0xff}`, **value 4 is written at `L00546` and nowhere else**, and value `0xff` at `L00545` and nowhere else, on the tail's `actor+0x50 == 0x17` arm, cleared only by `R0146` (`L05751`, `L05752`). Five further `0xff` stores at the same displacement are on other objects — `L05753` and `L05754` on the mover at `actor+0x154`, `L05755` on the position object at `actor+0x10`. So an actor stops acting for three unrelated reasons, not one: this bit, `ord+0x09 = 0xff`, and `actor+0x54 == 0x10`, which makes `R0037` return at `L05756` before the machine runs at all

**Confidence.** **High** (every cited instruction re-read from the raw image bytes at its own address, 65/65, and every cited branch target recomputed from the branch's own displacement, 10/10, by `tools/effectgate -mode asserts`; both switch tables are read out of the PE with no disassembler in the path; the rival reading "the gate is movement-specific" is excluded by the fixed order of two reads inside one routine, not by plausibility) / **Medium** for the completeness of the `ord+0x09` writer set — the sweep is complete for byte stores through a register base, and the wide-store check at displacements 6..9 (`EnumRefs re:`, 2765 hits, filtered to non-stack bases inside the order module, 20 candidates, all linked-list nodes from the allocator at `R0052`) came back empty, but neither can see a store through a pointer offset by address arithmetic, and both `ord+0x08` and `ord+0x09` sit inside the order object's serialized `0x94`-byte raw block, which is `AI-PROGRESS-034`'s own blind spot

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MAGIC-ACTKEY-080

**(rom.exe) The refusal is keyed to the spell id, and to nothing about the effect.** `actor+0x144` bit *n* is set by one instruction sequence in the image: `L05674`..`L05675` (read the byte `effect+0xc`, shift 1 left by it and merge it into the 32-bit mask at `+0x144`), in the attach `R0612`. `effect+0x0c` is stamped from `spell+0x8` by the arm that builds the effect — for spell 20, `L05757`/`L05758` (a copy of the byte `spell+0x8` to `effect+0xc`). The gate's immediate `0x00100000` is `1 << 20`, and entry 20 of the image's own 29-entry spell-name array at `L05034` is `stone_curse` (`tools/effectgate -mode spells`, `evidence/spell-bits.csv`). `R0016` reads none of `effect+0x3c` (kind), `+0x3d` (mode), `+0x40` (magnitude) or `+0x42` (duration): the effect record is not on the path at all, only the id bit it left on the actor. **Consequence for a consumer: implementing spell 20 from its `Effects` column produces `absorbtion=+5` and no immobilisation**, because the column supplies the kind (`MAGIC-EFFECT-015`) and the immobilisation comes from the id. Confirming `MAGIC-EFFECT-015`'s "a duration and no magnitude" for this arm: `L05759`…`L05760` stores `+0x3d` once (`L05761`, mode `duration` from `[L04012]`), `+0x42` four times, `+0x0c` and `+0x0e`, and contains **no store to `+0x40` and no occurrence of the displacement `0x40` at all**. The refusal ends by the ordinary countdown: `R0673` decrements `+0x42` (`L05762`, a decrement by 1) and at zero clears the same bit through `L05763`..`L05764` (the one-bit value `1 << effect+0xc` inverted and applied to the mask at `+0x144`). Before resistance the duration is `ftol(1.025^power × 10 × 16)` ticks — 160 at power 0, 262 at 20, 549 at 50, 1890 at 100 — then multiplied by `(100 − target+0xca)/100` with a floor of one tick (`MAGIC-SING-019` d)

**Confidence.** **High** for the keying (the immediate, the single bit-setting writer and the id stamp are named instructions, byte-asserted at their own addresses; the rival "keyed to an effect kind, mode, magnitude or spell column" is excluded by the *absence* of any such read on a routine read end to end) / **Medium** for the tick figures, which are EXP-0067's re-execution of the arm's arithmetic on the shipped `Data.bin` row rather than a measurement of a running game and move with that row

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MAGIC-MASKREAD-081

**(rom.exe) The reader set of `actor+0x144` is nine routines, and `MAGIC-ATTACH-016`'s enumeration names six of them.** Instrument: `EnumRefs disp:144 disp:145 disp:146 disp:147`, whole image, on the repaired function table — **125 / 0 / 7 / 0 hits**, the `disp:144` set being 63 distinct owners with 2 in orphan code. The field is a dword, so `0x145`…`0x147` are the displacements a narrower access could reach it through; the seven `disp:146` hits are word accesses in the code-region, code-region and code-region classes and none of them is the actor. Of the 125, nineteen have an actor base. Writers, all in the effect machinery: `L05765` zeroes it at spawn, `L05675` sets a bit, `L05766` and `L05767`/`L05768` clear one, `L05769` un-applies, `L05770` serializes. Readers: `L05771` the generic `1 << id` tester `R0254`; `L05213` and `L00543` the order machine, whose `this` is the actor because its only live caller passes its own `this` (`L05772`/`L05773`/`L00086`) and is reached through offset `0x18` of the three actor vtables, independently confirmed by a raw dword scan that found `R0037` stored at `L03030`, `L03031` and `L03032` and nowhere else; `L00788` and `L05774` the two AI sites `MAGIC-ATTACH-016` names; `L05214` `R0394`, testing `1 << spell+0x8` built at `L05775`/`L05776`; and `L05777` `R0293`, testing bit 15 as the byte test of 0x80 at `L05215`. **The three the ledger's enumeration omits are `R0016`, `R0394` and `R0293`, and the first of them is the routine that answers what the mask does.** The last is also the shape `INSTRUMENT.md` rule 8 warns about from the other side: a bit test folded into a sub-register makes the mask invisible to any `imm:` sweep — `EnumRefs imm:100000` returns 31 hits over 20 owners and finds the two literal bit-20 tests, while no immediate sweep for `0x8000` would ever have found a byte-wide test of 0x80 in the high byte

**Confidence.** **High** (the enumeration states its instrument and its blind spot, every hit at the field's own width and at the three narrower displacements was read, no hit was dropped for want of a function name, and the base-register identification rests on a caller read at instruction level plus a raw pointer scan rather than on a name) / **Medium** for the writer half: a displacement sweep cannot see a wholesale structure copy, and `R0210` restores the field from a save through `CArchive`

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MAGIC-AIBIT-082

**(rom.exe) Two AI decisions read bit 20, neither of them gates an action, and the second is one case of a general no-restack rule.** `R0225` returns a melee target cost — `L00781` (the constant 0xffffff) is its unreachable-target sentinel, so lower is preferred — and adds a flat 127 for a target that carries the bit: `L00788`/`L05778` (a test of the mask `+0x144` against 0x100000, jumping to `L05779` on zero, rel8 `0x0d`) and `L05780` (adding 0x7f), floored at 1 by `L05781`/`L00789`. `R1048` picks a spell for an AI caster by asking the book for ids `0xe`, `0x14`, `0xd` and `1` in that order (`R0017`), each gated on the caster's mana at `+0x9a`, and gives id `0x14` one extra condition the other three do not have: `L05774`/`L05782` (a test of the mask `+0x144` against 0x100000, jumping to `L05783` on non-zero, rel8 `0x05`), so the target must not already carry it. That pairing of `L05784` (the argument 0x14) with a test of `1 << 20` thirty-two bytes later fixes the bit index as the spell id independently of the attach's own shift. The same refusal generalised is `R0394`: `L05775`..`L05785` (the spell id byte `+0x8` turned into a one-bit value, tested against the mask `+0x144` with a jump on non-zero), for whatever spell the loop is holding. So the mask serves the AI as a "has this target already got this spell" index, and the action gate is the one consumer that reads a fixed bit out of it

**Confidence.** **High** (each test, its branch displacement and its arithmetic are named instructions, byte-asserted at their own addresses; the 0x14-argument/0x100000-test pairing is a second, independent fixing of the id-to-bit relation) / **Medium** that lower is better in `R0225`: the sentinel and the `+0x7f` are read, the routine's consumer was not

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MAGIC-EFFLOOK-083

**(rom.exe) The effect list at the drawable has one reader, `R0615`, and its complete call-site set is fourteen addresses over six routines, all in the presentation layer.** `R0615(this, kind)` walks the container at `drawable+0x124` -- data `+0x128` (`L05786`), count `+0x12c` read as a **dword** (`R0615`) -- one dword per element, and returns the index of the first element whose `element >> 16` equals the argument (`L05787`..`L05788`: the high 16 bits of a 32-bit entry compared with the argument) or -1 (`L04364`, the constant 0xffffffff). The sites are `L05789` (`R0379`, map draw); `L05790` and `L05791` (`R0509`, the `0x88` and `0x89` message arms); `L05792`, `L05793`, `L05794`, `L05795` (`R0552`); `L05796`, `L05797`, `L04337`, `L05798`, `L05799` (`R0553`); `L05800` (`R1106`); `L05801` (`R0558`, the projectile driver). Eleven pass a literal kind, and the only literals in the image are `0x30` (`stone_curse`) and `0x26` (`invisibility`); the other three take the kind from the wire or from `projectile+0x20`. `R0552`, `R0553` and `R1106` are vtable slots 4, 5 and 6 of the client unit class, read out of `.rdata` at `L05802`/`414`/`418` and `L05803`/`49c`/`4a0`, beside `R0586` at slot 2 (`MAGIC-MARK-059`). `MAGIC-ACTOR-066` names two of these six routines; `R0553` and `R1106` are the two it does not.

**Confidence.** **High** (the call-site set was located twice with instruments that have different blind spots -- `EnumRefs callto:R0615` on the repaired function table, and a raw scan of every executable section for the `E8 rel32` encoding whose recomputed target is `R0615` -- and the two agree address for address: 14 hits, 6 owners, 0 orphan. `EnumRefs callto:` on `R0553` and `R1106` returns only `RDATA-SLOT` rows, which `INSTRUMENT.md` rule 7 says is not an answer, so the slots were read out of the PE by virtual address instead. Blind spot: a call through a computed pointer is invisible to both instruments, so a seventh caller of that form cannot be excluded)

**Original status.** ● active

**Evidence.** [EXP-0182](../experiments/EXP-0182-effect-at-actor/)

### MAGIC-STONEDRAW-084

**(rom.exe) The `stone_curse` arm of the unit draw has two parts: it replaces the frame index with the facing, and it replaces the shade table with a greyscale one.** Both are gates on `R0615(this, 0x30)` in `R0552`, and both carry the same second condition, a compare of the byte at `+0x15a` with 0x2 and a jump on above. **Frame** (`L05737`): after the per-state arm has computed the frame index, `L05804`..`L05805` (read `+0x6c`, subtract 0x8, mask with 0xf) overwrite it with the facing at full 16-way resolution, and above 8 mirrors it (`L05806`..`L05807`: 0x10 minus the facing) with the mirror flag at stack offset +0x10 set to 1 (`L05808`); `drawable+0x6c` is the same field the routine's head reduces to 8 directions at `L05809`..`L05810`. Both values reach the blit at `L05811` and `L05812`. The identical arm is present a second time in `R0553` at `L05813`..`L05814`. **Shade** (`L05815`): the object at `stack+0x14`, selected earlier on `record+0x98` (`PAL-KEY-002`'s `Palette`), is overwritten. When `Palette == 0` it becomes `[[L05816] + 0x40]`, element 16 of the array whose elements 0..15 are the sixteen team tables (`PAL-OWN-007`, `TERR-LIGHT-064`); otherwise `R1107([record+0xac], 0x10, 5, 0)` builds one on the stack (`L05817`..`L05818`), against `(0x10, 2, 1)` on the ordinary paths at `L05819` and `L05820`. Mode 5 is the arm at `L05821`, read from the six-entry jump table at `L05822`, and it writes one grey per level from `(r+g+b) * level * 2 / 3 / levels`. This closes `TERR-LIGHT-064`'s Unknown on what the two mode-5 overrides are for. `MAGIC-ACTOR-066`'s single description of the arm as an animation hold covers the frame gate and not the shade gate.

**Confidence.** **High** for the two gates, their conditions, the frame arithmetic and the shade arguments: every cited instruction was re-read from the raw image bytes at its own address and every cited branch target recomputed from its own displacement, 93 assertions and 18 anchors with 0 mismatches, identically from both preserved roots; the mode table was read out of the PE by virtual address, which refuted the assumption that arm order follows address order. **Medium** for the statement that both shade forms render grey: mode 5's arm is read at instruction level, but element 16's own build (`R0395`) was not read here and rests on `PAL-OWN-007`. `drawable+0x15a`, which gates both parts, is **Unknown**

**Original status.** ● active

**Evidence.** [EXP-0182](../experiments/EXP-0182-effect-at-actor/)

### MAGIC-INVISOWN-085

**(rom.exe) The per-player bit that gates the invisibility sprite is an ownership test: an `invisibility` element hides the unit from every client except its owner's.** At `L05739`..`L05823` the drawable's owner index is `*(*(drawable+0x14)+0x4)` (`PAL-SHADE-012` identifies `drawable+0x14` as a `CPlayer`), the row is `[[session + 0x9b4] + 0x38]`, and `L05741` (a test of bit 0x8 in the byte of the tile word at the cell) decides between the blit at `L05824` through `vt+0x34` and a jump to the routine's epilogue at `L05825`, which destroys the stack shade table and returns without drawing. A unit with no `invisibility` element blits through `vt+0x14` at `L05826`. The row is the 32-entry `CWordArray` at `CPlayer+0x34` that `UNIT-VPLAYER-021` publishes, and `UNIT-VISBIT-044` establishes that bit 3 is set on exactly one entry, the local player's own. The same test appears at `L05827` and at the four further invisibility gates in `R0553` and `R1106` (`L05828`, `L05829`, `L05830`), and outside the spell at `L05831` in the map draw (`MAGIC-ACTOR-066`). `MAGIC-ACTOR-066` records the owning table as unidentified; it was already published in `claims/unit.md` on 2026-08-04, so this is a cross-ledger join and not a new identification.

**Confidence.** **High** for the instruction chain and for the reduction to an ownership test, which follows from `UNIT-VISBIT-044`'s write population rather than from corpus agreement. **Medium** for what the owner sees: the two blit slots are different call sites of the object at `drawable+0x194` and neither implementation was read, so that an invisible unit looks different to its owner rather than identical is inferred from the slot difference alone. The four gates in `R0553` and `R1106` were classified by their kind and by the bit test that follows, not read end to end

**Original status.** ● active

**Evidence.** [EXP-0182](../experiments/EXP-0182-effect-at-actor/)

### MAGIC-089

**(rom.exe) Heal and Drain Life carry an eight-output particle cohort with a constant, tile-scaled vertical step.** Each builder walks the previous mark array first, keeps only its own kind (20 or 30) whose phase is below 7, increments the phase, and copies it before any fresh record. Heal then subtracts `s` from `dy`; Drain adds it, where `s = trunc_toward_zero(1 + 32T/7)` and shipped `T={1,2,3}` gives `s={5,10,14}`. The fourth explicit builder argument, `32T`, supplies this term, not the changing countdown; `R0279` selects x87 truncation before `FISTP`. Thus old phase 6 becomes visible phase 7 and old phase 7 is dropped on the next rebuild. **G2:** the kind ids, phase bound, sign, `32T/7` law and carry-before-spawn order are engine code; `TileSize` and the sheets are data.

**Confidence.** High. The stack positions, kind/phase tests, stores and truncation helper are among `tools/spellvisual`'s 94 PE assertions, 94/94 on the install and both preserved roots, and the independent falsifier reconstructed the argument order and terminal phase.

**Original status.** ✔ promoted

**Evidence.** [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

### MAGIC-090

**(rom.exe) Heal and Drain Life append a 3–5-particle elliptical cohort on each rebuild whose unsigned countdown is at least 8.** One shared CRT call gives `N=rand()%3+3`; each particle consumes one further call and uses the same `a=rand()%360` degrees for `dx=trunc(cos(a)*16T)` and `depth=trunc(sin(a)*8T)`. Heal starts `dy=0`, kind 20; Drain starts `dy=-32T`, kind 30; both start phase 0. A countdown opened at 32 therefore spawns 25 cohorts, 75–125 particles total. Each cohort is visible at phases 0–7, eight cohorts overlap in steady state, and 24–40 particles can be live there. The process-wide RNG is `state=state*214013+2531011 mod 2^32`, returning `(state>>16)&0x7fff`, seeded once from `time(0)`; there is no private effect seed. Heal and Drain share cadence, ellipse and phase life, but differ in initial `dy`, carried motion sign and cast-projectile draw arm. **G2:** spawn gate, count, RNG order, angle modulus and multipliers are engine code.

**Confidence.** High. Both builders, the CRT recurrence/seed and the modulo operations are byte-asserted on all three roots; cohort totals are arithmetic over closed bounds, and the independent review reproduced them.

**Original status.** ✔ promoted

**Evidence.** [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

### MAGIC-091

**(rom.exe) Shield Component A is a sampled midpoint circle transformed into two ordered kind-44 records per source point.** The integer rasteriser starts `x=0,y=r,err=2-2r`, samples the first iteration and thereafter whenever its counter reaches the interval, and emits axis duplicates in order `(+x,+y),(-x,+y),(+x,-y),(-x,-y)` while `y>0`; radius zero emits nothing. A uses `r=16T`, interval 4, `p=countdown%90`, `theta=4p` degrees. For each sampled `(x,y)`, `q=trunc(y*28T/16T)` and `f=abs(4-trunc(abs(y)*5/16T))`, then A1 is `(dx=trunc(x*cos(theta)),dy=-q,depth=11T+trunc(x*sin(theta)),phase=f)` and A2 is `(dx=-trunc(x*sin(theta)),dy=-q,depth=11T+trunc(x*cos(theta)),phase=f)`. Both records use kind 44 and A1 precedes A2. **G2:** the rasteriser, radius, interval, transforms and record order are engine code.

**Confidence.** High for the instructions and formulas: every branch, constant and truncating store is asserted on all three roots and independently reconstructed. The committed Go phase hashes are a transcription, not proof that a host transcendental library rounds identically to x87 at every truncation boundary.

**Original status.** ✔ promoted

**Evidence.** [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

### MAGIC-092

**(rom.exe) Shield Component B is a second, independently rasterised pair of circles appended after all Component A records.** For `p=countdown%90`, stored qword `c45` has bits `0x3f96c16c16c16c17` and exact value `6405119470038039/2^58`. CRT startup requests precision `0x10000` under mask `0x30000`, whose translator sets x87 PC bits `0x0200`. A debugger observation after the initializer returns and before `WinMain` reads full control word `0x027f`, PC53 and round-to-nearest, on both lawful roots. A raw scan classifies all 127 FLDCW, FLDENV, FRSTOR, FSAVE/FNSAVE, FNINIT and FXRSTOR candidates. The two valid FSAVE resets each execute one call and then immediately restore the saved environment with FRSTOR before any branch or return; no lasting PC64 writer remains on the normal Shield path. The envelope therefore rounds every arithmetic instruction to 53-bit significand precision: `x1=round53(p*c45)`, `x2=round53(x1-1)`, `x3=abs(x2)`, `x4=round53(x3*float32(28T))`, `E=trunc(x4)`. `__ftol` changes only the rounding-control bits to truncation and restores the saved word. It then computes `F=4-floor(abs(p-45)/9)` and `R=trunc(sin(acos(E/(28T)))*16T)`. Exact-dyadic retention and division by 45 are not interchangeable: the shipped discriminants are `T=3,p=15:E=56`, `p=30:E=27`, and `p=60:E=28`. It rasterises radius `R` at interval `2F`, and for each source `(u,v)` appends B1 `(dx=u,dy=-(E+11T),depth=v,phase=F)` then B2 `(dx=u,dy=E-11T,depth=v,phase=F)`, both kind 44. The raw stack discriminator is explicit: after the interval push, `stack+0x34` aliases the pre-push `16T` argument at `[P+0x30]`, while `E`, stored at pre-push `[P+0x34]`, has moved to `stack+0x38`. At `p=0`, `F=-1` and `R=0`, so no invalid frame is drawn. For every shipped `T={1,2,3}`, this is the only zero-radius phase; at `p=45`, `E=0` and `R=16T`. Every emitted B record uses a frame in 0..4. The builder reads neither RNG nor previous marks. Rebuild precedes countdown decrement, so the low `0xffff` countdown descends modulo 90; recast detach/attach resets the cycle. The older one-circle-plus-envelope account is superseded. **G2:** every formula, append order and the 90-step cycle are engine code; sheet 44 is data. **The former `*E` radius clause, its derived zero-phase sets, exact-division envelope and exact-dyadic envelope are refuted below.**

**Confidence.** High for the instruction formula, stored constant, startup runtime control word on both roots, the 127-candidate raw and recursive all-opcode census, raw stack aliases, stores, state inputs and zero-radius edge. The PC53 model scans all 270 shipped rows: exact-dyadic retention disagrees at `T=3,p=15`, division at `p=30` and `p=60`, round-down at `p=30`, `p=60` and `p=75`, and round-toward-zero at `p=15`, `p=30`, `p=60` and `p=75`. Round-up agrees on the integer outputs, so the runtime control-word observation supplies that discriminator. A direct Shield-entry control-word witness was not obtained. Pixel identity at a transcendental truncation boundary still requires original x87 evaluation rather than the evidence table's Go transcription.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Original status.** ● active (corrected after fresh review)

**Evidence.** [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

**Amended.** The table ledger carried the status "● active (corrected after fresh review)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-093

**(rom.exe) Meteor Storm owns placement on the simulation side: 32 stages at effect ticks `0,3,...,93`, two ordered bounded-RNG calls per stage, and one stationary projectile per accepted cell.** Each stage takes x then y from `floor(rand()*(5+1)/32768)-2`, so each offset is in `[-2,3]`. A candidate is accepted only for `8 <= x <= mapWidth-9` and `8 <= y <= mapHeight-9`; a rejected edge cell is clamped internally but sends, applies and draws nothing. An accepted stage applies to occupants and sends picture 51 at the cell with lifetime 16. Repeated cells are allowed and create independent projectiles. The client receives a point, not a region, and consumes no placement RNG. **G2:** stage count/cadence, RNG bound/order, interior gate, picture and lifetime are engine code; map extent and picture-51 art are data.

**Confidence.** High. Producer, bounded RNG and sender instructions are asserted at their addresses on three roots; the stage schedule and acceptance chain were independently reconstructed.

**Original status.** ✔ promoted

**Evidence.** [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/)

### MAGIC-ITEMTRAIN-116

**A weapon-borne spell receives no immediate cast award; training comes from later event producers, and cadence follows accepted targets and damage ticks rather than releases.** Item context makes `actor+0x68 != 0`, so the `vt+0x68` half-mana award is skipped. Positive direct damage calls `vt+0x64` once per application; Drain Life uses the positive capped transfer; Slow id 28, Stone Curse and Curse id 27 use `trunc(target.healthMax*0.03)` once per accepted target before attachment; Poison Cloud id 8 calls on every nonzero signed tick while its recorded caster has health `>= 0`; negative caster health clears the pointer. Before the common sink, the damage-award caller requires a victim owner whose `+0x28` is nonzero and refuses when owner `+0x5c` and multiplayer `server+0x0c` are both nonzero. The damage award then becomes `trunc(target.XPvalue*0.5*amount/target.healthMax + 1)`. The common sink first requires recipient `typeID` in `[0x21,0x3f]`, then applies Mind scaling, relation refusal, the per-award level cap and the skill-100 ceiling. Thus transferred low-type Humans train nothing even when every later event gate passes. On the fighter route an ordinary-id rider requires positive physical damage and post-physical health above zero; Fire Ball alone is admitted whenever that conjunction fails, including after a miss or zero damage. The item spell runs before physical experience. That later award separately requires positive physical damage, the saved pre-hit target gate and post-spell health above `-10`, so rider damage can suppress it. The sink has no scaled-zero refusal and preserves signed amounts below its positive cap, so a custom negative-power Poison can heal by subtraction yet reduce slot and aggregate XP without lowering the level. One release can pay zero, one or many awards, and a refresh or Bless/Curse annihilation does not undo a pseudo-damage award already sent. **Prismatic Spray is route-dependent:** a caster item runs the fan during admission and the wrapper suppresses a duplicate; a fighter rider reaches only the refusing wrapper. Harmless buffs other than the three pseudo-damage arms train nothing. A delayed kill is another award with its own victim-owner gates and attribution history (`MAGIC-ITEMKILL-117`, `MAGIC-ATTRGATE-118`). The authored Data.bin corpus is 48 weapon-borne cells per root: 46 Human staffs plus Catapult and Ballista Unit weapons. The live shop is an additional runtime producer: unflagged generated weapons select ids `{1,11,13,14,20}` and store a price-derived random power capped at 100. Plain Unit `vt+0x60/+0x64/+0x68` are stubs, so those siege actors can release a rider but never train. The authored corpus exercises 18 distinct unclamped powers `{1,5,10,15,25,30,34,35,40,50,60,63,65,70,82,90,98,99}`; it is not a bound on generated runtime stock. The selected-inventory Cast command is not a third shipped Weapon route: its UI requires descriptor bits `0x10|0x01`, while a Weapon effect supplies bit 4 but its class descriptor can never supply the MagicItems-only bit 0. A crafted item order could reach the latent server arm, which is a G2 engine/client boundary.

**Confidence.** **High.** The raw-PE probe asserts 201 route, caller-gate, recipient-gate, item-Spell mutation, attribution, persistence, shop-producer and selected-item-control facts on the byte-identical executable; independent reference scans find two item-wrapper callers, three `Spell::Apply` callers, and a raw 42-candidate slot-`0x64` census selects seven actor award sites with none unclassified. The normal 32-bit-address forms are complete and a separate scan finds zero address-size-override candidates. An independent `disp:64` sweep finds 538 instructions / 213 owners image-wide and classifies all 36 actor/effect-module hits; its 12 loads contain no indirect call. `UNIT-OWNER-009` identifies owner `+0x28 == 0` as a human participant; owner `+0x5c` remains unnamed. Per-root data agreement is Medium by itself and supplies the mapping and authored equipment population, not route completeness; the shop table and formula are independently raw-byte asserted.

**Original status.** ● active

**Evidence.** [EXP-0191](../experiments/EXP-0191-itemcast-training/)

### MAGIC-ITEMKILL-117

**Item-spell delayed-kill credit is not an award for release; it consumes the victim's surviving attribution state.** The resolver leaves prior state for a null source or null source definition. A definition-bearing source with no owner clears `victim+0x40`; one with an owner writes itself, then bit-4 actors copy damage kind to `+0x48` while clear-bit actors write zero. Shipped direct-damage PointEffects and nonzero-domain AreaEffects normally replace admitted resolver state with the actual spell id, so an isolated direct spell kill reaches the actual Sphere with actor bit 4 set and current weapon skill with it clear. Point and Area tails both clear `+0x40` when the recorded caster's definition becomes null; a null caster owner retains prior state for Point but clears it for Area. Drain Life constructs neither envelope and writes no fresh attribution. Poison Cloud can seed id 8 during area application, but later continuous ticks do not refresh it. Drain and later Poison deaths can therefore inherit, lose or redirect credit according to prior and intervening combat; non-damaging offensive envelopes can overwrite or clear it too. The death processor calls `vt+0x60` once only when `victim+0x40` remains non-null. Its caller also requires a victim owner and refuses when owner `+0x5c` and multiplayer `server+0x0c` are both nonzero, but unlike damage award it has no owner-`+0x28` test. Low-type credited Humans still receive no award at the sink. `Unit::Serialize` stores `+0x40` as a raw `u32` and `+0x48` as a byte. **G2:** widening ids changes actor and save state; per-source attribution or tick refresh requires additional state. The earlier inference that portable actor identity also requires new state is withdrawn: the world LOAD lifecycle explicitly resolves the raw `+0x40` key through `R1108`, replacing a map hit and clearing a miss (`SAV-908`). The wire primitive is not the whole identity programme.

**Confidence.** **High** for each resolver writer/gate, envelope writer/clear, recipient gate and the two bypasses, all asserted from disjoint raw instructions; `MAGIC-ATTRGATE-118` owns the shape-specific tails. **Medium** for the eventual recipient after Drain, Poison or another intervening effect because runtime history decides it. Raw-pointer validity after load is unobserved.

**Original status.** ● active (amended)

**Evidence.** [EXP-0191](../experiments/EXP-0191-itemcast-training/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### MAGIC-ATTRGATE-118

**PointEffect and AreaEffect attribution are two different post-payload programmes; definition loss clears prior credit, and their null-owner outcomes differ.** A PointEffect enters the tail only when it recorded a caster and `PointEffect+0x41 != 0`; its builder sets `+0x41 = !spell+0x0a`, while cached `spell+0x0a` is one exactly for raw `Defensive == 1`. A now-null caster definition clears `victim+0x40`. Otherwise a non-null owner admits the recorded-caster and actual-id writes; a null owner leaves prior state. An AreaEffect instead excludes Wall of Earth before payload and Light after payload, then requires a nonzero byte from target `vt+0x20` and a recorded caster. A null definition or null owner clears `victim+0x40`; only a surviving owner admits the recorded-caster and actual-id writes. The virtual returns movement-domain byte `actor+0x4a`, not health, and the caller zero-extends it before testing: ordinary domains 1..3 and custom `0xfe` pass, while only zero suppresses the tail. All shipped direct-damage PointEffects are non-Defensive and ordinary shipped AreaEffect targets have nonzero movement domain, so their actual-id overwrite is normally reached. **G2:** changing `Defensive`, delivery shape, zero/nonzero movement domain or caster definition/owner lifetime can change delayed credit without changing damage; those are authored seams, while the gates and clears are engine code.

**Confidence.** **High** for both complete tails, the builder polarity, distinct clear arms and unsigned movement-domain identity, each raw-byte asserted and cross-checked against `DAT-ACT-006` and `TERR-MOVE-055`. **Medium** for shipped reachability of definition or owner loss between envelope build and payload: no such transition was observed at runtime.

**Original status.** ● active

**Evidence.** [EXP-0191](../experiments/EXP-0191-itemcast-training/)

### MAGIC-CADENCE-126

**A cast adds the Spell row's `Complication Level` only when state is `0x0d` or `0x0e` and `actor+0x64` still holds the Spell at recovery; retained order casts do, weapon-diverted caster casts do not.** After loading relax, inclusive `U[0,3]` and any equipped-Humanoid penalty, `L05832`..`L04383` tests the two states and then the pointer before calling the row parameter accessor with index 0 and adding its byte result to `actor+0x6c`. Parameter 0 is title 1 because the parameter array excludes the name column; title 1 is `Complication Level`. An ordinary retained order cast has null item pointer `actor+0x68`, so its Spell survives application and its uninterrupted interval is `charge + relax + U[0,3] + humanoidPenalty + ComplicationLevel + 2 actor ticks`. A weapon diversion enters state `0x0d` with both pointers set, then `L04547`..`L04548` clears them before recovery; its corresponding interval omits `ComplicationLevel`. Physical state 3 also skips the addend. Both shipped roots contain 28 Spell rows with distribution `3:6, 10:11, 20:9, 25:1, 30:1`. This narrows but still refutes the EXP-0067 report's statement that the column is never read on the cast path; that published report is not rewritten

**Confidence.** High for the state/pointer gates, pointer-clear ordering, accessor index and addition (named instructions and raw-byte anchors on both byte-identical executables; the two route outcomes discriminate a state-only rival) / Medium for the dual-root distribution

**Original status.** ● active

**Evidence.** [EXP-0234](../experiments/EXP-0234-action-cadence/)

### MAGIC-CADENCE-127

**Insufficient mana writes no failed-cast recovery; after a completed cast leaves `actor+0x136 = 1`, a retained refusal attempts admission on three actor ticks, skips one, then repeats.** In `R0268`, `L05833`..`L01999` compares cost against current mana and returns zero before the phase-0 arm changes phase, countdown or the stale completion byte. Re-arming progress 2 resets its counter and attempts in that tick; the next two progress ticks increment the counter to 1 and 2, bypass the completion test and attempt again. At 3, `L00094`..`L05834` consumes the still-set completion byte, clears progress and actor state, and no admission occurs; the pending order re-arms on the following tick. If the completion byte was already zero, the progress arm does not clear the state and refusal continues to be attempted each tick. No refusal path writes phase 7 or a countdown. Once admitted, an ordinary order cast uses `MAGIC-CADENCE-126`'s complication addend; a weapon-diverted cast uses its pointer-cleared form. Player and AI ownership does not select another executor after order 8 or 9 is armed

**Confidence.** High (cast admission, pending-order re-arm and progress arm 2 are read whole; their counter, latch and state writes fix both retry patterns, while the alternative failed-cast cooldown would require a phase/countdown writer absent between the compare and return) / Unknown for what every external command writer does if the retained order is replaced during these retries

**Original status.** ● active

**Evidence.** [EXP-0234](../experiments/EXP-0234-action-cadence/)

## Fire Wall and Poison Cloud overlap

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-CLOUDCLOCK-154 | A registered cloud can pulse at new counter zero; first paint, final pulse and cleanup are separate calls. | High | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-CLOUDOWNER-155 | Same-spell cell ownership does not select a single ticking area or payload. | High | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-CLOUDVISIT-156 | Fire Wall and Poison Cloud visit occupied cells without a per-target set; their inner routes give different repeat behavior. | High | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-POISONINPUT-157 | Poison Cloud supplies Token 8 to the general HP body; its parsed counter and scaled magnitude are different fields. | High / Medium | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-POISONREFRESH-158 | An incoming continuous same-id Poison effect changes only the retained counter, while first attachment copies and applies immediately. | High | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-POISONPHASE-159 | Continuous refresh can change Poison event phase, and first attachment can be followed by an actor-phase HP event in the same tick. | High / Medium | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-FIREPOISON-160 | Fire/Poison conflict removes a cell layer and preserves existing target attachments and area clocks in the bounded paths. | High / Medium | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-AREASOURCE-161 | The retained Poison producer and a target's later area-credit producer can differ. | High / Medium | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-AREAOVERLAP-162 | The finite Fire/Poison overlap population has reproducible conditional streams, not native acceptance. | Medium / Unknown | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |
| MAGIC-CLOUDEND-163 | Cloud layer removal can restore a saved cell's terrain and flags when its record becomes empty. | High | ✔ promoted | [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`) |

### MAGIC-CLOUDCLOCK-154

**A registered cloud can pulse at new counter zero; first paint, final pulse and cleanup are separate calls.** Complete `R0629..L05835` reads `+0x48` first: unset paints and returns. Otherwise old unsigned `+0x4c > 0` guards the decrement, and the new counter's modulo-16 zero branch enters `L05836`, including when the stored word is zero. The next call with old zero invokes cleanup and marks `+0x40`. With unmodified initial `V>0` and registration at relative tick 0, pulse ticks are `(V-1)%16+1`, then every 16 through `V`; count is `ceil(V/16)`, teardown tick `V+1`, inclusive driver calls `V+2`. `V=0` in forced cloud mode paints then tears down without a pulse. These are pulse-branch entries, not admitted target applications. Ring/blast behavior is outside this correction.

**Confidence.** High for this complete local branch and arithmetic. Both executable hashes agree; forced boundary cases 145–160 and ordinary selected states execute the original instructions. Synthetic forced cloud zero does not establish ordinary constructor reachability.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-CLOUDOWNER-155

**Same-spell cell ownership does not select a single ticking area or payload.** `R0647` appends areas; original `R0641` visits each remaining list member. Layer add writes the incoming pointer at `cell+0x14+4*index`, overwriting the previous same-spell slot. Query `R1081` returns that slot, and pulse `L05837` tests only nonzero before invoking `R0267` with the ticking area itself. Its own `+0x44` supplies the inner payload. Thus an area can apply where another area owns the layer within its scan square, even outside its own painted pattern. Cleanup `L05838` requires pointer equality; removal clears the slot and does not reinstate a saved prior owner. An older area's clock can continue after a newer owner removes their common cell gate.

**Confidence.** High for the named local list, stores and predicates, with same-spell/shifted/reverse-strength conditional witnesses. Whole-game producer lifetime and arbitrary insertion interleaving remain Unknown; no global overlapping-spell rule follows.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-CLOUDVISIT-156

**Fire Wall and Poison Cloud visit occupied cells without a per-target set; their inner routes give different repeat behavior.** In `R0629`, dx is the outer loop and dy the inner, both from `-radius` through `radius`. Every nonzero spell-layer query reads cell occupant `+4` through `R0035` and invokes the same ticking area's `R0267`. Only outer Token id 2 takes the footprint-square division; ids 3 and 8 dispatch inner virtual `+0x3c` unchanged. Fire's DirectDamage route can write HP for each reached cell; Poison's ordinary timed route first attaches then refreshes the same-id effect on later visits. Case 001 has two Fire areas and one occupied cell with 30 HP stores; case 002 aliases a 2×2 target footprint, of which the wall reaches two cells, with 60. This is not a claim that every 2×2 placement receives four hits.

**Confidence.** High for the complete local loops and id-2-only normalization gate. Conditional store counts hold with the fixture's nonterminal stationary target and controlled damage resolver; native HP totals and later target transitions are Unknown.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-POISONINPUT-157

**Poison Cloud supplies Token 8 to the general HP body; its parsed counter and scaled magnitude are different fields.** The selected lawful Poison effect text runs through complete parser `R1000`, mode parser `R0857` and integer helper `L05839`, conditionally yielding kind `+0x3c=6`, mode `+0x3d=2`, magnitude `+0x40=-2`, remaining `+0x42=128`. The numeric duration is shifted left four at `L05840`; Poison's builder scales only magnitude at `L05841..L05842`, stamps spell Token id 8 at `L05843`, and does not overwrite that counter. The common area tail copies envelope source `+0x3c` to inner Effect `+0x44` at `L05415`. Original virtual dispatch/copy/tick then reaches ANIM-074's Token-8 HP/notification branch. This identifies a spell producer, not a separate item equip/use producer.

**Confidence.** High for the literal stores and static builder-to-attachment dispatch. Medium for the joined valid-text parse: CString, valid-decimal sscanf and constructor support are explicit substitutes. Both input record hashes and conditional parses agree; malformed text, Stone Curse numeric parsing and a native cast witness remain outside scope.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-POISONREFRESH-158

**An incoming continuous same-id Poison effect changes only the retained counter, while first attachment copies and applies immediately.** `R0612` searches by Token `+0x0c`; when absent it calls copy `R1001`, which explicitly copies kind, mode, packed magnitude/counter, id and source `+0x44`, invokes copied effect virtual `+0x40` at `L05844`, then appends at `L05845`. When the id is already present and the incoming continuous bit is set, `L05846` is the only payload write: old `+0x42 = incoming +0x42`. It neither applies immediately in that repeat branch nor changes retained magnitude, source, kind or mode. The target therefore has one timed Poison record, not one per area, producer or cell. A later actor phase can still apply that record in the same world tick (`MAGIC-POISONPHASE-159`).

**Confidence.** High for the complete copy/search/refresh branches and finite same/different-source, equal/unequal-magnitude conditional executions. Only the continuously refreshed Poison record and ordinary target list state are claimed; arbitrary effect combinations and native source destruction remain Unknown.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-POISONPHASE-159

**Continuous refresh can change Poison event phase, and first attachment can be followed by an actor-phase HP event in the same tick.** The named world caller `R0193` invokes entity/area driver `R0426` at `L01866`, then actor driver `R0427` at `L01867`. The actor processes attachments before its health/action branch; terminal act `0x10` skips its tick. `R0673` tests old remaining counter modulo 8 before decrement. An attachment/refreshed counter of 128 is therefore eligible in the later actor phase of that same call order. First attachment has already applied once; a continuous repeat itself only refreshed. Staggered area pulses can reset the counter so intervals between actor-phase HP events are not an invariant eight ticks. At unrefreshed count 1, decrement to zero marks removal without another modulo-zero pulse.

**Confidence.** High for the named driver order and local counter branch. Medium for the complete conditional streams: the probe follows this driver path and skips later target health/action/regen processing. No native scheduling census, target death transition or arbitrary other-list population is claimed.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-FIREPOISON-160

**Fire/Poison conflict removes a cell layer and preserves existing target attachments and area clocks in the bounded paths.** Complete `R0636` queries the opposing spell layer: Fire Wall removes a present Poison layer; Poison encountering Fire removes its own layer. Both creation orders therefore leave Fire as the cell's admitted spell of this pair. The conflict calls per-cell removal, not removal of the other area list member or revocation of an attached Poison. In case 167 Poison attaches at tick 16, Fire clears its target-cell layer at 21, the attachment runs to expiry 143, and the original Poison area still reaches its own teardown at 241. Same-type overwrite behavior is separately `MAGIC-CLOUDOWNER-155`; other spell conflicts retain MAGIC-MAPLAYER-040's older confidence.

**Confidence.** High for the selected complete local conflict branches and their scope. Medium for the joined stationary-target event streams; native world evolution, other conflicts and arbitrary geometry are not established.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-AREASOURCE-161

**The retained Poison producer and a target's later area-credit producer can differ.** Effect copy carries source `+0x44`; continuous refresh does not rewrite it. Separately, after payload application, `R0267`'s established movement-domain/source/owner gates can write area `+0x3c` to target `+0x40` and inner Token id to target `+0x48` at `L05469`/`L05470`. In case 144, B's tick-21 refresh retains magnitude -4 and source A, then target credit reads B/id 8; the later actor HP event still uses A. Token-8 nonzero HP application consults retained `+0x44` before its `vt+0x64` call (`L05847`), independently of those target credit fields. MAGIC-ITEMTRAIN-116, MAGIC-ITEMKILL-117 and MAGIC-ATTRGATE-118 remain the authority for their later gates.

**Confidence.** High for the original copy/refresh/credit stores and conditioned source distinction. Medium for the joined event history; positive-health valid producers and admitted owner/domain gates are fixture conditions, and training sinks are substituted. Eventual native kill credit, awards and destroyed-source lifetime remain Unknown.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-AREAOVERLAP-162

**The finite Fire/Poison overlap population has reproducible conditional streams, not native acceptance.** EXP-0332 retains 178 lifecycle cases plus one parser vector per locale, over exactly 32 complete lifecycle bodies and 16 immediate helpers. Both locale streams agree. It distinguishes independent area clocks, latest cell ownership, cell visits, first attachment, continuous refresh, phase reset, cross-spell suppression and separate expiry under the recorded inputs. The probe admits only listed original instructions or explicit environment substitutes; complete body listings and executed-address coverage are separate. Map records, target/source state, object construction and some library calls are synthetic; Fire RNG/resolver, notifications and training are substituted, and target health/action/regen after the attachment loop is skipped. Powers are 0 and 30: zero is a disclosed deviation from the preregistered positive-power population.

**Confidence.** Medium for the finite joined conditional result. Unknown for an observed native event stream, all legal power/placement populations, native RNG totals, terminal transitions, producer destruction and broader scheduling. A native witness requires a proven isolated write boundary and was not performed.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

### MAGIC-CLOUDEND-163

**Cloud layer removal can restore a saved cell's terrain and flags when its record becomes empty.** In the complete `R1078..L05848`, pointer equality at `L05838` gates slot clearing, layer recount and recomputation. The empty-record test examines occupants `+4/+8/+0xc/+0x10`, layer count `+2` as well as the byte `+0x2c` at `L05849..L05850`. On the empty branch, `L05851` writes the saved payload byte to `map+key`, and `L05852` writes the saved flag byte to `map+key+0x10000`; later code preserves the existing fire-plane bit and unlinks the record. The old MAGIC-AREAEND-041 universal no-terrain-write wording omitted this helper tail. Clearing a last cell layer does not restore an older area pointer or revoke the target's separate Poison attachment.

**Confidence.** High for the complete local identity/empty-record predicates and named restoration stores. The occupied conditional fixtures deliberately retain their records, so empty-record deletion and its allocator/free helper effects are static-only, with native consequences unobserved.

**Original status.** ✔ promoted

**Evidence.** [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) (`evidence/`, `RESULT.md`)

## Fire Wall and Poison Cloud overlap (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-DELIVERY-170 | A simulation SpellTransport delivers by a saved signed countdown, not by the visible projectile reaching a collision point. | High | ● active | [EXP-0362](../experiments/EXP-0362-cast-delivery/) |
| MAGIC-CASTCLOCK-171 | Cast admission, prepared simulation effects and client flight use different clocks. | High / Medium | ● active | [EXP-0362](../experiments/EXP-0362-cast-delivery/) |
| MAGIC-TELEPORT-174 | A destination veto can leave an admitted teleport in place after mana was spent and cast presentation was produced. | High / Medium | ● active | [EXP-0363](../experiments/EXP-0363-blocked-teleport/) |

### MAGIC-DELIVERY-170

**A simulation SpellTransport delivers by a saved signed countdown, not by the visible projectile reaching a collision point.** ConstructorR0633 stores distance/speed into word+4c; Apply's delivery2 arm wraps its prepared payload and enqueues the transport. TickR0634 subtracts1, hands child+44 or alternate+48 to the effect collection when signed remaining<=0, clears both references and sets retirement+40. Its complete43-instruction body never moves Position or calls the client projectile. The PointEffect consumer separately calls nested effect virtual+3c on an admitted target. Both installed28-row spell tables assign delivery2 to IDs1,2,13,14 and delivery1 to the other24.

**Confidence.** High for the named instruction/data contract and 24 conditional paired serializer/load/tick cases; native scheduler ordering and moving/dead target outcomes Unknown

**Original status.** ● active

**Evidence.** [EXP-0362](../experiments/EXP-0362-cast-delivery/)

### MAGIC-CASTCLOCK-171

**Cast admission, prepared simulation effects and client flight use different clocks.** R0268 spends mana and sendsR0617/R0618's animation packet; its distance/speed or fixed5 value is a message field. Actor phase5 later reachesR0002/R0003 after charge. Apply independently builds a transport with distance/speed and overrides IDs13/14 to 10 atL03086. ID14 selects/applies each victim during admission atL03027, and the later unit-target wrapper excludes that ID to avoid duplication. MAGIC-CAST-003's queued-apply description and MAGIC-SING-019's universal no-queue shorthand were too broad; MAGIC-CASTTICK-030's call timing is payload preparation, not universal HP-change timing. Client MAGIC-CASTSPAWN-033 remains a separate presentation object.

**Confidence.** High for bounded call/operand ordering, table selection and transport consumer; Medium for end-to-end original cast timing without native process observation

**Original status.** ● active

**Evidence.** [EXP-0362](../experiments/EXP-0362-cast-delivery/)

### MAGIC-TELEPORT-174

**A destination veto can leave an admitted teleport in place after mana was spent and cast presentation was produced.** Mage book-cost subtractionL05055/L05057 precedes application. Teleport armL05853 callsR1055 and ignores its return before issuing the0x20 update atL05854. Relocation intersects map plane+20000 with mover+5 over the actor footprint; a matching bit atL05855 or the laterR0048 refusal returns before Position storesL05856/L05857. Neither early path refunds mana; value1 is also returned on this veto. The coordinate cast packet has picture60 for ID26 and the distinct client path owns its two sprites (MAGIC-CASTSPAWN-033). Fourteen paired original-fragment cases distinguish passable/masked/unmasked/placement-refused and mana59/60/100 with synthetic cost60; branch inversion changes the blocked result.

**Confidence.** High for named conditional instruction ordering and mutations; Medium for whole-game inference from separated fragments and explicit service stubs; native selected-tile/pixel acceptance Unknown

**Original status.** ● active

**Evidence.** [EXP-0363](../experiments/EXP-0363-blocked-teleport/)

## Fire Wall and Poison Cloud overlap (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-TARGETID-182 | PointEffect::Tick's attribution tail (`R0266`) reads its own `+44` (target) six times across its branches; two of those reads feed an unconditional dereference with no null check anywhere between function entry and the dereferencing ... | High | ● active | [EXP-0371](../experiments/EXP-0371-saved-effect-identity/) |
| MAGIC-TICKGATE-183 | Within the graph this experiment traced, cast-time construction is the only target-specific gate on the population PointEffect::Tick reaches; a separate, target-blind exclusion also exists, and one caller of the registrar ... | High | ● active (partially retracted) | [EXP-0371](../experiments/EXP-0371-saved-effect-identity/); [EXP-0388](../experiments/EXP-0388-effect-construction/EXP-0388.md) |
| MAGIC-SIBLINGREF-184 | AreaEffect's own `+44` and SpellTransport's own `+44`/`+48` are read by different functions than PointEffect's raw target dword, through the standard typed-reference mechanism rather than a raw identity dword — but the guard contrast ... | High / Medium | ● active | [EXP-0371](../experiments/EXP-0371-saved-effect-identity/) |
| MAGIC-187 | `R0641` walks a container whose `+0x4`/`+0x8` hold head/tail node pointers and whose nodes hold next/payload at `+0x0`/`+0x8`, through the standard GetHeadPosition/GetNext idiom that advances the walk's position before it reads ... | High / Medium / Unknown | ● active | [EXP-0374](../experiments/EXP-0374-ticked-effect-population/) |
| MAGIC-188 | `R0641` reads exactly two fields on a member before dispatching Tick — `member+0x3c` and the member's own vtable pointer at `member+0x00`, which the dispatch itself loads — plus, conditionally, `(*(member+0x3c))+0x14`, a field ... | High | ● active | [EXP-0374](../experiments/EXP-0374-ticked-effect-population/) |
| MAGIC-189 | `R0641`'s own body writes exactly one field on any member object — `member+0x3c = 0`, conditionally — and no other instruction in its complete 50-instruction body (`R0641`-`L05858`, 148 contiguous bytes) touches member ... | High | ● active | [EXP-0374](../experiments/EXP-0374-ticked-effect-population/) |
| MAGIC-190 | `R0641`'s own body never reads or writes a member's `+0x44` anywhere, and gates nothing on it; a member whose `+0x44` is null — including one left null by `SAV-1010`'s load-time identity-map miss — is not filtered, skipped, or ... | High | ● active | [EXP-0374](../experiments/EXP-0374-ticked-effect-population/) |
| MAGIC-197 | `R0634` (SpellTransport's own Tick) computes the identical shared-list container operand at both of its own registrar calls that `R0003` computes at both of its own — closing `MAGIC-187`'s own left-open Unknown — and ... | High | ● active | [EXP-0378](../experiments/EXP-0378-transport-scheduling/) |
| MAGIC-202 | After a ROM1 archive LOAD resolves a `SpellEffect`-lineage record, the archive's own load arm inserts it into a container by the identical generic primitive, and the identical operand-computation formula, `MAGIC-187`/`MAGIC-197`'s own ... | High / Medium / Unknown | ● active | [EXP-0380](../experiments/EXP-0380-transport-rebind/EXP-0380.md), `evidence/container-serialize-R1109-body.txt`, `evidence/wrapper-R1122-body.txt`, `evidence/document-loader-body.txt`, `evidence/container-serialize-call-sites.txt`, `evidence/store-load-split.txt`, `evidence/manager-lazy-alloc-and-call-window.txt`, `evidence/postload-dispatch-R1126-body.txt`, `evidence/owner-R1125-reachability.txt`, `evidence/readobject-wrapper-R1124-body.txt`; MAGIC-187, MAGIC-197, MAGIC-TICKGATE-183, SAV-DOC-053, SAV-774, SAV-1030, SAV-TOKENLOAD-093, SAV-TOKENLOAD-094 |
| MAGIC-209 | The STORE arm of `R1109` (the same container-level `Serialize` `MAGIC-202` already reads on its LOAD side) writes a `u32` element count and then calls the identical generic write wrapper (`R1110`) once per element, with no ... | High / Medium / Unknown | ● active | [EXP-0381](../experiments/EXP-0381-transport-store/EXP-0381.md), `evidence/container-serialize-R1109-body.txt`, `evidence/store-head-R1117-body.txt`, `evidence/store-getcount-R1127-body.txt`, `evidence/store-getnext-R1128-body.txt`, `evidence/store-writeu32-R0685-body.txt`, `evidence/writeobject-wrapper-R1110-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-187, MAGIC-202, SAV-982, SAV-PLAYER-028, SAV-1042, SAV-1047 |
| MAGIC-213 | `PointEffect`'s own real destructor deletes exactly one field, `+0x48`; `AreaEffect`'s own real destructor deletes exactly one field, `+0x44`; each is the class's own "inner Effect" payload, never the sibling offset the other class ... | High | ● active | [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/pointeffect-realdtor-L05905-body.txt`, `evidence/areaeffect-scalardtor-L05956-body.txt`, `evidence/areaeffect-realdtor-L05962-body.txt`, `evidence/spelleffect-realdtor-L05960-body.txt`, `evidence/furtherbase-realdtor-R1130-body.txt`, `evidence/vtable-slots.json`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-187, MAGIC-197, SAV-CLASSSER-174, SAV-CLASSSER-175, MAGIC-SIBLINGREF-184 |
| MAGIC-214 | `SpellTransport`'s own real destructor treats `+0x44` and `+0x48` as two independent owning pointers, each conditionally deleted through the pointee's own vtable `+0x04` slot and then cleared, in the identical single-field pattern ... | High / Medium | ● active | [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/spelltransport-realdtor-R1062-body.txt`, `evidence/spelltransport-tick-R0634-body.txt`, `evidence/spelltransport-ctor-R0633-body.txt`, `evidence/registrar-wrapper-R1113-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-213, MAGIC-197, MAGIC-TICKGATE-183, MAGIC-215, SAV-CASTCONT-1006, SAV-CLASSSER-175, SAV-1042 |
| MAGIC-215 | `SpellTransport`'s own constructor is the only construction site this experiment locates, and it never stores its own object argument at `frame+8` into more than one owning field: `+0x44` is written unconditionally from the argument ... | High / Medium | ● active | [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/spelltransport-ctor-R0633-body.txt`, `evidence/spelltransport-tick-R0634-body.txt`, `evidence/registrar-wrapper-R1113-body.txt`, `evidence/call-site-census.json`, `evidence/unresolved.json`; MAGIC-TICKGATE-183, MAGIC-197, MAGIC-187, MAGIC-214 |
| MAGIC-216 | The two primitives `MAGIC-187` left unread inside the removal path do exactly what that row inferred from their call shape: `R1111` locates a node by its own payload pointer without touching the payload, and `R1112` unlinks ... | High / Medium | ● active | [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/removal-findnode-R1111-body.txt`, `evidence/removal-unlinknode-R1112-body.txt`, `evidence/removal-bulk-clear-R1131-body.txt`, `evidence/removal-wrapper-R1118-body.txt`, `evidence/removal-unlink-dispatch-R1119-body.txt`, `evidence/call-site-census.json`, `evidence/unresolved.json`; MAGIC-187, MAGIC-189 |
| MAGIC-217 | `PointEffect`'s and `AreaEffect`'s own per-tick attribution tails dereference the recorded caster's own pointee exactly three times each, in the identical order: first, through the "definition" test callee `R0416`, with no ... | High / Medium | ● active | [EXP-0383](../experiments/EXP-0383-effect-token-reference/EXP-0383.md), `evidence/pointeffect-tick-attribution-R0266-body.txt`, `evidence/areaeffect-tick-attribution-R0267-body.txt`, `evidence/caster-definition-test-R0416-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; SAV-1011, MAGIC-ATTRGATE-118, MAGIC-TARGETID-182, MAGIC-188, MAGIC-189, MAGIC-AREATICK-036, SAV-1056 |

### MAGIC-TARGETID-182

**PointEffect::Tick's attribution tail (`R0266`) reads its own `+44` (target) six times across its branches; two of those reads feed an unconditional dereference with no null check anywhere between function entry and the dereferencing instruction.** Tick (vtable slot `+0x18`, `R1065`) is a bare one-instruction thunk that always calls `R0266(this)`; inside it, after an unrelated cleanup of the attached Effect payload's own source field (that Effect object's own `+44` at `this+0x48`, a different field on a different object — the mechanism `formats/magic/application.md`'s "Effect ownership and credit" section already documents), the raw sequence `L05859..645` loads `this+44` and dereferences it (the load of the target's vtable pointer through that field at `L05860`) to fetch the target's vtable before calling `+0x2c`, with no comparison or test of `this+44` anywhere before that load — unguarded site 1. Its return value gates the rest of the function: nonzero falls to `L05861`, where `this+44` is reloaded into a stack local (`L05862`) and that local is null-checked (`L05863`) and its own `+0x14` checked (`L05864`) before any further use — a guarded read, not a third unguarded site. Zero instead branches to `L05865`, where `this+44` is read and dereferenced a second time, again unconditionally (`L05866..722`, the load of the vtable pointer through that field), to invoke vtable `+0x34` — unguarded site 2, reached only on this path, not the entry path. That call's own return value gates a final read: nonzero reaches `L05867`, where `this+44` is read once more and passed, unchecked, as an argument to the Effect payload's own `+0x3c` virtual call (`L05868`) — read but never itself compared to zero on this path.

**Confidence.** High for the complete instruction sequence and the presence/absence of a comparison or test on `this+44` at each read, cross-checked against the Tick thunk's own body (`evidence/ranges.txt:271-387`). Whether the game ever reaches either unguarded site (`L05860`, always; `L05869`, when the first virtual call returns zero) with `this+44` null — i.e. whether a load-time identity-map miss (`SAV-1010`) on an in-flight PointEffect is reachable in any admitted native save/load population — is Unknown; that would fault at whichever site is reached, a falsifiable, unrun prediction naming two possible fault addresses, not one.

**Original status.** ● active

**Evidence.** [EXP-0371](../experiments/EXP-0371-saved-effect-identity/)

### MAGIC-TICKGATE-183

**Within the graph this experiment traced, cast-time construction is the only target-specific gate on the population PointEffect::Tick reaches; a separate, target-blind exclusion also exists, and one caller of the registrar (`R0641`) was not read.** Live constructor `R1046` stores its target argument into `this+44` unconditionally (`L05870`) and immediately dereferences it (`L05871..5e0`, reading target `+0x10` to copy its Position into the effect Position through `L05872`; the former target-registration gloss is retracted by SAV-1068) with no null check of its own; the target-typing this depends on is `MAGIC-SHAPE-008`'s, not independently re-derived here. The documented guarantee that a cast-time PointEffect always has a target is enforced earlier, at attachment (`formats/magic/application.md`'s "REQUIRES a target unit" / dropped-construction rule), not inside this constructor. `MAGIC-AREATICK-036` already reads that AreaEffect "is added to that list at `L05395` and `L03053`"; the intervening bytes (`L05873..L05874`) show `L05395` is fed by one local variable that both the PointEffect branch (`R1046`'s return copied to it at `L05875..L05876`) and the AreaEffect branch (`R0652`'s return copied the same way at `L03047..L13149`) write identically, so `L05395`'s call of R1113 is the same registration call for either constructed object, taken when `getParam(row,5)==1` (`L05877`). When the value is 2 instead, the object instead becomes the `+0x44` child of a new SpellTransport (`+0x48` remains zero on this constructor path, MAGIC-215) (`R0633` at `L03052`), and the transport itself is what reaches `L03053`'s call of R1113, matching `MAGIC-DELIVERY-170`'s later hand-off through the same transport Tick. When the value is neither 1 nor 2, `L05878` jumps straight to the function's epilogue (`L05879`): no transport construction and no registration occur in that suffix. The PointEffect/AreaEffect was already constructed before delivery selection; the former no-construction clause is retracted by MAGIC-225. This is a registration exclusion gated on the delivery-system field, not on the target. `R0647`, the append `R1113` calls, is a generic doubly-linked insert with no field access on the object being appended. `MAGIC-AREATICK-036` names `R0641` as a walker of the same shared list; this experiment did not disassemble it, so whether it applies any further filter before invoking a member's Tick is Unknown. So within the traced graph, the only target-specific exclusion is the one-time cast-time attachment gate; a PointEffect that later has its own `+44` rewritten to null by a load-time identity-map miss (`SAV-1010`) re-enters the same, otherwise target-unfiltered Tick population, whether it was originally admitted directly or by way of a SpellTransport.

**Confidence.** High for the constructor's own unconditional store/dereference, the shared-local-variable proof that PointEffect and AreaEffect reach the identical `R1113` call, the delivery-field exclusion at `L05878`, and the registrar's generic, field-blind append — each read as complete bodies. That the attachment gate is the only target-specific exclusion is bounded to the functions this experiment read; `R0641` in particular was not traced and could filter further. Reachability of the null-miss case itself is the same Unknown as `MAGIC-TARGETID-182`.

**Original status.** ● partially retracted (Position/admission wording; target gate retained)

**Evidence.** [EXP-0371](../experiments/EXP-0371-saved-effect-identity/); [EXP-0388](../experiments/EXP-0388-effect-construction/EXP-0388.md)

**Amended.** The table ledger carried the status "● partially retracted (Position/admission wording; target gate retained)". `retracted.md` records a correction against this claim.

### MAGIC-SIBLINGREF-184

**AreaEffect's own `+44` and SpellTransport's own `+44`/`+48` are read by different functions than PointEffect's raw target dword, through the standard typed-reference mechanism rather than a raw identity dword — but the guard contrast with PointEffect is narrower than a single function, and at least one further unguarded AreaEffect reader exists.** `MAGIC-AREAAPPLY-038` already reads AreaEffect's shared per-target tail `R0267` (called from AreaEffect's own Tick `R0629` and from `R0639`/`R0630` — 7 call sites, 3 owners) returning on a null caller-supplied target parameter at `L05410`. Inside that function, `this+0xc==2` (`L05880..3c1`) diverts to a path (`L05419`) that reads `this+44` but only passes it, unchecked, as a raw argument to `R1068` (no dereference of it in that arm); the ordinary `this+0xc!=2` fall-through (`L05412`) is where `this+44` is read twice and dereferenced unconditionally (`L05881`, the vtable-pointer load) to invoke the Effect payload's own virtual `+0x3c(target)` — an unguarded read, not a guarded one, contrary to this row's own earlier framing of the branch. That field is `SAV-CLASSSER-175`'s typed `ar<<*(+0x44)` object reference, resolved by the ordinary archive-reference loader on LOAD, not by `R1108`'s raw-dword map lookup — the fact this row adds is that connection between the serialization-side typing (`SAV-CLASSSER-174`/`-175`) and the consumer-side reader `MAGIC-AREAAPPLY-038` already named, not a guard-strength contrast with `MAGIC-TARGETID-182`. A separate AreaEffect vtable method, `R1114` (vtable `L05352` slot `+0x38`), also dereferences `this+44` unconditionally, with no null check anywhere in its own body (`L05882..371`); this experiment did not run a caller census on it, so whether it is ever reached is Unknown. SpellTransport's own `+44`/`+48` (`SAV-CASTCONT-1006`) are typed nested SpellEffect/AreaEffect references; besides its own Tick (`R0634`), which hands the non-null one to the same registrar (`R1113`) that `MAGIC-TICKGATE-183` shows a directly cast PointEffect or AreaEffect also reaches, `SAV-CASTCONT-1006`'s own "child repair" language already documents a post-load hook (`R1115`) touching the same fields; AreaEffect has an analogous post-load hook, `R1116`. No image-wide reader census was run for either sibling's own `+44`/`+48`; "read only by its own Tick" is retracted as overstated. One registrar and one Tick population reach all three siblings, but three different readers and two different persistence mechanisms serve their own `+44`.

**Confidence.** High for the reader-identity and mechanism contrast (typed archive reference vs. raw identity dword), built on already-published `MAGIC-AREAAPPLY-038`/`SAV-CLASSSER-175` plus this experiment's own `MAGIC-TARGETID-182`/`MAGIC-TICKGATE-183`, with the connecting fact (registrar/mechanism identity across all three) newly read here. Medium, not High, for any guard-strength contrast between the siblings: AreaEffect's own dereference at `L05881` is itself unguarded, and `R1114` is a second, uncensused unguarded reader. Whether AreaEffect's or SpellTransport's own typed-reference resolution can itself produce an unrepaired dangling or null reference under any admitted LOAD population is Unknown; this experiment traced consumer identity and mechanism, not that loader's own failure modes or a complete reader population.

**Original status.** ● active

**Evidence.** [EXP-0371](../experiments/EXP-0371-saved-effect-identity/)

### MAGIC-187

**`R0641` walks a container whose `+0x4`/`+0x8` hold head/tail node pointers and whose nodes hold next/payload at `+0x0`/`+0x8`, through the standard GetHeadPosition/GetNext idiom that advances the walk's position before it reads the member being processed; it removes a member from inside its body, by unlink-then-destroy, never later.** New here are the container and node layout, the advance-before-return idiom, the removal call chain and its `+0x04` slot identification, and the two-owner registrar census. Three of the facts this row restates are published: `MAGIC-AREATICK-036` carries the tick chain into this function by its instruction addresses (`L01866`, `L03066`), names `L05395`/`L03053` as the two sites that add an effect to the list, and places the removal at `L05392` inside this body gated on `member+0x40`; `MAGIC-AREAEND-041` publishes that same in-body reap together with the three Tick-callee end sites (`L05449`, `L03056`, `L05358`) that set `+0x40`. `R0647` (the append `MAGIC-TICKGATE-183` reads as "a generic doubly-linked insert with no field access on the object being appended") writes the layout: container `+0x8` (tail) `!=0` links `oldTail->+0x0 = newNode` (`L05883`) or, on an empty list, sets container `+0x4` (head) `= newNode` (`L05884`); either way container `+0x8 = newNode` (`L05885`) and `newNode->+0x8 = payload` (`L05886`); the allocator `R0651` independently confirms the same two node fields, initializing the carved node's `+0x0` to `0` and `+0x4` to the old tail (`L05887`/`L05888`), a field this walk never reads. `R0648` ("begin", called at `L05889`) stores the container into the local iterator's slot `0` and, via `R1117` (`return *(x+4)`), the head node into slot `4`; `R0650` (the shared core both `R0648` and `R0649`/"GetNext", called at `L05890`, delegate to) reads `*posRef` as the current node, writes `*posRef = curNode->+0x0` (advances to next) **before** reading and returning `curNode->+0x8` (payload) — the walk's live position has moved past a node before that node's loop-body processing (Tick, possible removal) runs, and the local iterator object is never passed into the removal call chain (only the container and the payload pointer are, at `R1118`) — no aliasing path exists for this function's same-pass removal, which always targets the member it is currently processing (`frame−0x4`), to corrupt the walk, independent of what the unread primitives below do internally. Removal of some *other* member during the same pass, for instance by Tick's callee (whose body is out of scope here), is **not** covered by that negative: the iterator's live position is the next node, and nothing read here excludes that node being the one removed. Removal happens at `L05392`, inside this function's body, calling `R1118(container, member)` on the identical pointer the walk uses: `L05891`/`L05892` recomputes `this+0x4`, the same value `L05893`/`L05894` computed and `L05895` pushed into `R0648`. `R1118` stores its incoming `this` at `frame−0xc` (`L05896`), loads exactly that back (`L05897`) and hands it to `R1119` (`L05898`), so the walk, the wrapper and the unlink routine all hold one and the same container pointer. `R1119` pushes `0` and the member pointer and calls `R1111` with that `this` (`L05899`), then, only when the result is non-null, pushes the result and calls `R1112` with the same `this` (`L05900`); **both callee bodies are unread, so that this pair unlinks the node is inference from the call shape — a lookup whose non-null result is handed to a second call on the same container — not a read fact.** `R1118` then unconditionally dispatches `member`'s vtable `+0x04` slot with argument `1` (the argument 1 supplied at `L05901`, the virtual call of slot `+4` at `L05902`), the scalar-deleting-destructor convention. For PointEffect specifically this is confirmed, not inferred: `EnumRefs callto:R1120` returns 1 hit / 0 owners, the sole reference the `.rdata` slot at `L05903`, which is `L05356+0x4` — PointEffect's vtable, established by `SAV-1010`'s citation of that same vtable's `+0x24` slot (`L05904`) and `MAGIC-TARGETID-182`'s of its `+0x18` slot (`R1065`) — and `R1120`'s body matches the exact shape (a call of L05905 for the real destructor, a mask of the incoming argument with 0x1, a conditional call of R1121) that `TRIG-TAKEITEM-038` independently found for four unrelated item vtables, both naming the identical operator-delete target `R1121` — the only row any ledger returns for that address. Entry: `R1113` (the registrar) has exactly 2 owner functions image-wide (`EnumRefs callto:R1113`: 4 hits / 2 owners); this experiment traced one, `R0003`, at both its call sites (`L05395`, `L03053`) and found its container operand computed as `*(*(X+0x14)+0x4)+4` with `X = *(L00285)` — and found the walk's operand, reached through `R0193(this=S) -> R0426(*(S+0x14)) -> R0641(*(*(S+0x14)+0x4))`, is the identical formula applied to `X = S` (`R0193`'s `this`). The second registrar owner, `R0634`, is not a fresh unknown: `MAGIC-SIBLINGREF-184` names it as SpellTransport's Tick, "which hands the non-null one to the same registrar (`R1113`) that `MAGIC-TICKGATE-183` shows a directly cast PointEffect or AreaEffect also reaches," and `MAGIC-DELIVERY-170` independently describes the same address's body handing a child off "to the effect collection" on delivery completion — this experiment did not re-derive `R0634`'s operand-computation instructions (a search over `claims/*.md` for `R0634`, `L03083`, `L03084` — `evidence/claim-search.txt` — surfaces these two rows and no operand-level detail), so whether its container operand matches the identical offset formula found for `R0003` remains unconfirmed for this owner specifically.

**Confidence.** High for the container/node layout (four independent writers/readers agree: insert, allocator, begin, GetNext), the advance-before-return idiom, the removal call chain's instructions through `R1118` and `R1119`, and the PointEffect-specific `+0x04` dispatch, each read as complete bodies, the last cross-corroborated by `TRIG-TAKEITEM-038`'s identical-pattern finding for four unrelated item vtables. That the two unread primitives `R1119` calls unlink the node is inference from the call shape, marked as inference in the row and on the canonical page, not part of this High. Medium for treating the registration and the walk as the same list at runtime: the two chains compute an identical offset formula, but this experiment read no caller of `R0193` and ran no `callto:` census on it, so it did not re-derive that function's `this`. Searching the pinned pre-publication ledgers for `R0193`, `R0426` and `L00285` returns four active rows that do bear on the identity: `UNIT-GATE-012` names `[L00285]` the server singleton; `SESS-TICK-004` places the `server+0x04` sub-tick increment at the head of `R0193`; `SESS-TICK-006` names `R0426(server+0x14)` — the exact hop traced here — as that function's per-sub-tick work; and `SAV-657` publishes `[L00285]+0x14` as the constructor-written literal `L05906` (`L05907`), which would make both formulas resolve through one fixed static object. Those four leave little room for the alternative that `R0193`'s `this` and `*(L00285)` are two distinct, same-shaped objects, but they are other experiments' evidence cited here, not this experiment's discrimination — hence Medium. Unknown: the internal unlink mechanics inside `R1111`/`R1112` (unread bodies); whether `R0634` (named by `MAGIC-SIBLINGREF-184`/`MAGIC-DELIVERY-170` as SpellTransport's Tick, feeding the same-named registrar) computes the identical offset formula this experiment found for `R0003`, which was not re-derived here.

**Evidence.** [EXP-0374](../experiments/EXP-0374-ticked-effect-population/)

### MAGIC-188

**`R0641` reads exactly two fields on a member before dispatching Tick — `member+0x3c` and the member's own vtable pointer at `member+0x00`, which the dispatch itself loads — plus, conditionally, `(*(member+0x3c))+0x14`, a field on whatever object `+0x3c` points at rather than on the member; neither of the two branches those `+0x3c`-derived reads control skips the Tick call; every member the walk reaches is Tick-dispatched unconditionally, so this function applies no filter beyond `MAGIC-TICKGATE-183`'s already-named cast-time attachment gate.** `L05908`/`L05909` (the field `+0x3c` compared with zero, jumping to `L05391` on zero) and, on the non-null arm, `L05910`/`L05911` (the field `+0x14` of the pointed-to object compared with zero, jumping to `L05391` on non-zero) are the walk's only two conditional branches between a member's arrival at the top of the loop body and the Tick dispatch; both targets, and the straight-through path when neither branch is taken, converge at the identical instruction, `L05391`, which reloads the member pointer from its frame local; the vtable load is the next instruction, `L05912`, `L05913` sets the `this` argument between it and the dispatch, and the dispatch is `L03082`, the virtual call of the slot at offset 0x18. The two `+0x3c`-derived reads gate only whether `member+0x3c` is zeroed first (`MAGIC-189`), never whether Tick executes; the `member+0x00` read is the dispatch's own operand fetch and controls no branch. This resolves the Unknown `MAGIC-TICKGATE-183` names verbatim — "whether it applies any further filter before invoking a member's Tick" — for this function specifically: no. The walk's own loop-termination check (the fetched member pointer is null) ends the walk when the list is exhausted; it is not a per-member skip, since null is never a member.

**Confidence.** High — the complete set of conditional branches between per-iteration entry and the Tick dispatch was read, and both targets traced to the identical call instruction. Bounded to this function's own body, per `MAGIC-TICKGATE-183`'s own framing: whether some other function excludes a member from reaching the list at all is the entry question (`MAGIC-187`), not this one.

**Original status.** ● active

**Evidence.** [EXP-0374](../experiments/EXP-0374-ticked-effect-population/)

### MAGIC-189

**`R0641`'s own body writes exactly one field on any member object — `member+0x3c = 0`, conditionally — and no other instruction in its complete 50-instruction body (`R0641`-`L05858`, 148 contiguous bytes) touches member memory as a write.** The function's sole write through the member pointer is `L05914` (a 32-bit store of zero at `member+0x3c`), reached only when `member+0x3c != 0` and `(*(member+0x3c))+0x14 == 0` (`MAGIC-188`'s two gating reads). Every other member-touching instruction in the body is a read — `L05908`/`L05915`/`L05910` (`+0x3c`, then `(*+0x3c)+0x14`), `L05912` (the vtable pointer itself), `L05916` (`+0x40`, the post-Tick removal flag) — or a call through the vtable (`L03082`) or into the shared removal helper (`L05392`); every remaining instruction operates on the local stack frame or the local iterator object. This extends `MAGIC-AREATICK-036`, which documents this write's own trigger condition but not its exclusivity, with that completeness: it is the only write, not merely a write. Whatever fields the opaque callees this function reaches — Tick's own vtable `+0x18` body, the removal chain (`R1118`, `R1119`, the unread `R1111`/`R1112`), and the deleting destructor `R1120`/`L05905` — themselves write on a member is outside this census, which is scoped to `R0641`'s own instruction stream only; `member+0x40` itself, the flag this function's own body only reads (`L05916`), is set by three separate Tick-callee end sites `MAGIC-AREAEND-041` already names (`L05449`, `L03056`, `L05358`), not by any instruction in this function.

**Confidence.** High — the complete, non-truncated function body (`evidence/functions.txt`, `R0641`-`L05858`, bounded by its own frame setup and return) was read and every instruction classified as a member-read, member-write, control-flow, local-frame access, or call; exhaustiveness follows directly from having read the whole body rather than a search over it.

**Original status.** ● active

**Evidence.** [EXP-0374](../experiments/EXP-0374-ticked-effect-population/)

### MAGIC-190

**`R0641`'s own body never reads or writes a member's `+0x44` anywhere, and gates nothing on it; a member whose `+0x44` is null — including one left null by `SAV-1010`'s load-time identity-map miss — is not filtered, skipped, or specially routed by this function's own instructions.** Cross-referenced against `MAGIC-189`'s exhaustive per-member read/write census, the only offsets `R0641` touches on the member pointer are `+0x00` (vtable), `+0x3c`, and `+0x40`; `+0x44` appears nowhere in its body. A member reaching this function's walk with `+0x44` null therefore receives the identical unconditional treatment `MAGIC-188` already establishes for every member — the two `+0x3c`-conditioned reads and the Tick dispatch neither test nor depend on `+0x44` in any way. This answers only what this function's own body does, per the question's own scope: whether Tick's own callee (vtable `+0x18`, the function `MAGIC-TARGETID-182` traces to `R0266`, whose two unguarded `+0x44` dereferences at `L05860`/`L05869` that claim already names as the fault sites) is ever actually invoked with a null `+0x44` present is the separate, still-open reachability question `MAGIC-TARGETID-182`/`MAGIC-TICKGATE-183` leave Unknown; this claim neither narrows nor extends that reachability, it only confirms nothing in `R0641`'s own body would intercept such a member before the dispatch those claims already show is unguarded.

**Confidence.** High that `R0641`'s own body contains no `+0x44` access, from the same complete-body read as `MAGIC-189`. The reachability of an actual null-`+0x44` member reaching this function at all remains the Unknown `MAGIC-TARGETID-182`/`SAV-1010` already carry, not narrowed here.

**Original status.** ● active

**Evidence.** [EXP-0374](../experiments/EXP-0374-ticked-effect-population/)

### MAGIC-197

**`R0634` (SpellTransport's own Tick) computes the identical shared-list container operand at both of its own registrar calls that `R0003` computes at both of its own — closing `MAGIC-187`'s own left-open Unknown — and SpellTransport's own vtable `+0x04` slot follows the identical scalar-deleting-destructor shape `MAGIC-187` confirmed only for PointEffect's, extending that confirmation to a second class.** Each of `R0634`'s own two calls to `R1113` (`L03083`, hand-off of `this+0x44`; `L03084`, hand-off of `this+0x48`) is immediately preceded by the pointer chain from the global `L00285` through the displacements `+0x14` and `+0x4` and an advance of 4 — identical, up to the carrying register, to the sequence at each of `R0003`'s own two calls (`L05395`, `L03053`) that `MAGIC-187` already reads as `*(*(X+0x14)+0x4)+4` with `X=*(L00285)`: same base literal, same two displacements, same order, only the carrying register free to differ. `SpellTransport`'s own constructor (`R0633`) installs vtable `L05917` (stored into the first field at `L05918`); that vtable's own `+0x04` slot is `L05919`, a distinct address from PointEffect's own `R1120` but the identical body shape `MAGIC-187` reads for it: a call of R1062 (the real destructor), a mask of the incoming argument with 0x1, a conditional call of R1121 (the same operator-delete target `MAGIC-187` and `TRIG-TAKEITEM-038` already name). Both facts read as complete instruction sequences against the lawful EN and RU images independently, confirmed byte-identical between them. Neither reads `AreaEffect`'s own `+0x04` slot, which was not read by `MAGIC-187`/EXP-0374 or by this experiment — no image-wide census of every ledger row for an AreaEffect vtable-slot address was run, so this is bounded to the two experiments that traced this call graph, not a claim that no row anywhere describes it.

**Confidence.** High for the two facts actually read, and for what they exclude: a transport-private or per-caster/per-session container would compute its own operand from a different base address or displacement chain than the shared list's own, and `MAGIC-187`'s own `callto:R1113` census (4 hits / 2 owners, image-wide) makes the four call sites read here — two in `R0634`, two in `R0003` — the complete population, not a sample; all four carry the identical base literal `L00285` and displacement chain `+0x14`/`+0x4`/`+4`, which excludes that alternative directly rather than by inference from the two functions sharing a registrar. Each read is a complete, EN/RU-cross-checked instruction sequence. This narrows `MAGIC-187`'s own "assumed, not shown, for any other member class" to specifically shown for a second class; it remains unshown for AreaEffect. It does not itself re-derive `MAGIC-187`'s own separate Medium (whether `R0193`'s own `this` equals `*(L00285)` at runtime) or read the unlink primitives `R1111`/`R1112`, both still Unknown as `MAGIC-187` already states.

**Original status.** ● active

**Evidence.** [EXP-0378](../experiments/EXP-0378-transport-scheduling/)

### MAGIC-202

**After a ROM1 archive LOAD resolves a `SpellEffect`-lineage record, the archive's own load arm inserts it into a container by the identical generic primitive, and the identical operand-computation formula, `MAGIC-187`/`MAGIC-197`'s own Tick/Apply-time registrar uses — not a separate post-archive pass, and not something left to the first Tick — though whether the load-time container and the Tick-time container are the identical runtime object is the same Medium gap `MAGIC-187` already carries, not re-derived here.** `R1109` (“the SpellEffect list” `SAV-DOC-053` already names, reached `R1122`→`R1109`) branches its own `IsLoading` arm at `L05920`/`L05921` (a call of R1123, then a jump on equal); its LOAD arm (`L05921`-`L05922`) clears the container first (`L05923`, a call of L05924 that itself calls L05925 on the same object), reads a `u32` element count (`L05926`, a call of R0686), then loops: `L05927` (a call of R1124) resolves one element through `CArchive::ReadObject` (`L05928`, `SAV-774`) with descriptor `L05929` pushed inside that wrapper, and `L05930` (a call of R0647 on the container, with the just-resolved pointer as argument) appends it — the same insert primitive `MAGIC-TICKGATE-183`/`MAGIC-187` already read as “a generic doubly-linked insert with no field access on the object being appended,” and one of only 3 owners EXP-0374's own image-wide `callto:R0647` census finds (`R1113`, the Tick/Apply registrar; `R1125`, a distinct class's own Serialize, reached only through a `.rdata`/`.data` vtable dword by this experiment's own narrower call-target/data-pointer scan, not traced further; and this function). `R1122` adds a literal `4` to its own incoming `this` before the call (`L05931`); the whole document loader `R0414` calls it from exactly two addresses, one on each side of `SAV-1030`'s own store/load split (`L05932`/`L05933`): `L05934`, inside the store arm (address below the split), and `L05935`, inside the load arm (address above it) — so a LOAD reaches this container exactly once, not twice, resolving what looked like a duplicate call site before the store/load boundary was checked directly. The load-arm call site is preceded by a lazy allocation of the manager pointer itself when absent (`L05936`-`L05937`: the field `Q+4` is compared with zero and a jump to L05937 on non-zero skips allocation when already present; otherwise a call of R0747 with the size 0x20, then constructor `L05938`, storing the result at `Q+4`) — first-touch construction, not a standing container guaranteed to already exist. The final operand computation at both the load-arm call site and `MAGIC-187`'s own Tick-time registrar site is the identical literal displacement chain: `D=frame−0x90` (the document loader's own cached `this`, `SAV-1030`'s own `S`) → `Q=*(D+0x14)` → `M=*(Q+4)` → operand `M+4` (the `+4` supplied by `R1122` itself) — matching `MAGIC-187`'s own `*(*(X+0x14)+0x4)+4` term for term, with `D` playing `X`'s own role. A separate pass over the same manager also exists and was read whole: `R1126` (31 instructions, dispatched through the manager's own vtable slot 0 per `SAV-TOKENLOAD-094`), computing `this+4` as its own container and walking it via the standard begin/GetNext idiom, calling each present element's own vtable `+0x24` slot — the Token-lineage post-load hook (`SAV-TOKENLOAD-093`) — but containing no call to `R0647` or any other write to the container's own head/tail fields anywhere in its body. That pass exists, and runs after the archive completes, but it does not register anything; whatever it walks was already inserted by the load arm above.

**Confidence.** High for the insertion mechanism itself — that a rebuilt `SpellEffect`-lineage record is appended into a container by the archive's own load arm, using the same primitive and, at the outer document-loader level, the identical operand-computation formula the Tick/Apply-time registrar uses, read as complete, EN/RU-cross-checked instruction sequences, with the load-arm/store-arm ambiguity resolved by locating `SAV-1030`'s own split address directly rather than assumed — and for the negative half: the one pass proven to run after the archive completes (`R1126`) performs no insertion, so “a parent/root registers children after the archive completes, as a separate pass” does not describe what happens; the pass that does exist dispatches, it does not register. Medium for treating the load-time container (`M`, reached from the document loader's own `this`, `D`) as the identical runtime object the Tick-time registrar reaches (`MAGIC-187`'s own `X`, `*(L00285)`) — the two chains compute the identical displacement formula off their own respective `this`, exactly as `MAGIC-187` already found for the Tick-time walker versus the Tick-time registrar, but this experiment did not itself trace any caller of `R0193` or `R0414` to compare `D` against `X` at runtime; this is a gap of the same shape as the one `MAGIC-187` names, over a different function's own `this`: `MAGIC-187` leaves open whether `R0193`'s `this` equals `*(L00285)`, while this row leaves open whether `R0414`'s does, and the rows narrowing the first do not bear on the second. Unknown: whether `R1125` (the third owner of `R0647`) inserts into this same container or a distinct one — this experiment's own scan found it reachable only through a vtable dword, with no direct call site, and did not read its body past confirming that shape.

**Original status.** ● active

**Evidence.** [EXP-0380](../experiments/EXP-0380-transport-rebind/EXP-0380.md), `evidence/container-serialize-R1109-body.txt`, `evidence/wrapper-R1122-body.txt`, `evidence/document-loader-body.txt`, `evidence/container-serialize-call-sites.txt`, `evidence/store-load-split.txt`, `evidence/manager-lazy-alloc-and-call-window.txt`, `evidence/postload-dispatch-R1126-body.txt`, `evidence/owner-R1125-reachability.txt`, `evidence/readobject-wrapper-R1124-body.txt`; MAGIC-187, MAGIC-197, MAGIC-TICKGATE-183, SAV-DOC-053, SAV-774, SAV-1030, SAV-TOKENLOAD-093, SAV-TOKENLOAD-094

### MAGIC-209

**The STORE arm of `R1109` (the same container-level `Serialize` `MAGIC-202` already reads on its LOAD side) writes a `u32` element count and then calls the identical generic write wrapper (`R1110`) once per element, with no class descriptor pushed by any instruction inside this STORE loop itself — unlike the LOAD arm's own per-element wrapper, which always pushes a literal expected-class descriptor before resolving; class identification on a STORE is instead a runtime lookup made inside the generic write call's own callee (`SAV-1047`'s own `GetRuntimeClass` dispatch and first-occurrence `WriteClass` push), never a literal this loop supplies itself.** `R1109`'s `IsLoading` branch (`L05939`/`L05920`) selects the arm not taken by `MAGIC-202`'s own LOAD reading; the STORE arm (`L05940`-`L05941`) reads the head node through `R1117` (the object as `this`; the same `return *(x+4)` primitive `MAGIC-187` already reads as `R0648`'s own "begin" wrapper's internal call — here it is the direct, sole call, not reached through that wrapper) into a local iterator (`L05942`-`L05943`), reads the element count through `R1127` (`return *(this+0xc)`, `L05944`) and writes it as a `u32` through `R0685` (`L05945`; the same width-4 helper `SAV-PLAYER-028` already names for `L08300`/`R0685`, confirmed here by its own 11-instruction body forwarding to `L05946`, an address no other claim row cites). It then loops while the iterator is non-null (`L05947`/`L05948`): `R1128` (`L05949`; the container as `this`, the container — the same `frame−0x18` slot `R1117`/`R1127` also read, reloaded at `L05950`, not the iterator — one stack argument, the iterator's own address `&frame−4` pushed at `L05951`) writes the current node's own `+0x0` field into `*position` — advancing past the current node — before returning the address of that same node's own `+0x8` payload slot (the identical advance-before-return idiom `MAGIC-187` already reads for a different pair of accessors, `R0648`/`R0650`, confirmed here from a second, distinct accessor function used only by this STORE-side loop); the caller dereferences that address (`L05952`) to get the element pointer and calls `R1110(ar, element)` (`L05953`-`L05954`). `R1110` is the thin wrapper `SAV-982` already names as forwarding the archive as `this` and one stack argument, to `R1129` without reading `R1129`'s own body (`SAV-1047` reads that body). No instruction in this loop pushes a literal class descriptor, in contrast to the LOAD arm's own per-element wrapper `R1124` (called at `L05927`), which always pushes a literal expected-class descriptor before calling `CArchive::ReadObject`; the visible call site `L05955`-`L05927` itself pushes only a local output-slot address and `ar`, not the literal, which `MAGIC-202` already places inside `R1124`'s own body (`SAV-1042` reads that body whole as a 10-instruction wrapper) — the corrected citation for where the LOAD-side literal actually lives.

**Confidence.** High for the STORE arm's own complete instruction sequence and every accessor it calls (`R1117`, `R1127`, `R1128`, `R0685`), each read whole, zero unresolved branches (`evidence/unresolved.json`) — this directly answers what the STORE arm emits per element: a generic write call carrying only the element pointer, with class identification deferred entirely to the callee (`SAV-1047`). EN/RU instruction-level agreement is asserted by `disasm.py`'s own `_assert_equal` check on the disassembled bodies, not read from `unresolved.json` (a per-function, EN-only unresolved-branch list); `evidence/program-digest.tsv` records the identical SHA-256 for both preserved roots, so on this pair of installs the EN/RU comparison cannot discriminate anything and adds no confirming power beyond the single image actually read. Medium/Unknown carried forward unchanged, not re-derived here: whether this container's own runtime identity at STORE time is the object `MAGIC-187`'s own Tick-time registrar reaches is the same gap `MAGIC-187`/`MAGIC-202` already name.

**Original status.** ● active

**Evidence.** [EXP-0381](../experiments/EXP-0381-transport-store/EXP-0381.md), `evidence/container-serialize-R1109-body.txt`, `evidence/store-head-R1117-body.txt`, `evidence/store-getcount-R1127-body.txt`, `evidence/store-getnext-R1128-body.txt`, `evidence/store-writeu32-R0685-body.txt`, `evidence/writeobject-wrapper-R1110-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-187, MAGIC-202, SAV-982, SAV-PLAYER-028, SAV-1042, SAV-1047

### MAGIC-213

**`PointEffect`'s own real destructor deletes exactly one field, `+0x48`; `AreaEffect`'s own real destructor deletes exactly one field, `+0x44`; each is the class's own "inner Effect" payload, never the sibling offset the other class uses; and `AreaEffect`'s own vtable `+0x04` slot, left assumed and not shown by `MAGIC-197`, is now read.** `PointEffect`'s own vtable (`L05356`) `+0x04` is `R1120` (`MAGIC-187`'s own subject); `AreaEffect`'s own vtable (`L05352`) `+0x04` is `L05956`, read here as raw data for the first time and confirmed the identical scalar-deleting-destructor shape `MAGIC-187`/`MAGIC-197` already read for the other two classes (a call of the real destructor, a mask of the argument with 1, a conditional call of R1121) — the third of three classes this area names, closing `formats/magic/application.md`'s own "assumed, not shown, for `AreaEffect`" sentence. `PointEffect`'s own real destructor (`L05905`, 40 instructions, `evidence/pointeffect-realdtor-L05905-body.txt`), read whole: after installing its own vtable, a single branch (`L05957`, a jump to `L05958` on equal) skips an entire dispatch-then-clear block when `this+0x48` is already null; when non-null, the pointee's own vtable `+0x04` slot is dispatched with argument `1` (the scalar-deleting-destructor convention `MAGIC-187` already names) and `this+0x48` is then written `0`; no other member offset is read or written anywhere in the body, whose own last call, at `L05959`, reaches `L05960` — not a tail call, since five more instructions follow it before the body ends at `L05961 ret`. `AreaEffect`'s own real destructor (`L05962`, 40 instructions, `evidence/areaeffect-realdtor-L05962-body.txt`) reproduces the identical shape at `this+0x44` instead — never reading `+0x48` anywhere in its own body — and calls the identical `L05960` at `L05963`, again its own last call and not a tail call. The address all three sibling real destructors this experiment reads call directly — `L05959`, `L05963`, and `MAGIC-214`'s own `L05964`, each a call read directly and none of them a tail call, not a raw call-site census — is `L05960` (9 instructions, `evidence/spelleffect-realdtor-L05960-body.txt`), attributed here to `SpellEffect`'s own real destructor by that shared-call-chain position rather than by any vtable or descriptor anchor; it installs no vtable of its own and does nothing but an unconditional forward to `R1130` — no member access at all. `R1130` (44 instructions, `evidence/furtherbase-realdtor-R1130-body.txt`), read one level further per this experiment's own stopping rule, installs a further base vtable (`L05965`) and touches only `+0x10` (released through `R0525`) and a sub-object at `+0x20` (torn down through `L05966`/`L05967`/`L05968`) — `+0x44`/`+0x48` appear nowhere in this body either. The base chain does not end here: `L05969` calls `L05970` with the same `this`, unread, so one further base destructor remains uncensused as a possible second deleter.

**Confidence.** High for all five instruction sequences (`PointEffect` real dtor, `AreaEffect` scalar and real dtor, `SpellEffect` real dtor, the further base dtor), each read whole, zero unresolved branches (`evidence/unresolved.json`), EN/RU-identical. This directly answers sub-question 2 of this experiment's own preregistration for the "inner Effect" field family: each class's own destructor deletes, not merely clears, its own inner-Effect field exclusively, and never reads the sibling offset the other class uses — read from the instruction stream, not assumed by symmetry between the two classes. Scope: this row is about the inner-Effect field only (`PointEffect`'s own `+0x48`, `AreaEffect`'s own `+0x44`, per `SAV-CLASSSER-174`/`SAV-CLASSSER-175`'s own typed archive-reference identification); `SpellTransport`'s own destructor treats a structurally different field pair and is `MAGIC-214`. This row does not determine whether the deleted inner-Effect pointer is ever independently shared elsewhere — `MAGIC-187`'s own image-wide registrar census (4 call sites, 2 owners, none of them this offset) already excludes it from the shared container specifically; a broader image-wide writer/reader census for these two offsets was not run by this experiment.

**Original status.** ● active

**Evidence.** [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/pointeffect-realdtor-L05905-body.txt`, `evidence/areaeffect-scalardtor-L05956-body.txt`, `evidence/areaeffect-realdtor-L05962-body.txt`, `evidence/spelleffect-realdtor-L05960-body.txt`, `evidence/furtherbase-realdtor-R1130-body.txt`, `evidence/vtable-slots.json`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-187, MAGIC-197, SAV-CLASSSER-174, SAV-CLASSSER-175, MAGIC-SIBLINGREF-184

### MAGIC-214

**`SpellTransport`'s own real destructor treats `+0x44` and `+0x48` as two independent owning pointers, each conditionally deleted through the pointee's own vtable `+0x04` slot and then cleared, in the identical single-field pattern `MAGIC-213` reads for `PointEffect`'s/`AreaEffect`'s own inner-Effect field, applied twice to two different fields in the one body — and, on the one delivery path this experiment traced, `SpellTransport::Tick`'s own hand-off/clear ordering, read here for instruction adjacency (`MAGIC-197`'s own subject is the registrar call sites and their operand computation, not this ordering), keeps this destructor's own dispatch from ever running against a field the shared container has already been given.** `R1062` (60 instructions, `evidence/spelltransport-realdtor-R1062-body.txt`), read whole, zero unresolved branches, EN/RU-identical: after installing vtable `L05917`, it runs the `+0x44` block first (`L05971`..`L05972`: single branch skips dispatch-and-clear entirely when already null; otherwise dispatch the pointee's own vtable `+0x04` with argument `1`, then `this+0x44 = 0`), then, independently, the `+0x48` block (`L05972`..`L05973`, the identical instruction pattern reproduced verbatim at a different offset and a different stack-local slot), before calling `L05960` (`MAGIC-213`'s own shared base) at `L05964`, this body's own last call and not a tail call — five more instructions follow it before the body ends at `L05974`, its return. Neither block's own five-instruction dispatch sequence (the argument 1 through the virtual call of slot `+4`) reads the other field: the `+0x44` block never references `+0x48` and vice versa — each field's own deletion is self-contained. This is the field pair `SAV-CASTCONT-1006`/`SAV-CLASSSER-175` already type as nested `SpellEffect`(`+0x44`, accepting `PointEffect` or `AreaEffect`)/`AreaEffect`(`+0x48`) archive references, and that `SpellTransport::Tick` (`R0634`, `MAGIC-197`'s own subject, re-read here) hands to the shared container's own registrar (`R1113`). Re-reading `R0634` for instruction adjacency: it hands off exactly one of `{+0x44, +0x48}` to the registrar (the primary if non-null, else the fallback), then clears both fields to null (`L05975`, `L05976`) before returning, always inside this one function body. The transient double reference ends when the handed-off field itself is cleared: three instructions after the call on the `+0x44` arm (`L03083` to `L05975`), four after on the `+0x48` arm (`L03084` to `L05976`) — the registrar wrapper `R1113` (11 instructions, re-read here) is a bare pass-through to the generic insert primitive `R0647`, itself already read by `MAGIC-187` as touching only the container's own head/tail/newNode fields, never the caller's own `SpellTransport` fields, so nothing between the hand-off and the clear can re-enter this object. By the time this destructor's own body could run afterward, the just-registered field is already null and its own dispatch for that field is skipped — no double-delete and no double-registration on this path. A second, narrower path this destructor's own reading covers that Tick does not: a `SpellTransport` destroyed before its own delivery countdown reaches zero still holds a live, unregistered child at `+0x44` (`MAGIC-215`'s own constructor reading: unconditional, no null check), and this destructor deletes it correctly — a genuine, exclusive single-owner path. What this row does not establish: a producer that leaves both `+0x44` and `+0x48` simultaneously non-null at the moment Tick's own countdown reaches zero. Tick's own body only hands off `+0x48` when `+0x44` IS null (an else-branch); if `+0x44` is non-null, `+0x48` is cleared without ever being read by the registrar or by this destructor's own dispatch for that field — an orphan-by-construction case if `+0x48` could ever be independently non-null at that moment. That fallback arm also carries no null guard of its own (`L05977`-`L03084`): when `+0x44` is null it hands `+0x48` to the registrar regardless of whether `+0x48` is itself null, so a transport with both fields null at expiry hands a null member to the shared container. The only writer of `+0x48` this experiment located is the constructor's own unconditional zero-write (`MAGIC-215`) and the LOAD arm's own `ReadObject` call (`SAV-1042`); no live, non-LOAD producer that sets `+0x48` non-zero while `+0x44` is also non-null was found within the functions this experiment read, and no image-wide census of every writer to a `SpellTransport` instance's own `+0x48` was run.

**Confidence.** High for the destructor's own complete instruction sequence and the field-independence within it, and for the hand-off/clear ordering in `SpellTransport::Tick` that prevents double-deletion/double-registration on the one traced delivery path — both functions read whole, the double-reference window measured directly from the instruction stream at three instructions after the call on the `+0x44` arm and four after on the `+0x48` arm. Medium for the overall answer to sub-question 2 for this class: it is complete for what the destructor does with each field independently, but whether either field can ever be simultaneously populated by a producer this experiment did not read is the named, unclosed Unknown above — the live alternative that keeps this row from High is a producer this experiment did not trace populating `+0x48` while `+0x44` is also non-null, which Tick's own else-branch shape would then silently orphan (cleared, neither deleted nor registered) — a case this evidence does not exclude.

**Original status.** ● active

**Evidence.** [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/spelltransport-realdtor-R1062-body.txt`, `evidence/spelltransport-tick-R0634-body.txt`, `evidence/spelltransport-ctor-R0633-body.txt`, `evidence/registrar-wrapper-R1113-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; MAGIC-213, MAGIC-197, MAGIC-TICKGATE-183, MAGIC-215, SAV-CASTCONT-1006, SAV-CLASSSER-175, SAV-1042

### MAGIC-215

**`SpellTransport`'s own constructor is the only construction site this experiment locates, and it never stores its own object argument at `frame+8` into more than one owning field: `+0x44` is written unconditionally from the argument, `+0x48` is always zero-written; the closest the traced graph comes to a doubly-referenced state is transient and confined to three instructions after `SpellTransport::Tick`'s own hand-off call on the `+0x44` arm and four after on the `+0x48` arm, ending when the handed-off field itself is cleared, both inside the one function body `MAGIC-214` already reads.** `R0633` (41 instructions, `evidence/spelltransport-ctor-R0633-body.txt`), read whole, matches `MAGIC-TICKGATE-183`'s own citation of this address as the class's single construction site (`L03052`), independently reproduced by this experiment's own raw call-site census (`evidence/call-site-census.json`: `spelltransport-ctor-R0633` target `R0633`, sites `[L03052]`, count 1 — one of three already-published counts this experiment's own method cross-checks before trusting it for a fresh target, per this experiment's own stopping rule). Both field writes below happen after the constructor's own unread base-class call returns (`L05978 call L05979`, not read by this experiment); its second argument, `frame+0xc`, is passed to that unread call and is not itself established as an object. Field order: install vtable `L05917` (`L05918`), then `this+0x44 = frame+8` (the argument, `L05980`, an unconditional store with no null test of the argument anywhere in the body), then `this+0x48 = 0` (`L05981`, also unconditional, a literal, not derived from any argument) — three instructions from the vtable install to the `+0x44` store, two more to the `+0x48` store, in that order — before computing the delivery countdown at `+0x4c` from `frame+0xc`, `frame+0x10` and `*(*(this+0x44)+0x10)` — the child pointer just stored at `L05980`, dereferenced through the field `+0x44` at `L05982` with no null test anywhere in between. No instruction in this constructor's own body writes a non-null value to `+0x48`. Cross-referenced against `MAGIC-214`'s own destructor reading and this experiment's own re-read of `SpellTransport::Tick` (`R0634`, `MAGIC-197`'s own subject) for instruction adjacency: the only place in the traced call graph where `+0x44` and `+0x48` could name objects simultaneously owned by the transport AND already handed to the shared container is inside `Tick`'s own body, between the registrar call (`L03083` or `L03084`) and the unconditional clear of both fields (`L05975`/`L05976`) — always inside the one function, never re-entrant, since the registrar wrapper `R1113` is a bare pass-through with no access to the caller's own fields (`MAGIC-214`). So the transiently-doubly-referenced state sub-question 1 of this experiment's own preregistration asks after does occur, in exactly this one narrow window, and collapses unconditionally before `Tick` returns.

**Confidence.** High that the constructor is the sole producer of `+0x44`/`+0x48` this experiment locates and that its own two writes are as described — a complete instruction-level read with a reproduced single-call-site census. High that no producer this experiment traced (constructor, `Tick`, and the LOAD arm `SAV-1042` already reads) stores one pointer into two owning positions that survive past the producing function's own return. Medium for the sub-question's own full wording, "any code path in the image": this row answers it for the graph this experiment's own preregistration named, not for an image-wide census of every writer to these two offsets; a producer entirely outside this traced graph is not excluded.

**Original status.** ● active

**Evidence.** [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/spelltransport-ctor-R0633-body.txt`, `evidence/spelltransport-tick-R0634-body.txt`, `evidence/registrar-wrapper-R1113-body.txt`, `evidence/call-site-census.json`, `evidence/unresolved.json`; MAGIC-TICKGATE-183, MAGIC-197, MAGIC-187, MAGIC-214

### MAGIC-216

**The two primitives `MAGIC-187` left unread inside the removal path do exactly what that row inferred from their call shape: `R1111` locates a node by its own payload pointer without touching the payload, and `R1112` unlinks that node from the doubly-linked list and hands the node — not the payload — to a container method (`L05983`, unread; graded Medium below as a node free by position), also without touching the payload. Removal is: locate the node holding that payload; unlink and free that node only if one is found; then, independently, delete the payload whenever the payload pointer is non-null (`R1118`'s own dispatch, already published). The two steps carry different guards, and a payload not in the container is deleted without ever being unlinked. A second caller of the removal wrapper exists, a bulk remove-all routine, raising this experiment's own direct-call census of the wrapper's caller population to at least 2, beyond `MAGIC-187`'s single-caller reading.** `R1111` (find-node-by-payload, 36 instructions, `evidence/removal-findnode-R1111-body.txt`), read whole: walks the list from a given node's own `+0x0` (next) or, when no start node is supplied, the container's own `+0x4` (head), and at each candidate compares node `+0x8` (the payload slot `MAGIC-187` already names) against the caller's own supplied payload pointer through an unread comparator (`L05984`, out of scope past this confirmation, per this experiment's own stopping rule) — returns the matching node pointer or `0`. No instruction anywhere in the body writes through the payload pointer; the only writes are to the local walk state. `R1112` (unlink-node, 41 instructions, `evidence/removal-unlinknode-R1112-body.txt`), read whole: a standard doubly-linked unlink — rewrites the container's own head (`+0x4`) to the node's own next when the node given is the current head, otherwise rewrites the node's own predecessor's own next; symmetric for the container's own tail (`+0x8`) against the node's own prev, the field this experiment newly locates at node `+0x4` (`MAGIC-187` names only `+0x0`/`+0x8`, next and payload). It then calls `L05983` once (`L05985`), handing it the node pointer itself, not the payload; `L05983`'s own body is unread past confirming it is called at the position a node free would occupy (graded Medium below). No instruction anywhere in this body reads or writes through the payload pointer at any offset. `R1118`'s own already-published body (`MAGIC-187`'s subject, re-read here at `evidence/removal-wrapper-R1118-body.txt`) shows `L05986`/`L05987` (the frame local −4 compared with zero, jumping to L05988 on equal) guarding the payload's own vtable `+0x04` dispatch, argument `1`, on the payload pointer itself being non-null — a different condition from the unlink dispatcher's own guard at `L05989`/`L05990` inside `R1119` (`evidence/removal-unlink-dispatch-R1119-body.txt`, the same shape `MAGIC-187` already describes, re-read here for its own instructions). This confirms the two-different-guards shape stated above from each primitive's own instructions, with the delete step being the payload's own scalar-deleting destructor already read by `MAGIC-187`/`MAGIC-213`/`MAGIC-214`. `MAGIC-187`'s own inference — "that this pair unlinks the node is inference from the call shape... not a read fact" — is now a read fact for the unlink itself. A raw `E8 rel32` call-site census of `R1118`, keyed to that one target address image-wide (`evidence/call-site-census.json`), finds 2 call sites in 2 owner functions, where `MAGIC-187` named only `L05392` (the walker's own same-pass reap): `L05392` (confirmed, matching `MAGIC-187`) and `L05991`, inside a second function, `R1131` (20 instructions, `evidence/removal-bulk-clear-R1131-body.txt`), read whole — a bulk remove-all: it calls `L05992` (a head-member accessor, not read past this call site) to fetch a node, and while non-null, calls `R1118` on it, then re-fetches the next remaining node, looping until the container reports none — clearing the entire shared container one member at a time through the identical removal-plus-delete wrapper, rather than a separate bulk-free path. An indirect or vtable-dispatched caller is not excluded by this direct-call instrument, so this raises the wrapper's own known caller population to at least 2, not a complete count.

**Confidence.** High for both primitive bodies (complete, zero unresolved, EN/RU-identical) and for what they read, which directly answers sub-question 3 of this experiment's own preregistration: removal is locate-then-conditionally-unlink, then, independently, delete-when-non-null — two steps under different guards, read from `R1119`'s and `R1118`'s own guard instructions; a payload not in the container is deleted without being unlinked. High for the bulk-remove-all caller and its own reuse of the identical wrapper, read as a complete body and a reproduced direct-call-site census. This closes `formats/magic/application.md`'s own "neither primitive's body was read" sentence. Medium for `L05983` (the node's own free, called on the node pointer at the position a free would occupy, not read past that confirmation), the comparator `L05984`, and `L05992` (the head-member accessor `R1131` calls, not read past confirming it returns that member) — none of these three callees were read past confirming they are called at the position their names imply. Medium also for the wrapper's own caller population: the census is a raw `E8 rel32` direct-call scan, which cannot see an indirect or vtable-dispatched caller — this experiment's own `call-site-census.json` reports zero direct sites for `pointeffect-scalardtor-R1120`, an address `MAGIC-187`'s own `EnumRefs` census found reached through a `.rdata` vtable slot, the identical blind spot applied to `R1118`; 2 is a lower bound, not a complete population.

**Original status.** ● active

**Evidence.** [EXP-0382](../experiments/EXP-0382-shared-effect-owner/EXP-0382.md), `evidence/removal-findnode-R1111-body.txt`, `evidence/removal-unlinknode-R1112-body.txt`, `evidence/removal-bulk-clear-R1131-body.txt`, `evidence/removal-wrapper-R1118-body.txt`, `evidence/removal-unlink-dispatch-R1119-body.txt`, `evidence/call-site-census.json`, `evidence/unresolved.json`; MAGIC-187, MAGIC-189

### MAGIC-217

**`PointEffect`'s and `AreaEffect`'s own per-tick attribution tails dereference the recorded caster's own pointee exactly three times each, in the identical order: first, through the "definition" test callee `R0416`, with no liveness check on the caster performed before that call — that dereference is still gated by `this->+0x3c != 0` and by `+0x41` (`PointEffect`) or `vt+0x20` (`AreaEffect`), only not by the caster's own `+0x14`; the other two dereferences, in both tails, are themselves the caster's own `+0x14==0` liveness check — no further field of the caster is ever dereferenced in either tail, and every later use of the caster's own pointer value (storing it into another field, or passing it to the final notify call) is preceded by at least one of those two checks.** `R0266` (`PointEffect`, 116 instructions, whole, zero unresolved, `evidence/pointeffect-tick-attribution-R0266-body.txt`, `MAGIC-TARGETID-182`'s own subject for `+0x44`, re-read here for `+0x3c`) and `R0267` (`AreaEffect`, 131 instructions, whole, zero unresolved, `evidence/areaeffect-tick-attribution-R0267-body.txt`, not previously committed whole by any claim) both reach `this->+0x3c` only after an unrelated block that dereferences a payload object's own vtable `+0x3c` slot (a virtual call, `L05993`/`L05868` and `L05413` — the unrelated `Effect`-class vtable slot at the same numeric offset, not this field; `PointEffect`'s own `L05868` lies on the `L05865` arm taken when the earlier, unrelated `+0x2c` call returns zero, which never reaches the caster-attribution block below) and the caster-attribution block proper: a comparison of `this+0x3c` with zero (`L05994`, `L05995`) gates entry; the load of `this+0x3c` and the call of R0416 (`L05996`-`L05997`, `L05998`-`L05999`) is the first dereference of the caster's own pointee — `R0416` (12 instructions, whole, `evidence/caster-definition-test-R0416-body.txt`): a comparison of `+0x3c` with zero producing a boolean — tests the caster's *own* `+0x3c` field (a different field on the caster's own class, the "definition" `SAV-1011` already names from this same callee's usage inside `Spell::Apply`), not `SpellEffect`'s own `+0x3c`. Only after this call returns does either tail read the caster's own `+0x14`: once to decide whether the credit write proceeds (a comparison of `caster+0x14` with zero, `L06000` / `L06001`) and once more, independently, immediately before the final notify call (the same comparison again, `L06002` / `L06003`). `PointEffect`'s own tail additionally re-tests `this->+0x3c` for null a second time (`L06004`) between those two liveness checks; `AreaEffect`'s own tail does not repeat that null test at the equivalent point, going straight from the credit write to the final liveness check — one asymmetry. The zero-`+0x14` arms also differ, as `MAGIC-ATTRGATE-118` publishes: `AreaEffect` clears `victim+0x40` (`L06005`-`L06006`), and `PointEffect` leaves it unchanged (`L06007`). Neither affects the guard order. Both tails reach the identical external call at the end, a call of R0226 on the object at `[L00004]` (a global singleton, the same address in both functions) and the caster/target pair pushed as arguments — `MAGIC-ATTRGATE-118`'s own delayed-kill-credit branch, confirmed here to be the identical mechanism for both classes, not assumed from the shared row text. Grep-verified: every `0x3c`-displacement instruction in both bodies is accounted for above, split between three classes of dereference (payload-vtable-slot virtual calls, unrelated; `this->+0x3c` field reads, which only ever retrieve the caster pointer value; and the one caster-pointee dereference inside `R0416`). `MAGIC-ATTRGATE-118` already documents the outcome of these same two gates ("a now-null caster definition clears victim+0x40... only a surviving owner admits the recorded-caster... writes") for both classes; this row adds the guard order the outcome-level description does not state — the "definition" test is not itself gated by the liveness check that gates everything downstream of it. Neither tail writes `this->+0x3c` anywhere in its own body (`SAV-1056`). This is not the only guard in the citation graph: `MAGIC-188`/`MAGIC-189` already read the shared-list Tick driver (`R0641`, cited, not re-read here) as performing the identical `caster+0x14==0` test on a member's own `+0x3c` and clearing it to null immediately before dispatching that member's own Tick call, every simulation tick, for every member the walk reaches — so the unguarded `R0416` dereference this row finds is reached with a stale `+0x3c` only if the recorded caster becomes invalid strictly between the driver's own per-member check and that same member's own Tick call completing, on the same tick; a caster already dead going into a given tick has its recording cleared by the driver before either tail runs at all. Whether a caster can die inside that window — for instance as a side effect of an earlier member's own Tick call in the same list walk, if that earlier call's own damage application can reduce a Unit's own health to zero synchronously — was not traced by this experiment: no Tick callee's own body was read for this question, and this row does not claim either the possibility or its absence.

**Confidence.** High for the guard order itself — a complete, zero-unresolved instruction-level read of both tails, cross-checked by an explicit accounting of every `0x3c`-displacement instruction in each. This directly answers what each reader does with a non-null `+0x3c`: one specific dereference (feeding `R0416`) is not preceded by this codebase's own `pointee+0x14` liveness convention; every other one is. Medium for the consequence this experiment's own sub-question 3 ultimately asks after — what happens when the named object no longer exists — since whether a freed `Unit`'s own memory is ever reused, unmapped, or merely flagged dead by this same `+0x14` convention (never independently confirmed as a general law by this experiment, a bounded observation only) is an object-lifetime question outside a static instruction read; no in-flight-population census was run, the same Unknown `SAV-1011` already leaves open, and the driver-level gate cited above narrows but does not close the reachability question. The `R0416` dereference's own reachability against a genuinely dangling caster pointer is therefore stated as a read fact about guard order, not as a demonstrated fault.

**Original status.** ● active

**Evidence.** [EXP-0383](../experiments/EXP-0383-effect-token-reference/EXP-0383.md), `evidence/pointeffect-tick-attribution-R0266-body.txt`, `evidence/areaeffect-tick-attribution-R0267-body.txt`, `evidence/caster-definition-test-R0416-body.txt`, `evidence/unresolved.json`, `evidence/program-digest.tsv`; SAV-1011, MAGIC-ATTRGATE-118, MAGIC-TARGETID-182, MAGIC-188, MAGIC-189, MAGIC-AREATICK-036, SAV-1056

## The caster's own `+0x14`: what "liveness" names on the write side

`EXP-0387` was allocated ids `221`..`224` of `claims/magic.md` (4 ids) and spent
`MAGIC-221`, 1 id. **`MAGIC-222`..`MAGIC-224` (3 ids) are returned unused**, never
reissued: the brief's Part 3 questions (how a picked spell becomes a cast order,
and whether a creature's cast spends mana/respects cooldown/range/fails
differently than a human's) resolved to `R0209`'s own six-target dispatch
table (`R0018` and `R0209`'s own kind-9 arms), read into one row, and
a parity question graded Medium/Unknown against the same shared, actor-type-agnostic
apply routine (`Spell::Cast`, already read whole by this page's own "Cast sequence"
section) — a single row was enough to carry both findings, and the mana/`Defensive`
filter Part 3's third question asks about is `MAGIC-AI-012`'s own pre-existing text,
not a new one. A returned range stays retired; the next free `magic.md` id is
therefore past it, at `MAGIC-225`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-219 | A writer this experiment locates does zero an already-constructed object's own `+0x14` at runtime, inside a removal/detach routine (`R1132`) distinct from every construction path `SAV-1058` reads and from every function ... | High / Medium / Unknown | ● active | [EXP-0384](../experiments/EXP-0384-effect-token-field14/EXP-0384.md), `evidence/call-site-census.json`, `evidence/actor-owner-status.tsv`, `evidence/population.txt`, `evidence/token-removal-R1132-body.txt`, `evidence/token-ownerctor-R1135-body.txt`, `evidence/spellapply-ownerwrite-L06023-window.txt`, `evidence/item-site-L06012-window.txt`, `evidence/item-site-L06013-window.txt`, `evidence/item-site-L06014-window.txt`, `evidence/building-site-L06015-window.txt`, `evidence/unit-ctor1-R0901-body.txt`, `evidence/unit-ctor2-R0501-body.txt`, `evidence/unit-ctor4-L06020-body.txt`, `evidence/unit-ctor5-L06021-body.txt`; SAV-1058, MAGIC-188, MAGIC-189, MAGIC-217, MAGIC-AREATICK-036, UNIT-OWNER-009, UNIT-CLASS-023, UNIT-T9CTOR-110, ITEM-CLASS-001, HERO-DEATH-026 |
| MAGIC-236 | Continuation of `MAGIC-219`: the `Unit` constructor census, the `Spell::Apply` owner store and the saved-corpus owner cross-check. | High | ● active | [EXP-0384](../experiments/EXP-0384-effect-token-field14/EXP-0384.md), `evidence/call-site-census.json`, `evidence/actor-owner-status.tsv`, `evidence/population.txt`, `evidence/token-removal-R1132-body.txt`, `evidence/token-ownerctor-R1135-body.txt`, `evidence/spellapply-ownerwrite-L06023-window.txt`, `evidence/item-site-L06012-window.txt`, `evidence/item-site-L06013-window.txt`, `evidence/item-site-L06014-window.txt`, `evidence/building-site-L06015-window.txt`, `evidence/unit-ctor1-R0901-body.txt`, `evidence/unit-ctor2-R0501-body.txt`, `evidence/unit-ctor4-L06020-body.txt`, `evidence/unit-ctor5-L06021-body.txt`; SAV-1058, MAGIC-188, MAGIC-189, MAGIC-217, MAGIC-AREATICK-036, UNIT-OWNER-009, UNIT-CLASS-023, UNIT-T9CTOR-110, ITEM-CLASS-001, HERO-DEATH-026 |

### MAGIC-219

**A writer this experiment locates does zero an already-constructed object's own `+0x14` at runtime, inside a removal/detach routine (`R1132`) distinct from every construction path `SAV-1058` reads and from every function `HERO-DEATH-026`'s own combat-death teardown traces — so the write-side evidence does not discriminate between `UNIT-OWNER-009`'s own "owning Player" meaning and a liveness-adjacent meaning for `+0x14`; this row grades that discrimination Unknown, not Medium. The caster `MAGIC-188`/`MAGIC-217`/`MAGIC-AREATICK-036` read as `caster+0x14==0` is, per `SAV-1011`'s own naming, a `SpellEffect+0x3c` object generally described as `Unit`/`Human`, but `UNIT-T9CTOR-110` shows at least one cast path (the building-associated type9 arm) constructs a distinct `VirtualCaster` object instead and copies its own `+0x14` from the source `Building`'s own `+0x14` at construction — so the caster's own concrete class is Unknown in the general case, not established as `Unit`/`Human` by either cited row.** `R1132` (157 instructions, whole, zero unresolved, `evidence/token-removal-R1132-body.txt`, 2 direct callers in the raw `E8` census, `L06008`/`L06009`, neither read by this experiment) walks a list off its own single argument's `+0x20` field (guarded by that argument's own `+0x50`/`+0x3d` fields, an iterator built through `R0027`/`R0028`/`R0029`) and, for each list member reached, unregisters it from two global tables (`R0865` on `[L04624]`, `R1133` on `[L00240]+4`), writes `[member+0x14]=0` (`L06010`) and `[member+0x94]=0xffce` (`L06011`, a `word`, `s16 -50`) before advancing. `HERO-DEATH-026` already publishes `actor+0x94` as the same signed-16-bit health field this teardown's own decay machine drives below `-10`/`-20`/`-40`/`-600`, so a member walked here is shaped like an actor and this write is consistent with force-setting its health to a below-`-10` sentinel as part of removal — inferred from the matching field width and an already-published field identity elsewhere, not proven: `R1132`'s own body installs no vtable and reads no class descriptor for the walked member, so its concrete class (`Unit`, `Human`, or another list element sharing this field shape) is not fixed by this function alone. `HERO-DEATH-026`'s own three-stage teardown (`claims/hero.md:51`) is a traced call chain — `R0864`'s tick, `R0208`'s teardown, `R0867`'s decay, `R0868`'s id release — and `R1132` is not one of its own named callees; its silence on `+0x14` is silence within that traced chain, not a claim that chain is the complete set of functions touching a dying or removed actor's own fields, and this row does not read it as such. `SAV-1058`'s own call-site census (`evidence/call-site-census.json`) finds `R0923`/`R1009`/`R1134`/`R1135` — `Token`'s own four constructors — called from sites across the image outside the `SpellEffect`/`Effect` lineages' own chains. Three of `R0923`'s own 13 sites (`L06012`, `L06013`, `L06014`) and one of `R1009`'s own 7 (`L06015`) were read in an earlier pass of this row only by address-range proximity as "consistent with... though not itself proof of" a `Unit`/`Human` writer; a short linear window at each (`evidence/item-site-L06012-window.txt`, `-L06013-`, `-L06014-`, `evidence/building-site-L06015-window.txt`) shows each instead installs `Item`'s own vtable (`L04580`, at `L06016`/`L06017`/`L06018`) or `Building`'s own (`R0492`, at `L06019`) — already-published class identities: `ITEM-CLASS-001`'s own `Item::Item` at `R0884` matches this row's own `L06012` site to the byte, and names the same `Token::Token`/vtable/field sequence. `Item` and `Building` are therefore two further `Token`-lineage classes (`ITEM-CLASS-001` already states this for `Item`), each reached through the shared zero-constructor and writing no further `+0x14` of their own in the window read — evidence that the shared construction-time zero is not scoped to the six classes `SAV-1058` names, but not evidence about `Unit`/`Human` specifically; this row withdraws the earlier "consistent with... `Unit`/`Human`" reading of these four sites. The `Unit` constructor census, the `Spell::Apply` owner store and the corpus cross-check continue in `MAGIC-236`.

**Confidence.** High for the writer census within the population read: `SAV-1058`'s own eighteen fresh whole-body reads plus `EXP-0382`'s/`SAV-1011`'s own already-published bodies account for every construction/destruction/post-load-hook instruction across the six `SpellEffect`/`Effect`-lineage classes; the four `Unit` constructors, the removal routine and the `Spell::Apply` store are each read whole or as a bounded window and zero-unresolved; the `Item`/`Building` reclassification is a direct vtable read, not an inference. Unknown for the owner-versus-liveness discrimination this row set out to make: a runtime writer that zeros a member's own `+0x14` does exist (`R1132`), so `+0x14` is not exclusively a construction-time default, and this experiment did not trace `R1132`'s own two callers, what removal condition reaches it, or whether that condition is death specifically or a broader despawn/detach — so the finding is equally consistent with "owner cleared on removal" and a liveness-adjacent flag, and rules out neither. Medium for the reconciliation attempted in an earlier pass of this row: `HERO-DEATH-026`'s own text is read in full for its own named field set, but it traces one specific call chain, not every function touching a dying or removed actor, so its silence on `+0x14` bounds only that chain. High for the corpus's own 100%-resolved finding, bounded as before to 70 admitted saves — it still shows no saved `Human` or `Unit` carries a zero owner-reference, which is compatible with either meaning of the constructed/removed zero, not itself discriminating.

**Original status.** ● active

**Evidence.** [EXP-0384](../experiments/EXP-0384-effect-token-field14/EXP-0384.md), `evidence/call-site-census.json`, `evidence/actor-owner-status.tsv`, `evidence/population.txt`, `evidence/token-removal-R1132-body.txt`, `evidence/token-ownerctor-R1135-body.txt`, `evidence/spellapply-ownerwrite-L06023-window.txt`, `evidence/item-site-L06012-window.txt`, `evidence/item-site-L06013-window.txt`, `evidence/item-site-L06014-window.txt`, `evidence/building-site-L06015-window.txt`, `evidence/unit-ctor1-R0901-body.txt`, `evidence/unit-ctor2-R0501-body.txt`, `evidence/unit-ctor4-L06020-body.txt`, `evidence/unit-ctor5-L06021-body.txt`; SAV-1058, MAGIC-188, MAGIC-189, MAGIC-217, MAGIC-AREATICK-036, UNIT-OWNER-009, UNIT-CLASS-023, UNIT-T9CTOR-110, ITEM-CLASS-001, HERO-DEATH-026

### MAGIC-236

**Continuation of `MAGIC-219`: the `Unit` constructor census, the `Spell::Apply` owner store and the saved-corpus owner cross-check.** It follows the `Item` and `Building` reclassification on `MAGIC-219`: `Unit`'s own construction is instead read directly: four of the six constructors `UNIT-CLASS-023` already names for vtable `L00001` (`R0901`, `R0501`, `L06020`, `L06021` — two further sites for that vtable, `L06022`/`R1136`, are not read here) each call one of `Token`'s own four constructors (`R0901`/`R0501` call the first zero-constructor `R0923`; `L06020` calls the second zero-constructor `R1009`; `L06021` calls the owner constructor `R1135`) and then install `L00001` (`L00494`/`L00496`/`L00500`/`L00502`), writing no separate `+0x14` of their own — direct, whole-body evidence, not an address-range inference, that `Unit`'s own `+0x14` is `Token`'s own inherited field, set at construction by the same shared layer `SAV-1058` reads for `SpellEffect`/`Effect`. `Human`'s own constructor is not separately read by this experiment (out of scope; `Human` and `Unit` share `UNIT-OWNER-009`'s own field naming, not independently re-derived here). `Spell::Apply` (`R0003`, `MAGIC-TICKGATE-183`'s own subject) stores a newly constructed object's own `+0x14` from a local at `L06023` (`evidence/spellapply-ownerwrite-L06023-window.txt`), after copying `+0x96`/`+0xa6`/`+0xbe` from a second object and registering the new object with `[L00240]` (the same global `R1132` unregisters a walked member from) — a further, previously uncited located writer, read only as a short window (`Spell::Apply` is not one of this experiment's own six classes and is already partially read elsewhere for a different field). `UNIT-OWNER-009` (`claims/unit.md:52`) already publishes the spawn/placement-time assignment writer for a live actor's own `+0x14` (manager lookup `R0425` at `L01836`, storing the resulting `Player` pointer at `actor+0x14` at `L06024`) — unaffected by this row's own findings. A corpus cross-check (`experiments/EXP-0384-effect-token-field14/corpus.sh`, `evidence/actor-owner-status.tsv`): across the same 70-save admitted population `SAV-1059` reads, every `Human`/`Unit` "owner"-kind reference — 773 `Human` + 1247 `Unit`, 2020 total — is `resolved`; zero are `zero` or `unresolved`.

**Confidence.** High for the clauses moved here, with the scope and bounds `MAGIC-219` states for them; the Unknown and Medium clauses stay on `MAGIC-219`.

**Original status.** ● active

**Evidence.** [EXP-0384](../experiments/EXP-0384-effect-token-field14/EXP-0384.md), `evidence/call-site-census.json`, `evidence/actor-owner-status.tsv`, `evidence/population.txt`, `evidence/token-removal-R1132-body.txt`, `evidence/token-ownerctor-R1135-body.txt`, `evidence/spellapply-ownerwrite-L06023-window.txt`, `evidence/item-site-L06012-window.txt`, `evidence/item-site-L06013-window.txt`, `evidence/item-site-L06014-window.txt`, `evidence/building-site-L06015-window.txt`, `evidence/unit-ctor1-R0901-body.txt`, `evidence/unit-ctor2-R0501-body.txt`, `evidence/unit-ctor4-L06020-body.txt`, `evidence/unit-ctor5-L06021-body.txt`; SAV-1058, MAGIC-188, MAGIC-189, MAGIC-217, MAGIC-AREATICK-036, UNIT-OWNER-009, UNIT-CLASS-023, UNIT-T9CTOR-110, ITEM-CLASS-001, HERO-DEATH-026

## Direct cast-admission construction fields

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-225 | The direct PointEffect admission caller supplies fields that the containing SpellTransport never copies into itself. | High / Medium | ✔ promoted | [EXP-0388](../experiments/EXP-0388-effect-construction/EXP-0388.md), `evidence/point-admission.txt`, `evidence/delivery-admission.txt`, `evidence/vectors.tsv`, `evidence/controlled-calls.tsv`, `evidence/call-boundaries.tsv`; SAV-1068, SAV-1069, MAGIC-ATTRGATE-118, MAGIC-PIC-026, MAGIC-215 |

### MAGIC-225

**The direct PointEffect admission caller supplies fields that the containing SpellTransport never copies into itself.** In `R0003`'s selected PointEffect branch, the constructor return receives `+0c=Spell+08` at `L06025`, `+41=(Spell+0a==0)` at `L06026` (MAGIC-ATTRGATE-118's existing polarity), and `+0e=2*(Spell+08)+9` at `L06027` (MAGIC-PIC-026's existing expression). Delivery 1 repeats the type write at `L06028` and appends the PointEffect at `L05395`. Delivery 2 constructs SpellTransport `R0633` at `L03052`, passing caster Position and the child; it appends the transport at `L03053`. After that constructor returns, the caller's only direct transport store is the existing `+4c` override, outside the selected Token/common fields. Thus no direct caller store fills transport `+0c` or either object's `+08`, nor copies the child's type/common fields into the transport. Its constructor stores remain `+0e/+40/+41=0/0/1` under the probe controls; the child retains its own fields. The registrar wrapper/insert directly manipulate the node and container, not payload fields; node allocation is an explicit unexpanded boundary. A delivery value other than 1 or 2 skips wrapping/registration in this suffix after child construction, correcting MAGIC-TICKGATE-183's no-construction wording. The chosen transport constructor stores the child only in `+44` and zeroes `+48`, as MAGIC-215 already establishes; the earlier `+44/+48` wording is also retracted for this call path.

**Confidence.** High for direct field sources and exclusions within the two bounded caller windows, with four constructor and eight admission runs producing 16 checked object states. Medium for native admission completeness; allocation/failure/exception paths, transitive external-helper effects, arbitrary aliases, AreaEffect/alternate construction and outer-Apply prehistory are outside the execution population. First native SAVE and all intervening writers are Unknown; explicit unrun predictions and reachable falsifiers are in the write-up.

**Original status.** ✔ promoted

**Evidence.** [EXP-0388](../experiments/EXP-0388-effect-construction/EXP-0388.md), `evidence/point-admission.txt`, `evidence/delivery-admission.txt`, `evidence/vectors.tsv`, `evidence/controlled-calls.tsv`, `evidence/call-boundaries.tsv`; SAV-1068, SAV-1069, MAGIC-ATTRGATE-118, MAGIC-PIC-026, MAGIC-215

## Temporary caster owner and group

Every quoted instruction is asserted byte for byte on both editions' `rom.exe`, which are identical
(162 rows, 0 mismatches). The instrument is capstone disassembly of address ranges plus direct `CALL rel32`
cross-references and displacement sweeps (`+0x70` over `L06029..L06030`, `+0x14` over `L02478..L06030`);
indirect calls are not followed.

`EXP-0443` was allocated ids `235`..`236` of `claims/magic.md` (2 ids) and spent both: `MAGIC-235`, and
`MAGIC-236`, the continuation of `MAGIC-219` that the card limit required when this ledger was rewritten as
cards. The next free `magic.md` id is `MAGIC-237`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-235 | A temporary caster built by R0914 has owner (+0x14) and group (+0x70) zero from its constructors, no call site of its two wrappers assigns either, and the Prismatic Spray selector reads both without a null test. | High / Medium / Unknown | ● active (amended) | [EXP-0443](../experiments/EXP-0443-footprint-slot/EXP-0443.md) |

### MAGIC-235

`R0914` builds the object that states `0x0d` and `0x0e` use as a temporary caster. It allocates `0x198`
bytes and calls `R0901` (`L06031`), one of the constructors `MAGIC-219` lists for the `Unit` vtable
`L00001`. It then zeroes `+4`, appends the object to the list at `this+0x2c` (`L06032`), sets the position
from its arguments through `R0289` (`L06033`), builds a `0x1c`-byte and a `0x14`-byte member, stores the
spell at `+0x64` (`L06034`), writes `+0x88 = 0x1e` (`L06035`) as well as the bytes `+0x134` and `+0x135` (`1`), and
returns. It makes no store at `+0x14` or `+0x70`. Two wrappers set the state: `R1137` writes `0x0d` at
`L06036` and `+0x5c`, and `R1138` writes `0x0e` at `L06037` as well as the bytes `+0x60` and `+0x61`.

The zero values come from the constructors. `R0901` calls the base constructor `R0923` (`L06038`),
which stores `[ECX+0x14] = 0` (`L06039`), and the initialiser `R0185` (`L00495`), which stores
`[EDX+0x70] = 0` (`L06040`) and the defaults `+0x49 = 1` and `+0x4a = 1` (`L00821`, `L00820`).
`UNIT-OWNER-009` names `actor+0x14` the owning Player.

The wrappers have ten direct call sites: `L06041`, `L06042` and `L06043` in `R0262`; `L06044`,
`L06045` and `L06046` in `R0156`; `L06047` and `L06048` in `R1139`; and `L06049` and
`L06050` in `R0458`, where they run for a cell record whose byte `+0x2c` is non-zero and not `0x1a`
(`MOVE-088`). None of the ten sites stores the returned pointer. After each call the next instruction is a jump to
the caller's exit or loop (`L06051`, `L06052`, `L06053`, `L06054`, and for `L06048` the loop test at
`L06055`), a return (`L06056`, `L06057`, `L06058`), or, in `R0458`, a path to `L06059`, which
reloads `frame−8` over the result. At the jumps to `L00892` and at the returns the return value still holds the new object,
so it is the return value of the enclosing function. The direct callers of `R0262` (`L06060`, `L06061`,
`L01870`), of `R0156` (`L01860`) and of `R1139` (`L06062`) were read one instruction past the
call, and none uses EAX there; `R0431` returns it (`L06063`), and its callers and everything after those
one-instruction reads were not read. Nothing at the ten sites stores `+0x14` or `+0x70`, or calls `R0152` on
the new object.

`R0152(group, actor)` is the group assignment. It removes an existing membership through `R1140`
when `actor+0x70` is non-zero, appends the actor to the group list (`L06064`), stores the group at `actor+0x70`
(`L06065`) and copies `actor+0x14` to `group+0x44` (`L06066`). It has 15 direct call sites in 11 routines: `R0064`, `R0823`, `R1141`, `R0061`,
`R0066`, `R0151`, `R0062`, `R0003`, `R0432`, `R1142` and
`R1143`; none is a wrapper, a wrapper caller or `R0914`. In `L06029..L06030` seven instructions store at displacement `0x70`: the zero above, the
store and the clear in `R0152` and `R1140`, and four others in classes with other vtables
(`L06067`, `L06068`, `L06069`, `L06070`).

The selector `R0109(session, caster, primary, out, cap)` calls `R0126(caster, primary)` at `L00878`.
That routine loads `caster+0x14` (`L06071`) and reads byte `+4` of it (`L06072`), then does the same for the
primary, and indexes a table with the two bytes (`L06073`..`L04127`), with no null test. The selector then passes `[caster+0x70]`
(`L06074`) to `R0110`, whose first read of it is a comparison of the field `+0xc` (`L06075`), again with no null
test. `MAGIC-SPRAY-135` gives the route from states `0x0d` and `0x0e` to this selector.

Of the three live readings, a group or owner inherited from the script's actor or the spell's origin caster
and a fixed neutral group are not found: no store or call in the population names either. The remaining reading
is that the caster has no owner and no group when it casts. A cast of spell 14 by such a caster reads address
4 at `L06072`. That is a static prediction and was not run.

**Confidence.** High for the builder's body, the zero stores of the constructors, the ten call sites and the
two unguarded reads, each an asserted instruction read whole. Medium for the absence of a later assignment:
the population is the ten call sites, the 15 direct call sites of `R0152`, the `+0x70` sweep over
`L06029..L06030` and the `+0x14` sweep over `L02478..L06030`, which excludes the wrapper module and
`L01836`; a store through a computed pointer, a `REP MOVS` copy of an object, an indirect call and code outside
the sweeps are not covered.

**Unknown.** Whether a wrapper owner is one of the script-instant creators named in `MAGIC-SPRAY-135`; only
the state values `0x0d` and `0x0e` link them. The runtime result of spell 14 cast by a temporary caster. Whether
a state routine of the new object assigns an owner on a later tick. Where the returned pointer goes beyond the
one-instruction reads at the callers named above.

**Amended.** The selector's unguarded reads are located by `MAGIC-241`: the owner read at `L06071` in `R0126`, called at selector entry, comes before the group read named above. The rest of the claim stands.

## Slot-drawn creature casts: aim, execution and retention

Evidence is a static read of `rom.exe` (one image on both lawful installs, sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`); no process was run. The instrument is capstone disassembly of address ranges, direct `CALL rel32` cross-references and displacement sweeps; indirect calls are not followed. Every instruction quoted below is asserted on both editions (`evidence/asserts-en.txt`, `evidence/asserts-ru.txt`: 170 rows, 0 mismatches). `ord` is the order block at `actor+0x158` and `mover` the block at `actor+0x154`. `MAGIC-221` already records the arm table, the Teleport overwrite and the `ord+0x60 = 1` store; the claims below add the aim of each arm, the failure behaviour and the retention test. The next free `magic.md` id is `MAGIC-241`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-237 | A slot-drawn spell is aimed by its `R0209` arm: 7 ids at the engage victim, 9 at the caster, 8 at the victim's cell, Acid Stream and Teleport at the cell next to the caster, 2 at nothing. | High / Medium | ✔ promoted | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md) |
| MAGIC-238 | A drawn Acid Stream or Teleport orders a kind-9 cast at the cell one step from the caster toward the victim; Teleport's nine-cell search is overwritten by that store. | High / Medium | ✔ promoted | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md) |
| MAGIC-239 | Kind-8/9 reach or facing failure selects approach rather than release; approach with a non-empty route retries, while route-search failure or the point obstacle branch can replace the pending kind. | High / Medium | ✔ promoted (amended) | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), [EXP-0497](../experiments/EXP-0497-out-of-range-cast/EXP-0497.md) |
| MAGIC-240 | `ord+0x60` is a retention flag read by the two cast install arms: zero runs the stop-and-reset at install, nonzero keeps the order armed; the eight stores of 1 found in the AI module are in creature cast selectors. | High / Medium | ✔ promoted | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md) |

### MAGIC-237

`R0209(this, actor, target, spellId)` receives `target` unchanged from `R0009`'s own second argument (`L06076` reads the stack slot `+0x20`, `L01739` calls R0209). After a non-null `Spellbook::Get` (`MAGIC-221`) its 28-entry table at `L05156` selects one of six arms. Each arm's aim, read from its instructions (`evidence/spell-arms.tsv` lists the arm beside each spell's name, Max Range and Defensive column):

- Arm `L05158`, ids 1, 11, 13, 14, 20, 27, 28 (Fire Arrow, Drain Life, Lightning, Prismatic Spray, Stone Curse, Curse, Slow): `R0018(actor, target, spell)` with the engage victim (`L05158`..`L06077`).
- Arm `L05159`, ids 5, 6, 10, 15, 16, 18, 22, 23, 24 (the Protections, Heal, Invisibility, Shield, Bless and Haste; `evidence/spell-arms.tsv` names each id): `R0018(actor, actor, spell)`. The arm passes the caster twice (`L06078`, `L06079`), so the caster is its own target (`L06080`, a call of R0018).
- Arm `L05160`, ids 2, 3, 7, 8, 12, 17, 19, 21 (Fire Ball, Wall of Fire, Freezing Cloud, Poison Cloud, Light, Darkness, Wall of Earth, Meteor Storm): a kind-9 order at the victim's packed cell. `R0019` (reading the 16-bit word at `+2`) reads it from the victim's position object (`L05160`..`L06081`), `ord+0x3c` takes it (`L06082`), and the shared tail stores `ord+0x30` and `ord+0x14 = spell+9` (`L06083`..`L06084`).
- Arm `L05161`, id 9 (Acid Stream), and arm `L05162`, id 26 (Teleport): a cell next to the caster (`MAGIC-238`).
- Arm `L05163`, ids 4 and 25 (Fire Sacrifice, Control Spirit): no order, only `ord+0x60 = 1`.

So Shield, Bless, Haste and Invisibility are aimed at the caster; Darkness at the victim's cell; Stone Curse at the victim; Acid Stream at the cell of `MAGIC-238`. None of the 28 arms is aimed at an ally. The ally search lives in `R0018` and runs only for a null target (`L06085`/`L06086`: a null test of the target and a jump to `L06087` on non-zero); it passes 1 or 0 by the spell's cached Defensive byte (`L06088`, a zero test). The caster arm never passes a null target. The victim arm's seven ids all have an empty Defensive column (`EXP-0067` catalogue, `Defensive=-1`), so a null target there would take the hostile search, not the ally one. Given the catalogue's Defensive column and a non-null target, the ally-selecting branch is not taken for any of the 28 slot-path ids. Light (id 12) and Teleport (id 26) carry a Defensive value; Light still goes to the victim's cell and Teleport to its own arm.

**Confidence.** High for the arm table (read off both roots, identical), each arm's aim and the single-push structure of the caster arm; the seven, nine, eight and two-id counts are the table's own. Medium for the caster being the whole target set of a Defensive spell: the arms were read whole, but `R0018`'s ally search was read only to the branch that selects it. The Defensive column values come from committed catalogue evidence, not from a fresh read of Data.bin.

**Unknown.** The cell arms (`L05160`, the tail at `L05161`) read `target+0x10` and call `R0051` with no null test, so a null target with a drawn cell spell would read low memory. Of the fourteen direct call sites of `R0009`, nine follow a zero test of the passed value (`L01989`, `L06089`, `L00586`, `L02072`, `L02073`, `L00842`, `L01990`, `L01708`, `L01991`), two supply a null in a branch the preceding zero test and jump on equal makes unreachable (`L06090`, `L00333`), `L00336` compares with a frame-held value, `L01992` passes the working target that `R0009` has just tested nonzero, and `L01988` (command state 3) passes `ord+0xc` with no test in its routine. Whether `ord+0xc` can be zero there was not traced.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), `evidence/asserts-en.txt`, `evidence/listing.txt`, `evidence/spell-arms.tsv`

### MAGIC-238

Both ids end at the shared tail `L05161`. It calls `R0051(plane, caster, victim)` (`L05167`), which classifies the direction from the caster's position to the victim's into one of eight and returns it in the top three bits of the returned byte (`L06091` clears bit 0 and `L06092` shifts left by 4; `L06093` shifts right by 5 and leaves 0..7). The tail adds the caster's cell bytes (`R0299` x, `R0300` y) to the step bytes `plane+0x58eb0+dir` and `plane+0x58eb8+dir` (`L06094`, `L06095`), stores kind 9 (`L06096`) and the word `(y << 8) + x` at `ord+0x3c` (`L05168`), then `ord+0x30 = spell` (`L06097`), `ord+0x14 = spell+9` (`L06084`) and `ord+0x60 = 1` (`L06098`). The plane initialiser `R0116` stores the step tables at `L00472`..`L00473`: dx = 0, 1, 1, 1, 0, -1, -1, -1 and dy = -1, -1, 0, 1, 1, 1, 0, -1. Every step is one cell, so the aimed cell is adjacent to the caster; when the victim stands one cell away on one of the eight lines it is the victim's own cell. The reach byte is the Max Range column: 3 for Acid Stream, 1 for Teleport.

Teleport's arm (`L05162`) first reads the victim's and the caster's packed cells and calls `R0167` (`L06099`). When the result is above 2 it scans the nine cells around the victim, (-1..+1, -1..+1), with `R0144(caster, cell)` (`L06100`) and, on the first cell that passes, stores kind 9, that cell, the spell and the reach (`L06101`..`L06102`). Every path then reaches `L05164`, falls into the tail and executes the store at `L05168`, which replaces `ord+0x3c`. The scan therefore cannot change the final order's cell. The arm contains no call to `R0179` (`MAGIC-221`).

**Confidence.** High for the tail, the overwrite and the step tables (complete-body read, instruction assertions on both roots, `R0051` read whole). Medium for the eight-way meaning of the direction byte: its sector boundaries were read as branches on the signed deltas and not tabulated against positions.

**Unknown.** What `R0167` and `R0144` return was not read; it does not affect the final cell. What the Teleport application does with the final cell is `MAGIC-TELEPORT-174`'s territory. The tail adds one step to the caster's packed cell with no clamp or passability test (`L06103`..`L05168`); edge wrap and the effect of an out-of-map cell were not read.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), `evidence/asserts-en.txt`, `evidence/listing.txt`, `evidence/spell-arms.tsv`

### MAGIC-239

The order machine `R0016` selects the kind arm only while order+9
is zero. Kind 8 at `L00108` bypasses distance and facing for self-target;
otherwise `R0041` gates installation at `L06105`. A false result
calls actor approach `R0042` at `L06107`. Kind 9 at `L05042`
uses `R0086`; a false result calls point approach `R0087`
at `L06112`. Their exact metrics are MAGIC-REACH-179 and
MAGIC-REACH-180. Forty-nine original-instruction controls in EXP-0497
stop at approach or install and corroborate these branches without a
service stub.

`R0042` passes the target and order reach to `R0043`.
Non-centred caster position calls the step-follow helper. At a centre,
distance within reach selects a direction/turn request and returns;
distance above reach selects path work. The path stores mover+0x7c
from the target pointer argument at `L13495/L01899`; reach is the
separate byte compared at `L00580`. AI-373 covers the subsequent
search, list and step preparation.

`R0087` passes cell and reach to `R0178`. Non-centred
caster position calls the step-follow helper. At a centre, whole-cell
distance within nonzero reach requests direction/turn and returns before
path search; distance above reach reaches path work. Both wrappers store
progress 3 and order+0x15 = 0 if the position centre predicate returns
zero, and end with actor action 1. Progress arm 3 can clear progress at
the next centre, allowing the kind arm to be evaluated again.

Complete point helper `R0178` also contains an obstacle branch.
After a centred out-of-range case reaches dynamic-route refresh, a blocked
candidate footprint and a nonzero external `R0402` verdict reach
a direction comparison. Wrong facing writes pending kind 10 at
`L13340`; matching facing writes pending kind 0 at `L13496`.
These local stores do not clear the parent cast state or queued Spell.

Route-search failure is a second path that replaces the pending kind. The
point helper stores mover+0x98 = 1 at `L01720` when `actor+0x168` is zero
after its full search `R0053`; the actor helper stores it at
`L00161` when the static path list is empty (AI-373). AI-373 also names
the near search `R0055`, an external callee, as a setter. The
executor epilogue is the flag's sole consumer (AI-ROUTE-045). For an
`actor+0x50` other than 1, 0xa and 0x17 it stores ord+0x08 = 0 at
`L00114` and calls reacquisition `R0004`, which writes kind 6, 0xb
or 0 (AI-350). The manual setters write parent state 0xd/0xe at
`L13497/L13498`, so an admitted manual cast whose route search fails
loses its pending cast kind in the same executor pass. The four complete
approach/wrapper bodies are the bounded store population; external callees
are not covered by that negative. MAGIC-253 establishes the ordinary manual
parent reissue and progress preservation.

The body of `R0268` has no range/facing test and its mage mana
refusal remains MAGIC-CADENCE-127's retained retry contract. This does
not close additional checks in its callees or the unit-tick gates.

**Confidence.** High for the selected gate branches, four complete
approach/wrapper bodies' direct stores, the conditional obstacle writes and
the two route-failure flag stores. The flag's consumer is the published
High AI-ROUTE-045 and AI-350 contract, not re-read here. Raw disassembly
and finite original-instruction execution exclude immediate install at the
selected distant non-self gates. Medium for approach, re-evaluation and
eventual movement as gameplay: path services, turn/step helpers and native
scheduling are not executed. The obstacle verdict is an untraced call, not
a claim that every blocked destination reaches this branch; which
destinations leave the route search empty is not established either.

**Unknown.** Native arrival and release, repeated path failure, sound/text,
removed or dead targets, the obstacle verdict and multi-tick parent outcome.
After a route failure, whether the next parent state-0xd/0xe evaluation
reissues kind 8/9 over the reacquisition order, or `R0004` leaves the
command state, is not established. A read of `R0004`'s `actor+0x50`
stores and of the parent evaluation's caller cadence, or a native trace of
order+0x08 and `actor+0x50` across ticks after a failed search, would settle
it. Whether cast-admission callees or unit-tick gates add range, facing or
owner checks remains unclosed. No whole-image absence is claimed.

**Amended.** The former mover+0x7c reach operand is corrected to the target
pointer; the non-centred branch reads caster position, not target position.
The unconditional pending-kind-retention consequence is narrowed by two
paths: route-search failure through mover+0x98 and the point obstacle
branch. The ordinary failed reach still selects approach, not immediate
release. The two correction entries preserve the former
wording; the kind-8 metric is independently MAGIC-REACH-179.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md),
[EXP-0497](../experiments/EXP-0497-out-of-range-cast/EXP-0497.md),
`evidence/instructions.tsv`, `evidence/orders.tsv`.

### MAGIC-240

Both cast install arms read `ord+0x60` after installing the cast. The kind-8 arm stores progress 2, `actor+0x54 = 0x0d`, `actor+0x5c` = target, `actor+0x64` = spell and `ord+0x5c = 1` (`L00034`..`L00035`), then loads `ord+0x60` (`L06124`) and, when it is 0, calls `R0401` (`L06125`); a nonzero value jumps to the common tail (`L06126`, a jump to `L00010` on non-equal). The kind-9 install `R0013` stores progress 2, state `0x0e`, the cell bytes at `actor+0x60`/`+0x61`, the spell and `ord+0x5c = 0` (`L00036`..`L00037`), tests `ord+0x60` at `L06127` and jumps to its exit (`L06128`) when it is nonzero. Its zero path holds the same stop and reset stores inline (`L06129`).

`R0401` (read whole) stores `ord+0x14 = actor+0x12c`, `ord+0x60 = 0`, `mover+0x7c = 0` and `actor+0x54 = 0`, then `actor+0x50 = 0x0c` (`L06130`), `ord+8 = 0` (`L06131`) and `ord+0x50 = 1` (`L06132`). It does not store `ord+9`, `actor+0x5c`, `actor+0x64` or `ord+0x5c`. With `ord+0x60 == 0` the install therefore clears the order kind and the actor state in the same call and leaves progress 2 behind. The next pass takes progress arm 2 (`L00093`), which writes `actor+0x54` as 0x0d when `ord+0x5c` is nonzero and 0x0e otherwise (`L00093`..`L06133`), so the actor re-enters the cast state with the fields the install left. With `ord+0x60 != 0` the order kind survives the install, and after the cast completes (progress cleared at `L00095`) the kind arm runs again for the same order.

The stores of the literal 1 at displacement `0x60` over the whole image (`evidence/scan-disp60-store1.txt`, only the immediate 32-bit store forms at `+0x60`) number ten, none of them a register-source store: `L06134` and `L06135` outside the AI module (the second stores to a stack local), and eight in the bodies of creature cast selectors, `L06098` (`R0209`), `L00061` and `L00063` (`R0015`), `L06136` and `L06137` (`R0113`), `L06138` (`R0393`), `L06139` and `L06140` (`R0394`). The stores of zero at `ord+0x60` read are `L06141`, `L06142` and `L06143` in `R0401` and `L06129` in the inlined reset of `R0013`.

`MAGIC-CADENCE-126` and `MAGIC-CADENCE-127` already describe retained order casts and the progress-arm retry; this claim adds the flag that makes an order retained. A creature's drawn cast is therefore a retained order. Between AI slots the same cast order is re-armed and executed again after each completed cast, until another writer replaces the order (`AI-376`).

**Confidence.** High for the reads, the two arms' branch, `R0401`'s stores and the ten literal stores (complete-body reads and a whole-image displacement sweep for the stores). Medium for the retained-repeat consequence: it composes the arm, the progress arm and the unit tick's state dispatch (`MAGIC-CADENCE-127`) and no run was observed.

**Unknown.** Register-source stores to `ord+0x60` are outside the sweep: `L01143`, a 32-bit store of a register value at `+0x60`, in `R0305` is one that `AI-FOLLOWHEAL-118` cites, and its register value was not traced. Readers of `ord+0x60` outside the three sites named were not enumerated, so "the arms" is the population read, not a census. Stores of other nonzero values were not swept. Which value player-issued cast orders leave in `ord+0x60` was not traced.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), `evidence/asserts-en.txt`, `evidence/listing.txt`, `evidence/spell-arms.tsv`

## Fire_Ball burst remainder

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-245 | Fire_Ball's transport target point is the footprint centre minus one sub-unit per axis for a target of size above 1 (`L06144`), and the cell centre of the cast's x and y bytes for a null target (`R0289`). | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| MAGIC-246 | The area effect's size byte `+0x49` is spell parameter 9 (1 for Fire_Ball), so the blast arm `R0630` applies the blast to cells -size to +size on both axes: a 3 x 3 area for Fire_Ball. | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| MAGIC-247 | The Ballista's weapon spell in the EN data is Fire_Ball with value 40 (Catapult: 70), taking the weapon-spell path and not a cast divert. | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |

### MAGIC-245

- Branch on the target's size byte (target slot `+0x1c`, `L06145` to `L06146`; `evidence/disasm-rider-position-L13150.txt`). Above 1: per axis, word = cell * 256 + fine + (size - 1) * 128 - 1 (`R0862` for x, `R0863` for y, a mask with 0xffff and a subtraction of 1), unpacked by `L06144` into cell (high byte) and fine (low byte); the packed word at `+2` is rewritten and the dword at `+8` stays from the earlier copy `R1010` (`evidence/disasm-centre-L06144.txt`, `evidence/disasm-size-offset-R0862.txt`). That is the centre of the size x size footprint minus one sub-unit.
- Null target (`L06147`): `R0289(position, x byte, y byte, global L04624)` gives cell (x, y), fine x and y 128 and the dword at `+8` = the global (`evidence/disasm-null-target-point-R0289.txt`). A point cast takes the bytes from the actor's `+0x60` and `+0x61` (call `L00008`, `evidence/disasm-point-cast-L13151.txt`).

**Confidence.** High for the arithmetic of both branches (instructions); Medium that a shot reaches the null branch only through a point cast: callers of `R0003` were read in `EXP-0441` and here only for the point cast.

**Unknown.** A saved Fire_Ball with a multi-cell or null target; none was read.

### MAGIC-246

- The constructor call `R0652` at `L06148` receives spell parameter 9 (`push 9` at `L05103`) and stores it as the size byte `+0x49`; duration `+0x4c` = parameter 11 shifted left by 4 (`evidence/disasm-area-effect-ctor-call-L13152.txt`). Fire_Ball's row has parameter 9 = 1 (`EXP-0432` `evidence/spell-pictures.txt`).
- The tick `R0629` takes the blast arm `R0630` when the stage tests `R0631` and `R0632` fail (`ANIM-114`). `R0630` sends the burst message (`L03054`, `evidence/disasm-effect-sender-R0635.txt`), applies `R0267` to the cells from -size to +size (calls at `L05407`, `L05408`, `L05409`) and sets the reap flag (`L03056`) (`evidence/disasm-blast-arm-R0630.txt`, `evidence/disasm-area-tick-R0629.txt`).
- The client's 3 x 3 registration is constant (`ANIM-123`), so a customised parameter 9 would change the simulated area and not the client's registration.

**Confidence.** High for the parameter flow and loop bounds (instructions); Medium for the per-cell effect `R0267`, not re-read here.

**Unknown.** Parameter 9 for spells other than Fire_Ball.

### MAGIC-247

- `EXP-0428` `shot-classes.csv`: Units row 27 (Ballista, class "Catapult 2", Projectile 6, charge 2, relax 50, ShootDelay 0) has weapon `Boulder Thrower{castSpell=Fire_Ball:40}`, `weapon_spell` true and `cast_divert` false; the Catapult (row 26, Projectile 5) has `Fire_Ball:70`. The EN weapon-spell table of `EXP-0191` lists the same two rows.
- Both therefore reach the Fire_Ball rider path of `ANIM-115`.

**Confidence.** High for the EN data: table rows of the EN `Data.bin`. Medium for the RU rows, which these files do not list.

**Unknown.** The RU weapon spells; none were read here.

## Prismatic Spray caster pointers, heading and facing

Evidence is a static read of `rom.exe` (one image on both lawful installs, sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`); no process was run. Every quoted instruction is asserted on both editions (`evidence/listings/asserts-en.txt`, `evidence/listings/asserts-ru.txt`, 0 differences). `EXP-0452` was allocated ids `241`..`244` of `claims/magic.md` and spent `MAGIC-241`..`MAGIC-243`; `MAGIC-244` is unused.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-241 | A spell 14 cast by a R0914-built caster reads the null owner +0x14 unguarded in R0126, called at selector entry before the group read and the empty-list test; a null primary faults earlier. | Medium / Unknown | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |
| MAGIC-242 | The Prismatic Spray heading is an 8-way multiple of 0x20 from the signs and a 2:1 magnitude test of the fine centres, each P + (n-1)*128 including the sub-cell; the selector ranks by edge gap and turn cost. | High / Medium | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |
| MAGIC-243 | No facing test is found in the three routines read (R0268, R0269, the selector); book and scroll casts are gated upstream by executor row 8 or 9 and a weapon cast by rows 5 and 6, not row 2 (Medium). | High / Medium / Unknown | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |

### MAGIC-241

Actor states `0x0d` and `0x0e` call `R0268`. For id 14 (`L03028`, a compare with 0xe) the code calls `R0904`, which reads no pointer, then `R0269`, which selects the item arm (`caster+0x68` non-zero) or the book arm (`R0253`) and calls the selector `R0109` (`L06149`) with the cap. The delivery test of `MAGIC-DELIVERY-170` takes the branch for id 14, so the early return at `L06150` for a caster with `+0x3c == 0` is not reached.

A caster built by `R0914` has `+0x14 = 0` (`L06039`), `+0x70 = 0` (`L06040`), `+0x3c = 0` (`L06151`) and health `+0x94 = 0x1e` (`L06152`); `MAGIC-235` records that no site assigns owner or group. The selector reads in this order:

1. `[primary+0x158]` at `L06153`: a null primary (state `0x0e`) faults here, address `0x158`;
2. `R0126(caster, primary)` at `L00878`, which reads `[caster+0x14]` at `L06071` and then `[owner+4]` at `L06072`: with a null owner it faults at address `4`, with no test;
3. `[caster+0x70]` passed to `R0110` at `L06074`, whose first read is `[group+0xc]` at `L06075`: a second fault site, not reached after the first;
4. the empty-list return at `L06154`, after all of these.

The three frames read (`R0269`, the actor tick `R0037`, the teardown `R0208`) have FuncInfo with nTryBlocks 0 (`evidence/listings/tryblocks.txt`).

**Confidence.** Medium: the stores of zero and the read order are asserted instructions on both roots; the fault is a static prediction, since no cast was run and shipped scripts contain no spell-14 instant (`MAGIC-SPRAY-135`). Unknown for a caster created any other way, and for any exception filter outside the three frames.

**Unknown.** Whether a script cast of spell 14 is ever built in play. What the process does on that fault.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-selector.txt`, `evidence/listings/rom-spray.txt`, `evidence/listings/rom-caster.txt`, `evidence/listings/tryblocks.txt`

### MAGIC-242

`R0051(caster, candidate)` calls `R0862` and `R0863` for each object (`L06155`..`L06156`). Each returns the fine point P from the position record (`R0165`, `R0166`) plus `(n-1) << 7` for the footprint size n (`L06157`, `L06158`), so a multi-cell footprint is measured from its centre and the sub-cell bytes are included (`MOVE-087`). The two differences feed a sign test and a 2:1 magnitude test (a comparison of the two magnitudes and a doubling of one, `L06159`..`L06160`), giving a raw index 0..15; the tail (`L06161`..`L06092`) adds 1 when non-zero, clears bit 0 and shifts left 4. The heading is therefore one of eight bytes, multiples of 0x20; a fractional offset moves it only through the sign and the 2:1 comparison.

`R0115(facing, heading)` returns the absolute byte difference folded to at most 0x80. The selector score is `((edgeDistance << 8) + turnCost) & 0xffff`, with the edge gap from `R0036` and the facing byte `[caster+0x154][0]`. Edge distance dominates; turn cost breaks ties. It is a ranking term, not a gate (`MAGIC-243`).

**Confidence.** High for the heading formula, the fold and the score form: asserted instructions. Medium for the edge-gap routine's footprint handling, read only through its call.

**Unknown.** The body of `R0036` is listed (`rom-heading.txt`, `R0036`..`L06162`) but its footprint handling was not decoded. Whether the 16-bit mask wraps in play.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-heading.txt`, `evidence/listings/rom-fine.txt`

### MAGIC-243

`R0268`, `R0269` and the selector `R0109` contain no compare of the facing byte against a heading with a branch to a refusal on any arm; the only use is the turn cost of `MAGIC-242`. The gates are upstream:

- book and scroll casts: executor row 8 (`L00108`, unit target; `R0041` is skipped when the caster is the target) or row 9 (`L05042`, point target; `R0086`), installed by the setters at `R0011` and `R0012` (opcodes `0x1e`, `0x1f`, `0x25`, `0x26`). `R0041(world, actor, target, range)` requires the facing byte `[actor+0x154][0]` to equal the bearing from `R0051` and an edge gap within range;
- weapon and caster-item casts (state 3 to `0x0d` at `L05238`, `L05239`): attack rows 5 (`L00101`) and 6 (`L00581`) call `R0041` and install action 3; row 2 (`L06163`) installs action 3 with no `R0041` call.

**Confidence.** High for the absence in the three routines, read over the listed bodies (`rom-spray.txt`, `rom-selector.txt` through `L06164`, the only facing-byte reads being the two turn-cost sites) and for the gates of rows 5, 6, 8 and 9: asserted instructions. Medium for any facing test outside those three routines. Medium for the arm-to-row mapping, composed from the setters and the actor slot.

**Unknown.** Which code issues pending order 2. Callers of the setters other than those read.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-cast-states.txt`, `evidence/listings/rom-setters.txt`, `evidence/listings/rom-death.txt`, `evidence/listings/rom-heading.txt`

## Control Spirit template, difficulty and the item-cast route

Evidence is a static read of `rom.exe` (one image on both lawful installs, sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`) and of `Data.bin` from `world.res` on each root; no process was run. Every quoted instruction is asserted on both editions (`evidence/anchors.tsv`, 105 anchors, 0 differences).

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-249 | The Control Spirit arm builds the new actor from the first Units row named `Ghost`; the corpse supplies seven listed stores and the row the rest; EN and RU rows are identical. | High / Medium | ● active | [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md) |
| MAGIC-250 | The mission difficulty is not applied to the raised actor by the arm or by the constructor chain; only the corpse's copied to-hit and defence words can carry it. | Medium | ● active | [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md) |
| MAGIC-251 | The item entry reaches the Control Spirit arm for id 25; the arm consumes the corpse before placement, finds it by cell, never reads or writes the target object and leaves the corpse on the dead list. | High / Medium | ● active | [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md) |
| MAGIC-252 | No authored Data.bin weapon or item, no executable shop weapon-spell table entry and no raw-scanned file carries spell id 25 on EN or RU; a shop-generated Scroll or Book can. | Medium | ● active | [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md) |

### MAGIC-249

**The raised actor is the first exact-name Units row `Ghost` plus seven stores from the corpse.** The arm pushes the `CString` built from the literal at `L05127` (`Ghost` on both roots) and calls `R0501` (`L03125`). The constructor reaches `R0184(name)`. A name of length zero skips the loader. The loop runs the Units ordinals from `0x1a` upward, skipping `0x1c`..`0x3e`, and compares the name to the row name at `row+4` with the CRT string compare, exact and case-sensitive (`R0751`); the probe's name match is case-sensitive as well and finds the same single row. The first match wins; its ordinal goes to `+0xc` and its row to `+0x3c`. The Humans table is not consulted. The Units table has 118 entries; the names `Ghost`, `Ghost.2`, `Ghost.3` and `Ghost.4` are adjacent entries, and the exact name selects only the unsuffixed one, which is also the loader's first match. The row has 55 parameters and empty equipment strings. The whole Units table (118 rows) is identical on EN and RU.

From the corpse the arm makes seven stores: reaction as `corpse/2 + 1` (`+0x86`), Mind (`+0x88`), Spirit (`+0x8a`), health maximum as `corpse/2` (`+0x96`), health set equal to it (`+0x94`), the first 16-bit word at `+0xa6` and the first 16-bit word at `+0xbe`. Both are 16-bit moves, so only the first word of each live block is copied. `+0xa6` is the to-hit block and `+0xbe` is the defence block; the corpse's defence word was already halved by the first dying tick (`L04104`, a signed shift right by 1). The owner (`+0x14`) is set to the caster's owner after placement. The arm reads no spell power. The new actor is registered in the global actor list and the owner's list, and is added to a new `0x48`-byte group object that is added to the owner's list.

**Confidence.** High for the resolution rule, the Humans exclusion, the copied fields and the EN and RU row identity: each is a named instruction (`evidence/anchors.tsv`) or a census row (`evidence/controlspirit.txt`). Medium for the claim that every other stat comes from the row: the `R0184` to `R0180` row streaming is taken from earlier experiments and was not re-read here.

**Unknown.** Whether code other than the arm reaches `Ghost.2` to `Ghost.4`. What the later words of the two blocks hold for the raised actor.

**Evidence.** [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md), `evidence/controlspirit.txt`, `evidence/anchors.tsv`, `evidence/exe.tsv`

### MAGIC-250

**No difficulty adjustment is applied to the raised actor in the arm or in the constructor chain.** The ALM spawner `R0151` applies the difficulty after its own `R0501` call (`L06165..L06166`), reading the session value at `[[L00285]+0x84]`: value 1 scales health maximum, value 3 adds `0x32` to the `+0xa6` and `+0xbe` words and scales health maximum, value 2 does nothing. The Control Spirit arm contains no read of that field; its `L00285` loads are the dead-list reads. The direct-call closures searched (`tools/ghidra/ConeNameRefs.java`, argument `L00285`, each root expanded separately) start at `R0501` (294 functions) and at `R0976`, `R0932` and `R0979`, which are actor vtable `L00001` slots `+0x3c`, `+0x40` and `+0x54`, plus `R0850` (136); none contains an instruction that names `L00285` (`evidence/closures.txt`). The arm's other callees were not searched for `L00285`: `R1145` (about 36 functions in its closure alone, among the 329 that closure holds), `R0059`, `R0411`, `R0032`, `R0149`, `R0198`, `R0152`, `R0044` and `R0063`; the vtable slots of `R0037` and `R0654` were not searched either. A reader that reaches the field through another base register or a copied pointer would not be seen by an instruction search on the global's address.

**Confidence.** Medium. The closures follow direct calls only, with 61, 15 and 36 computed sites not followed in the three large ones, and the search population is the one listed above. Because the corpse's own to-hit and defence words are copied, a corpse that was itself spawned under difficulty 3 passes that adjustment on.

**Unknown.** Computed call sites in the closures. The callees and vtable slots not searched. Whether any later tick rescales a raised actor.

**Evidence.** [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md), `evidence/closures.txt`, `evidence/anchors.tsv`

### MAGIC-251

**An item cast of spell id 25 reaches the arm, and the arm body does not read or write the target object.** `R0002` refuses only spell id 14 (`L06167`, a compare with 0xe) and otherwise calls `R0003(caster, target, 0, 0)`. Control Spirit's two damage parameters are -1 on both roots, so the damage gate takes the no-damage jump table, whose entry 25 is the arm (`L05128`); id 25 is the only id that dispatches there on both roots. Two callers reach `R0002`: `L00007` in the actor tick (caster item state `0xd`, target at `+0x5c`) and `L03045` in the melee strike `R0246` (fighter rider). Both pass target and caster as pointers. After the call the tick path clears `+0x64` and `+0x68` and sets `+0x58 = 7`, and the tick path deletes the objects at `+0x64` and `+0x68` through a virtual deleting call; nothing in the range `L06168..L06169` reads `+0x5c`; the melee path reads `target+0x94` and calls `R0561` and a virtual method at `+0x64`, and the strike requires damage above zero and a living target (`L06170..L06171`), so the target is not the consumed corpse. The arm selects the corpse by cell: the first actor on the dead list with stage `+0x13c == 2` whose packed cell word equals the target cell (`R1146`), the target cell being built in the entry's prelude (`L06172..L06146`) from the target position, which reads `target+0x10` and calls the virtual method at `+0x1c` (`R0863`, `R0862`), or from x and y when the target is null. The arm body itself never reads or writes the target object. A stage-2 corpse released its cells at teardown (`HERO-DEATH-026`), so a living actor can share its cell.

The arm stores health `-10001` and stage `5` into the corpse first (`L06173`, `L06174`), then constructs and places. Placement `R1145` is called with radius zero: it tries only the corpse's own cell. If placement fails, the new actor is deleted and the arm exits; the corpse stays consumed and nothing is raised. The `0x198`-byte allocation is null-tested at `L06175`, and that test skips only the constructor, so a null `this` would reach placement (`L06176..L06177`). The allocator `R1147` was not read, so whether it can return null is Unknown.

The arm does not unlink the corpse: none of the direct-call closures of the arm's callees contains a call to the list unlink `R1133` other than the group re-link under `R0152`, and the decay routines `R0867`, `R0860` and `R0208` contain none either (`evidence/closures.txt`). The dead list is saved whole (`SAV-DOC-053`); `SAV-DEADLOAD-127` and `SAV-DEADLOAD-128` already record stage-5 corpses and this route as a stage cause.

**Confidence.** High for the dispatch, the id-14-only refusal, the cell selection, the order of effects, the pointer use after the call and the placement-failure exit: each is a named instruction. Medium for the dead-list persistence, which is an absence over direct-call closures with computed sites not followed, and for the allocation-failure path, whose allocator was not read.

**Unknown.** Whether a script reads the dead list. Whether placement at radius zero fails in ordinary play (the blocking planes were not enumerated).

**Evidence.** [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md), `evidence/anchors.tsv`, `evidence/closures.txt`

### MAGIC-252

**The searched population (Data.bin census, executable shop tables, raw scan) carries no spell id 25.** The census read the cells of eleven Data.bin collections on both roots. It found 48 `castSpell` cells (46 in Humans, 2 in Units); the Magic collection's one `castSpell` occurrence is a definition entry, not a cell. Each resolves by name to a Spells ordinal, with 0 unresolved and 0 naming ordinal 25. The string `Control Spirit` occurs once per root, as the Spells row itself. The authored MagicItems collection (49 rows) carries no `castSpell` cell. The shop weapon-spell tables in the executable at `L06178` list spell ids 20 and 11 (flagged) and 1, 13, 14, 20 and 11 (unflagged), none 25. The table words and their two reading sites are asserted in `evidence/controlspirit.txt` and `evidence/anchors.tsv`. A raw ASCII scan of every file under both install trees (139 EN, 42 RU), case-insensitive, finds the identifier form `control_spirit` only in the executable, and the display form `Control Spirit` in the Data.bin container `world.res` on both roots and in the EN `Map Editor.exe` (`evidence/controlspirit.txt`). Shop-generated Scrolls (spell ids 1 to 27) and Books (a per-school set that includes 25) can carry id 25 within the shop price window (`SHOP-CONSUME-073`); those are not an authored population and were not bounded per shop.

**Confidence.** Medium: a bounded census of authored Data.bin cells, tables in the executable and a raw string scan. Saves, the compressed ALM and LM containers and script content were not read, so an id-25 reference there is not excluded; the shop generator's per-shop admission of id 25 was not decided.

**Unknown.** Which shops admit an id-25 Scroll or Book under their price cap. Whether a compressed ALM or LM container, a save or a script gives a weapon or item id 25.

**Evidence.** [EXP-0464](../experiments/EXP-0464-control-spirit/EXP-0464.md), `evidence/controlspirit.txt`

## Manual cast retention and parent reach reload

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-253 | Admitted manual actor/point setters clear order+0x60; parent states 13/14 reissue kind 8/9 and reload stored Spell+9 without changing the progress byte or that flag. | High | ✔ promoted | [EXP-0497](../experiments/EXP-0497-out-of-range-cast/EXP-0497.md) |

### MAGIC-253

The admitted non-null Spell branches of `R0011` and `R0012`
write EBX to order+0x60 at `L13499` and `L13500`. EBX is zeroed at the
function heads `L13501` and `L13502`. The 67 instructions between each
head and its admitted slice contain no EBX writer and five calls, which are
assumed to preserve EBX by the callee-saved convention. The store therefore
replaces both zero and nonzero previous values with zero. Their
target/Spell/range and progress boundary is otherwise the contract AI-356
establishes.

`R0008` dispatches parent states 13 and 14 through the table at
`L00322` to `L01993` and `L13503`. With a non-null actor target,
state 13 calls `R0018`, whose non-null branch writes kind 8 and
loads reach from Spell+9 at `L05040/L00075`. State 14 writes kind 9
and loads the same byte at `L13504/L13505`. These reached parent paths
leave order+9 and order+0x60 unchanged.

Each setter slice is entered with the EBX value that the executed head
`xor` leaves from a nonzero sentinel. A sentinel-entered control per form
stores the sentinel at order+0x60, so the stored value is the register, not
a constant. Eighteen original-instruction cases use both setter forms,
previous flag values 0/1/7 and progress 0/2/3. The setter copies Spell+9 =
33 to order+0x14. Spell+9 is then set to 35 before the first and 31 before
the second complete parent execution; order+0x14 reads 35 and then 31
after the two passes. The parent retains progress, leaves the active victim
unchanged and retains flag zero. No external service is substituted. Forty-nine
selected child-gate cases corroborate the reused self/facing/distance
contracts; they stop before approach or install and do not witness release.

**Confidence.** High for the named local stores, parent representation chain
and finite original-instruction contrasts. Retaining the previous nonzero
flag is excluded by the listing (head `xor`, no EBX writer before the store)
and the sentinel control; the five skipped calls rest on the callee-saved
convention, not on execution. Fixing reach at setter time is excluded by the
order+0x14 reads after each parent pass, each differing from the value
stored before it. Clearing progress at these parent evaluations is excluded
by the executed parent passes. These are admitted setter slices, not a full
command admission or native game run. Both locale executables are one image.

**Unknown.** Whether current power is recalculated before an individual
order; null-target selection, native event timing, route completion,
release success and feedback. Single-cast cleanup composes this zero flag
with MAGIC-240; no native repetition count is observed.

## Cast delivery start positions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-261 | The normal CUnit/CAirUnit cast producer starts at cached caster centre plus class ShootOffset minus Center, times 8; an empty array or picture 60 uses an unscaled Selection fallback. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-262 | Human state 8 uses idle weapon/shield name selection in R0551; the class-store path maps an unshielded empty-handed mage to 23 and staff categories to 24. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-263 | EN/RU units.reg has three distinct human ShootOffset arrays for mage/xbowman, mage_st and archer; the other 12 human class IDs use delta (-48,-57), with exact eight-direction integer vectors established. | High / Medium | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-264 | Among the 12 EN/RU Units rows with positive spell slots, Goblin_Sling.4, Orc_Bow.4 and Bat_Sonic.4 have ShootOffset arrays; the other nine use the normal cast producer's Selection fallback. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-265 | The inspected cast and cell-effect producers have different origins: Teleport forces the fallback and its second object's raw point is destination-derived; three message arms use cell centres. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-266 | Direct simulation delivery-2 admission copies caster Position into SpellTransport; its tick leaves Position unchanged, while the nested PointEffect copies the target Position. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |
| MAGIC-267 | The visible projectile SAV writer stores current +08/+0c as Prj<ID>/x and /y; its sixteen leaves have no separate launch-point copy from +28/+2c, so a moved projectile does not preserve its original point there. | High | ● active | [EXP-0502](../experiments/EXP-0502-spell-launch-point/EXP-0502.md) |

### MAGIC-261

`R0620` occupies slot `+5c` in CUnit vtable `L02468` and
CAirUnit vtable `L02585`. The inspected hero constructor creates CUnit.
The action-8 driver calls that slot at class ShootDelay; picture 60 uses
clock zero. The all-byte executable dword search finds two pointers to the
producer, both in those vtables. This is not a census of relative calls.

The complete origin window `L13393..L13394` uses class ID `caster+20`,
class array `L02113`, facing `caster+6c`, and cached centre
`A=(caster+58,caster+5c)`. With `i=(facing-8)&14`, nonempty ShootOffset
`class+ec` and picture other than 60 give
`A+8*(ShootOffset[i:i+2]-Center)`. Center is `class+34/+38`.
Otherwise each axis is `A+trunc((Selection2-Selection1)/2)-Center`, with
Selection at `class+84/+88/+8c/+90`. The fallback neither multiplies by 8
nor adds Selection1. The loader stores Selection values unscaled.

REG-UNITS-049 supplies the registry key-to-offset mapping: CenterX/Y at
`class+34/+38` and the ShootOffset CArray at `class+e8`, with data at
`+ec` and size at `+f0`. The fragment controls use that mapping; they
verify the arithmetic, not the key-to-offset mapping.

The producer writes new object `+08/+0c` and copies them into `+28/+2c`.
`R0614` derives A from cached `P28/P2c` plus
`128*(TileSize-1)` through the two TileSize getters. The unit-shot path in
SAV-1142 instead reads caster `+08/+0c`; those inputs are not interchangeable
without the geometry step. The complete origin window has no call, frame
lookup or equipment read. Equipment can select the class upstream.

**Confidence.** High for the exact integer formula and direct input population.
Both branches execute in 18,816 original-fragment controls across 28 class/row
inputs, 28 even pictures and eight directions. Translation/frame/equipment
controls and odd-facing aliases agree; both output copies agree. Static
memory operands exclude direct equipment and frame terms in this window.

**Unknown.** Native first-visible-frame position and external cache ordering
were not observed. Alternate producers and arbitrary memory aliases are
outside the population.

### MAGIC-262

`R0551` requires `caster+18c & 1`, reads slot-0 weapon `+15c`, and
uses `(byte[weapon+6]&31)-1` to select `main.res::text/heropicture.txt`.
Absent weapon uses index 0, `unarmed`. Slot-1 shield `+160` appends `_`.
Mage flag `+18c & 2` changes the exact name `unarmed` to `mage`.
In this selector, cast action 8 uses the idle weapon/shield name selection.
Action 6 bypasses it and uses `mage_st` or `unarmed`; this is the death control.

The subsequent original name-to-class chain maps `mage` to 23, `mage_st`
to 24, sword names to 3/4/5, axes to 7/8/9, clubs to 10/11, pikes to
12/13, archer/bowman to 14, xbowman to 15 and unarmed names to 1/2.
Shipped Data.bin Weapons rows 13 `Staff` and 14 `Shaman Staff` correspond
to name indices 12/13, both `mage_st`, on EN and RU. Unsupported suffix
combinations do not acquire a class meaning from this claim.
HERO-APPEAR-040 through HERO-APPEAR-046 supply the body-art route;
the geometry still comes from the selected units.reg class.

The cast-message window `L05328..L13395` sets action 8 without calling
the selector. At ShootDelay, the held class matches this selection when
the selector last wrote it for the current equipment and no other
`caster+20` writer ran since.

**Confidence.** High for the selector's static branches, name-to-ID stores
and isolated original name controls. The original selector and table accessor execute 144 times:
two mage-flag values, two shield values, twelve weapon cases and states 0/8/6.
Cast and idle names agree in all controls. The class chain has explicit
matching names and ID stores.

**Unknown.** The `caster+20` writer population was not enumerated. Hero
creation writes the message class `[L03223]` at `L02143`, then calls
the selector at `L03209`. The selector can return at `L04246` ->
`L03172` without a class write when `+18c & 2` is set and the body name
is unchanged. A writer census and native observation of the held class at
ShootDelay would settle cast-time selection freshness. Native
equipment-message freshness, changes during wind-up and unsupported names
were not observed. Successful class/art allocation is the boundary;
no asset pixels are claimed.

### MAGIC-263

Both shipped `graphics.res::units/units.reg` members are byte-identical.
The 16 human body IDs are 1,2,3,4,5,7,8,9,10,11,12,13,14,15,23,24.
All have Center `(64,78)`, Selection `(48,48,80,90)` and TileSize 1.
The effective ShootOffset arrays, in eight stored pairs, are:

| Class IDs | ShootOffset |
|---|---|
| 23,15 | 57,75;45,66;44,53;52,43;68,42;80,50;81,62;73,73 |
| 24 | 52,79;36,64;37,46;54,33;75,34;91,46;91,65;75,79 |
| 14 | 61,90;36,80;28,58;40,39;65,30;85,40;96,60;87,81 |

The other twelve IDs have empty arrays after the inspected parent fallback.
MAGIC-261 then yields `(-48,-57)` in every direction. For these TileSize-1
classes A equals cached `P28/P2c`. Original `L02828` resolves
N/NE/E/SE/S/SW/W/NW to facing 0/2/4/6/8/10/12/14, with positive y south.
The corresponding pair indices are 8/10/12/14/0/2/4/6; odd facings use the
preceding even pair. Exact integer deltas are in the evidence vectors and
the canonical presentation table. Teleport yields the fallback for all 16.

**Confidence.** High for registry values, parent provenance, original direction
labels and integer deltas. The parser consumes full registry members and
the original origin fragment verifies every human class in eight directions.
Medium for eight fine units per map pixel, retained from SAV-1130 and
UNIT-STRUCTDELIVERY-065; no native pixel witness upgrades that scale.

**Unknown.** Native first draw and art-dependent pixel extent remain unobserved.
The registry File descriptor is geometry provenance; dynamically selected
hero art is a separate route.

### MAGIC-264

The full EN/RU Data.bin parser consumes all 118 stored Units rows;
56 carry parameters and 12 have a positive spell ID in the three spell slots.
These rows and their class IDs are Goblin_Pike.4 (64), Goblin_Sling.4 (79),
Orc_Sword.4 (80), Orc_Bow.4 (65), Ogre.4 (66), Troll.4 (68),
Bat_Sonic.4 (70), Ghost.4 (69), Bee.4 (73), Squirrel.4 (74),
Foot_Animated.4 (75) and Turtle.4 (76). Derived rows agree across locales.

Classes 79 and 65 use pairs
`59,77;44,67;42,53;51,42;66,39;80,45;85,59;77,72`.
Class 70 uses eight `(64,64)` pairs and Center `(64,64)`, giving zero delta
for a normal non-Teleport cast. The other nine rows have no effective array
and use their own Selection/Center fallback, not necessarily zero.
The evidence records all twelve classes' geometry and vectors.
MAGIC-238 remains the authority for creature cast selection and aimed cells.

**Confidence.** High for this stored population and its inputs to MAGIC-261.
The parser joins each positive spell-slot row to its shipped graphics class;
original-fragment execution checks all twelve geometries in eight directions.

**Unknown.** A stored spell slot does not prove native visible-cast reachability
for that row. Modded rows and unexamined construction paths are outside scope.

### MAGIC-265

The population is all 28 shipped Spells rows joined to projectiles.reg,
normal producer `R0620`, and the three inspected client arms at
`L03226`, `L13396` and `L13397`. Normal pictures with nonzero
cast lifetime are 10,12,20,30,34,36,60. They share MAGIC-261's origin
window; picture 60 alone forces its Selection fallback.

Teleport creates a second object through copy constructor `L03008`,
which calls drawable copy constructor `R2159`. The producer then
uses copied action target `+88/+8c` plus the same Selection delta for its
raw `+08/+0c` at `L13398/L13399`. It preserves source `+28/+2c`
through construction. Driver tail `L13400/L13401` later copies raw
current coordinates into those cached fields. Thus the second object's raw
point is destination-derived, while its construction cache is source-derived.
The caster-centre placement clause of MAGIC-CASTSPAWN-033 is partially
retracted. This does not determine the first rendered Teleport point.

The odd-picture 0x86 arm writes
`(256*msg[0d]+128,256*msg[0e]+128)`. Source-cell opcode 0x8b writes
`(256*msg[0a]+128,256*msg[0b]+128)`. Source-cell 0x8c picture 36 uses
the low/high bytes of packed source word `msg+0e`, each multiplied by 256
and increased by 128. These message arms do not use the normal class origin.
MAGIC-DELIVER-035 supplies the no-client-ID sender boundary.

**Confidence.** High for the enumerated registry join and raw coordinate stores,
copy paths and driver-tail sync. Whole copy constructors have no unresolved
branches in the recursive instruction capture. Direct coordinate windows
distinguish caster, destination and message-cell alternatives.

**Unknown.** Global producer completeness, native reachability, external
geometry/draw ordering and the first displayed Teleport point remain unknown.
A native capture or complete draw-order trace would settle that presentation
boundary; no native game was run here.

### MAGIC-266

The direct-admission window `L05107..L03052` inside `R0003`
supplies caster Position `+10` from that function's `[ebp+8]` caster
argument at `L13402..L13403` to `R0633`.
MAGIC-DELIVERY-170 supplies the enclosing dispatch and caster argument.
The base chain through `L05979`, `R1009` and `R1010`
copies coordinate and
binding fields of that 12-byte Position; offsets 6/7 are not copied.
Class ShootOffset, equipment, frame and visible projectile
coordinates are not inputs to this direct Position copy. The nested
PointEffect constructor `R1046` separately supplies target Position
to assignment helper `L05872`. SAV-1068 retains the padding boundary.

The inspected `R0634` tick body has 43 reached instructions and
does not move Position. MAGIC-DELIVERY-170 supplies its countdown role;
MAGIC-225 and SAV-1068 retain admission/allocation/alias boundaries.
SAV-TOKENPOS-074, SAV-1054 and SAV-CASTCONT-1006 establish persistence in
the common Token/SpellEffect prefix. Token serializer `R0950`
calls `R1494` for the 12 Position bytes at body offset 0.

**Confidence.** High for the direct pointer and Position-copy chain and
unchanged tick field. Every captured body has zero unresolved branches.
This establishes simulation position, separately from visible origin.

**Unknown.** Native ordering, arbitrary aliasing, later writes outside this
tick body and native SAVE/LOAD continuation were not tested.

### MAGIC-267

The visible writer window `L13404..L13405` stores projectile `+08/+0c`
as top-level `Prj<ID>/x` and `/y`, with z and thirteen other action leaves.
`Projectiles/IDs` indexes those records. The sixteen-leaf list in
SAV-PROJSTORE-428 contains no dedicated leaf sourced from `+28/+2c`, which
MAGIC-261 initializes as coordinate copies. SAVE records the current point,
not an immutable origin. A moving Fire Arrow/Fire Ball can save a later
point; stationary Lightning/Prismatic retain the initial raw point under
the inspected driver in MAGIC-DELIVER-035.

SAV-1130 and SAV-1142 supply coordinate/layout and related unit-shot evidence.
SAV-PROJLOAD-429 supplies the loader's defaults and post-load helper;
this writer claim does not upgrade native restoration. Simulation transport
persistence uses the separate Position prefix in MAGIC-266, not these leaves.

**Confidence.** High for the complete writer leaf population, exact source
fields and absence of a dedicated coordinate-copy leaf within that writer.
The absence claim is limited to those sixteen leaves.

**Unknown.** Native resave, first post-load frame, cache restoration and
continued flight were not observed. A native SAVE/LOAD in flight would settle
those behaviors.

## Spell-object light

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-269 | The CProjectile light selector deposits vertex stamps for pictures 10, 12, 13, 34 and 36 during the client view rebuild. | High | ● active (amended) | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| MAGIC-270 | Lightning and Prismatic Spray overwrite the four vertices of each admitted drawn-path cell with the low byte of 10 times phase. | High | ● active | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| MAGIC-271 | Fire Arrow, Fire Ball flight and its explosion use clipped uniform point stamps with radii 0, 1 and an explosion phase table. | High | ● active | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| MAGIC-272 | Spell light is rebuilt per client frame; normal caster construction gives Lightning and Prismatic Spray 13 successful phase steps before shared cleanup. | High | ● active (amended, partially retracted) | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md); [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-273 | Dynamic lighting gates point stamps and bit-clear terrain lighting; bolt light still reaches units with the option off, while native activation and visible results of the L10964 & 2 set branch are Unknown. | High / Unknown | ● active (amended, partially retracted) | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| MAGIC-274 | The measured spell light is client draw state; the direct Prj SAV program saves 16 source scalars and omits the light grids and drawn-path array. | High / Medium | ● active | [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |

### MAGIC-269

R2160 is CProjectile vt+0x40 (slot L13406, vtable L02587). R1657 calls that slot in the primary store at L13407 and the deferred projectile store at L13408. The frame painter R0379 calls the rebuild at L13409. The unsigned selector uses pictures 10..36 and 27 index bytes. Five arms are reached over all 256 byte picture IDs; the jump table at L13410 is not in the committed evidence. Only 10, 12, 13, 34 and 36 reach a light store. Other integer pictures take this method's default.

The original-x86 probe runs all 256 byte picture IDs at phase 0 with both lighting options enabled and one supplied valid path point. Only those five IDs change view+0xa8. MAGIC-PIC-026 identifies their spell names; MAGIC-BOLTGATE-069 identifies the two drawn paths. The source is separate from the standing light, darkness and wall_of_fire cell bits.

**Confidence.** High for this complete local selector, vtable slot and named callers. The owned EN/RU executable is one byte-identical input, not two independent witnesses.

**Unknown.** R1657 has a third R1095 call at L13411 over a second store: radius 1, level selected by an 11-arm value >> 1 switch at L13412, default no stamp. This identified point source is unattributed. Other object classes and indirect lighting methods were not enumerated. Their own dispatch populations would settle the global question.

**Amended.** The target-count clause is narrowed to the five arms reached by the committed evidence; claims/retracted.md records the former wording. The five light picture IDs, vtable slot and named callers stand.

### MAGIC-270

Pictures 34 and 36 use the point array at object+0x114, count +0x118 and stride 8. For each signed point, c = pointX >> 5 and r = R0377(pointX, pointY). The method admits 0 <= c <= visibleColumns and 0 <= r <= visibleRows+4. With stamp stride visibleColumns+7, it writes (c+3,r+3), (c+4,r+3), (c+3,r+4) and (c+4,r+4), each as u8(10 * object.phase). These are path-cell vertices, not a radial halo at the object's position.

The stores at L13413/L13414/L13415/L13416 overwrite prior bytes. They do not add or take a minimum. Adjacent path cells share vertices. Empty paths and off-view columns deposit nothing. The original method and inverse-row helper run on a synthetic planar view; path controls distinguish an empty list, adjacent cells, an off-view point and overwriting an existing zero stamp with 40.

**Confidence.** High for the complete local arm and its executed stores. MAGIC-BOLTLIST-071 (superseded in part) and MAGIC-BOLTSTILL-072 supply path storage and stationary-object context.

**Unknown.** Native random path geometry, fog admission and overlapping object order were not observed. A native cast with initialized view grids would settle the resulting pixels.

### MAGIC-271

Picture 10 calls R1095(worldCellX,worldCellY,0,16). Picture 12 calls it with radius 1 and level 16. Picture 13 uses this phase table:

| Phase | Radius | Level |
|---:|---:|---:|
| 0 | 1 | 16 |
| 1 | 2 | 8 |
| 2,3 | 3 | 0 |
| 4 | 3 | 8 |
| 5 | 3 | 16 |
| 6 | 3 | 24 |
| 7 | 3 | 32 |
| 8 | 3 | 40 |
| 9,10 | 3 | 46 |

The helper subtracts view scroll. For nonnegative i,j <= radius it accepts i*i+j*j < radius*(radius+1), with threshold 1 at radius 0. Relative to the unpadded source cell, the X choices are 4+i and 3-i, and Y choices are 4+j and 3-j. Each reflected vertex gets the same byte; each write is separately viewport-clipped. There is no radial falloff. Unclipped radii 0,1,2,3 touch 4,12,32,52 distinct vertices. The original initializer at L13417..L13418 fills all 1681 square-distance entries; the original helper executes against it.

**Confidence.** High for the named point helper, phase table and executed vertex populations. This corrects the square-radius wording of MAGIC-UNITLIGHT-057 (partially retracted).

**Unknown.** Synthetic explosion phases above 10 take a default in the probe, but are not established native frames. Native edge pixels and overlapping source order require observation.

### MAGIC-272

R1657 copies the previous stamp at L13419 when L10964 == 2. It clears view+0xa8 to 255 at L13420 and its write count to zero at L13421 before object vt+0x40 calls. It finishes by calling R0608 at L13422. In the original reset control, four stamped vertices become unstamped and ambient 48 gives unit level 12 without a source. The current frame does not subtract an expired object's old contribution.

For pictures 34/36 from normal caster construction, R0558 writes phases 4,3,2,1,0,1,2,1,0,1,2,3,4 on successful calls 1..13 (MAGIC-CASTSPAWN-033, MAGIC-BOLTSTILL-072). The stamp bytes on that route are 40,30,20,10,0,10,20,10,0,10,20,30,40. The probe supplies actionsegments = 13; it does not enumerate every construction route. The direct 0x8b message route takes actionsegments from msg+0xf: 5 for spells 13 and 14, starting actionphase at -1 and giving phases 0,4,3,2,1 (MAGIC-DELIVER-035, MAGIC-281). A loaded object resumes from its saved actionsegments and actionphase (MAGIC-274). The common tail decrements actionsegments and returns 1; the next call with zero returns 0. The shared updater R0334 calls vt+0x3c at L02474, collects zero-return IDs, unlinks their store nodes and calls the scalar destructor at L13423. That cleanup arm has no picture filter. Removal prevents later deposits by that object. SAV-1133 supports this positive shared cleanup route.

Fire Arrow and Fire Ball flight retain their fixed stamp levels while present. On normal caster construction, MAGIC-CASTSPAWN-033 supplies distance-derived countdowns dist/200 for picture 10 and dist/384 for 12; the direct message route instead supplies its message counter (MAGIC-DELIVER-035). The normal picture-13 explosion receives 22 ticks (MAGIC-BURSTLIFE-034, amended only for staged-area cadence), with its two-tick sheet clock in ANIM-PROJ-025. The positive shared cleanup route applies when their driver returns zero. MAGIC-BOLTSTILL-072 supplies the existing bolt countdown authority. The bolt probe executes counter/phase instructions while skipping only the separate geometry call.

**Confidence.** High for the frame reset, positive cleanup route and executed bolt phase/countdown path with the normal caster counter. The named direct-message counter and loaded-counter route are scoped by MAGIC-DELIVER-035 and MAGIC-274.

**Unknown.** Exact native spawn/removal frame ordering and scheduling were not observed. A timed native cast is required to join successful driver calls to visible frames.

**Amended.** The universal 13-step clause is narrowed to normal caster construction; claims/retracted.md records the former wording. The direct 0x8b and loaded-object routes use their own supplied or saved counters. The direct sequence formerly written 4,3,2,1,0 is also partially retracted: MAGIC-281 executes its initializer and measures 0,4,3,2,1. Frame reset and positive shared cleanup stand.

### MAGIC-273

VIDEO-077 identifies Lighting at L06417 and Object animations at L05659. R1095 skips point stamps when Lighting is zero. Object animations controls its invalidation rectangle, not its stamp stores. The picture34/36 arm of R2160 reads neither flag. All four flag combinations give the same bolt vertices and original unit-grid merge. MENU-074 identifies the options OK path that clears Lighting when Animation is zero.

R0608 initializes view+0xb0 from ambient>>2. Each stamped cell maps each corner to max(stamp-32,0), substitutes ambient for a 255 corner, sums and shifts right 4, then takes min(result,ambient>>2). Four 255 corners leave the prior initialized/cell-bit value. Four equal bolt bytes at phases 0..4 give levels 0,0,0,0,2 at ambient 48. The merge has no Lighting predicate.

The software terrain call at L13424 requires nonzero write count, L10964 & 2 clear and Lighting nonzero. With that bit clear and Dynamic lighting off, bolt light reaches units and not the ground. R1816 selects terrain+0x18 at an unstamped corner and min(stamp,terrainByte) at a stamped corner. The L10964 & 2 set branch is selected at L13425..L13426; calls at L13427 and L13428 pass view+0xa8 without a Lighting test in the inspected local dispatch/call paths. R0541 uses a present stamp directly, falling back to terrain+0x18 at 255. R2161 constructs update regions from current/previous stamps; it is not a colour converter. The persistent terrain plane is not overwritten by these deposits.

The Lighting test at L13429 skips WallFire art and rejoins before the ordinary object reads of view+0xb0 at L10388/L10389. TERR-SPR-065 and TERR-LIGHT-061 (partially retracted) keep the Medium boundary for unit-body passes that force zero or interpret the argument differently. The same grid does not guarantee identical lighting for every drawable pass.

**Confidence.** High for the named local gates, stores, branch, stamp-else-terrain fallback, consumer arithmetic and executed synthetic controls. Native activation and the visible result of the L10964 & 2 set branch are Unknown. TERR-FAMILY-187 found no enabling writer in its file-backed embedded-address scan for L10968..L10969 and records the startup write of 0 at L10967; computed pointers, bulk copies and native lifecycle remain outside that search. No original game was run.

**Unknown.** Native activation and the visible result of the L10964 & 2 set branch, other options, fog admission, outer frame suppression, exact overlap ordering and native pixels. A native selector/lifecycle observation and casts under both Lighting states would settle the corresponding boundaries.

**Amended.** The hardware-mode label and Medium visible-result clause are withdrawn; claims/retracted.md records the former wording. The selector-bit branch and local fallback remain High. TERR-FAMILY-187 bounds native activation, and the bit-set visible result is Unknown.

### MAGIC-274

R2160 reads the object's attached client view at +0xe0 and deposits into view+0xa8. R0608 derives view+0xb0 for drawing. These methods write viewport light and invalidation, not simulation lighting. MAGIC-DELIVERY-170 is the separate simulation delivery authority; the measured presentation effect does not replace it.

The direct Prj writer in R0084, L13430..L08365, and loader in R0099, L13431..L13432, preserve 16 scalars: x,y,z,picture,dir,phase,lastaction,action,actiondir,actiontarget,actionx,actiony,actionz,actionphase,actionsegments,actionspell. Manager IDs and FreeIndex are separate leaves. The field program omits the viewport stamp and unit-level grids, and the point array +0x114/+0x118. The loader reattaches the client view, inserts into the projectile store and rebinds the drawable. SAV-PROJSTORE-428, SAV-PROJLOAD-429 and SAV-1130 supply the typed field authorities. SAV-914 and SAV-915 bound the direct consumer/producer program and retain their unexpanded-helper Unknown.

**Confidence.** High for the positive client writes and complete direct Prj scalar field program. Medium for absence of independent light persistence and simulation lighting outside these routines: other alias/helper populations were not independently enumerated.

**Unknown.** Identical Lightning/Prismatic light and random geometry after original LOAD before the first client tick. The bounded SAV population in SAV-1152 has no picture14+ witness. A live-object SAV, original LOAD, frame observation, post-load ticks and resave comparison would settle that boundary. Other simulation or SAV routes require their own enumeration.

## Bolt figure arithmetic and drawing

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-275 | Both inspected bolt producers pass the projectile display point first and the resolved target display point second; rotation subtracts the first generated sample, then truncates to coordinate words. | High | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-276 | The bolt walk admits only deflection magnitude greater than 2 and abscissa step greater than stored 0.15; accepted steps alternate sign, clamp the ordinate and stop at abscissa at least stored 0.7. | High | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-277 | The bolt builder inserts midpoint knots with -0.5, fits overlapping quadratic triples, and samples each triple every six canonical x units over a half-open interval. | High | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-278 | The bolt builder rejects a completed sampled figure only when a sampled ordinate exceeds the stored 0.15-times-length band; retry consumes the continued random stream and rebuilds every point. | High | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-279 | The bolt walk reads the current thread's CRT seed, shared with same-thread callers; the bounded direct-call census names potential consumers but does not determine the native interval between bolt ticks. | High / Medium / Unknown | ✔ promoted (amended) | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-280 | The selected software bolt drawer stamps each stored point once, in list order, using the own-table forward .16a receiver; sequential table blending and rectangle clipping preserve that order. | High / Unknown | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-281 | Normal caster, direct 0x8b and source-cell 0x8c bolt routes have different initial phase/countdown pairs; every live action-1 driver call invokes geometry after phase selection and before countdown decrement. | High / Unknown | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| MAGIC-282 | Finite bolt length uses a scaled PC64 hypot with explicit binary64 spills and restores the incoming control word; later arithmetic inherits that word, whose precision and rounding can change integer points. | High / Unknown | ✔ promoted | [EXP-0504](../experiments/EXP-0504-bolt-figure/) |

### MAGIC-275

The original producer calls L05528 and L05527 pass four doubles in this
order. Each SUB wraps to 32 bits before the result is interpreted as signed.

```
Ax = s32(projectile+50)
Ay = s32wrap(projectile+54 - projectile+10 - projectile+68)
Bx = s32(target+60)
By = s32wrap(target+64 - target+68 - target+10)
```

FILD dword and FSTP qword preserve each resulting integer exactly. Picture
34 passes tag 34; picture 36 passes the original victim-list index modulo 7.
Twenty-two original-producer controls execute construction, zero-size array
clear and target hash lookup without substitution. Independent field changes
and signed-wrap cases distinguish the endpoint order and every operand.
MAGIC-261 and MAGIC-265 supply the upstream raw construction origins; these
display fields are later inputs, not interchangeable raw world coordinates.

Let R round each arithmetic instruction under the incoming x87 control word,
and S store binary64 under that word. MAGIC-282 defines L. The horizontal end
is H=S(R(Ax+L)); the walk uses D=S(R(H-Ax)), not an algebraically substituted
L. Real deltas are dx=S(R(Bx-Ax)) and dy=S(R(By-Ay)). Rotation stores
c=S(R(dx/L)), s=S(R(dy/L)). For final canonical arrays X and Y:

```
u = R(X[i]-X[0]); v = R(Y[i]-Y[0])
rx = R(R(R(c*u)+X[0])-R(s*v))
ry = R(R(R(c*v)+R(s*u))+Y[0])
```

The original R0279 helper saves CW, ORs its RC bits with 0x0c00, FISTPs
to signed i64, restores CW and returns the low/high dwords. The point stores
only the low 16 bits of each result. The record is i16 x, i16 y, i16 zero,
u8 tag, u8 zero. This is truncation then word wrapping, not screen clipping.
The origin is the projectile end; the first generated Y can differ from Ay
by floating evaluation residue. Subtracting the literal Ay instead of Y[0]
does not implement the displayed instruction sequence.

**Confidence.** High for local argument order, conversions and ordered
rotation operations. The original instruction capture and native fragment
controls agree. This does not establish native cache refresh order.

### MAGIC-276

R1091 clears output arrays and starts raw arrays t=[0], q=[0]. It
sets verticalScale=S(R(L*0.03)) and band=S(R(L*0.15)). All constants are
stored binary64 values; MAGIC-282 names their bits and precision boundary.
The first rand selects sign +1 for an odd return and -1 for an even return.
Each attempt consumes two further draws, even when it is not admitted:

```
v = (rand()%7)*sign
d = S(R((rand()%50)*stored_0.01))
if abs(v)>2 and d>stored_0.15:
    qnext = clamp(R(q[-1]+v), -3, +3)
    tnext = S(R(t[-1]+d))
    append(qnext, tnext)
    sign = S(R(sign*stored_minus_1))
continue while t[-1] < stored_0.7
```

The initial sign draw is at L05538, the two attempt draws at L05539 and
L05540. The magnitude and step comparisons branch on less-or-equal at
L13433/L13434. Magnitudes 0,1,2 and step residues 0..15 are rejected.
The clamp and signed-scale store precede the next attempt. If the last t is
at least 1, the builder removes that last raw t/q pair, then appends (1,0).
Otherwise it appends (1,0) directly. The terminal raw endpoint is not the
last stamped point; MAGIC-277 defines sampling.

**Confidence.** High for this complete local walk and finite predicates.
Original instructions distinguish strict admission, alternating accepted
steps and independent pair consumption. No finite maximum retry count is
claimed.

### MAGIC-277

Transform raw point j by the ordered, stored operations:
Px=S(R(R(t[j]*D)+Ax));
Py=S(R(R(R(t[j]*0)+R(q[j]*verticalScale))+Ay)).
The local knot arrays start with P0 and P1. For each j from 2 through the
penultimate raw point, append M then Pj, where each component is
M=S(R(previous-R(R(Pj-previous)*stored_minus_half))). Here previous is
the last knot P(j-1). The exact factor -0.5 creates a midpoint; it does not
scale a quadratic coefficient. Append the canonical terminal endpoint last.

R1092 receives overlapping triples K[0:3], K[2:5], K[4:7], in that
order. It computes coefficients A,B,C for a quadratic y=(A*x+B)*x+C.
The reproducible coefficient program has 43 ordered operations and ten
binary64 stores. It preserves the determinant cancellation. With tN denoting
its arithmetic temporaries, outputs are A=S(t42), B=S(t41), C=S(t43). The coefficient
program and the complete instruction operands are in the evidence. A
closed-form polynomial is a real-arithmetic interpretation, not permission
to reorder those operations.

Each triple starts with held x=first.x. While held x<third.x, evaluate
R(R(R(R(A*heldX)+B)*heldX)+C), store y to binary64, and append binary64 x/y
to generator arrays +94/+a8, count +98/+ac. The next held x is
R(storedPreviousX-stored_minus_6). It is retained on the x87 stack for the
comparison and polynomial; x's array store is a separate rounding boundary.
The last endpoint is excluded. Shared triple starts are included once.
Six is horizontal canonical spacing, not arc length; the final interval can
be shorter, and rotation/truncation can change integer spacing or duplicate
points. Drawing does not insert intermediate stamps.

**Confidence.** High for local knot order, coefficient instruction program
and half-open six-unit sampling. Twenty native helper controls cover four
distinct triples and five CW values. All 33 PC53-nearest helper sample y
words match the independent 43-operation evaluator bit for bit.

### MAGIC-278

After all quadratic triples, L13435..L13436 visits every sampled Y in
array +a8, count +98. It accepts abs(R(Y[i]-Ay))<=stored band; it retries
when the comparison is greater, not when equal. The band is
S(R(L*stored_0.15)), not a screen-space clipping rectangle.

The recursive call at L05553 receives the same four endpoint arguments.
The next invocation clears its point/raw arrays and consumes the continued
CRT stream; no rewind or per-bolt seed occurs. It replaces the complete
sampled output rather than deleting offending points. Rotation and word
conversion follow acceptance. No deduplication or point clipping appears
in these complete producer/generator/rotation bodies. Screen clipping is
per stamp in MAGIC-280.

**Confidence.** High for the bounded pass order and strict retry predicate.
Among 300 native supplied-state controls, 58 encounter one rejected figure;
none exceeds one retry in that population. This is not a global retry bound.

### MAGIC-279

Original rand R0179 and srand R0291 both call R0228. TlsGetValue
uses index L13437; on null the getter allocates 0x74 bytes and installs a
block with TlsSetValue. Seed is block+14. Initializer L12226 writes 1.
The recurrence is state=214013*state+2531011 modulo 2^32; return is
(state>>16)&0x7fff. CRT-created threads initialize their own block and do
not inherit the creator's seed.

Three direct srand sites take timeGetTime at L12227 and L12228, and
time(0) at L00376. The walk has no local seed. Synchronous routes R0260
and R0454 call simulation tick, client dispatch and message 0x401 on the
same thread. The separate callback R0261 is created through CRT R0457
and has its own block. Client/simulation sharing therefore depends on route.

Raw E8 scanning covers all executable PE sections. It retains 88 rand sites:
87 inventory-boundary matches across 44 owners and orphan L13438, whose
local start/fallthrough is decoded but native reachability is unknown. It
also retains three srand sites, 44 integer-range wrapper calls and one float
wrapper call. The whole-file dword scan finds no rand/srand/getter entry
pointer; computed addresses remain outside this population.

Potential same-thread consumers include Heal R1103, Drain R1104,
music selection R2050 and shuffle R1230, AI spellbook R0009,
idle turn R0205, roaming R0155, and the measured map/actor/item range
wrapper callers. A refused walk attempt and each regenerated Prismatic link
also advance the stream. These are possible consumers, not a runtime interval
list. One figure is reproducible from the pre-generation thread seed plus
the numeric inputs/CW. Cast, object and tick identities alone do not fix it.

**Confidence.** High for TLS, recurrence, initialization and the named seed
operand sources. Medium for the bounded potential-caller classification.
Unknown for the executed caller sequence, thread assignment and seed of a
native bolt. A trace of thread IDs, srand arguments, every rand seed/return
and consecutive walk entries would settle the chosen mode's interval.

**Amended.** The thread-assignment Unknown becomes Medium: the bolt walk is
placed on the main thread by exclusion (MAGIC-283), and the separate callback
R0261 is the server loop, whose launcher has no reference (SESS-082). The executed caller sequence and the seed at a native bolt stay
Unknown; SESS-083 gives the reseed points.

### MAGIC-280

Each arm of R0556 visits point indices 0..count-1 and sends one call
to sprite vt+18: (i16x-8,i16y-8,frame,0,0). Fixed sheets are record 34
lightnin, five 16x16 frames, and record 36 chain, thirty-five 16x16 frames.
Frame34=phase. Frame36=phase+5*u8tag; generator tag is victim index modulo 7.
All points of one link share its frame, while distinct links can differ.
There is no facing fold, interpolation, b-sibling pass or trail in these arms.
The thirteen phase values are 4,3,2,1,0,1,2,1,0,1,2,3,4 for actionphase 1..13.
Outside that range the driver retains the previous phase. MAGIC-281 names
construction-dependent sequences.

The selected own-table .16a receiver is R1785, forward decoder R1518.
It uses sprite+1c and the current framebuffer. For each literal source word W,
it adds read16(sourceTable+W) to a destination-table u16 modulo 65536.
Normal destination offset is 2*entries+2*old+((W<<8)&0x1e0000).
Low-memory selector 1 uses 2*(old>>3)+((W<<5)&0x3c000).
For installed even words, L=(W>>9)&15; normal source palette scaling is
(L+1)/16 and destination scaling (15-L)/16, separately quantized and packed.
Low-memory source uses /18 and destination row L, not the normal extra row.
SPR16A-080 and PAL-MODE4-010 supply table-generation authorities.

Each stamp uses active [left,right) x [top,bottom) globals L01503..14,
buffer L01168 and stride L01511. Whole exterior stamps return; crossing
stamps trim forward literal runs and skip excluded source words. No point
list entry changes. Later stamps blend with earlier output.
Normal projectile insertion joins view+9d4 to software painter R0379's
collection pass: buckets ascending, then node+0 chains; payload vt+28 is
called at L02994 with (0,0,0). This follows selector-3 shadows and precedes
retained-area/selector-3 body passes, markers/bars and shroud.

**Confidence.** High for the joined software path and finite measurement:
22 stamp captures, 17 ramp cases, three collection cases and 1120 installed
frame/clip/memory-mode decoder comparisons. EN/RU executables are one witness.
Unknown for native framebuffer words, packing masks, clip, stride, tables,
current collection and frame gates. Capture them before/after drawing and
byte-compare the resulting buffer to settle native pixels and overlap order.

### MAGIC-281

The original driver tests nonzero actionsegments before geometry. On action 1
it increments actionphase, selects/retains phase, invokes vt+50, then decrements
actionsegments and returns 1. With zero it returns 0 before either operation.
Other actions bypass geometry. The thirteen-entry ramp is bounded unsigned;
an actionphase outside 1..13 preserves the preceding phase.

| Named construction | Initial actionphase | Counter | Successful phase values |
|---|---:|---:|---|
| Normal caster, pictures 34/36 | 0 | 13 | 4,3,2,1,0,1,2,1,0,1,2,3,4 |
| Direct 0x8b, spells 13/14 | -1 | message 5 | 0,4,3,2,1 |
| Source-cell 0x8c, picture 36 | -1 | 13 | 0,4,3,2,1,0,1,2,1,0,1,2,3 |
| Loaded Prj | saved value | saved counter | Ramp indexed by saved value+call, otherwise retained phase |

The direct initializer writes -1 at L02569; source-cell writes -1 at
L02570 and 13 at L03020. The base constructor initializes phase to 0.
N positive saved remaining calls do not imply N new ramp entries. The next
zero-counter call returns 0; shared updater R0334 collects that result,
unlinks the node and calls the scalar destructor at L13423.

Producer admission still matters: unresolved picture-34 target retains its
previous list; picture 36 clears then appends only resolved victim entries,
in their original index order. The schedule promises a vt+50 invocation,
not a resolved target or a visible frame. MAGIC-BOLTLIST-071 (superseded in
part) and MAGIC-274 supply list/loaded-scalar boundaries.

**Confidence.** High for executed original route initializers and 96 driver
calls across 17 controls. Geometry is captured at a substituted vt+50 stub;
these controls establish call timing and state, not random path pixels.
Unknown for native spawn/update/draw/removal ordering, redraw counts per
tick and loaded first draw. A timed native cast or live-object LOAD trace
would settle the selected route's visible-frame schedule.

### MAGIC-282

L05533 delegates the finite hypot to L13439. For differences of
signed-i32 endpoints and incoming masked precision exception, it installs
CW 0x133f: PC64, nearest, all exceptions masked. It selects m=max(abs(dx),
abs(dy)); m=0 restores caller CW and returns +0. For m>0, R64 means x87
64-bit-significand arithmetic, nearest; S53 means a separate binary64 store.

```
a=S53(R64(abs(dx)/m)); b=S53(R64(abs(dy)/m))
q=S53(R64(R64(a*a)+R64(b*b)))
h=S53(R64(sqrt(q)))
(fm,em)=frexp(m); (fh,eh)=frexp(h)
p=S53(R64(fm*fh)); ep=encodedExponent(p)-1022
n=ep+em+eh
top16=(oldTop16(p)&0x800f)|((n+1022)<<4)
```

Normal frexp replaces exponent bits with 1022 and reports oldExponent-1022.
The last helper rebuilds p's exponent, preserving sign/fraction; it is not
another multiply. FSQRT runs with CW 0x037f, then restores 0x133f; the masked
precision path can add an equivalent binary64 spill. Hypot stores length,
restores caller CW, then reloads length. It must not be replaced by one
binary64 sqrt(dx*dx+dy*dy).

Later operations use caller PC/RC, including coefficient arithmetic, stores,
walk comparisons, midpoint operations and rotation. __ftol temporarily sets
RC to truncation and restores it. Stored constants' bits are:

| Value | Binary64 bits |
|---|---|
| -6 | c018000000000000 |
| 0.15 | 3fc3333333333333 |
| 0.03 | 3f9eb851eb851eb8 |
| 1 | 3ff0000000000000 |
| -1 | bff0000000000000 |
| 0.01 | 3f847ae147ae147b |
| 2 | 4000000000000000 |
| 3 | 4008000000000000 |
| -3 | c008000000000000 |
| 0.7 | 3fe6666666666666 |
| -0.5 | bfe0000000000000 |

**Confidence.** High for the finite masked length sequence, exact operand
bits and conditional arithmetic recipe. Three hundred original native x86
fragment controls use six endpoint pairs, ten seeds and five explicit CWs.
Against CW027f, integer point lists change in 35/60 PC64, 38/60 round-down,
32/60 round-up and 41/60 truncating-round controls. Each completed host run
executes original geometry, arrays, rand, hypot and conversion; allocation,
free and TLS access are controlled substitutions. No application entry runs.

**Unknown.** The CW at a native bolt entry was not observed. MAGIC-092
records prior startup CW027f, not a bolt-entry witness. A native breakpoint
must capture CW before R1090, alongside endpoints, TLS seed and target
resolution order. Cache scheduling, arbitrary IEEE-special/unmasked inputs
and native buffer identity are outside these controls.

## Random stream placement and call forms

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MAGIC-283 | Of the MAGIC-279 census, 87 `rand` sites, 44 range-wrapper calls, the float-wrapper call and 3 `srand` sites run on the main thread; orphan `L13438` is not placed. | High / Medium | ✔ promoted | [EXP-0507](../experiments/EXP-0507-random-stream/) |
| MAGIC-284 | `R0861(n)` returns 0 without drawing when n is 0 and otherwise draws once for `rand()*(n+1)` shifted right 15; `R1715(n)` is `1+R0861(n-1)`; `L13564` is `rand()/32767.0`; dice `L13565` has no caller. | High | ✔ promoted | [EXP-0507](../experiments/EXP-0507-random-stream/) |
| MAGIC-285 | Every consumer family named for the shared stream draws through raw `rand()` with an inline scale, the range wrapper or the float wrapper; existing claims give each family's arithmetic, and the random item kind and school timers are added. | High / Medium | ✔ promoted | [EXP-0507](../experiments/EXP-0507-random-stream/) |
| MAGIC-286 | A music list replace draws its start track as `rand()%n` after the stop and before the order build; an order build draws 2n times only in Random Order mode. | High / Medium | ✔ promoted | [EXP-0507](../experiments/EXP-0507-random-stream/) |

### MAGIC-283

Input: the 88 `rand` sites, 3 `srand` sites, 44 range-wrapper calls and one
float-wrapper call of MAGIC-279, each re-verified as an `E8` call to its
target.

| Target | Direct path from an anchor | By exclusion, stored roots | By exclusion, some roots unreferenced | Not placed |
|---|---:|---:|---:|---:|
| `rand` | 19 | 60 | 8 | 1 (`L13438`) |
| range wrapper `R0861` | 24 | 14 | 6 | 0 |
| float wrapper `L13564` | 1 | 0 | 0 | 0 |
| `srand` | 3 | 0 | 0 | 0 |

The anchors are InitInstance `R0326`, OnIdle `R0560` and the frame
window procedure `R0701` (SESS-082). A site is placed by exclusion when its
owner is reached only from functions stored in data (vtable slots, callback
tables) or from roots without a reference, and no launched secondary thread's
closure contains it. 13 `rand` sites, 24 range-wrapper calls, the float-wrapper
call and `srand` `L00376` also lie in the unlaunched server loop's closure;
each of them also has a direct path from OnIdle or the frame window procedure.
The roots without a reference include the dice wrapper `L13565`
(MAGIC-284).

**Confidence.** High for the 47 direct-path placements. Medium for the
exclusion placements, which inherit SESS-082's bound on library and
unresolved computed calls.

**Unknown.** The owner and reachability of `L13438`.

### MAGIC-284

- `R0861(n)`: `n == 0` returns 0 at `L13566` with no call. Otherwise one
  `rand()` call; `rand()*(n+1)` is divided by 32768 with truncation toward zero
  (`cdq`, `and edx, 0x7fff`, `add`, `sar 15`). The product is a 32-bit
  `imul` at `L13567`. For 0 ≤ n ≤ 65537 that is
  `floor(rand()*(n+1)/32768)`, range 0..n; above 65537 the product can wrap.
  n = -1 draws and returns 0.
- `R1715(n)` passes `n-1` and adds 1, so n = 1 returns 1 without drawing.
- `L13564` is one draw, `rand()/32767.0` (`[L13568]` = 32767.0), range
  0.0..1.0 inclusive.
- `L13565(k, s)` sums k calls of `R0861(s-1)` and adds k (k dice of s
  sides); no reference to it exists.

A consumer that counts draws therefore counts no draw for a zero-width range.

### MAGIC-285

Forms: raw `rand()` with an inline scale; the range wrapper (MAGIC-284); the
float wrapper. The inline scales are `n*rand()/32767` (multiply by `0x80010003`
or `idiv` by `0x7fff`), `n*rand()/AImgr[0]` with `AImgr[0] = 0x8000`
(AI-RANGE-102), `rand()/511` (`0x80402011`), `rand()/10` (`0x66666667`),
`rand()/65` (`0x7e07e07f`), shifts and masks, signed `rand()%k` by a
constant, and unsigned `rand()%n`.

| Family | Owners (sites) | Form | Authority |
|---|---|---|---|
| AI roam, idle turn, guard and aggression rolls | `R0155` (1), `R0205` (2), `R0258`, `R0259` (1 each), `R0024` (2), `R0161` (2), `R0175` (2), `R0146` (2) | raw, `n*rand()/AImgr[0]` | AI-ROAM-025, AI-TURN-104, AI-JITTER-103, AI-DEADROLL-078, AI-GUARD-012, AI-POST-095 |
| AI spellbook slot pick | `R0009` (5), `R0393` (5), `R0394` (5), `R0293` (1) | raw against thresholds, `n*rand()/AImgr[0]` | AI-341, MAGIC-221, MAGIC-AI-012, MAGIC-AIBIT-082, MAGIC-MASKREAD-081 |
| combat damage, hit, Bless and Curse | `R0265` (6), buildings `R0653` (1) | range wrapper | HERO-DAMAGE-022, UNIT-STRUCTDAMAGE-064 |
| Meteor, Heal and Drain, bolt walk | `R0639` (2), `R1103`, `R1104` (2 each), `R1091` (3) | range wrapper; raw `rand()%3`, `rand()%360`; raw parity, `rand()%7`, `rand()%50` | MAGIC-093, MAGIC-CLOUD-065, MAGIC-276, MAGIC-279 |
| mission load drop pick, scatter, seating; death gold; trigger return | `R0065` (3), `R0945` (4 and the item kind), `R1145` (2), `R0208` (2), `R0948` (1) | range wrapper | MISSION-DROP-002, ITEM-SPAWN-027, MISSION-SEAT-012, TRIG-RETURN-042, MISSION-DROP-074 |
| shop stock | `R1779` (3), `R1711` (2), `R1038` (2), `R1031` (2), `R1035` (2), `R1627` (2), `R1715` | range wrapper | SHOP-RNG-008, SHOP-EFFALT-071, ITEM-155, SAV-775 |
| random item | `L13569` (float), `L09853` (1), `L13570` (3) | float then range wrapper | this card |
| town square | `R1906`, `R1383`, `R1489`, `R1924`, `R1925`, `R1907`, `R1916`, `L11546`, `R0706` | raw, `n*rand()/32767`, `rand()%100` | TOWN-505, TOWN-MARKER-508 |
| tavern | `R1902` (3) | raw, `rand()/16 + 3000` | TOWN-409 |
| shop interior | `L09858` (2), `L09664` (1), `R1539` (1) | raw, `n*rand()/32767`, `rand()%5` | SHOP-109, SHOP-ANIMATION-085 |
| school | `L11522`, `L11521` (2 each), `R1971` (3), `L13571`, `L13572` (2 each) | raw, `n*rand()/32767`; `rand()/10` | TOWN-499, this card |
| music | `R2050` (1), `R1230` (2 per swap) | raw, unsigned `rand()%n` | MAGIC-286, VIDEO-MUSIC-056 |
| voice speaker and recording | `R2089` (3), `R0578`, `R0579` (1 each) | raw, `n*rand()/32767`; `rand()>>13`, `rand()>>14` | VIDEO-067, ANIM-119 |
| item stars | `R2171` (2 per pair) | raw, `rand()/511 + 8` | ITEM-STARPIX-098, SESS-083 |

Within one routine the draws run in address order along the executed branch;
`evidence/sites.tsv` gives each site's instruction window.

- **Random item kind.** `L13569` draws u from `L13564` and selects
  `Weapon` (u < 0.4), `Armor` (u < 0.65), `Shield` (u < 0.8) or `Potion`
  (`[L13573]`, `[L13574]`, `[L13575]`), then calls the generator
  `L09853` with that kind; the generator's range calls are
  `R0861(12)+1` at `L13576` and, in `L13570`, `R0861(v-lo)+lo`
  with v from `L13577`, `R0861(5)` and `R0861(16)`. Callers: the scatter `R0945`,
  `L08216` (two), `L13578` and `L09853`.
- **School timers.** `L13571` and `L13572` each draw `rand()/10` into a
  delay global (`L13579`, `L13580`) on first use, and again when more
  than delay + 3000 ms have passed since the stamp while `+0x324 & 3` is clear (`L13581`, `L13582`,
  from `timeGetTime`); both are reached from the school paint `R1487`.

**Confidence.** High for the forms, read at each call site. Medium that the
family list is complete: it is bounded by the MAGIC-279 direct-call census.

### MAGIC-286

- **Replace** `R2050(list)`, when the player buffer `+0x9c` is set: stop
  `R1233`; copy the list; n = its count; one draw, `rand()%n` (unsigned
  `div`) at `L13583`; then `R1232(start)`.
- **Order build** `R1232(value)`: calls `R1230(+0x20)`, which writes the
  identity order 0..n-1 and, when the mode word `+0x20` is non-zero (Random
  Order), runs n swaps, each drawing a = `rand()%n` then b = `rand()%n` and
  exchanging `order[a]` and `order[b]` (VIDEO-MUSIC-056). It then sets the
  current position to the index holding `value`, so the start track does not
  depend on the order.
- **Stop** `R1233`: when `+0x9c`, `+0x10` and `+0x18` are all non-zero it
  calls `R1232(+0x1c)`, an order build.
- A replace therefore draws 1 time with Random Order off, and 1 + 2n times
  with it on, plus any order build its stop makes. The Random Order setter
  path through `L03131` calls `R1230` and `R1232` directly.
- Callers of `R2050`: nine screen routines reached from the frame window
  procedure (`R1315`, `R1317`, `R0816`, `R1318`, `R1319`,
  `R1320`, `L08084`, `R0909`, `R2087`). The service `L13584`,
  run from OnIdle (SESS-082), calls the stop at `L13585`. All run on the
  main thread.

**Confidence.** High for the draw counts and order, read whole. Medium for
when the service's stop call runs.

**Unknown.** The service's stream-end conditions that reach `L13585`.
