package game

import (
	"fmt"
	"math/rand"

	"github.com/phxenix-w/gotestgame/utils"
)

type Monster struct {
	Name           string
	Level          int
	Max_HP         int
	Current_HP     int
	LowAttack      int
	HighAttack     int
	AttackSpeed    float32
	CritChance     float32
	Accuracy       int
	Range          int
	XP_Min         int
	XP_Max         int
	Gold_Drop_Min  int
	Gold_Drop_Max  int
	Item_Drops     Rarity
	Weapon_Drops   Rarity
	Artifact_Drops Rarity
	Monster_Rarity Rarity
}

var Zombie = Monster{
	Name:           "Zombie",
	Level:          1,
	Max_HP:         50,
	Current_HP:     50,
	LowAttack:      5,
	HighAttack:     15,
	AttackSpeed:    0.5,
	CritChance:     3.0,
	Accuracy:       80,
	Range:          5,
	XP_Min:         10,
	XP_Max:         15,
	Gold_Drop_Min:  5,
	Gold_Drop_Max:  10,
	Item_Drops:     Rarity{25, 5, 0, 0, 0},
	Weapon_Drops:   Rarity{5, 0, 0, 0, 0},
	Artifact_Drops: Rarity{0, 0, 0, 0, 0},
	Monster_Rarity: RarityCommon,
}

var Skeleton = Monster{
	Name:           "Skeleton",
	Level:          1,
	Max_HP:         60,
	Current_HP:     60,
	LowAttack:      10,
	HighAttack:     20,
	AttackSpeed:    1.0,
	CritChance:     8.5,
	Accuracy:       90,
	Range:          6,
	XP_Min:         20,
	XP_Max:         25,
	Gold_Drop_Min:  10,
	Gold_Drop_Max:  15,
	Item_Drops:     Rarity{30, 10, 5, 0, 0},
	Weapon_Drops:   Rarity{10, 2, 0, 0, 0},
	Artifact_Drops: Rarity{0, 0, 1, 0, 0},
	Monster_Rarity: RarityUncommon,
}

var Dwarf = Monster{
	Name:           "Dwarf",
	Level:          1,
	Max_HP:         30,
	Current_HP:     30,
	LowAttack:      25,
	HighAttack:     45,
	AttackSpeed:    1.5,
	CritChance:     12.5,
	Accuracy:       99,
	Range:          3,
	XP_Min:         10,
	XP_Max:         15,
	Gold_Drop_Min:  30,
	Gold_Drop_Max:  40,
	Item_Drops:     Rarity{10, 2, 0, 0, 0},
	Weapon_Drops:   Rarity{0, 0, 0, 0, 0},
	Artifact_Drops: Rarity{0, 0, 3, 0, 0},
	Monster_Rarity: RarityRare,
}

var Goblin = Monster{
	Name:           "Goblin",
	Level:          1,
	Max_HP:         70,
	Current_HP:     70,
	LowAttack:      20,
	HighAttack:     25,
	AttackSpeed:    2.5,
	CritChance:     15,
	Accuracy:       85,
	Range:          2,
	XP_Min:         15,
	XP_Max:         25,
	Gold_Drop_Min:  150,
	Gold_Drop_Max:  250,
	Item_Drops:     Rarity{0, 25, 0, 0, 0},
	Weapon_Drops:   Rarity{0, 0, 0, 0, 0},
	Artifact_Drops: Rarity{0, 0, 10, 0, 0},
	Monster_Rarity: RarityEpic,
}

var ZombieKing = Monster{
	Name:           "★ Zombie King ★",
	Level:          1,
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
	Monster_Rarity: RarityLegendary,
}

var SkeletonKing = Monster{
	Name:           "★ Skeleton King ★",
	Level:          3,
	Max_HP:         310,
	Current_HP:     310,
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
	Monster_Rarity: RarityLegendary,
}

//refreshes the max hp of a monster, basically spawning a new one
func (m *Monster) RefreshHP() {
	m.Current_HP = m.Max_HP
}

//prints the info about a monster in a nice format
func (m *Monster) Info() string {
	return `Monster Info for ` + m.Name + `
Level: ` + fmt.Sprint(m.Level) + `
Max HP: ` + fmt.Sprint(m.Max_HP) + `
Current HP: ` + fmt.Sprint(m.Current_HP) + `
Attack Values: ` + fmt.Sprint(m.LowAttack) + `-` + fmt.Sprint(m.HighAttack) + `
Attack Speed: ` + fmt.Sprint(m.AttackSpeed) + `
Crit Chance: ` + fmt.Sprint(m.CritChance) + `%
Range: ` + fmt.Sprint(m.Range) + `
XP Values: ` + fmt.Sprint(m.XP_Min) + `-` + fmt.Sprint(m.XP_Max)
}

//picks a random monster to fight against
func GetMonster(s Stage) Monster {
	var monsterList []Monster
	var r Rarity

	n := utils.GetRandomNumber(100)

	if n <= 5 {
		r = RarityLegendary
	} else if n <= 10 {
		r = RarityEpic
	} else if n <= 25 {
		r = RarityRare
	} else if n <= 50 {
		r = RarityUncommon
	} else {
		r = RarityCommon
	}

	for _, m := range s.Monsters {
		if m.Monster_Rarity == r {
			monsterList = append(monsterList, m)
		}
	}

	return PickRandomMonster(monsterList)
}

//picks a random monster
func PickRandomMonster(ml []Monster) Monster {
	utils.GetNewRandomSeed()
	n := rand.Intn(len(ml))
	return ml[n]
}
