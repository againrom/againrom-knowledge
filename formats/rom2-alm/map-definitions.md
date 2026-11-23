# ROM2 map definition tables

Companion of the [ROM2 ALM map page](format.md). The tables give, for the installed maps (EN 59, RU 83: the `.alm`
files of the root folder and the `.alm` members of `scenario.res`), which Data.bin row each type-6 definition id,
each type-4 kind and each sack item names. The Data.bin read is the one inside each root's `world.res`; that the
program scans these collections is Medium (R2-ENGINE-038, R2-ENGINE-1512). A blank cell means the root holds no
record for that key. Counts are records (units), elements (structures) or items.

## Type-6 definition id to Data.bin row

Flag 0x10 clear: the Units row whose serverID (slot 55) equals the id. Flag 0x10 set: the Humans row whose
serverID (slot 24) equals the id. EN has 241 distinct pairs, RU 250; every record resolves and no id matches two rows
(R2-ASSET-316, R2-ENGINE-1512).

| Arm | Id | EN row | EN name | EN records | RU row | RU name | RU records | Claim |
|---|---:|---:|---|---:|---:|---|---:|---|
| Units | 600 | 64 | Goblin_Pike | 33 | 64 | Goblin_Pike | 82 | R2-ASSET-316 |
| Units | 601 | 65 | Goblin_Pike.2 | 36 | 65 | Goblin_Pike.2 | 96 | R2-ASSET-316 |
| Units | 602 | 66 | Goblin_Pike.3 | 62 | 66 | Goblin_Pike.3 | 120 | R2-ASSET-316 |
| Units | 603 | 67 | Goblin_Pike.4 | 64 | 67 | Goblin_Pike.4 | 160 | R2-ASSET-316 |
| Units | 604 | 69 | Goblin_Sling | 19 | 69 | Goblin_Sling | 53 | R2-ASSET-316 |
| Units | 605 | 70 | Goblin_Sling.2 | 46 | 70 | Goblin_Sling.2 | 136 | R2-ASSET-316 |
| Units | 606 | 71 | Goblin_Sling.3 | 59 | 71 | Goblin_Sling.3 | 118 | R2-ASSET-316 |
| Units | 607 | 72 | Goblin_Sling.4 | 43 | 72 | Goblin_Sling.4 | 135 | R2-ASSET-316 |
| Units | 608 | 74 | Orc_Sword | 28 | 74 | Orc_Sword | 68 | R2-ASSET-316 |
| Units | 609 | 75 | Orc_Sword.2 | 40 | 75 | Orc_Sword.2 | 56 | R2-ASSET-316 |
| Units | 610 | 76 | Orc_Sword.3 | 95 | 76 | Orc_Sword.3 | 155 | R2-ASSET-316 |
| Units | 611 | 77 | Orc_Sword.4 | 62 | 77 | Orc_Sword.4 | 144 | R2-ASSET-316 |
| Units | 612 | 79 | Orc_Bow | 17 | 79 | Orc_Bow | 41 | R2-ASSET-316 |
| Units | 613 | 80 | Orc_Bow.2 | 60 | 80 | Orc_Bow.2 | 82 | R2-ASSET-316 |
| Units | 614 | 81 | Orc_Bow.3 | 63 | 81 | Orc_Bow.3 | 114 | R2-ASSET-316 |
| Units | 615 | 82 | Orc_Bow.4 | 41 | 82 | Orc_Bow.4 | 155 | R2-ASSET-316 |
| Units | 616 | 84 | Ogre | 23 | 84 | Ogre | 38 | R2-ASSET-316 |
| Units | 617 | 85 | Ogre.2 | 56 | 85 | Ogre.2 | 83 | R2-ASSET-316 |
| Units | 618 | 86 | Ogre.3 | 54 | 86 | Ogre.3 | 122 | R2-ASSET-316 |
| Units | 619 | 87 | Ogre.4 | 46 | 87 | Ogre.4 | 91 | R2-ASSET-316 |
| Units | 620 | 89 | Troll | 9 | 89 | Troll | 16 | R2-ASSET-316 |
| Units | 621 | 90 | Troll.2 | 19 | 90 | Troll.2 | 39 | R2-ASSET-316 |
| Units | 622 | 91 | Troll.3 | 42 | 91 | Troll.3 | 73 | R2-ASSET-316 |
| Units | 623 | 92 | Troll.4 | 64 | 92 | Troll.4 | 108 | R2-ASSET-316 |
| Units | 624 | 94 | Bat_Sonic | 61 | 94 | Bat_Sonic | 135 | R2-ASSET-316 |
| Units | 625 | 95 | Bat_Sonic.2 | 75 | 95 | Bat_Sonic.2 | 207 | R2-ASSET-316 |
| Units | 626 | 96 | Bat_Sonic.3 | 260 | 96 | Bat_Sonic.3 | 531 | R2-ASSET-316 |
| Units | 627 | 97 | Bat_Sonic.4 | 226 | 97 | Bat_Sonic.4 | 617 | R2-ASSET-316 |
| Units | 628 | 99 | Ghost | 35 | 99 | Ghost | 79 | R2-ASSET-316 |
| Units | 629 | 100 | Ghost.2 | 33 | 100 | Ghost.2 | 68 | R2-ASSET-316 |
| Units | 630 | 101 | Ghost.3 | 56 | 101 | Ghost.3 | 124 | R2-ASSET-316 |
| Units | 631 | 102 | Ghost.4 | 89 | 102 | Ghost.4 | 229 | R2-ASSET-316 |
| Units | 632 | 104 | Bee | 73 | 104 | Bee | 271 | R2-ASSET-316 |
| Units | 633 | 105 | Bee.2 | 179 | 105 | Bee.2 | 278 | R2-ASSET-316 |
| Units | 634 | 106 | Bee.3 | 157 | 106 | Bee.3 | 362 | R2-ASSET-316 |
| Units | 635 | 107 | Bee.4 | 153 | 107 | Bee.4 | 364 | R2-ASSET-316 |
| Units | 636 | 109 | Squirrel | 78 | 109 | Squirrel | 128 | R2-ASSET-316 |
| Units | 637 | 110 | Squirrel.2 | 131 | 110 | Squirrel.2 | 237 | R2-ASSET-316 |
| Units | 638 | 111 | Squirrel.3 | 218 | 111 | Squirrel.3 | 353 | R2-ASSET-316 |
| Units | 639 | 112 | Squirrel.4 | 165 | 112 | Squirrel.4 | 293 | R2-ASSET-316 |
| Units | 640 | 114 | Foot_Animated | 24 | 114 | Foot_Animated | 62 | R2-ASSET-316 |
| Units | 641 | 115 | Foot_Animated.2 | 47 | 115 | Foot_Animated.2 | 81 | R2-ASSET-316 |
| Units | 642 | 116 | Foot_Animated.3 | 89 | 116 | Foot_Animated.3 | 163 | R2-ASSET-316 |
| Units | 643 | 117 | Foot_Animated.4 | 79 | 117 | Foot_Animated.4 | 144 | R2-ASSET-316 |
| Units | 644 | 119 | Turtle | 52 | 119 | Turtle | 115 | R2-ASSET-316 |
| Units | 645 | 120 | Turtle.2 | 70 | 120 | Turtle.2 | 138 | R2-ASSET-316 |
| Units | 646 | 121 | Turtle.3 | 125 | 121 | Turtle.3 | 250 | R2-ASSET-316 |
| Units | 647 | 122 | Turtle.4 | 209 | 122 | Turtle.4 | 308 | R2-ASSET-316 |
| Units | 648 | 124 | Dragon | 7 | 124 | Dragon | 11 | R2-ASSET-316 |
| Units | 649 | 125 | Dragon.2 | 9 | 125 | Dragon.2 | 18 | R2-ASSET-316 |
| Units | 650 | 126 | Dragon.3 | 16 | 126 | Dragon.3 | 30 | R2-ASSET-316 |
| Units | 651 | 127 | Dragon.4 | 54 | 127 | Dragon.4 | 135 | R2-ASSET-316 |
| Units | 652 | 129 | Urd | 1 | 129 | Urd | 1 | R2-ASSET-316 |
| Units | 653 | 131 | Orc_Shaman.1 | 3 | 131 | Orc_Shaman.1 | 13 | R2-ASSET-316 |
| Units | 654 | 132 | Orc_Shaman.2 | 31 | 132 | Orc_Shaman.2 | 69 | R2-ASSET-316 |
| Units | 655 | 133 | Orc_Shaman.3 | 16 | 133 | Orc_Shaman.3 | 43 | R2-ASSET-316 |
| Units | 656 | 134 | Orc_Shaman.4 | 11 | 134 | Orc_Shaman.4 | 87 | R2-ASSET-316 |
| Units | 657 | 136 | F_Zombie.1 | 44 | 136 | F_Zombie.1 | 48 | R2-ASSET-316 |
| Units | 658 | 137 | F_Zombie.2 | 54 | 137 | F_Zombie.2 | 82 | R2-ASSET-316 |
| Units | 659 | 138 | F_Zombie.3 | 66 | 138 | F_Zombie.3 | 84 | R2-ASSET-316 |
| Units | 660 | 139 | F_Zombie.4 | 75 | 139 | F_Zombie.4 | 164 | R2-ASSET-316 |
| Units | 661 | 141 | A_Zombie.2 | 31 | 141 | A_Zombie.2 | 50 | R2-ASSET-316 |
| Units | 662 | 142 | A_Zombie.3 | 52 | 142 | A_Zombie.3 | 68 | R2-ASSET-316 |
| Units | 663 | 143 | A_Zombie.4 | 54 | 143 | A_Zombie.4 | 125 | R2-ASSET-316 |
| Units | 664 | 145 | F_Skeleton.1 | 39 | 145 | F_Skeleton.1 | 47 | R2-ASSET-316 |
| Units | 665 | 146 | F_Skeleton.2 | 36 | 146 | F_Skeleton.2 | 56 | R2-ASSET-316 |
| Units | 666 | 147 | F_Skeleton.3 | 75 | 147 | F_Skeleton.3 | 107 | R2-ASSET-316 |
| Units | 667 | 148 | F_Skeleton.4 | 89 | 148 | F_Skeleton.4 | 232 | R2-ASSET-316 |
| Units | 668 | 150 | A_Skeleton.1 | 13 | 150 | A_Skeleton.1 | 20 | R2-ASSET-316 |
| Units | 669 | 151 | A_Skeleton.2 | 12 | 151 | A_Skeleton.2 | 32 | R2-ASSET-316 |
| Units | 670 | 152 | A_Skeleton.3 | 34 | 152 | A_Skeleton.3 | 59 | R2-ASSET-316 |
| Units | 671 | 153 | A_Skeleton.4 | 50 | 153 | A_Skeleton.4 | 143 | R2-ASSET-316 |
| Units | 672 | 155 | M_Skeleton.3 | 20 | 155 | M_Skeleton.3 | 33 | R2-ASSET-316 |
| Units | 673 | 156 | M_Skeleton.4 | 53 | 156 | M_Skeleton.4 | 112 | R2-ASSET-316 |
| Units | 674 |  |  |  | 158 | Necromancer | 1 | R2-ASSET-316 |
| Units | 675 |  |  |  | 159 | Druid | 1 | R2-ASSET-316 |
| Units | 677 | 161 | MercenaryOrc | 1 | 161 | MercenaryOrc | 1 | R2-ASSET-316 |
| Units | 679 | 163 | Orc_ShamanLeader | 5 | 163 | Orc_ShamanLeader | 10 | R2-ASSET-316 |
| Units | 680 | 164 | Druid_Mercenary | 1 | 164 | Druid_Mercenary | 1 | R2-ASSET-316 |
| Units | 681 | 165 | Druid_Preacher | 1 | 165 | Druid_Preacher | 1 | R2-ASSET-316 |
| Units | 682 | 166 | Druid_Male2 | 21 | 166 | Druid_Male2 | 29 | R2-ASSET-316 |
| Units | 683 | 167 | Druid_Female2 | 14 | 167 | Druid_Female2 | 20 | R2-ASSET-316 |
| Units | 684 | 168 | Druid_Leader2 | 3 | 168 | Druid_Leader2 | 6 | R2-ASSET-316 |
| Units | 685 | 169 | Druid_Male3 | 14 | 169 | Druid_Male3 | 29 | R2-ASSET-316 |
| Units | 686 | 170 | Druid_Female3 | 13 | 170 | Druid_Female3 | 27 | R2-ASSET-316 |
| Units | 687 | 171 | Druid_Leader3 | 10 | 171 | Druid_Leader3 | 17 | R2-ASSET-316 |
| Units | 688 | 172 | Druid_Male4 | 15 | 172 | Druid_Male4 | 16 | R2-ASSET-316 |
| Units | 689 | 173 | Druid_Female4 | 13 | 173 | Druid_Female4 | 17 | R2-ASSET-316 |
| Units | 690 | 174 | Druid_Leader4 | 8 | 174 | Druid_Leader4 | 16 | R2-ASSET-316 |
| Units | 691 | 178 | Necro_Male1 | 8 | 178 | Necro_Male1 | 12 | R2-ASSET-316 |
| Units | 692 | 179 | Necro_Female1 | 4 | 179 | Necro_Female1 | 7 | R2-ASSET-316 |
| Units | 693 | 180 | Necro_Leader1 | 6 | 180 | Necro_Leader1 | 9 | R2-ASSET-316 |
| Units | 694 | 181 | Necro_Male2 | 13 | 181 | Necro_Male2 | 17 | R2-ASSET-316 |
| Units | 695 | 182 | Necro_Female2 | 11 | 182 | Necro_Female2 | 13 | R2-ASSET-316 |
| Units | 696 | 183 | Necro_Leader2 | 11 | 183 | Necro_Leader2 | 16 | R2-ASSET-316 |
| Units | 697 | 184 | Necro_Male3 | 18 | 184 | Necro_Male3 | 21 | R2-ASSET-316 |
| Units | 698 | 185 | Necro_Female3 | 17 | 185 | Necro_Female3 | 17 | R2-ASSET-316 |
| Units | 699 | 186 | Necro_Leader3 | 13 | 186 | Necro_Leader3 | 16 | R2-ASSET-316 |
| Units | 700 | 187 | Necro_Male4 | 19 | 187 | Necro_Male4 | 31 | R2-ASSET-316 |
| Units | 701 | 188 | Necro_Female4 | 19 | 188 | Necro_Female4 | 32 | R2-ASSET-316 |
| Units | 702 | 189 | Necro_Leader4 | 19 | 189 | Necro_Leader4 | 38 | R2-ASSET-316 |
| Units | 703 | 200 | Wolf | 109 | 200 | Wolf | 160 | R2-ASSET-316 |
| Units | 704 | 201 | Wolf.2 | 76 | 201 | Wolf.2 | 150 | R2-ASSET-316 |
| Units | 705 | 202 | Wolf.3 | 125 | 202 | Wolf.3 | 272 | R2-ASSET-316 |
| Units | 706 | 203 | Wolf.4 | 167 | 203 | Wolf.4 | 302 | R2-ASSET-316 |
| Units | 707 | 205 | Spider | 40 | 205 | Spider | 90 | R2-ASSET-316 |
| Units | 708 | 206 | Spider.2 | 67 | 206 | Spider.2 | 164 | R2-ASSET-316 |
| Units | 709 | 207 | Spider.3 | 142 | 207 | Spider.3 | 266 | R2-ASSET-316 |
| Units | 710 | 208 | Spider.4 | 114 | 208 | Spider.4 | 313 | R2-ASSET-316 |
| Units | 711 | 210 | Succubus | 19 | 210 | Succubus | 34 | R2-ASSET-316 |
| Units | 712 | 211 | Succubus.2 | 12 | 211 | Succubus.2 | 29 | R2-ASSET-316 |
| Units | 713 | 212 | Succubus.3 | 23 | 212 | Succubus.3 | 99 | R2-ASSET-316 |
| Units | 714 | 213 | Succubus.4 | 57 | 213 | Succubus.4 | 134 | R2-ASSET-316 |
| Units | 715 | 215 | Dino | 21 | 215 | Dino | 36 | R2-ASSET-316 |
| Units | 716 | 216 | Dino.2 | 47 | 216 | Dino.2 | 91 | R2-ASSET-316 |
| Units | 717 | 217 | Dino.3 | 69 | 217 | Dino.3 | 151 | R2-ASSET-316 |
| Units | 718 | 218 | Dino.4 | 98 | 218 | Dino.4 | 168 | R2-ASSET-316 |
| Units | 719 | 220 | Sheep | 72 | 220 | Sheep | 98 | R2-ASSET-316 |
| Units | 801 | 68 | Goblin_Pike.5 | 5 | 68 | Goblin_Pike.5 | 39 | R2-ASSET-316 |
| Units | 802 | 73 | Goblin_Sling.5 | 3 | 73 | Goblin_Sling.5 | 32 | R2-ASSET-316 |
| Units | 803 | 78 | Orc_Sword.5 | 2 | 78 | Orc_Sword.5 | 27 | R2-ASSET-316 |
| Units | 804 | 83 | Orc_Bow.5 | 2 | 83 | Orc_Bow.5 | 28 | R2-ASSET-316 |
| Units | 805 | 88 | Ogre.5 | 1 | 88 | Ogre.5 | 18 | R2-ASSET-316 |
| Units | 806 |  |  |  | 93 | Troll.5 | 16 | R2-ASSET-316 |
| Units | 807 | 98 | Bat_Sonic.5 | 19 | 98 | Bat_Sonic.5 | 118 | R2-ASSET-316 |
| Units | 808 | 103 | Ghost.5 | 4 | 103 | Ghost.5 | 30 | R2-ASSET-316 |
| Units | 809 | 108 | Bee.5 | 3 | 108 | Bee.5 | 99 | R2-ASSET-316 |
| Units | 810 |  |  |  | 113 | Squirrel.5 | 26 | R2-ASSET-316 |
| Units | 811 | 118 | Foot_Animated.5 | 4 | 118 | Foot_Animated.5 | 53 | R2-ASSET-316 |
| Units | 812 | 123 | Turtle.5 | 1 | 123 | Turtle.5 | 46 | R2-ASSET-316 |
| Units | 813 | 128 | Dragon.5 | 2 | 128 | Dragon.5 | 13 | R2-ASSET-316 |
| Units | 814 | 135 | Orc_Shaman.5 | 2 | 135 | Orc_Shaman.5 | 34 | R2-ASSET-316 |
| Units | 815 | 140 | F_Zombie.5 | 1 | 140 | F_Zombie.5 | 53 | R2-ASSET-316 |
| Units | 816 | 144 | A_Zombie.5 | 1 | 144 | A_Zombie.5 | 8 | R2-ASSET-316 |
| Units | 817 | 149 | F_Skeleton.5 | 5 | 149 | F_Skeleton.5 | 23 | R2-ASSET-316 |
| Units | 818 | 154 | A_Skeleton.5 | 5 | 154 | A_Skeleton.5 | 24 | R2-ASSET-316 |
| Units | 819 | 157 | M_Skeleton.5 | 4 | 157 | M_Skeleton.5 | 32 | R2-ASSET-316 |
| Units | 820 |  |  |  | 175 | Druid_Male5 | 2 | R2-ASSET-316 |
| Units | 821 | 176 | Druid_Female5 | 1 | 176 | Druid_Female5 | 3 | R2-ASSET-316 |
| Units | 822 | 177 | Druid_Leader5 | 2 | 177 | Druid_Leader5 | 4 | R2-ASSET-316 |
| Units | 823 | 190 | Necro_Male5 | 2 | 190 | Necro_Male5 | 12 | R2-ASSET-316 |
| Units | 824 | 191 | Necro_Female5 | 1 | 191 | Necro_Female5 | 7 | R2-ASSET-316 |
| Units | 825 | 192 | Necro_Leader5 | 2 | 192 | Necro_Leader5 | 9 | R2-ASSET-316 |
| Units | 826 | 204 | Wolf.5 | 2 | 204 | Wolf.5 | 28 | R2-ASSET-316 |
| Units | 827 | 209 | Spider.5 | 5 | 209 | Spider.5 | 60 | R2-ASSET-316 |
| Units | 828 | 214 | Succubus.5 | 5 | 214 | Succubus.5 | 28 | R2-ASSET-316 |
| Units | 829 | 219 | Dino.5 | 4 | 219 | Dino.5 | 31 | R2-ASSET-316 |
| Units | 830 | 193 | Necro_Fire | 11 | 193 | Necro_Fire | 11 | R2-ASSET-316 |
| Units | 831 | 194 | Necro_Air | 10 | 194 | Necro_Air | 10 | R2-ASSET-316 |
| Units | 832 | 195 | Necro_Earth | 10 | 195 | Necro_Earth | 10 | R2-ASSET-316 |
| Units | 833 | 196 | Necro_Water | 11 | 196 | Necro_Water | 11 | R2-ASSET-316 |
| Units | 834 | 197 | Necro_Astral | 8 | 197 | Necro_Astral | 8 | R2-ASSET-316 |
| Units | 835 | 198 | GoodNecroLeader | 1 | 198 | GoodNecroLeader | 1 | R2-ASSET-316 |
| Units | 836 | 199 | GoodNecro | 1 | 199 | GoodNecro | 1 | R2-ASSET-316 |
| Units | 999 | 26 | Catapult | 10 | 26 | Catapult | 11 | R2-ASSET-316 |
| Humans | 1 | 1 | Man_Unarmed | 2 | 1 | Man_Unarmed | 2 | R2-ASSET-316 |
| Humans | 407 | 75 | NPC07_4 | 4 | 75 | NPC07_4 | 4 | R2-ASSET-316 |
| Humans | 408 | 76 | NPC08_4 | 2 | 76 | NPC08_4 | 2 | R2-ASSET-316 |
| Humans | 410 | 78 | NPC10_4 | 4 | 78 | NPC10_4 | 4 | R2-ASSET-316 |
| Humans | 412 | 80 | NPC12_4 | 4 | 80 | NPC12_4 | 4 | R2-ASSET-316 |
| Humans | 1003 | 86 | A_PeasantGuard1 | 2 | 86 | A_PeasantGuard1 | 2 | R2-ASSET-316 |
| Humans | 1004 |  |  |  | 87 | U_PeasantLeader1 | 1 | R2-ASSET-316 |
| Humans | 1006 | 89 | F_PeasantGuard2 | 1 | 89 | F_PeasantGuard2 | 1 | R2-ASSET-316 |
| Humans | 1008 | 91 | U_PeasantLeader2 | 2 | 91 | U_PeasantLeader2 | 2 | R2-ASSET-316 |
| Humans | 1014 | 97 | F_PeasantGuard4 | 2 | 97 | F_PeasantGuard4 | 2 | R2-ASSET-316 |
| Humans | 1020 | 103 | F_BrigandLeader1 | 1 | 103 | F_BrigandLeader1 | 1 | R2-ASSET-316 |
| Humans | 1051 | 134 | F_Knight3 | 3 | 134 | F_Knight3 | 3 | R2-ASSET-316 |
| Humans | 1052 | 135 | A_Knight3 | 2 | 135 | A_Knight3 | 3 | R2-ASSET-316 |
| Humans | 1053 | 136 | H_Knight3 | 1 | 136 | H_Knight3 | 1 | R2-ASSET-316 |
| Humans | 1060 |  |  |  | 143 | F_KnightLeader4 | 2 | R2-ASSET-316 |
| Humans | 1084 | 167 | U_FemalePeasant2 | 2 | 167 | U_FemalePeasant2 | 2 | R2-ASSET-316 |
| Humans | 1085 | 168 | U_FemalePeasant3 | 2 | 168 | U_FemalePeasant3 | 2 | R2-ASSET-316 |
| Humans | 1086 | 169 | U_FemalePeasant4 | 2 | 169 | U_FemalePeasant4 | 2 | R2-ASSET-316 |
| Humans | 2000 | 178 | CH_Scrakan | 2 | 178 | CH_Scrakan | 2 | R2-ASSET-316 |
| Humans | 2001 | 179 | CH_Recruit | 1 | 179 | CH_Recruit | 1 | R2-ASSET-316 |
| Humans | 2003 | 181 | CH_Elf | 1 | 181 | CH_Elf | 1 | R2-ASSET-316 |
| Humans | 2005 | 183 | CH_Orcologist | 1 | 183 | CH_Orcologist | 1 | R2-ASSET-316 |
| Humans | 2006 | 184 | CH_UncleGeneral | 1 | 184 | CH_UncleGeneral | 1 | R2-ASSET-316 |
| Humans | 2007 | 185 | CH_KingZombie | 1 | 185 | CH_KingZombie | 1 | R2-ASSET-316 |
| Humans | 2008 | 186 | CH_Ranger | 2 | 186 | CH_Ranger | 2 | R2-ASSET-316 |
| Humans | 2014 | 192 | CH_KingZombieUnmasked | 1 | 192 | CH_KingZombieUnmasked | 1 | R2-ASSET-316 |
| Humans | 2015 | 193 | CH_MageMonsterTrap | 2 | 193 | CH_MageMonsterTrap | 2 | R2-ASSET-316 |
| Humans | 2016 | 194 | CH_MageAmuletKeeper | 1 | 194 | CH_MageAmuletKeeper | 1 | R2-ASSET-316 |
| Humans | 2017 | 195 | CH_SmithBewitched | 1 | 195 | CH_SmithBewitched | 1 | R2-ASSET-316 |
| Humans | 2018 | 196 | CH_SmithDispelled | 2 | 196 | CH_SmithDispelled | 2 | R2-ASSET-316 |
| Humans | 2019 | 197 | CH_AtamanRunaway | 1 | 197 | CH_AtamanRunaway | 1 | R2-ASSET-316 |
| Humans | 2020 | 198 | CH_AtamanCamped | 1 | 198 | CH_AtamanCamped | 1 | R2-ASSET-316 |
| Humans | 2021 | 199 | CH_SwordCarrier | 1 | 199 | CH_SwordCarrier | 1 | R2-ASSET-316 |
| Humans | 2022 | 200 | CH_Xara | 1 | 200 | CH_Xara | 1 | R2-ASSET-316 |
| Humans | 2023 | 201 | CH_Sister | 1 | 201 | CH_Sister | 1 | R2-ASSET-316 |
| Humans | 2024 | 202 | CH_Son | 1 | 202 | CH_Son | 1 | R2-ASSET-316 |
| Humans | 2025 | 203 | CH_Sphinx | 1 | 203 | CH_Sphinx | 1 | R2-ASSET-316 |
| Humans | 2026 | 204 | CH_Girl | 1 | 204 | CH_Girl | 1 | R2-ASSET-316 |
| Humans | 2027 | 205 | CH_PyramidFire | 1 | 205 | CH_PyramidFire | 1 | R2-ASSET-316 |
| Humans | 2028 | 206 | CH_PyramidWater | 1 | 206 | CH_PyramidWater | 1 | R2-ASSET-316 |
| Humans | 2029 | 207 | CH_PyramidEarth | 1 | 207 | CH_PyramidEarth | 1 | R2-ASSET-316 |
| Humans | 2030 | 208 | CH_PyramidAir | 1 | 208 | CH_PyramidAir | 1 | R2-ASSET-316 |
| Humans | 2100 | 209 | 2F_Brigand1 | 13 | 209 | 2F_Brigand1 | 18 | R2-ASSET-316 |
| Humans | 2101 | 210 | 2A_Brigand1 | 15 | 210 | 2A_Brigand1 | 26 | R2-ASSET-316 |
| Humans | 2102 | 211 | 2M_Brigand1 | 12 | 211 | 2M_Brigand1 | 21 | R2-ASSET-316 |
| Humans | 2103 | 212 | 2F_BrigandLeader1 | 8 | 212 | 2F_BrigandLeader1 | 12 | R2-ASSET-316 |
| Humans | 2104 | 213 | 2F_Brigand2 | 31 | 213 | 2F_Brigand2 | 51 | R2-ASSET-316 |
| Humans | 2105 | 214 | 2A_Brigand2 | 59 | 214 | 2A_Brigand2 | 88 | R2-ASSET-316 |
| Humans | 2106 | 215 | 2M_Brigand2 | 43 | 215 | 2M_Brigand2 | 72 | R2-ASSET-316 |
| Humans | 2107 | 216 | 2F_BrigandLeader2 | 26 | 216 | 2F_BrigandLeader2 | 44 | R2-ASSET-316 |
| Humans | 2108 | 217 | 2U_PeasantMale | 69 | 217 | 2U_PeasantMale | 114 | R2-ASSET-316 |
| Humans | 2109 | 218 | 2U_PeasantFemale | 62 | 218 | 2U_PeasantFemale | 103 | R2-ASSET-316 |
| Humans | 2110 | 219 | 2F_PeasantGuard | 59 | 219 | 2F_PeasantGuard | 71 | R2-ASSET-316 |
| Humans | 2111 | 220 | 2A_PeasantGuard | 53 | 220 | 2A_PeasantGuard | 65 | R2-ASSET-316 |
| Humans | 2112 | 221 | 2U_PeasantLeader | 45 | 221 | 2U_PeasantLeader | 61 | R2-ASSET-316 |
| Humans | 2113 | 222 | 2F_Knight1 | 20 | 222 | 2F_Knight1 | 22 | R2-ASSET-316 |
| Humans | 2114 | 223 | 2A_Knight1 | 13 | 223 | 2A_Knight1 | 20 | R2-ASSET-316 |
| Humans | 2115 | 224 | 2H_Knight1 | 2 | 224 | 2H_Knight1 | 2 | R2-ASSET-316 |
| Humans | 2116 | 225 | 2M_Knight1 | 5 | 225 | 2M_Knight1 | 12 | R2-ASSET-316 |
| Humans | 2117 | 226 | 2F_KnightLeader1 | 7 | 226 | 2F_KnightLeader1 | 8 | R2-ASSET-316 |
| Humans | 2118 | 227 | 2F_Knight2 | 3 | 227 | 2F_Knight2 | 13 | R2-ASSET-316 |
| Humans | 2119 | 228 | 2A_Knight2 | 20 | 228 | 2A_Knight2 | 27 | R2-ASSET-316 |
| Humans | 2120 | 229 | 2H_Knight2 | 12 | 229 | 2H_Knight2 | 13 | R2-ASSET-316 |
| Humans | 2121 | 230 | 2M_Knight2 | 15 | 230 | 2M_Knight2 | 27 | R2-ASSET-316 |
| Humans | 2122 | 231 | 2F_KnightLeader2 | 4 | 231 | 2F_KnightLeader2 | 7 | R2-ASSET-316 |
| Humans | 2123 | 232 | 2F_Knight3 | 22 | 232 | 2F_Knight3 | 43 | R2-ASSET-316 |
| Humans | 2124 | 233 | 2A_Knight3 | 28 | 233 | 2A_Knight3 | 74 | R2-ASSET-316 |
| Humans | 2125 | 234 | 2H_Knight3 | 9 | 234 | 2H_Knight3 | 16 | R2-ASSET-316 |
| Humans | 2126 | 235 | 2M_Knight3 | 14 | 235 | 2M_Knight3 | 39 | R2-ASSET-316 |
| Humans | 2127 | 236 | 2F_KnightLeader3 | 6 | 236 | 2F_KnightLeader3 | 20 | R2-ASSET-316 |
| Humans | 2128 | 237 | 2F_Knight4 | 29 | 237 | 2F_Knight4 | 57 | R2-ASSET-316 |
| Humans | 2129 | 238 | 2A_Knight4 | 36 | 238 | 2A_Knight4 | 122 | R2-ASSET-316 |
| Humans | 2130 | 239 | 2H_Knight4 | 17 | 239 | 2H_Knight4 | 30 | R2-ASSET-316 |
| Humans | 2131 | 240 | 2M_Knight4 | 19 | 240 | 2M_Knight4 | 85 | R2-ASSET-316 |
| Humans | 2132 | 241 | 2F_KnightLeader4 | 12 | 241 | 2F_KnightLeader4 | 96 | R2-ASSET-316 |
| Humans | 2133 | 242 | 2F_Hunter2 | 12 | 242 | 2F_Hunter2 | 17 | R2-ASSET-316 |
| Humans | 2134 | 243 | 2A_Hunter2 | 21 | 243 | 2A_Hunter2 | 28 | R2-ASSET-316 |
| Humans | 2135 | 244 | 2F_Hunter3 | 8 | 244 | 2F_Hunter3 | 26 | R2-ASSET-316 |
| Humans | 2136 | 245 | 2A_Hunter3 | 15 | 245 | 2A_Hunter3 | 37 | R2-ASSET-316 |
| Humans | 2137 | 246 | 2F_Hunter4 | 12 | 246 | 2F_Hunter4 | 25 | R2-ASSET-316 |
| Humans | 2138 | 247 | 2A_Hunter4 | 12 | 247 | 2A_Hunter4 | 23 | R2-ASSET-316 |
| Humans | 2139 | 248 | 2F_Sword3_NPC | 22 | 248 | 2F_Sword3_NPC | 22 | R2-ASSET-316 |
| Humans | 2140 | 249 | 2F_Axe3_NPC | 14 | 249 | 2F_Axe3_NPC | 14 | R2-ASSET-316 |
| Humans | 2141 | 250 | 2F_Club3_NPC | 15 | 250 | 2F_Club3_NPC | 15 | R2-ASSET-316 |
| Humans | 2142 | 251 | 2F_Pike3_NPC | 13 | 251 | 2F_Pike3_NPC | 13 | R2-ASSET-316 |
| Humans | 2143 | 252 | 2F_Bow3_NPC | 22 | 252 | 2F_Bow3_NPC | 22 | R2-ASSET-316 |
| Humans | 2144 | 253 | 2F_Knight3_73 | 19 | 253 | 2F_Knight3_73 | 19 | R2-ASSET-316 |
| Humans | 2145 | 254 | 2A_Knight3_73 | 7 | 254 | 2A_Knight3_73 | 7 | R2-ASSET-316 |
| Humans | 2146 | 255 | 2H_Knight3_73 | 4 | 255 | 2H_Knight3_73 | 4 | R2-ASSET-316 |
| Humans | 2147 | 256 | 2M_Knight3_73 | 12 | 256 | 2M_Knight3_73 | 17 | R2-ASSET-316 |
| Humans | 2148 | 257 | 2F_KnightLeader3_73 | 4 | 257 | 2F_KnightLeader3_73 | 4 | R2-ASSET-316 |
| Humans | 2149 |  |  |  | 258 | 2F_Knight2_51 | 3 | R2-ASSET-316 |
| Humans | 2150 |  |  |  | 259 | 2A_Knight2_51 | 2 | R2-ASSET-316 |
| Humans | 2152 | 261 | 2M_Knight2_51 | 1 | 261 | 2M_Knight2_51 | 1 | R2-ASSET-316 |
| Humans | 10230 | 31 | Mage_Hero | 1 | 31 | Mage_Hero | 1 | R2-ASSET-316 |

