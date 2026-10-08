# Claim method index

Each claim ID in the ledgers carries one label for how its finding was acquired. The label is computed from the claim text and the claims that share its experiments. It is not a confidence grade.

- **observation.** The claim's own row and card cite no code location. The finding comes from runtime observation or from reading installed files.
- **static analysis.** The claim's own row or card cites a code location: a routine name, an address inside a game executable, or a partial address pattern. An address counts in any written form: a long hex number, a six- or seven-digit hex number that lies in an executable section and has address context (a hex letter, a tool query, a file path or a neighbouring address), or the little-endian bytes of one. The finding rests at least in part on reading the executable's code.
- **mixed.** The claim cites no code location itself, but an experiment cited in its Evidence also produced a claim that does. Observation and code reading were both applied to the same question.

Retracted claims keep their label and show their ledger status.

Claims labelled: 3005. Observation: 95. Static analysis: 2485. Mixed: 425.

## Ledger ai

| ID | Method | Status |
|---|---|---|
| AI-354 | static analysis | ● active |
| AI-355 | static analysis | ● active |
| AI-356 | static analysis | ● active |
| AI-357 | static analysis | ● active |
| AI-349 | static analysis | ● active |
| AI-350 | static analysis | ● active |
| AI-351 | static analysis | ● active |
| AI-352 | static analysis | ● active (amended, partially retracted) |
| AI-353 | static analysis | ● active |
| AI-ACTIVITY-324 | static analysis | ✔ promoted |
| AI-FORMOWNER-314 | static analysis | ✔ promoted |
| AI-FORMCMD-315 | static analysis | ✔ promoted |
| AI-FORMTRIGGER-316 | static analysis | ✔ promoted |
| AI-FORMACTIVE-317 | static analysis | ✔ promoted |
| AI-STRUCTUSE-306 | static analysis | ✔ promoted |
| AI-SPELLIDENT-286 | static analysis | ✔ promoted |
| AI-SPELLPOP-287 | static analysis | ✔ promoted |
| AI-SPELLCAP-288 | static analysis | ✔ promoted |
| AI-SPELLGUARD-289 | static analysis | ✔ promoted |
| AI-SPELLITEM-290 | static analysis | ✔ promoted |
| AI-QUICKASSIGN-278 | static analysis | ● active |
| AI-QUICKINVOKE-279 | static analysis | ● active |
| AI-QUICKOWNER-280 | static analysis | ● active |
| AI-QUICKSAVE-281 | static analysis | ● active |
| AI-RETREAT-270 | static analysis | ● active |
| AI-RETREAT-271 | static analysis | ● active |
| AI-RETREAT-272 | static analysis | ● active |
| AI-RETREAT-273 | static analysis | ● active |
| AI-RETREAT-274 | static analysis | ● active |
| AI-RETREAT-275 | static analysis | ● active |
| AI-SPRAY-266 | static analysis | ● active |
| AI-SPRAY-267 | static analysis | ● active |
| AI-INPUT-121 | static analysis | ● active |
| AI-SELECT-122 | mixed | ● active |
| AI-PANEL-123 | static analysis | ● active |
| AI-MINIMAP-124 | mixed | ● active |
| AI-KEY-125 | static analysis | ● active (amended) |
| AI-CURSOR-126 | static analysis | ● active |
| AI-INPUT-127 | mixed | ● active |
| AI-FILTER-001 | static analysis | ● active |
| AI-ACQUIRE-002 | static analysis | ● active |
| AI-LIST-003 | static analysis | ● active |
| AI-DIPLO-004 | static analysis | ● active (amended, superseded) |
| AI-DIPLO-005 | static analysis | ● active |
| AI-SIGHT-006 | static analysis | ● active (amended, partially retracted, superseded, contested) |
| AI-GUARD-007 | static analysis | ● active |
| AI-TICK-008 | static analysis | ● active (contested) |
| AI-GROUP-009 | static analysis | ● active |
| AI-ORDER-010 | static analysis | ● active |
| AI-STATE-011 | static analysis | ● active (amended, superseded) |
| AI-GUARD-012 | static analysis | ● active (amended, partially retracted, superseded) |
| AI-PATROL-013 | static analysis | ● active (partially retracted) |
| AI-RADIUS-014 | static analysis | ● active (amended, superseded, contested) |
| AI-AUTHOR-015 | static analysis | ● active (amended, superseded) |
| AI-DIFF-016 | static analysis | ● active |
| AI-PATROL-017 | static analysis | ● active |
| AI-PATROL-018 | static analysis | ● active |
| AI-PATROL-019 | static analysis | ● active |
| AI-GROUPCMD-020 | static analysis | ● active (amended, superseded) |
| AI-GUARD-021 | static analysis | ● active (amended) |
| AI-SWARM-022 | static analysis | ● active |
| AI-MOVE-023 | static analysis | ● active (superseded) |
| AI-SWARM2-024 | static analysis | ● active (amended, superseded) |
| AI-ROAM-025 | static analysis | ● active (amended) |
| AI-WITHDRAW-026 | static analysis | ● active |
| AI-WITHDRAW-027 | static analysis | ● active |
| AI-WITHDRAW-028 | static analysis | ● active (amended, partially retracted) |
| AI-CLASS-029 | static analysis | ● active |
| AI-CLASS-030 | static analysis | ● active |
| AI-ORDER-031 | static analysis | ● active (amended) |
| AI-CMD-032 | static analysis | ● active |
| AI-CMD-033 | static analysis | ● active (amended) |
| AI-PROGRESS-034 | static analysis | ● active (amended, superseded) |
| AI-DEAD-035 | static analysis | ● active |
| AI-DEAD-036 | static analysis | ● active |
| AI-FORM-037 | static analysis | ● active (amended, partially retracted) |
| AI-SPREAD-038 | static analysis | ● active |
| AI-ORDER-039 | static analysis | ● active (amended, partially retracted) |
| AI-PURSUE-040 | static analysis | ● active |
| AI-BREAK-041 | static analysis | ● active |
| AI-POST-042 | static analysis | ● active (amended, partially retracted, superseded) |
| AI-STATE-043 | static analysis | ● active (amended) |
| AI-THREAT-044 | static analysis | ● active (amended, superseded) |
| AI-ROUTE-045 | static analysis | ● active |
| AI-CENSUS-046 | mixed | ● active |
| AI-CENSUS-047 | static analysis | ● active |
| AI-CENSUS-048 | mixed | ● active |
| AI-ROOT-049 | mixed | ● active |
| AI-CLICK-050 | static analysis | ● active (partially retracted) |
| AI-CLICK-051 | static analysis | ● active (amended, superseded) |
| AI-CURSOR-052 | static analysis | ● active (amended, partially retracted) |
| AI-PANEL-053 | static analysis | ● active (partially retracted, superseded) |
| AI-CMD-054 | static analysis | ● active |
| AI-STRIKE-055 | static analysis | ● active |
| AI-RETAL-056 | static analysis | ● active |
| AI-ARBITER-057 | static analysis | ● active |
| AI-RAND-058 | static analysis | ● active (amended, superseded) |
| AI-KEYMOD-059 | static analysis | ● active |
| AI-PANEL-060 | static analysis | ● active |
| AI-PANEL-061 | static analysis | ● active (partially retracted) |
| AI-MINIMAP-062 | static analysis | ● active |
| AI-SURFACE-063 | static analysis | ● active (partially retracted, superseded) |
| AI-FANOUT-064 | static analysis | ● active |
| AI-SELECT-065 | static analysis | ● active |
| AI-FACE-066 | static analysis | ● active (amended) |
| AI-FACE-067 | static analysis | ● active (amended) |
| AI-GROUPSEE-068 | static analysis | ● active |
| AI-SCORE-069 | static analysis | ● active |
| AI-PREF-070 | static analysis | ● active |
| AI-COST-071 | static analysis | ● active |
| AI-REACH-072 | static analysis | ● active (amended) |
| AI-327 | static analysis | ● active |
| AI-328 | static analysis | ● active |
| AI-329 | static analysis | ● active |
| AI-FLIER-073 | static analysis | ● active |
| AI-GRPGUARD-074 | static analysis | ● active (amended, partially retracted) |
| AI-RADFREEZE-075 | static analysis | ● active (amended) |
| AI-STAND-076 | static analysis | ● active (amended) |
| AI-REISSUE-077 | static analysis | ● active |
| AI-DEADROLL-078 | static analysis | ● active |
| AI-GATE-079 | mixed | ● active (amended, partially retracted) |
| AI-CLOCK-080 | static analysis | ● active |
| AI-LOS-081 | static analysis | ● active |
| AI-DIPLO-082 | static analysis | ● active |
| AI-DIPLO-083 | static analysis | ● active (partially retracted) |
| AI-DIPLO-084 | static analysis | ● active |
| AI-DIPLO-085 | static analysis | ● active |
| AI-DIPLO-086 | static analysis | ● active |
| AI-LOS-087 | static analysis | ● active |
| AI-LOS-088 | static analysis | ● active |
| AI-LOS-089 | static analysis | ● active |
| AI-LOS-090 | static analysis | ● active |
| AI-LOS-091 | static analysis | ● active |
| AI-SIGHT-092 | static analysis | ● active |
| AI-SIGHT-093 | static analysis | ● active |
| AI-SIGHT-094 | static analysis | ● active |
| AI-POST-095 | static analysis | ● active |
| AI-POST-096 | static analysis | ● active |
| AI-POST-097 | static analysis | ● active |
| AI-STANCE-098 | static analysis | ● active |
| AI-LOAD-099 | static analysis | ● active |
| AI-GATE-100 | static analysis | ● active |
| AI-CENTRE-101 | static analysis | ● active |
| AI-RANGE-102 | static analysis | ● active |
| AI-JITTER-103 | static analysis | ● active |
| AI-TURN-104 | static analysis | ● active |
| AI-CANDCOUNT-105 | static analysis | ● active |
| AI-CANDLIST-106 | static analysis | ● active |
| AI-SWARM2GATE-107 | static analysis | ● active |
| AI-CMDSET45-108 | static analysis | ● active |
| AI-BB4SWEEP-109 | static analysis | ● active |
| AI-CANDBYTE-110 | static analysis | ● active |
| AI-DEFEND-111 | static analysis | ● active |
| AI-FOLLOW-112 | static analysis | ● active |
| AI-FOLLOWTAB-113 | static analysis | ● active |
| AI-FOLLOWGAP-114 | static analysis | ● active |
| AI-FOLLOWRANGE-115 | static analysis | ● active |
| AI-FOLLOWSET-116 | static analysis | ● active |
| AI-FOLLOWAUTH-117 | static analysis | ● active |
| AI-FOLLOWHEAL-118 | static analysis | ● active |
| AI-FOLLOWDEATH-119 | static analysis | ● active |
| AI-SCRIPTATTACK-120 | static analysis | ● active (amended, superseded) |
| AI-MINIMAP-156 | static analysis | ● active (amended, partially retracted) |
| AI-MINIMAP-157 | static analysis | ● active |
| AI-PANEL-158 | static analysis | ● active |
| AI-CURSOR-172 | static analysis | ● active (amended, partially retracted) |
| AI-CURSOR-173 | static analysis | ● active |
| AI-CURSOR-174 | static analysis | ● active |
| AI-CURSOR-175 | static analysis | ● active |
| AI-CURSOR-176 | static analysis | ● active (amended) |
| AI-CURSOR-177 | static analysis | ● active |
| AI-CURSOR-178 | static analysis | ● active |
| AI-CURSOR-188 | static analysis | ● active |
| AI-CURSOR-189 | static analysis | ● active |
| AI-CURSOR-190 | static analysis | ● active |
| AI-CURSOR-191 | static analysis | ● active |
| AI-CURSOR-192 | static analysis | ● active |
| AI-CURSOR-193 | static analysis | ● active |
| AI-CURSOR-194 | static analysis | ● active |
| AI-CURSOR-195 | static analysis | ● active |
| AI-CURSOR-196 | static analysis | ● active |
| AI-CURSOR-202 | static analysis | ● active |
| AI-CURSOR-203 | static analysis | ● active |
| AI-CURSOR-204 | static analysis | ● active |
| AI-CURSOR-205 | static analysis | ● active |
| AI-CURSOR-206 | static analysis | ● active |
| AI-CURSOR-207 | static analysis | ● active |
| AI-CURSOR-208 | static analysis | ● active |
| AI-CURSOR-209 | static analysis | ● active (amended) |
| AI-CURSOR-218 | static analysis | ● active |
| AI-CURSOR-219 | static analysis | ● active |
| AI-CURSOR-220 | static analysis | ● active |
| AI-CURSOR-221 | static analysis | ● active |
| AI-CURSOR-222 | static analysis | ● active |
| AI-CURSOR-223 | static analysis | ● active |
| AI-CURSOR-224 | static analysis | ● active |
| AI-CURSOR-225 | static analysis | ● active |
| AI-CURSOR-226 | static analysis | ● active |
| AI-CURSOR-227 | static analysis | ● active |
| AI-CURSOR-228 | static analysis | ● active |
| AI-CURSOR-229 | static analysis | ● active |
| AI-CURSOR-230 | static analysis | ● active |
| AI-CURSOR-231 | static analysis | ● active |
| AI-CURSOR-232 | static analysis | ● active |
| AI-CURSOR-233 | static analysis | ● active (amended, partially retracted) |
| AI-CURSOR-234 | static analysis | ● active (amended, partially retracted) |
| AI-CURSOR-235 | static analysis | ● active |
| AI-CURSOR-236 | static analysis | ● active |
| AI-CURSOR-242 | static analysis | ● active |
| AI-CURSOR-243 | static analysis | ● active |
| AI-ORDER-294 | static analysis | ● active |
| AI-332 | static analysis | ● active (amended) |
| AI-335 | static analysis | ● active (amended) |
| AI-336 | static analysis | ● active |
| AI-337 | static analysis | ● active |
| AI-340 | static analysis | ● active |
| AI-341 | static analysis | ● active |
| AI-360 | static analysis | ● active |
| AI-361 | static analysis | ● active |
| AI-362 | static analysis | ● active |
| AI-363 | static analysis | ● active |
| AI-364 | static analysis | ● active |
| AI-365 | static analysis | ● active |
| AI-366 | static analysis | ● active |
| AI-367 | static analysis | ● active |
| AI-370 | static analysis | ● active |
| AI-371 | static analysis | ● active |
| AI-372 | static analysis | ● active |
| AI-373 | static analysis | ● active |
| AI-374 | static analysis | ● active |
| AI-375 | static analysis | ● active (amended) |
| AI-394 | static analysis | ● active |
| AI-395 | static analysis | ● active |
| AI-376 | static analysis | ✔ promoted (amended) |
| AI-377 | static analysis | ✔ promoted |
| AI-378 | static analysis | ● active (amended) |
| AI-381 | static analysis | ● active |
| AI-382 | static analysis | ● active |
| AI-383 | static analysis | ● active |
| AI-384 | static analysis | ● active |
| AI-385 | static analysis | ● active |
| AI-387 | static analysis | ● active |
| AI-390 | static analysis | ● active |
| AI-391 | static analysis | ● active |
| AI-392 | static analysis | ● active |
| AI-393 | mixed | ● active |
| AI-397 | static analysis | ✔ promoted (branch candidate) |
| AI-398 | static analysis | ✔ promoted (branch candidate) |
| AI-405 | static analysis | ● active |
| AI-406 | static analysis | ● active |
| AI-407 | static analysis | ● active |
| AI-408 | static analysis | ● active |
| AI-409 | static analysis | ● active |
| AI-410 | static analysis | ● active |
| AI-411 | static analysis | ● active |
| AI-412 | static analysis | ● active (partially retracted) |
| AI-413 | static analysis | ● active |
| AI-414 | static analysis | ● active |
| AI-415 | static analysis | ● active |
| AI-416 | static analysis | ● active |
| AI-417 | static analysis | ● active |
| AI-418 | static analysis | ● active |

## Ledger alm

| ID | Method | Status |
|---|---|---|
| ALM-T9CONTROL-175 | static analysis | ✔ promoted |
| ALM-T9ACTOR-176 | static analysis | ✔ promoted |
| ALM-HDR-001 | static analysis | ● active (amended) |
| ALM-SEC-002 | static analysis | ● active (amended) |
| ALM-SEC-003 | static analysis |  |
| ALM-SEC-004 | static analysis | ● active (amended) |
| ALM-TRL-005 | mixed | ✖ retracted |
| ALM-HDR-006 | mixed | ✖ retracted |
| ALM-LOC-007 | mixed | ● active |
| ALM-META-008 | mixed | ● active (amended) |
| ALM-META-009 | static analysis |  |
| ALM-META-010 | mixed | ● active (amended) |
| ALM-GRID-011 | mixed |  |
| ALM-GRID-012 | static analysis | ● active (amended) |
| ALM-GRID-013 | static analysis | ● active |
| ALM-GRID-014 | static analysis | ● active (amended) |
| ALM-TERR-015 | static analysis | ● active (amended) |
| ALM-TERR-016 | static analysis | ● active (amended) |
| ALM-CNT-017 | static analysis | ● active (amended) |
| ALM-UNIT-018 | mixed | ● active (amended) |
| ALM-OBJ-019 | static analysis | ● active (amended, contested) |
| ALM-GRP-020 | mixed | ● active (amended) |
| ALM-TRIG-021 | mixed | ● active (amended) |
| ALM-TRIG-022 | static analysis | ● active (amended) |
| ALM-CODE-023 | static analysis | ● active |
| ALM-META-024 | static analysis | ● active (amended) |
| ALM-META-025 | static analysis | ● active (amended) |
| ALM-META-026 | static analysis | ● active (amended) |
| ALM-META-027 | mixed |  |
| ALM-META-028 | static analysis | ● active (amended) |
| ALM-HDR-029 | mixed | ✖ retracted (superseded by ALM-FRAME-031) |
| ALM-TRL-030 | static analysis | ✖ retracted |
| ALM-FRAME-031 | static analysis |  |
| ALM-GRID-032 | static analysis | ● active |
| ALM-PLACE-033 | static analysis | ● active |
| ALM-OBJ-034 | static analysis | ● active (amended) |
| ALM-CLS-035 | static analysis | ● active (amended) |
| ALM-CLS-036 | static analysis | ● active (amended) |
| ALM-CLS-037 | static analysis | ● active (amended) |
| ALM-CLS-038 | static analysis | ● active (amended) |
| ALM-OWN-039 | static analysis | ● active |
| ALM-UNIT-040 | static analysis | ● active (amended) |
| ALM-GRP-041 | static analysis |  |
| ALM-CLS-042 | mixed | ● active (amended) |
| ALM-TERR-043 | static analysis | ● active |
| ALM-TRIG-044 | static analysis | ● active |
| ALM-TRIG-045 | static analysis | ● active |
| ALM-TRIG-046 | static analysis | ● active |
| ALM-TRIG-047 | static analysis |  |
| ALM-UNIT-048 | mixed | ● active |
| ALM-TRIG-049 | mixed | ● active (amended) |
| ALM-TRIG-050 | static analysis | ● active (amended) |
| ALM-CLS-051 | static analysis | ● active |
| ALM-CLS-052 | static analysis | ● active |
| ALM-CLS-053 | static analysis | ● active |
| ALM-CLS-054 | static analysis | ● active (corrected) |
| ALM-REQ-055 | static analysis | ● active |
| ALM-REQ-056 | static analysis |  |
| ALM-ORD-057 | static analysis |  |
| ALM-META-058 | static analysis | ● active |
| ALM-RDR-059 | static analysis | ● active (amended) |
| ALM-CORP-060 | mixed | ● active |
| ALM-OBJ-061 | static analysis | ● active |
| ALM-OBJ-062 | static analysis | ● active |
| ALM-CLS-063 | static analysis | ● active |
| ALM-LIM-064 | static analysis | ● active |
| ALM-SACK-065 | static analysis | ● active |
| ALM-SACK-066 | static analysis | ● active |
| ALM-LIM-067 | static analysis | ● active (amended) |
| ALM-ORD-068 | static analysis | ● active |
| ALM-PLAYER-069 | static analysis | ● active (amended) |
| ALM-MODE-070 | static analysis | ● active |
| ALM-EFFREC-071 | mixed | ● active |
| ALM-EFFLINK-072 | mixed | ● active |
| ALM-EFFPOP-073 | mixed | ● active |
| ALM-M40-074 | mixed | ● active |
| ALM-TAILMAP-079 | static analysis | ● active |
| ALM-TAILHOLD-080 | static analysis | ● active |
| ALM-TAILDIR-081 | static analysis | ● active |
| ALM-TAILU16-082 | static analysis | ● active |
| ALM-TAILRUN-083 | static analysis | ● active |
| ALM-TAILVER-084 | static analysis | ● active |
| ALM-SCALAR-087 | static analysis | ● active (amended) |
| ALM-SCALAR-088 | static analysis | ● active (amended) |
| ALM-SCALAR-089 | mixed | ● active (amended) |
| ALM-CPLAYER-090 | static analysis | ● active (amended) |
| ALM-META-091 | static analysis | ● active |
| ALM-META-092 | static analysis | ● active |
| ALM-CORP-093 | static analysis | ● active |
| ALM-HEADER-097 | static analysis | ● active |
| ALM-HEADER-098 | mixed | ● active |
| ALM-WRITER-099 | static analysis | ● active |
| ALM-STREAM-100 | static analysis | ● active |
| ALM-CENSUS-101 | mixed | ● active |
| ALM-EDSCALAR-102 | static analysis | ● active |
| ALM-FLAGPATH-109 | static analysis | ✔ promoted |
| ALM-FLAGCORP-110 | mixed | ✔ promoted |
| ALM-PLACESTREAM-111 | static analysis | ✔ promoted |
| ALM-PLACEEDITOR-112 | static analysis | ✔ promoted |
| ALM-127 | static analysis | ✔ promoted |
| ALM-128 | static analysis | ✔ promoted |
| ALM-129 | static analysis | ✔ promoted |
| ALM-130 | static analysis | ✔ promoted |
| ALM-TILEMAIN-121 | static analysis | ✔ promoted |
| ALM-TILEVIEW-122 | static analysis | ✔ promoted |
| ALM-139 | static analysis | ✔ promoted |
| ALM-140 | static analysis | ✔ promoted |
| ALM-COPY-151 | static analysis | ✔ promoted |
| ALM-EDITOR-152 | static analysis | ✔ promoted |
| ALM-EDITOR-153 | static analysis | ✔ promoted |
| ALM-ALIAS-154 | static analysis | ✔ promoted |
| ALM-METACOPY-183 | static analysis | ✔ promoted |
| ALM-METALIFE-184 | static analysis | ✔ promoted |
| ALM-METAEDITOR-185 | static analysis | ✔ promoted |
| ALM-METADEFAULT-186 | static analysis | ✔ promoted (amended control population) |
| ALM-COUNT-195 | static analysis | ● active |
| ALM-STALE-196 | static analysis | ● active |
| ALM-HDRREAD-197 | static analysis | ● active |
| ALM-SCALAR-198 | static analysis | ● active |
| ALM-READERS-199 | mixed | ● active |
| ALM-FRONTIER-200 | static analysis | ✔ promoted |
| ALM-INGEST-201 | static analysis | ● active |
| ALM-EDITORHDR-202 | static analysis | ✔ promoted |
| ALM-EDITORSCALAR-203 | static analysis | ✔ promoted |
| ALM-EDITCOUNT-204 | static analysis | ✔ promoted |
| ALM-TRIGSTORAGE-163 | static analysis | ✔ promoted |
| ALM-TRIGEDITOR-164 | static analysis | ✔ promoted |
| ALM-TRIGZERO-165 | static analysis | ✔ promoted |
| ALM-TRIGMETA-166 | static analysis | ✔ promoted |
| ALM-RECVSOURCE-207 | static analysis | ✔ promoted |
| ALM-RECVREAD-208 | static analysis | ✔ promoted |
| ALM-RECVOBJ-209 | static analysis | ✔ promoted |
| ALM-SUBMIT-210 | static analysis | ✔ promoted |
| ALM-EDITTIME-205 | static analysis | ✔ promoted |
| ALM-HDRLIFE-206 | mixed | ✔ promoted |
| ALM-CATAPULT-213 | mixed | ● active |
| ALM-LOOT-214 | mixed | ● active |
| ALM-VIEW-215 | mixed | ● active |

