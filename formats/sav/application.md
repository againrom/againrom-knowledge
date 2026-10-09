# SAV label and application state

[Format reference](format.md) · [Campaign](campaign.md) · [REG encoding](../reg/format.md)

## Physical tail

Let `bodyEnd=16+blobBytes`, the first file byte after the compressed blob.
The tail is uncompressed and has three consecutive parts.

| File range | Contents | Endpoint rule |
|---|---|---|
| `[bodyEnd,bodyEnd+0x100)` | Label buffer | Exactly 256 bytes |
| `[bodyEnd+0x100,storeEnd)` | `&YA1` inline key/value store | `storeEnd=bodyEnd+0x100+0x18+32*R+4+poolLen` |
| `[storeEnd,EOF)` | Campaign record | Counted [campaign grammar](campaign.md#wire-programme) |

`R` is the store's record count. REG header, kind bits, records and pool
encoding apply directly; state entries have no CArchive tags. — SAV-EXT-009,
SAV-EMB-004 (its run-to-EOF extent clause is partially retracted by
SAV-TAILEXT-062), SAV-TAILEXT-062, SAV-CAMPTAIL-070, REG-FMT-017, REG-099

Both located save dialogs read the label's fixed buffer, but the
selection-to-writer path passes only the NUL-terminated string. Bytes after
the first NUL can be stale destination-buffer residue on fresh SAVE; they
have no source-preservation guarantee. — SAV-LABELTAIL-236

At the one SAVE dialog construction site read, label buffer offset 0 is
explicitly reset to an empty string before construction; whether every
construction path does this was not enumerated. — SAV-SAVELABEL-1016

The SAVE dialog's in-memory staging area for this buffer is twice the
on-disk buffer's width; only the first half reaches disk, matching the
256-byte label in the physical tail above. — SAV-SAVELABEL-1017

A lawful save in the preserved owner corpus carries a byte above 0x7F as
live, pre-terminator label content; a reader must not assume the label
buffer holds only printable ASCII. — SAV-SAVELABEL-1018

## State roots and production

The baseline world producer has nine roots:
`Character`, `CurrentState`, `Fog`, `GameOptions`, `Inventory`, `Objects`,
`Projectiles`, `SpellBook`, `View`. It can additionally emit live
`Objects/Group0` through `Group9` and decimal top-level `Prj<id>` roots.
City/no-world records omit Fog and Projectiles on the located producer route.
Record/leaf counts therefore depend on state, not on one fixed schema count.
— SAV-EMB-004 (extent clause partially retracted; key names stand),
SAV-TAILEXT-062, SAV-PROJSTORE-428, SAV-915

The application consumers load stack-local registries, then read the campaign
from the same file, consume named state and destroy the registry. The longer
consumer additionally restores Fog and Projectiles. They do not enumerate
all unknown entries. Ordinary SAVE builds a fresh empty registry from current
application/UI/world values and writes it before the campaign; it does not
pass the loaded registry through. — SAV-914, SAV-915

## SpellBook state

| Key | Wire kind / bytes | Meaning / runtime source |
|---|---|---|
| `SpellBook/Shortcuts` | Kind 6, 16-byte integer array | Four signed zero-based spell indices, F5/F6/F7/F8; controller `+64,+68,+6c,+70` |
| `SpellBook/Pressed` | Integer | Current spell, controller `+60` |
| `SpellBook/IsOpen` | Integer | Visibility |

The controller is reached through campaign `+ec`. Shortcut `-1` is unbound.
Restore copies four dwords without checking array count or individual indices.
Missing/empty returns can reach a null read; arrays shorter than four elements
do not cover the four reads. The canonical writer supplies exactly 16 bytes.
A fifth element covers the reads but is not the canonical writer shape.
— SAV-ORIGFAULT-335, AI-QUICKSAVE-281, SAV-916

The shortcut wire meanings are established. Original reload of populated
bindings and later malformed-index behavior remain Unknown; restore has no
local per-index check. — AI-QUICKSAVE-281, SAV-916

## Fog

| Key | Type | Meaning |
|---|---|---|
| `Fog/FirstState` | int32 | State carried by the first run |
| `Fog/Data` | int32[] | Alternating-state run lengths over `W*H` tile words |

The serialized bit is tile-plane bit 15, explored at least once. Bit 14 is
not saved. The scan order is `index=col+row*W`; run lengths sum to `W*H`.
— SAV-FOG-061, TERR-TILE-079 (amended, superseded), ANIM-TICK-011, TERR-EDGE-024

Read sequence:

1. Construct the ALM tile plane, whose authored cells have bit 15 clear.
2. Start in `FirstState` and consume each run in row-major tile order.
3. Set bit 15 on cells in the set state and flip state at each run boundary.

Write runs where `tile[i] & 0x8000` changes, storing the first state separately.
The simulation document contains no full tile plane; the application tail
supplies explored state after terrain reconstruction. The renderer's OR of
four corner words can draw a different extent from the set cells;
that rendering extent is not established by SAV framing.
— SAV-FOG-061, TERR-FOG-087 (partially retracted), TERR-FOG-145, SAV-LOAD-057

## Projectiles

World SAVE always emits the Projectiles root, even for an empty manager.

| Key | Kind | Value source |
|---|---:|---|
| `Projectiles/FreeIndex` | 2 | Manager's u16 allocator field widened to integer |
| `Projectiles/IDs` | 6 | Every live node's u16 ID widened to a four-byte array element |

For every ID it also emits `Prj<decimal id>` with sixteen kind-2 leaves in
this order:

```text
x y z picture dir phase lastaction action actiondir actiontarget
actionx actiony actionz actionphase actionsegments actionspell
```

An empty manager still emits both leaves. No rule requires
`FreeIndex=max(ID)+1`; preserve the allocator field separately from the IDs.
— SAV-PROJSTORE-428, SAV-PROJCORP-430

The writer sources `x/y` from current projectile `+08/+0c`. It emits no
dedicated launch-point leaf from cached `+28/+2c`. A moving projectile can
therefore save a point different from its origin; SAVE does not freeze the
original cast point in a second field. Native in-flight SAVE/LOAD and
post-load cache restoration remain unobserved. — MAGIC-267

LOAD is less restrictive than the producer. Missing FreeIndex defaults to 0;
missing IDs leaves the constructed vector empty. Kind 6 yields `size/4`
elements, taking each low u16; kind 2 has a one-element compatibility arm.
Other kinds reach the typed-getter error. Each listed ID causes allocation,
sixteen defaulted leaf reads, terrain/world/ID binding, insertion and a
post-load helper. The tree is producer-complete but not loader-mandatory.
— SAV-PROJLOAD-429

LOAD stores the low u16 of FreeIndex into the live counter before it builds
the records, and SAVE writes that counter back, so a round trip keeps it and
the next shot takes ID FreeIndex. — SAV-1191

The counter's other writers are the client-world constructor (0) and six
record insertions (old + 1, wrapping at 65536). No found writer resets it at
a mission entry or a collect-all clear; the client-world constructor is
reached only from one virtual method.
— SAV-1189, SAV-1190

The sixteen leaves are the whole saved record. The projectile trail, which the
driver fills for pictures 10 and 12, is not saved and starts empty after LOAD. The start point is not saved
either: `x/y` is the current point. A dead shooter's pending damage is not part
of the projectile and is never applied; the record itself keeps flying. —
SAV-1193, SAV-1192, SAV-1133

## Unknown entries and extension survival

The raw registry loader copies framed records/pool bytes without name
validation. With correct ordering, noncolliding unrelated names do not alter
the selected known lookups. Flag bit 4 selects case-sensitive bsearch; without
it the 15-byte linear comparison folds ASCII case under the no-locale condition.
Duplicates, false sorted flags and incompatible kinds can change lookup.
— REG-099, REG-100, REG-101

A raw-registry write can preserve unrelated entries. Ordinary application
SAVE reconstructs its registry from named/live producers; loss of entries
that no producer requests is a Medium-confidence persistence inference.
A complete ordinary native LOAD/SAVE cycle and hidden helper effects remain
Unknown. — REG-102, SAV-914, SAV-915

| Candidate extension location | Located original read/write behavior |
|---|---|
| Label after first NUL | Read as buffer, forwarded as string; fresh buffer residue has no source provenance |
| After campaign endpoint | Located loaders perform no further read/EOF check; fresh writer truncates and has no suffix-copy path |
| Decoded alignment byte | Zero/one byte selected by parity; no source-pad copy |

The located paths supply no general preservation contract for extension data
in these three locations. Other load/save routes, ignored fields, internal
raw residue, runtime serializers and multiplayer shapes remain unspecified.
— SAV-LABELTAIL-236, SAV-PHYSUFFIX-237, SAV-DECPAD-238, SAV-EXTSURV-239