## Type-4 kind to Data.bin Buildings row

The kind indexes the Buildings collection by position (R2-ENGINE-1289); every kind has a row. Kind 0x21 holds the
28-byte elements (EN 41, RU 170, R2-ASSET-317, R2-ASSET-318). EN has 176 distinct kinds, RU 178.

| Kind | Row name | Parameters 0..5 | EN elements | RU elements | Claim |
|---:|---|---|---:|---:|---|
| 1 | Goblin Hut | 3;2;4;300;55;63 | 65 | 135 | R2-ASSET-317 |
| 2 | Goblin Hut 2 | 3;2;4;300;55;63 | 65 | 135 | R2-ASSET-317 |
| 3 | Orc Hut | 3;2;4;300;23;63 | 56 | 103 | R2-ASSET-317 |
| 4 | Orc Hut 2 | 3;2;4;300;23;63 | 46 | 93 | R2-ASSET-317 |
| 5 | Mill | 3;3;4;1000;186;511 | 14 | 21 | R2-ASSET-317 |
| 6 | Mill 2 | 3;3;4;1000;250;511 | 10 | 16 | R2-ASSET-317 |
| 7 | Mill 3 | 3;3;4;1000;506;511 | 13 | 20 | R2-ASSET-317 |
| 8 | Well | 2;2;0;1000;15;15 | 50 | 117 | R2-ASSET-317 |
| 9 | Pizant Hut | 3;3;5;1000;510;511 | 21 | 48 | R2-ASSET-317 |
| 10 | Pizant Hut 2 | 3;3;5;1000;506;511 | 26 | 53 | R2-ASSET-317 |
| 11 | Church | 4;3;7;30000;3830;4095 | 19 | 24 | R2-ASSET-317 |
| 12 | Tower | 2;2;5;30000;15;15 | 51 | 72 | R2-ASSET-317 |
| 13 | Tower 2 | 2;2;5;30000;15;15 | 32 | 40 | R2-ASSET-317 |
| 14 | Cave | 3;3;3;1000;63;63 | 26 | 77 | R2-ASSET-317 |
| 15 | Magic Well 2 | 2;2;0;3;15;15 | 31 | 59 | R2-ASSET-317 |
| 16 | Magic Well 3 | 3;3;0;3;511;511 | 21 | 52 | R2-ASSET-317 |
| 17 | Human House 3a | 3;3;5;2000;186;511 | 24 | 46 | R2-ASSET-317 |
| 18 | Human House 3b | 3;3;5;2000;447;511 | 22 | 40 | R2-ASSET-317 |
| 19 | Human House 2a | 3;3;5;2000;511;511 | 21 | 31 | R2-ASSET-317 |
| 20 | Human House 2b | 3;3;5;2000;447;511 | 17 | 31 | R2-ASSET-317 |
| 21 | Ogre House | 3;3;4;500;447;511 | 42 | 75 | R2-ASSET-317 |
| 22 | Ogre House 2 | 3;3;4;500;447;511 | 33 | 60 | R2-ASSET-317 |
| 23 | Ogre House 3 | 3;3;4;500;447;511 | 35 | 61 | R2-ASSET-317 |
| 24 | Bee House | 1;1;2;100;1;1 | 92 | 207 | R2-ASSET-317 |
| 25 | Bee House 2 | 1;1;2;100;1;1 | 76 | 171 | R2-ASSET-317 |
| 26 | Leg's House | 2;1;3;100;3;3 | 106 | 214 | R2-ASSET-317 |
| 27 | Magic Symbol | 3;3;3;0;0;511 | 57 | 68 | R2-ASSET-317 |
| 28 | Switch | 1;1;1;1;1;1 | 20 | 28 | R2-ASSET-317 |
| 29 | Switch 2 | 1;1;1;1;1;1 | 16 | 22 | R2-ASSET-317 |
| 30 | Ruins | 3;1;0;0;5;7 | 14 | 85 | R2-ASSET-317 |
| 31 | Ruins 2 | 2;2;0;0;13;15 | 27 | 49 | R2-ASSET-317 |
| 32 | Ruins 3 | 1;1;0;0;1;1 | 48 | 105 | R2-ASSET-317 |
| 33 | Vertical Wooden Bridge | 1;1;0;0;1;1 | 41 | 170 | R2-ASSET-317 |
| 34 | Shop | 3;3;4;0;447;511 | 10 | 30 | R2-ASSET-317 |
| 35 | Shop 2 | 3;3;4;0;447;511 | 15 | 52 | R2-ASSET-317 |
| 36 | Hanging Peasant | 2;1;0;0;3;3 | 34 | 63 | R2-ASSET-317 |
| 38 | Horisontal Bridge | 6;4;0;0;16515135;16777215 | 41 | 61 | R2-ASSET-317 |
| 39 | King's Grave 1 | 2;2;0;0;15;15 | 45 | 60 | R2-ASSET-317 |
| 40 | King's Grave 2 | 2;2;0;0;15;15 | 53 | 73 | R2-ASSET-317 |
| 41 | Common Grave 1 | 1;1;0;0;1;1 | 45 | 147 | R2-ASSET-317 |
| 42 | Common Grave 2 | 1;1;0;0;1;1 | 110 | 180 | R2-ASSET-317 |
| 43 | Inn 1 | 3;3;3;30000;191;511 | 2 | 3 | R2-ASSET-317 |
| 44 | Inn 2 | 3;3;3;30000;511;511 | 2 | 3 | R2-ASSET-317 |
| 45 | Inn 3 | 3;3;3;30000;447;511 | 3 | 3 | R2-ASSET-317 |
| 46 | Guard Tower 1 | 2;3;5;30000;63;63 | 28 | 42 | R2-ASSET-317 |
| 47 | Guard Tower 2 | 2;3;5;30000;63;63 | 28 | 42 | R2-ASSET-317 |
| 48 | Teleport Pad | 3;3;5;0;0;511 | 45 | 56 | R2-ASSET-317 |
| 49 | Training School 1 | 3;4;3;30000;4090;4095 | 1 | 5 | R2-ASSET-317 |
| 50 | Training School 2 | 3;4;3;30000;4090;4095 |  | 2 | R2-ASSET-317 |
| 51 | Training School 3 | 3;4;3;30000;3578;4095 | 3 | 9 | R2-ASSET-317 |
| 52 | Vertical Bridge | 4;6;0;0;10066329;16777215 | 24 | 33 | R2-ASSET-317 |
| 53 | Camp Fire | 1;1;1;0;1;1 | 135 | 240 | R2-ASSET-317 |
| 54 | Skrakan Tower | 3;3;5;0;511;511 |  | 3 | R2-ASSET-317 |
| 55 | Skrakan Tower 2 | 4;3;5;0;1791;4095 | 2 | 5 | R2-ASSET-317 |
| 56 | Mage's tower | 2;2;5;30000;15;15 | 46 | 77 | R2-ASSET-317 |
| 57 | Castle | 11;4;5;30000;2132801535;2132801535 | 3 | 5 | R2-ASSET-317 |
| 58 | Sphinx | 3;2;3;0;63;63 | 11 | 20 | R2-ASSET-317 |
| 59 | Sphinx 2 | 3;2;3;0;63;63 | 7 | 14 | R2-ASSET-317 |
| 60 | Sphinx 3 | 3;2;3;0;63;63 | 6 | 9 | R2-ASSET-317 |
| 61 | Sphinx 4 | 3;2;3;0;63;63 | 2 | 5 | R2-ASSET-317 |
| 62 | Sphinx 5 | 3;2;3;0;63;63 | 5 | 11 | R2-ASSET-317 |
| 63 | Sphinx 6 | 3;2;3;0;63;63 | 3 | 5 | R2-ASSET-317 |
| 64 | Sphinx 7 | 3;2;3;0;63;63 | 2 | 7 | R2-ASSET-317 |
| 65 | Sphinx 8 | 3;2;3;0;63;63 | 6 | 12 | R2-ASSET-317 |
| 66 | Multiplayer Church | 4;3;7;0;3830;4095 | 28 | 109 | R2-ASSET-317 |
| 67 | Multiplayer Inn 1 | 3;3;3;0;191;511 | 5 | 13 | R2-ASSET-317 |
| 68 | Multiplayer Inn 2 | 3;3;3;0;511;511 | 6 | 16 | R2-ASSET-317 |
| 69 | Multiplayer Inn 3 | 3;3;3;0;447;511 | 14 | 39 | R2-ASSET-317 |
| 70 | Pointer 1 | 1;1;0;0;1;1 | 5 | 5 | R2-ASSET-317 |
| 71 | Pointer 2 | 1;1;0;0;1;1 | 1 | 1 | R2-ASSET-317 |
| 72 | Pointer 3 | 1;1;0;0;1;1 | 5 | 5 | R2-ASSET-317 |
| 73 | Pointer 4 | 1;1;0;0;1;1 | 1 | 1 | R2-ASSET-317 |
| 74 | Pointer 5 | 1;1;0;0;1;1 | 6 | 6 | R2-ASSET-317 |
| 75 | Pointer 6 | 1;1;0;0;1;1 | 3 | 3 | R2-ASSET-317 |
| 76 | Altar 1 | 2;2;0;1;15;15 | 6 | 8 | R2-ASSET-317 |
| 77 | Altar 2 | 2;2;0;1;15;15 | 7 | 8 | R2-ASSET-317 |
| 78 | Necro House 1 | 3;2;5;500;63;63 | 26 | 35 | R2-ASSET-317 |
| 79 | Necro House 1 | 3;2;5;500;63;63 | 26 | 33 | R2-ASSET-317 |
| 80 | Necro House 1 | 3;2;5;500;63;63 | 17 | 25 | R2-ASSET-317 |
| 81 | Shed 1 | 3;3;5;300;511;511 | 21 | 29 | R2-ASSET-317 |
| 82 | Shed 2 | 3;3;5;300;250;511 | 20 | 29 | R2-ASSET-317 |
| 83 | Shed 3 | 3;3;5;300;510;511 | 22 | 30 | R2-ASSET-317 |
| 84 | Orc Shaman House 1 | 3;2;5;400;63;63 | 9 | 18 | R2-ASSET-317 |
| 85 | Orc Shaman House 2 | 3;3;5;400;186;511 | 16 | 29 | R2-ASSET-317 |
| 86 | Orc Shaman House 3 | 3;2;5;400;63;63 | 6 | 17 | R2-ASSET-317 |
| 87 | Big Human House 1 | 4;4;5;800;61439;65535 | 11 | 17 | R2-ASSET-317 |
| 88 | Big Human House 2 | 4;4;5;800;65534;65535 | 5 | 12 | R2-ASSET-317 |
| 89 | Big Human House 3 | 4;4;5;800;65534;65535 | 5 | 9 | R2-ASSET-317 |
| 90 | Kaarg Shop 1 | 4;4;5;30000;61439;65535 | 3 | 4 | R2-ASSET-317 |
| 91 | Kaarg Shop 2 | 4;4;5;30000;28671;65535 | 5 | 5 | R2-ASSET-317 |
| 92 | Kaarg Shop 3 | 4;4;5;30000;65535;65535 | 1 | 1 | R2-ASSET-317 |
| 93 | Multiplayer Kaarg Shop 1 | 4;4;5;0;61439;65535 | 4 | 10 | R2-ASSET-317 |
| 94 | Multiplayer Kaarg Shop 2 | 4;4;5;0;28671;65535 | 6 | 18 | R2-ASSET-317 |
| 95 | Multiplayer Kaarg Shop 3 | 4;4;5;0;65535;65535 | 4 | 13 | R2-ASSET-317 |
| 96 | Kaarg Inn 1 | 3;3;3;30000;447;511 | 2 | 2 | R2-ASSET-317 |
| 97 | Kaarg Inn 2 | 3;3;3;30000;191;511 | 1 | 1 | R2-ASSET-317 |
| 98 | Kaarg Inn 3 | 3;3;3;30000;511;511 | 4 | 5 | R2-ASSET-317 |
| 99 | Multiplayer Kaarg Inn 1 | 3;3;3;0;447;511 | 6 | 14 | R2-ASSET-317 |
| 100 | Multiplayer Kaarg Inn 2 | 3;3;3;0;191;511 | 4 | 11 | R2-ASSET-317 |
| 101 | Multiplayer Kaarg Inn 3 | 3;3;3;0;511;511 | 3 | 10 | R2-ASSET-317 |
| 102 | Druid Shop 1 | 4;3;7;30000;3839;4095 | 1 | 1 | R2-ASSET-317 |
| 103 | Druid Shop 2 | 4;3;7;30000;4095;4095 | 1 | 1 | R2-ASSET-317 |
| 104 | Druid Shop 3 | 4;3;7;30000;4095;4095 | 3 | 3 | R2-ASSET-317 |
| 105 | Multiplayer Druid Shop 1 | 4;3;7;0;3839;4095 | 1 | 6 | R2-ASSET-317 |
| 106 | Multiplayer Druid Shop 2 | 4;3;7;0;4095;4095 | 8 | 10 | R2-ASSET-317 |
| 107 | Multiplayer Druid Shop 3 | 4;3;7;0;4095;4095 | 2 | 12 | R2-ASSET-317 |
| 108 | Druid Inn 1 | 4;3;7;30000;1791;4095 | 6 | 6 | R2-ASSET-317 |
| 109 | Druid Inn 2 | 4;3;7;30000;1791;4095 | 2 | 2 | R2-ASSET-317 |
| 110 | Druid Inn 3 | 4;3;7;30000;2047;4095 | 3 | 3 | R2-ASSET-317 |
| 111 | Multiplayer Druid Inn 1 | 4;3;7;0;1791;4095 | 4 | 6 | R2-ASSET-317 |
| 112 | Multiplayer Druid Inn 2 | 4;3;7;0;1791;4095 | 5 | 15 | R2-ASSET-317 |
| 113 | Multiplayer Druid Inn 3 | 4;3;7;0;2047;4095 | 3 | 8 | R2-ASSET-317 |
| 114 | Barbarian Hut 1 | 3;2;3;300;55;63 | 22 | 30 | R2-ASSET-317 |
| 115 | Barbarian Hut 2 | 3;2;3;300;63;63 | 15 | 31 | R2-ASSET-317 |
| 116 | Barbarian Hut 3 | 3;2;3;300;63;63 | 33 | 42 | R2-ASSET-317 |
| 117 | Crater | 4;4;0;0;1632;65535 | 7 | 11 | R2-ASSET-317 |
| 118 | Vigvam 1 | 2;2;3;0;15;15 | 21 | 29 | R2-ASSET-317 |
| 119 | Vigvam 2 | 2;2;3;0;15;15 | 19 | 31 | R2-ASSET-317 |
| 120 | Vigvam 3 | 2;2;3;0;15;15 | 16 | 27 | R2-ASSET-317 |
| 121 | Open Grave 1 | 3;3;3;0;443;511 | 37 | 71 | R2-ASSET-317 |
| 122 | Open Grave 2 | 3;3;3;0;254;511 | 47 | 94 | R2-ASSET-317 |
| 123 | Open Grave 3 | 3;3;3;0;251;511 | 39 | 72 | R2-ASSET-317 |
| 124 | Tent 1 | 2;2;3;80;15;15 | 16 | 25 | R2-ASSET-317 |
| 125 | Tent 2 | 2;2;3;80;15;15 | 19 | 27 | R2-ASSET-317 |
| 126 | Tent 3 | 2;2;3;80;15;15 | 13 | 21 | R2-ASSET-317 |
| 127 | Big tent 1 | 3;3;3;100;511;511 | 9 | 19 | R2-ASSET-317 |
| 128 | Big tent 2 | 3;3;3;100;511;511 | 1 | 8 | R2-ASSET-317 |
| 129 | Big tent 3 | 3;3;3;100;511;511 | 1 | 2 | R2-ASSET-317 |
| 130 | Torture 1 | 2;2;3;0;15;15 | 21 | 36 | R2-ASSET-317 |
| 131 | Torture 2 | 2;2;3;0;15;15 | 18 | 39 | R2-ASSET-317 |
| 132 | Torture 3 | 2;2;3;0;15;15 | 19 | 34 | R2-ASSET-317 |
| 133 | Sclep 1 | 3;3;3;0;447;511 | 28 | 40 | R2-ASSET-317 |
| 134 | Sclep 2 | 3;3;3;0;511;511 | 32 | 48 | R2-ASSET-317 |
| 135 | Sclep 3 | 3;3;3;0;511;511 | 21 | 28 | R2-ASSET-317 |
| 136 | Kaarg Tower 1 | 2;1;5;30000;3;3 | 18 | 38 | R2-ASSET-317 |
| 137 | Kaarg Tower 2 | 2;1;5;30000;3;3 | 28 | 65 | R2-ASSET-317 |
| 138 | Dragon's nest | 3;3;4;0;16;511 | 26 | 62 | R2-ASSET-317 |
| 139 | Necro Tower | 3;3;8;30000;511;511 | 6 | 13 | R2-ASSET-317 |
| 140 | Druid House 1 | 4;3;7;30000;1791;4095 | 20 | 30 | R2-ASSET-317 |
| 141 | Druid House 2 | 4;3;7;30000;1791;4095 | 13 | 27 | R2-ASSET-317 |
| 142 | Druid House 3 | 4;3;7;30000;1791;4095 | 19 | 29 | R2-ASSET-317 |
| 143 | Grave Doors 1 | 5;1;3;0;27;31 | 14 | 23 | R2-ASSET-317 |
| 144 | Grave Doors 2 | 5;1;3;0;27;31 | 11 | 17 | R2-ASSET-317 |
| 145 | BlackSmith 1 | 4;3;4;500;4095;4095 | 7 | 13 | R2-ASSET-317 |
| 146 | BlackSmith 2 | 4;3;4;500;4095;4095 | 17 | 21 | R2-ASSET-317 |
| 147 | StoneHenge 1 | 1;2;4;0;15;15 | 10 | 14 | R2-ASSET-317 |
| 148 | StoneHenge 2 | 2;2;4;0;15;15 | 12 | 14 | R2-ASSET-317 |
| 149 | StoneHenge 3 | 2;2;4;0;15;15 | 10 | 13 | R2-ASSET-317 |
| 150 | StoneHenge 4 | 2;2;4;0;15;15 | 11 | 14 | R2-ASSET-317 |
| 151 | StoneHenge 5 | 2;2;4;0;15;15 | 8 | 10 | R2-ASSET-317 |
| 152 | StoneHenge 6 | 2;2;4;0;15;15 | 9 | 11 | R2-ASSET-317 |
| 153 | StoneHenge 7 | 2;2;4;0;15;15 | 7 | 10 | R2-ASSET-317 |
| 154 | StoneHenge 8 | 2;2;4;0;15;15 | 11 | 13 | R2-ASSET-317 |
| 155 | Barbarian Hut 4 | 4;3;4;300;3327;4095 | 6 | 10 | R2-ASSET-317 |
| 156 | Barbarian Hut 5 | 4;3;4;300;2047;4095 | 11 | 23 | R2-ASSET-317 |
| 157 | Barbarian Hut 6 | 4;4;4;300;61439;65535 | 18 | 25 | R2-ASSET-317 |
| 158 | Barbarian Hut 7 | 4;4;4;300;53239;65535 | 7 | 10 | R2-ASSET-317 |
| 159 | Barbarian Hut 8 | 4;4;4;300;61431;65535 | 16 | 22 | R2-ASSET-317 |
| 160 | Barbarian Hut 9 | 4;4;4;300;61431;65535 | 24 | 28 | R2-ASSET-317 |
| 161 | Kaarg Hut 1 | 3;3;3;300;511;511 | 14 | 26 | R2-ASSET-317 |
| 162 | Kaarg Hut 2 | 3;3;3;300;511;511 | 8 | 17 | R2-ASSET-317 |
| 163 | Kaarg Hut 3 | 3;3;3;300;511;511 | 7 | 21 | R2-ASSET-317 |
| 164 | Kaarg Hut 4 | 3;3;3;300;511;511 | 8 | 25 | R2-ASSET-317 |
| 165 | Kaarg Hut 5 | 3;3;3;300;511;511 | 5 | 15 | R2-ASSET-317 |
| 166 | Kaarg Hut 6 | 3;3;3;300;511;511 | 5 | 15 | R2-ASSET-317 |
| 167 | Mine 1 | 3;3;3;300;511;511 | 12 | 18 | R2-ASSET-317 |
| 168 | Mine 2 | 3;3;3;300;511;511 | 9 | 14 | R2-ASSET-317 |
| 169 | Mine 3 | 3;3;3;300;511;511 | 13 | 18 | R2-ASSET-317 |
| 170 | Ore 1 | 2;2;3;0;15;15 | 24 | 35 | R2-ASSET-317 |
| 171 | Ore 2 | 2;2;3;0;15;15 | 23 | 29 | R2-ASSET-317 |
| 172 | Ore 3 | 2;2;3;0;15;15 | 18 | 26 | R2-ASSET-317 |
| 173 | StoneCrusher 1 | 3;2;3;300;63;63 | 3 | 6 | R2-ASSET-317 |
| 174 | StoneCrusher 2 | 3;2;3;300;63;63 | 7 | 8 | R2-ASSET-317 |
| 175 | StoneCrusher 3 | 2;3;3;300;63;63 | 5 | 5 | R2-ASSET-317 |
| 176 | Pyramid 1 | 3;3;3;0;191;511 | 3 | 6 | R2-ASSET-317 |
| 177 | Pyramid 2 | 3;2;3;0;63;63 | 4 | 4 | R2-ASSET-317 |
| 178 | Pyramid 3 | 3;2;3;0;63;63 | 1 | 2 | R2-ASSET-317 |
| 179 | Pyramid 4 | 3;2;3;0;63;63 | 1 | 2 | R2-ASSET-317 |