## Ledger anim

| ID | Method | Status |
|---|---|---|
| ANIM-CLOCK-001 | static analysis | ● active |
| ANIM-STATE-002 | static analysis | ● active (amended) |
| ANIM-PHASE-003 | static analysis | ● active (amended) |
| ANIM-RUN-004 | static analysis | ● active |
| ANIM-MSG-005 | static analysis | ● active (partially retracted) |
| ANIM-DIR-006 | static analysis | ● active |
| ANIM-DEATH-007 | static analysis | ● active (amended, partially retracted) |
| ANIM-OBJ-008 | static analysis | ● active (amended, partially retracted) |
| ANIM-IDLE-009 | static analysis | ● active |
| ANIM-IDLE-010 | mixed | ● active |
| ANIM-TICK-011 | static analysis | ● active |
| ANIM-VT-012 | static analysis | ● active |
| ANIM-WALK-013 | static analysis | ● active |
| ANIM-WALK-014 | static analysis | ● active |
| ANIM-WALK-015 | static analysis | ● active |
| ANIM-AMBIENT-016 | static analysis | ● active |
| ANIM-PACE-017 | static analysis | ● active |
| ANIM-ARM-018 | static analysis | ● active |
| ANIM-BLOW-019 | static analysis | ● active (amended, partially retracted) |
| ANIM-NUM-020 | static analysis | ● active (amended) |
| ANIM-NUM-021 | static analysis | ● active |
| ANIM-SND-022 | static analysis | ● active (amended, partially retracted) |
| ANIM-STATE-023 | static analysis | ● active |
| ANIM-CLOCK-024 | static analysis | ● active |
| ANIM-094 | static analysis | ✔ promoted |
| ANIM-095 | static analysis | ✔ promoted (amended) |
| ANIM-096 | static analysis | ✔ promoted (amended) |
| ANIM-119 | static analysis | ● active |
| ANIM-120 | static analysis | ● active |
| ANIM-PROJ-025 | static analysis | ● active |
| ANIM-PROJ-026 | static analysis | ● active |
| ANIM-CAST-027 | static analysis | ● active |
| ANIM-PHASECLOCK-028 | static analysis | ● active (amended, partially retracted) |
| ANIM-BOLTDRAW-034 | static analysis | ✔ promoted |
| ANIM-BOLTRAMP-035 | static analysis | ✔ promoted |
| ANIM-AREAPHASE-029 | static analysis | ● active |
| ANIM-AREAOFFSET-030 | static analysis | ● active |
| ANIM-AREANOSTATE-031 | static analysis | ● active |
| ANIM-AREATRANS-032 | static analysis | ● active |
| ANIM-WALLFIREFRAME-033 | static analysis | ● active |
| ANIM-PARK-039 | static analysis | ● active |
| ANIM-044 | static analysis | ✔ promoted |
| ANIM-045 | mixed | ✔ promoted |
| ANIM-046 | mixed | ✔ promoted |
| ANIM-047 | mixed | ✔ promoted (partially retracted) |
| ANIM-071 | static analysis | ● active |
| ANIM-072 | static analysis | ● active |
| ANIM-073 | mixed | ● active |
| ANIM-074 | static analysis | ● active (amended) |
| ANIM-075 | static analysis | ● active (amended, partially retracted) |
| ANIM-REGISTER-083 | static analysis | ✔ promoted |
| ANIM-CATEGORY-084 | static analysis | ✔ promoted |
| ANIM-CELL-085 | static analysis | ✔ promoted |
| ANIM-AIRPASS-086 | static analysis | ✔ promoted |
| ANIM-DRAWGATE-087 | static analysis | ✔ promoted |
| ANIM-WALKORDER-088 | mixed | ✔ promoted |
| ANIM-097 | static analysis | ✔ promoted |
| ANIM-098 | static analysis | ✔ promoted |
| ANIM-099 | static analysis | ✔ promoted |
| ANIM-100 | static analysis | ✔ promoted |
| ANIM-101 | static analysis | ● active |
| ANIM-102 | static analysis | ● active |
| ANIM-103 | static analysis | ● active (amended) |
| ANIM-109 | static analysis | ● active |
| ANIM-110 | static analysis | ● active (amended) |
| ANIM-111 | static analysis | ● active |
| ANIM-112 | static analysis | ● active |
| ANIM-113 | static analysis | ● active |
| ANIM-114 | static analysis | ● active (amended) |
| ANIM-115 | static analysis | ● active (amended) |
| ANIM-121 | static analysis | ● active |
| ANIM-122 | static analysis | ● active |
| ANIM-123 | static analysis | ● active |
| ANIM-124 | static analysis | ● active |
| ANIM-125 | static analysis | ● active |
| ANIM-126 | static analysis | ● active (amended) |
| ANIM-127 | static analysis | ● active (amended) |
| ANIM-128 | static analysis | ● active |
| ANIM-105 | static analysis | ● active |
| ANIM-106 | static analysis | ● active |
| ANIM-107 | static analysis | ● active |
| ANIM-108 | static analysis | ● active |
| ANIM-116 | static analysis | ● active |
| ANIM-117 | static analysis | ● active (amended, partially retracted) |
| ANIM-118 | static analysis | ● active |
| ANIM-129 | static analysis | ✔ promoted |
| ANIM-130 | static analysis | ✔ promoted |
| ANIM-131 | static analysis | ✔ promoted |
| ANIM-132 | static analysis | ✔ promoted |
| ANIM-133 | static analysis | ✔ promoted |
| ANIM-134 | static analysis | ✔ promoted |
| ANIM-135 | static analysis | ✔ promoted |
| ANIM-136 | static analysis | ✔ promoted |

## Ledger databin

| ID | Method | Status |
|---|---|---|
| DAT-LOC-001 | static analysis | ● active |
| DAT-OBJ-002 | static analysis | ● active (amended) |
| DAT-GRAM-003 | static analysis | ● active |
| DAT-SCHEMA-004 | static analysis | ● active |
| DAT-SCHEMA-007 | static analysis | ● active |
| DAT-HUMANS-008 | static analysis | ● active (amended, partially retracted) |
| DAT-HUMANS-009 | static analysis | ● active |
| DAT-ACT-006 | static analysis | ● active (amended, partially retracted) |
| DAT-BLD-005 | static analysis | ● active |
| DAT-MATMASK-020 | static analysis | ● active |
| DAT-ITEMNAME-010 | mixed | ● active |
| DAT-NAMES-016 | static analysis | ● active |
| DAT-DOC-021 | mixed | ✔ promoted |

## Ledger dialogue

| ID | Method | Status |
|---|---|---|
| DLG-WIN-001 | static analysis | ● active (amended) |
| DLG-PATH-002 | static analysis | ● active (partially retracted, amended) |
| DLG-ABSENT-003 | static analysis | ● active |
| DLG-EMPTY-004 | static analysis | ● active (amended) |
| DLG-LIFE-005 | static analysis | ● active (amended, partially retracted) |
| DLG-READ-006 | static analysis | ● active |
| DLG-MARKUP-007 | static analysis | ● active (amended) |
| DLG-FACE-008 | static analysis | ● active |
| DLG-WRAP-009 | static analysis | ● active (amended, partially retracted) |
| DLG-LANG-010 | mixed | ● active |
| DLG-READER-011 | mixed | ● active (amended) |
| DLG-STOP-012 | static analysis | ● active |
| DLG-DIM-013 | static analysis | ● active (amended, partially retracted) |
| DLG-CLOCK-014 | static analysis | ● active |
| DLG-DRAW-015 | static analysis | ● active (amended) |
| DLG-ENTRY-016 | static analysis | ● active |
| DLG-MODAL-017 | static analysis | ● active |
| DLG-NPCTAG-018 | static analysis | ● active |
| DLG-NPCTAG-019 | static analysis | ● active (amended, partially retracted) |
| DLG-FIGURE-020 | static analysis | ● active (amended) |
| DLG-FIGURE-021 | static analysis | ● active |
| DLG-SPEAKER-022 | static analysis | ● active |
| DLG-SPEAKER-023 | static analysis | ● active |
| DLG-DRESS-024 | mixed | ● active (amended, partially retracted) |
| DLG-SPEAKER-041 | static analysis | ● active |
| DLG-SYNTH-042 | static analysis | ● active |
| DLG-FACEBYTE-043 | static analysis | ● active |
| DLG-MSGNUM-025 | static analysis | ● active |
| DLG-MISSION-026 | static analysis | ● active |
| DLG-TAGARM-027 | static analysis | ● active (amended) |
| DLG-SOUND-028 | static analysis | ● active |
| DLG-ZEROARM-029 | static analysis | ● active |
| DLG-INNVOICE-030 | static analysis | ● active (amended) |
| DLG-PANEL-035 | static analysis | ✔ promoted (amended) |
| DLG-PORTRAIT-036 | static analysis | ✔ promoted (amended) |
| DLG-RECT-037 | static analysis | ● active |
| DLG-LINE-038 | static analysis | ✔ promoted (amended) |
| DLG-BUTTON-039 | static analysis | ✔ promoted (amended) |
| DLG-KEYS-040 | static analysis | ● active (amended) |
| DIALOGUE-054 | static analysis | ✔ promoted |
| DIALOGUE-055 | static analysis | ✔ promoted |
| DIALOGUE-056 | static analysis | ✔ promoted |
| DIALOGUE-057 | static analysis | ✔ promoted |
| DIALOGUE-044 | static analysis | ✔ promoted |
| DIALOGUE-045 | static analysis | ✔ promoted |
| DIALOGUE-048 | static analysis | ● active |
| DIALOGUE-049 | static analysis | ● active |
| DIALOGUE-050 | static analysis | ● active |
| DIALOGUE-051 | mixed | ● active |
| DIALOGUE-060 | static analysis | ✔ promoted |
| DIALOGUE-061 | static analysis | ✔ promoted |
| DIALOGUE-062 | mixed | ✔ promoted (amended) |
| DIALOGUE-063 | static analysis | ✔ promoted |
| DIALOGUE-064 | static analysis | ✔ promoted |
| DIALOGUE-065 | static analysis | ✔ promoted |
| DIALOGUE-068 | static analysis | ✔ promoted |
| DIALOGUE-069 | static analysis | ✔ promoted |
| DIALOGUE-070 | static analysis | ✔ promoted |
| DLG-REPORT-072 | observation | ● active |
| DLG-VOICE-074 | mixed | ● active |
| DLG-SPEAKER-075 | mixed | ● active |
| DLG-VOICE-076 | mixed | ● active |

## Ledger fame

| ID | Method | Status |
|---|---|---|
| FAME-HDR-001 | static analysis | ✔ promoted |
| FAME-REC-002 | mixed | ● active (amended) |
| FAME-NAME-003 | mixed | ● active (partially retracted) |
| FAME-SCORE-004 | mixed | ● active (partially retracted) |
| FAME-UNK-005 | mixed | ● active (amended) |
| FAME-DEFAULT-006 | mixed | ● active |
| FAME-WRITE-007 | static analysis | ● active |
| FAME-DEFAULT-008 | static analysis | ● active (amended) |
| FAME-READER-009 | static analysis | ✔ promoted |
| FAME-STRING-010 | static analysis | ✔ promoted |
| FAME-INSERT-011 | static analysis | ✔ promoted |
| FAME-DISPLAY-012 | static analysis | ✔ promoted |
| FAME-TAILS-013 | static analysis | ✔ promoted |
| FAME-PRODUCER-014 | static analysis | ✔ promoted |
| FAME-SEED-015 | static analysis | ✔ promoted |
| FAME-BOUNDARY-016 | static analysis | ✔ promoted |
| FAME-021 | static analysis | ✔ promoted |
| FAME-022 | static analysis | ✔ promoted |
| FAME-023 | static analysis | ✔ promoted |
| FAME-024 | static analysis | ✔ promoted |
| FAME-025 | static analysis | ✔ promoted |
| FAME-026 | static analysis | ● active (amended) |
| FAME-027 | static analysis | ✔ promoted |
| FAME-028 | static analysis | ✔ promoted |
| FAME-029 | static analysis | ● active (partially retracted) |
| FAME-030 | static analysis | ● active |
| FAME-031 | static analysis | ● active |
| FAME-032 | static analysis | ● active |
| FAME-033 | static analysis | ● active |
| FAME-034 | static analysis | ● active |

## Ledger hero

| ID | Method | Status |
|---|---|---|
| HERO-DEFEAT-136 | static analysis | ● active |
| HERO-STAT-001 | static analysis | ● active |
| HERO-COST-002 | static analysis | ● active |
| HERO-BUY-003 | static analysis | ● active |
| HERO-BUDGET-004 | static analysis | ● active |
| HERO-HP-005 | static analysis |  |
| HERO-MP-006 | static analysis | ● active |
| HERO-SIGHT-007 | static analysis | ● active (amended, contested) |
| HERO-SPEED-008 | static analysis | ● active |
| HERO-SKILL-009 | static analysis | ● active |
| HERO-XP-010 | static analysis | ✖ retracted |
| HERO-COMBAT-011 | static analysis |  |
| HERO-RESIST-012 | static analysis | ● active |
| HERO-CLASS-013 | static analysis | ● active (corrected) |
| HERO-ORDER-014 | static analysis | ● active |
| HERO-CAP-015 | static analysis | ● active |
| HERO-MOD-016 | static analysis | ● active |
| HERO-EQUIP-017 | static analysis | ● active (weapon inverse, direct-write scope and selector provenance corrected) |
| HERO-ARMOUR-018 | static analysis |  |
| HERO-EFFECT-019 | static analysis | ● active |
| HERO-CLASS-020 | static analysis | ● active |
| HERO-REGEN-021 | static analysis | ● active (partially retracted/amended) |
| HERO-DAMAGE-022 | static analysis | ● active (amended) |
| HERO-CADENCE-023 | static analysis | ● active (amended; complete-period formula and worked intervals retracted) |
| HERO-TARGET-024 | static analysis | ● active (amended, contested) |
| HERO-REACH-025 | static analysis | ● active |
| HERO-DEATH-026 | static analysis | ● active (amended) |
| HERO-KILL-027 | static analysis | ● active (amended) |
| HERO-AGGRO-028 | static analysis | ● active (amended) |
| HERO-DMG2-029 | static analysis | ● active |
| HERO-CLAMP-030 | static analysis | ● active |
| HERO-AUTOHIT-031 | static analysis | ● active |
| HERO-HEALTH-032 | static analysis | ● active (amended) |
| HERO-FOLD-033 | static analysis | ● active (amended) |
| HERO-DERIVE-034 | static analysis | ● active |
| HERO-FOLD-035 | static analysis | ● active (instruction counts corrected) |
| HERO-STATDMG-036 | static analysis | ● active |
| HERO-BARE-037 | static analysis | ● active |
| HERO-SHEET-038 | static analysis | ● active |
| HERO-START-039 | static analysis | ● active |
| HERO-APPEAR-040 | static analysis | ● active |
| HERO-APPEAR-041 | static analysis | ● active |
| HERO-APPEAR-042 | static analysis | ● active |
| HERO-APPEAR-043 | static analysis | ● active |
| HERO-APPEAR-044 | static analysis | ● active (amended) |
| HERO-APPEAR-045 | static analysis | ● active |
| HERO-APPEAR-046 | static analysis | ● active |
| HERO-APPEAR-047 | static analysis | ● active |
| HERO-APPEAR-048 | static analysis | ● active |
| HERO-APPEAR-049 | static analysis | ● active |
| HERO-APPEAR-050 | static analysis | ● active |
| HERO-APPEAR-051 | static analysis | ● active |
| HERO-APPEAR-052 | static analysis | ● active (amended) |
| HERO-APPEAR-053 | static analysis | ● active |
| HERO-APPEAR-054 | static analysis | ● active |
| HERO-APPEAR-055 | static analysis | ● active |
| HERO-APPEAR-056 | static analysis |  |
| HERO-FIGURE-057 | static analysis | ● active |
| HERO-FIGURE-058 | static analysis | ● active |
| HERO-FIGURE-059 | static analysis | ● active |
| HERO-FIGURE-060 | static analysis | ● active |
| HERO-FIGURE-061 | static analysis | ● active (amended) |
| HERO-FIGURE-062 | static analysis | ● active (corrects `HERO-APPEAR-056`) |
| HERO-FIGURE-063 | static analysis | ● active (amended) |
| HERO-FIGURE-064 | static analysis | ● active |
| HERO-DWELL-065 | static analysis | ● active (amended) |
| HERO-FINISH-066 | static analysis | ● active |
| HERO-DYETICK-067 | static analysis | ● active |
| HERO-REVIVE-068 | static analysis | ● active |
| HERO-DECAY-069 | static analysis | ● active |
| HERO-ZERO-070 | static analysis | ● active |
| HERO-HP-071 | static analysis | ● active |
| HERO-HP-072 | static analysis | ● active |
| HERO-SKILLUP-073 | static analysis | ● active |
| HERO-SKILLGATE-074 | mixed | ● active (corrected) |
| HERO-SKILLLOSS-075 | static analysis |  |
| HERO-SKILLBUY-076 | static analysis | ● active (amended) |
| HERO-XP-077 | static analysis | ● active |
| HERO-DOLL-078 | static analysis | ● active |
| HERO-START-081 | static analysis | ✔ promoted (amended) |
| HERO-NAME-079 | static analysis | ● active |
| HERO-TYPED-080 | static analysis | ● active |
| HERO-CHARGEN-082 | static analysis | ● active |
| HERO-CHARGEN-083 | static analysis | ● active (amended) |
| HERO-CHARGEN-084 | static analysis | ● active |
| HERO-CHARGEN-085 | static analysis | ● active |
| HERO-GENERAL-086 | static analysis | ● active (corrected after two independent falsifications) |
| HERO-GENERAL-087 | mixed | ● active |
| HERO-GENERAL-088 | static analysis | ● active (corrected after two independent falsifications) |
| HERO-GENERAL-089 | mixed | ● active (expanded after independent falsification) |
| HERO-GENERAL-090 | static analysis | ● active |
| HERO-GENERAL-091 | static analysis | ● active |
| HERO-GENERAL-092 | static analysis | ● active (corrected after independent falsification) |
| HERO-GENERAL-093 | static analysis | ● active (corrected after four independent falsifications) |
| HERO-ITEMSKILL-096 | static analysis | ● active |
| HERO-104 | static analysis | ● active |
| HERO-CADENCE-112 | static analysis | ● active |
| HERO-CADENCE-113 | static analysis | ● active |
| HERO-CADENCE-114 | static analysis | ● active |
| HERO-CADENCE-115 | static analysis | ● active (start/reach clause corrected) |
| HERO-JOIN-120 | static analysis | ● active |
| HERO-JOIN-121 | mixed | ● active |
| HERO-JOIN-122 | mixed | ● active |
| HERO-JOIN-123 | mixed | ● active |
| HERO-JOIN-124 | mixed | ● active |
| HERO-JOIN-125 | mixed | ● active |
| HERO-JOIN-126 | static analysis | ● active |
| HERO-JOIN-127 | mixed | ● active |
| HERO-JOIN-128 | mixed | ● active |
| HERO-FIGURE-144 | static analysis | ● active |
| HERO-DYINGTICK-145 | static analysis | ● active |
| HERO-CROSSHOLD-146 | static analysis | ● active |

## Ledger inv

| ID | Method | Status |
|---|---|---|
| INV-CORPUS-001 | mixed | ● active |
| INV-SIG-002 | mixed | ● active |
| INV-SIG-003 | static analysis | ● active |
| INV-SIG-004 | mixed | ● active |
| INV-SCOPE-005 | mixed | ● active |

## Ledger item

