# Projectile trajectories

[Reference](format.md)

## Projectile trajectories

Settled by `MAGIC-BOLTGATE-069`, `MAGIC-BOLTSHAPE-070` (partially retracted and superseded for interpolation), `MAGIC-BOLTLIST-071` (superseded in part in the ledger),
`MAGIC-BOLTSTILL-072` (lifetime scope narrowed), `MAGIC-TRAIL-073`, `MAGIC-BOLTEND-074`. Section 12 gives the
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

**The raw source is a bounded random walk; the drawn points are quadratic samples.**
MAGIC-BOLTSHAPE-070's unqualified no-interpolation clause and -0.5 smoothing
attribution are partially retracted. The exact local passes are MAGIC-275
through MAGIC-278 and MAGIC-282. Generation begins at the projectile display
point and points toward the resolved target display point.

```
initial sign       +1 for odd rand(), -1 for even
attempt            v=(rand()%7)*sign; d=stored_0.01*(rand()%50)
admission          abs(v)>2 and stored d>stored_0.15
accepted step      clamp(q+v,-3,3); t+=d; append; sign*=-1
walk stop          t>=stored_0.7
raw end            if last t>=1 remove last pair; append (1,0)
knots              first two raw points, then midpoint/current for each interior point,
                   then the terminal endpoint
sampling           overlapping triples 0,1,2 then 2,3,4; quadratic y=(A*x+B)*x+C;
                   x starts at triple first.x, adds 6, remains strictly below third.x
acceptance         every sampled abs(y-Ay)<=stored_0.15*length
retry              clear and rebuild all points, with the continued random stream
rotation           accepted samples only; __ftol truncation, then low-word stores
clipping           later per-stamp screen rectangle; no point-list clipping
```

The -0.5 operations form midpoint knots as previous-(current-previous)*(-0.5),
in that instruction order. Six units are canonical horizontal spacing, not
arc length. The terminal endpoint is excluded from sampling; it is a fit
constraint. Rotation/truncation can change adjacent integer spacing and make
duplicate points. The sampled band is inclusive, not the former strict bound.
Every arithmetic/store boundary uses the rules below. — MAGIC-276, MAGIC-277,
MAGIC-278

**The list is rebuilt every tick and does not accumulate.**

```
picture 34   resolve actiontarget; generate one set with tag 34;
             SetSize(generatedCount) and overwrite  -> the list is REPLACED
             an unresolvable target leaves the previous list in place
picture 36   SetSize(0)  -> the list is CLEARED
             then, per entry of the projectile's own word array at +0xac (count +0xb0):
               resolve the id, generate one set with tag = (index mod 7), append
```

The generator uses the full source-to-target segment as its constraint; no
argument carries a partial length. A resolved live call produces the complete
sampled figure, whose terminal fit endpoint is excluded. Native first-draw
and cache-refresh order remain Unknown. — MAGIC-277, MAGIC-281 The `+0xac` array is copied word by word
onto the projectile by the cast spawner from the caster's own array of the same offsets.

`MAGIC-SPRAY-134`…`MAGIC-SPRAY-137` close that array's producer. After applying the selected list, simulation opcode `0x8a`
or `0x8c` sends `count+1` words: caster identity or packed source cell first, then every selected
victim id in list order. The matching client arm clears the source's embedded word array, appends
those ids, writes its count and primary, and only then reaches the spawner copy above. The previously
proposed `L05035` `SetSize` site belongs to a different list operation.

**These two projectiles do not move.** The driver computes a per-axis step before its switch, but
only the default arm applies it. Pictures 34 and 36 take an arm that writes the sheet frame and
nothing else, so `x`, `y` and `z` keep the spawner's values for the object's whole life and
normal caster construction's `actionsegments = 13` is a countdown, not a
division of a distance. Direct and loaded routes have their supplied/saved
counters; MAGIC-BOLTSTILL-072's lifetime scope is narrowed by MAGIC-281. A moving target is
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

**Customisation limits.** The two picture ids, both comparisons, the six draw arms, the eleven shape
constants, the two moduli, the tag stride 5, the 8-pixel centring, the trail cap 6 and the 13-step
frame ramp are all `.text` or `.rdata`. A third spell with a drawn path, a straight bolt, a longer
trail or a differently sized bolt sheet all require an engine change; none is reachable from
`Data.bin` or `projectiles.reg`.