## Sack item to Data.bin row

The item id low word splits into A (bits 8-11) and D (bits 0-4); for A = 14 the row is the low byte. A selects the
table (R2-ENGINE-1513): 1 Weapons, 2 Shields, 3..13 Armors, 14 MagicItems (table names by use, Medium). The row's Slot parameter equals A for
every weapon, shield and armour (R2-ASSET-319). Items: EN 139, RU 327.

| A | Table | D or row | Row name | EN items | RU items | Claim |
|---:|---|---:|---|---:|---:|---|
| 1 | Weapons | 2 | Dagger | 1 | 1 | R2-ASSET-319 |
| 1 | Weapons | 4 | Long Sword | 1 | 3 | R2-ASSET-319 |
| 1 | Weapons | 6 | Two Handed Sword | 1 | 1 | R2-ASSET-319 |
| 1 | Weapons | 10 | Morning Star | 1 | 1 | R2-ASSET-319 |
| 1 | Weapons | 13 | Staff |  | 8 | R2-ASSET-319 |
| 1 | Weapons | 15 | Pike |  | 1 | R2-ASSET-319 |
| 1 | Weapons | 18 | Axe |  | 1 | R2-ASSET-319 |
| 1 | Weapons | 19 | Two Handed Axe |  | 1 | R2-ASSET-319 |
| 1 | Weapons | 21 | Long Bow | 1 | 3 | R2-ASSET-319 |
| 1 | Weapons | 22 | Crossbow |  | 2 | R2-ASSET-319 |
| 2 | Shields | 3 | Soft Small | 1 | 1 | R2-ASSET-319 |
| 2 | Shields | 5 | Large |  | 1 | R2-ASSET-319 |
| 2 | Shields | 6 | Soft Large | 3 | 4 | R2-ASSET-319 |
| 4 | Armors | 1 | Ring | 4 | 7 | R2-ASSET-319 |
| 5 | Armors | 2 | Amulet | 4 | 7 | R2-ASSET-319 |
| 6 | Armors | 3 | Hat | 1 | 2 | R2-ASSET-319 |
| 6 | Armors | 4 | Cap |  | 1 | R2-ASSET-319 |
| 6 | Armors | 5 | Low Hat | 2 | 2 | R2-ASSET-319 |
| 6 | Armors | 6 | Helm | 2 | 2 | R2-ASSET-319 |
| 6 | Armors | 7 | Soft Helm | 3 | 3 | R2-ASSET-319 |
| 6 | Armors | 9 | Full Helm |  | 1 | R2-ASSET-319 |
| 6 | Armors | 10 | Plate Helm | 2 | 4 | R2-ASSET-319 |
| 7 | Armors | 14 | Dress | 1 | 4 | R2-ASSET-319 |
| 7 | Armors | 15 | Soft Mail | 2 | 5 | R2-ASSET-319 |
| 7 | Armors | 16 | Chain Mail | 1 | 1 | R2-ASSET-319 |
| 7 | Armors | 17 | Scale Mail | 1 | 1 | R2-ASSET-319 |
| 8 | Armors | 11 | Cloak |  | 2 | R2-ASSET-319 |
| 8 | Armors | 19 | Plate Cuirass | 7 | 10 | R2-ASSET-319 |
| 9 | Armors | 22 | Plate Bracers | 3 | 4 | R2-ASSET-319 |
| 10 | Armors | 23 | Gloves | 2 | 3 | R2-ASSET-319 |
| 10 | Armors | 26 | Scale Gauntlets | 2 | 3 | R2-ASSET-319 |
| 12 | Armors | 27 | Shoes | 1 | 3 | R2-ASSET-319 |
| 12 | Armors | 30 | Plate Boots | 3 | 6 | R2-ASSET-319 |
| 14 | MagicItems | 1 | Book Air | 1 | 2 | R2-ASSET-319 |
| 14 | MagicItems | 2 | Book Water |  | 1 | R2-ASSET-319 |
| 14 | MagicItems | 3 | Book Fire | 3 | 3 | R2-ASSET-319 |
| 14 | MagicItems | 4 | Book Earth | 3 | 3 | R2-ASSET-319 |
| 14 | MagicItems | 5 | Book Astral |  | 2 | R2-ASSET-319 |
| 14 | MagicItems | 46 | SuperScroll Invisibility |  | 16 | R2-ASSET-319 |
| 14 | MagicItems | 47 | SuperScroll Protection from Air |  | 14 | R2-ASSET-319 |
| 14 | MagicItems | 53 | SuperScroll Protection from Earth |  | 55 | R2-ASSET-319 |
| 14 | MagicItems | 57 | SuperScroll Teleport |  | 46 | R2-ASSET-319 |
| 14 | MagicItems | 65 | Potion Body | 20 | 21 | R2-ASSET-319 |
| 14 | MagicItems | 66 | Potion Reaction | 15 | 16 | R2-ASSET-319 |
| 14 | MagicItems | 67 | Potion Mind | 17 | 18 | R2-ASSET-319 |
| 14 | MagicItems | 68 | Potion Spirit | 15 | 16 | R2-ASSET-319 |
| 14 | MagicItems | 77 | Quest Documents | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 78 | Quest Banner | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 79 | Quest Crown1 | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 80 | Quest Crown12 | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 81 | Quest Crown13 | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 82 | Quest Ingredient | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 83 | Quest Treasure | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 84 | Quest Amulet | 2 | 2 | R2-ASSET-319 |
| 14 | MagicItems | 85 | Quest Map | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 86 | Quest Head | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 92 | Quest RuneA | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 93 | Quest RuneE | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 94 | Quest RuneF | 1 | 1 | R2-ASSET-319 |
| 14 | MagicItems | 95 | Quest RuneW | 1 | 1 | R2-ASSET-319 |

