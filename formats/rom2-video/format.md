# ROM2 completion report and movie consumers

These contracts describe selected EN/RU single-player client consumers. They
do not specify the third-party SMK decoder, complete campaign scheduling or
rendered media equivalence.

## Selection and source families

The selected fresh campaign route reaches type-1 mission ID 10 after the
initial town interaction. Its success report acknowledgement calls the DLL
LeaveLocation output route. The enabled movie dispatcher indexes cutpaths.txt,
uses stem 1 `teleport`, tries numbered files 01..99, skips missing resources
and stops the batch on player return zero. Startup logos have a separate
dispatcher. The source families and index meaning are conditional native
contracts. — R2-ENGINE-073, R2-ENGINE-074, R2-ENGINE-075

The success report uses general flat text slot 0x8c. The selected startup
join identifies main.res:text/main.txt row 140 and dialogs.txt rows 43/154.
Report initialization stores the supplied text pointer and constructs labels
with commands 0x41d and 0x446. These sources differ from the cached mission
text used by the selected failure report. Live table mutation, report
font/background identities and complete input bindings remain Unknown.
— R2-ENGINE-093

Both measured video archives contain teleport/01.smk with 6753596 bytes;
the EN/RU payload hashes differ. Only this one key exists among the 99
selected numbered teleport SMK keys in each archive. Both same-stem REG
companions are 476 bytes and have equal hashes. This bounded archive-node
measurement does not establish loose override precedence or media equality.
— R2-ENGINE-096

## Companion numeric plan

The selected open replaces the SMK suffix with `.reg`. It requests Common
startx, starty, nFadings and nPanaramings with zero defaults. Fading and
Panaraming names are one-based; preserve the native spelling Panaraming.
— R2-ENGINE-077

Selected getters mask kind with 0xe and shift right one. Result 1 selects
integer DWORD +4; result 2 selects double QWORD +4/+8. The movie consumer
stores double endpoints as float values in a 16-byte segment array. The
measured first companion has 14 records and no pool bytes. Common gives
startx=0, starty=0 and nFadings=2. Its absent nPanaramings requests the zero
default if lookup returns missing. — R2-ENGINE-097

| Segment | Start frame | End frame | Start fade | End fade |
|---|---:|---:|---:|---:|
| Fading1 | 0 | 5 | 0 | 1 |
| Fading2 | 161 | 166 | 1 | 0 |

Fading triggers when descriptor+0x374 equals the next start frame, passes
end minus start and the stored endpoints, then advances the plan index.
Panaraming sets steps at start equality, clears them and advances at end
equality; after SmackNextFrame its step adds to destination coordinates.
The first measured companion has no Panaraming records. Complete lookup,
sort/malformed-input acceptance and the actual visual interpretation of
these plans remain Unknown. — R2-ENGINE-097

## Backing stream and return boundaries

The selected resource source retains member base, length and cursor. A
zero-count read positions the backing stream at base plus cursor before
returning zero. Nonzero reads clamp to the member's remaining length.
Movie open passes the source's backing handle to SmackOpen with flags
0xff000 and argument -1. Codec reads through that handle are not shown to
use the member clamp. Physical member/archive EOF, aliases and truncation
behavior remain Unknown. — R2-ENGINE-096

The outer player removes messages with PeekMessageA. IDs 0x12, 0x10,
0x201, 0x204, 0x100 and 0x104 return zero after player cleanup; 0x12/0x10
are translated/dispatched first. Other queued messages are dispatched and
the loop continues. With no queued message the player calls its tick with
argument 1. A terminating tick and an unsuccessful open both clean up and
return one. Thus one is a batch continuation value, not proof of successful
decode. The dispatcher clears its temporary presentation bit and retains
the seen bit set by this method. — R2-ENGINE-094

The active tick waits through the named SmackWait import, processes frame
plans, calls SmackDoFrame and uses presentation receiver slots +0x64/+0x80.
An unfocused buffer returns one without advancing; a receiver failure
returns zero. After processing, equality of descriptor DWORD +0x374 with
DWORD +0xc minus one returns zero. Otherwise SmackNextFrame and pan step
run. Descriptor-to-header binding, physical EOF, codec errors, hardware
receiver behavior, audio synchronization and media fidelity remain Unknown.
— R2-ENGINE-095