## Light cast by spell objects

The client CProjectile light method deposits into a viewport vertex grid for
five pictures: 10 Fire Arrow, 12 Fire Ball flight, 13 its explosion,
34 Lightning and 36 Prismatic Spray. The client view rebuild calls it in
the primary and deferred object stores. This complete local selector is
not an enumeration of other object classes. The rebuild has a third
point-stamp call over a second store; attribution of that radius-1 source
is Unknown. — MAGIC-269

Lightning and Prismatic Spray visit every 8-byte drawn-path record. Set
`c = signedPointX >> 5` and obtain `r` from the original screen-to-row
helper, which accounts for terrain projection. Admit `0 <= c <= visCols`
and `0 <= r <= visRows+4`. With stamp stride `visCols+7`, write the four
vertices `(c+3,r+3)`, `(c+4,r+3)`, `(c+3,r+4)`, `(c+4,r+4)` to
`u8(10 * phase)`. Adjacent path cells share vertices. Empty or off-view
paths deposit nothing. Each store overwrites previous light; it does not
add or take a minimum. This is light along the drawn path, with no radial
halo around the stationary projectile. — MAGIC-270

Fire Arrow uses radius 0 and stamp 16. Fire Ball flight uses radius 1
and stamp 16. The explosion uses this table. — MAGIC-271

| Explosion phase | Radius | Stamp |
|---:|---:|---:|
| 0 | 1 | 16 |
| 1 | 2 | 8 |
| 2,3 | 3 | 0 |
| 4 | 3 | 8 |
| 5 | 3 | 16 |
| 6 | 3 | 24 |
| 7 | 3 | 32 |
| 8 | 3 | 40 |
| 9,10 | 3 | 46 |

For nonnegative `i,j <= radius`, a point stamp accepts
`i*i+j*j < radius*(radius+1)`, with threshold 1 at radius 0. After
subtracting scroll, the reflected vertex offsets from the unpadded source
cell are X=`4+i` or `3-i`, Y=`4+j` or `3-j`. Each accepted vertex gets
the same byte, clipped separately to the viewport. There is no radial
falloff. Unclipped radii 0,1,2,3 touch 4,12,32,52 vertices. — MAGIC-271

The client rebuild copies the previous stamp when `L10964 == 2`, then
clears the stamp grid to 255 before objects deposit again and rebuilds
unit light. Under normal caster construction,
Lightning and Prismatic Spray have 13 successful driver steps with phases
`4,3,2,1,0,1,2,1,0,1,2,3,4` and stamps
`40,30,20,10,0,10,20,10,0,10,20,30,40`. The following zero-countdown call
returns finished; the shared updater unlinks and destroys the object.
The direct `0x8b` message route supplies `actionsegments` from its message
counter: 5 for spells 13 and 14, starting `actionphase=-1`, giving
`0,4,3,2,1`, then removal. The former direct sequence in MAGIC-272 is
partially retracted; MAGIC-281 measures its original initializer.
A loaded object resumes from its saved `actionsegments` and `actionphase`.
Fire Arrow and Fire Ball retain their fixed levels while present; the
normal explosion receives 22 ticks and follows its two-tick sheet clock.
Normal caster construction uses distance-derived counters `dist/200` for
Fire Arrow and `dist/384` for Fire Ball flight; a direct message uses its
supplied counter. Each successful driver call decrements the counter.
Exact visible frame ordering remains Unknown.
— MAGIC-272 (partially retracted for the direct sequence), MAGIC-281, MAGIC-DELIVER-035, MAGIC-274

Dynamic lighting gates point deposits. Object animations gates their
rectangle invalidation. Neither flag is read by the direct Lightning or
Prismatic stores, and their unit-grid merge still runs with both flags
off. With `L10964 & 2` clear, Dynamic lighting gates the terrain-light
pass. With that bit clear and Dynamic lighting off, bolt light reaches
units but not the ground. The `L10964 & 2` set branch passes the stamp
through the inspected terrain dispatch/call paths without that local
Lighting test. The local branch and fallback are High. Native activation
and the visible result of the bit-set branch are Unknown:
TERR-FAMILY-187 found no enabling writer in its bounded address-form
search, and startup writes 0.
Native pixels and outer option/frame gates remain Unknown.
— MAGIC-273 (hardware-mode label partially retracted), TERR-FAMILY-187