| ID | Method | Status |
|---|---|---|
| ITEM-USE-112 | static analysis | ✔ promoted |
| ITEM-USE-113 | static analysis | ✔ promoted |
| ITEM-USE-114 | static analysis | ✔ promoted |
| ITEM-VALUE-115 | static analysis | ✔ promoted |
| ITEM-CLASS-001 | static analysis | ● active |
| ITEM-DEF-002 | static analysis | ● active (partially retracted) |
| ITEM-STACK-003 | static analysis | ● active (amended, partially retracted) |
| ITEM-CONT-004 | static analysis | ● active |
| ITEM-LOAD-005 | static analysis | ● active |
| ITEM-EQUIP-006 | static analysis | ● active (amended, partially retracted) |
| ITEM-CMD-007 | static analysis | ● active (amended, partially retracted) |
| ITEM-DROP-008 | static analysis | ● active |
| ITEM-PICK-009 | static analysis | ● active (amended, superseded) |
| ITEM-SACK-010 | static analysis | ● active (contested) |
| ITEM-SACK-011 | static analysis | ● active (contested) |
| ITEM-DEATH-012 | static analysis | ● active (amended, partially retracted) |
| ITEM-SPAWN-013 | static analysis | ● active (partially retracted) |
| ITEM-SAVE-014 | static analysis | ● active (amended) |
| ITEM-CARRY-015 | mixed | ● active |
| ITEM-PICK-016 | static analysis | ● active |
| ITEM-SCALE-017 | static analysis | ● active (partially retracted) |
| ITEM-DMGCOL-018 | static analysis | ● active (partially retracted) |
| ITEM-LADDER-019 | static analysis | ● active |
| ITEM-DMGFACT-020 | static analysis | ● active |
| ITEM-WEAPCOL-021 | static analysis | ✔ promoted (amended) |
| ITEM-PANEL-022 | static analysis | ● active |
| ITEM-APPEAR-023 | static analysis | ● active |
| ITEM-APPEAR-024 | static analysis | ● active (amended) |
| ITEM-APPEAR-025 | static analysis | ● active |
| ITEM-SPAWN-026 | static analysis | ● active |
| ITEM-SPAWN-027 | static analysis | ● active |
| ITEM-OWNED-028 | static analysis | ● active |
| ITEM-CODE-029 | static analysis | ● active |
| ITEM-HUMEQ-030 | static analysis | ● active |
| ITEM-ARMSLOT-031 | static analysis | ● active |
| ITEM-ARMFILL-032 | static analysis | ● active |
| ITEM-ARMFOLD-033 | static analysis | ● active (amended, partially retracted) |
| ITEM-CORPSE-034 | static analysis | ● active |
| ITEM-SUIT-035 | static analysis | ● active (partially retracted, superseded) |
| ITEM-PICT-046 | static analysis | ● active |
| ITEM-PICT-047 | static analysis | ● active |
| ITEM-PICT-048 | static analysis | ● active |
| ITEM-PICT-049 | static analysis | ● active (amended) |
| ITEM-PICT-050 | static analysis | ● active |
| ITEM-PICT-051 | static analysis | ● active |
| ITEM-PICT-052 | static analysis | ● active |
| ITEM-DISPNAME-036 | static analysis | ● active |
| ITEM-NAMEKEY-037 | static analysis | ● active |
| ITEM-NAMEPOP-038 | mixed | ● active (amended) |
| ITEM-NAMEMISS-039 | static analysis | ● active (amended) |
| ITEM-NAMEPARSE-040 | static analysis | ● active |
| ITEM-BRACE-041 | mixed | ● active (amended) |
| ITEM-NAMELIMIT-042 | static analysis | ● active |
| ITEM-DOC-053 | static analysis | ● active |
| ITEM-DOC-054 | static analysis | ✔ promoted (amended) |
| ITEM-DOC-069 | static analysis | ✔ promoted |
| ITEM-WEAR-055 | static analysis | ● active (amended, partially retracted) |
| ITEM-WEAR-056 | static analysis | ● active (amended, partially retracted) |
| ITEM-WEAR-057 | static analysis | ● active |
| ITEM-WEAR-058 | static analysis | ● active |
| ITEM-CASTSTATE-056 | static analysis | ● active (amended, partially retracted) |
| ITEM-EFFGRAM-070 | static analysis | ● active |
| ITEM-EFFPOP-071 | mixed | ● active |
| ITEM-EFFOBJ-072 | static analysis | ● active |
| ITEM-EFFMODE-073 | mixed | ● active |
| ITEM-EFFSPLIT-074 | static analysis | ● active |
| ITEM-EFFDISP-075 | static analysis | ● active |
| ITEM-CASTLINK-076 | static analysis | ● active |
| ITEM-EFFSAVE-077 | static analysis | ● active |
| ITEM-EFFARM-146 | static analysis | ● active |
| ITEM-EFFKEY-147 | static analysis | ● active |
| ITEM-EFFSYM-148 | static analysis | ● active |
| ITEM-EFFSHIP-149 | mixed | ● active |
| ITEM-AUTHCAST-086 | mixed | ● active |
| ITEM-AUTHDROP-087 | mixed | ● active |
| ITEM-DRAGON-088 | mixed | ● active |
| ITEM-DRAGDROP-089 | mixed | ● active |
| ITEM-MAGVAL-090 | static analysis | ● active (amended, partially retracted) |
| ITEM-PRODUCER-091 | static analysis | ● active |
| ITEM-STARFLAG-096 | static analysis | ● active |
| ITEM-STARSURF-097 | static analysis | ● active |
| ITEM-STARPIX-098 | mixed | ● active |
| ITEM-STARPHASE-099 | static analysis | ● active |
| ITEM-STARCOMP-100 | static analysis | ● active |
| ITEM-WHOLE-128 | static analysis | ● active |
| ITEM-MERGE-129 | static analysis | ● active |
| ITEM-GROUNDMOVE-130 | static analysis | ● active |
| ITEM-EQUIPMOVE-131 | mixed | ● active |
| ITEM-SPELLMOVE-132 | static analysis | ● active |
| ITEM-136 | static analysis | ● active |
| ITEM-137 | static analysis | ● active |
| ITEM-138 | static analysis | ● active |
| ITEM-PRICETAG-144 | static analysis | ● active |
| ITEM-PICKTEXT-145 | static analysis | ● active |
| ITEM-152 | static analysis | ● active |
| ITEM-154 | static analysis | ● active |
| ITEM-155 | static analysis | ● active |
| ITEM-156 | static analysis | ● active |
| ITEM-157 | static analysis | ✔ promoted |
| ITEM-158 | static analysis | ✔ promoted |
| ITEM-159 | static analysis | ✔ promoted |
| ITEM-161 | static analysis | ✔ promoted |
| ITEM-162 | static analysis | ✔ promoted |

## Ledger magic

| ID | Method | Status |
|---|---|---|
| MAGIC-REACH-178 | static analysis | ● active |
| MAGIC-REACH-179 | static analysis | ● active |
| MAGIC-REACH-180 | static analysis | ● active |
| MAGIC-REACH-181 | static analysis | ● active |
| MAGIC-CONSUME-142 | static analysis | ✔ promoted |
| MAGIC-CONSUME-143 | static analysis | ✔ promoted |
| MAGIC-CONSUME-144 | static analysis | ✔ promoted |
| MAGIC-SPRAY-134 | static analysis | ● active |
| MAGIC-SPRAY-135 | static analysis | ● active (amended) |
| MAGIC-SPRAY-136 | static analysis | ● active |
| MAGIC-SPRAY-137 | mixed | ● active |
| MAGIC-SPELL-001 | static analysis | ● active |
| MAGIC-BOOK-002 | static analysis | ● active |
| MAGIC-CAST-003 | static analysis | ● active (partially retracted) |
| MAGIC-POWER-004 | static analysis | ● active (partially retracted) |
| MAGIC-DMG-005 | static analysis | ● active |
| MAGIC-RESIST-006 | static analysis | ● active (amended) |
| MAGIC-ITEM-007 | static analysis | ● active (amended, partially retracted) |
| MAGIC-SHAPE-008 | static analysis | ● active (partially retracted) |
| MAGIC-EFFMODE-009 | static analysis | ● active |
| MAGIC-MIND-010 | static analysis | ● active |
| MAGIC-SPIRIT-011 | static analysis | ● active |
| MAGIC-AI-012 | static analysis | ● active |
| MAGIC-221 | static analysis | ● active |
| MAGIC-CEIL-013 | static analysis | ● active (amended, superseded) |
| MAGIC-ARM-014 | static analysis | ● active |
| MAGIC-EFFECT-015 | static analysis | ● active (amended) |
| MAGIC-ATTACH-016 | static analysis | ● active (amended, partially retracted, superseded) |
| MAGIC-TARGET-017 | static analysis | ● active (partially retracted) |
| MAGIC-TRAIN-018 | static analysis | ● active (partially retracted) |
| MAGIC-SING-019 | static analysis | ● active (amended, partially retracted) |
| MAGIC-AUTOCAST-020 | static analysis | ● active |
| MAGIC-AUTOCAST-021 | static analysis | ● active |
| MAGIC-STAFF-022 | static analysis | ● active |
| MAGIC-SPELLHOP-023 | static analysis | ● active |
| MAGIC-ICON-024 | static analysis | ● active |
| MAGIC-ICON-025 | static analysis | ● active |
| MAGIC-PIC-026 | static analysis | ● active |
| MAGIC-PIC-027 | static analysis | ● active (amended, superseded) |
| MAGIC-CAST-028 | static analysis | ● active (partially retracted) |
| MAGIC-CASTANIM-029 | static analysis | ● active |
| MAGIC-CASTTICK-030 | static analysis | ● active (partially retracted) |
| MAGIC-BURST-031 | static analysis | ● active |
| MAGIC-SENDER-032 | static analysis | ● active |
| MAGIC-CASTSPAWN-033 | static analysis | ● active |
| MAGIC-BURSTLIFE-034 | static analysis | ● active (amended, partially retracted) |
| MAGIC-DELIVER-035 | static analysis | ● active |
| MAGIC-AREATICK-036 | static analysis | ● active |
| MAGIC-AREAPULSE-037 | static analysis | ● active (amended, partially retracted) |
| MAGIC-AREAAPPLY-038 | static analysis | ● active (amended, partially retracted) |
| MAGIC-AREACELL-039 | static analysis | ● active (amended, partially retracted) |
| MAGIC-MAPLAYER-040 | static analysis | ● active |
| MAGIC-AREAEND-041 | static analysis | ● active (amended, partially retracted) |
| MAGIC-WALLEARTH-042 | static analysis | ● active (amended) |
| MAGIC-AREATARGET-043 | static analysis | ● active (amended) |
| MAGIC-LIGHTDARK-044 | static analysis | ● active |
| MAGIC-WALLBLOCK-045 | static analysis | ● active |
| MAGIC-AREACOST-046 | static analysis | ● active |
| MAGIC-FIREDIV-047 | static analysis | ● active (amended, partially retracted) |
| MAGIC-RING-048 | static analysis | ✔ promoted |
| MAGIC-BOLTGATE-069 | static analysis | ✔ promoted |
| MAGIC-BOLTSHAPE-070 | static analysis | ✔ promoted |
| MAGIC-BOLTLIST-071 | static analysis | ✔ promoted (superseded) |
| MAGIC-BOLTSTILL-072 | static analysis | ✔ promoted |
| MAGIC-TRAIL-073 | static analysis | ✔ promoted |
| MAGIC-BOLTEND-074 | static analysis | ✔ promoted |
| MAGIC-AREADRAW-049 | static analysis | ● active |
| MAGIC-OVERLAY-050 | static analysis | ● active |
| MAGIC-OVERLAYART-051 | static analysis | ● active |
| MAGIC-CLOUDLIFE-052 | static analysis | ● active |
| MAGIC-CELLSET-053 | static analysis | ● active |
| MAGIC-AREARADIUS-054 | static analysis | ● active |
| MAGIC-LIGHTDRAW-055 | static analysis | ● active |
| MAGIC-LIGHTLEVEL-056 | static analysis | ● active |
| MAGIC-UNITLIGHT-057 | static analysis | ● active |
| MAGIC-WALLFIRE-058 | static analysis | ● active |
| MAGIC-MARK-059 | static analysis | ✔ promoted |
| MAGIC-MARK-060 | static analysis | ✔ promoted |
| MAGIC-MARK-061 | static analysis | ✔ promoted |
| MAGIC-PROT-062 | static analysis | ✔ promoted |
| MAGIC-SHIELD-063 | static analysis | ● active (superseded) |
| MAGIC-BLESS-064 | static analysis | ✔ promoted |
| MAGIC-CLOUD-065 | static analysis | ● active (partially retracted) |
| MAGIC-ACTOR-066 | static analysis | ✔ promoted (amended, superseded) |
| MAGIC-ACTGATE-079 | static analysis | ● active |
| MAGIC-ACTKEY-080 | static analysis | ● active |
| MAGIC-MASKREAD-081 | static analysis | ● active |
| MAGIC-AIBIT-082 | static analysis | ● active |
| MAGIC-EFFLOOK-083 | static analysis | ● active |
| MAGIC-STONEDRAW-084 | static analysis | ● active |
| MAGIC-INVISOWN-085 | static analysis | ● active |
| MAGIC-089 | static analysis | ✔ promoted |
| MAGIC-090 | mixed | ✔ promoted |
| MAGIC-091 | mixed | ✔ promoted |
| MAGIC-092 | mixed | ● active (amended, partially retracted) |
| MAGIC-093 | mixed | ✔ promoted |
| MAGIC-ITEMTRAIN-116 | mixed | ● active |
| MAGIC-ITEMKILL-117 | static analysis | ● active (amended, partially retracted) |
| MAGIC-ATTRGATE-118 | mixed | ● active |
| MAGIC-CADENCE-126 | static analysis | ● active |
| MAGIC-CADENCE-127 | static analysis | ● active |
| MAGIC-CLOUDCLOCK-154 | static analysis | ✔ promoted |
| MAGIC-CLOUDOWNER-155 | static analysis | ✔ promoted |
| MAGIC-CLOUDVISIT-156 | static analysis | ✔ promoted |
| MAGIC-POISONINPUT-157 | static analysis | ✔ promoted |
| MAGIC-POISONREFRESH-158 | static analysis | ✔ promoted |
| MAGIC-POISONPHASE-159 | static analysis | ✔ promoted |
| MAGIC-FIREPOISON-160 | static analysis | ✔ promoted |
| MAGIC-AREASOURCE-161 | static analysis | ✔ promoted |
| MAGIC-AREAOVERLAP-162 | mixed | ✔ promoted |
| MAGIC-CLOUDEND-163 | static analysis | ✔ promoted |
| MAGIC-DELIVERY-170 | static analysis | ● active |
| MAGIC-CASTCLOCK-171 | static analysis | ● active |
| MAGIC-TELEPORT-174 | static analysis | ● active |
| MAGIC-TARGETID-182 | static analysis | ● active |
| MAGIC-TICKGATE-183 | static analysis | ● active (partially retracted) |
| MAGIC-SIBLINGREF-184 | static analysis | ● active |
| MAGIC-187 | static analysis | ● active |
| MAGIC-188 | static analysis | ● active |
| MAGIC-189 | static analysis | ● active |
| MAGIC-190 | static analysis | ● active |
| MAGIC-197 | static analysis | ● active |
| MAGIC-202 | static analysis | ● active |
| MAGIC-209 | static analysis | ● active |
| MAGIC-213 | static analysis | ● active |
| MAGIC-214 | static analysis | ● active |
| MAGIC-215 | static analysis | ● active |
| MAGIC-216 | static analysis | ● active |
| MAGIC-217 | static analysis | ● active |
| MAGIC-219 | static analysis | ● active |
| MAGIC-236 | static analysis | ● active |
| MAGIC-225 | static analysis | ✔ promoted |
| MAGIC-235 | static analysis | ● active (amended) |
| MAGIC-237 | static analysis | ✔ promoted |
| MAGIC-238 | static analysis | ✔ promoted |
| MAGIC-239 | static analysis | ✔ promoted |
| MAGIC-240 | static analysis | ✔ promoted |
| MAGIC-245 | static analysis | ● active |
| MAGIC-246 | static analysis | ● active |
| MAGIC-247 | mixed | ● active |
| MAGIC-241 | static analysis | ● active |
| MAGIC-242 | static analysis | ● active |
| MAGIC-243 | static analysis | ● active |
| MAGIC-249 | static analysis | ● active |
| MAGIC-250 | static analysis | ● active |
| MAGIC-251 | static analysis | ● active |
| MAGIC-252 | static analysis | ● active |

## Ledger menu

| ID | Method | Status |
|---|---|---|
| MENU-COMBAT-017 | static analysis | ● active (partially retracted) |
| MENU-COMBAT-018 | mixed | ● active |
| MENU-COMBAT-019 | mixed | ● active |
| MENU-ASSET-001 | static analysis | ✔ promoted |
| MENU-ASSET-002 | static analysis | ✔ promoted |
| MENU-MASK-003 | static analysis | ✔ promoted |
| MENU-MASK-004 | static analysis | ✔ promoted |
| MENU-GEOM-005 | static analysis | ✔ promoted |
| MENU-GEOM-006 | mixed | ✔ promoted |
| MENU-STATE-007 | static analysis | ● active |
| MENU-STRTAB-008 | static analysis | ● active |
| MENU-DOC-009 | static analysis | ● active |
| MENU-ESC-010 | static analysis | ✔ promoted (partially retracted) |
| MENU-ITEM-011 | static analysis | ✔ promoted |
| MENU-ITEM-012 | static analysis | ✔ promoted |
| MENU-KEY-013 | static analysis | ✔ promoted |
| MENU-ART-014 | static analysis | ✔ promoted |
| MENU-STOP-015 | static analysis | ● active |
| MENU-INPUT-016 | static analysis | ● active (partially retracted) |
| MENU-CURSOR-046 | static analysis | ● active |
| MENU-051 | static analysis | ● active (amended) |
| MENU-052 | static analysis | ● active (amended) |
| MENU-053 | static analysis | ● active (amended) |
| MENU-054 | static analysis | ● active (amended) |
| MENU-055 | static analysis | ● active (amended) |
| MENU-056 | static analysis | ● active (amended) |
| MENU-057 | static analysis | ● active |
| MENU-058 | static analysis | ● active |
| MENU-059 | static analysis | ● active |
| MENU-060 | static analysis | ● active (amended) |
| MENU-061 | static analysis | ● active (amended) |
| MENU-062 | static analysis | ● active (amended) |
| MENU-063 | static analysis | ● active (amended) |
| MENU-064 | static analysis | ● active |
| MENU-065 | static analysis | ● active |
| MENU-066 | static analysis | ● active |
| MENU-067 | static analysis | ● active |
| MENU-068 | static analysis | ● active |
| MENU-069 | static analysis | ● active |
| MENU-096 | static analysis | ● active |
| MENU-070 | static analysis | ● active |
| MENU-071 | static analysis | ● active |
| MENU-072 | static analysis | ● active |
| MENU-073 | static analysis | ● active |
| MENU-074 | static analysis | ● active |
| MENU-075 | static analysis | ● active |
| MENU-076 | static analysis | ● active |
| MENU-077 | static analysis | ● active |
| MENU-078 | static analysis | ● active |
| MENU-079 | static analysis | ● active |
| MENU-080 | static analysis | ● active |
| MENU-081 | static analysis | ● active (amended) |
| MENU-082 | static analysis | ● active (amended) |
| MENU-083 | static analysis | ● active |
| MENU-084 | static analysis | ● active |
| MENU-085 | static analysis | ● active |
| MENU-086 | static analysis | ✔ promoted (branch candidate) |
| MENU-087 | static analysis | ✔ promoted (branch candidate) |
| MENU-088 | static analysis | ✔ promoted (amended, partially retracted, branch candidate) |
| MENU-089 | static analysis | ✔ promoted (branch candidate) |
| MENU-090 | static analysis | ✔ promoted (branch candidate) |
| MENU-091 | static analysis | ✔ promoted (branch candidate) |
| MENU-092 | static analysis | ✔ promoted (branch candidate) |
| MENU-093 | static analysis | ✔ promoted (branch candidate) |
| MENU-094 | static analysis | ✔ promoted (branch candidate) |
| MENU-095 | static analysis | ✔ promoted (branch candidate) |
| MENU-099 | static analysis | ✔ promoted (branch candidate) |
| MENU-100 | static analysis | ✔ promoted (branch candidate) |
| MENU-101 | static analysis | ✔ promoted (branch candidate) |
| MENU-102 | static analysis | ✔ promoted (branch candidate) |
| MENU-103 | static analysis | ✔ promoted (branch candidate) |
| MENU-104 | static analysis | ✔ promoted (branch candidate) |
| MENU-105 | static analysis | ✔ promoted (branch candidate) |
| MENU-106 | static analysis | ✔ promoted (branch candidate) |
| MENU-107 | static analysis | ✔ promoted (branch candidate) |
| MENU-108 | static analysis | ✔ promoted (branch candidate) |
| MENU-109 | static analysis | ✔ promoted (branch candidate) |
| MENU-110 | static analysis | ✔ promoted (branch candidate) |
| MENU-111 | static analysis | ✔ promoted (branch candidate) |
| MENU-112 | static analysis | ✔ promoted (branch candidate) |
| MENU-113 | static analysis | ✔ promoted (branch candidate) |
| MENU-114 | static analysis | ✔ promoted (branch candidate) |

## Ledger mission

| ID | Method | Status |
|---|---|---|
| MISSION-MSGLINE-056 | static analysis | ● active |
| MISSION-MSGLINE-057 | static analysis | ● active |
| MISSION-MSGPOST-058 | static analysis | ● active |
| MISSION-SCATTER-053 | static analysis | ✔ promoted |
| MISSION-DEFEAT-045 | static analysis | ● active |
| MISSION-DEFEAT-046 | static analysis | ● active |
| MISSION-START-001 | static analysis | ● active (amended, partially retracted, superseded) |
| MISSION-DROP-002 | static analysis | ● active (amended, superseded) |
| MISSION-WIN-003 | mixed | ● active (amended, partially retracted) |
| MISSION-VIP-004 | static analysis | ● active (amended, partially retracted) |
| MISSION-TEXT-005 | static analysis | ● active |
| MISSION-ARM-006 | static analysis | ● active |
| MISSION-DEF-007 | static analysis | ● active (amended) |
| MISSION-SLOT-008 | static analysis | ● active |
| MISSION-M10-009 | mixed | ● active |
| MISSION-TYP-010 | mixed | ● active (amended, partially retracted) |
| MISSION-DROP-011 | static analysis | ● active (amended) |
| MISSION-SEAT-012 | static analysis | ● active (amended, partially retracted) |
| MISSION-END-013 | static analysis | ● active (partially retracted, amended) |
| MISSION-LATCH-014 | static analysis | ● active |
| MISSION-PATH-015 | static analysis | ● active (partially retracted, amended) |
| MISSION-STOP-016 | static analysis | ● active (partially retracted, amended) |
| MISSION-ROOT-017 | mixed | ● active (amended, partially retracted) |
| MISSION-VIP-018 | mixed | ● active |
| MISSION-VIEW-019 | static analysis | ● active |
| MISSION-VIEW-020 | static analysis | ● active |
| MISSION-DOC-021 | static analysis | ● active |
| MISSION-MONEY-022 | static analysis | ✔ promoted |
| MISSION-DOC-023 | static analysis | ✔ promoted |
| MISSION-M30-024 | observation | ● active (amended) |
| MISSION-LOSE-025 | observation | ● active |
| MISSION-CURE-026 | observation | ● active (amended) |
| MISSION-PAIR-027 | observation | ● active |
| MISSION-TYP-028 | observation | ● active |
| MISSION-VICTORY-029 | mixed | ● active |
| MISSION-VICTORY-030 | static analysis | ● active |
| MISSION-VICTORY-031 | mixed | ● active |
| MISSION-VICTORY-032 | static analysis | ● active |
| MISSION-VICTORY-033 | mixed | ● active |
| MISSION-VICTORY-034 | static analysis | ● active |
| MISSION-VICTORY-035 | static analysis | ● active (partially retracted, amended) |
| MISSION-VICTORY-036 | mixed | ● active |
| MISSION-DEFEAT-059 | static analysis | ● active (amended) |
| MISSION-LOAD-060 | static analysis | ● active |
| MISSION-LOAD-061 | static analysis | ● active (amended) |
| MISSION-063 | static analysis | ● active |
| MISSION-064 | static analysis | ● active |
| MISSION-065 | static analysis | ● active |
| MISSION-066 | static analysis | ● active |
| MISSION-067 | static analysis | ● active |
| MISSION-068 | static analysis | ● active |
| MISSION-069 | static analysis | ● active |
| MISSION-REPORT-070 | observation | ● active |
| MISSION-MAP-073 | static analysis | ● active |
| MISSION-DROP-074 | static analysis | ● active |

