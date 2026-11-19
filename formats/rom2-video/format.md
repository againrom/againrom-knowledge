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
contracts. R2-ENGINE-073 is partially retracted only for its world-map
record-choice clause; the clauses here stand. — R2-ENGINE-073, R2-ENGINE-074,
R2-ENGINE-075

The success report uses general flat text slot 0x8c. The selected startup
join identifies main.res:text/main.txt row 140 and dialogs.txt rows 43/154.
Report initialization stores the supplied text pointer and constructs labels
with commands 0x41d and 0x446. These sources differ from the cached mission
text used by the selected failure report. Live table mutation remains
Unknown; the report's font, frame and input are in Outcome pages below.
— R2-ENGINE-093

## Outcome pages

Both outcome reports are popups placed by the snap of R2-ENGINE-608: the
width becomes 96k + 8 and the height 64k + 104, placed by the snap at
x = (screen width - w) >> 1, y = (screen height - h) >> 1 (arithmetic High;
that this is the drawn rectangle is Medium, the rectangle setter `L2.00637`
being unread). Children are page-relative. Measured in EN; RU is Medium except where noted.

| Page | Built | Placed size | Children (page-relative) |
|---|---|---|---|
| victory | (128,200,512,380) | 296x168, at (172,156) on 640x480 (Medium) | text (40,32,232,80); buttons (40,84,248,108), (40,108,248,132) |
| first failure | (32,100,608,380) | 488 wide, height from the text, then 64 more below | text area (40,56,448,144) before `L2.01495`; title above it (Unknown offset); buttons (61,H-96,427,H-72), (61,H-72,427,H-48) |

— R2-ENGINE-1300, R2-ENGINE-1301

Showing a report remaps the whole screen once through row 3 of table
`L2.07326` (Medium that this darkens it: the table's writer is unread; the
amount is Unknown), then the page paints `interface/lm.256` frames 0..8 on
its rectangle less 8 pixels at the right and bottom: a level-6 shadow pass
offset by 8, the four 48x48 corners, the 96x48 top and bottom edges every 96
pixels, the 48x64 side edges every 64 pixels and the 96x64 fill. The
children are painted after the frame (Medium: the child-paint route
`L2.00579` is unread). A text child repaints the frame under its own
rectangle before its words (High). — R2-ENGINE-1302, R2-ASSET-244

Every word uses font1. The victory words (main row 140) are anchored at the
centre of their rectangle's top edge (the anchor is High; the centring by
the font method is Medium); the failure title (slot 0x8d) is left-aligned
and the failure body is wrapped. Captions are dialogs rows 43 and 154
(victory) and indices 0x2c and 0x23 (failure). The rows main 140 and
dialogs 43 and 154 carry the Medium of R2-ENGINE-093. Normal colour is
table `L2.05264`, whose values are Unknown. RU rows reach glyphs through
the font1 mapper of R2-ENGINE-052. — R2-ENGINE-1303

| Input | Victory page | Failure page |
|---|---|---|
| release inside fires the button (High); Enter and the `~` hot letter fire it in the button's own key arm (High); that a key reaches the focused button, and the key value of the letter, are Medium | Victory! sends 0x41d (Leave route), Continue sends 0x446 | first button (0x445): main menu; Load Game (0x446): load page |
| Esc | closes with 0x446; the application's next step is Unknown | as Load Game |
| Tab, arrows | focus moves; Up and Down step focus | focus moves |
| first focus | the first child by `L2.03078` (1,0); that it is a button is Medium (the text child is child 1; the step body is read for 80 instructions) | first button (id 4, High) |

The Load Game button is set by slot +0x1c (Medium: disabled) when no save
file is found. — R2-ENGINE-1304, R2-ENGINE-1305

The victory report plays sound entry 0xe and the failure report entry
0x10 (sfx.reg Sfx14 and Sfx16, Medium) right after the page is shown; no
click sound was found in the button bodies read. — R2-ENGINE-1306

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

## Teleport movie presentation

`teleport/01.smk` is SMK2, 640x360, 167 frames at 15 frames a second, with
one 22050 Hz 16-bit stereo compressed track; EN and RU members differ in
bytes but not in these fields. — R2-ASSET-243

The player opens the movie for a 640x360 target (the one call in the player
reached through the dispatcher) with its destination at the screen's
640x480 centring offset plus 60 rows. Its doubling branch is taken only
when twice the SMK width or twice its height fits the target. The player
does not take its doubling branch for a 640x360 SMK (High, EN); that the
blit then draws 1:1 is Medium (the meaning of SmackBlitOpen flags 1 and 2 is
unread). The destination is (0,60) on 640x480 (High) and is centred on the
640x360 target in larger modes: (80,120) on 800x600 and (192,204) on
1024x768. A 480-high SMK would be placed at the offset itself. The source
rectangle starts at the companion's startx, starty. Before the first file
and after the batch the dispatcher calls the fill with colour 0 over
`L2.07341..L2.00629` and presents (Medium that it covers the screen and
runs: the early-return flag and the rectangle values are unread), so the
area around the picture is read as black (Medium). The doubling test and
the destinations are read in EN only; RU is Medium outside the three equal
player bodies. — R2-ENGINE-1307

Sound plays through the movie library with its own DirectSound choice
(argument 0) and sound switched on; the game's sound object gets one call
pair before the batch and one call after it (Medium: pause and resume). A
key press, a left or right mouse press, close or quit ends the player with
zero, which ends the whole teleport batch. The victory route then
continues as in the campaign flow, after the same fill (Medium). — R2-ENGINE-1308
