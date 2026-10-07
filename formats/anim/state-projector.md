# State projection and a retained corpse stage

This reference classifies the 51 direct E8 sites calling `R0059` in
the single EN/RU ROM1 code image. It supplies original local predicates and
arguments. It does not prove that each path occurs in native play with a
particular drawable, or enumerate computed calls. — ANIM-129

## Receipt and admission

Server stage is byte `actor+0x13c`; client stage is byte `drawable+0x15a`.
They need not be equal before a message. The client saves its old stage,
copies server stage only under mask `0x8`, and runs the stage switch even
when that bit is absent. Old 1/current 1 calls hurt event 2; old 0/current 1
calls event 3. Voice admission is a further condition. — ANIM-095, ANIM-126

A projector call is not necessarily a state-sync message. Recipient 0
iterates the player list and selects a human player or the actor's owner.
An explicit recipient selects that player. For a nonowner, object-kind and
class branches reduce the supplied mask with `0x507b` or `0x50fb`, and may
clear `0x1000`; maximum mana zero clears `0x2`. These reductions preserve
stage bit `0x8`. The ordinary message send requires
`effective_mask & 0xff7fff7f != 0`. Bits `0x80` and `0x800000` have separate
sends. Thus mask 0 and mask `0x80` alone do not reach the client stage
switch. A dynamic dirty mask must first be resolved and filtered. — ANIM-129

The repeated-stage consequence requires an ordinary state sync actually
received for the existing drawable, old stage 1, and current stage still 1.
Without bit `0x8` the message leaves that stage intact. With bit `0x8` the
server stage controls it. A health or server-stage guard alone does not
establish the old client stage. — ANIM-129, ANIM-133

## Per-site predicates

`A` is the actor passed to the projector; `P` is a player. `owner` means
`A+0x14`. Masks below are supplied masks, before the common reductions.
`dirty` means dword `A+0x150`. `P+0x34` is the player's primary actor.
An address is an original instruction site, not a semantic sender name.

The classifications are bounded local conclusions:

- **Retain**: the supplied mask omits `0x8`. If its ordinary message is
  admitted and reaches an existing stage-1 drawable, it preserves stage 1.
- **Conditional**: mask carries or can carry `0x8`; stage and alias/lifetime
  must be established at dispatch. This is an unresolved repeated-stage path.
- **Away**: the selected source explicitly supplies stage 0 or 2..5.
  Dispatch to old client stage 1 can occur, but it is not a 1-to-1 transition.
- **New**: the object is constructed on the selected path. It does not show
  that an existing actor can take that path; previous drawable/id reuse is Unknown.
- **Separate**: no ordinary state sync is emitted by that supplied mask.
- **Fanout**: recipient recursion repeats its input; it supplies no new trigger.

### Player and command paths

Each row states a local gate. Upstream command admission, opaque helper
effects and collection membership remain Unknown where named. — ANIM-132

| Site | Actor and recipient | Mask | Original gate or producer | Classification |
|---|---|---|---|---|
| `L12974` | player lookup's `P+0x34`; 0 | `ffffffff` | string arm `L13176`, player and primary nonnull; stores attributes then calls actor slot `0x50` | Conditional; derive side effects unresolved |
| `L12975` | input A; 0 | `ffffffff` | owner change and owner-dependent visibility-bit clearing; no local health/stage gate | Conditional |
| `L03739` | `P+0x34`; P | `ffffffff` | primary nonnull at `L12976`; preceding entry helpers | Conditional; preceding helper effects unresolved |
| `L03848` | constructed A; P | `ffffffff` | existing primary returns; construct, assign id, attach primary | New |
| `L03266` | `P+0x34`; P | `ffffffff` | server stage zero at `L03267..L03268`; item lookup/predicate succeeds | Away to 0 |
| `L12977` | `P+0x34`; 0 | `bf7fff7f` | separate string arm `L13177`, mode 1; `R1562` helper | Conditional; bypasses preceding stage-zero gate |
| `L12978` | iterator A from `P+0x20`; 0 | `bf7fff7f` | same arm, mode 2, nonnull iterator; helper | Conditional; membership/helper effects unresolved |
| `L12979` | `P+0x34`; owner | `bf7fff7f` | string arm `L13178`, mode 1, `A+0x140` nonnull | Conditional |
| `L12980` | `P+0x34`; owner | `bf7fff7f` | string arm `L13179`, mode 1, `A+0x140` nonnull; loop 1..28 | Conditional |
| `L03270` | lookup A; 0 | local bits OR dirty | opcode `0x22`, nonnull lookup and object-selection branches; helper writes contribute dirty | Conditional; fixed local bits omit 8, dirty unresolved |
| `L12981` | `P+0x34`; P | `ffffffff` | opcode `0xbe`, player lookup, primary exists | Conditional |
| `L12982` | loaded A; owner | `ffffffff` | same opcode with no primary; nonnull load result, id/owner assignment | Conditional; load identity/lifetime unresolved |
| `L12983` | returned constructed A; owner | `ffffffff` | opcode `0x49`, session `+0xc==0`, source and primary nonnull, index in bounds, `R0066` result nonnull | New |
| `L03269` | lookup A; owner | `1f001304` | opcode `0x3d`, session `+0xc==0`, nonnull lookup, index 0..5, cost and slot `0x30` checks | Retain; upstream/virtual admission unresolved |
| `L12984` | `P+0x34`; 0 | `ffffffff` | `P+0x3d!=0`, `P+0x28==0`, `P+0x3c>=2`, session `+0xc!=0`, human P, signed health below -53; calls `R0823` first | Conditional; intervening restoration unresolved |
| `L12985` | newly constructed A; 0 | `ffffffff` | `R0066` constructor path, `A+0xe!=0`, input receiver `+0x2c!=0` | New |
| `L12986` | same new A; owner | `ffffffff` | same path, receiver `+0x2c==0`; owner copied from input actor | New |