## Ledger move

| ID | Method | Status |
|---|---|---|
| MOVE-SEARCH-001 | static analysis | ● active |
| MOVE-COST-002 | static analysis | ● active |
| MOVE-TERM-003 | static analysis | ● active (amended, partially retracted) |
| MOVE-ROUTE-004 | static analysis | ● active |
| MOVE-PLANE-005 | static analysis | ● active |
| MOVE-PARAM-006 | static analysis | ● active |
| MOVE-CLAIM-007 | static analysis | ● active |
| MOVE-WAIT-008 | static analysis | ● active |
| MOVE-TICK-009 | static analysis | ● active (partially retracted) |
| MOVE-STEP-010 | static analysis | ● active (amended, partially retracted) |
| MOVE-SPEED-011 | static analysis | ● active |
| MOVE-REFRESH-012 | static analysis | ● active |
| MOVE-TICK-013 | static analysis | ● active |
| MOVE-TICK-014 | static analysis | ● active (contested) |
| MOVE-TICK-015 | static analysis | ● active |
| MOVE-ID-016 | static analysis | ● active |
| MOVE-TICK-017 | static analysis | ● active |
| MOVE-ALT-018 | static analysis | ● active |
| MOVE-ALT-019 | static analysis | ● active |
| MOVE-ALT-020 | static analysis | ● active (amended) |
| MOVE-ALT-021 | static analysis | ● active |
| MOVE-ALT-022 | static analysis | ● active |
| MOVE-ORDER-023 | static analysis | ● active (partially retracted) |
| MOVE-DOM-024 | static analysis | ● active |
| MOVE-DOM-025 | static analysis | ● active |
| MOVE-DOM-026 | mixed | ● active |
| MOVE-DOM-027 | static analysis | ● active |
| MOVE-DOM-028 | static analysis | ● active |
| MOVE-RATE-029 | static analysis | ● active |
| MOVE-GROUP-030 | static analysis | ● active (amended, partially retracted) |
| MOVE-TURN-031 | static analysis | ● active (partially retracted) |
| MOVE-CLOCK-032 | static analysis | ● active |
| MOVE-LIMIT-033 | static analysis | ● active |
| MOVE-DIR-034 | static analysis | ● active |
| MOVE-GATE-035 | static analysis | ● active |
| MOVE-FORM-036 | static analysis | ● active |
| MOVE-GROUP-037 | static analysis | ● active |
| MOVE-AREA-038 | static analysis | ● active |
| MOVE-GATE-039 | static analysis | ● active |
| MOVE-STEP-040 | static analysis | ● active |
| MOVE-EFFLIST-041 | static analysis | ● active |
| MOVE-072 | static analysis | ● active |
| MOVE-073 | static analysis | ● active |
| MOVE-TURN-044 | static analysis | ✔ promoted |
| MOVE-RATE-052 | static analysis | ✔ promoted |
| MOVE-RATE-053 | static analysis | ✔ promoted |
| MOVE-RATE-054 | static analysis | ✔ promoted |
| MOVE-RATE-055 | static analysis | ✔ promoted |
| MOVE-EVENT-060 | static analysis | ✔ promoted |
| MOVE-EVENT-061 | static analysis | ✔ promoted |
| MOVE-080 | static analysis | ● active |
| MOVE-081 | static analysis | ● active |
| MOVE-082 | static analysis | ● active |
| MOVE-083 | static analysis | ● active |
| MOVE-084 | static analysis | ● active |
| MOVE-085 | static analysis | ● active |
| MOVE-086 | static analysis | ● active |
| MOVE-087 | static analysis | ● active |
| MOVE-088 | static analysis | ● active (amended) |
| MOVE-089 | static analysis | ● active |
| MOVE-090 | static analysis | ● active |
| MOVE-093 | static analysis | ● active |
| MOVE-094 | static analysis | ● active |
| MOVE-095 | static analysis | ● active |
| MOVE-096 | static analysis | ● active |
| MOVE-097 | static analysis | ● active |
| MOVE-098 | static analysis | ● active |
| MOVE-099 | static analysis | ● active |
| MOVE-100 | static analysis | ● active |
| MOVE-105 | static analysis | ✔ promoted |
| MOVE-106 | static analysis | ✔ promoted |
| MOVE-107 | static analysis | ✔ promoted |

## Ledger pal

| ID | Method | Status |
|---|---|---|
| PAL-FILE-001 | static analysis | ● active |
| PAL-KEY-002 | static analysis | ● active |
| PAL-NAME-003 | static analysis | ● active |
| PAL-TIER-004 | static analysis | ● active |
| PAL-FACE-005 | static analysis | ● active |
| PAL-XFORM-006 | mixed | ● active |
| PAL-OWN-007 | static analysis | ● active (amended) |
| PAL-CORP-008 | mixed | ● active |
| PAL-LIMIT-009 | static analysis | ● active |
| PAL-MODE4-010 | static analysis |  |
| PAL-PROJ-011 | static analysis | ● active |
| PAL-SHADE-012 | static analysis | ● active (amended) |
| PAL-SHADE-013 | static analysis | ● active (amended) |
| PAL-FIGURE-014 | static analysis | ● active (amended) |
| PAL-SLOT-015 | mixed | ● active |
| PAL-BAND-016 | static analysis | ● active |
| PAL-LIMIT-017 | static analysis | ● active |
| PAL-RULE-021 | static analysis | ● active |
| PAL-FIRST-022 | static analysis | ● active |
| PAL-JOIN-023 | static analysis | ● active |
| PAL-BLIT-024 | static analysis | ● active |

## Ledger party

| ID | Method | Status |
|---|---|---|
| PARTY-OWN-001 | static analysis | ● active |
| PARTY-ROSTER-002 | static analysis | ● active (amended, superseded) |
| PARTY-FLAG-003 | static analysis | ● active (amended, partially retracted) |
| PARTY-CULL-004 | static analysis | ● active |
| PARTY-CARRY-005 | static analysis | ● active (amended) |
| PARTY-LOSS-006 | static analysis | ● active |
| PARTY-MERC-007 | mixed | ● active (amended) |
| PARTY-SESSION-008 | static analysis | ● active (partially retracted) |
| PARTY-GROUP-009 | static analysis | ● active (amended) |
| PARTY-ORIGIN-010 | static analysis | ● active (partially retracted) |
| PARTY-WRITE-011 | static analysis | ● active (amended) |
| PARTY-INSTALL-012 | static analysis | ● active |
| PARTY-GATE-013 | static analysis | ● active |
| PARTY-PERSIST-014 | static analysis | ● active (amended) |
| PARTY-MONEY-015 | static analysis | ● active (partially retracted) |
| PARTY-MONEY-016 | static analysis | ✔ promoted |
| PARTY-ADDHERO-017 | static analysis | ● active |
| PARTY-MONEY-018 | mixed | ● active (partially retracted) |
| PARTY-MONEY-024 | static analysis | ✔ promoted |
| PARTY-JOIN-025 | static analysis | ● active |
| PARTY-ENDCULL-026 | static analysis | ● active |
| PARTY-BAND-027 | static analysis | ✖ retracted |
| PARTY-PERSIST-028 | static analysis | ● active (amended, superseded) |
| PARTY-JOINCORPUS-029 | mixed | ● active (amended, partially retracted) |
| PARTY-M20-030 | mixed | ● active |
| PARTY-M20-031 | static analysis | ● active |
| PARTY-M20-032 | mixed | ● active |

## Ledger reg

| ID | Method | Status |
|---|---|---|
| REG-099 | static analysis | ✔ promoted |
| REG-100 | static analysis | ✔ promoted |
| REG-101 | static analysis | ✔ promoted |
| REG-102 | static analysis | ✔ promoted |
| REG-LOC-016 | mixed | ● active (amended) |
| REG-FMT-017 | mixed | ● active (amended) |
| REG-UNITS-018 | static analysis | ● active (amended) |
| REG-ROSTER-019 | mixed | ● active (partially retracted) |
| REG-OBJ-020 | mixed | ● active (amended) |
| REG-OBJ-039 | observation | ● active |
| REG-STR-021 | mixed | ● active (amended) |
| REG-STR-040 | observation | ● active (partially retracted) |
| REG-PROJ-022 | mixed | ● active (amended) |
| REG-PROJ-041 | observation | ● active |
| EDITOR-023 | mixed | ● active |
| REG-VAL-024 | mixed | ● active (amended) |
| REG-VAL-025 | mixed | ● active (amended) |
| REG-VAL-026 | mixed | ● active (amended) |
| REG-VAL-027 | mixed | ● active (amended) |
| REG-VAL-028 | mixed | ● active (amended) |
| REG-VAL-029 | mixed | ● active (amended) |
| REG-VAL-030 | mixed | ● active (partially retracted) |
| REG-FMT-031 | static analysis | ● active |
| REG-REC-032 | static analysis | ● active (amended, partially retracted) |
| REG-KIND-033 | static analysis | ● active |
| REG-KIND-034 | static analysis | ● active (partially retracted) |
| REG-DBL-035 | static analysis | ● active |
| REG-TEXT-036 | static analysis | ● active (amended) |
| REG-TEXT-037 | mixed | ● active |
| REG-LOC-038 | static analysis | ● active |
| REG-MAT-042 | observation | ● active |
| REG-VAL-043 | observation | ● active |
| REG-KEY-044 | static analysis | ● active (amended) |
| REG-KEY-045 | static analysis | ● active |
| REG-OBJ-046 | static analysis | ● active |
| REG-OBJ-047 | static analysis | ● active |
| REG-UNITS-049 | static analysis | ● active |
| REG-UNITS-050 | static analysis | ● active (amended) |
| REG-UNITS-051 | static analysis | ● active |
| REG-ROSTER-052 | static analysis | ● active |
| REG-CUT-053 | static analysis | ● active |
| REG-KEY-054 | static analysis | ● active (partially retracted) |
| REG-NAME-055 | static analysis | ● active |
| REG-KIND-056 | static analysis | ● active |
| REG-SFX-057 | mixed | ● active |
| REG-NPC-058 | static analysis | ● active (amended) |
| REG-SCN-059 | static analysis | ● active (amended) |
| REG-AI-060 | mixed | ● active |
| REG-UNITS-061 | static analysis | ✔ promoted (partially retracted) |
| REG-SCN-062 | static analysis | ● active (amended, partially retracted) |
| REG-SCN-063 | static analysis | ● active |
| REG-SCN-064 | static analysis | ● active (amended, partially retracted) |
| REG-SCN-065 | static analysis | ● active |
| REG-GMAP-066 | static analysis | ● active (amended, partially retracted) |
| REG-SCN-067 | static analysis | ● active |
| REG-STR-080 | static analysis | ● active |
| REG-STR-081 | static analysis | ● active |
| REG-STR-082 | static analysis | ● active |
| REG-PICT-083 | static analysis | ● active |
| REG-PROJ-086 | static analysis | ● active (amended, superseded) |
| REG-PROJ-087 | static analysis | ● active |
| REG-NPC-088 | static analysis | ● active |
| REG-NPC-089 | static analysis | ● active |
| REG-NPC-090 | static analysis | ● active |
| REG-NPC-091 | static analysis | ● active |
| REG-DESC-096 | mixed | ● active |
| REG-SCN-097 | static analysis | ● active |
| REG-SCN-098 | static analysis | ● active |
| REG-104 | static analysis | ✔ promoted |
| REG-105 | static analysis | ✔ promoted |
| REG-106 | static analysis | ✔ promoted |
| REG-107 | mixed | ✔ promoted |
| REG-108 | static analysis | ✔ promoted |
| REG-109 | static analysis | ✔ promoted |
| REG-110 | static analysis | ✔ promoted |
| REG-111 | static analysis | ✔ promoted |
| REG-INNCLOSE-116 | mixed | ● active |
| REG-118 | static analysis | ● active (partially retracted) |
| REG-INNSTAGE-117 | static analysis | ● active |

## Ledger res

| ID | Method | Status |
|---|---|---|
| RES-MAGIC-001 | mixed | ✔ promoted |
| RES-HDR-002 | static analysis | ✔ promoted |
| RES-HDR-003 | mixed | ✔ promoted (amended) |
| RES-HDR-004 | mixed | ✔ promoted (amended) |
| RES-HDR-005 | static analysis | ✔ promoted |
| RES-HDR-006 | mixed | ● active |
| RES-NODE-007 | static analysis | ✔ promoted |
| RES-NODE-008 | mixed | ✔ promoted |
| RES-TREE-009 | mixed | ✔ promoted (amended) |
| RES-SCOPE-010 | mixed | ✔ promoted |
| RES-NODE-011 | static analysis | ● active (amended) |
| RES-HDR-012 | mixed | ● active (amended, superseded) |
| RES-HDR-013 | mixed | ● active (amended, partially retracted) |
| RES-NODE-014 | mixed | ● active (amended) |
| RES-SCOPE-015 | static analysis | ● active |
| RES-NODE-016 | static analysis | ● active |
| RES-HDR-017 | static analysis | ● active |
| RES-HDR-018 | static analysis | ● active (partially retracted) |
| RES-NODE-019 | static analysis | ● active (partially retracted) |
| RES-CODE-020 | static analysis | ● active (partially retracted) |
| RES-TEXT-021 | static analysis | ● active |
| RES-TEXT-022 | mixed | ● active |
| RES-LOOKUP-023 | static analysis | ● active (partially retracted) |
| RES-IDENT-024 | static analysis | ● active (amended) |
| RES-OPEN-026 | static analysis | ● active |
| RES-PATH-025 | static analysis | ● active (partially retracted) |
| RES-OPEN-027 | static analysis | ● active |
| RES-GEOM-028 | mixed | ● active |
| RES-HDR-029 | mixed | ● active |
| RES-HDR-030 | mixed | ● active |
| RES-SET-032 | static analysis | ● active |
| RES-ORDER-033 | static analysis | ● active |
| RES-IDENT-034 | static analysis | ● active |
| RES-MASK-035 | static analysis | ● active |
| RES-CASE-036 | static analysis | ● active |
| RES-DIR-037 | static analysis | ● active |
| RES-COLL-038 | mixed | ● active |
| RES-ACCEPT-031 | mixed | ● active (amended) |
| RES-HDR-001 | mixed | ● active (amended) |
| RES-039 | static analysis | ✔ promoted |
| RES-040 | static analysis | ✔ promoted |
| RES-041 | static analysis | ✔ promoted |

## Ledger rom2-asset

| ID | Method | Status |
|---|---|---|
| R2-ASSET-001 | observation | ✔ promoted |
| R2-ASSET-002 | observation | ✔ promoted |
| R2-ASSET-003 | observation | ✔ promoted |
| R2-ASSET-004 | observation | ✔ promoted |
| R2-ASSET-005 | observation | ✔ promoted |
| R2-ASSET-006 | observation | ✔ promoted |
| R2-ASSET-007 | observation | ✔ promoted |
| R2-ASSET-008 | observation | ✔ promoted |
| R2-ASSET-009 | observation | ✔ promoted |
| R2-ASSET-010 | observation | ✔ promoted |
| R2-ASSET-011 | observation | ✔ promoted |
| R2-ASSET-012 | observation | ✔ promoted |
| R2-ASSET-017 | mixed | ✔ promoted |
| R2-ASSET-018 | mixed | ✔ promoted |
| R2-ASSET-019 | static analysis | ✔ promoted |
| R2-ASSET-020 | static analysis | ✔ promoted |
| R2-ASSET-021 | static analysis | ✔ promoted |
| R2-ASSET-022 | static analysis | ✔ promoted |
| R2-ASSET-023 | static analysis | ✔ promoted |
| R2-ASSET-024 | static analysis | ✔ promoted |
| R2-ASSET-025 | mixed | ✔ promoted |
| R2-ASSET-026 | static analysis | ✔ promoted |
| R2-ASSET-027 | static analysis | ✔ promoted (partially retracted: maxima attribution only; see retracted.md) |
| R2-ASSET-029 | static analysis | ✔ promoted |
| R2-ASSET-030 | mixed | ✔ promoted |
| R2-ASSET-031 | mixed | ✔ promoted |
| R2-ASSET-032 | mixed | ✔ promoted |
| R2-ASSET-033 | static analysis | ✔ promoted |

## Ledger rom2-engine

| ID | Method | Status |
|---|---|---|
| R2-ENGINE-001 | mixed | ● active |
| R2-ENGINE-002 | mixed | ● active |
| R2-ENGINE-003 | mixed | ● active (amended, partially retracted) |
| R2-ENGINE-004 | mixed | ● active |
| R2-ENGINE-005 | mixed | ● active (amended, partially retracted) |
| R2-ENGINE-006 | mixed | ● active (amended, partially retracted) |
| R2-ENGINE-007 | mixed | ● active (amended, partially retracted) |
| R2-ENGINE-008 | static analysis | ● active (amended, partially retracted) |
| R2-ENGINE-009 | mixed | ● active (amended) |
| R2-ENGINE-010 | mixed | ● active |
| R2-ENGINE-011 | mixed | ● active |
| R2-ENGINE-017 | static analysis | ● active |
| R2-ENGINE-018 | static analysis | ● active |
| R2-ENGINE-019 | mixed | ● active |
| R2-ENGINE-020 | static analysis | ● active |
| R2-ENGINE-021 | static analysis | ● active |
| R2-ENGINE-022 | static analysis | ● active |
| R2-ENGINE-023 | mixed | ● active |
| R2-ENGINE-024 | static analysis | ● active |
| R2-ENGINE-025 | static analysis | ● active |
| R2-ENGINE-026 | static analysis | ● active |
| R2-ENGINE-027 | static analysis | ● active |
| R2-ENGINE-028 | static analysis | ● active |
| R2-ENGINE-029 | static analysis | ● active |
| R2-ENGINE-030 | static analysis | ● active |
| R2-ENGINE-031 | static analysis | ● active |
| R2-ENGINE-032 | mixed | ● active |
| R2-ENGINE-033 | static analysis | ● active |
| R2-ENGINE-034 | static analysis | ● active |
| R2-ENGINE-035 | static analysis | ● active |
| R2-ENGINE-036 | mixed | ● active |
| R2-ENGINE-037 | static analysis | ● active (amended) |
| R2-ENGINE-038 | static analysis | ● active |
| R2-ENGINE-039 | mixed | ● active |
| R2-ENGINE-040 | mixed | ● active |
| R2-ENGINE-041 | mixed | ● active |
| R2-ENGINE-042 | static analysis | ● active |
| R2-ENGINE-043 | static analysis | ● active |
| R2-ENGINE-044 | static analysis | ● active |
| R2-ENGINE-045 | static analysis | ● active |
| R2-ENGINE-046 | static analysis | ● active |
| R2-ENGINE-047 | static analysis | ● active |
| R2-ENGINE-048 | static analysis | ● active |
| R2-ENGINE-053 | static analysis | ● active |
| R2-ENGINE-054 | static analysis | ● active |
| R2-ENGINE-055 | static analysis | ● active |
| R2-ENGINE-056 | static analysis | ● active |
| R2-ENGINE-057 | static analysis | ● active (amended) |
| R2-ENGINE-058 | static analysis | ● active |
| R2-ENGINE-059 | static analysis | ● active (amended) |
| R2-ENGINE-060 | static analysis | ● active |
| R2-ENGINE-049 | static analysis | ✔ promoted (branch candidate) |
| R2-ENGINE-050 | static analysis | ✔ promoted (branch candidate) |
| R2-ENGINE-051 | static analysis | ✔ promoted (branch candidate) |
| R2-ENGINE-052 | static analysis | ✔ promoted (branch candidate) |
| R2-ENGINE-069 | static analysis | ✔ promoted |
| R2-ENGINE-070 | static analysis | ✔ promoted |
| R2-ENGINE-071 | static analysis | ✔ promoted |
| R2-ENGINE-072 | static analysis | ✔ promoted |
| R2-ENGINE-073 | mixed | ● active |
| R2-ENGINE-074 | static analysis | ● active |
| R2-ENGINE-075 | mixed | ● active |
| R2-ENGINE-076 | mixed | ● active |
| R2-ENGINE-077 | static analysis | ● active |
| R2-ENGINE-085 | static analysis | ✔ promoted |
| R2-ENGINE-086 | static analysis | ✔ promoted |
| R2-ENGINE-087 | static analysis | ✔ promoted |
| R2-ENGINE-093 | static analysis | ● active |
| R2-ENGINE-094 | static analysis | ● active |
| R2-ENGINE-095 | static analysis | ● active |
| R2-ENGINE-096 | static analysis | ● active |
| R2-ENGINE-097 | static analysis | ● active |
| R2-ENGINE-105 | static analysis | ● active |
| R2-ENGINE-106 | static analysis | ● active |
| R2-ENGINE-107 | static analysis | ● active |
| R2-ENGINE-108 | static analysis | ● active |
| R2-ENGINE-123 | static analysis | ✔ promoted |
| R2-ENGINE-124 | static analysis | ✔ promoted |
| R2-ENGINE-125 | static analysis | ✔ promoted |
| R2-ENGINE-126 | mixed | ✔ promoted |
| R2-ENGINE-127 | mixed | ✔ promoted |
| R2-ENGINE-135 | mixed | ● active |
| R2-ENGINE-136 | static analysis | ● active |
| R2-ENGINE-137 | static analysis | ● active |
| R2-ENGINE-138 | static analysis | ● active |
| R2-ENGINE-139 | static analysis | ● active |
| R2-ENGINE-140 | static analysis | ● active |
| R2-ENGINE-141 | mixed | ● active |
| R2-ENGINE-143 | static analysis | ● active |
| R2-ENGINE-144 | static analysis | ● active |
| R2-ENGINE-145 | static analysis | ● active |
| R2-ENGINE-146 | static analysis | ● active |
| R2-ENGINE-147 | static analysis | ● active |
| R2-ENGINE-148 | mixed | ● active |
| R2-ENGINE-149 | static analysis | ● active |
| R2-ENGINE-150 | static analysis | ● active |
| R2-ENGINE-159 | static analysis | ● active |
| R2-ENGINE-160 | static analysis | ● active |
| R2-ENGINE-161 | static analysis | ● active |
| R2-ENGINE-151 | static analysis | ✔ promoted |
| R2-ENGINE-152 | static analysis | ✔ promoted |
| R2-ENGINE-153 | static analysis | ✔ promoted |
| R2-ENGINE-154 | static analysis | ✔ promoted |
| R2-ENGINE-155 | static analysis | ✔ promoted |
| R2-ENGINE-156 | static analysis | ✔ promoted |
| R2-ENGINE-157 | static analysis | ✔ promoted |
| R2-ENGINE-158 | mixed | ✔ promoted |
| R2-ENGINE-167 | static analysis | ● active |
| R2-ENGINE-168 | static analysis | ● active |
| R2-ENGINE-191 | static analysis | ✔ promoted (amended) |
| R2-ENGINE-192 | static analysis | ✔ promoted |
| R2-ENGINE-199 | static analysis | ✔ promoted |
| R2-ENGINE-200 | static analysis | ✔ promoted |

