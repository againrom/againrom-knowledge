# Weapon rows and identity lookups

[Reference](format.md)

## Authored cells and instance fields

Weapons uses 1-based definition rows. Both installed collections store
count 28 and serialize 27 rows, each with 17 parameters. Shape[0],
Material[1], Slot[4], absorbtion[10] and Other Parameter[16] have measured
authored values, but no operational reader was located in the named
construction, fill, price, Equip and descriptor bodies. This bounded result
does not establish image-wide absence. Instance Shape/Material come from
supplied bytes or name parsing; the Weapon picture class is literal 1.
— DAT-WEAPCELLS-023, ITEM-219

The group-C tail at runtime row+0x28 is a CStringArray. All 54 installed
Weapons arrays are empty. Its serializer and destructor process strings;
the semantic use of a nonempty element remains Unknown. Settle it by
tracing an embedded-array accessor and an element's consuming decision.
— DAT-STRARRAY-022

## Suitability and range

Parameter 15 sutableFor bit 0 produces fighter descriptor bit 2; bit 1
produces mage descriptor bit 4. The client wear predicate uses those bits.
Bit 0 also filters ordinary Weapon shelf candidates; collector mode 6
skips that class filter. Death tests the whole dword against zero before
transferring the weapon into the corpse container. Individual consumers
of higher bits remain Unknown; complete a row-typed field census to settle
them. — ITEM-220

The packed-code factory's Weapon constructor binds the decoded row and
fills byte +0x50 from parameter 11 @.range, or 1 for -1. The base range
pair, tag 38, requires suitability bit 0 and +0x50 above 1. The non-class-14
castSpell text arm constructs an item from its code and reads that item's
+0x50 under main[123], independently of the Spell's +9 range. It supplies
Staff 5 and Shaman Staff 4. Stream receiver and final painter behavior
remain Medium; a reconstructed range need not equal an archived live byte.
— ITEM-221

## Identity joins

The loader reads a u32 count, resizes Weapons, and dispatches entries
1..count-1. Backward name lookup returns row 0 on failure. Group-C GetAt
computes data+row*0x3c without a size test. Weapon LOAD rebinds the archived
Token row byte; it also restores a separate packed code and Shape/Material
bytes. No local mask or row-bound check guards that rebinding. Behavior
after invalid rows remains Unknown beyond these local bodies. — ITEM-222

The named consumer keys and missing-key results are ITEM-223.

| Consumer | Key and established missing-key behavior |
|---|---|
| Inventory and equipment pictures | Packed material/class/Shape/row forms a seven-digit leaf; selected icon failure is below |
| Hero animation class | heropicture[row-1], shield suffix and class-name chain; unmatched name retains class 1; the text getter has no file-local bound |
| Combat swing sound | Drawn units.reg class Sound[0]; zero id or null buffer requests nothing |
| Physical projectile | Drawn-class record +0xd4; missing target builds no shot; missing projectile resource remains Unknown |
| Shop speech | Eligible effect tag, book value or packed class/row with variant; missing-file playback remains Unknown |
| Shop lists | Collection rows, Shape/material masks and price, then class filtering |
| SAV | Archived row byte plus separate packed code and Shape/Material bytes; local lookup is unchecked |
| Scripts | Compiled low-word packed item code; search miss gives null, false or distance sentinel 255 according to the check |

The table is a set of named paths, not a complete resource-dependency
graph. Resolve missing animation, projectile, worn-layer and speech
behavior by tracing each exact loader and consumer error path. — ITEM-223

## Material masks and picture misses

Random selector L13570 retries all-zero rows, zero-mask Shapes and clear
material bits. Shop walk R0963 skips a clear bit before the price
window. The examined direct constructor, fill and Weapon LOAD bodies have
no local mask test. Further callback validation remains Unknown; follow
the complete path for a concrete shape/material/row to settle it.
— ITEM-224

Diagnostic icon open returns failure without selecting a replacement
picture. The examined inventory sprite constructor instead enters an error
message call and helper R0754. Its complete termination/callback outcome
and the other figure-layer paths remain Unknown; read those callees and
their installed dispatch bindings to settle them. A missing node therefore
has no established universal blank-picture fallback. — ITEM-225
