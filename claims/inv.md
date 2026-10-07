# INV — install corpus and file signatures

Claims about what the lawful install contains and how each file family
identifies itself. Format of this file: [registry.md](registry.md).

## Install inventory and signatures

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| INV-CORPUS-001 | The GOG install holds 120 files; the manifest lists each with its size and SHA-256. | High | ● active | [EXP-0000](../experiments/EXP-0000-inventory/) |
| INV-SIG-002 | All 11 `*.res` files and `KIDS.LM` begin `26 59 41 31` (`&YA1`): one container family. | High | ● active | [EXP-0000](../experiments/EXP-0000-inventory/) |
| INV-SIG-003 | All 10 maps begin `4d 37 52 00` (`M7R␀`) followed by the `u32` value 20. | High | ● active | [EXP-0000](../experiments/EXP-0000-inventory/) |
| INV-SIG-004 | Saves begin `41 73 67 26` (`Asg&`); the next 2 bytes vary per save. | High / Medium | ● active | [EXP-0000](../experiments/EXP-0000-inventory/) |
| INV-SCOPE-005 | `Map Editor.opt` is an OLE2 compound document and `unins000.dat` an Inno Setup log; both are out of scope. | High | ● active | [EXP-0000](../experiments/EXP-0000-inventory/) |

### INV-CORPUS-001

The authoritative list, with size and SHA-256 per file, is the EXP-0000
manifest. The count excludes a tooling settings file once
written into the read-only install and since removed; an earlier count of 121
included it. Regenerating the inventory without it left all 120 hashes
unchanged.

**Confidence.** High. The inventory is an enumeration with hashes: a listing of
what is present, with no model between observation and claim.

### INV-SIG-002

**Confidence.** High. A direct byte observation over every such file in the
install. The binary confirms one container family: both readers validate the
one magic with a `CMP` (`RES-CODE-020`).

### INV-SIG-003

The value 20 is not a version number: it is the length of the file header,
which `R0465` reads and uses (`ALM-FRAME-031`).

**Confidence.** High. A direct byte observation over every map in the install.

### INV-SIG-004

The 2 varying bytes are the low half of the descriptor offset at `+0x04`,
which scales with body size (`SAV-PTR-003`).

**Confidence.** High for the magic, a direct observation over all 4 saves.
Medium for the meaning of the 2 varying bytes, which rests on `SAV-PTR-003`,
itself Medium on 4 samples of one map.

### INV-SCOPE-005

**Confidence.** High. Each file is identified by its own self-describing
signature. Nothing is claimed about their contents.
