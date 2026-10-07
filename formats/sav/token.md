# SAV Token and Position

[Format reference](format.md) · [Object programmes](objects.md)

## Token head

Every Token-derived body begins with these 37 bytes, serialized by
`R0950`. Player and Diary do not have this head.
— SAV-TOKEN-034, SAV-MEMBER-036 (its unread-classes Unknown superseded by
SAV-EMBED-039)

| Body offset | Type | Source | Meaning |
|---:|---|---|---|
| `0` | raw 12 | `*(this+10)` | Position object below |
| `12` | u32 | `+04` | Runtime creation-order ID |
| `16` | u8 | `+0c` | Registry/definition selector; class-specific use |
| `17` | u16 | `+0e` | Type word; Human LOAD compares it with `0x21` |
| `19` | u32 | `+08` | Low word is the map-unit ID on the mapped producer path; on an Item, see [Item Token `+08`](#item-token-08) |
| `23` | u16 | `+18` | Recipient publication mask in located send paths |
| `25` | u32 | `+1c` | Item value/price; other classes' meanings Unknown |
| `29` | u32 | `this` | Object's saved-address identity key |
| `33` | u32 | `+14` | Saved-address reference, actor owner on the actor path |

Both archive arms transfer the stated widths. Coincident runtime offsets in
another object are not the same field. Token `+18` is also Unit `+18` because
Unit calls Token without adjusting `this`. The mask is not a class discriminator. — SAV-OBJ-014, SAV-ID-015, SAV-TOKEN-034,
SAV-636 (its no-named-reader-or-writer clause is partially retracted; the byte
identity stands), SAV-653, SAV-654, SAV-678, ITEM-VALUE-115,
SHOP-CONSUME-073

`+14`'s "actor owner" meaning is established for the actor (Unit/Human) path
by a spawn/placement-time assignment writer, not by construction: Token has
four own constructors, not two — two zero-constructors, a copy constructor
and an owner constructor — and every construction path traced for the
`SpellEffect`/`Effect` lineages (`SpellEffect`, `PointEffect`, `AreaEffect`,
`SpellTransport`, `Effect`, `Effect_DirectDamage`) reaches one of these four;
the two zero-constructors write the field to zero, the copy constructor
copies it from a source object, and the owner constructor writes it from an
argument. `Unit`'s own constructors reach the same shared layer and write no
`+0x14` of their own. A separate runtime writer, outside every construction
path, zeros an already-constructed object's own `+0x14` inside a
removal/detach routine — so the field is not exclusively a construction-time
default, and whether that runtime zero is the same "owner" meaning cleared
on removal or a liveness-adjacent meaning is Unknown. Over the full current
admitted save corpus, nonzero is observed only on `Effect` (18 of 192
instances, one save, an independently confirmed original) and resolves,
every instance, to the `Human` object that structurally contains that
`Effect` — reached through an equipped `Armor`, `Weapon` or `Shield` record
in that `Human`'s own body — a self-referential owning-actor back-pointer,
not a `Player` and not the separate `SpellEffect+0x3c` caster field. Four of
the other five classes show one saved instance each, all the constructed
zero (one of the four drawn from a ROM1 resave of a produced document rather
than an unmodified original); `SpellEffect` has no record in the admitted
corpus at all. SAV-1058 is partially retracted for claiming both SpellEffect
bases call the first Token constructor; its two Token `+14` zero stores
remain valid. — SAV-1058, SAV-1059, SAV-1060, MAGIC-219

## Item Token `+08`

On an Item, Token `+08` is a pending pickup-announcement flag. Pickup sets it to
1 and a merge ORs it into the retained Item. The compact item-record writer
publishes it as record bit `0x40` and then stores 0. The client's
carried-container arm posts a "Picked up" line for that bit. Both archive arms
copy the word literally, so a saved 1 survives LOAD until its first
publication, the only clearing writer located. — SAV-1114, SAV-1115

## PointEffect and SpellTransport cast construction

The selected PointEffect chain reaches Token `R0923`; the selected
SpellTransport chain reaches Token `R1009`. Both reach the three-field
Token reset and zero `+14`. Their direct constructors do not write Token
`+08` or `+0c`. The embedded-list construction at `+20` does not overlap
those fields. The shared SpellEffect layer writes type word `+0e=0`.
These are direct instruction facts, not universal first-SAVE values.
— SAV-1068

PointEffect copies target Position through `L05872`; SpellTransport copies
its Position argument through `R1010`. The native transport caller passes
caster Position. Both copy `+00..+05` and `+08..+0b`, leaving destination
Position `+06/+07` untouched. Token's saved-address key is the allocation's
address, not a constant field default. Allocator contents, alias/callback
writes and changes before the first SAVE remain Unknown. Direct admission
fills PointEffect `+0c` from the Spell; it does not fill the transport's own
`+0c`, and neither selected path fills its own `+08`. — SAV-1068, MAGIC-225

## Position

| Position offset | Type | Meaning |
|---:|---|---|
| `+00` | u8 | Cell X |
| `+01` | u8 | Cell Y |
| `+02` | u16 | Explicit packed cell `(cellY<<8)\|cellX` |
| `+04` | u8 | Sub-cell X; named constructors initialize `0x80` |
| `+05` | u8 | Sub-cell Y; named constructors initialize `0x80` |
| `+06` | u16 | Bytes left unmanaged by the named constructors/copies |
| `+08` | u32 | Caller-supplied terrain pointer/key |

Accessors return `fullX=256*cellX+subX`, `fullY=256*cellY+subY`.
The u16 spanning `+00/+01` and explicit `+02` encode the same cell in the
normal producer relation. The archive copies all 12 bytes, including bytes
that constructors/copies omit. LOAD overwrites the fresh current-terrain
pointer before terrain reconstruction. — SAV-TOKENPOS-074, SAV-TOKENPTR-075

No named constructor/copy/lifecycle path assigns a meaning to `+06/+07`.
Zero is a Medium-confidence authoring placeholder for those two bytes only;
a preserving reader retains them. Aliased/bulk copies and unwitnessed runtime
paths remain outside that conclusion. — SAV-TOKENPOS-074, SAV-TOKENLOAD-095

## Identity map

Token LOAD binds the saved key to the object just constructed in the map at
`[L00285]+0x88`. The trailing reference uses `R1371`: missing key
becomes null. Player's own key and Diary's trailing reference use the same
map at their named sites. A writer assigns a unique nonzero key per defined
object and references the target's key. Archive indices and map-unit IDs are
separate. The null-on-miss rule is narrowed to these named resolver sites;
Position, cell and order repair use the different policies below.
— SAV-TOKEN-034, SAV-PTRMAP-035, SAV-ARCHREL-253

Reference-repair policies differ by field:

| Field or family | Hit | Miss / zero |
|---|---|---|
| Token trailing `+14`, named Player/Diary/Group references | Replace with live pointer | Null at their named resolver |
| Position terrain `+08` | Replace with live terrain | Retain saved word |
| Ten cell-payload object fields | Replace with live pointer | Retain saved word |
| Stage-zero order keys and mover `+7c` | Replace with live pointer | Retain saved word |
| World-lifecycle Unit `+40/+64/+68` | Replace with live pointer | Null |

Use the exact call-site rule, not one generic fixup for all u32 fields.
Position's nonzero key must agree with the terrain record for the named
world-rebind route. Zero or a dummy pointer is not a general substitute.
An unresolved nonzero Position key remains unresolved; coordinates or record
order do not supply a replacement.
— SAV-IDCENSUS-254, SAV-TOKENLOAD-092, SAV-CELLLOAD-112, SAV-HUMRESUME-460
(its reference-repair shorthand is superseded by these field-specific rows),
SAV-GRPLOAD-560, SAV-847, SAV-908

## Reach of Position repair

World LOAD registers terrain, then dispatches live/dead actors and top-level
SpellEffects through their lifecycle hooks. Sacks load after the earlier
live-list construction and do not receive that pass. The complete loader has
no Building/loaded-Sack manager callback and establishes no nested Item/Effect
walk. City/no-world LOAD constructs no terrain and calls no such manager.
— SAV-TOKENLOAD-093, SAV-TOKENLOAD-094, SAV-CELLLOAD-112, SAV-DEADLOAD-125

Building, Outpost, Tavern, Shop and Sack serializers contain no direct Position
rebind/clamp call after their base body. Building/Sack tick bodies are empty.
Building placement/removal reads cell X/Y; Sack reads the packed cell, with
terrain supplied separately. Nested callbacks, later resave and actual first
Position use remain Unknown. — SAV-POSLOAD-140

## Session entry and later sends

The later session join walks Building and Sack managers. Both senders read
cell and sub-cell through full-coordinate getters. Sack sends full words;
Building shifts right eight and emits cell bytes. Neither sender uses packed
`+02`, unmanaged `+06` or terrain `+08`. Earlier Position use can precede
these mandatory-prefix consumers. — SAV-POSTLOAD-221 (its no-world corpus
count is superseded by SAV-SUFF-300; the sender result stands)

Before entry, nonzero campaign `+6b8` can call the ordinary wrapper
`R0147`; nonzero server `+2c` admits the same sub-tick used by normal pacing.
It increments the counter, drains commands, dispatches world entries and the
global tick list. The zero campaign arm executes no tick. Queued/computed
callbacks can precede the later sender even though Building/Sack `vt+18`
bodies are no-ops. — SAV-FIRSTTICK-348, SAV-FIRSTTICK-349,
SAV-FIRSTTICK-350

Actor entry mask `-1` can admit equipment/container traversal, subject to
class, recipient ownership and type gates. An Item's Effects require nonzero
appearance `+40`; plain Item `+1c==-1` bypasses the common Effect walk.
Automatic Sack entry does not visit its container. Pickup drains quantity-one
Items, can deep-copy Effects on a split, deletes the source container, nulls
Sack `+40`, destroys the Sack and requests another gated actor send.
— SAV-POSTLOAD-222, SAV-POSTLOAD-223

## Recipient publication mask

| Operation | Token `+18` rule |
|---|---|
| `L08643` | Return intersection with Player `+2c` |
| `R0676` | Return 1 if that intersection is zero |
| `L08644` | OR recipient bits into the field |
| `R1624` reset | Zero `+18`, `+1c` and `+4` |
| Located revoke sites | AND with complement of recipient bits |

One revoke site immediately republishes. Other send paths use a dedup test,
a positive intersection test or inline OR; a null recipient can publish to
every bit. Zero therefore records current revoked/unpublished state, not
proof that publication never happened. Terminal stage does not determine the
mask; revoke and republish can change it during the actor lifetime.
Universal fog-of-war meaning and the complete class lifetime remain Unknown.
— SAV-678, SAV-653, SAV-664
