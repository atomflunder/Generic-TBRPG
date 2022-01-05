package game

var ZombieKing = Monster{
	Name:           "★ Zombie King ★",
	Max_HP:         300,
	Current_HP:     300,
	LowAttack:      10,
	HighAttack:     25,
	AttackSpeed:    0.6,
	CritChance:     3.0,
	Accuracy:       85,
	Range:          6,
	XP_Min:         80,
	XP_Max:         105,
	Gold_Drop_Min:  90,
	Gold_Drop_Max:  130,
	Item_Drops:     Rarity{100, 90, 30, 0, 0},
	Weapon_Drops:   Rarity{0, 40, 0, 0, 0},
	Artifact_Drops: Rarity{25, 0, 0, 0, 0},
}

var SkeletonKing = Monster{
	Name:           "★ Skeleton King ★",
	Max_HP:         210,
	Current_HP:     210,
	LowAttack:      25,
	HighAttack:     50,
	AttackSpeed:    1.3,
	CritChance:     10.0,
	Accuracy:       95,
	Range:          8,
	XP_Min:         105,
	XP_Max:         130,
	Gold_Drop_Min:  120,
	Gold_Drop_Max:  160,
	Item_Drops:     Rarity{100, 100, 60, 0, 0},
	Weapon_Drops:   Rarity{0, 90, 0, 0, 0},
	Artifact_Drops: Rarity{0, 0, 75, 0, 0},
}

//gets the stages boss, or the zombie king as a default if the stage does not have a boss
func GetStageBoss(s Stage) Monster {
	switch s.Name {
	case "Foothills":
		return ZombieKing
	case "Old Bridge":
		return SkeletonKing
	}
	return ZombieKing
}