The unit-grid merge maps each corner to `max(stamp-32,0)`, substitutes
ambient for an unstamped 255 corner, sums and divides by 16, and caps
the result at `ambient >> 2`. Four unstamped corners leave the initialized
or cell-bit value. Ordinary object/unit consumers use this grid, while
special unit passes retain their separate behavior. Bit-clear terrain uses
`min(stamp,terrainByte)` at stamped corners. The bit-set branch uses the
stamp directly, falling back to the terrain byte for an unstamped corner.
Neither deposit changes the persistent terrain plane.
— MAGIC-273 (hardware-mode label partially retracted)

The measured light is client draw state. The direct Prj SAV program saves
16 projectile source scalars, including phase and countdown, and omits the
viewport light grids and drawn-path array. Absence of independent light
persistence outside this program is Medium. Identical light and random
geometry after native LOAD of a live Lightning or Prismatic object remain
Unknown; a live-object SAV and original LOAD/post-load observation are
required. — MAGIC-274

## Exact bolt evaluation inputs

Both inspected producer calls pass Ax,Ay,Bx,By,tag in that order:

```
Ax=s32(P[50]); Ay=s32wrap(P[54]-P[10]-P[68])
Bx=s32(T[60]); By=s32wrap(T[64]-T[68]-T[10])
tag34=34; tag36=originalVictimIndex%7
```

Each SUB wraps to 32 bits before FILD. FILD dword to FSTP qword is exact.
These are displayed fields; MAGIC-261/MAGIC-265's raw construction points
require the intervening geometry/cache steps. The figure starts at the
projectile end. — MAGIC-275

For finite signed-i32 endpoint differences and masked precision exceptions,
length uses this ordered scaled hypot. R64 rounds to a 64-bit x87
significand, nearest; S53 is a separate binary64 store:

```
m=max(abs(dx),abs(dy)); if m==0 return +0
a=S53(R64(abs(dx)/m)); b=S53(R64(abs(dy)/m))
q=S53(R64(R64(a*a)+R64(b*b)))
h=S53(R64(sqrt(q)))
(fm,em)=frexp(m); (fh,eh)=frexp(h)
p=S53(R64(fm*fh)); ep=encodedExponent(p)-1022
n=ep+em+eh
replace p.top16 with (p.top16&0x800f)|((n+1022)<<4)
store binary64 length; restore incoming CW; reload length
```

Normal frexp changes exponent bits to 1022 and returns oldExponent-1022.
Hypot installs CW133f. FSQRT temporarily uses CW037f, then restores CW133f;
the length return restores caller CW. A binary64 sqrt(dx*dx+dy*dy) can erase
the original spill/rounding distinctions. — MAGIC-282

Let R apply the caller's PC/RC to each arithmetic instruction and S perform
a binary64 store under its RC. No reassociation is implied:

```
dx=S(R(Bx-Ax)); dy=S(R(By-Ay))
H=S(R(Ax+L)); D=S(R(H-Ax))
verticalScale=S(R(L*stored_0.03)); band=S(R(L*stored_0.15))
Px=S(R(R(t*D)+Ax))
Py=S(R(R(R(t*0)+R(q*verticalScale))+Ay))
Mx=S(R(previousX-R(R(Px-previousX)*stored_minus_half)))
My=S(R(previousY-R(R(Py-previousY)*stored_minus_half)))
c=S(R(dx/L)); s=S(R(dy/L))
u=R(X[i]-X[0]); v=R(Y[i]-Y[0])
rx=R(R(R(c*u)+X[0])-R(s*v))
ry=R(R(R(c*v)+R(s*u))+Y[0])
```

The first generated Y can differ from Ay by floating residue. The rotation
uses Y[0]. __ftol saves CW, sets RC bits to truncation, FISTPs i64 and
restores CW. Coordinates keep only the low 16 bits, then the drawer
sign-extends them. — MAGIC-275, MAGIC-277, MAGIC-282

The quadratic coefficient program below uses input (x0,y0,x1,y1,x2,y2).
Every arithmetic instruction has a separate R; FCHS changes the sign exactly.
S marks the ten original binary64 FSTP stores, including the output stores.
At PC53 nearest in the measured finite population, those stores preserve
the arithmetic results. — MAGIC-277,
MAGIC-282