The numeric dispatch above uses original byte table `L00564`, dword table
`L03265`, and opcode bias 2. It requires command byte `+4==0`. It does not
assign gameplay names to those opcodes. — ANIM-132

### Projection and actor-lifetime paths

The death, bleed and entry paths are prior controls. Decay and regeneration
are distinguished from repeated-stage receipt by their actual masks and
stage producers. — ANIM-126, ANIM-130, ANIM-132

| Site | Actor and recipient | Mask | Original gate or producer | Classification |
|---|---|---|---|---|
| `L03228` | same A; selected P | inherited | recipient-zero loop, human P or owner | Fanout |
| `L03521` | input A; explicit P | `ffffffff` | human P; `(A.visibility & P.visibility)==0`; actor slot `0x2c` nonzero | Conditional; full re-projection of server stage |
| `L03740` | `P+0x34`; P | `ffffffff` | entry projection; primary nonnull | Conditional; prior entry control |
| `L03789` | world-list A; P | `ffffffff` | entry projection; nonnull iterator | Conditional; prior entry control |
| `L03791` | dead-list A; P | `ffffffff` | entry projection; server stage below 5 | Conditional; prior entry control |
| `L12987` | `P+0x34`; P | `ffffffff` | owner's projection; primary nonnull | Conditional; prior entry control |
| `L12988` | owner's list A; P | `ffffffff` | nonnull iterator; clears this player's visibility bits first | Conditional; prior entry control |
| `L12989` | tick A; 0 | `409` | dying branch, server stage 0; writes stage 1 before send | Conditional; prior death control |
| `L08224` | tick A; 0 | `1` | act !=`0x10`, health >0 and below max, regeneration divisor nonzero, full counter divisible by 4 | Retain; positive health does not test client stage |
| `L08225` | tick A; 0 | `2` | act !=`0x10`, initial health >0, mana below max; positive-health branch | Retain; no four-tick gate on this mana arm |
| `L03113` | tick A; owner | `1` | act !=`0x10`, health <0, full counter divisible by 4; subtracts one | Retain; prior bleed control |
| `L03271` | input A; 0 | `20` | placement-search helper succeeds; clear off-map bit and register A | Retain; upstream actor admission unresolved |
| `L03272` | input A; 0 | `20` | retained-cell search or radius-3 search succeeds; clear off-map bit and register A | Retain; upstream actor admission unresolved |
| `L12990` | input A; 0 | `a08000` | nonnull removed object; detach/delete it, then project actor | Retain; ordinary bits `208000` survive for owner; other-recipient filtering matters |
| `L03277` | teardown A; owner | `80` | A is owner's primary; increment player `+0x54` | Separate |
| `L03273` | decay A; 0 | `9` | post-decrement health below -600 and stage changed; writes stage 5 | Away to 5 |
| `L03274` | decay A; 0 | `9` | stage changed after thresholds below -10/-20/-40; health at least -600 | Away to 2..4 |
| `L12991` | award A; 0 | `1f001304` | actor type `0x21..0x3f`, skill raised; derive slot `0x50` | Retain; skill gate per HERO-SKILLUP-073 |
| `L12992` | award A; 0 | `4` | same actor type, no skill raised, scaled award nonzero | Retain; upstream award admission unresolved |
| `L12993` | newly constructed stock A; input P | `ffffffff` | nonnull stock-array entry after constructors/id assignment | New |
| `L12994` | receiver's `+0x74` A; owner | `a08000` | item loop ends; recalculates money/load then projects actor | Retain; upstream shop actor/lifetime unresolved |
| `L12995` | receiver's `+0x74` A; owner | `a08000` | purchase loop ends or money insufficiency; recalculates load then projects actor | Retain; upstream shop actor/lifetime unresolved |
| `L09201` | input A; 0 | `ffffffff` | centred position; fetched cell type `0x1a`; both `R0048` checks succeed; relocates and registers A | Conditional; path admission/helper side effects unresolved |
| `L12996` | input A; 0 | `ffffffff` | footprint-block and placement gates reach relocation body; releases reservation, writes position, registers A | Conditional; outer invocation at stage 1 unresolved |

