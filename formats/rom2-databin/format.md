<a id="rom2-databin-database--identity-survey"></a>

# ROM2 Data.bin definition database

Client `world.res:data/data.bin` and server `world_srv.res:data/data.bin`
use eight ordered groups. Their grammar matches [ROM1 Data.bin](../databin/format.md)
with a 14-byte group-C raw block replacing ROM1's ten-byte block.
The root `templates.bin` has no established complete grammar.
— R2-ASSET-029, R2-ASSET-032

<a id="positive-control"></a><a id="worldres--world_srvres-result-the-grammar-corpus-wide"></a><a id="per-group-ghidra-identity"></a><a id="per-group-result"></a>

## Primitives and group order

Integers are little-endian. CString uses the established u8 length form
(0xff selects u16 length), followed by bytes. A title array has a u16 count;
a parameter array has a u16 count and u32 values; a collection has a u32
stored count. Groups C–H reserve null entry 0 and serialize entries
`1..count-1`. — R2-ASSET-029; shared primitive reference: DAT-GRAM-003

| Group | Collections | Entry after CString name | Stored counts, client/server | Title count |
|---|---|---|---|---:|
| A | Shapes, Materials | Nine binary64 values | 7/7, 16/16 | 11 |
| B | Magic | Parameter array | 50/50 | 30 |
| C | Armors, Shields, Weapons | Parameter array, raw14, second dword array | 31,10,28 on both | 18 |
| D | MagicItems | Parameter array, raw1, CString | 97/97 | 4 |
| E | Units | Parameter array, two CStrings | 242/242 | 64 |
| F | Humans | Parameter array, ten CStrings | 290/299 | 28 |
| G | Buildings | Parameter array | 180/180 | 9 |
| H | Spells | Parameter array, CString | 35/35 | 24 |

Write each group's title array once, followed by its collections in table
order. B and G have the same record shape; collection position identifies
their distinct roles. Group C's constructor initializes seven u16 values in
its raw block; wider field meanings are not supplied by that initialization.
— R2-ASSET-029

<a id="column-title-schema-vs-dat-schema-007"></a>

## Parameter and title schema

| Collection | Numeric parameter width |
|---|---:|
| Magic | 28 |
| Armors / Shields / Weapons | 17 |
| MagicItems | 2 |
| Units | 62 |
| Humans | 26 |
| Buildings | 6 |
| Spells | 22 |

These are installed parameterized-row widths, not replacement values for the
stored counts. Slot i is title i+1 because title 0 names the entry. Units has
63 titles after the name title: a shared 44-title prefix and one-title suffix
surround an 18-title block replacing ROM1's eleven-title treasure/spell block.
The additional titles include serverID, knownSpells and the five school skills.
Other title arrays retain their ROM1 forms; the Spells radius title differs
only by literal quote characters. — R2-ASSET-031

<a id="the-2556-byte-clientserver-delta-located"></a>

## Client/server values

The client payload is 149971 bytes and server payload 152527 bytes.
The 2556-byte size difference is Units +10 and Humans +2546. Units retains
the same stored count; Humans increases from 290 to 299. Other group spans
are equal; span equality alone is not a general value-equality statement.
— R2-ASSET-012, R2-ASSET-030

<a id="allods2exes-own-loading-code-q5-which-file-does-the-map-readers"></a><a id="type-id-resolution-consult"></a>

## Read and write sequence

1. Process groups A–H in order, reading each shared title array.
2. Read each collection count and the correct all-entry or reserved-zero loop.
3. Parse each exact entry programme, including group C's raw14 and second
   counted dword array.
4. Check the final endpoint against the input extent. Do not use the
   installed collection counts as parsing constants.

Structural emission uses the same order and primitive lengths. Preserve
opaque raw fields and array values; no complete semantic default set is
established. The binary loader has a once-only guard and tries loose Data.bin,
then its relative path inside World.res. Failure of both loads eleven named
`.txt` source tables and rewrites Data.bin. The actual ALM type-ID resolver
link to this table remains Unknown. — R2-ASSET-026, R2-ASSET-029

<a id="templatesbin-result"></a><a id="a2serverexe-classification-q5"></a><a id="not-yet-surveyed"></a>

## templates.bin and Unknowns

The 61802-byte root templates.bin is a distinct input. Applying either known
Data.bin group-C width does not complete it: the walk diverges in Shapes
entry 112 at byte 35004. Its competing head interpretations and data past
entry 111 remain Unknown. — R2-ASSET-005, R2-ASSET-032

No literal templates.bin reference is established in allods2.exe or
a2server.exe under the named ASCII/case/UTF-16 forms. ROM2 Map Editor.exe
contains the literal filename, but its consumer is untraced. Computed paths
remain possible. The broader absence wording of R2-ASSET-025 is narrowed to
the actual client/server literal-search scope. — R2-ASSET-025, R2-ASSET-032

The client's substantive Data.bin serializer bodies have no established
matching server functions. Generic shared helper shapes do not identify a
server consumer or disprove a separately compiled one. Complete runtime
server/classification relationships, entry semantics beyond the mapped
schema, editor consumption and original-compatible arbitrary-value writing
remain Unknown. — R2-ASSET-033

## Native NPC key source frontier

Selected EN/RU name parsers read native parameter slots 55 and 24, transfer
the low WORD to server-unit+0x14c, and the client update can copy it to the
NPC lookup word+0x1dc. The second parser replaces a WORD above 10000 by
(word/10)%1000. These positions match Units/Humans `serverID`, but the native
tables' load-time identity with these Data.bin groups remains unclosed.
This schema association is Medium. — R2-ENGINE-072

Stored serverID, transformed NPC WORD, ALM definition ID and NPC registry
section ID are distinct identities until a source join proves equivalence.
Authored key constraints, collisions after truncation/transformation and
original-compatible arbitrary-value binding remain Unknown.
— R2-ENGINE-072

## Selected construction input records

Measured character input stems Start_MF, Start_FF, Start_MM and Start_FM each
match one stored Humans name in the two byte-identical client payloads and
the distinct preserved RU server payload. Their parameter24 values are
10210..10213, respectively. The known low-WORD and greater-than-10000
transformation would yield 21 for each, conditional on the unclosed native
table load-time join and name normalization. Stored parameter values and
runtime object words remain distinct identities. No EN server payload or
complete constructor acceptance was measured. — R2-ENGINE-126