```
t1=R(y1*x0)
t2=R(y1-y0)
t3=R(x1*y0)
t4=R(y2*x1)
t5=R(t1-t3)
t6=R(x2*y1)
t7=R(y0-y2)
t8=R(x0*x0)
t9=R(t4-t6)
t10=R(x1*x1)
t11=R(x2*x2)
t12=R(x0-x2)
t13=R(x2-x1)
t14=R(x2*y0)
t15=R(y2*x0)
t16=R(y2-y1)
t17=R(x1-x0)
t18=R(t12*S(t10))
t19=R(t13*S(t8))
t20=R(t5*S(t11))
t21=R(t9*S(t8))
t22=R(t14-t15)
t23=R(x2*S(t2))
t24=R(x1*S(t7))
t25=R(t17*S(t11))
t26=R(t18+t19)
t27=R(S(t11)*S(t2))
t28=R(t22*S(t10))
t29=R(t23+t24)
t30=R(S(t10)*S(t7))
t31=R(t20+t21)
t32=R(x0*S(t16))
t33=R(t27+t30)
t34=R(S(t8)*S(t16))
t35=R(t26+t25)
t36=R(t29+t32)
t37=R(t31+t28)
t38=R(t33+t34)
t39=R(t36/S(t35))
t40=R(t37/S(t35))
t41=R(t38/S(t35))
t42=-t39
t43=-t40
A=S(t42); B=S(t41); C=S(t43)
heldX=first.x
while heldX<third.x:
    storedX=S(heldX)
    storedY=S(R(R(R(R(A*heldX)+B)*heldX)+C))
    append(storedX,storedY)
    heldX=R(storedX-stored_minus_6)
```

Constants are binary64: -6=c018000000000000, 0.15=3fc3333333333333,
0.03=3f9eb851eb851eb8, 1=3ff0000000000000, -1=bff0000000000000,
0.01=3f847ae147ae147b, 2=4000000000000000, 3=4008000000000000,
-3=c008000000000000, 0.7=3fe6666666666666, -0.5=bfe0000000000000.
The CW at native bolt entry remains Unknown; prior startup CW027f does not
capture that entry. Native fragment controls distinguish PC64 and three
rounding alternatives from PC53 nearest by changed integer points.
— MAGIC-282

The random seed is current-thread TLS block+14, not a bolt field. Initial
seed is 1; rand advances state=214013*state+2531011 modulo 2^32 and returns
(state>>16)&0x7fff. Three named reseeds take timeGetTime or time(0). Heal,
Drain, music, AI and range-wrapper consumers can share this thread. The
bounded direct-call population has 88 rand sites, including one preserved
orphan, plus 44 range-wrapper and one float-wrapper caller. The executed
interval between two bolt ticks is Unknown. A pre-generation seed is enough
for one isolated figure with endpoints and CW; cast/object/tick identity
alone is insufficient. — MAGIC-279

The bolt walk is placed on the main thread by exclusion (Medium): its owner
is reached only through stored pointers, and no launched secondary thread
reaches it by direct calls and classified computed calls. The server loop,
the only secondary thread that reaches rand by direct calls, has a launcher
with no reference. — MAGIC-283, SESS-082

| Construction route | Initial actionphase | Successful calls | Phase sequence |
|---|---:|---:|---|
| Normal caster 34/36 | 0 | 13 | 4,3,2,1,0,1,2,1,0,1,2,3,4 |
| Direct 0x8b, spells 13/14 | -1 | 5 | 0,4,3,2,1 |
| Source-cell 0x8c, picture 36 | -1 | 13 | 0,4,3,2,1,0,1,2,1,0,1,2,3 |
| Loaded Prj | saved | saved positive counter | Ramp when incremented actionphase is 1..13; otherwise retained phase |

Every live action-1 driver call invokes geometry after increment/phase
selection and before counter decrement. The next zero-counter call returns
finished. Missing targets can prevent a fresh set; picture 34 retains its
old list, picture 36 clears before resolved links are appended. Exact
native spawn/update/draw/removal frame order remains Unknown. — MAGIC-281

The drawer supplies one own-table forward .16a stamp per stored point, in
list order, with no interpolation. Normal insertion joins the software
painter's projectile collection: ascending hash buckets, node chains, then
ascending points. Later stamps blend with current earlier output and clip
each rectangle separately. Native clip/stride/tables, packing, current
collection and pre-draw framebuffer are additional reproduction inputs.
— MAGIC-280
