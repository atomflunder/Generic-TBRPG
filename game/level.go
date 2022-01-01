package game

import "fmt"

var (
	Level1  = 1
	Level2  = 50
	Level3  = 120
	Level4  = 300
	Level5  = 450
	Level6  = 640
	Level7  = 910
	Level8  = 1200
	Level9  = 1640
	Level10 = 2200
	Level11 = 2820
	Level12 = 3400
	Level13 = 4030
	Level14 = 4800
	Level15 = 5750
	Level16 = 6750
	Level17 = 7890
	Level18 = 9030
	Level19 = 10800
	Level20 = 13000
)

var AllLevels = [20]int{Level1, Level2, Level3, Level4, Level5, Level6, Level7, Level8, Level9,
	Level10, Level11, Level12, Level13, Level14, Level15, Level16, Level17, Level18, Level19, Level20}

//checks a character level
func CheckLevel(c *Character) int {
	var charLevel int

	for _, y := range AllLevels {
		if c.XP >= y {
			charLevel += 1
		}
	}

	return charLevel
}

//updates a character level
func UpdateLevel(c *Character) {
	l := CheckLevel(c)
	if c.Level != l {
		c.Level = l
		LevelUp(c)
		fmt.Println("Congratulations, you have leveled up to level " + fmt.Sprint(l) + ". Your stats have increased and you healed to full health.")
	}
}

//levels up a characters stats
func LevelUp(c *Character) {
	c.Max_HP += 10
	c.Current_HP = c.Max_HP
	switch c.Class {
	case "Barbarian":
		c.Strength += 10
		c.Dexterity += 5
		c.Intelligence += 2
	case "Rogue":
		c.Strength += 2
		c.Dexterity += 10
		c.Intelligence += 5
	case "Mage":
		c.Strength += 2
		c.Dexterity += 5
		c.Intelligence += 10
	}
}

//gets a 10% xp penalty of the current level, down to the minimum. returns the lost xp
func ApplyXPPenalty(c *Character) int {
	l := CheckLevel(c)
	min := AllLevels[l-1]

	levelRange := AllLevels[l] - AllLevels[l-1]
	penalty := levelRange / 10

	var lostXP int

	if (c.XP - penalty) <= min {
		lostXP = c.XP - min
		c.XP = min
	} else {
		lostXP = penalty
		c.XP -= penalty
	}
	return lostXP

}
