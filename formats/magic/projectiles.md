# Projectile trajectories

[Reference](format.md)

## Projectile trajectories

Settled by `MAGIC-BOLTGATE-069`, `MAGIC-BOLTSHAPE-070`, `MAGIC-BOLTLIST-071` (superseded in part in the ledger),
`MAGIC-BOLTSTILL-072`, `MAGIC-TRAIL-073`, `MAGIC-BOLTEND-074`. Section 12 gives the
object and its lifetime; this section gives what is drawn between the two ends of a flight, which
is per-picture and is not the object's own sprite for two of the seven ids.

**Two picture ids draw a path and no other does.** The projectile calls its own `vt+0x50` once per
tick, on the path it takes when `actionsegments` is still non-zero and its `action` byte is 1, the
value the cast spawner writes. That routine's whole body is behind two comparisons on
the picture: 34 (`lightnin`, Lightning) and 36 (`chain`, Prismatic Spray). Every other picture
returns having touched nothing, and the geometry generator is reachable from nowhere else in the
image.

**A path is a list of 8-byte records held on the object.** The list is a `CArray` at object
`+0x110`: data pointer `+0x114`, count `+0x118`, capacity `+0x11c`.

```
+0x00  i16 x        screen coordinate
+0x02  i16 y        screen coordinate
+0x04  i16 zero     no reader located
+0x06  u8  tag      34 for picture 34; the chain-link index mod 7 for picture 36
+0x07  u8  zero
```

**The geometry is a bounded random walk.** Given the two endpoints, the generator takes their
distance, builds the shape on a horizontal segment of that length, then rotates every point onto
the real direction. The walk itself:

```
deflection step   (rand() % 7) * randomSign,  applied only when its magnitude is >= 2
abscissa step     (rand() % 50) * 0.01,       applied only when it is >= 0.15
deflection sum    clamped to [-3, +3]
step scale        multiplied by -1 each accepted step
walk ends         when the abscissa passes 0.7
smoothing         one pass over every second point, factor -0.5
acceptance        every point's distance from the straight line must be < 0.15 * length;
                  on failure the whole walk is generated again with the same arguments
```

So the drawn figure is ragged, different on every call, and confined to a lens of half-width
`0.15 * length` about the straight caster-to-target line. None of these numbers is data: the ten
constants are `.rdata` doubles and the two moduli are immediates.

**The list is rebuilt every tick and does not accumulate.**

```
picture 34   resolve actiontarget; generate one set with tag 34;
             SetSize(generatedCount) and overwrite  -> the list is REPLACED
             an unresolvable target leaves the previous list in place
picture 36   SetSize(0)  -> the list is CLEARED
             then, per entry of the projectile's own word array at +0xac (count +0xb0):
               resolve the id, generate one set with tag = (index mod 7), append
```

The generator always spans the whole source-to-target segment, and no argument carries a partial
length, so the whole figure exists from the first tick. The `+0xac` array is copied word by word
onto the projectile by the cast spawner from the caster's own array of the same offsets.

`MAGIC-SPRAY-134`…`MAGIC-SPRAY-137` close that array's producer. After applying the selected list, simulation opcode `0x8a`
or `0x8c` sends `count+1` words: caster identity or packed source cell first, then every selected
victim id in list order. The matching client arm clears the source's embedded word array, appends
those ids, writes its count and primary, and only then reaches the spawner copy above. The previously
proposed `L05035` `SetSize` site belongs to a different list operation.

**These two projectiles do not move.** The driver computes a per-axis step before its switch, but
only the default arm applies it. Pictures 34 and 36 take an arm that writes the sheet frame and
nothing else, so `x`, `y` and `z` keep the spawner's values for the object's whole life and
`actionsegments = 13` is a countdown of ticks, not a division of a distance. A moving target is
followed by regenerating the figure onto its new position, not by travelling toward it.

**Nothing is created at the end.** On the tick `actionsegments` is already 0 the driver returns
before its switch and before the `vt+0x50` call; that path allocates nothing and sends nothing.
The list is a member of the object and ceases with it. Neither spell has a second picture:
`2*spellId + 9` gives 35 and 37, and no `projectiles.reg` row defines either on either root.

**The Fire Arrow and Fire Ball trail is the opposite mechanism.** Those two pictures move normally
and keep a `CObArray` at object `+0x138` (data `+0x13c`, count `+0x140`) of **past** positions,
each packed as `(y << 16) or x` and appended after the object has been moved for that tick. When
the count reaches **6** the oldest entry is removed first. So that trail accumulates, is bounded at
six, and expires one entry at a time — where a bolt's list is wholesale and vanishes with the
object.

**Customisation limits.** The two picture ids, both comparisons, the six draw arms, the ten shape
constants, the two moduli, the tag stride 5, the 8-pixel centring, the trail cap 6 and the 13-step
frame ramp are all `.text` or `.rdata`. A third spell with a drawn path, a straight bolt, a longer
trail or a differently sized bolt sheet all require an engine change; none is reachable from
`Data.bin` or `projectiles.reg`.