## Ledger rom2-session

| ID | Method | Status |
|---|---|---|
| R2-SESSION-001 | static analysis | ● active |
| R2-SESSION-002 | static analysis | ● active |
| R2-SESSION-003 | static analysis | ✔ promoted |
| R2-SESSION-004 | static analysis | ● active |
| R2-SESSION-005 | static analysis | ● active |
| R2-SESSION-006 | static analysis | ● active |
| R2-SESSION-007 | static analysis | ● active |
| R2-SESSION-008 | mixed | ● active |
| R2-SESSION-009 | static analysis | ● active |
| R2-SESSION-010 | static analysis | ✔ promoted |
| R2-SESSION-017 | static analysis | ✔ promoted |
| R2-SESSION-018 | static analysis | ✔ promoted |
| R2-SESSION-019 | static analysis | ✔ promoted |
| R2-SESSION-020 | static analysis | ✔ promoted |
| R2-SESSION-021 | static analysis | ✔ promoted |
| R2-SESSION-022 | static analysis | ✔ promoted |
| R2-SESSION-023 | static analysis | ● active |
| R2-SESSION-031 | static analysis | ● active (branch candidate) |
| R2-SESSION-032 | static analysis | ● active (branch candidate) |
| R2-SESSION-033 | static analysis | ● active (branch candidate) |
| R2-SESSION-034 | static analysis | ● active (branch candidate) |
| R2-SESSION-035 | static analysis | ● active (branch candidate) |
| R2-SESSION-036 | static analysis | ● active (branch candidate) |
| R2-SESSION-037 | static analysis | ● active (branch candidate) |
| R2-SESSION-047 | static analysis | ● active |
| R2-SESSION-048 | static analysis | ● active |
| R2-SESSION-049 | static analysis | ● active |
| R2-SESSION-050 | static analysis | ● active |
| R2-SESSION-051 | static analysis | ● active |
| R2-SESSION-052 | static analysis | ● active |
| R2-SESSION-059 | static analysis | ● active |
| R2-SESSION-060 | static analysis | ● active (amended) |
| R2-SESSION-067 | static analysis | ● active |
| R2-SESSION-068 | static analysis | ● active |
| R2-SESSION-071 | static analysis | ● active |
| R2-SESSION-072 | static analysis | ● active |
| R2-SESSION-075 | observation | ● active (branch candidate) |
| R2-SESSION-076 | observation | ● active (branch candidate) |
| R2-SESSION-077 | observation | ● active (branch candidate) |
| R2-SESSION-078 | observation | ● active (branch candidate) |
| R2-SESSION-079 | observation | ● active (branch candidate) |
| R2-SESSION-083 | mixed | ✔ promoted (amended) |
| R2-SESSION-091 | mixed | ✔ promoted |

## Ledger sav

| ID | Method | Status |
|---|---|---|
| SAV-994 | static analysis | ✔ promoted |
| SAV-995 | static analysis | ✔ promoted |
| SAV-996 | static analysis | ✔ promoted |
| SAV-997 | static analysis | ✔ promoted |
| SAV-998 | mixed | ✔ promoted |
| SAV-982 | static analysis | ✔ promoted |
| SAV-983 | static analysis | ✔ promoted |
| SAV-984 | static analysis | ✔ promoted |
| SAV-970 | static analysis | ✔ promoted |
| SAV-971 | static analysis | ✔ promoted |
| SAV-972 | static analysis | ✔ promoted |
| SAV-934 | static analysis | ✔ promoted |
| SAV-935 | static analysis | ✔ promoted |
| SAV-936 | static analysis | ✔ promoted |
| SAV-958 | static analysis | ✔ promoted |
| SAV-959 | static analysis | ✔ promoted |
| SAV-960 | static analysis | ✔ promoted |
| SAV-961 | static analysis | ✔ promoted |
| SAV-962 | static analysis | ✔ promoted |
| SAV-926 | static analysis | ✔ promoted |
| SAV-927 | static analysis | ✔ promoted |
| SAV-928 | static analysis | ✔ promoted |
| SAV-929 | static analysis | ✔ promoted |
| SAV-930 | static analysis | ✔ promoted |
| SAV-931 | mixed | ✔ promoted |
| SAV-932 | mixed | ✔ promoted |
| SAV-918 | static analysis | ✔ promoted |
| SAV-919 | static analysis | ✔ promoted |
| SAV-920 | mixed | ✔ promoted |
| SAV-914 | static analysis | ✔ promoted |
| SAV-915 | static analysis | ✔ promoted |
| SAV-916 | static analysis | ✔ promoted |
| SAV-906 | static analysis | ✔ promoted |
| SAV-907 | static analysis | ✔ promoted |
| SAV-908 | static analysis | ✔ promoted |
| SAV-909 | static analysis | ✔ promoted |
| SAV-910 | static analysis | ✔ promoted |
| SAV-898 | static analysis | ✔ promoted |
| SAV-899 | static analysis | ✔ promoted |
| SAV-900 | mixed | ✔ promoted |
| SAV-890 | static analysis | ✔ promoted |
| SAV-891 | static analysis | ✔ promoted |
| SAV-892 | static analysis | ✔ promoted |
| SAV-893 | mixed | ✔ promoted |
| SAV-MOVRATE-866 | static analysis | ✔ promoted |
| SAV-854 | static analysis | ✔ promoted |
| SAV-855 | static analysis | ✔ promoted |
| SAV-842 | static analysis | ✔ promoted |
| SAV-843 | static analysis | ✔ promoted |
| SAV-844 | static analysis | ✔ promoted |
| SAV-845 | static analysis | ✔ promoted |
| SAV-846 | static analysis | ✔ promoted |
| SAV-847 | static analysis | ✔ promoted |
| SAV-848 | mixed | ✔ promoted |
| SAV-PLAYERIDENT-830 | static analysis | ✔ promoted |
| SAV-PLAYERPOP-831 | mixed | ✔ promoted |
| SAV-SACKENTRY-590 | static analysis | ✔ promoted |
| SAV-SACKREMOVE-591 | static analysis | ✔ promoted |
| SAV-SACKPLANES-592 | static analysis | ✔ promoted |
| SAV-SACKCALLER-593 | static analysis | ✔ promoted |
| SAV-CELLENTRY-582 | static analysis | ✔ promoted |
| SAV-CELLFAIL-583 | static analysis | ✔ promoted (amended) |
| SAV-CELLLEAVE-584 | static analysis | ✔ promoted |
| SAV-CROSSNEXT-585 | static analysis | ✔ promoted |
| SAV-GRPNEW-576 | static analysis | ✔ promoted |
| SAV-GRPALLOC-577 | static analysis | ✔ promoted |
| SAV-GRPCMD-578 | static analysis | ✔ promoted |
| SAV-GRPFIRSTSAVE-579 | static analysis | ✔ promoted |
| SAV-GRPDISPATCH-568 | static analysis | ✔ promoted |
| SAV-GRPMUTATE-569 | static analysis | ✔ promoted |
| SAV-GRPPATROL-570 | static analysis | ✔ promoted |
| SAV-PATROLCURSOR-571 | static analysis | ✔ promoted |
| SAV-GRPSAVENEXT-572 | static analysis | ✔ promoted |
| SAV-GRPLOAD-560 | static analysis | ✔ promoted |
| SAV-GRPOWNER-561 | static analysis | ✔ promoted |
| SAV-GRPIDENT-562 | static analysis | ✔ promoted |
| SAV-GRPAI-563 | static analysis | ✔ promoted |
| SAV-EQUIPORDER-552 | static analysis | ✔ promoted (amended) |
| SAV-EQUIPEFFECT-553 | static analysis | ✔ promoted |
| SAV-EQUIPCALL-554 | static analysis | ✔ promoted |
| SAV-EQUIPOBS-555 | static analysis | ✔ promoted |
| SAV-ACTORBIND-544 | static analysis | ✔ promoted |
| SAV-ACTORDISPLAY-545 | static analysis | ✔ promoted (amended) |
| SAV-ACTORCTOR-546 | static analysis | ✔ promoted |
| SAV-ACTORINPUT-547 | static analysis | ✔ promoted (amended, superseded) |
| SAV-ACTORLIMIT-548 | static analysis | ✔ promoted |
| SAV-MINCHILD-536 | observation | ✔ promoted |
| SAV-REGENWIDTH-528 | static analysis | ✔ promoted |
| SAV-REGENSTORE-529 | static analysis | ✔ promoted |
| SAV-REGENFAULT-530 | static analysis | ✔ promoted |
| SAV-REGENORDER-531 | static analysis | ✔ promoted |
| SAV-REGENWIRE-532 | static analysis | ✔ promoted |
| SAV-CITYMOVE-512 | static analysis | ✔ promoted |
| SAV-CITYSALE-513 | static analysis | ✔ promoted |
| SAV-CITYRETURN-514 | static analysis | ✔ promoted |
| SAV-CITYDERIVE-515 | static analysis | ✔ promoted |
| SAV-CITYSTORE-516 | static analysis | ✔ promoted |
| SAV-CHILDBOUNDARY-496 | observation | ✔ promoted |
| SAV-HUMALLOC-504 | static analysis | ✔ promoted |
| SAV-HUMNEW-505 | static analysis | ✔ promoted |
| SAV-HUMSEL-506 | mixed | ✔ promoted |
| SAV-HUMNEWSAVE-507 | static analysis | ✔ promoted |
| SAV-WHEADINIT-520 | static analysis | ✔ promoted |
| SAV-WHEADLOAD-521 | static analysis | ✔ promoted |
| SAV-WHEADMAP-522 | static analysis | ✔ promoted |
| SAV-WHEADHERO-523 | static analysis | ✔ promoted |
| SAV-WHEADWATCH-524 | static analysis | ✔ promoted |
| SAV-WHEADLIMIT-525 | static analysis | ✔ promoted |
| SAV-INITGUARD-488 | observation | ✔ promoted |
| SAV-HUMRUNTIME-476 | static analysis | ✔ promoted |
| SAV-HUMRUN-444 | mixed | ✔ promoted |
| SAV-HUMLOAD-445 | static analysis | ✔ promoted |
| SAV-HUMFOLD-446 | static analysis | ✔ promoted |
| SAV-HUMEQUIP-447 | static analysis | ✔ promoted |
| SAV-HUMMUT-448 | static analysis | ✔ promoted |
| SAV-HUMGAPS-449 | static analysis | ✔ promoted |
| SAV-HUMRESUME-460 | static analysis | ✔ promoted (amended, superseded) |
| SAV-HUMPROJECT-461 | static analysis | ✔ promoted |
| SAV-HUMTICK-462 | static analysis | ✔ promoted |
| SAV-HUMSTRIKE-463 | static analysis | ✔ promoted |
| SAV-HUMINDEX-464 | static analysis | ✔ promoted |
| SAV-HUMFIRST-465 | mixed | ✔ promoted |
| SAV-HDR-001 | mixed | ● active (amended, superseded) |
| SAV-VER-002 | static analysis | ● active (amended, partially retracted) |
| SAV-PTR-003 | mixed | ● active (amended) |
| SAV-EMB-004 | mixed | ● active (amended, partially retracted) |
| SAV-MAP-005 | mixed | ● active (amended) |
| SAV-UNK-006 | mixed | ● active (amended) |
| SAV-PACK-007 | mixed | ● active (amended) |
| SAV-SIZE-008 | mixed | ● active (amended) |
| SAV-EXT-009 | mixed | ● active (amended) |
| SAV-STREAM-010 | mixed | ● active (amended) |
| SAV-BLOCK-011 | static analysis | ● active |
| SAV-BLOCK-012 | static analysis | ● active |
| SAV-STREAM-013 | mixed | ● active (amended, superseded) |
| SAV-OBJ-014 | mixed | ● active (amended) |
| SAV-ID-015 | static analysis | ● active (amended) |
| SAV-OBJ-016 | static analysis | ● active (amended) |
| SAV-CELLREC-017 | static analysis | ● active (amended) |
| SAV-TAIL-018 | static analysis | ● active (amended, superseded) |
| SAV-FRAME-021 | static analysis | ● active |
| SAV-CODEC-022 | static analysis | ● active |
| SAV-SHAPE-023 | static analysis | ● active (amended) |
| SAV-ROSTER-024 | static analysis | ● active |
| SAV-HEAD-025 | static analysis | ● active |
| SAV-TRAIL-026 | static analysis | ● active |
| SAV-FLAG-027 | static analysis | ● active |
| SAV-PLAYER-028 | static analysis | ● active |
| SAV-OBF-029 | static analysis | ● active |
| SAV-CITY-030 | static analysis | ● active |
| SAV-SESS-031 | static analysis | ● active |
| SAV-CELLREC-032 | static analysis | ● active (amended, superseded) |
| SAV-CLASS-033 | static analysis | ● active |
| SAV-TOKEN-034 | static analysis | ● active (amended) |
| SAV-PTRMAP-035 | static analysis | ● active (partially retracted) |
| SAV-MEMBER-036 | static analysis | ● active (amended, superseded) |
| SAV-BLDG-037 | static analysis | ● active |
| SAV-OBFCEN-038 | static analysis | ● active |
| SAV-EMBED-039 | static analysis | ● active |
| SAV-WLIST-040 | static analysis | ● active |
| SAV-SPELLBK-041 | static analysis | ● active |
| SAV-DIARY-042 | static analysis | ● active |
| SAV-HUMAN-043 | static analysis | ● active |
| SAV-SPELL-044 | static analysis | ● active |
| SAV-UNITLEN-045 | static analysis | ● active |
| SAV-EFFCHAIN-046 | static analysis | ● active (amended, partially retracted) |
| SAV-SERPOP-047 | mixed | ● active (amended) |
| SAV-OWNER-048 | static analysis | ● active |
| SAV-UNITFLD-049 | static analysis | ● active (amended) |
| SAV-CARRY-050 | static analysis | ● active (amended, partially retracted) |
| SAV-DEATH-051 | static analysis | ● active (amended, partially retracted) |
| SAV-TOPLVL-052 | static analysis | ● active (amended, partially retracted) |
| SAV-DOC-053 | static analysis | ● active (amended, partially retracted) |
| SAV-PLDIARY-054 | static analysis | ● active |
| SAV-SEEN-055 | static analysis | ✖ retracted |
| SAV-TERRKEY-056 | static analysis | ● active |
| SAV-LOAD-057 | static analysis | ● active |
| SAV-GRPORD-058 | static analysis | ● active |
| SAV-HERO-059 | static analysis | ● active |
| SAV-GRPFLD-060 | static analysis | ● active (amended) |
| SAV-FOG-061 | static analysis | ● active |
| SAV-TAILEXT-062 | static analysis | ● active (amended) |
| SAV-HEROXP-063 | static analysis | ● active (amended) |
| SAV-HEROSKILL-064 | static analysis | ● active |
| SAV-HEROID-065 | mixed | ● active |
| SAV-CAMPTAIL-070 | static analysis | ✔ promoted |
| SAV-CAMPPROG-071 | mixed | ✔ promoted |
| SAV-CAMPPOS-072 | static analysis | ✔ promoted |
| SAV-CAMPMARK-073 | static analysis | ● active (partially retracted) |
| SAV-TOKENPOS-074 | static analysis | ● active (amended) |
| SAV-TOKENPTR-075 | static analysis | ● active (amended) |
| SAV-CAMPAIGN-076 | mixed | ● active |
| SAV-CAMPAIGN-077 | static analysis | ● active |
| SAV-CAMPAIGN-078 | static analysis | ● active |
| SAV-CAMPAIGN-079 | static analysis | ● active |
| SAV-CAMPAIGN-080 | static analysis | ● active |
| SAV-CAMPAIGN-081 | mixed | ● active |
| SAV-CAMPAIGN-082 | mixed | ● active |
| SAV-CAMPAIGN-083 | mixed | ● active |
| SAV-CAMPAIGN-084 | mixed | ● active |
| SAV-CAMPAIGN-085 | mixed | ● active |
| SAV-CAMPAIGN-086 | static analysis | ● active (amended, superseded) |
| SAV-CAMPAIGN-087 | mixed | ● active |
| SAV-CAMPAIGN-088 | static analysis | ● active |
| SAV-CELLLOAD-108 | static analysis | ● active |
| SAV-CELLLOAD-109 | static analysis | ● active |
| SAV-CELLLOAD-110 | static analysis | ● active |
| SAV-CELLLOAD-111 | static analysis | ● active |
| SAV-CELLLOAD-112 | static analysis | ● active |
| SAV-CELLLOAD-113 | mixed | ● active |
| SAV-TOKENLOAD-092 | static analysis | ● active |
| SAV-TOKENLOAD-093 | static analysis | ● active |
| SAV-TOKENLOAD-094 | static analysis | ● active |
| SAV-TOKENLOAD-095 | mixed | ● active |
| SAV-TOKENLOAD-096 | mixed | ● active |
| SAV-DEADLOAD-124 | static analysis | ● active |
| SAV-DEADLOAD-125 | static analysis | ✔ promoted |
| SAV-DEADLOAD-126 | mixed | ✔ promoted |
| SAV-DEADLOAD-127 | mixed | ● active |
| SAV-DEADLOAD-128 | mixed | ● active |
| SAV-DEADLOAD-129 | static analysis | ● active |
| SAV-DEADLOAD-130 | mixed | ✔ promoted |
| SAV-DEADLOAD-131 | static analysis | ● active |
| SAV-DEADLOAD-132 | mixed | ● active |
| SAV-POSLOAD-140 | static analysis | ✔ promoted |
| SAV-UNITPROG-156 | static analysis | ✔ promoted (amended) |
| SAV-UNITCORP-157 | mixed | ● active |
| SAV-TAGSCAN-158 | mixed | ● active |
| SAV-CLASSSER-172 | static analysis | ● active |
| SAV-CLASSSER-173 | static analysis | ● active |
| SAV-CLASSSER-174 | static analysis | ● active |
| SAV-CLASSSER-175 | static analysis | ● active |
| SAV-CLASSSER-176 | static analysis | ● active |
| SAV-CLASSSER-177 | mixed | ● active |
| SAV-CONTSER-188 | mixed | ● active |
| SAV-CONTSER-189 | static analysis | ● active |
| SAV-CONTSER-190 | mixed | ● active |
| SAV-CONTSER-191 | mixed | ● active |
| SAV-CONTSER-192 | mixed | ● active |
| SAV-CONTSER-193 | static analysis | ● active |
| SAV-CONTSER-194 | mixed | ● active |
| SAV-PRODROOT-204 | static analysis | ● active |
| SAV-PRODPOP-205 | static analysis | ● active |
| SAV-PRODGEN-206 | static analysis | ● active |
| SAV-PRODDIRECT-207 | mixed | ● active |
| SAV-LOADRESAVE-208 | mixed | ● active |
| SAV-PRODNOHIT-209 | mixed | ● active |
| SAV-POSTLOAD-220 | static analysis | ✔ promoted (partially retracted, amended) |
| SAV-POSTLOAD-221 | static analysis | ✔ promoted (partially retracted, superseded) |
| SAV-POSTLOAD-222 | static analysis | ✔ promoted |
| SAV-POSTLOAD-223 | static analysis | ✔ promoted |
| SAV-LABELTAIL-236 | static analysis | ● active |
| SAV-PHYSUFFIX-237 | static analysis | ● active |
| SAV-DECPAD-238 | mixed | ● active |
| SAV-EXTSURV-239 | mixed | ● active |
| SAV-WRITER-284 | observation | ● active |
| SAV-WORLDWRITER-316 | observation | ● active |
| SAV-ORIGLOAD-332 | mixed | ✔ promoted |
| SAV-ORIGCHOOSER-333 | mixed | ✔ promoted |
| SAV-SAVEDIR-334 | static analysis | ✔ promoted |
| SAV-ORIGFAULT-335 | static analysis | ✔ promoted |
| SAV-FULLREAD-252 | observation | ✔ promoted |
| SAV-ARCHREL-253 | observation | ✔ promoted |
| SAV-IDCENSUS-254 | observation | ✔ promoted |
| SAV-READPOP-255 | observation | ✔ promoted |
| SAV-RECON-268 | mixed | ● active (amended, partially retracted) |
| SAV-RECON-269 | mixed | ● active |
| SAV-RECON-270 | mixed | ● active |
| SAV-SUFF-300 | mixed | ● active |
| SAV-SUFF-301 | mixed | ● active |
| SAV-SUFF-302 | mixed | ● active |
| SAV-SUFF-303 | mixed | ● active |
| SAV-FIRSTTICK-348 | static analysis | ● active |
| SAV-FIRSTTICK-349 | static analysis | ● active |
| SAV-FIRSTTICK-350 | mixed | ● active |
| SAV-BOUNDARY-364 | observation | ✔ promoted |
| SAV-SPELLBK-365 | observation | ✔ promoted |
| SAV-EFFECTGRAPH-366 | observation | ✔ promoted |
| SAV-BOUNDARY-367 | observation | ✔ promoted |
| SAV-WRITERAUDIT-380 | observation | ✔ promoted |
| SAV-ORIGWRITER-396 | observation | ✔ promoted |
| SAV-ORIGMIN-397 | observation | ✔ promoted |
| SAV-ORIGCANON-398 | observation | ✔ promoted |
| SAV-ORIGVALUE-399 | observation | ✔ promoted |
| SAV-ORIGMISSION-400 | observation | ✔ promoted |
| SAV-LIVEFIELD-412 | observation | ✔ promoted |
| SAV-LIVEPROD-413 | observation | ✔ promoted |
| SAV-LIVEREADER-414 | observation | ✔ promoted |
| SAV-LIVEPGD-415 | observation | ✔ promoted |
| SAV-PROJSTORE-428 | mixed | ✔ promoted |
| SAV-PROJLOAD-429 | static analysis | ✔ promoted |
| SAV-PROJCORP-430 | mixed | ✔ promoted |
| SAV-WORLDSTATE-431 | mixed | ✔ promoted |
| SAV-WORLDFRONT-432 | mixed | ✔ promoted |
| SAV-WORLDDIAG-433 | mixed | ✔ promoted |
| SAV-598 | static analysis | ✔ promoted |
| SAV-599 | static analysis | ✔ promoted |
| SAV-600 | static analysis | ✔ promoted |
| SAV-601 | mixed | ✔ promoted |
| SAV-602 | static analysis | ✔ promoted |
| SAV-606 | static analysis | ✔ promoted |
| SAV-607 | static analysis | ✔ promoted |
| SAV-608 | mixed | ✔ promoted (partially retracted) |
| SAV-609 | static analysis | ✔ promoted |
| SAV-614 | mixed | ✔ promoted |
| SAV-615 | mixed | ✔ promoted |
| SAV-616 | mixed | ✔ promoted |
| SAV-617 | static analysis | ✔ promoted |
| SAV-622 | mixed | ✔ promoted |
| SAV-623 | mixed | ✔ promoted |
| SAV-624 | mixed | ✔ promoted |
| SAV-625 | mixed | ✔ promoted |
| SAV-626 | mixed | ✔ promoted |
| SAV-627 | mixed | ✔ promoted |
| SAV-628 | static analysis | ✔ promoted |
| SAV-629 | static analysis | ✔ promoted |
| SAV-630 | static analysis | ✔ promoted |
| SAV-631 | static analysis | ✔ promoted (amended) |
| SAV-632 | static analysis | ✔ promoted |
| SAV-633 | static analysis | ✔ promoted |
| SAV-634 | static analysis | ✔ promoted |
| SAV-635 | mixed | ✔ promoted (amended) |
| SAV-636 | static analysis | ✔ promoted (amended, partially retracted) |
| SAV-646 | static analysis | ✔ promoted |
| SAV-647 | static analysis | ✔ promoted (amended, partially retracted) |
| SAV-648 | static analysis | ✔ promoted |
| SAV-649 | static analysis | ✔ promoted |
| SAV-650 | static analysis | ✔ promoted |
| SAV-651 | static analysis | ✔ promoted (partially retracted) |
| SAV-652 | static analysis | ✔ promoted (partially retracted) |
| SAV-653 | static analysis | ✔ promoted |
| SAV-654 | mixed | ✔ promoted |
| SAV-655 | mixed | ✔ promoted |
| SAV-656 | mixed | ✔ promoted |
| SAV-657 | static analysis | ✔ promoted |
| SAV-658 | static analysis | ✔ promoted |
| SAV-662 | static analysis | ✔ promoted |
| SAV-663 | static analysis | ✔ promoted |
| SAV-664 | static analysis | ✔ promoted |
| SAV-665 | static analysis | ✔ promoted |
| SAV-666 | static analysis | ✔ promoted (amended) |
| SAV-667 | static analysis | ✔ promoted (amended) |
| SAV-668 | static analysis | ✔ promoted (amended, partially retracted, superseded) |
| SAV-669 | static analysis | ✔ promoted (amended, partially retracted) |
| SAV-670 | static analysis | ✔ promoted |
| SAV-671 | static analysis | ✔ promoted |
| SAV-672 | static analysis | ✔ promoted (amended) |
| SAV-694 | static analysis | ✔ promoted (amended, branch candidate) |
| SAV-758 | static analysis | ✔ promoted |
| SAV-726 | static analysis | ✔ promoted (branch candidate) |
| SAV-710 | static analysis | ✔ promoted (branch candidate) |
| SAV-711 | static analysis | ✔ promoted (branch candidate) |
| SAV-678 | static analysis | ✔ promoted (branch candidate) |
| SAV-774 | static analysis | ● active |
| SAV-775 | static analysis | ● active (partially retracted) |
| SAV-776 | static analysis | ● active (amended, partially retracted) |
| SAV-790 | static analysis | ✔ promoted |
| SAV-791 | static analysis | ✔ promoted (amended) |
| SAV-792 | static analysis | ✔ promoted |
| SAV-793 | static analysis | ✔ promoted |
| SAV-794 | static analysis | ✔ promoted |
| SAV-795 | mixed | ✔ promoted |
| SAV-796 | static analysis | ✔ promoted |
| SAV-797 | mixed | ✔ promoted |
| SAV-GRPORDER-806 | mixed | ● active |
| SAV-GRPLIST-807 | static analysis | ● active |
| SAV-TURNLOAD-822 | static analysis | ✔ promoted |
| SAV-946 | static analysis | ✔ promoted |
| SAV-947 | static analysis | ✔ promoted |
| SAV-948 | static analysis | ✔ promoted |
| SAV-949 | mixed | ✔ promoted |
| SAV-LOADREG-878 | static analysis | ✔ promoted |
| SAV-LOADHOOK-879 | static analysis | ✔ promoted |
| SAV-FIRSTMOVE-880 | static analysis | ✔ promoted |
| SAV-CASTCONT-1006 | static analysis | ● active |
| SAV-1010 | static analysis | ● active (partially retracted) |
| SAV-1011 | static analysis | ● active |
| SAV-SAVELABEL-1016 | static analysis | ● active |
| SAV-SAVELABEL-1017 | static analysis | ● active |
| SAV-SAVELABEL-1018 | mixed | ● active |
| SAV-1026 | static analysis | ● active |
| SAV-1027 | static analysis | ● active |
| SAV-1028 | static analysis | ● active |
| SAV-1029 | static analysis | ● active |
| SAV-1030 | static analysis | ● active |
| SAV-1031 | static analysis | ● active |
| SAV-1032 | static analysis | ● active |
| SAV-1033 | static analysis | ● active |
| SAV-1034 | static analysis | ● active |
| SAV-1037 | mixed | ● active (amended) |
| SAV-1042 | static analysis | ● active |
| SAV-1043 | static analysis | ● active (amended) |
| SAV-1047 | static analysis | ● active |
| SAV-1048 | static analysis | ● active |
| SAV-1050 | static analysis | ● active |
| SAV-1054 | static analysis | ● active |
| SAV-1055 | static analysis | ● active |
| SAV-1056 | static analysis | ● active |
| SAV-1058 | static analysis | ● active (partially retracted) |
| SAV-1059 | mixed | ● active |
| SAV-1060 | static analysis | ● active |
| SAV-1066 | static analysis | ● active |
| SAV-1068 | static analysis | ✔ promoted |
| SAV-1069 | static analysis | ✔ promoted |
| SAV-1070 | static analysis | ● active |
| SAV-1071 | static analysis | ● active |
| SAV-1074 | static analysis | ● active |
| SAV-1078 | static analysis | ● active |
| SAV-1079 | static analysis | ● active |
| SAV-1084 | static analysis | ● active |
| SAV-1085 | static analysis | ● active |
| SAV-1086 | mixed | ● active |
| SAV-1087 | static analysis | ● active |
| SAV-1088 | static analysis | ● active |
| SAV-1089 | static analysis | ● active |
| SAV-1091 | static analysis | ● active |
| SAV-1092 | static analysis | ● active |
| SAV-1093 | mixed | ● active |
| SAV-1094 | mixed | ● active |
| SAV-1095 | mixed | ● active |
| SAV-1096 | mixed | ● active |
| SAV-1097 | static analysis | ● active |
| SAV-1098 | mixed | ● active |
| SAV-1099 | mixed | ● active |
| SAV-1100 | mixed | ● active |
| SAV-1101 | mixed | ● active |
| SAV-1102 | static analysis | ● active |
| SAV-1103 | static analysis | ● active |
| SAV-1104 | static analysis | ● active |
| SAV-1105 | static analysis | ● active |
| SAV-1106 | static analysis | ● active |
| SAV-1107 | static analysis | ● active |
| SAV-1111 | mixed | ● active |
| SAV-1112 | static analysis | ● active |
| SAV-1113 | static analysis | ● active |
| SAV-1114 | static analysis | ● active |
| SAV-1115 | static analysis | ● active |
| SAV-1116 | static analysis | ● active |
| SAV-1118 | static analysis | ● active |
| SAV-1119 | static analysis | ● active |
| SAV-1120 | mixed | ● active |
| SAV-1121 | mixed | ● active |
| SAV-1123 | observation | ● active |
| SAV-1125 | static analysis | ✔ promoted |
| SAV-1126 | static analysis | ✔ promoted |
| SAV-1127 | mixed | ✔ promoted |
| SAV-1128 | static analysis | ✔ promoted |
| SAV-1129 | static analysis | ● active |
| SAV-1130 | static analysis | ● active |
| SAV-1131 | static analysis | ● active |
| SAV-1132 | static analysis | ● active |
| SAV-1133 | static analysis | ● active |
| SAV-1134 | mixed | ● active |
| SAV-1135 | static analysis | ● active |
| SAV-1136 | static analysis | ● active |
| SAV-1137 | static analysis | ● active |
| SAV-1138 | static analysis | ● active |
| SAV-1139 | mixed | ● active |
| SAV-1140 | static analysis | ● active |
| SAV-1141 | mixed | ● active |
| SAV-1142 | static analysis | ● active |
| SAV-1143 | mixed | ● active (amended) |
| SAV-1144 | static analysis | ● active |
| SAV-1145 | static analysis | ● active (amended) |
| SAV-1146 | static analysis | ● active |
| SAV-1147 | static analysis | ● active |
| SAV-1149 | static analysis | ● active (amended) |
| SAV-1150 | static analysis | ● active |
| SAV-1151 | static analysis | ● active |
| SAV-1152 | mixed | ● active |
| SAV-1153 | static analysis | ● active (amended) |
| SAV-1154 | static analysis | ● active (amended) |
| SAV-1155 | static analysis | ● active (amended) |
| SAV-1156 | static analysis | ● active |
| SAV-1157 | static analysis | ● active |
| SAV-1159 | static analysis | ● active |
| HERO-MODDK-161 | static analysis | ● active |
| HERO-DKIDX-162 | static analysis | ● active |
| SAV-1161 | static analysis | ● active |
| SAV-1163 | mixed | ● active |
| SAV-1164 | mixed | ● active |
| SAV-1165 | static analysis | ● active |
| SAV-1166 | static analysis | ● active |
| SAV-1169 | static analysis | ● active |
| SAV-1170 | mixed | ● active |
| SAV-1171 | static analysis | ● active |
| SAV-1172 | static analysis | ● active (amended) |
| SAV-1173 | static analysis | ● active |
| SAV-1180 | static analysis | ✔ promoted |