## Kind 0x21 extension dwords

The two dwords of each 28-byte element (R2-ASSET-318). The loader replaces Buildings parameters 0 and 1 with them
when the sum of their low bytes is above 0 (R2-ENGINE-1289).

| First | Second | EN elements | RU elements | Claim |
|---:|---:|---:|---:|---|
| 2 | 2 | 3 | 3 | R2-ASSET-318 |
| 2 | 3 | 2 | 4 | R2-ASSET-318 |
| 2 | 4 |  | 2 | R2-ASSET-318 |
| 2 | 8 |  | 2 | R2-ASSET-318 |
| 3 | 2 |  | 3 | R2-ASSET-318 |
| 3 | 3 |  | 5 | R2-ASSET-318 |
| 3 | 4 | 4 | 5 | R2-ASSET-318 |
| 3 | 5 | 1 | 4 | R2-ASSET-318 |
| 3 | 6 | 1 | 3 | R2-ASSET-318 |
| 3 | 10 |  | 1 | R2-ASSET-318 |
| 3 | 12 |  | 1 | R2-ASSET-318 |
| 4 | 2 | 1 | 4 | R2-ASSET-318 |
| 4 | 4 | 1 | 9 | R2-ASSET-318 |
| 4 | 5 | 2 | 10 | R2-ASSET-318 |
| 4 | 6 |  | 7 | R2-ASSET-318 |
| 4 | 7 |  | 5 | R2-ASSET-318 |
| 4 | 8 |  | 2 | R2-ASSET-318 |
| 4 | 9 | 1 | 1 | R2-ASSET-318 |
| 4 | 11 |  | 1 | R2-ASSET-318 |
| 4 | 13 |  | 4 | R2-ASSET-318 |
| 4 | 16 |  | 2 | R2-ASSET-318 |
| 5 | 2 | 4 | 6 | R2-ASSET-318 |
| 5 | 3 |  | 14 | R2-ASSET-318 |
| 5 | 4 | 2 | 2 | R2-ASSET-318 |
| 5 | 5 | 1 | 1 | R2-ASSET-318 |
| 5 | 6 |  | 1 | R2-ASSET-318 |
| 5 | 7 |  | 1 | R2-ASSET-318 |
| 6 | 2 |  | 4 | R2-ASSET-318 |
| 6 | 3 |  | 15 | R2-ASSET-318 |
| 6 | 4 | 2 | 2 | R2-ASSET-318 |
| 7 | 2 | 4 | 5 | R2-ASSET-318 |
| 7 | 3 | 2 | 6 | R2-ASSET-318 |
| 8 | 2 | 4 | 5 | R2-ASSET-318 |
| 8 | 3 |  | 10 | R2-ASSET-318 |
| 8 | 4 | 1 | 1 | R2-ASSET-318 |
| 9 | 2 | 1 | 1 | R2-ASSET-318 |
| 9 | 3 | 1 | 6 | R2-ASSET-318 |
| 10 | 2 |  | 1 | R2-ASSET-318 |
| 10 | 3 | 2 | 5 | R2-ASSET-318 |
| 10 | 4 |  | 1 | R2-ASSET-318 |
| 11 | 3 | 1 | 2 | R2-ASSET-318 |
| 12 | 3 |  | 1 | R2-ASSET-318 |
| 16 | 2 |  | 1 | R2-ASSET-318 |
| 17 | 4 |  | 1 | R2-ASSET-318 |
