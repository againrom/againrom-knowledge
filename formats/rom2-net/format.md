<a id="rom2-session-wire-frame--header-survey"></a>

# ROM2 session record header

The client and server receive an eight-byte header before each socket payload.
The four named fields below have the same header-relative offsets in both
programs. Payload and opcode grammar remain Unknown. — R2-SESSION-003

<a id="result"></a>

## Header

| Offset | Type | Field | Rule |
|---:|---|---|---|
| 0 | u16 LE | Payload length | Socket receive continuation accepts 1..142 |
| 2 | raw 2 | Unknown | No meaning assigned by this reference |
| 4 | u8 | Codec selector | Zero skips decoding |
| 5 | u16 LE | Codec input size | Passed as the codec input-length argument |
| 7 | u8 | Passthrough | Copied into the post-decode record; not read by the codec |

The DirectPlay handler accepts total lengths 8..150. Its maximum is
`8+142`; its total-length test does not establish that a zero payload passes
the socket path. The codec call and returned decoded size are bounded by
142 bytes. — R2-SESSION-003

<a id="not-yet-surveyed"></a>

## Receive sequence

1. Accumulate the eight header bytes.
2. Apply the transport's length rule and read the payload.
3. If the codec selector is nonzero, pass the codec input size to decoding.
4. Preserve the passthrough byte in the record sent to the next queue.

No complete sending/encoding procedure is established: the two unnamed
header bytes, codec grammar, queue payload, possible opcode fields and any
ROM1 counterpart remain Unknown. — R2-SESSION-003
