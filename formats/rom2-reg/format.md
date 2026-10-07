<a id="rom2-inline-registry-ya1-nested--identity-survey"></a>

# ROM2 inline REG store (`&YA1`)

Nested ROM2 registries use the inline REG envelope. It is distinct from the
tail-registry RES envelope despite the shared magic. Resource context and
record geometry select the parser. — R2-ASSET-006

<a id="result"></a>

## Layout

| Position | Width | Data |
|---|---:|---|
| 0 | 24 bytes | Six-dword header; magic `&YA1`, root fields, record count at 0x10 |
| 0x18 | 32×R bytes | Record array |
| 0x18+32×R | 4 bytes | Pool byte length |
| 0x1c+32×R | poolLen bytes | Pool |

`registryEnd = 24 + 32*R + 4 + poolLen`. This framing covers the preserved
nested registries. The installed `data/ai.reg` and `data/map.reg` lengths are
156 and 1020 bytes in both world archives; equal lengths do not prove equal
values. — R2-ASSET-006

## Read and write order

Read the six header dwords, R records, u32 pool length and pool bytes.
The corresponding structural emitter writes those same regions in order.
Raw record and pool values must remain opaque where ROM2 type/lookup semantics
have not been established. [ROM1 REG](../reg/format.md) supplies a layout
cross-reference, not an authority for unverified ROM2 value behavior.
REG-REC-032's ROM1 record layout is retained; its lookup clause is partially
retracted and cannot be imported as an unconditional ROM2 rule.
— R2-ASSET-006, REG-FMT-031, REG-REC-032

<a id="not-yet-surveyed"></a>

## Unknowns

General ROM2 record kinds, value interpretation, key-set differences, native
lookup, sort and application/writer acceptance are unspecified. Recognition by a
nested `&YA1` payload is broader than a `.reg` extension filter.

## Selected movie companion consumer

The selected native movie-open path replaces the SMK suffix with `.reg` and
loads an inline store. Common requests startx, starty, nFadings and
nPanaramings, with zero defaults. Fading1..N requests startframe/endframe
and floating startfade/endfade; Panaraming1..N requests startframe/endframe
and stepx/stepy. Preserve the native spelling Panaraming. — R2-ENGINE-077

The first selected teleport companion in both installs has 14 records and
a zero-byte pool, totaling 476 bytes under the envelope above. The native
lower open path then supplies the SMK resource to its named SmackOpen import.
General lookup/type semantics, malformed companion acceptance, full frame-plan
application, rendering and audio behavior remain Unknown. — R2-ENGINE-077

## Selected numeric branches and frame plan

The movie's selected getters mask kind with 0xe and shift right one.
Result 1 loads integer DWORD +4; result 2 loads double QWORD +4/+8.
Missing records return the caller's default. These branches do not establish
a general ROM2 lookup/comparator or acceptance grammar. The movie stores
double fade endpoints as float values in 16-byte segment records.
— R2-ENGINE-097

The measured first companion has Common/startx=0, starty=0 and nFadings=2;
nPanaramings is absent. Fading1 stores start/end 0/5 and doubles 0/1;
Fading2 stores 161/166 and 1/0. There are no Panaraming records in this
companion. Fading triggers at exact descriptor start equality and advances
the plan index; Panaraming sets/clears stored steps at start/end equality
and updates destination coordinates after the next-frame call. Missed
equality, visual interpolation, malformed lookup and timing remain Unknown.
— R2-ENGINE-097
