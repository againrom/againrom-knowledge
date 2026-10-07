# SAV encoding

[Format reference](format.md)

## Word codec

The compressed blob is `u32 outWords` followed by run/literal packets. A word
is two bytes in original order. The decoded allocation is `2 * outWords` bytes.
— SAV-FRAME-021, SAV-CODEC-022

| Opcode byte `n` | Following bytes | Output |
|---|---|---|
| `0x00..0x7f` | `2*n` | Copy `n` literal words |
| `0x80..0xff` | 2 | Repeat that word `n & 0x7f` times |

Read packets until the compressed source span ends. There is no end marker.
The original loop is source-bounded; exact output-size agreement is a
structural check. `0x00` is a zero-word literal, `0x7f` a 127-word literal,
`0x80` a zero-word run and `0x81` a one-word run. All are legal decoder arms.
Do not swap the
two bytes of a repeated word. — SAV-PACK-007, SAV-EXT-009, SAV-CODEC-022

The original encoder chooses a run when the next two words are equal and a
literal otherwise. Pad the logical input to even length first and write its
word count before the packets. Its own buffer is only `2 * inWords` bytes,
with no expansion allowance. — SAV-CODEC-022

## Counts and strings

`Count(n)` is the generic MFC collection count. `list32<T>` instead uses an
unconditional u32 count followed by that many typed `objref<T>` operations.
Player, dead-actor, Building, SpellEffect and Sack lists use `list32`; Group
counts also use plain u32. — SAV-WLIST-040, SAV-DOC-053 (partially retracted
for its 7-of-18, Unit-failure and terminal-padding clauses; the list shape
stands)

| Primitive | Minimal length prefix | Payload |
|---|---|---|
| `Count`, `n < 0xffff` | `u16 n` | Determined by the collection |
| `Count`, `n >= 0xffff` | `u16 0xffff`, `u32 n` | Determined by the collection |
| ANSI `CString`, `n < 0xff` | `u8 n` | `n` string bytes |
| ANSI `CString`, `0xff <= n < 0xfffe` | `u8 0xff`, `u16 n` | `n` string bytes |
| ANSI `CString`, `n >= 0xfffe` | `u8 0xff`, `u16 0xffff`, `u32 n` | `n` string bytes |

`ff fe ff` is the CString reader's Unicode-mode sentinel, not length 65534.
The Unicode arm then reads a length and two bytes per character. Minimal ANSI
`CString("10.alm")` is `06 31 30 2e 61 6c 6d`; no NUL is part of that payload.
The [campaign marker string](campaign.md#markers) has its own grammar.
— SAV-MAP-005, SAV-FULLREAD-252, SAV-758

Readers admit nonminimal encodings: `ff ff 01 00 00 00` is generic count 1;
`ff 01 00` and `ff ff ff 01 00 00 00` are CString length 1. These consume
6, 3 and 7 prefix bytes respectively. Primitive acceptance does not establish
large-allocation or complete malformed-SAV acceptance. — SAV-758

## CArchive object framing

Objects and class descriptors share one index counter starting at 1. Indices
are assigned in stream order, not separately per class. — SAV-STREAM-013
(its "null arm unobserved" clause is superseded by SAV-ARCHREL-253; the shared
counter stands), SAV-ARCHREL-253

| u16 tag | Meaning | Following content |
|---|---|---|
| `0x0000` | Null object reference | None |
| `0xffff` | New class and its first object | `u16 schema`, `u16 nameLen`, `nameLen` ASCII bytes, object body |
| `0x8000 \| classIndex` | New instance of a known class | That class's object body |
| Plain nonzero object index | Back-reference to an existing object | No body |

A new class consumes the next index; its first instance immediately consumes
the next. Each later new instance consumes one index. A back-reference aliases
the previously loaded object. The located typed reader checks derivation from
the requested base class for a new class, a known class and a back-reference.
For example, `objref<Effect>` admits `Effect_DirectDamage`, but not an unrelated
class. — SAV-STREAM-010, SAV-STREAM-013 (null-arm clause superseded),
SAV-ARCHREL-253, SAV-774

Plain-index back-references, Unicode strings, noncanonical nonzero presence
bytes and marker-absent trailers have defined reader arms. Their primitive
framing does not establish acceptance of every complete file using them.
— SAV-FULLREAD-252, SAV-ARCHREL-253

## Two identity systems

| System | Defined by | Used by |
|---|---|---|
| CArchive class/object index | Object operation and tag order | Class reuse and object aliasing |
| Saved-address key | Explicit u32 in Token, Player, Spell or terrain body | Field-specific reference repair |

Neither is the Token creation-order ID, map-unit ID, Player slot ID or list
position. Saved-address keys are file-local and can be reminted on resave. See [Token](token.md#identity-map) and
[required relations](writing.md#required-relations). — SAV-ID-015,
SAV-ARCHREL-253, SAV-HEROID-065, SAV-ORIGCANON-398

## Parsing boundaries

Dispatch the [object programme](objects.md) at every archive call and honor
its counts, CString lengths and presence bytes. Inheritance invokes a base
programme directly. Embedded objects also serialize directly; they do not
necessarily emit their runtime class descriptor.

An Effect's 44-byte body does not imply another Effect follows it: it belongs
to its owner's list. A class-tag-shaped word or trailer marker inside a raw
member remains payload; scanning for it does not identify an archive boundary.
— SAV-EFFCHAIN-046 (its tag-scan support clause is partially retracted; the
44-byte programme and list placement stand), SAV-TAGSCAN-158, SAV-FULLREAD-252
