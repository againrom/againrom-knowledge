# SAV object programmes

[Format reference](format.md) · [Encoding](encoding.md)

Each row is an ordered body programme after its archive tag. `Base` means a
direct call to that base serializer; it adds no tag. Widths exclude the tag
and any newly introduced class descriptor. Offsets name the runtime source.
— SAV-MEMBER-036 (its unread-classes Unknown superseded by SAV-EMBED-039),
SAV-CLASSSER-172, SAV-CONTSER-188

## Token-derived bodies

| Class | Ordered body after its base |
|---|---|
| `Token` | [37-byte Token head](token.md#token-head) |
| `VirtualCaster` | Token; u8 `+3c`; raw 6 from buffer pointed to by `+40` |
| `Unit` | [Unit programme](actors.md#unit-programme), beginning with Token |
| `Humanoid` | Unit; raw 24 from `+1cc`; `objref` at `+198+4*i`, `i=1..12`; `objref` at `+1e4` |
| `Human` | Humanoid; own store arm adds no bytes |
| `Effect` | Token; u8 `+3c`; u8 `+3d`; u32 `+40`; u8 `+0c` |
| `Effect_DirectDamage` | Effect; raw 24 from `+48` (shares its shape with the [attack block](actors.md#unit-programme); only `+13/+14/+15` have a located writer here) |
| `SpellEffect` | Token; u8 `+40`; u8 `+41` |
| `PointEffect` | SpellEffect; `objref<Effect>` at `+48`; raw u32 `+44` |
| `AreaEffect` | SpellEffect; u8 `+48,+49,+4a,+4b`; u16 `+4c`; `objref<Effect>` at `+44` |
| `SpellTransport` | SpellEffect; `objref<SpellEffect>` at `+44`; `objref<AreaEffect>` at `+48`; u16 `+4c` |
| `Item` | Token; `list32<Effect>` at `+20`; u16 `+40,+42`; u8 `+44,+45,+46`; u16 `+48,+4a`; u8 `+47` |
| `Shield` | Item; raw 22 from `+50` (its own store arm is empty; this transfer is inherited/shared) |
| `Armor` | Item; raw 22 from `+52`; u8 `+50` |
| `Weapon` | Item; raw 24 from `+52`; raw 22 from `+6a`; u8 `+50`; `objref` at `+80` |
| `Building` | Token; raw 22 from `+52`; u8 `+40`; u16 `+42,+44,+46`; u8 `+48,+60,+61`; u32 `+64,+68` |
| `Outpost` | Building; u32 `+84,+88,+80,+8c`; `Count(n)`; raw `8*n` from embedded `+6c` |
| `Tavern` | Building; u32 `+9c` |
| `Shop` | Building; u32 `+70`, the value cap |
| `Sack` | Token; u32 `+3c`; direct [container](items.md#container) at `*(+40)` |

Sack's gold scalar is current `+3c`, the same source its pickup body reads.
Its named serializer calls Token serialization, writes/reads four gold bytes
and serializes the existing container. This direct join does not establish a
native post-LOAD pickup or a queue-time SAVE result (MENU-095).

These sequences are the schema-1 `.data` Token lineage. PointEffect LOAD
immediately remaps its trailing raw `+44` key: hit installs the pointer;
miss writes null. No extra bytes are consumed. Outpost's array has only the
`CObject` runtime descriptor; its element meanings remain Unknown.

`SpellEffect`'s own row, and the three rows built on it, list every field
that class family's own `Serialize` touches beyond `Token`'s own head. The
shared base these rows reach first, `Token::Serialize`, touches only
Token-head fields; it never touches `+3c`. — SAV-CLASSSER-173,
SAV-CLASSSER-174, SAV-CLASSSER-175, SAV-1054

For the selected PointEffect and SpellTransport cast constructors, the common
SpellEffect bytes are explicit stores: `+40=0` and `+41=1` in both reached
base variants. PointEffect's direct caller replaces `+41` with 1 exactly when
cached Spell `+0a` is zero; the transport caller does not replace its own
common bytes. A child and its transport therefore need not carry the same
`+41`. These constructor/admission sources do not establish first-SAVE
invariants: allocation/alias boundaries and later writes remain Unknown.
— SAV-1069, MAGIC-ATTRGATE-118, MAGIC-225

Direct delivery-2 admission supplies the caster's Position `+10` to
SpellTransport's Position-copy constructor chain. Its inspected tick body
does not move that Position. The nested PointEffect separately copies the
target Position. The transport's caster-origin Position persists as the
12-byte Token prefix at body offset 0; it is separate from the visible
projectile's class-dependent start. Native ordering, aliases and SAVE/LOAD
continuation remain Unknown. — MAGIC-266

`Effect_DirectDamage`'s raw 24 bytes at `+48` are not a separate shape: the
same constructor family that builds the [attack block](actors.md#unit-programme)
at live `+a6` and base `+114`, and the same resolver pair that reads the
live `+a6` copy, reach this span too. The positions those resolvers read
(`+00`, `+0e`..`+15`) carry that block's own `toHit`/`dmgBase`/`dmgSpread`/
`active` names, structurally present but with no located writer here; the
six skill `u16` at `+02`..`+0d` are named from the shared constructor alone,
and no located reader touches them at either embedding. Only `+13/+14/+15`
(damage minimum, spread, school) have a located writer for this class, by
five near-identical setters dispatched from the spell-apply routine; LOAD
copies the raw 24 bytes as one transfer and does not re-run them, so a
reloaded effect's damage triple is frozen at save time regardless of which
construction path built the object. `+16/+17` have no located writer beyond
the shared constructor's own partial zero-init and no located reader at
all. The class's descriptor creator field (`SAV-TOKENLOAD-093`) is a bare
wrapper around the same constructor a live cast uses directly. The archive's
own dispatch mechanism that reads a class's creator field during LOAD is
traced generically for the `SpellEffect` lineage — `CArchive::ReadObject`
resolves the class descriptor and calls through its own `+0xC` field
(`SAV-1055`) — but `SAV-1055`'s own raw descriptor-dword read does not cover
this class's own descriptor, so that generic mechanism is corroborated for
`Effect_DirectDamage` only by `SAV-TOKENLOAD-093`'s own separate reading,
named above. One
corpus-witnessed record (n=1 of 31 SHA-distinct preserved saves) agrees:
`+00`/`+0e..+12` are `0x00`, `+13/+14/+15` are `0x04/0x04/0x01`, and
`+16/+17` hold nonzero save-time heap residue, `0xc6/0x02`.
— SAV-MEMBER-036 (its unread-classes Unknown superseded by SAV-EMBED-039),
SAV-HUMAN-043, SAV-CLASSSER-172, SAV-CLASSSER-173, SAV-CLASSSER-174,
SAV-CLASSSER-175, SAV-CLASSSER-176, SHOP-SAVE-015, SAV-1032, SAV-1033, SAV-1034

## Fixed extents and variable extents

| Class | Body bytes |
|---|---:|
| Spell | 9 |
| Token | 37 |
| SpellEffect | 39 |
| VirtualCaster, Effect | 44 |
| Effect_DirectDamage | 68 |
| Building | 77 |
| Tavern, Shop | 81 |
| PointEffect | At least 45 |
| AreaEffect | At least 47 |
| SpellTransport | At least 45 |
| Outpost | `95+8*n`, or `99+8*n` with a wide Count |

The minima for reference-bearing records use null references. A reference can
instead introduce a class and nested body. Item is 53 bytes **including** its
u32 Effect-list count, plus the encoded Effect reference entries:
`37 Token + 4 count + 12 scalar bytes`. Unit has a 603-byte fixed part and is
at least `609+L` with its name of length `L` and three null references;
Human adds at least 24+26 bytes.
Counts, strings, presence flags and nested references move the endpoint.
— SAV-MEMBER-036 (its `Human` clause amended by SAV-HUMAN-043's thirteen
references; its unread-classes Unknown superseded by SAV-EMBED-039),
SAV-HUMAN-043, SAV-UNITLEN-045, SAV-CLASSSER-173, SAV-CLASSSER-176

Building's Token map-unit ID joins the authored type-4 record. Its 77-byte body
is fixed. — SAV-BLDG-037

A structure destroyed in play keeps its Building record. Word `+0x42` is 0, the
maximum `+0x44` and the other fields are unchanged, and the record count equals
the map's authored type-4 count. The footprint cell keeps its Building key and
block rows in the one observed case. In the 7,232 records and the one destroyed
structure's record and cell, no ruin flag appeared; no whole-document search for
another marker was run. An authored placement word of zero also stores HP 0, so the
type-4 word is needed to tell the two apart. No negative HP was observed in one
destruction event, which cannot tell the unclamped physical writer from the clamped
consumer. That LOAD restores the HP word is inferred from the serializer, not observed. — SAV-1163, SAV-1164, SAV-1165

The health word that the Building publication arm sends is word `+0x42`, or 1000
when `+0x44` is zero; the arm skips a Building whose word `+0x0e` is zero. Whether
LOAD invokes an introduction is Unknown. — SAV-1166

## Non-Token schema-1 bodies

| Class | Ordered body |
|---|---|
| `Player` | [Player programme](player.md#wire-fields) |
| `Diary` | `Count(n)`, raw `4*n`; `Count(m)`, raw `2*m`; u32 owner key |
| `Spell` | u8 `+08,+09,+0a`; u16 `+0c`; u32 `this` identity key |
| `Spellbook` | u32 `+18`; u32 `n` from array `+04`; `n-1` references for indices `1..n-1` |
| `TableLine` | CString at `+4`; direct CDWordArray at `+8` |
| `CMultiShopShelf` | Empty |
| `CMultiShopInstance` | Empty |
| `CMultiShopTemplate` | Empty |

Spellbook index zero is omitted in both directions. Each slot has its own
reference operation; new Spell objects and aliases follow the intended graph.
— SAV-DIARY-042, SAV-SPELL-044, SAV-SPELLBK-041, SAV-CONTSER-188

TableLine's grammar is for the descriptor-created base object. Other vtables
advertise the same descriptor and some use different serializers; those
unregistered derived programmes are not covered by this row. The three
CMultiShop slot-2 bodies are `RET 4`. Shop's own programme excludes stock.
— SAV-CONTSER-189, SAV-CONTSER-193, SHOP-SAVE-015

## Schema-0 bodies

| Class | Ordered body |
|---|---|
| `CStringArray` | `Count(n)`, `n` CStrings |
| `CDWordArray` | `Count(n)`, raw `4*n` |
| `CWordArray` | `Count(n)`, raw `2*n` |
| `CByteArray` | `Count(n)`, raw `n` |
| `CMapStringToString` | `Count(n)`, `n` pairs of CString key and CString value |
| `CMapStringToOb` | `Count(n)`, `n` pairs of CString key and `objref<CObject>` value |
| `CDib` | BMP file header[14], info/palette[`bfOffBits-14`], pixels[`imageBytes`] |

Both maps walk current bucket chains without sorting. CMapStringToOb values
participate in the shared archive index. Diary directly dispatches its two
arrays and therefore emits no class tag for them at those sites.
— SAV-CONTSER-189, SAV-CONTSER-190, SAV-CONTSER-191, SAV-CONTSER-192,
SAV-DIARY-042

CDib flushes the archive and writes through the underlying file. Store writes
`BM`, `bfOffBits=0x36+4*paletteCount`, `bfSize=bfOffBits+imageBytes`, a 40-byte
info header, palette and pixels. LOAD requires `BM`, obtains the middle size
from `bfOffBits`, and uses nonzero `biSizeImage` or the padded-row formula for
pixels. It does not validate `bfSize`. — SAV-CONTSER-192

## Embedded collections

| Source | Body | Meaning |
|---|---|---|
| Unit `+15c` | `Count(n)`, `n*u16` | Static packed-cell route |
| Unit `+178` | Same | Dynamic packed-cell route |
| `*(*(Unit+158)+90)` | Same | Actor order's own patrol path |
| Group `+20` | Same | Meaning Unknown |
| `*(*(Group+3c)+4c)` | Same, after raw AI80 | Group AI path |
| `*(Unit+140)` | Spellbook, presence-gated | Known spells |
| Diary `+04/+18` | CDWordArray / CWordArray | [Diary](diary.md) |

The unnamed word-list class uses vtable `L08172`, constructor `R1545`
and serializer `R1546`. Each saved element is node `+8`; nodes themselves
are not saved. Runtime append grows the pool by ten nodes and teardown frees
it. LOAD uses that same append path. The Unit route elements are packed cells
`(y<<8)|x`; actor-order and Group paths have separate identities.
— SAV-EMBED-039, SAV-WLIST-040, MOVE-ROUTE-004, SAV-630, SAV-631, SAV-632, SAV-633, SAV-634,
SAV-GRPPATROL-570

## Programme scope

The tables cover the known creator-backed schema-1 game and schema-0 runtime
programmes. Schema `0xffff` descriptors with null creators are not
stream-creatable. Runtime-built, copied and aliased descriptors, plus an
unregistered derived object advertising TableLine's descriptor, remain outside
these programmes. — SAV-CLASS-033, SAV-SERPOP-047, SAV-READPOP-255,
SAV-CONTSER-189

A base programme can execute within a derived body without its own class tag.
An embedded programme can likewise execute without a descriptor. Absence of a
class name from a save does not make that programme unreachable.
— SAV-PRODDIRECT-207

## Saved cast continuation

SpellTransport's final u16 at runtime+4c is its remaining delivery counter.
The serializer preserves it beside the two typed children. Its own post-load
method calls the Position rebind and dispatches each child's own post-load
hook, without resetting the counter.
Tick subtracts1 and, when the signed result is<=0, enqueues the child, clears
both child fields and marks the transport retired. Natural game0018 contains
counter4 in its pending PointEffect/DirectDamage transport. Conditional isolated
store/load/tick execution agrees; native full-process resumption and first-frame
scheduler ordering remain Unknown. — SAV-CASTCONT-1006

The two typed children are bound during LOAD to a freshly resolved archive
reference, not a raw copy of an on-disk value: the serializer's own LOAD arm
calls the same typed-reference resolver used for every other archive
reference in this format, once per field, with each field's own class
descriptor. The post-load method's own child repair is entirely
the child's own doing — it null-guards each field and, only when present,
dispatches that child's own post-load hook; it does not itself write either
field. Both fields are therefore set exactly once, during the LOAD arm, and
never rewritten afterward on this path. — SAV-1042

The class remains a single-digest corpus witness (`game0018.sav`) after an
independent re-run of the population `SAV-1037` already searched (129 paths,
76 SHA-distinct); no writer path in that population emits a second one. The
one witnessed record is resolved into a live object by the archive's own
per-element load loop for the container holding it, through the identical
typed-reference mechanism named above. — SAV-1043

A SpellTransport, and the child it hands off on delivery, are held alive by
the same single container every directly-cast PointEffect/AreaEffect uses —
not a transport-private list — confirmed by instruction comparison at all of
Tick's own registrar calls, not assumed from the shared function name.
— MAGIC-197

Over a corpus more than double the size searched when the above was
established (129Asg& save paths / 76 SHA-distinct, against the prior 55/31,
a population that includes at least a dozen documented or self-named
project-produced fixtures, none of which contributes a hit), the witness
count for SpellTransport, PointEffect and Effect_DirectDamage is unchanged:
one distinct digest, the same `game0018.sav`. n=1 stays n=1. AreaEffect,
previously a zero-witness class, now has one: a record in a separate save,
ROM1's own output but written from a document this project had modified
before ROM1 loaded it. Neither the record's own graph position (a direct
cast, or a nested reference reachable through `SpellTransport+0x48`'s own
`objref<AreaEffect>` above) nor its fields were decoded; not graded.
— SAV-1037

For a load-time binding to resolve at all, a writer must have put something
on the wire for it. The generic archive object writer takes a null pointer
and a real object pointer down two different paths: null becomes a bare
zero-value tag with no class information at all; a real object either
becomes a compact back-reference index, if this exact pointer already
occurred anywhere earlier in the archive, or, on its own first occurrence,
the object's own class descriptor followed by a recursive dispatch into that
object's own Serialize — but only when that class descriptor has not itself
already been registered by an earlier object or class write anywhere in the
archive. Objects and classes share one running index counter, not two:
class registration (`CArchive::WriteClass`) reads, stores into and
post-increments the identical counter object registration
(`CArchive::WriteObject`) uses, so a class already written once, by any
object, is thereafter referenced by every later object of that class as the
same compact back-reference index a repeated object pointer gets, never a
second name write; the LOAD side mirrors this in the identical shape for
both `ReadClass` and `ReadObject`. The same generic mechanism, not a
container-specific or class-specific one, serves both the container's own
per-element writes and a SpellTransport's own two child writes. Within the
five writer/reader bodies this mechanism was traced through, a null
child and a never-otherwise-assigned back-reference index are the identical
wire value, by construction of the running index counter's own starting
point; no other code that writes into that counter, or might reset it
mid-archive, was searched for, and whether the allocator itself guarantees
a never-indexed slot's zero value, versus every traced write path merely
producing that result, is Unknown. Every write call therefore writes
something; nothing is silently omitted for a null child. — SAV-1047

SpellTransport's own STORE arm calls this generic writer unconditionally on
both typed child fields, with no null test anywhere in the arm; whichever
child is absent at STORE time, the writer still runs and still writes the
null case above, not nothing. The class hierarchy a load-time typed
reference depends on (SpellTransport is a SpellEffect, one hop) is
confirmed by a direct read of the class descriptor chain, independent of
the base-Serialize-call route used elsewhere in this section. — SAV-1048

Neither PointEffect's own post-load hook nor AreaEffect's own contains any
call to the shared list's registrar or its append primitive, in either
class's own complete body. This extends the SpellTransport finding above —
LOAD-time field binding is a one-time event, and the post-load method's own
repair is entirely the child's own doing — to the other two classes in this
family: no post-load hook anywhere in this family re-registers a
doubly-referenced child into the shared container, each hook only
null-guards its own class-specific field and, when present, dispatches that
field's own post-load hook once. Across all three classes, LOAD-time field
binding and LOAD-time container membership are each set exactly once, and
no post-load hook revisits which owns which. — SAV-1050