## Ledger session

| ID | Method | Status |
|---|---|---|
| SESS-DEFEAT-064 | static analysis | ● active |
| SESS-DEFEAT-065 | static analysis | ● active |
| SESS-INPUT-037 | static analysis | ● active (partially retracted) |
| SESS-OBJ-001 | static analysis | ● active |
| SESS-PHASE-002 | static analysis | ● active |
| SESS-SCREEN-003 | static analysis | ● active |
| SESS-TICK-004 | static analysis | ● active |
| SESS-CLOCK-005 | static analysis | ● active |
| SESS-TICK-006 | static analysis | ● active |
| SESS-IDLE-007 | static analysis | ● active |
| SESS-CMD-008 | static analysis | ● active (amended) |
| SESS-LOAD-009 | static analysis | ● active |
| SESS-MAP-010 | static analysis | ● active (partially retracted, amended, superseded) |
| SESS-END-011 | static analysis | ● active (amended, superseded) |
| SESS-LOSE-012 | static analysis | ● active |
| SESS-HERO-013 | static analysis | ● active |
| SESS-HERO-014 | static analysis | ● active (amended, superseded) |
| SESS-CMD-015 | static analysis | ● active |
| SESS-CMD-016 | static analysis | ● active |
| SESS-PARAM-017 | static analysis | ● active (partially retracted) |
| SESS-PACE-018 | static analysis | ● active (amended) |
| SESS-IDLE-019 | static analysis | ● active |
| SESS-PAUSE-020 | static analysis | ● active |
| SESS-CLOCK-021 | static analysis | ● active |
| SESS-TIMER-022 | static analysis | ● active |
| SESS-STATE-023 | static analysis | ● active |
| SESS-DIPLO-024 | static analysis | ● active |
| SESS-DIPLO-025 | static analysis | ● active |
| SESS-TICK-026 | static analysis | ● active |
| SESS-TICK-027 | static analysis | ● active |
| SESS-VIEW-028 | static analysis | ● active |
| SESS-VIEW-029 | static analysis | ● active |
| SESS-VIEW-030 | static analysis | ● active |
| SESS-VIEW-031 | static analysis | ● active (amended) |
| SESS-DEFNAME-032 | static analysis | ● active |
| SESS-NAMEGATE-033 | static analysis | ● active |
| SESS-START-034 | static analysis | ✔ promoted (amended) |
| SESS-PLAYER-035 | static analysis | ● active |
| SESS-START-036 | static analysis | ● active |
| SESS-COMPOSE-056 | static analysis | ● active |
| SESS-COMPOSE-057 | static analysis | ● active |
| SESS-COMPOSE-058 | static analysis | ● active |
| SESS-COMPOSE-059 | static analysis | ● active |
| SESS-COMPOSE-060 | static analysis | ● active |
| SESS-072 | static analysis | ✔ promoted |
| SESS-073 | static analysis | ✔ promoted |
| SESS-074 | static analysis | ✔ promoted |
| SESS-075 | static analysis | ✔ promoted |

## Ledger shop

| ID | Method | Status |
|---|---|---|
| SHOP-ANIMATION-081 | static analysis | ✔ promoted |
| SHOP-ANIMATION-082 | static analysis | ✔ promoted |
| SHOP-ANIMATION-083 | static analysis | ✔ promoted |
| SHOP-ANIMATION-084 | static analysis | ✔ promoted |
| SHOP-ANIMATION-085 | static analysis | ✔ promoted |
| SHOP-CONSUME-073 | static analysis | ✔ promoted |
| SHOP-CONSUME-074 | static analysis | ✔ promoted |
| SHOP-CLS-001 | static analysis | ● active |
| SHOP-OBJ-002 | static analysis | ● active (superseded, partially retracted) |
| SHOP-ENTRY-003 | static analysis | ● active (amended) |
| SHOP-CAP-004 | static analysis | ● active |
| SHOP-GEN-005 | static analysis | ● active (partially retracted) |
| SHOP-POOL-006 | static analysis | ● active (partially retracted) |
| SHOP-MAGIC-007 | static analysis | ● active (amended, partially retracted) |
| SHOP-RNG-008 | static analysis | ● active (amended) |
| SHOP-BUY-009 | static analysis | ● active (amended) |
| SHOP-SELL-010 | static analysis | ● active (amended) |
| SHOP-PRICE-011 | static analysis | ● active (amended) |
| SHOP-LIFE-013 | static analysis | ● active |
| SHOP-LIFE-014 | static analysis | ● active (amended) |
| SHOP-SAVE-015 | static analysis | ● active |
| SHOP-ENTRY-016 | static analysis | ● active (amended, contested) |
| SHOP-ROUND-017 | static analysis | ● active (partially retracted) |
| SHOP-NPC-012 | static analysis | ● active (amended, superseded) |
| SHOP-MISSION-018 | static analysis | ● active |
| SHOP-MISSION-019 | static analysis | ● active |
| SHOP-MISSION-020 | mixed | ● active (amended, partially retracted) |
| SHOP-POOL-021 | static analysis | ✔ promoted (amended, partially retracted) |
| SHOP-TOWN-022 | static analysis | ● active |
| SHOP-TOWN-023 | static analysis | ● active |
| SHOP-TRAY-024 | static analysis | ● active |
| SHOP-TRAY-025 | static analysis | ● active (amended, partially retracted) |
| SHOP-TRAY-026 | static analysis | ● active |
| SHOP-TRAY-027 | static analysis | ● active |
| SHOP-DUP-028 | static analysis | ● active (amended, partially retracted) |
| SHOP-DOC-029 | mixed | ✔ promoted |
| SHOP-SCREEN-030 | static analysis | ● active |
| SHOP-SCREEN-031 | static analysis | ● active |
| SHOP-SCREEN-032 | static analysis | ● active |
| SHOP-SCREEN-033 | static analysis | ● active |
| SHOP-SCREEN-034 | static analysis | ● active (partially retracted) |
| SHOP-SCREEN-035 | static analysis | ● active (amended) |
| SHOP-SCREEN-036 | static analysis | ● active (amended) |
| SHOP-SCREEN-037 | static analysis | ● active |
| SHOP-SCREEN-038 | static analysis | ● active |
| SHOP-SCREEN-039 | static analysis | ● active (amended) |
| SHOP-USABLE-040 | static analysis | ● active |
| SHOP-FIGURE-041 | static analysis | ● active |
| SHOP-FIGURE-042 | static analysis | ● active |
| SHOP-PICKER-043 | static analysis | ● active (amended, partially retracted) |
| SHOP-VIEW-044 | static analysis | ● active |
| SHOP-TIP-045 | static analysis | ● active |
| SHOP-MERCHANT-046 | static analysis | ● active |
| SHOP-SHELF-047 | static analysis | ● active |
| SHOP-MONEY-048 | static analysis | ● active |
| SHOP-LIMIT-049 | mixed | ● active |
| SHOP-050 | static analysis | ● active |
| SHOP-051 | static analysis | ● active |
| SHOP-052 | static analysis | ● active |
| SHOP-EFFPOOL-061 | mixed | ● active |
| SHOP-EFFWEIGHT-062 | mixed | ● active |
| SHOP-EFFPAY-063 | mixed | ● active |
| SHOP-EFFRANGE-064 | mixed | ● active |
| SHOP-EFFCAST-065 | mixed | ● active |
| SHOP-EFFPRICE-066 | mixed | ● active (partially retracted) |
| SHOP-EFFORDER-067 | mixed | ● active |
| SHOP-EFFRETRY-068 | mixed | ● active |
| SHOP-EFFCAP-069 | mixed | ● active |
| SHOP-EFFBASE-070 | mixed | ● active |
| SHOP-EFFALT-071 | static analysis | ● active |
| SHOP-EFFSEED-072 | mixed | ● active |
| SHOP-102 | static analysis | ✔ promoted |
| SHOP-103 | static analysis | ✔ promoted |
| SHOP-104 | static analysis | ✔ promoted |
| SHOP-105 | static analysis | ✔ promoted |
| SHOP-096 | static analysis | ● active |
| SHOP-097 | static analysis | ● active |
| SHOP-098 | static analysis | ● active |
| SHOP-099 | static analysis | ● active |
| SHOP-100 | static analysis | ● active |
| SHOP-101 | static analysis | ● active |
| SHOP-114 | static analysis | ● active (branch candidate) |
| SHOP-115 | static analysis | ● active (branch candidate) |
| SHOP-116 | static analysis | ● active (branch candidate) |
| SHOP-106 | static analysis | ● active (branch candidate) |
| SHOP-107 | static analysis | ● active (branch candidate) |
| SHOP-108 | static analysis | ● active (branch candidate) |
| SHOP-109 | static analysis | ● active (branch candidate) |
| SHOP-110 | static analysis | ● active (branch candidate) |
| SHOP-111 | static analysis | ● active (branch candidate) |
| SHOP-112 | static analysis | ● active (branch candidate) |

## Ledger spr16a

| ID | Method | Status |
|---|---|---|
| SPR16A-CURSOR-046 | mixed | ● active |
| SPR16A-STRUCT-001 | mixed | ● active (amended) |
| SPR16A-RLE-002 | static analysis | ● active (amended by EXP-0356) |
| SPR16A-RLE-003 | mixed | ● active |
| SPR16A-PIX-004 | mixed | ✖ retracted |
| SPR16A-PIX-005 | mixed | ✖ retracted |
| SPR16A-PAL-006 | mixed | ✖ retracted |
| SPR16A-PAL-008 | static analysis | ● active |
| SPR16A-PIX-009 | mixed | ✖ retracted |
| SPR16A-PIX-010 | mixed | ✖ retracted |
| SPR16A-PIX-011 | static analysis | ● active (amended by EXP-0356) |
| SPR16A-FONT-007 | mixed | ● active (amended) |
| SPR16A-TRLR-012 | static analysis | ● active (amended; .16 clauses partially retracted) |
| SPR16A-FONT-013 | static analysis | ● active |
| SPR16A-FONT-014 | static analysis | ● active (partially retracted) |
| SPR16A-FONT-015 | static analysis | ● active |
| SPR16A-RDR-017 | static analysis | ● active (partially retracted) |
| SPR16A-FONT-018 | static analysis | ● active (amended) |
| SPR16A-FONT-019 | static analysis | ● active |
| SPR16A-FONT-020 | static analysis | ● active (amended) |
| SPR16A-FONT-021 | mixed | ● active |
| SPR16A-FONT-022 | mixed | ● active (amended) |
| SPR16A-TXT-023 | static analysis |  |
| SPR16A-BOUND-016 | mixed | ● active |
| SPR16A-PROJ-024 | mixed | ● active |
| SPR16A-ALPHA-025 | static analysis |  |
| SPR16A-PROJ-026 | mixed | ● active |
| SPR16A-ICON-027 | mixed | ● active |
| SPR16A-CAST-028 | mixed | ● active |
| SPR16A-MARK-029 | mixed |  |
| SPR16A-PART-030 | static analysis | ✔ promoted |
| SPR16A-031 | mixed | ✔ promoted |
| SPR16A-CURSOR-061 | static analysis | ● active |
| SPR16A-CURSOR-067 | static analysis | ● active |
| SPR16A-070 | static analysis | ✔ promoted |
| SPR16A-071 | static analysis | ✔ promoted |
| SPR16A-072 | static analysis | ✔ promoted |
| SPR16A-073 | static analysis | ✔ promoted |
| SPR16A-078 | static analysis | ✔ promoted |
| SPR16A-079 | static analysis | ✔ promoted |
| SPR16A-080 | mixed | ✔ promoted |
| SPR16A-081 | static analysis | ✔ promoted |
| SPR16A-082 | static analysis | ✔ promoted |
| SPR16A-083 | mixed | ✔ promoted |
| SPR16A-FONT-094 | static analysis | ● active |

## Ledger spr256

| ID | Method | Status |
|---|---|---|
| SPR256-CURSOR-046 | static analysis | ● active |
| SPR256-STRUCT-001 | mixed | ● active |
| SPR256-COUNT-002 | observation | ● active (amended, superseded) |
| SPR256-PAL-003 | observation | ● active |
| SPR256-VAR-004 | observation | ● active |
| SPR256-EXC-005 | observation | ● active (amended) |
| SPR256-CORPUS-006 | observation | ● active (amended) |
| SPR256-RLE-007 | observation | ● active (amended) |
| SPR256-RLE-008 | observation | ● active (amended) |
| SPR256-RLE-009 | observation | ● active (amended) |
| SPR256-RLE-010 | observation | ● active |
| SPR256-PAL-011 | mixed | ● active |
| SPR256-PAL-012 | static analysis | ● active (amended, superseded) |
| SPR256-PAL-013 | mixed | ● active (amended) |
| SPR256-OVL-014 | static analysis | ● active (amended, superseded) |
| SPR256-OVL-015 | mixed | ● active |
| SPR256-TRLR-016 | mixed | ● active (amended) |
| SPR256-EXC-017 | mixed | ● active |
| SPR256-PAL-018 | observation | ● active |
| SPR256-RLE-019 | mixed | ● active |
| SPR256-RLE-020 | static analysis | ● active (amended) |
| SPR256-RLE-022 | mixed | ● active (amended) |
| SPR256-EXC-020 | static analysis | ● active |
| SPR256-TRLR-021 | mixed | ● active |
| SPR256-FRAME-023 | mixed | ● active |
| SPR256-UNIT-024 | static analysis | ● active |
| SPR256-STR-040 | static analysis | ● active |
| SPR256-STR-041 | static analysis | ● active |
| SPR256-EQUIP-042 | mixed | ● active |
| SPR256-PICT-043 | mixed | ● active |
| SPR256-KEY-044 | static analysis | ● active (amended) |
| SPR256-DOLL-045 | static analysis | ● active |
| SPR256-077 | static analysis | ● active |
| SPR256-061 | static analysis | ✔ promoted |
| SPR256-062 | mixed | ✔ promoted |
| SPR256-063 | static analysis | ✔ promoted |
| SPR256-064 | mixed | ✔ promoted |
| SPR256-065 | mixed | ✔ promoted |