### Token and attachment paths

Token ids are the original selector byte at token `+8`, read through table
`L03242`. An effect's original callbacks can change dirty; a dynamic mask
does not imply either a send or an unchanged stage. — ANIM-130, ANIM-131

| Site | Actor and recipient | Mask | Original gate or producer | Classification |
|---|---|---|---|---|
| `L03249` | token target A; 0 | dirty | token 6; player relationship gate, health >-10 and regeneration divisor nonzero; add capped restoration | Conditional; dirty starts 0, health helper sets 1, crossing zero adds `428` and stage 0; upstream target kind Unknown |
| `L03259` | token target A; 0 | `1` | token 11; actor target, positive `min(request,health+10)`; subtract amount after caster callback | Retain; target need not have health >0 |
| `L03260` | token caster A; 0 | `1` | same accepted token-11 transfer; add amount to caster via health helper | Retain even if helper has reset server stage; literal mask omits 8; caster-stage reachability Unknown |
| `L05854` | token caster A; 0 | `20` | token 26; calls relocation helper, then sends without testing its return | Retain; helper itself also has a full-mask site |
| `L03275` | selected dead-list A; 0 | `9` | token 25; matching cell and server stage 2; writes health -10001 and stage 5 | Away to 5 |
| `L03276` | newly constructed A; 0 | `bf7fff7f` | same token-25 path; constructor, placement and owner-list registration | New; does not project the consumed actor |
| `L03261` | attachment's input A; 0 | dirty | accepted flags `effect+0x3d & (byte[L04012] OR byte[L04013])`; continuous flag set; duration divisible by 8; callback slot `0x40` | Conditional; dirty cleared before callback |
| `L03262` | attachment's input A; 0 | dirty | accepted flags; duration <=9600, decrement reaches 0, continuous flag clear; callback slot `0x44` | Conditional; dirty cleared before callback |
| `L03263` | attachment's input A; 0 | dirty | opposite effect pair `0x17/0x1b` found; callback `0x44`, notify and remove | Conditional; dirty cleared at entry |
| `L03264` | attachment's input A; 0 | dirty | no opposite-pair early branch; same effect refresh/replace or new attachment, or direct non-timed callback `0x40` | Conditional; dirty depends on selected callbacks |

## Health restoration and stage

The health helper always ORs dirty with `0x1`. It records whether health was
nonpositive, adds the signed word delta and caps at maximum health. If that
operation crosses to positive health, the revival helper ORs dirty with
`0x428` and sets stage 0. The token-6 caller sends dirty, so restoration that
leaves health nonpositive preserves a stage-1 client's stage; restoration
that crosses zero supplies stage 0. A literal mask-1 caller, such as the
token-11 caster send, does not transmit the helper's stage reset. The next
message may supply it. — HERO-REVIVE-068, ANIM-130

Synthetic execution of the original local instructions uses stage 1,
health -5, maximum 100 and dirty 0. Deltas 3, 5 and 6 yield respectively
`(-2,1,1)`, `(0,1,1)` and `(1,0x429,0)` for health, dirty and stage.
The position call and subsequent derive/registration callbacks are outside
those controls. They prove the local branch split, not native scheduling.
— ANIM-130

## Client calls and remaining scope

Between the stage copy and switch, `R0589` receives two scalar arguments
and returns a scalar, with only stack stores. The `R0544` call copies
12 bytes into `drawable+0xe4..0xef`, outside stage `+0x15a`. The preceding
slot-`0xc` call uses the campaign record at campaign screen `+0x548`, obtained
through the application singleton's slot `0x7c`. The native record vtable
`L10123` maps that slot to `R0678`, a scalar getter of receiver `+4`, with
no stores or calls. These three calls do not clear the stage for their native
receiver bindings. An overridden receiver or an aliased object outside those
bindings is not excluded. — ANIM-133

High covers named local instructions, literal arguments and the restoration
split. Medium covers the finite site classification and conditional
repeated-stage feasibility. Computed callers, indirect callback dirty/stage
writes, unresolved aliases, drawable/id reuse, message ordering and native
frequency remain Unknown. Native playback/device audibility is outside this
static population. — ANIM-129, ANIM-131, ANIM-132, ANIM-133
