# SAV Diary

[Format reference](format.md) · [Player](player.md) · [Actors](actors.md)

## Wire fields

| Order | Wire | Runtime member |
|---:|---|---|
| 1 | `Count(n)`, raw `4*n` | Embedded CDWordArray at `+04` |
| 2 | `Count(m)`, raw `2*m` | Embedded CWordArray at `+18` |
| 3 | u32 saved-address reference | `+2c` |

Counts are independent. LOAD sizes each array from its own count, reads its
payload, then resolves `+2c` through the saved-address map: hit installs the
pointer, miss/zero becomes null. It enforces neither equal lengths nor a
complementary element relation. — SAV-DIARY-042, SAV-PLDIARY-054, SAV-847

A Player owns one Diary at `+40`, dispatched directly in its suffix. A
Humanoid/Human can also reference a separate Diary at actor `+1e4` through
CArchive. Actor-owned Diaries default to a null owner; the two ownership
routes are not interchangeable. — SAV-HUMAN-043, SAV-846

The actor member's typed LOAD can return an existing archive object or a new
factory result. The new object enters the archive table before virtual
serialization; the type test permits a compatible descendant. The member
does not imply a unique allocation or exclusive ownership. — SAV-984

## Initialization and index domain

A fresh Diary sizes both arrays from the Units collection count and initializes
dword 0 and a word of 1024. That capacity does not establish a single actor index
domain: base Unit construction writes a Units ordinal into actor `+0c`,
while Human construction writes a Humans ordinal. The mutation rejects an
index above 63 only when actor virtual `+30` is nonzero: false for Unit,
true for Humanoid/Human. — SAV-667, SAV-668 (its unconditional index bound
and per-unit-type reading are superseded by SAV-844), SAV-844

The array accessors perform pointer arithmetic without a length check. No
safe malformed count/index relationship follows from the field widths.
— SAV-842

## Mutation and notification

The located mutation first decrements a nonzero word, then increments the
dword only while its unsigned old value is at most 16. From defaults,
17 admitted calls yield `(17,1007)` and 18 yield `(17,1006)`. Thus
`word=1024-dword` is not an invariant. LOAD must preserve both arrays independently. — SAV-842, SAV-847

The located caller is actor teardown under removed actor `T+54==16`.
With `S=T+40`, it requires S type word `+0e` in 33..63 and signed health
`+94>=0`. The receiver is `S->Player(+14)->Diary(+40)` and argument is T.
Standard damage can record its source in T `+40` before death; this does not
establish a universal lethal-blow/every-kill interpretation.
— SAV-843

That Diary selection occurs after the source's `+48` and `+60` callbacks.
The manager rereads the current victim `+40`, then source `+14` and Player
`+40`; it does not use the separate actor-owned Diary. On an admitted prefix,
the progress callback can refuse for its victim-owner gate while the later
Player counter and Diary operation still run. Distinct source/Player/Diary
controls execute those actual stores, with arithmetic and notification
returns supplied. Their joined history is Medium. — SAV-960

Earlier removal/Effect callbacks and later progress/notification effects
remain lifetime boundaries. The measured clear applies to the removed
victim's own attached Effect sources, not every reference to that actor.
First native post-LOAD attribution and actual source destruction remain
Unknown. — SAV-962

Non-null Diary `+2c` admits notification when the dword increments to
2,4,6,8,10,12,14,16. The downstream packet builder reads the supplied Player's
own array and writes 17 words under opcode 186:

1. For `i>=64`, obtain `typeID` and `face` from Units row i.
2. For `typeID=64..80`, select output word `typeID - 64`.
3. OR `min(lowByte(dword[i]>>1),7) << (4*(face-1))` into that word.
4. If `Player+68>10`, select all `0xffff` instead.

Null owner suppresses notification without suppressing mutation. — SAV-845

## Client receipt and local refresh

The Player ID selects a transport connection before serialization. The
count-word packet serializes only opcode, dword count plus the words: 39 bytes
for 17 words. Its concrete receiver reads the count, then the words, into a static
packet. Queue/refill/flush and native delivery remain separate boundaries.
— SAV-994

Opcode186 replaces the receiving client object's CWordArray at `+3f58`.
It sizes from packet count and copies words; it performs no second Player
lookup. The client constructor creates this array, and the frontend's
`+d0` pointer supplies the dispatch receiver. This client array is separate
from both serialized Diary arrays. — SAV-995

The shared character/unit panel reads the cached word at
`drawable+20 - 64`, when that signed index is in range. Its `drawable+e0`
pointer supplies the client object. The byte at drawable `+24` selects a
nibble by an x86 shift; the result replaces the earlier visibility fallback.
A separate override forces7. Drawing and pointer-position text gates consume
the selector. Nibble15 remains 15; some drawing gates require exact7.
These are conditional client reads, not proof of every drawable's binding
or a native redraw schedule. — SAV-996

The panel reads the selector as one unclamped nibble level L: 0..7 from the Diary path, 15 from the 0xffff send. Health and Mana
need L>=1, Sight and Speed 2, attack and damage 3, armour and defence 4,
primary statistics 5, resistances 6, skills 7. The local player's own units
read 7 without the table, and a Human or hero reads no table nibble.
— UNIT-145, UNIT-146, UNIT-147

The client array is not serialized: SAV carries the two Diary arrays and the
original rebuilds the array from opcode186. The four mission-151 owner saves
hold all-zero Diaries, so the level the owner reported on enemy Humans did not
come from a Diary count. — SAV-1172, SAV-1173

Normal-return setup with nonnull Player `+34` calls the builder with that
Player, independently of the setup argument's optional initialization arm.
This extends the conditional opcode4 resume relation. Later attributed
events retain the Diary mutation's existing notification gate. The other
located builder call belongs to a text-command prefix branch, so it does
not establish an additional ordinary LOAD/event producer. — SAV-997

The opcode186 arm directly resizes/copies; it does not directly draw, post
a message or invalidate a panel. Native delivery, static-packet interleaving,
callback preservation, first-after-LOAD refresh and visible delay remain
Unknown. — SAV-998

## Remaining scope

The bounded actor receiver search reaches member SAVE and virtual destruction,
but no ordinary actor-owned array consumer. It includes 85 complete bodies
and the 56 Humanoid/Human slots through `+6c`. Whole-actor callees, callbacks
and arbitrary aliases remain outside transitive closure. Shared accessor
calls need not receive a Diary. — SAV-854, SAV-855, SAV-982

The progress expression `actor+1cc+4*k` overlaps the Diary pointer word when
`k=6`; its two branches do not locally impose an upper index of 5. This is
conditional scalar access to the pointer word, with native index reach and
pointer change unobserved. It is not a read of either Diary array. The
bounded negative therefore supplies no pointer-preservation rule. — SAV-983

Indirect/computed/inlined access and unclassified receivers remain open.
The word's gameplay meaning, full event attribution and the first post-LOAD
consumer are Unknown. The wire grammar requires neither an owner self-reference
nor equal/complementary arrays. Preserve both arrays and the archive relation;
the local copy route supplies no normal actor consumer. — SAV-843, SAV-847,
SAV-982, SAV-984
