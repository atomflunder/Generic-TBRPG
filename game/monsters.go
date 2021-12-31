package game

import "fmt"

type Monster struct {
	Name        string
	Max_HP      int
	Current_HP  int
	LowAttack   int
	HighAttack  int
	AttackSpeed float32
	Range       int
	XP_Min      int
	XP_Max      int
	Item_Drops  Rarity
}

var Zombie = Monster{
	Name:        "Zombie",
	Max_HP:      50,
	Current_HP:  50,
	LowAttack:   5,
	HighAttack:  15,
	AttackSpeed: 0.5,
	Range:       5,
	XP_Min:      10,
	XP_Max:      15,
	Item_Drops:  Rarity{25, 5, 0, 0, 0},
}

//refreshes the max hp of a monster, basically spawning a new one
func RefreshMonsterHP(m *Monster) {
	m.Current_HP = m.Max_HP
}

//prints the info about a monster in a nice format
func MonsterInfo(m *Monster) string {
	return `Monster Info for ` + m.Name + `
Max HP: ` + fmt.Sprint(m.Max_HP) + `
Current HP: ` + fmt.Sprint(m.Current_HP) + `
Attack Values: ` + fmt.Sprint(m.LowAttack) + `-` + fmt.Sprint(m.HighAttack) + `
Attack Speed: ` + fmt.Sprint(m.AttackSpeed) + `
Range: ` + fmt.Sprint(m.Range) + `
XP Values: ` + fmt.Sprint(m.XP_Min) + `-` + fmt.Sprint(m.XP_Max)
}
