# Unknowns

[Reference](format.md)

## Remaining boundaries

Complete effect-producer and target-class coverage beyond the named item/
spell paths, the full writer set for actor+0x4c bit 2 and candidate-list
membership beyond the specified acquisition routes remain Unknown.
Item construction and modifiers are described in [ITEM](../item/format.md),
spell lifetime in [MAGIC](../magic/format.md) and acquisition in
[AI](../ai/targeting.md). Their named paths do not establish universal coverage.

The persistent 100/50 invariant in HERO-REGEN-021 is retracted: those values
are constructor/template initialization, and saves can restore different
periods at `+0x98/+0x9e`. — HERO-REGEN-021, SAV-REGENWIRE-532

The [purchase arithmetic table](experience.md) specifies the first signed,
storage-width and x87 conversion boundaries. Native execution downstream of
unexecuted crossings remains Unknown, and the six-slot total can cross a
dword boundary earlier depending on the other slots. Derive clamps do not
remove these transitions. — HERO-SKILLUP-073, HERO-SKILLBUY-076,
HERO-GENERAL-089

The school Train widget directly produces purchases for slots 1–5, not
General. Equipment numbers follow ITEM-LADDER-019 through ITEM-WEAPCOL-021;
the ladder/column clauses of ITEM-SCALE-017 and ITEM-DMGCOL-018 are
retracted. Character-sheet damage is base through base+spread, while the
server transmits the raw pair. — HERO-SHEET-038