## Ledger tavern

| ID | Method | Status |
|---|---|---|
| MERC-TYPE-001 | static analysis | ● active (amended) |
| MERC-SHELF-002 | static analysis | ● active (partially retracted) |
| MERC-HIRE-003 | static analysis | ● active |
| MERC-PRICE-004 | static analysis | ● active |
| MERC-LEVEL-005 | static analysis | ● active |
| MERC-CMD-007 | static analysis | ● active |
| MERC-DEATH-006 | static analysis | ● active (amended) |
| MERC-POOL-011 | static analysis | ● active |
| MERC-POOL-012 | static analysis | ● active |
| TAVERN-MERCVOICE-008 | mixed | ● active (amended) |
| TAVERN-ORDER-015 | static analysis | ● active |
| TAVERN-TALKPIC-016 | static analysis | ● active |
| TAVERN-TALKSTATS-017 | static analysis | ● active |
| TAVERN-CLICK-019 | static analysis | ● active |
| TAVERN-BUTTON-020 | static analysis | ● active |
| TAVERN-FIGURE-021 | static analysis | ● active |
| TAVERN-LINES-022 | static analysis | ● active |
| TAVERN-023 | static analysis | ✔ promoted |
| TAVERN-024 | static analysis | ✔ promoted |

## Ledger terrain

| ID | Method | Status |
|---|---|---|
| TERR-LOC-001 | mixed | ● active (partially retracted) |
| TERR-LOAD-002 | static analysis | ● active (amended) |
| TERR-IDX-003 | static analysis | ● active (partially retracted) |
| TERR-SEM-004 | static analysis | ● active |
| TERR-VER-005 | mixed | ● active |
| TERR-ANIM-006 | static analysis | ● active |
| TERR-ANIM-007 | static analysis | ● active |
| TERR-ANIM-008 | static analysis | ● active (amended, superseded) |
| TERR-ANIM-009 | static analysis | ● active |
| TERR-ANIM-010 | mixed | ● active |
| TERR-LIGHT-011 | static analysis | ● active (amended, superseded) |
| TERR-LIGHT-012 | static analysis | ● active |
| TERR-LIGHT-013 | static analysis | ● active (amended) |
| TERR-LIGHT-014 | static analysis | ● active (amended, partially retracted) |
| TERR-LIGHT-015 | static analysis | ● active (amended) |
| TERR-LIGHT-016 | mixed | ● active (amended) |
| TERR-DIRT-017 | static analysis | ● active |
| TERR-LIGHT-018 | static analysis | ● active (amended) |
| TERR-LIGHT-019 | static analysis | ● active |
| TERR-LIGHT-020 | mixed | ● active |
| TERR-LIGHT-021 | static analysis | ● active |
| TERR-LIGHT-022 | static analysis | ● active |
| TERR-LIGHT-023 | static analysis | ● active (amended) |
| TERR-EDGE-024 | static analysis | ● active |
| TERR-EDGE-025 | static analysis | ● active |
| TERR-EDGE-026 | static analysis | ● active |
| TERR-GRID-027 | mixed | ● active |
| TERR-LIGHT-028 | static analysis | ● active |
| TERR-LIGHT-029 | static analysis | ● active |
| TERR-LIGHT-030 | static analysis | ● active (amended, partially retracted) |
| TERR-GEOM-031 | static analysis | ● active (amended) |
| TERR-GEOM-032 | static analysis | ● active |
| TERR-GEOM-033 | static analysis | ● active |
| TERR-GEOM-034 | static analysis | ● active (amended) |
| TERR-GEOM-035 | static analysis | ● active |
| TERR-GEOM-036 | static analysis | ● active (amended, superseded) |
| TERR-FOG-037 | static analysis | ● active (amended, superseded) |
| TERR-SPR-038 | static analysis | ● active (amended) |
| TERR-SPR-039 | static analysis | ● active (amended) |
| TERR-SPR-040 | static analysis | ● active (amended) |
| TERR-SPR-041 | static analysis | ● active (amended) |
| TERR-SPR-042 | static analysis | ● active (amended, superseded) |
| TERR-SPR-043 | static analysis | ● active |
| TERR-TILE-044 | static analysis | ● active (amended, partially retracted) |
| TERR-SPR-047 | static analysis | ● active (amended) |
| TERR-SPR-048 | static analysis | ● active (partially retracted) |
| TERR-PASS-049 | static analysis | ● active |
| TERR-PASS-050 | static analysis | ● active (partially retracted) |
| TERR-PASS-051 | static analysis | ● active (amended, superseded) |
| TERR-COST-052 | static analysis | ● active (amended) |
| TERR-PASS-053 | static analysis | ● active |
| TERR-MOVE-054 | static analysis | ● active |
| TERR-MOVE-055 | static analysis | ● active (amended) |
| TERR-MOVE-056 | static analysis | ● active |
| TERR-MOVE-057 | static analysis | ● active (amended, superseded) |
| TERR-MOVE-058 | static analysis | ● active (amended) |
| TERR-LIGHT-059 | static analysis | ● active |
| TERR-LIGHT-060 | static analysis | ● active |
| TERR-LIGHT-061 | static analysis | ● active |
| TERR-LIGHT-062 | static analysis | ● active |
| TERR-LIGHT-063 | static analysis | ● active |
| TERR-LIGHT-064 | static analysis | ● active (amended, partially retracted) |
| TERR-SPR-065 | static analysis | ✔ promoted (partially retracted) |
| TERR-SPR-066 | static analysis | ● active |
| TERR-SPR-067 | static analysis | ● active (amended) |
| TERR-STRUCT-068 | static analysis | ● active (amended, superseded) |
| TERR-STRUCT-069 | static analysis | ● active |
| TERR-STRUCT-070 | static analysis | ● active (amended, partially retracted) |
| TERR-STRUCT-071 | static analysis | ● active (amended) |
| TERR-STRUCT-072 | static analysis | ● active |
| TERR-PASS-073 | static analysis | ● active (amended, partially retracted) |
| TERR-STRUCT-074 | static analysis | ● active (amended) |
| TERR-STRUCT-075 | static analysis | ● active (superseded, contested) |
| TERR-STRUCT-076 | static analysis | ● active |
| TERR-STRUCT-077 | static analysis | ● active (amended, superseded) |
| TERR-STRUCT-078 | static analysis | ● active |
| TERR-TILE-079 | static analysis | ● active (amended, superseded) |
| TERR-FOG-080 | static analysis | ● active (superseded) |
| TERR-FOG-081 | static analysis | ● active |
| TERR-FOG-082 | static analysis | ● active (partially retracted) |
| TERR-FOG-083 | static analysis | ● active |
| TERR-FOG-084 | static analysis | ● active |
| TERR-FOG-085 | static analysis | ● active |
| TERR-FOG-086 | static analysis | ● active (partially retracted) |
| TERR-FOG-087 | static analysis | ● active (partially retracted) |
| TERR-FOG-088 | static analysis | ● active (amended) |
| TERR-FOG-089 | static analysis | ● active |
| TERR-STRUCT-090 | static analysis | ● active |
| TERR-STRUCT-100 | static analysis | ● active |
| TERR-STRUCT-101 | static analysis | ● active |
| TERR-STRUCT-102 | static analysis | ● active |
| TERR-STRUCT-103 | static analysis | ● active |
| TERR-STRUCT-104 | static analysis | ✔ promoted (partially retracted) |
| TERR-STRUCT-105 | static analysis | ● active |
| TERR-STRUCT-106 | static analysis | ● active |
| TERR-STRUCT-107 | static analysis | ● active |
| TERR-LIGHT-108 | static analysis | ● active |
| TERR-LIGHT-109 | static analysis | ● active |
| TERR-LIGHT-110 | static analysis | ● active |
| TERR-LIGHT-111 | static analysis | ● active |
| TERR-LIGHT-112 | static analysis | ● active |
| TERR-LIGHT-113 | static analysis | ● active |
| TERR-LIGHT-114 | static analysis | ● active |
| TERR-SIGHT-115 | static analysis | ● active |
| TERR-SIGHT-116 | static analysis | ● active |
| TERR-FOG-117 | static analysis | ● active |
| TERR-FOG-118 | static analysis | ● active |
| TERR-LIGHT-119 | static analysis | ● active |
| TERR-LIGHT-120 | static analysis | ● active |
| TERR-LIGHT-121 | static analysis | ● active |
| TERR-LIGHT-122 | static analysis | ● active |
| TERR-LIGHT-123 | static analysis | ● active |
| TERR-LIGHT-124 | static analysis | ● active |
| TERR-LIGHT-125 | mixed | ● active |
| TERR-LIGHT-126 | static analysis | ● active (amended) |
| TERR-LIGHT-127 | static analysis | ● active |
| TERR-LIGHT-128 | mixed | ● active |
| TERR-SHDW-129 | static analysis | ● active |
| TERR-SHDW-130 | static analysis | ● active |
| TERR-SHDW-131 | static analysis | ● active |
| TERR-SHDW-132 | static analysis | ● active |
| TERR-SHDW-133 | static analysis | ● active |
| TERR-SHDW-134 | static analysis | ● active |
| TERR-SHDW-135 | static analysis | ● active |
| TERR-SHDW-136 | static analysis | ● active |
| TERR-SPR-137 | static analysis | ✔ promoted (partially retracted) |
| TERR-SPR-138 | static analysis | ✔ promoted (partially retracted) |
| TERR-SPR-139 | static analysis | ✔ promoted (partially retracted) |
| TERR-SPR-140 | static analysis | ● active |
| TERR-SPR-141 | static analysis | ● active |
| TERR-FOG-142 | static analysis | ● active |
| TERR-LIGHT-143 | static analysis | ● active |
| TERR-SPR-144 | static analysis | ● active |
| TERR-FOG-145 | static analysis | ● active |
| TERR-CELLREC-146 | static analysis | ● active |
| TERR-FOOTPRINT-147 | static analysis | ● active (amended, partially retracted) |
| TERR-PASS-148 | static analysis | ● active |
| TERR-LIGHT-149 | static analysis | ● active |
| TERR-LIGHT-150 | static analysis | ● active |
| TERR-LIGHT-151 | static analysis | ● active |
| TERR-LOAD-152 | static analysis | ● active |
| TERR-STREAM-157 | static analysis | ● active |
| TERR-TILECONTROL-163 | static analysis | ✔ promoted |
| TERR-WATERBOUND-164 | static analysis | ✔ promoted |
| TERR-GFXBOUND-165 | static analysis | ✔ promoted |
| TERR-SLOTADMIT-166 | static analysis | ✔ promoted |
| TERR-WATERPASS-167 | static analysis | ✔ promoted |
| TERR-RLEBIT-168 | static analysis | ✔ promoted |
| TERR-TILECLEAR-169 | static analysis | ✔ promoted |
| TERR-DRAWSTAMP-170 | static analysis | ✔ promoted |
| TERR-DRAWGATE-171 | static analysis | ✔ promoted |
| TERR-METALIFE-175 | static analysis | ✔ promoted |
| TERR-FAMILY-187 | static analysis | ✔ promoted |
| TERR-FAMILY-188 | mixed | ✔ promoted |
| TERR-191 | static analysis | ● active |
| TERR-192 | static analysis | ● active (amended) |
| TERR-193 | static analysis | ● active (amended) |
| TERR-194 | static analysis | ● active |
| TERR-195 | static analysis | ● active |
| TERR-196 | static analysis | ● active |
| TERR-198 | static analysis | ● active |
| TERR-200 | static analysis | ● active |
| TERR-202 | static analysis | ● active |
| TERR-204 | static analysis | ● active |
| TERR-GMAP-206 | mixed | ● active |
| TERR-STRUCT-207 | mixed | ● active |
| TERR-PLACE-208 | static analysis | ● active |
| TERR-STRUCT-210 | static analysis | ● active |
| TERR-STRUCT-211 | mixed | ● active |
| TERR-STRUCT-212 | mixed | ● active |

## Ledger text

| ID | Method | Status |
|---|---|---|
| TEXT-CONV-001 | static analysis | ● active (partially retracted) |
| TEXT-LANG-002 | static analysis | ● active |
| TEXT-INDEX-003 | static analysis | ● active |
| TEXT-FIT-004 | mixed | ● active (superseded) |
| TEXT-IN-005 | static analysis | ● active |
| TEXT-LOWER-006 | static analysis | ● active |
| TEXT-DOM-010 | static analysis | ● active |
| TEXT-ALIAS-011 | static analysis | ● active |
| TEXT-SEL0-012 | static analysis | ● active |
| TEXT-FIT2-013 | mixed | ● active |
| TEXT-CAP-018 | mixed | ● active |
| TEXT-API-007 | static analysis | ● active |
| TEXT-FONT3-008 | static analysis | ● active |
| TEXT-TILDE-009 | static analysis | ● active (amended, partially retracted, superseded) |
| TEXT-FONT-015 | static analysis | ● active |
| TEXT-FONT2-016 | mixed | ● active |
| TEXT-TILDE2-017 | static analysis | ● active |
| TEXT-079 | static analysis | ● active |
| TEXT-065 | static analysis | ● active |
| TEXT-066 | static analysis | ● active |
| TEXT-067 | static analysis | ● active |
| TEXT-068 | static analysis | ● active |
| TEXT-069 | mixed | ● active |
| TEXT-070 | static analysis | ● active |
| TEXT-071 | static analysis | ● active |
| TEXT-078 | static analysis | ● active |
| TEXT-ROOT-014 | mixed | ● active |
| TEXT-ITEMNAME-019 | mixed | ● active |
| TEXT-PATTERN-020 | static analysis | ● active |
| TEXT-BATTLEROOT-063 | mixed | ● active |
| TEXT-STRTAB-023 | static analysis | ● active (partially retracted) |
| TEXT-NAMETAB-026 | mixed | ● active (amended) |
| TEXT-STRTAB-030 | static analysis | ● active |
| TEXT-STRTAB-031 | static analysis | ● active |
| TEXT-UI-032 | mixed | ● active |
| TEXT-UI-033 | static analysis | ● active |
| TEXT-UI-034 | static analysis | ● active |
| TEXT-UI-035 | static analysis | ● active |
| TEXT-UI-036 | mixed | ● active (amended) |
| TEXT-UI-037 | static analysis | ● active |
| TEXT-UI-038 | static analysis | ● active |
| TEXT-UI-039 | static analysis | ● active |
| TEXT-UI-040 | static analysis | ● active |
| TEXT-UI-041 | static analysis | ● active |
| TEXT-UI-042 | static analysis | ● active |
| TEXT-UI-043 | mixed | ● active |
| TEXT-UI-044 | static analysis | ● active |
| TEXT-UI-045 | static analysis | ● active |
| TEXT-UI-046 | mixed | ● active |
| TEXT-UI-047 | mixed | ● active |
| TEXT-HOVER-048 | static analysis | ● active |
| TEXT-HOVERSET-049 | static analysis | ● active |
| TEXT-HOVERCHAR-050 | static analysis | ● active |
| TEXT-HOVERROOM-051 | static analysis | ● active |
| TEXT-HOVERTEXT-052 | static analysis | ● active |
| TEXT-HOVERPAINT-053 | static analysis | ● active |
| TEXT-080 | static analysis | ● active |
| TEXT-081 | static analysis | ● active (amended) |
| TEXT-082 | static analysis | ● active (amended) |
| TEXT-083 | static analysis | ● active |
| TEXT-084 | static analysis | ● active (amended) |
| TEXT-085 | static analysis | ● active (amended) |
| TEXT-096 | static analysis | ● active |
| TEXT-097 | static analysis | ● active |
| TEXT-098 | static analysis | ● active |
| TEXT-105 | static analysis | ● active |
| TEXT-106 | static analysis | ● active |
| TEXT-CHARGEN-027 | static analysis | ● active |
| TEXT-CHARGEN-028 | static analysis | ● active |
| TEXT-CHARGEN-029 | static analysis | ● active |
| TEXT-NAMEIN-024 | static analysis | ● active (partially retracted) |
| TEXT-COLL-025 | mixed | ● active |
| TEXT-073 | static analysis | ● active |
| TEXT-074 | static analysis | ● active |
| TEXT-075 | static analysis | ● active |
| TEXT-076 | static analysis | ● active |
| TEXT-077 | static analysis | ● active |
| TEXT-SAVELABEL-054 | static analysis | ● active |
| TEXT-SAVELABEL-055 | static analysis | ✖ retracted |
| TEXT-SAVELABEL-057 | static analysis | ● active |
| TEXT-SAVELABEL-058 | static analysis | ● active |
| TEXT-SAVELABEL-059 | static analysis | ● active (partially retracted) |
| TEXT-SAVELABEL-060 | static analysis | ● active |
| TEXT-SAVELABEL-061 | mixed | ● active (amended) |
| TEXT-086 | static analysis | ● active |
| TEXT-087 | mixed | ● active |
| TEXT-088 | static analysis | ● active (amended) |
| TEXT-089 | mixed | ● active |
| TEXT-090 | static analysis | ● active |
| TEXT-091 | static analysis | ● active |
| TEXT-092 | static analysis | ● active |
| TEXT-094 | static analysis | ● active |
| TEXT-095 | static analysis | ● active |
| TEXT-099 | static analysis | ● active |
| TEXT-100 | static analysis | ● active |
| TEXT-108 | static analysis | ● active |
| TEXT-109 | mixed | ● active |
| TEXT-110 | mixed | ● active |

## Ledger town

