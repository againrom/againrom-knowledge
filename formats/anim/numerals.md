# Damage numerals

[Reference](format.md)

## Damage numerals

A landed blow's client arm (opcode `0x73`) builds a floating numeral when the damage-number display option is enabled
(`view+0xaa4`, default on, toggled by Ctrl+L) and the victim's health is dropping; a `0x24`-byte record per numeral is merged rather than
duplicated on a repeat hit (`ANIM-NUM-020`). The merge key is `+0x1c` (victim drawable) **and**
`+0x20`, which is the campaign's own sub-tick counter (`SESS-TICK-026`), fetched at construction
through the campaign-object idiom `SESS-OBJ-001` names (`ANIM-071`). Two blows on the same victim
merge into one growing numeral only while the client's local sub-tick mirror has not advanced
between them — not merely while both remain inside the numeral's 1000 ms display window, which is
a separate, longer-lived clock. A merge changes only `+0x04` (the displayed number) and re-formats
the string in the same step; birth time, drift, colour and both keys are untouched (`ANIM-071`):

```
found in the list      +0x04 += new damage; re-format "%d" immediately         ANIM-071
not found in the list   append (growable array, capacity doubles, cap 0x400)   ANIM-071
```

**Damage notifications.** `R0561` is called from melee/general
strike, Building strike, both area direct-damage applies, and the general
effect applier's Token-8 branch (`ANIM-BLOW-019`, whose band-operand and
call-free state-sync clauses are withdrawn; `ANIM-074`). Poison
produces Token8 (`MAGIC-POISONINPUT-157`); an item equip/use producer remains
Unknown. First attachment and continuous ticking call base Effect virtual
`+0x40`, reaching the Token-8 HP body (`MAGIC-POISONREFRESH-158`,
`MAGIC-POISONPHASE-159`). Nonzero computed damage calls the notification
routine; zero skips both the health write and notification. The no-Poison-
notification inference in `ANIM-075` is retracted. This local call does not
establish a drawable numeral or client timing. Fire Wall uses direct damage;
missile and melee use the strike/countdown dispatch. Other indirect
notification paths remain Unknown.

**Pacing.** Lifetime uses `timeGetTime() - birth > 1000 ms`; drift is a fixed `-2`
screen units per `0x401` tick (`ANIM-NUM-020`). Slower nominal cadence therefore covers
fewer ticks and less total drift:

```
speed (ticks/s)   8     10    12    14    16    20    24    28    32
ms/tick           125   100   83    71    62    50    41    35    31
nominal ticks     8-9   10-11 12-13 14-15 16-17 20-21 24-25 28-29 32-33
```

These are nominal alignment counts, not realised bounds: `ANIM-PACE-017`'s catch-up
loop can move actual counts in either direction after a stall (`ANIM-073`). The earlier
"a third as far" aside has no named reference speed; slowest/default gives 1:2,
slowest/fastest gives 1:4, and slowest/24-tps gives 1:3 at the nominal lower counts.

**Message ordering.** `R0509` pops command messages from `L00625` and
dispatches on their opcode byte. Its shared post-dispatch tail compares the handled
opcode with the caller's requested low byte and returns 1 on a match. Empty-queue
polling also depends on that byte and the run bit; an exceptional queue-discard branch
returns 0. The normal pacer runs simulation, pumps for `0x64`, then separately posts
`0x401` to the window-message dispatcher `R0333` (`ANIM-072`, `ANIM-CLOCK-001`,
`AI-CURSOR-177`). Matching completion applies a sub-tick broadcast before presentation.
The transport of server `0x64`/`0x73` broadcasts into this queue remains untraced here.
The paused path separately calls the pump with zero and issues a guarded `0x402` post;
the five idle outcomes and paused command work are in `SESS-IDLE-019` and
`SESS-PAUSE-020` (`ANIM-072`).
