# Claim registry — ROM2 SESSION (multiplayer socket, message surface, state ownership, save and character files)

Level 2 ledger. Index: [registry.md](registry.md). Legend: `✔ promoted` (also in a spec) ·
`● active` · `✖ retracted`. IDs are permanent.

This ledger is a static read of `gameversions/rom2-ru/allods2.exe` and `a2server.exe` (never
executed, no network connection opened): PE import/section tables, an independent byte-level
ordinal and call-site scan, Ghidra decompilation, and named literal-string xrefs; two rows also
cross-check `gameversions/en/rom.exe`, scoped to one function each. Two rows (`R2-SESSION-003`,
`R2-SESSION-010`) describe a structural byte layout (a wire message frame; a record-dispatch
arity) and are cross-published on canonical format pages — see
`experiments/EXP-2002-session-model/format-surfaces.tsv`. The rest describe process and session
behaviour, not a stored or wire layout. This ledger cites `claims/rom2-engine.md` and
`claims/rom2-asset.md` by id where an existing row already covers an address or population this
ledger reaches, and never restates their evidence. `R2-SESSION-001` through `022` are reserved
(`pipeline/ALLOCATIONS.md`); `011`–`016` are retired unused, recorded in
`experiments/EXP-2002-session-model/EXP-2002.md`. `R2-SESSION-017` through `022` (topic Save and
character files) read the save routines, the `scenario.dll` save record and the character-file
routines of `allods2.exe` statically; see `experiments/EXP-2005-rom2-save-format/format-surfaces.tsv`. The widths and operands behind `R2-SESSION-017` through `021` are listed in
`experiments/EXP-2005-rom2-save-format/evidence/operands.tsv`; the decompiled bodies are not
committed and are reproducible only by running `tools/ghidra/DecompAddrs.java` at the cited addresses.

## Transports and sockets

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-001 | Both binaries import the whole `WSOCK32.DLL` surface by ordinal; a decompiler-independent byte scan finds the same complete call-site population as the reference tool, plus one site the tool missed. | High | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-002 | DirectPlay is a second, live transport: `a2server.exe` instantiates it through COM and one selector field routes `StartServer` and `Connect` between the socket path and a 12-method wrapper. | High | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |

### R2-SESSION-001

Both `allods2.exe` and `a2server.exe` import the entire `WSOCK32.DLL` surface by ordinal, not by name; a decompiler-independent byte scan finds the same complete call-site population as the decompiler-reference tool, plus one call site the reference tool missed. `tools/ordxref` reads `IMAGE_ORDINAL_FLAG` directly from every `FirstThunk`/`OriginalFirstThunk` entry: `allods2.exe` has 23 `WSOCK32.DLL` slots, `a2server.exe` has 19, and in both binaries all of them are by-ordinal, none by name (`evidence/allods2-wsock32-ordinals.tsv`, `evidence/a2server-wsock32-ordinals.tsv`; the already-committed `evidence/{allods2,a2server}.exe.pescan.tsv` show the same import rows independently). The ordinal-to-name label on each row (`socket`, `connect`, `recv`, …) is Ghidra's own bundled WSOCK32 ordinal database, not re-derived by this experiment; only the ordinal-vs-name fact itself is read byte-level with no decompiler involved. A separate brute-force scan of every `0xE8` (near CALL) opcode in every executable section, matched against the thunk stubs found above, finds a complete, fully-attributed call-site population in both binaries (0 unattributed): `allods2.exe` 50 call sites across 17 owner functions, `a2server.exe` 35 call sites across 12 owner functions (`evidence/{allods2,a2server}-wsock32-callsites.tsv`). Reconstructing the pairing by address order and by which WSOCK32 functions and call counts each owner has, all 12 of `a2server.exe`'s owners map one-to-one onto 12 of `allods2.exe`'s 17 (identical membership and count, e.g. both binaries' listen-setup function calls `socket`+`ioctlsocket`+`bind`+`listen`, exactly 4 sites, and both binaries' connect-setup function calls `socket`+`bind`+`connect`, exactly 3 sites — confirmed by direct decompilation of `allods2.exe R2.0124`/`a2server.exe R2.0125` and `R2.0126`/`R2.0127`, structurally identical bodies). `allods2.exe`'s remaining 5 owners are a client-only module absent from `a2server.exe`: `R2.0128` (a second, separate `WSAStartup` call site), `R2.0129` (`closesocket`+`WSACleanup`), `R2.0130` (`closesocket`, `socket(2,3,1)` — `AF_INET`/`SOCK_RAW`/`IPPROTO_ICMP`, a raw socket, confirmed by direct decompilation —, `setsockopt`×2, `gethostbyname`, `inet_addr`, `inet_ntoa`), `R2.0131` (`sendto`, `recvfrom`, `WSAGetLastError`×2), and `R2.0020` (one more `WSAStartup` call site). `R2.0020`'s call site is present in the byte scan and absent from the decompiler-reference tool's own output (`evidence/allods2-wsock32.txt` lists only `R2.0128` and `R2.0132` as `WSAStartup` owners); this experiment found no `WSACleanup` call site attributed to `R2.0020` anywhere in the image, so it does not confirm a paired `WSAStartup`/`WSACleanup` for that function, only the one `WSAStartup` site. The session/listen socket itself is `socket(2,1,6)` (`AF_INET`/`SOCK_STREAM`/`IPPROTO_TCP`) in both binaries, read directly from the two structurally identical listen-setup functions above