| ID | Method | Status |
|---|---|---|
| TOWN-467 | static analysis | ● active |
| TOWN-468 | static analysis | ● active |
| TOWN-469 | static analysis | ● active |
| TOWN-470 | mixed | ● active |
| TOWN-439 | static analysis | ✔ promoted |
| TOWN-440 | static analysis | ✔ promoted |
| TOWN-441 | static analysis | ✔ promoted |
| TOWN-442 | static analysis | ✔ promoted |
| TOWN-443 | static analysis | ✔ promoted |
| TOWN-444 | static analysis | ✔ promoted |
| TOWN-445 | static analysis | ✔ promoted |
| TOWN-446 | static analysis | ✔ promoted |
| TOWN-447 | static analysis | ✔ promoted |
| TOWN-427 | static analysis | ✔ promoted |
| TOWN-428 | static analysis | ✔ promoted |
| TOWN-429 | mixed | ✔ promoted |
| TOWN-430 | mixed | ✔ promoted |
| TOWN-431 | static analysis | ✔ promoted |
| TOWN-432 | static analysis | ✔ promoted |
| TOWN-433 | static analysis | ✔ promoted |
| TOWN-434 | static analysis | ✔ promoted |
| TOWN-415 | static analysis | ✔ promoted |
| TOWN-416 | static analysis | ✔ promoted |
| TOWN-417 | static analysis | ✔ promoted |
| TOWN-418 | static analysis | ✔ promoted |
| TOWN-419 | static analysis | ✔ promoted |
| TOWN-420 | static analysis | ✔ promoted |
| TOWN-407 | static analysis | ✔ promoted |
| TOWN-408 | static analysis | ✔ promoted |
| TOWN-409 | static analysis | ✔ promoted |
| TOWN-410 | static analysis | ✔ promoted |
| TOWN-411 | static analysis | ✔ promoted |
| TOWN-412 | static analysis | ✔ promoted |
| TOWN-413 | static analysis | ✔ promoted |
| TOWN-414 | static analysis | ✔ promoted |
| TOWN-399 | static analysis | ✔ promoted |
| TOWN-400 | static analysis | ✔ promoted |
| TOWN-401 | static analysis | ✔ promoted |
| TOWN-402 | static analysis | ✔ promoted |
| TOWN-403 | static analysis | ✔ promoted |
| TOWN-404 | static analysis | ✔ promoted |
| TOWN-405 | static analysis | ✔ promoted |
| TOWN-406 | static analysis | ✔ promoted |
| TOWN-001 | static analysis | ● active |
| TOWN-002 | static analysis | ● active |
| TOWN-003 | static analysis | ● active |
| TOWN-004 | static analysis | ● active (amended) |
| TOWN-005 | static analysis | ● active |
| TOWN-006 | static analysis | ● active |
| TOWN-007 | mixed | ● active |
| TOWN-008 | mixed | ● active |
| TOWN-009 | static analysis | ● active |
| TOWN-010 | static analysis | ● active (partially retracted, superseded) |
| TOWN-011 | static analysis | ● active |
| TOWN-012 | mixed | ● active (partially retracted) |
| TOWN-013 | static analysis | ● active |
| TOWN-014 | static analysis | ● active |
| TOWN-015 | static analysis | ● active (amended, partially retracted) |
| TOWN-016 | static analysis | ● active |
| TOWN-017 | static analysis | ● active |
| TOWN-018 | static analysis | ● active (amended) |
| TOWN-019 | static analysis | ✖ retracted |
| TOWN-020 | static analysis | ● active (amended, partially retracted) |
| TOWN-021 | static analysis | ● active |
| TOWN-022 | static analysis | ● active (partially retracted) |
| TOWN-023 | static analysis | ● active |
| TOWN-024 | static analysis | ● active |
| TOWN-025 | static analysis | ● active |
| TOWN-026 | mixed | ● active |
| TOWN-036 | static analysis | ● active (amended) |
| TOWN-037 | static analysis | ● active |
| TOWN-038 | static analysis | ● active (amended) |
| TOWN-039 | static analysis | ● active |
| TOWN-040 | static analysis | ● active (amended, partially retracted, superseded) |
| TOWN-041 | static analysis | ● active |
| TOWN-042 | static analysis | ● active |
| TOWN-043 | static analysis | ● active (amended, partially retracted, superseded) |
| TOWN-044 | static analysis | ● active |
| TOWN-045 | static analysis | ● active (amended) |
| TOWN-061 | static analysis | ● active (amended, partially retracted, superseded) |
| TOWN-062 | static analysis | ● active |
| TOWN-063 | static analysis | ● active |
| TOWN-064 | static analysis | ● active |
| TOWN-065 | static analysis | ✖ retracted |
| TOWN-066 | static analysis | ● active (partially retracted) |
| TOWN-067 | static analysis | ● active |
| TOWN-068 | static analysis | ● active (partially retracted) |
| TOWN-086 | static analysis | ● active |
| TOWN-087 | static analysis | ● active |
| TOWN-088 | static analysis | ● active |
| TOWN-089 | static analysis | ● active |
| TOWN-090 | static analysis | ● active |
| TOWN-091 | static analysis | ● active (amended, partially retracted) |
| TOWN-092 | static analysis | ● active |
| TOWN-093 | static analysis | ● active (amended, partially retracted) |
| TOWN-094 | static analysis | ● active |
| TOWN-095 | static analysis | ● active |
| TOWN-096 | static analysis | ● active |
| TOWN-116 | static analysis | ● active |
| TOWN-117 | static analysis | ● active |
| TOWN-118 | static analysis | ● active |
| TOWN-119 | static analysis | ● active (amended, superseded) |
| TOWN-120 | static analysis | ● active |
| TOWN-121 | static analysis | ● active |
| TOWN-122 | static analysis | ● active |
| TOWN-123 | static analysis | ● active |
| TOWN-124 | static analysis | ● active |
| TOWN-GENERAL-106 | static analysis | ● active |
| TOWN-GENERAL-107 | static analysis | ● active (amended) |
| TOWN-GENERAL-108 | static analysis | ● active (amended) |
| TOWN-136 | static analysis | ● active |
| TOWN-137 | static analysis | ● active |
| TOWN-138 | static analysis | ● active |
| TOWN-139 | static analysis | ● active |
| TOWN-140 | static analysis | ● active |
| TOWN-141 | static analysis | ● active |
| TOWN-142 | static analysis | ● active |
| TOWN-143 | static analysis | ● active |
| TOWN-144 | static analysis | ● active |
| TOWN-145 | mixed | ● active |
| TOWN-146 | static analysis | ● active |
| TOWN-147 | static analysis | ● active |
| TOWN-148 | static analysis | ● active |
| TOWN-149 | static analysis | ● active |
| TOWN-150 | static analysis | ● active |
| TOWN-151 | static analysis | ● active |
| TOWN-152 | static analysis | ● active (amended) |
| TOWN-153 | static analysis | ● active |
| TOWN-154 | static analysis | ● active |
| TOWN-155 | static analysis | ● active |
| TOWN-156 | static analysis | ● active |
| TOWN-157 | static analysis | ● active |
| TOWN-158 | static analysis | ● active (partially retracted) |
| TOWN-159 | static analysis | ● active (partially retracted, superseded) |
| TOWN-160 | static analysis | ● active |
| TOWN-161 | static analysis | ● active (partially retracted) |
| TOWN-162 | static analysis | ● active |
| TOWN-163 | static analysis | ● active |
| TOWN-164 | static analysis | ● active |
| TOWN-165 | static analysis | ● active |
| TOWN-181 | static analysis | ● active |
| TOWN-182 | static analysis | ● active |
| TOWN-183 | static analysis | ● active (amended, partially retracted) |
| TOWN-184 | static analysis | ● active (amended, partially retracted) |
| TOWN-185 | static analysis | ● active (amended) |
| TOWN-186 | static analysis | ● active (amended) |
| TOWN-187 | static analysis | ● active |
| TOWN-188 | static analysis | ● active |
| TOWN-206 | static analysis | ● active |
| TOWN-207 | static analysis | ● active |
| TOWN-208 | static analysis | ● active |
| TOWN-209 | static analysis | ● active (amended, superseded) |
| TOWN-210 | static analysis | ● active |
| TOWN-211 | static analysis | ● active (partially retracted) |
| TOWN-214 | static analysis | ● active (partially retracted) |
| TOWN-215 | static analysis | ● active (partially retracted) |
| TOWN-216 | static analysis | ● active (amended) |
| TOWN-217 | static analysis | ● active (amended) |
| TOWN-222 | static analysis | ● active |
| TOWN-223 | static analysis | ● active (amended) |
| TOWN-224 | static analysis | ● active (amended, partially retracted) |
| TOWN-232 | static analysis | ● active (amended, partially retracted) |
| TOWN-233 | static analysis | ● active (amended) |
| TOWN-234 | static analysis | ● active |
| TOWN-235 | static analysis | ● active |
| TOWN-236 | static analysis | ● active (partially retracted) |
| TOWN-242 | static analysis | ● active |
| TOWN-243 | static analysis | ● active |
| TOWN-244 | static analysis | ● active (partially retracted) |
| TOWN-245 | static analysis | ● active |
| TOWN-246 | static analysis | ● active |
| TOWN-252 | static analysis | ● active |
| TOWN-253 | static analysis | ● active |
| TOWN-254 | static analysis | ● active |
| TOWN-255 | static analysis | ● active |
| TOWN-256 | static analysis | ● active |
| TOWN-257 | static analysis | ● active |
| TOWN-258 | static analysis | ● active |
| TOWN-259 | static analysis | ● active (amended, partially retracted) |
| TOWN-260 | static analysis | ● active |
| TOWN-261 | static analysis | ● active |
| TOWN-264 | mixed | ● active |
| TOWN-265 | mixed | ● active |
| TOWN-280 | static analysis | ● active |
| TOWN-281 | static analysis | ● active |
| TOWN-282 | static analysis | ● active (partially retracted) |
| TOWN-283 | static analysis | ● active (amended) |
| TOWN-284 | static analysis | ● active |
| TOWN-266 | static analysis | ● active |
| TOWN-312 | static analysis | ● active |
| TOWN-313 | static analysis | ● active |
| TOWN-314 | static analysis | ● active |
| TOWN-328 | static analysis | ● active |
| TOWN-329 | static analysis | ● active |
| TOWN-330 | static analysis | ● active |
| TOWN-331 | static analysis | ● active |
| TOWN-332 | static analysis | ● active |
| TOWN-333 | static analysis | ● active |
| TOWN-344 | static analysis | ● active |
| TOWN-345 | static analysis | ● active |
| TOWN-346 | static analysis | ● active |
| TOWN-347 | static analysis | ● active (amended, partially retracted) |
| TOWN-348 | static analysis | ● active |
| TOWN-349 | static analysis | ● active |
| TOWN-350 | static analysis | ● active |
| TOWN-351 | static analysis | ● active (amended, partially retracted) |
| TOWN-352 | static analysis | ● active |
| TOWN-353 | static analysis | ● active |
| TOWN-354 | static analysis | ● active |
| TOWN-355 | static analysis | ● active |
| TOWN-356 | static analysis | ● active |
| TOWN-357 | static analysis | ● active |
| TOWN-358 | static analysis | ● active |
| TOWN-359 | mixed | ● active |
| TOWN-360 | static analysis | ● active |
| TOWN-372 | static analysis | ● active |
| TOWN-373 | static analysis | ● active |
| TOWN-379 | static analysis | ● active |
| TOWN-380 | static analysis | ● active |
| TOWN-381 | static analysis | ● active |
| TOWN-382 | static analysis | ● active |
| TOWN-383 | static analysis | ● active |
| TOWN-384 | static analysis | ● active |
| TOWN-385 | static analysis | ● active |
| TOWN-391 | static analysis | ● active |
| TOWN-392 | static analysis | ● active |
| TOWN-393 | static analysis | ● active |
| TOWN-453 | static analysis | ● active |
| TOWN-OPTIONS-457 | mixed | ✔ promoted |
| TOWN-AUTOHEAL-458 | mixed | ✔ promoted |
| TOWN-GRAPHICS-459 | mixed | ✔ promoted |
| TOWN-SMOOTH-460 | static analysis | ✔ promoted |
| TOWN-473 | static analysis | ● active |
| TOWN-474 | static analysis | ● active |
| TOWN-475 | static analysis | ● active |
| TOWN-476 | static analysis | ● active |
| TOWN-477 | static analysis | ● active |
| TOWN-478 | static analysis | ● active |
| TOWN-479 | static analysis | ● active |
| TOWN-480 | static analysis | ● active |
| TOWN-481 | static analysis | ● active |
| TOWN-482 | static analysis | ● active |
| TOWN-483 | static analysis | ● active |
| TOWN-484 | static analysis | ● active |
| TOWN-485 | static analysis | ● active (branch candidate) |
| TOWN-486 | static analysis | ● active (branch candidate) |
| TOWN-487 | static analysis | ● active (branch candidate) |
| TOWN-489 | static analysis | ● active (branch candidate) |
| TOWN-490 | mixed | ● active (branch candidate) |
| TOWN-491 | static analysis | ● active (branch candidate) |
| TOWN-492 | static analysis | ● active (branch candidate) |
| TOWN-493 | static analysis | ● active (branch candidate) |
| TOWN-494 | static analysis | ● active (branch candidate) |
| TOWN-495 | static analysis | ● active (branch candidate) |
| TOWN-496 | mixed | ● active (branch candidate) |
| TOWN-497 | static analysis | ● active (branch candidate) |
| TOWN-503 | static analysis | ● active (branch candidate) |
| TOWN-504 | static analysis | ● active (branch candidate) |
| TOWN-505 | static analysis | ● active (branch candidate) |
| TOWN-506 | static analysis | ● active (branch candidate) |
| TOWN-507 | static analysis | ● active (branch candidate) |
| TOWN-499 | static analysis | ● active (branch candidate) |
| TOWN-500 | static analysis | ● active (branch candidate) |
| TOWN-501 | static analysis | ● active (branch candidate) |
| TOWN-502 | static analysis | ● active (branch candidate) |
| TOWN-MARKER-508 | static analysis | ● active |

## Ledger trigger

| ID | Method | Status |
|---|---|---|
| TRIG-EVAL-001 | static analysis | ● active (contested) |
| TRIG-STORE-002 | static analysis | ● active |
| TRIG-COND-003 | static analysis | ● active |
| TRIG-ACT-004 | static analysis | ● active |
| TRIG-GROUP-005 | static analysis | ● active |
| TRIG-CMP-006 | static analysis | ● active |
| TRIG-FIRE-007 | static analysis | ● active |
| TRIG-SAVE-008 | static analysis | ● active (partially retracted, amended) |
| TRIG-END-009 | static analysis | ● active (partially retracted, amended, superseded) |
| TRIG-BIND-010 | static analysis | ● active (amended) |
| TRIG-REC-011 | static analysis | ● active |
| TRIG-INI-012 | static analysis | ● active |
| TRIG-DROP-013 | static analysis | ● active (amended) |
| TRIG-DIST-014 | static analysis | ● active |
| TRIG-COUNT-015 | static analysis | ● active |
| TRIG-GRPLIST-016 | static analysis | ● active |
| TRIG-REAP-017 | static analysis | ● active (amended) |
| TRIG-EMPTY-018 | mixed | ● active |
| TRIG-DIPLO-019 | static analysis | ● active |
| TRIG-DIPLO-020 | static analysis | ● active |
| TRIG-DIPLO-021 | mixed | ● active |
| TRIG-SACK-022 | static analysis | ● active |
| TRIG-MSG-023 | static analysis | ● active |
| TRIG-DROPALL-024 | static analysis | ● active (amended) |
| TRIG-GIVEALL-025 | static analysis | ● active |
| TRIG-CAT-026 | mixed | ● active |
| TRIG-ADDITEM-027 | static analysis | ● active (amended, superseded) |
| TRIG-MONEY-028 | static analysis | ● active |
| TRIG-NEAREST-029 | static analysis | ● active |
| TRIG-PARAM-030 | static analysis | ● active |
| TRIG-ESCORT-031 | static analysis | ● active |
| TRIG-TARGETID-032 | mixed | ● active |
| TRIG-CAST-033 | static analysis | ● active |
| TRIG-EFFECTTIME-034 | mixed | ● active (amended, superseded) |
| TRIG-CELLTAIL-035 | static analysis | ● active (amended, partially retracted) |
| TRIG-PROPERTY-036 | static analysis | ● active |
| TRIG-CLOSURE-037 | mixed | ✔ promoted (amended, superseded) |
| TRIG-TAKEITEM-038 | static analysis | ● active |
| TRIG-XFERITEM-039 | static analysis | ● active |
| TRIG-ITEMTEST-040 | static analysis | ● active |
| TRIG-OFFMAP-041 | static analysis | ● active |
| TRIG-RETURN-042 | static analysis | ● active |
| TRIG-MAPGROUP-043 | static analysis | ● active |
| TRIG-CASTACTOR-044 | static analysis | ● active (amended, partially retracted) |
| TRIG-CELLEFFECT-045 | static analysis | ● active |
| TRIG-INSTCENSUS-046 | mixed | ● active |
| TRIG-GRPARM-047 | static analysis | ● active |
| TRIG-GRPLIMIT-048 | static analysis | ● active |
| TRIG-MSGCORPUS-049 | mixed | ● active |
| TRIG-MSGLIMIT-050 | static analysis | ● active |
| TRIG-CHECK-051 | static analysis | ● active |
| TRIG-CHECK-052 | static analysis | ● active |
| TRIG-CHECK-053 | static analysis | ● active |
| TRIG-CHECK-054 | static analysis | ● active |
| TRIG-BIND-063 | static analysis | ✔ promoted |
| TRIG-HEROORD-075 | static analysis | ● active |
| TRIG-HEROTPL-076 | static analysis | ● active |
| TRIG-HEROBIND-077 | static analysis | ● active |
| TRIG-HEROFAIL-078 | static analysis | ● active |
| TRIG-DRAGON-079 | mixed | ● active |
| TRIG-BRIGAND-080 | mixed | ● active |
| TRIG-VIP-081 | mixed | ● active |

## Ledger unit

| ID | Method | Status |
|---|---|---|
| UNIT-T9CTOR-110 | static analysis | ✔ promoted |
| UNIT-T9LIFE-111 | static analysis | ✔ promoted |
| UNIT-STRUCTUSE-090 | static analysis | ✔ promoted |
| UNIT-STRUCTUSE-091 | static analysis | ✔ promoted |
| UNIT-STRUCTCONT-078 | static analysis | ✔ promoted |
| UNIT-STRUCTNEXT-079 | static analysis | ✔ promoted |
| UNIT-STRUCTZERO-080 | static analysis | ✔ promoted |
| UNIT-STRUCTBOUND-081 | static analysis | ✔ promoted |
| UNIT-STRUCTCELL-070 | static analysis | ● active |
| UNIT-AREAVISIT-071 | static analysis | ● active |
| UNIT-AREADIRECT-072 | static analysis | ● active |
| UNIT-AREAHP-073 | static analysis | ● active |
| UNIT-STRUCTDETACH-074 | static analysis | ● active |
| UNIT-AREAPOP-075 | mixed | ● active |
| UNIT-STRUCTORDER-062 | static analysis | ● active |
| UNIT-STRUCTREACH-063 | static analysis | ● active |
| UNIT-STRUCTDAMAGE-064 | static analysis | ● active |
| UNIT-STRUCTDELIVERY-065 | static analysis | ● active |
| UNIT-STRUCTSTOP-066 | static analysis | ● active |
| UNIT-M10CELL-054 | static analysis | ● active |
| UNIT-M10ENTRY-055 | static analysis | ● active |
| UNIT-M10CAST-056 | static analysis | ● active |
| UNIT-M10LIFE-057 | static analysis | ● active |
| UNIT-STREAM-001 | static analysis | ● active (contested) |
| UNIT-DIFF-002 | static analysis | ● active |
| UNIT-DERIVE-003 | static analysis | ● active |
| UNIT-CTOR-004 | static analysis | ● active |
| UNIT-EQUIP-005 | static analysis | ● active |
| UNIT-COMBAT-006 | static analysis | ● active (partially retracted) |
| UNIT-SPELL-007 | static analysis | ● active |
| UNIT-PROTO-008 | static analysis | ● active |
| UNIT-OWNER-009 | static analysis | ● active (partially retracted, superseded) |
| UNIT-PANEL-010 | static analysis | ● active (superseded) |
| UNIT-PANEL-011 | static analysis | ● active |
| UNIT-GATE-012 | static analysis | ● active |
| UNIT-GATE-013 | static analysis | ● active |
| UNIT-GATE-014 | static analysis | ● active (amended) |
| UNIT-COMBAT-015 | static analysis | ● active |
| UNIT-WIDTH-016 | static analysis | ● active |
| UNIT-ROOT-017 | mixed | ● active |
| UNIT-M10-018 | mixed | ● active |
| UNIT-ABSORB-019 | mixed | ● active |
| UNIT-HOVER-020 | static analysis | ● active (partially retracted) |
| UNIT-VPLAYER-021 | static analysis | ● active (superseded) |
| UNIT-VPLAYER-022 | static analysis | ● active |
| UNIT-CLASS-023 | static analysis | ● active |
| UNIT-RTTI-024 | mixed | ● active |
| UNIT-APPEAR-030 | static analysis | ● active |
| UNIT-APPEAR-031 | static analysis | ● active |
| UNIT-FIGURE-032 | static analysis | ● active (amended) |
| UNIT-GATE-033 | static analysis | ● active |
| UNIT-PLACE-034 | static analysis | ● active |
| UNIT-PICT-035 | static analysis | ● active |
| UNIT-PICT-036 | static analysis | ● active |
| UNIT-PICT-037 | static analysis | ● active |
| UNIT-PICT-038 | static analysis | ● active |
| UNIT-NAME-039 | static analysis | ● active |
| UNIT-NAMETAB-041 | mixed | ● active |
| UNIT-INSTNAME-040 | static analysis | ● active |
| UNIT-SPRITE-042 | static analysis | ● active |
| UNIT-VISBIT-044 | static analysis | ● active |
| UNIT-PLACESKILL-086 | static analysis | ● active |
| UNIT-PLACERESIST-087 | static analysis | ● active |
| UNIT-PLACEIDLE-088 | static analysis | ● active |
| UNIT-PLACEGEAR-098 | static analysis | ✔ promoted |
| UNIT-PLACEFRONTIER-099 | static analysis | ✔ promoted |
| UNIT-REGEN-130 | static analysis | ● active |
| UNIT-BIND-131 | static analysis | ● active |
| UNIT-STRUCTUSE-137 | static analysis | ● active |
| UNIT-STRUCTHIT-138 | static analysis | ● active |
| UNIT-140 | static analysis | ● active |
| UNIT-141 | static analysis | ● active |
| UNIT-142 | mixed | ● active |
| UNIT-143 | mixed | ● active |
| UNIT-144 | static analysis | ● active |
| UNIT-145 | static analysis | ● active |
| UNIT-146 | static analysis | ● active |
| UNIT-147 | static analysis | ● active (amended) |
| UNIT-148 | static analysis | ● active |

## Ledger video

| ID | Method | Status |
|---|---|---|
| VIDEO-MUSIC-001 | observation | ✔ promoted |
| VIDEO-MUSIC-002 | observation | ✔ promoted |
| VIDEO-MUSIC-003 | observation | ✔ promoted |
| VIDEO-MUSIC-004 | observation | ✔ promoted |
| VIDEO-MUSIC-005 | observation | ✔ promoted |
| VIDEO-MUSIC-006 | observation | ✔ promoted |
| VIDEO-MUSIC-007 | observation | ✔ promoted |
| VIDEO-MUSIC-008 | observation | ✔ promoted |
| VIDEO-MUSIC-009 | observation | ✔ promoted |
| VIDEO-MUSIC-010 | observation | ✔ promoted |
| VIDEO-MUSIC-011 | observation | ✔ promoted |
| VIDEO-MUSIC-012 | observation | ✔ promoted |
| VIDEO-SFX-013 | mixed | ✔ promoted |
| VIDEO-SFX-014 | mixed | ✔ promoted |
| VIDEO-SFX-015 | mixed | ✔ promoted |
| VIDEO-SFX-016 | mixed | ✔ promoted |
| VIDEO-SFX-017 | mixed | ✔ promoted |
| VIDEO-SFX-018 | mixed | ✔ promoted |
| VIDEO-SFX-019 | mixed | ✔ promoted |
| VIDEO-SFX-020 | mixed | ✔ promoted |
| VIDEO-SFX-021 | mixed | ✔ promoted |
| VIDEO-029 | observation | ✔ promoted |
| VIDEO-030 | observation | ✔ promoted |
| VIDEO-031 | observation | ✔ promoted (amended, partially retracted) |
| VIDEO-032 | observation | ✔ promoted |
| VIDEO-033 | observation | ✔ promoted |
| VIDEO-034 | observation | ✔ promoted |
| VIDEO-035 | observation | ✔ promoted |
| VIDEO-036 | observation | ✔ promoted |
| VIDEO-045 | observation | ✔ promoted (amended, partially retracted) |
| VIDEO-046 | observation | ✔ promoted |
| VIDEO-047 | observation | ✔ promoted |
| VIDEO-048 | observation | ✔ promoted |
| VIDEO-049 | observation | ✔ promoted |
| VIDEO-050 | observation | ✔ promoted |
| VIDEO-051 | observation | ✔ promoted |
| VIDEO-SFX-053 | observation | ✔ promoted |
| VIDEO-SFX-054 | observation | ✔ promoted |
| VIDEO-MUSIC-056 | mixed | ✔ promoted |
| VIDEO-OPTIONS-057 | mixed | ✔ promoted |
| VIDEO-SFX-058 | observation | ✔ promoted |
| VIDEO-SFX-059 | observation | ✔ promoted |
| VIDEO-SFX-060 | observation | ✔ promoted |
| VIDEO-MUSIC-061 | static analysis | ● active |
| VIDEO-MUSIC-062 | static analysis | ● active |
| VIDEO-MUSIC-063 | static analysis | ● active (amended) |
| VIDEO-MUSIC-064 | static analysis | ● active |
| VIDEO-MUSIC-065 | static analysis | ● active |
| VIDEO-MUSIC-066 | static analysis | ● active |
| VIDEO-067 | static analysis | ● active |
| VIDEO-068 | static analysis | ● active |
| VIDEO-069 | static analysis | ● active |
| VIDEO-070 | static analysis | ● active |
| VIDEO-071 | static analysis | ● active (amended, partially retracted) |
| VIDEO-072 | static analysis | ● active (amended) |
| VIDEO-073 | static analysis | ● active (amended) |
| VIDEO-074 | static analysis | ● active |
| VIDEO-081 | static analysis | ● active |
| VIDEO-082 | static analysis | ● active |
| VIDEO-083 | static analysis | ● active |
| VIDEO-084 | static analysis | ● active |
| VIDEO-075 | static analysis | ● active |
| VIDEO-076 | static analysis | ● active |
| VIDEO-077 | static analysis | ● active |
| VIDEO-078 | static analysis | ● active |
| VIDEO-079 | static analysis | ● active |
| VIDEO-SFX-085 | static analysis | ✔ promoted |
| VIDEO-SFX-086 | static analysis | ✔ promoted |
| VIDEO-SFX-087 | static analysis | ✔ promoted |
| VIDEO-SFX-088 | static analysis | ● active |
| VIDEO-SFX-089 | static analysis | ● active |
| VIDEO-SFX-090 | static analysis | ● active |
| VIDEO-SFX-091 | static analysis | ● active |
