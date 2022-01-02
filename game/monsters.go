package game

import (
	"fmt"
	"math/rand"

	"github.com/phxenix-w/gotestgame/utils"
)

type Monster struct {
	Name           string
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
	Monster_Rarity Rarity
}

var Zombie = Monster{
	Name:           "Zombie",
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
	Monster_Rarity: rarityCommon,
}

var Skeleton = Monster{
	Name:           "Skeleton",
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
	Monster_Rarity: rarityCommon,
}

var Dwarf = Monster{
	Name:           "Dwarf",
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
	Monster_Rarity: rarityCommon,
}

var AllMonsters = []Monster{
	Zombie, Skeleton, Dwarf,
}

//refreshes the max hp of a monster, basically spawning a new one
func (m *Monster) RefreshHP() {
	m.Current_HP = m.Max_HP
}

//prints the info about a monster in a nice format
func (m *Monster) Info() string {
	return `Monster Info for ` + m.Name + `
Max HP: ` + fmt.Sprint(m.Max_HP) + `
Current HP: ` + fmt.Sprint(m.Current_HP) + `
Attack Values: ` + fmt.Sprint(m.LowAttack) + `-` + fmt.Sprint(m.HighAttack) + `
Attack Speed: ` + fmt.Sprint(m.AttackSpeed) + `
Crit Chance: ` + fmt.Sprint(m.CritChance) + `%
Range: ` + fmt.Sprint(m.Range) + `
XP Values: ` + fmt.Sprint(m.XP_Min) + `-` + fmt.Sprint(m.XP_Max)
}

//picks a random monster
func PickRandomMonster(ml []Monster) Monster {
	utils.GetNewRandomSeed()
	n := rand.Intn(len(ml))
	return ml[n]
}