**Confidence.** High (two independent instruments — a decompiler-reference tool and a raw byte scan that never resolves a function boundary — agree on a complete, zero-unattributed population in both binaries; this directly rules out H1b, "imported by name": no by-name thunk exists, and rules out H1d, a `WS2_32.DLL`/Winsock 2 alternative: no such import exists in either table. This corrects this row's own prior text, which stated the opposite — "imported by name only" — an error, not a hypothesis)

### R2-SESSION-002

DirectPlay is a second, live transport selected at runtime, not dead vocabulary: `a2server.exe` instantiates it through COM, and the same selector field routes both `StartServer` and `Connect` between the socket path and a 12-method DirectPlay wrapper. `a2server.exe R2.0133` calls `ole32.dll!CoCreateInstance` twice (retrying with a fallback interface id on failure), storing the resulting interface pointer at `this+0x5c4`; the three 16-byte operands read as raw GUIDs, not decoded against any external name table, are `{D1EB6D20-8923-11D0-9D97-00A0C90A43CB}` (class id argument), `{0AB1C531-4745-11D1-A7A1-0000F803ABFC}` (first interface id) and `{133EFE41-32DC-11D0-9CFB-00A0C90A43CB}` (fallback interface id) (`evidence/a2server-directplay-guids.txt`). Neither import table needs a dedicated DirectPlay module for this to be real: COM instantiates by class id through the already-imported `OLE32.DLL`, so the prior text's "neither import table names a DirectPlay-specific module" was not, by itself, evidence the code was unreachable — that reading is retracted. A named driver class `CLlDriver` carries exactly 13 distinct `CLlDriver::`-prefixed method-name strings (32 literal occurrences total) resolving to exactly 12 distinct owner functions in both binaries alike — `StartServerDP` and `StartServerDp` are two log strings inside the same one function, not two methods (`evidence/{allods2,a2server}-clldriver-methods.txt`; every corresponding string pair carries the identical hash across binaries). A cruder raw-substring count agrees and adds cross-binary symmetry this structured count does not by itself show: the bare literal `CLlDriver` occurs 38 times and `DirectPlay` 13 times in each binary's string table, identically in `allods2.exe` and `a2server.exe` (`evidence/{allods2,a2server}-hat-strings/STRXREF_SUMMARY.md`, independently recounted this pass). Two of those methods, `StartServer` (`a2server.exe R2.0134`) and `Connect` (`R2.0135`), each read the same `this`-relative field, `+0x4f0`, and branch on it: `==4` calls the socket-path helper (`R2.0125` / `R2.0127`, `R2-SESSION-001`'s listen/connect functions), any other value calls the `*Dp` (DirectPlay) counterpart (`R2.0136` / `R2.0137`) — read directly from both functions' decompilation, not inferred. `StartServerDp` (`R2.0136`) calls the DirectPlay object's own vtable at slots `+0x38`, `+0x60` and `+0x18` (capability query, session/player setup, then a create-with-event call); `HandleMessageDp` (`R2.0138`) calls a fourth slot, `+0x68`, in a send-retry loop that sleeps and retries while the call returns one specific error code — a slot the prior pass's finding did not name. `srv.cfg`'s shipped `Protocol=WSOCK_TCPIP` (`R2-SESSION-008`) selects the socket path for this lawful install; DirectPlay is the code the selector does not choose here, not code with no path to run at all

**Confidence.** High (the CLlDriver method population, the selector field and its two independent call sites, and the CoCreateInstance call are all read directly from decompiled bodies in both binaries where population claims are made; this rules out H1c/H2's "DirectPlay is unreachable/vestigial" reading — a real runtime selector routes to it from at least two call sites, on a real COM object)

## Message surface

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-003 | The 8-byte wire header has four fixed fields (payload length, codec selector, codec input size, passthrough byte) read identically in both binaries; the 150-byte bound is the 142-byte cap plus the header. | High | ✔ promoted | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-004 | The receive path ends in a bounded record queue gated by the codec-selector byte, not an opcode dispatcher; its six-field queue state sits at a different offset than previously stated. | High / Unknown | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-005 | `allods2.exe`'s inbound path applies the same XOR transform against the same key table at the same structural position as `a2server.exe`; the call site's arithmetic does not match the earlier +8/-8 reading. | High / Unknown | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |

### R2-SESSION-003

The 8-byte wire header has four fields at fixed offsets — payload length, codec selector, codec input size, and a passthrough byte — read identically (once each transport's own record-copy displacement is accounted for) in both binaries, and a fifth bound the prior pass treated as a separate constant is the same 142-byte cap expressed in different units. Relative to the header's own start: `+0` is a `u16` payload length, valid observed range 1..142 (0 or ≥143 ends the receive loop, `R2-SESSION-004`); `+4` is a `u8` codec selector (`0` skips decoding entirely — the branch `R2-SESSION-004` gates on); `+5` is a `u16` codec input size, passed as the decoder's own input-length argument; `+7` is a `u8` passthrough byte the decoder never reads, copied verbatim into the post-decode record. `a2server.exe`'s `CBufferManager::ReceiveData` (`R2.0139`) reads these at record-buffer offsets `+0xc`/`+0xd`/`+0xf` (the wire header is copied to `record+8` earlier in the same receive loop, `R2.0140`, so `record+8+4=+0xc` etc.); `allods2.exe`'s counterpart (`R2.0141`) reads the same four fields at header-relative `+4`/`+5`/`+7` directly, through its own record-pointer convention — both readings independently confirmed by decompilation, not assumed to agree because the functions share source. The 142/143 payload-length bound is read from `a2server.exe R2.0140`'s own receive-loop continuation test (continues while accumulated header bytes `<8`, or the payload-length field is both nonzero and `<0x8f`; stops otherwise) and separately from `R2.0142`'s own decoder call, which both takes `0x8e` (142) as a bound argument and has its return value checked against `0x8e` again after the call. `a2server.exe`'s DirectPlay message handler (`HandleMessageDp`, `R2.0138`) bounds its own total message length to `7 < len < 0x97`, i.e. 8..150 — this is not a fifth, independent constant: 150 = 142 (the same payload cap) + 8 (the header this transport has not yet split off), the same bound restated in total-message units instead of payload-only units. `R2.0143` (`allods2.exe`) and `R2.0140` (`a2server.exe`) are the same source function compiled twice, not two independently designed implementations; the prior pass's "four independent call sites in two independently compiled binaries" framing is dropped as a coincidence argument, since shared source is the simpler explanation for shared constants

**Confidence.** High (the field layout is read directly at both transports' copy points in both binaries, and the 142-vs-150 relationship is arithmetic, not a coincidence needing an independence argument; N4's own census sweep, unchanged by this pass, found the immediate 142 in only 10 of 8,413 and 10 of 8,928 functions respectively — about 0.12% each — which is the discriminator against "an unrelated buffer-size constant that happens to recur")

### R2-SESSION-004

The function immediately reachable at the end of the receive path in both binaries is a bounded record queue gated by the codec-selector byte, not an opcode dispatcher; its six-field queue state sits at a different offset than previously stated, independently and repeatedly confirmed from both the enqueue and dequeue bodies. `a2server.exe`'s `R2.0139` and `allods2.exe`'s `R2.0141` are both self-identified by the literal ASCII string `CBufferManager::ReceiveData` (one xref each). Both test the `+4` codec-selector byte (`R2-SESSION-003`) as a 2-arm branch — decode when nonzero, skip straight to enqueue when zero — the branch a prior pass recorded, then withdrew as unreproduced (`HYPOTHESES.md`); this pass finds it at exactly this position in both binaries, gating a codec decision, not a message-type opcode, so the withdrawal itself was the error. Both reject an oversized record (`R2-SESSION-003`) and otherwise hand it to a free-list-backed queue. Reading `a2server.exe R2.0139` (enqueue) and `R2.0144` (dequeue) directly and repeatedly, the six queue fields sit at `this+0x268` (head), `+0x26c` (tail), `+0x270` (count), `+0x274` (free-list head), `+0x278` (block-allocator base) and `+0x27c` (block-allocator capacity), immediately followed with no gap by the critical section at `+0x280` — both functions agree with each other and with themselves on every read. This experiment's brief and its confidence review both name `this+0x264` for the same six fields; this pass finds no reference to `+0x264` in either function and cannot reconcile the two numbers — it publishes the directly-read `+0x268` and records the discrepancy as unresolved, not silently adopting either source. The dequeue function has exactly ten distinct caller functions in each binary, seventeen call sites each, zero unattributed (`evidence/{a2server,allods2}-dequeue-callers.txt`) — this specific count from the prior pass reproduces exactly. Neither enqueue nor dequeue function reads or branches on a message-type/opcode field

**Confidence.** High for the enqueue-not-dispatch shape, the codec-selector 2-arm branch (resolving `HYPOTHESES.md`'s withdrawal), and the ten-callers-each count, all read directly and repeatedly in this pass / the `+0x268` field range is High-confidence as this pass's own reading but is an explicitly flagged, unreconciled discrepancy against the brief's and review's `+0x264` / Unknown, unchanged, for where or whether a message-type dispatch happens downstream of the queue

### R2-SESSION-005

`allods2.exe`'s client-side inbound path applies the same XOR transform, against the same key table, at the same structural position, as `a2server.exe`'s server-side path; the one call site's own arithmetic does not match the buffer-plus-eight/length-minus-eight reading this experiment's own review gives it. `allods2.exe R2.0145` and `a2server.exe R2.0146` are byte-identical (byteLen 122, instrCount 42, matching `mnemHash`/`normHash`, matching small-immediate operands `16,12,80,8`) — the same source function compiled into both binaries. Both XOR the record in chunks of up to 80 bytes, restarting the key at its own first byte for every chunk, against an 80-byte table: `a2server.exe`'s table is at file offset `0x234968` (VA `L2.00363`, `.data`), `allods2.exe`'s is at file offset `0x201230` (VA `L2.00364`, `.data`) — different addresses, identical content, sha256 `04863743fcdaac1cdb18a121608be367a3f8b3e06c35c1eeecd0c5e466a33c8d` for both (`evidence/xor-key-crossbinary.txt`, address-independent hashing; independently re-verified by a direct read of both file offsets). `R2.0145` has exactly two callers in `allods2.exe`: `R2.0147` (`CLlDriver::SendData`) and `R2.0143` at call site `L2.00365` — the same receive loop `R2-SESSION-003`/`R2-SESSION-004` already read, immediately before the call into `ReceiveData`, mirroring `R2.0146`'s two callers in `a2server.exe` (`R2.0148`/`SendData`, and `R2.0140`'s receive loop, immediately before its own `ReceiveData` call). This experiment's own confidence review reads that call site as passing "the buffer advanced by eight and the length reduced by eight." Ghidra's decompiler did not surface it either way (`R2.0145();`, no visible operands), so this pass read the raw instructions directly instead (`evidence/allods2-xorcall-disasm.txt`, `tools/ghidra/DisasmFn.java`). The call passes exactly two 4-byte arguments (the stack pointer is raised by 8 immediately after the call, which confirms a 2-dword cdecl cleanup). One is unambiguously a pointer: `rec+16` — pure address arithmetic, no dereference — from a one-line accessor `R2.0149` (`this+8`), compounded by the call site's own addition of 8 to the accessor result, where `rec` is this function's own `[+0x26c]` field. The other is the raw 4-byte value stored at `rec+0xa0`, through a second one-line accessor `R2.0150` (`*(this+0xa0)+8`) whose own internal `+8` the call site's own subtraction of 8 from the accessor result exactly cancels, net displacement 0; this pass does not independently establish whether the value stored there is itself a pointer or a length — the review's own separate finding about a differently-scoped consumer function describes a `record+0xa0` field there as holding a decoded length, but that is a different object (a queued `CBufferManager` record, not this function's own per-connection accumulation buffer) and this pass does not confirm the two share a layout. Either way, the review's specific "buffer advanced by eight, length reduced by eight" arithmetic does not match what this pass reads: the two net displacements are 0 and +16, not +8 and −8

**Confidence.** High for "does `allods2.exe` apply an equivalent transform to `a2server.exe`'s" (identical function body, identical key-table content, identical structural call position, all read directly, ruling out both a missing step and a different key on the client side) and for "the review's specific +8/−8 arithmetic does not match the raw disassembly" (this holds regardless of which argument is a pointer and which is a length) / Unknown for which of the two arguments is the pointer and which (if either) is a length, and what either points at or counts

## Master server and configuration

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-006 | `a2server.exe`'s master-server (hat) connection path is unreachable in this build: the gate is permanently false with no writer, and the retry poller never bootstraps a first attempt. | High | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-007 | `a2server.exe` recognises the case-folded key `hataddress` in a 24-entry `srv.cfg` key table, and this pass traces the complete path from that config line to the address `R2.0151` connects to. | High | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-008 | The shipped `srv.cfg` for this lawful install: `Protocol=WSOCK_TCPIP`, `IPAddress=127.0.0.1:8001`, `IPAddress2=127.0.0.1:8002`, `HatAddress=127.0.0.1`, `Save=server`, `ChrBase=chr1\` | High | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |
| R2-SESSION-009 | `Save=server` names the server as the persistence target in the shipped config, but static reading does not establish which side computes simulation state. | High / Unknown | ● active | [EXP-2002](../experiments/EXP-2002-session-model/) |

### R2-SESSION-006

`a2server.exe`'s master-server ("hat") connection path is unreachable in this shipped build, not merely non-fatal on failure — the gate is permanently false with no writer anywhere in the image, and the poller that would retry it never bootstraps a first attempt either. the global at `L2.00366` has exactly four references in the whole image, all reads, no write: `R2.0151` (the gate itself), `R2.0152`, `R2.0153` and `R2.0154` (`evidence/a2server-refto-hatgate.txt`, a complete reference-manager scan, zero orphan hits). The address sits past `a2server.exe`'s `.data` section's raw file data (`VirtualSize 0xe297c` against `SizeOfRawData 0x16c00`, `evidence/a2server.exe.pescan.tsv`) — the loader zero-fills it, so absent a write through a pointer this scan cannot see, the gate (`==4`) is never satisfied. `R2.0151` in fact has two early-return checks, not the one previously described: the the global at `L2.00366` gate, and a second, immediately inside it, that returns success without attempting anything when `[L2.00367]==1` (already connected). Its sole caller, the poller `R2.0155`, sends its 15-second keepalive (`R2.0156`) conditionally — only once `GetTickCount()` has advanced 15,000 ms past a per-object timestamp, not unconditionally as previously stated; this pass's own decompilation shows one visible argument (`1`) to that call, and does not independently confirm a second. Separately from the gate, the poller's own retry-with-backoff arm is itself gated on a nonzero previous-attempt timestamp (`this+0x1d8`) that this function only ever writes as a result of an attempt having already happened — nothing in this function bootstraps a first attempt from the zero state. The backoff (`[L2.00368] * 60000`) and inactivity (`30000`) constants read exactly as previously stated. No call to a process-termination primitive appears in either function, so the H4a alternative (failure is fatal) remains ruled out

**Confidence.** High (the discriminator is now two independently-sufficient reasons the gated path never runs in this build — a permanently-zero, unwritten gate value, and a poller that cannot self-bootstrap — read directly from a complete reference scan and two decompiled functions; this is a stronger, different statement than "non-fatal on failure", which remains true but understates it)

### R2-SESSION-007

`a2server.exe` recognises the case-folded key `hataddress` in a 24-entry `srv.cfg` key table, and this pass traces the complete path from that config line to the address `R2.0151` connects to. The prior negative searched only an exact-case byte sequence; `hataddress`, lower-cased, exists at file-relative address `L2.00369` (len 10), one arm of a chained `[settings]`-section key comparison inside `R2.0157`. Reading that function's control flow directly and counting every `R2.0158(<key>)`-guarded arm chained under the `[settings]` section guard gives 24 distinct key names, not 25 — cross-checked by an independent signal: the section-name literal itself (`len=8`, matching `"settings".length()`) is referenced exactly 24 times in the function, once per guarded arm (`evidence/a2server-srvcfg-keys.txt`, hash16 frequency count). Three more keys (generically: two ban-list keys and one reporting key) are recognised outside this chain, as bare top-level keys, not part of the 24-entry `[settings]` table. The `hataddress` arm assigns global the global at `L2.00370`; `R2.0151` reads that same global at `L2.00371` and passes it as the sole argument to `R2.0159` — `CLlDriver::PrepareForConnect` by address, per `R2-SESSION-002`'s method census. That is the complete, address-verified path from the configuration line to the connect target; the prior claim's own guess (a linear key table, case-folded) was directionally right but its cited globals (the global at `L2.00372`/the global at `L2.00373`) were a misattribution — those hold `ipaddress`/`ipaddress2`, used to build the address this server advertises, not the one it connects to for the hat

**Confidence.** High (a negative bounded to one exact-case search is not evidence of a global absence — the case-insensitive table exists, and every step from config key to connect-target argument is read directly by address; the discriminator this row now answers is "was any mechanism located", not "does one specific mechanism exist". The 24-key count is this pass's own two-signal recount and is published as such; it was not re-derived from, and does not match, a differing count on record elsewhere for the same table)

### R2-SESSION-008

The shipped `srv.cfg` for this lawful install: `Protocol=WSOCK_TCPIP`, `IPAddress=127.0.0.1:8001`, `IPAddress2=127.0.0.1:8002`, `HatAddress=127.0.0.1`, `Save=server`, `ChrBase=chr1\` — a direct key-value read of the shipped configuration file (`evidence/srv-cfg-facts.txt`), not inferred from any binary. These six are the shipped subset; `a2server.exe`'s own `[settings]` key table recognises 24 keys total (`R2-SESSION-007`)

**Confidence.** High (direct read of six named keys from the shipped file, not sampled or inferred)

### R2-SESSION-009

`Save=server` (`R2-SESSION-008`) names the server as the target of persistence for this shipped config, but static reading does not establish which side computes tick/movement/combat/spellcast/item-transfer state, and the two cross-referenced facts available do not discriminate this either. `claims/rom2-engine.md`'s `R2-ENGINE-008` (already published, cited not restated) establishes that a byte-identical-skeleton session `Serialize` routine (mnemHash+instrCount+byteLen match, changed small-immediate operands) exists in `rom.exe`, `allods2.exe` and `a2server.exe` alike — present in all three, so it does not by itself indicate which ROM2 binary calls it live, or on which side persistence is actually driven. `experiments/EXP-2001-engine-sharing/evidence/classification.tsv` classifies two ROM1 movement/terrain anchor functions (`R0437`, `R0438`, 255 instructions each) as `absent` under its own cross-binary instrument — but that same instrument's own already-published correction history (`R2-ENGINE-008`'s partial retraction, `claims/retracted.md`) shows an `absent` classification can mean "this exact byte/mnemonic shape was not found," not "no ROM2 code performs this computation"; a rewritten or differently-compiled equivalent would classify absent under this instrument without meaning the computation is missing

**Confidence.** High for the `Save=server` config value alone (`R2-SESSION-008`) / Unknown for who computes vs. applies tick/movement/combat/spellcast/item-transfer state, which this experiment's static reading does not reach; the two cross-referenced facts are reported as context only, and neither is treated as discriminating between the thin-client, lockstep, or mixed models (H3a/b/c)

## ALM record dispatch

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-010 | The ALM record-header primitive is byte-identical across `rom.exe`, `allods2.exe` and `a2server.exe`, each called once by its own per-record dispatch loop; dispatcher arity is 10 in `rom.exe` against 13 in ROM2. | High | ✔ promoted | [EXP-2002](../experiments/EXP-2002-session-model/) |

### R2-SESSION-010

The shared ALM record-header primitive is a three-way byte-identical match across `rom.exe`, `allods2.exe` and `a2server.exe`, each called exactly once by its own binary's per-record dispatch loop; `rom.exe`'s own dispatcher has an arity of 10 against ROM2's 13, matching `R2-ASSET-003`'s already-published record-type count from the data side. `a2server.exe R2.0160`, `allods2.exe R2.0003` and `rom.exe R0467` share one census signature exactly (byteLen 33, instrCount 14, matching `mnemHash`/`normHash`, matching small immediates `28,1000,20,8,32,60`) — a three-way match, one binary more than this row previously established. Each has exactly one caller in its own image, and in each case that caller is the binary's own ALM per-record dispatch loop (`a2server.exe R2.0161`, `allods2.exe R2.0002`, `rom.exe R0478`; `evidence/{a2server,allods2,rom}-almhdr-callers.txt`, one hit each, zero orphan). Counting each dispatcher's own distinct `COMPUTED_JUMP` targets directly (`tools/ghidra/SwitchCount.java`, which counts jump-table fan-out without recovering case values or requiring decompilation): `allods2.exe` and `a2server.exe` both have 13 (previously established); `rom.exe R0478` has exactly 10 (`evidence/rom-switchcount.txt`, a fresh, narrowly-scoped read of this one function this pass, no other ROM1 fact drawn from it). A dense, zero-based jump table with 10 arms in `rom.exe` and 13 in both ROM2 binaries, on a byte-identical header-parsing primitive, means ROM2's record-type surface (`R2-ASSET-003`'s already-published 13, types `0..12`) carries exactly three record types a ROM1 `.alm` cannot supply — the code-side count matches the data-side count exactly, and it directly answers the brief's question 5 for its code half

**Confidence.** High (the three-way primitive match and the single-caller-from-own-dispatcher pattern are read directly and identically in all three binaries; the arity comparison that was previously deferred is now run, and its result — 10 vs 13, matching the already-published 13-vs-10 data-side count — is what makes this a genuinely new code-side fact rather than only "the two ROM2 binaries agree with each other")

## Save and character files

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-017 | A ROM2 single-player `game*.sav` uses the ROM1 envelope with magic `Bsg&` (0x26677342): same 16-byte header, version 0x0BAD0002, word-codec blob and 256-byte label at the blob end. | High | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |
| R2-SESSION-018 | The ROM2 document serializer reproduces the ROM1 head, Player-list position, world-present byte, world-half order and 400-byte trailer; ROM2 adds one object list after the last world list. | High / Medium | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |
| R2-SESSION-019 | ROM2 session raw data is 10,374 bytes against ROM1's 4,374, and ROM2 terrain is a sparse key list over cell indices 0x807 up to 0xedee (exclusive) plus a sub-object, not the ROM1 cell table. | High / Medium | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |
| R2-SESSION-020 | The ROM2 Player body keeps ROM1's 51-byte scalar prefix widths and adds a 2,560-byte raw block, a 36-byte raw suffix and no Diary; Player grows 112 to 2,688 runtime bytes. | High / Medium | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |
| R2-SESSION-021 | After the label ROM2 keeps the ROM1 `&YA1` store and its key names, then writes a count-prefixed pair list, the `scenario.dll` `ScenarioSave` record and nine length-prefixed blocks in place of ROM1's campaign record. | High / Unknown | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |
| R2-SESSION-022 | A ROM2 `.a2c` character file is a `0x04507989` magic plus six 16-byte-headed sections, scrambled with a per-section 15-bit key and a doubling checksum; both lawful files decode to end of file. | High / Medium | ✔ promoted | [EXP-2005](../experiments/EXP-2005-rom2-save-format/) |

### R2-SESSION-017

`allods2.exe` writes a ROM2 save in `R2.0162` and reads it in `R2.0121`. The writer emits `u32 0x26677342`, a zero `u32`, `u32 0x0BAD0002`, a `u32` compressed length and the blob; it then seeks back and patches the second dword with the blob end. The reader compares the first dword with 0x26677342 (`Invalid save file.` otherwise), discards the second, requires the third to exceed 0x0BAD0001 (`Outdated save file.` otherwise), reads the length and the blob and decodes it. The blob is `u32 outWords` followed by packets: opcode below 0x80 copies that many literal words, opcode 0x80 and above repeats one word `n & 0x7f` times; the encoder (`R2.0163`) starts a run when the next two words are equal. This is the packet grammar of ROM1 `formats/sav/encoding.md`. `rom.exe` holds the immediate 0x26677341 at four sites (two writer and reader pairs plus two slot-list routines); `allods2.exe` holds 0x26677342 at the corresponding four sites and 0x26677341 at none.

The 256-byte label follows the blob at the stored blob end. Two slot-list routines (`R2.0164`, `R2.0165`, reached through data slots `L2.00374` and `L2.00375`) open each `game*.sav`, read eight bytes, compare the magic, seek to the second dword and read 256 bytes. The single-player save routine `R2.0166` calls `R2.0162`, reopens the file, seeks to the end and writes 256 bytes. `R2.0162` itself writes a label only when a flag at save-object offset `+0x74` is set: the 28-character text `Server Multiplayer save file.` padded with zeroes to 256 bytes. Slot names are `game%ld.sav`, `game%d.sav` with zero-padded forms, and `game0000.sav`; the slot lister skips `game9999.sav` and `game9998.sav` in one of its two modes.

**Confidence.** High. The magic and version are immediate operands in the writer and reader; the codec grammar is read from `R2.0167` and `R2.0163`. Alternative D1 (identical container and codec) is correct except for the magic word; alternatives D2 and D3 are excluded for the envelope and codec. No ROM2 save was decoded: this is a static read.

**Unknown.** Whether a multiplayer save (the label path under the flag at `+0x74`) differs in layout; whether the compressed length and the stored blob end are always consistent in a written file; which byte values the label holds beyond the first NUL; the meaning of the zero second header dword on write (the writer patches it after the blob).

### R2-SESSION-018

The document serializer is `R2.0120`, one routine for both directions, called by the save routine `R2.0162` and the load routine `R2.0121`. Its call sequence is compared with ROM1's `R0414` in `rom.exe`. Both start with two `u32`, one `CString` and eleven `u32`, then a mission `u32` and a difficulty `u32` (ROM2's reader stores the difficulty only when it is 1..3). Both then serialize the Player list, a dead-actor list through a virtual slot, a `u8` world-present flag and, when it is nonzero, the world half. Both write two `u32` of 0xBADFACE1 and end with a 400-byte raw block (`R2.0168`). ROM1 reads the second word only when the first equals 0xBADFACE1 (compare at `L08053`); `allods2.exe` has two 0xBADFACE1 immediates, both pushes in the writer, and its reader reads both `u32` unconditionally without comparing either. In ROM2 the world half is five `list32`-style calls in this order: `R2.0169`, `R2.0170`, terrain `R2.0171`, session `R2.0001`, `R2.0172`, followed by `R2.0173`, an extra counted object list with no ROM1 counterpart in `R0414`. The ROM1 order is Buildings, SpellEffects, terrain, session, Sacks.

On load the ROM2 world-half branch constructs the map from the named scenario through the ALM loader (`R2.0002`, 0x3b8-byte object) before reading terrain.

**Confidence.** High for the head, Player-list, flag, trailer and the position of every world-half call (read from both call sequences and the decompiled bodies); Medium for the identity of the three leading and trailing lists (Buildings, SpellEffects, Sacks by position only; their class programmes were not read).

**Unknown.** Whether the second 0xBADFACE1 written by ROM2 is a constant or a stored global; the element class of `R2.0173`; whether its list is empty in a typical single-player save; the semantics of the eleven head `u32`.

### R2-SESSION-019

Session `R2.0001` transfers, in order: 4,000 bytes at session `+0xc784`, 1,000 at `+0xe6c4`, 48 at `+0x08`, 400 at `+0xa728`, 4,908 at `+0xa8bc`, a `u8` at `+0xa48`, a `u8` at `+0xa49`, a `u32` at `+0xa4c` and three `u32` at `+0xbe0c`, `+0xbe10`, `+0xbe14`: 10,374 bytes. ROM1's session block (`formats/sav/world.md`) is 400, 1,000, 48, 400, 2,508, then the same 1 + 1 + 4 + 12: 4,374 bytes. The 4,000 against 400 and 4,908 against 2,508 differences are the only changed widths. 4,908 minus 8 is 4,900, which equals 70 times 70, and ROM1's 2,508 minus 8 is 50 times 50; the 70-by-70 reading is inference. The slot-list reading of the 4,000-byte block as 1,000 signed slots (ROM1: 100 slots in 400 bytes) is inference from width only.

Terrain `R2.0171` writes: for cell index `i` from 0x807 while `i < 0xedee`, if the byte at `terrain + 0x20000 + i` exceeds 0x0f, the `u32` key `(i << 16) | (that byte << 8) | byte(terrain + 0x10000 + i)` is added to a list; the list is written (`R2.0174`), the sub-object at `terrain + 0x54084` is serialized through its virtual slot `+8`, and the terrain's own address follows as a `u32` identity key. On load the keys are scattered back into the two byte planes. The ROM1 cell rows of `2 + 54 * count + 4` bytes do not appear in this routine.

**Confidence.** High for the session raw widths and offsets and for the terrain key formula and index bounds (decompiled body, both directions); Medium for the 70-by-70 and 1,000-slot readings (arithmetic only). The ROM1 comparison reads `formats/sav/world.md`, not ROM1 save bytes.

**Unknown.** The list writer's count form (`R2.0174`); the sub-object's layout; every session field meaning in ROM2.

### R2-SESSION-020

The Player body is the virtual slot `+8` of vtable `L2.00376` (`R2.0085`). It writes: a `CString` from `+0x18`, `u16` `+4`, `u32` `+8`, raw 8 at `+0x10`, `u8` `+0xa44`, `u32` `+0x2c`, `u16` `+0x30`, `u32` `+0x3c` combined by exclusive-or with the constant 0x5c073f4d, `u8` `+0x40`, `u8` `+0x41`, `u32` `+0xa48` combined by exclusive-or with the same constant 0x5c073f4d, `u32` `+0xa50`, two `u16` saturated at 0x7fff from `+0xa54` and `+0xa4c`, `u32` `+0xa5c`, `u32` `+0x38`, the Player's own address as `u32`, then raw 2,560 bytes at `+0x44`. Counting widths this prefix has the same 51 bytes and the same type order as ROM1's (`formats/sav/player.md`); the runtime offsets differ. Suffix order is then the group list `R2.0122` (count `u32`, bodies by `R2.0123`) and raw 36 bytes from `R2.0175`. ROM1's suffix is a group count and bodies, raw 32 bytes and a Diary body; no Diary call appears in the ROM2 routine. The ROM1 and ROM2 class censuses over the MFC runtime-class tables list `Diary` and `Tavern` only in ROM1 and `Inn` and `Pointer` only in ROM2; `Player` is 112 bytes in ROM1 and 2,688 in ROM2 (an increase of 0xa10, which holds the 0xa00 block).

**Confidence.** High for the field order, widths and the missing Diary call (decompiled body, both directions); Medium for the class-census sizes as an object-size measure (read from the allocation size in the class table, not a layout).

**Unknown.** The meaning of the 2,560-byte block (its size equals a `.a2c` section B decoded size, `R2-SESSION-022`; the identity is inference), the 36-byte suffix, and the 51-byte prefix field meanings beyond ROM1's, which ROM2 does not establish.

### R2-SESSION-021

`R2.0166` (the single-player save driver) writes, after the label: the `&YA1` store from `R2.0176` (magic 0x31415926, header of six `u32`, 32-byte records, a pool length `u32` and the pool, `R2.0177`), filled by `R2.0178`, `R2.0179`, `R2.0180` and `R2.0181` from key names that include `Character`, `GameOptions` (`Wimpy`, `ShowHP`, `FlyingHP`, `Formation`, `Speed`, `ShowTimeFlow`), `SpellBook` (`IsOpen`, `Pressed`, `Shortcuts`), `Objects` (`Selection`, `Group0`..`Group9`), `Inventory/IsOpen`, `Projectiles` (`FreeIndex`, `Prj<id>`) and `Fog` (`FirstState`, `Data`). Those are the ROM1 roots of `formats/sav/application.md`. The store is followed by `R2.0182`: a `u32` count and that many pairs of two `u32` (`R2.0183`), then two `u32`; then the `scenario.dll` export `ScenarioSave` through the pointer at `L2.00377`, which `allods2.exe` fills with `GetProcAddress(module, 12)` at `L2.00378` after loading `scenario.dll` by name; then nine blocks from `R2.0184`, each `u16`, `u16`, `u32` length and that many raw bytes.

`ScenarioSave` (`scenario.dll` RVA 0x3585) writes: raw 4,096 bytes from the DLL data at VA `D2.00003`; a 4-by-4 grid of 20-byte cells (320 bytes); a `u32` count from a list; for each list node a `u32`, a `u32` and 16 raw bytes; and a `u32` index of the node equal to a DLL global (zero when none). `ScenarioLoad` (ordinal 13) is the mirror. ROM1's campaign record (`formats/sav/campaign.md`) has a different counted grammar; no ROM1 campaign routine was compared byte for byte here.

**Confidence.** High for the order and widths of the store, the pair list, the nine blocks and the `ScenarioSave` record (decompiled and disassembled bodies); Unknown for every field meaning in them. The claim that the nine blocks replace ROM1's campaign record is Medium-grade inference from position.

**Unknown.** What the pair list, the nine blocks and the 4,096-byte DLL block hold; which `Prj`, `Group` and `Fog` shapes ROM2 uses beyond the key names; whether the key set is equal to ROM1's (names were read, values were not).

### R2-SESSION-022

The writer is `R2.0185` and the reader `R2.0186`. The writer first opens the file for reading and loads any existing sections, then rewrites it with the sections supplied by the caller replacing the old ones. A file is `u32 0x04507989` followed by sections in the order A, B, C, D, E, F, tags `0xAAAAAAAA`, `0x55555555`, `0x40A40A40`, `0xDE0DE0DE`, `0x41392521`, `0x3A5A3A5A`. Each section has a 16-byte header: `u32 tag`, `u32 length`, `u16` zero, `u16 key`, `u32 checksum`. The key is one `R2.0021(0x7fff)` value per section (15 bits). The body (A, B, C, D, F) is scrambled: XOR byte `i` with bits 16..23 of a 32-bit register that starts as `key | key << 16`, shifts left one bit per byte and has `key` OR-ed back in after every 16th byte; the checksum is `c = c * 2 + byte` over the plain bytes (E: over the decoded 0x34 field bytes). A is 52 bytes: `u32 id0`, `u32 id1`, a zero `u32`, a NUL-terminated string within `[0xc, 0x2c)`, and eight more bytes. B decodes through a byte-run codec (`R2.0187`: `u32 outBytes`, then opcode below 0x80 copies that many bytes, 0x80 and above repeats one byte `n & 0x7f` times) to 2,560 bytes. C is a raw block (72 bytes in both files; it holds small nonzero values, 0 to 19, and its stored checksum is 0). D and F are item lists: a 9-byte header whose `u16` at offset 7 is the payload length, then records of a `u16`, a `u8` and a body. E holds 0x34 field bytes in 16 fields (its header length is not its extent: first file header 59, extent 60 = 0x34 plus 8 pad bytes; second file 58 and 58; the extent rule places F correctly in both); each field has a transform chain (XOR or add of constants and earlier fields) and a pad byte from `R2.0021(0xff)` is inserted before a field when its key bit is set (fields use bits 0..8, 14, 13 and 9..13 in offset order; bit 13 pads two fields); the header length is not read.

The file name is `<decimal id0><decimal id1>.a2c` (`%u%u.a2c`); the character-creation routine `R2.0188` fills the pair from `QueryPerformanceCounter` (`timeGetTime` as fallback for the low dword) and clears 0xa00 bytes. The replay `tools/r2a2c/r2a2c.py` on both lawful files: both checksums verify on all six sections, section lengths sum to the file size (510 and 527 bytes), section B decodes to 2,560 bytes in both, the file name equals the id pair in both, and D parses to its stated payload length (three non-empty records in the first file: kinds 12, 6, 8; twelve empty records in the second). `R2.0189` scans `*.a2c` files for a first dword 0x68436c42 (`BlCh`), reads a flags `u32` at file offset 0x4c and deletes the file when bit 0 is set and bit 2 clear, unless a game-mode field at `+0x63c` equals 2; no writer of that variant was found.

**Confidence.** High for the container, scramble, checksum, run codec, section order and the id-pair file name (decompiled writer and reader, replayed on two files). The checksum keeps only the last 32 plain bytes, so a verifying checksum excludes error only there, and C's checksum is 0 over zero tail bytes and does not discriminate; the scramble of A, B and D rests on structure (B decodes to exactly 2,560 bytes, the A id pair equals the file name, D framing consumes its stated length) plus the tail checksum, and C's scramble on its small decoded values; Medium for the item-record framing (three non-empty records replayed) and for the id pair being a performance-counter value (read from one routine); Unknown for every field meaning in A..F, for section B's 2,560 bytes beyond size, and for the `BlCh` variant.

**Unknown.** Section C; section E field meanings; item record kinds; the A bytes at offsets 0x2c and 0x2f (observed 0x00/0x40 and 4 in two files); server-side validation was located but not read: `Cheating detected! Player: %s file: %u%u.a2c` is referenced at `L2.00379` inside `R2.0119`, and `Character is too strong for this map.` at `L2.00380` and `L2.00381` in a different routine.


## Initial campaign bank and first continuation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-023 | In the matching ROM2 DLL code, NewGame returns a cleared 1,024-DWORD bank except slot 768=10; selected first Leave paths have explicit bank and availability transitions. | High | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |

### R2-SESSION-023

NewGame export 11 at D2.00004 clears 4096 bytes at bank base D2.00003, rebuilds the pointer catalog, clears availability, adds/enters the initial type-2 ID-1 record and initializes two grids outside the bank. The complete selected export and its catalog/list helpers leave bank slot 768=10 and every other slot zero. This applies at DLL return. Later client character-choice code writes slots 776/781 from flag bits 0x40/0x80. R2-SESSION-021 supplies the bank's raw scenario-save block.

The selected initial EnterInn/TalkTo kind-3 unlock changes availability rather than the bank. Type-2 ID-1 Leave removes the initial record and clears current, with movie output remaining minus one. Mission Enter clears slots 752..767, as R2-ENGINE-048 establishes.

Ordinary mission-10 Leave sets slot 773 to zero. For i=0..19, nonzero 532+i becomes 1 when incoming 512+i is zero or 2 otherwise; zero 532+i stays zero. Every 512+i becomes zero. It removes the current record, clears current and sets completed slot 906 to 1. The established ID-10 case adds type-1 ID-20 availability and outputs 1. The common divisible-by-ten rule advances slot 768 by ten; incoming 10 becomes 20. A nonzero slot 775 first invokes a separate available-list restoration branch, so the resulting list must not be inferred under arbitrary incoming bank state.

**Confidence.** High for the bounded DLL post-export state and stated conditional first-transition algebra. The full NewGame span, catalog/list calls, bank-address translation, type-2 branch, ordinary loop, selected ID jump-table slot and stage/return branches are preserved with exact image/range hashes. The matching raw DLL text/data sections are one code dependency, not independent locale confirmation. No global bank-writer absence is asserted.

**Unknown.** Complete whole-client post-NewGame state, every script/network writer, arbitrary restored catalog values, complete party carryover and owner-save instances. The states are static derivations, not executed snapshots.

## Native campaign record and selected client tail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-031 | The native scenario save/load prefix transfers 4096 raw bank bytes and sixteen 20-byte grid cells before the counted list and index; field meanings remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-032 | ScenarioSave writes each availability node as ID, kind and raw 16 bytes, then a zero-default index matched by current-object pointer. | High / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-033 | The ordinary ScenarioLoad path selects catalog pointers by kind/ID and deduplicates pointers; it does not copy incoming per-record raw 16 bytes into catalog objects. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-034 | The native current-index reader walks by equality and dereferences its resulting cursor without a local range/null check; complete invalid-input responses remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-035 | Selected EN/RU pair tails transfer a u32 count, two u32 per 60-byte runtime row and two final u32; the other row fields are outside those pair helpers. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-036 | Nine selected EN/RU client tail cells transfer two u16, a u32 byte length and buffer bytes; nonzero incoming length controls reader allocation and raw I/O. | High / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |
| R2-SESSION-037 | Native EN/RU ordinal-12/13 bindings connect the selected client tails to the scenario DLL; the selected reader gates that tail on client mode 2 and a nonzero argument. | High / Unknown | ● active (branch candidate) | [EXP-2015](../experiments/EXP-2015-rom2-campaign-save/) |

### R2-SESSION-031

ScenarioSave at D2.00102 and ScenarioLoad at D2.00103 are complete export bodies.
Their first transfers are 4096 bytes at D2.00003 and a nested 4-by-4 loop over
D2.00104, with row stride 80, cell stride 20 and 20 bytes per cell. These are raw
source/destination addresses, not derived bank values or reconstructed grid
fields. The list count and final index each transfer four bytes. R2-SESSION-032
and R2-SESSION-033 distinguish record writing from restoration.

**Confidence.** High for explicit byte counts, order, addresses and loop algebra;
Unknown for field meanings and producer completeness. The complete paired bodies,
internal transfer boundaries and section-based range hashes rule out a different
prefix width on these paths. Matching DLL text/data and selected range hashes are
one shared dependency, not independent locale evidence. The instrument is Python
3.12.10, pefile 2024.8.26 and Capstone 5.0.7 in x86 32-bit mode, not a function
inventory or a whole-image census.

**Unknown.** Bank and grid producers, arbitrary values, hidden reference meanings,
file-I/O response and complete live-state coverage. Native reads/writes were not
executed.

### R2-SESSION-032

The availability export returns list base D2.00105. The current export returns
the object pointer at D2.00007. ScenarioSave reads the list count at runtime +12,
walks nodes from list +4 through node +0, and writes each node's object pointer
payload through record helper D2.00106. That helper transfers runtime +4 as a
u32 ID, runtime +0 as a u32 kind and 16 bytes at runtime +8. The raw-address
helper D2.00107 returns its input address unchanged. Record width is 24 bytes.

The writer initializes the output index to zero. Each pointer equality with
D2.00007 replaces it with that node's zero-based index; no match leaves zero.
The pointer itself is not emitted by this scenario-record helper.

**Confidence.** High for the exact local program, widths, linked-list relation and
pointer equality; Unknown for raw payload meaning and unmatched-current policy.
Complete writer, record, count/iterator and export ranges exclude an ordinal-only
interpretation of the writer's matching operation. The bounded manifest does not
exclude other native producers or identity relations elsewhere.

**Unknown.** Catalog-object payload producers, live raw values, repeated object
pointers outside the selected load path, and whether an unmatched current pointer
is a reachable normal state.

### R2-SESSION-033

After the raw prefix and incoming count, ScenarioLoad clears availability and
calls catalog rebuild D2.00005. That complete routine clears catalog list D2.00032
and makes 49 append requests with distinct fixed object addresses. Successful
membership still depends on the excluded allocation path. Their object constructors
and current payload values are not established by these addresses.

Each incoming 24-byte record is read into a temporary by D2.00108. Its kind
getter reads +0 and its ID getter reads +4. Kind equal to 1 calls D2.00014;
every other kind calls D2.00031, whose catalog filter requires kind 2. Each filter
scans the catalog for matching kind and ID, tests the selected object's pointer
against availability through D2.00021/D2.00109, and appends only an absent pointer.
The comparator compares pointer DWORDs. The filter can append every matching
catalog pointer; no unique kind/ID assumption is required.

Only the temporary kind and ID feed these filters. Its incoming raw 16 bytes
are not copied into the selected catalog objects by the complete ScenarioLoad
body or these filters. Duplicates and unmatched keys can change the rebuilt count
and therefore the relation between an incoming ordinal and a rebuilt ordinal.

**Confidence.** High for kind routing, pointer deduplication and the absence of a
raw-payload transfer in this complete bounded data-flow program. Medium for full
restored membership, because node allocation, catalog-object producers and I/O
are excluded dependencies. Unknown for their failure/default responses. This
bounded result excludes a verbatim per-record reconstruction model on the selected
path; it does not exclude live catalog values or derived/default catalog payloads.

**Unknown.** Catalog payload sources and mutation, allocation/destruction behavior,
unrecognized kinds and IDs as valid game states, invalid-input responses and any
full live-World restoration relation.

### R2-SESSION-034

The complete ScenarioSave body defaults current index to zero; a first-node match
also produces zero. The complete ScenarioLoad body reads that DWORD, takes the
availability head cursor and starts a counter at zero. It advances while the
counter differs from the incoming index. At equality it calls list-at D2.00110,
which computes cursor +8, dereferences that location and stores the resulting
object pointer at D2.00007. List-next D2.00111 similarly dereferences its incoming
cursor while advancing.

No local count/range/null comparison precedes those cursor dereferences in the
complete endpoint and selected iterator bodies. With a nonempty rebuilt list,
index zero selects its first object. The file alone cannot distinguish this from
a writer's unmatched-pointer default. Empty lists, excessive indices and malformed
counts were not supplied to the native code.

**Confidence.** High for the exact local selection program and bounded absence of
those tests; Unknown for complete invalid-input response. The population is the
whole endpoint and its count/first/next/at helpers, decoded through all internal
transfers. Indirect file-I/O exceptions and excluded allocation code can stop the
path earlier. This is not a universal no-validation claim or an executed crash
claim.

**Unknown.** Native read/exception handling, invalid counts/indices, null cursor
response, whole-client recovery and valid-state current identity after a list
whose incoming records normalize differently.

### R2-SESSION-035

Selected pair writers are EN L2.00382 and RU R2.0182; readers are EN L2.00383
and RU L2.00384. They use an array object at pair-object +36. Its count helper reads +8;
the index-address helpers EN L2.00385 and RU L2.00386 return
DWORD[arrayObject + 4] + index * 60. The +4 field holds the row-buffer pointer.
Writer/reader record helpers transfer two DWORDs at that row's
+0 and +4. The remaining 52 runtime bytes are outside those helpers. The tail
ends with DWORDs at pair-object +4 and +8.

The reader takes the incoming count through the array-resize path and then reads
the two fields into each selected row. Resizing retains, constructs or removes
rows through unexpanded row helpers. Its current field values cannot be inferred
from serialized width or the old row allocation alone.

**Confidence.** High for widths, runtime addressing and selected field-copy data
flow in both measured clients. Medium for complete array membership under valid
allocation and row-helper assumptions; Unknown for remaining field values and
invalid-input responses. Paired full helper bodies and index arithmetic support
the 60-byte runtime width independently of the eight-byte stored record width.
No global absence or record semantics are inferred from this ratio.

**Unknown.** Pair-field meanings, other row fields, row constructors, source
membership producers, generic allocation/error behavior and full-world identity.

### R2-SESSION-036

Writers EN L2.00387 and RU R2.0184 transfer WORDs at cell +0 and +2, a DWORD
length at +4, then length raw bytes from the pointer at +8 when length is nonzero.
The pointer is not emitted. Readers EN L2.00388 and RU L2.00389 first branch on
old length: nonzero old length frees a nonnull old pointer and clears length and
pointer. They read the two WORDs and incoming length, then allocate and read raw
bytes when incoming length is nonzero. The cells are 12 runtime bytes apart.

The selected client loops visit exactly nine cells, starting at client +0x4dc
in EN and +0x4f4 in RU. Their local transfer program has no payload grammar or
length-limit comparison. Generic allocation and virtual file I/O remain outside
this read; their checks and failure responses are not excluded.

**Confidence.** High for the complete local helpers, loop bound, byte transfer and
buffer-versus-pointer relation. Unknown for field meanings, native validation,
buffer producers and accepted length population. Native helpers were not executed;
this is not evidence that arbitrary lengths are accepted.

**Unknown.** Buffer content grammar, sources, allocation/overflow/short-read
behavior, zero-length cells with inconsistent prior pointer state and downstream
consumers.

### R2-SESSION-037

Positive native LoadLibraryA sequences use the scenario.dll literal and store the
module handle at EN L2.00390 and RU L2.00391. GetProcAddress with ordinals 12 and
13 reads those same handle slots and writes EN function slots L2.00392/L2.00393
and RU slots L2.00377/L2.00394. DLL exports identify ScenarioSave and ScenarioLoad
at D2.00102 and D2.00103.

The selected writer tail calls pair-save, ScenarioSave and the nine-cell loop in
that order after the registry-store writer. The measured call to ScenarioSave is
EN L2.00395 / RU L2.00396. The selected reader tail similarly calls pair-load,
ScenarioLoad and the nine-cell loop; the measured scenario call is EN L2.00397 /
RU L2.00398. Reader gates require client DWORD +0x5d8 in EN or +0x63c in RU to
equal 2 and its selected argument to be nonzero. Pair-object offsets are +0x598
and +0x5f8 respectively.

**Confidence.** High for the positive native module/ordinal/slot chain, selected
calls, gate and order. Unknown for other driver arms, registry-store values and
complete source production. Binding anchors are not full binder or all-call
proof. Distinct client hashes/addresses are separate measurements; shared DLL
code is not a second independent scenario implementation.

**Unknown.** Whole-driver control flow, other save/load routes, complete key-value
production, native file acceptance and post-load gameplay/presentation. No full
World serializer or networking claim is made.

## Campaign departure bank transitions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-047 | Every ordinary departure zeroes slot 773, normalizes 532..551 against 512..531, zeroes 512..531, removes and clears current and stores 1 at slot 896+ID, all before any ID case. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-048 | Per-ID bank stores: 20 sets 532, 552; 40 sets 773=23; 50 sets 771, 534, 554; 60 sets 537, 557, 774; 70 sets 535, 555, 777, 770; 80 sets 778; 100 sets 538, 558; only the output 3 store is gated. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-049 | Slot 768 gains ten when the incoming ID is divisible by ten, whether or not the ID has a case body, and the stage switch runs only when the resulting slot 768 is 30, 40, ..., 110. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-050 | Stage cases and ID cases store immediate values into an eight-entry, 20-byte-stride array at D2.00112..D2.00007, outside the bank; the field meaning and consumers are Unknown. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-051 | Ordinary departure reads bank slots 512+i, 532+i, 768, 772, 775..781 as gates and writes 512..551, 552, 554, 555, 557, 558, 768, 770, 771, 773..775, 777, 778 and 896+ID; no other bank slot is written in the measured routine. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-052 | The store of 1 at slot 896+ID in the measured ordinary routine has no bound check on the ID argument, so an ID of 128 or above, or a negative ID, addresses outside the 1,024-DWORD bank. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-SESSION-059 | EN text/town.txt holds each of the sections #npc22talk30 and #npc2108talk31 once; their keys match the stage-30 kind-3 packed entries NPC 22 topic 30 and NPC 2108 topic 31. | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |
| R2-SESSION-060 | In EN, all 4 EnterInn call windows (call sites L2.00399, L2.00400, L2.00401, L2.00402) contain a compare of client member +0x5d8 with 2; the RU gate was not measured. | Low | ● active (amended) | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |

### R2-SESSION-047

Entry of the ordinary routine (D2.00012) stores zero to D2.00113 (slot 773) before it tests slot 775. A nonzero slot 775 runs the restoration of R2-ENGINE-149 and joins the common prefix at D2.00114. The prefix loops i=0..19: a nonzero slot 532+i becomes 2 when slot 512+i is nonzero and 1 when it is zero; a zero slot 532+i is unchanged; slot 512+i is then set to zero in every case. The routine calls the find helper for current, calls the remove helper only when a node is found, stores zero to current (D2.00007) and stores 1 to D2.00115+4*ID (slot 896+ID). The ID switch at D2.00024 follows. A case store to slot 532 or later therefore replaces the loop result for that slot.

R2-SESSION-023 states the same prefix for incoming ID 10. This claim states it for every ordinary ID.

**Confidence.** High: the instructions D2.00012..D2.00116 and D2.00114..D2.00117 are read to the switch. The matching EN and RU raw sections are one code dependency, not independent locale confirmation.

**Unknown.** Bank contents at the call, the other writers of these slots and the readers of slot 896+ID.

### R2-SESSION-048

branch-matrix.tsv rows for ordinary departure, with slot N the DWORD at D2.00003+4N: ID 20 stores slot 532=1 and slot 552=2 after its optional mission-21 add. ID 40 stores slot 773=23. ID 50 stores slot 771=1, slot 534=1 and slot 554=3, after an optional append gated by slot 780. ID 60 stores slot 537=1, slot 557=2 and slot 774=1. ID 70 stores slot 535=1 and slot 555=3, then slot 777=1 and slot 770=1 on both paths of its output gate (slot 777 zero and slot 778 nonzero stores output 3; every other combination skips only that store). ID 80 stores slot 778=1 on both paths of its output gate (slot 778 zero and slot 777 nonzero stores output 3). ID 100 stores slot 538=1 and slot 558=2. IDs 10, 30, 31, 90 and 110 and every ID without a case body store no bank slot in the case.

The case-to-slot relation is not arithmetic: IDs 20, 50, 60, 70 and 100 map to slots 532, 534, 537, 535 and 538.

**Confidence.** High for the immediate stores and their gates as read in the case bodies. Slot meanings are not resolved.

**Unknown.** Readers of these slots, the meaning of the values 1, 2, 3 and 23, and any later writer.

### R2-SESSION-049

At D2.00026 the routine divides the ID argument (the frame slot `+0xc` off the frame base) by ten. A zero remainder adds ten to slot 768 (D2.00118). The test reads the argument, not the switch target, so an ID without a case body that is divisible by ten still advances slot 768. The routine then subtracts 30 from slot 768 and, when the unsigned result exceeds 80, returns. Otherwise the byte table D2.00028 and the dword table D2.00029 select a stage case; values that are not multiples of ten inside 30..110 select the return target.

R2-SESSION-023 gives the first case, incoming ID 10 raising slot 768 from 10 to 20. No claim covers which IDs occur in play or how slot 768 changes outside this routine.

**Confidence.** High for the conditional rule as read in the instructions D2.00026..D2.00027.

**Unknown.** Other writers of slot 768 and the values the client presents.

### R2-SESSION-050

Entry n of the array is at D2.00112+0x14*n. Stage case 30 stores 10000 at field +4 of entries 0..3; stage 40 stores 22000 at the same fields. Stage cases 50..110 store 60000, 150000, 400000, 800000, 1500000, 5000000 and 10000000 at field +4 of entries 0..7. Stage 90 also stores 12000 and stage 100 stores 40000 at field +0 of entries 0, 1 and 3. ID cases 30, 40, 60, 80 and 90 store immediate words at field +0x10 of entries 0, 1 and 3. ID 50 loops over entries 4..7: field +0 is 499 except 0 for entry 6, field +8 is 100 for entries 4 and 7 and 20 for entries 5 and 6, field +0xC is 2 for entries 4 and 7 and 1 for entries 5 and 6, and field +0x10 receives one immediate word per entry. The array ends at the current pointer (D2.00007).

R2-SESSION-023 places two grids outside the bank. This claim lists the stores of the departure routine and does not name the structure.

**Confidence.** High for the stores listed, read as immediates or as stores of locals that the same loop sets from immediates under the entry-index tests recorded in branch-matrix.tsv.

**Unknown.** Field meaning, consumers, why entry 2 is skipped in the stage 90 and 100 loops, and the client side.

### R2-SESSION-051

The gate reads are slot 512+i and 532+i (prefix loop), 768 (stage switch), 772 (ID 20), 775 (restoration), 776 and 781 (restoration tail), 777 and 778 (IDs 70 and 80), 779 (ID 110) and 780 (ID 50). The writes are 512+i and 532+i through the prefix loop, 552, 554, 555, 557 and 558 and slots 532, 534, 535, 537 and 538 through the ID cases, 768, 770, 771, 773, 774, 775 (zero in restoration), 777, 778 and 896+ID. The restoration tail is selected by slot 776 and slot 781: both nonzero appends record D2.00074, 776 only appends D2.00075, 781 only appends D2.00076 and neither appends D2.00077; every arm then appends five common records, so only four records differ by arm.

R2-SESSION-023 notes that the client writes slots 776 and 781 after character choice. This claim states which slots the departure routine reads.

**Confidence.** Medium: the set is read from the call-reachable instructions of one routine and the generated matrix. The loop writes are index-based, so the list is a summary of the matrix and the prefix, not a per-address proof of each loop slot.

**Unknown.** The writers of slots 772, 775, 776, 779, 780 and 781 and the identity of the records appended by the restoration.

### R2-SESSION-052

At D2.00119 the routine loads the ID argument and at D2.00120 stores 1 to [edx*4+D2.00115] with no compare of the ID. The ID compare that follows (ID minus 10 against 100) selects the switch and does not guard this store. The bank spans D2.00003..D2.00032, so ID 127 addresses slot 1023, ID 128 addresses the first DWORD of the catalog list object at D2.00032 and a negative ID addresses lower memory. IDs of 10 to 110 are the ones with case bodies in branch-matrix.tsv.

**Confidence.** Medium: the missing compare is a local instruction fact. No claim covers which IDs the client passes or how the objects after the bank are laid out.

**Unknown.** The runtime ID range and the effect of an out-of-range store.

### R2-SESSION-059

Instrument: byte search of the EN main.res text/town.txt member for section headers. #npc22talk30 occurs once (4471 bytes, hash prefix 5effd57da3cc, 24 tag lines, parts 1..16) and #npc2108talk31 occurs once (3421 bytes, hash prefix 39c832e02925, 17 tag lines, parts 1..8). Both carry the flag words IAMFEMALE, IAMFIGHTER, IAMMAGE and IAMMALE as part selectors. The client tag parser compares the keywords iammale and NPC at 3 string cells per locale (EN L2.00403, L2.00404, L2.00405; RU L2.00406, L2.00407, L2.00408). The header key #npc<NPC>talk<topic> equals the packed entries of stage 30 in the EnterInn case body; this is an inferred key scheme, not a read of the consumer.

**Confidence.** Medium: section counts and sizes are measured. The key scheme is an inference from two matches, and no code that builds the header string was read. RU town.txt was not read. Register-based indirect calls remain unresolved in 14 of the 20 distinct new client bodies, so no decoded client body is a closed control-flow graph. Method: the experiment exceeded its preregistered inputs cap (20 reported and 24 conservative against 16), so this claim is bounded by an incomplete method record.

No negative control exists for the key scheme: no header that fails the scheme was searched, so the match is two positive matches only and Medium is the ceiling.

**Unknown.** Which stage the client passes for ID 2, the RU sections, how a part is chosen from the flag words, and the choice-action producer, which was not located.

### R2-SESSION-060

Instrument: raw scan for calls through the EnterInn import cell L2.00409 (ordinal 9) in EN allods2.exe finds 4 call sites in bodies of 161, 489, 466 and 528 instructions; each body contains a compare of member +0x5d8 with 2. The RU scan finds 4 call sites (L2.00410, L2.00411, L2.00412, L2.00413) in bodies with the same instruction counts, but the gate detector did not match there; that is a detector result, not evidence of absence. The detector hard-coded member +0x5d8. A later reading of RU windows reports the comparison of the member at offset +0x63c with the constant 2 at L2.00414, L2.00415 and L2.00416; that reading is not a generator product and the generator was not rerun with the RU offset. Two TALK call windows per locale were decoded (EN L2.00417, L2.00418; RU L2.00419, L2.00420); 3 further TalkTo call sites per locale were excluded as outside the adjacent TALK handlers.

**Confidence.** Low: window and gate extraction is a pattern match, the compare is an instruction fact only, and the RU encoding was not read by hand. Register-based indirect calls remain unresolved in 14 of the 20 distinct new client bodies, so no decoded client body is a closed control-flow graph. Method: the experiment exceeded its preregistered inputs cap (20 reported and 24 conservative against 16), so this claim is bounded by an incomplete method record.

**Unknown.** The RU gate (offset +0x63c reported by reading, not tabulated), which of the 4 windows serves ID 2, the meaning of member +0x5d8, and the 3 excluded TalkTo call sites per locale.

**Amended.** R2-SESSION-072 measures the member compare and unequal bypass in the selected first RU caller only. Its other three RU windows and member semantics remain Unknown.

## Selected initialization local flags and unresolved insertion field

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-067 | Selected initialization stores, through its symbolic saved entry receiver, +0x74 from signed argument<2, +0x70 from argument>0 and +0x1a8 from the fresh +0x74; tested value 2 gives local values 0, 1 and 0. | High | ● active | [EXP-2021](../experiments/EXP-2021-rom2-startup-gate/) |
| R2-SESSION-068 | Session DWORD+0x174 before selected map-actor insertion remains Unknown: local initialization stores, loaded-pointer aliases and opaque/indirect boundaries prove no composed numeric value. | Unknown | ● active | [EXP-2021](../experiments/EXP-2021-rom2-startup-gate/) |

### R2-SESSION-067

In reused initialization EN L2.00221 / RU R2.0010, signed less-than and greater-than comparisons, with their boolean results zero-extended, store DWORD+0x74 as argument<2 and DWORD+0x70 as argument>0. The following +0x74 test stores DWORD+0x1a8=0 for zero and 1 otherwise. If the tested stack argument is 2, these local stores produce 0, 1 and 0 respectively. The receiver is reloaded from the saved entry receiver at frame offset -0x150. The offsets are relative to that symbolic saved entry receiver; the claim does not establish a Session that persists through the earlier opaque and indirect calls. R2-ENGINE-153 supplies the verified wrapper's literal argument 2; these fields are distinct from DWORD+0x174.

**Confidence.** High for the decoded signed predicates, store widths and local branch values in both complete 197-instruction bodies. There is no intervening call between each comparison and its store or between the +0x70 store and +0x1a8 branch.

**Unknown.** Actual stack values and saved-receiver persistence after earlier opaque calls, player-visible mode, ordinary-start entry and the numeric +0x174 value remain Unknown. The locally tested argument is not inferred from a function name.

### R2-SESSION-068

R2-ENGINE-167's explicit indexed stores stop at byte +0x16f and its complete selected initializer has no memory displacement +0x174. R2-ENGINE-168's selected first target writes a fixed global coordinate after an opaque direct call. Neither establishes the live Session field unchanged. The initializer's indirect call EN L2.00421 / RU L2.00422 and the unexpanded eligible sites permit unknown effects: the first target's own call EN L2.00343 / RU L2.00344 and the call with receiver Session+0x78 at EN L2.00345 / RU L2.00347. The saved receiver is also stored to a fixed global DWORD (EN L2.00331 / RU L2.00332), whose relation to the Session is unproved. Its pointer-mediated store EN L2.00423 / RU L2.00424 uses [Session+0x7c]+0x0c, whose alias relation is unproved.

R2-ENGINE-153's wrapper calls the map caller only on initialization return zero. R2-ENGINE-156's insertion condition includes `(Session DWORD+0x174==0 || owner DWORD+0x2c==0)`. The selected initializer's lookup failure arm locally returns 2 and its accepted continuation returns 0; actual lookup/callback results and intervening map-path effects remain unread. No composed numeric +0x174 value joins those selected gates.

**Confidence.** Unknown for the numeric field before insertion and positive ordinary-start reachability. Eight distinct native input keys and one new paired body do not exclude the live external, indirect or alias alternatives. The bounded local observations do not establish a global writer census.

**Unknown.** Incoming +0x174, actual stack/field values, runtime receiver/global/pointee aliases, external and indirect mutations, callback effects, ordinary player-start identity and accepted actors.

## Selected first EnterInn caller

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-071 | The selected EN/RU EnterInn setup pushes frame-0x14 then frame-0x94 and leaves ECX at frame-0x94; no scalar catalog or stage value is produced in that local setup. | Medium | ● active | [EXP-2022](../experiments/EXP-2022-rom2-entry-stage/) |
| R2-SESSION-072 | In the selected first EnterInn caller, EN member+0x5d8 and RU member+0x63c are compared with 2; the unequal direct branch bypasses the local call setup. | Medium | ● active | [EXP-2022](../experiments/EXP-2022-rom2-entry-stage/) |

### R2-SESSION-071

The selected calls are EN L2.00399 and RU L2.00410; their EnterInn identity is
inherited from the published locator, not a new import-cell or table read.
Their final local fallthrough blocks calculate and push frame-0x14, then frame-0x94. ECX holds
the second address at the call instruction. No scalar catalog or stage value
is produced by this setup. Two pushes do not establish full callee arity or
a receiver convention.

The exact enclosing ranges, EN L2.00425..L2.00426 and RU
L2.00427..L2.00428, match their recorded 12-digit hash prefixes and each hold
161 decoded instructions. Each has one unread indirect switch. The earlier
selector feeds separate direct calls whose effects are unread; no direct
return or argument edge produces the terminal frame addresses. No direct branch
in either admitted range targets the setup interior or call; unread indirect
destinations still prevent an all-path proof. No callee
body or switch table was opened on a guessed stage model.

**Confidence.** Medium for the bounded local setup: exact native ranges and
hashes reproduce the locator and synthetic controls reject scalar, caller
argument, call-return and alias counterexamples. This is one dependency
measured in two locales, not two independent witnesses. Both complete byte
ranges have unresolved CFG edges; no body-level pairing is claimed. Six
install keys were charged before access, with one generated listing display
additionally retained as a retrospective conservative helper receipt.

**Unknown.** The buffer contract and contents, full ABI, stage source,
catalog/stage relationship, effects across unread callees and aliases,
indirect destinations, all-path entry into the setup, runtime reachability,
roster and player route.

### R2-SESSION-072

The final local compares are EN L2.00429 and RU L2.00415. They compare a
member of a pointer loaded from frame-0xb0, the same local receiver slot, with 2, at
locale-specific offsets+0x5d8 and+0x63c. Their unequal branches at EN
L2.00430 and RU L2.00431 go to L2.00432 and L2.00433, bypassing the local
EnterInn setup. Earlier compares in the same complete caller occur at EN
L2.00434 and RU L2.00414. This extends R2-SESSION-060's unmeasured RU clause
for the selected first caller only. The slot receives entry ECX at EN L2.00435
and RU L2.00436; the final pointer loads are EN L2.00437 and RU L2.00438.

**Confidence.** Medium for these local instruction/branch facts. Both exact
selected ranges were read completely and direct branch targets are verified
instruction boundaries. The indirect switch remains unresolved; the other
three RU EnterInn windows and intervening callee effects were not read.

**Unknown.** The member's semantic name, the values it takes at runtime,
changes across calls or aliases, the other three RU windows, and whether this
path serves any particular catalog record or stage.

## Four original save-point corpus

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-075 | Four frozen original ROM2 SAVs have consistent Bsg&/0x0bad0002 envelopes and exact codec output; decoded lengths are 4486 for A/C and 72178 for B/D. | High | ● active (branch candidate) | [EXP-2023](../experiments/EXP-2023-rom2-owner-save-corpus/) |
| R2-SESSION-076 | The four measured physical tails tile exactly from label end through inline REG, one pair, a scenario record and nine zero-length buffers; raw payload meanings remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2023](../experiments/EXP-2023-rom2-owner-save-corpus/) |
| R2-SESSION-077 | A/C store mission 0/current kind 2 ID 1; B/D store mission 10/current kind 1 ID 10, supporting bounded city-like/mission-like keys without a proven native current pointer. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2023](../experiments/EXP-2023-rom2-owner-save-corpus/) |
| R2-SESSION-078 | Player-list counts are 1/8/1/8; the first Player in each file has Group count 1, where unsupported ROM2 Group/actor programmes leave active-party membership Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2023](../experiments/EXP-2023-rom2-owner-save-corpus/) |
| R2-SESSION-079 | A/C decoded documents are byte-equal despite campaign-tail differences; B/D's parsed first Player states match while head/bank/opaque bytes differ, without causal or actor-state interpretation. | High / Unknown | ● active (branch candidate) | [EXP-2023](../experiments/EXP-2023-rom2-owner-save-corpus/) |

### R2-SESSION-075

Four stable original snapshots have physical sizes A=6348, B=26886, C=6372 and
D=26503. The separate frozen prereg precedes every original content access.
Each capture streams SHA256 and an exclusive external copy after a flushed
before-read charge; original size/mtime before and after agree. The copied
revisions verify against those hashes before each parser run. Symbols have no
filename role or temporal meaning.

Python 3.12 and the bounded standard-library probe read the 16-byte envelope.
All four magic words are 0x26677342, versions are 0x0bad0002 and stored blob
ends equal 16 plus blob lengths. Blob lengths are 794, 21024, 794 and 20753.
Source-bounded literal/run decoding consumes every compressed byte and emits
4486, 72178, 4486 and 72178 bytes, exactly twice each declared word count.
No native executable, other save or asset is an input. Four capture and four
frozen parse keys use 8 of the 16-key cap; replays reuse those revisions.

**Confidence.** High for the four captured revisions' numeric framing and extent
agreement. Independent literal/run, truncated packet, output underflow/overflow,
blob-end and dual-domain byte-coverage controls discriminate a permissive scan
or discarded suffix. This is bounded structural decoding, not native execution.

**Unknown.** Whole document meaning, native load/acceptance, changed future
original revisions, multiplayer scope and save-point chronology.

### R2-SESSION-076

At the exact end of each envelope's 256-byte label, the inline store magic is
0x31415926. Store sizes are A/C=742, B=1050 and D=938; record counts are 22/28/22/28
and pool lengths are 10/126/10/14. R2-SESSION-021 and R2-ASSET-006 supply the
framing. No magic search identifies this boundary.

The following pair count is 1 and its row is (1,1), followed by two zero DWORDs.
Every scenario record has 4096 bank bytes, sixteen raw 20-byte grid cells,
a counted list of 24-byte ID/kind/raw16 records, and one current index.
Availability counts are 1/1/2/1. Every index is zero. Nine following cells each
have two zero WORDs and a zero byte length; the ninth ends exactly at EOF.
Bank offsets are 1828/22366/1828/21983. Physical coverage partitions every input
byte; raw extents remain explicitly distinguished from understood fields.

**Confidence.** High for exact physical order, counts and offsets on these four
revisions, with paired prior writer/reader grammar and independent synthetic
pool-length/file-truncation negatives. Unknown for raw payload semantics.
Native malformed-file responses are not measured by these controls.

**Unknown.** Registry values and references, pair meanings, grid cells,
availability raw16 payloads, general buffer grammar and full live restoration.

### R2-SESSION-077

A/C have document mission 0 and first incoming availability kind 2/ID 1.
B/D have document mission 10 and their sole incoming record kind 1/ID 10.
All current indices are zero. C additionally holds kind 1/ID 10 availability.
The exact nonzero banks are A:{768=10}, B:{754=1,768=10,769=1},
C/D:{768=10,769=1}; every other bank slot is zero in these revisions.

The comparison classifies A/C as city-like stored keys and B/D as mission-like
stored keys. R2-SESSION-023's selected initial kind-2/ID-1 and ordinary mission
paths and R2-ENGINE-048's location-ID relation support this bounded interpretation.
R2-SESSION-032..034 distinguish incoming index zero from a proven current-pointer
match and from native catalog-restored membership. No original code was run.

**Confidence.** High for literal header/record/bank values. Medium for the bounded
city-like/mission-like interpretation through the selected prior location grammar.
Unknown for the native current pointer and the owner's unmapped descriptions.
Synthetic mismatched kind/ID/mission and out-of-list index vectors remain Unknown;
label changes cannot influence classification.

**Unknown.** Which dialogue fired, filename-to-owner-description roles,
chronology, causality, catalog payload restoration, actual UI mode and universal
campaign schedules. Equal stage slot 768 does not establish equal complete state.

### R2-SESSION-078

The Player-list counts are A/C=1 and B/D=8. Only the first reference is reached:
new schema-1 class Player consumes shared class index 1, then object index 2 under shared archive-index framing
(SAV-ARCHREL-253, SAV-758). Those indices are protocol inference, not stored
index words. R2-SESSION-020 supplies its scalar widths, raw2560 block and following Group count.
Every first Player has Group count 1. The parsed decoded prefixes end at A/C=2707
and B/D=2708; the remaining 1779/69470 bytes are explicitly opaque.

The first Player name is represented by length/hash only. Its raw2560 hash and
saved self key 254852752 and stored +0x38 key 262270312 agree across the four.
No actor target is parsed or resolved. Archive object index, saved-address key,
list position and active-client role are separate questions. ROM1 Group, Token,
Unit and Human programmes and field meanings were not imported into ROM2.

**Confidence.** High for the literal first class/schema, prefix, counts, key
values and stopping offsets. First class/object indices are inferred through
the shared archive framing, corroborated only for this first object. Unknown for actor membership/state. Independent opaque fake
presence, edge, same-offset-key and relocation changes cannot produce a resolved
party. Unsupported class and unproved plain-index reference controls stop rather
than resynchronize. These test bounded reporting, not a ROM2 actor programme.

**Unknown.** Inline Group body, subsequent Players, actor references, hero/current
party, owned/equipped objects, stats, progression, world-present discriminator,
world objects, session, trailer and alignment inside the opaque remainder.

### R2-SESSION-079

The generator holds all four decoded documents from the individually ledgered
frozen reads and compares them directly in memory. A/C's 4486 bytes are equal,
although physical labels, inline-store hashes, bank slot 769 and availability
differ. Their decoded
SHA256 is 5224b9f1440e00b16dfa021486b8cceefd39d0523dea9f57dc03208c06556dcc.
B/D have equal parsed first Player state; first head DWORD B=11 and D=1 and
bank slot 754 B=1 and D=0 differ. Their opaque document hashes and inline-store
extents/hashes also differ. All six pairs are represented in the delta table.

Across all four, the first Player stored +0xa44 is A/C=1 versus B/D=2 and +0x41
is A/C=0 versus B/D=1. The name length/hash, raw2560 hash and other admitted
scalar values agree. The XOR-unmasked +0x3c and +0xa48 values are 1000 and 0;
no ROM2 gold, hero or entry-latch meaning is assigned.

**Confidence.** High for byte equality, literal parsed differences and hash
comparisons on the frozen population. A synthetic one-byte bank754 loss must
produce the exact bank.754 delta; an independent expected-value control kills
a wrapper mutation that erases the returned bank value. Unknown for opaque/actor-state meaning.

**Unknown.** The causes and order of differences, equal live gameplay state,
actor inventory/equipment/stats/progression and interpretation of opaque bytes.

## Four frozen first-Group boundaries

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-SESSION-083 | Walking the published first-Player prefix in four frozen revisions fits group count 1 and Group starts at decoded offsets 2707 or 2708; member offset, extent and identity remain Unknown. | Medium | ✔ promoted | [EXP-2024](../experiments/EXP-2024-rom2-group-serialization/) |

### R2-SESSION-083

The exact frozen inputs A, B, C and D are verified by SHA256 from the same
buffers consumed by the instrument. Envelope lengths and source-bounded word
codec lengths agree. Their decoded sizes are 4486, 72178, 4486 and 72178.
Only the first serialized Player is walked, through its matching descriptor
digest, CString, published 51 scalar bytes and raw 2560 bytes.

A and C fit group count 1 at decoded offset 2703 and a first Group start at
2707. B and D fit count 1 at 2704 and start at 2708. The head CString length
explains the one-byte offset difference. A and C have equal decoded-document
hashes despite distinct physical revision hashes. This equality establishes
no time, filename, dialogue or owner-observation mapping. R2-SESSION-075,
R2-SESSION-078 and R2-SESSION-079 already report framing, count 1, the prefix
stopping offsets and A/C document equality. This result independently
reproduces those values and adds their link to the selected RU native inline
Group dispatch; count 1 is not a new corpus discovery.

**Confidence.** Medium for the conditional corpus fit. The scalar programme
is published authority; the selected RU native list confirms direct inline
Group dispatch. Both fresh replays reproduce each measured boundary and input
hash. The parser stops before the first unresolved embedded Group programme.
The offset/count fit and local dispatch supply no native identity relation
or complete Group encoding authority. Instrument sanity checks add no native
member-grammar evidence.

**Unknown.** The embedded Group programmes, first member offset, complete
Group extent and remaining Players; which Player is current; actual hero or
party; original acceptance, actor/World LOAD and post-load behavior.
