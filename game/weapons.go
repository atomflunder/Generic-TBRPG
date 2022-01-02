package game

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

type Weapon struct {
	Name        string
	Description string
	Rarity      Rarity
	LowAttack   int
	HighAttack  int
	AttackSpeed float32
	CritChance  float32
	Accuracy    int
	Range       int
	ReqStr      int
	ReqDex      int
	ReqInt      int
	BuyPrice    int
	SellPrice   int
}

var RustyClub = Weapon{
	Name:        "Rusty Club",
	Description: "Starter Weapon for the Barbarian.",
	Rarity:      rarityCommon,
	LowAttack:   15,
	HighAttack:  40,
	AttackSpeed: 1.2,
	CritChance:  5.0,
	Accuracy:    95,
	Range:       10,
	ReqStr:      15,
	ReqDex:      5,
	ReqInt:      0,
	BuyPrice:    30,
	SellPrice:   1,
}

var RustyDagger = Weapon{
	Name:        "Rusty Dagger",
	Description: "Starter Weapon for the Rogue",
	Rarity:      rarityCommon,
	LowAttack:   5,
	HighAttack:  30,
	AttackSpeed: 1.7,
	CritChance:  15.0,
	Accuracy:    85,
	Range:       6,
	ReqStr:      5,
	ReqDex:      10,
	ReqInt:      5,
	BuyPrice:    30,
	SellPrice:   1,
}

var RustyStaff = Weapon{
	Name:        "Rusty Staff",
	Description: "Starter Weapon for the Mage",
	Rarity:      rarityCommon,
	LowAttack:   10,
	HighAttack:  35,
	AttackSpeed: 1.4,
	CritChance:  6.0,
	Accuracy:    95,
	Range:       14,
	ReqStr:      5,
	ReqDex:      5,
	ReqInt:      10,
	BuyPrice:    30,
	SellPrice:   1,
}

var Broadsword = Weapon{
	Name:        "Broadsword",
	Description: "A two-handed sword, very sharp and dangerous",
	Rarity:      rarityUncommon,
	LowAttack:   25,
	HighAttack:  50,
	AttackSpeed: 1.5,
	CritChance:  10.0,
	Accuracy:    95,
	Range:       12,
	ReqStr:      15,
	ReqDex:      15,
	ReqInt:      5,
	BuyPrice:    350,
	SellPrice:   50,
}

var AllWeapons = []Weapon{
	RustyClub, RustyDagger, RustyStaff, Broadsword,
}

//picks a random weapon from the list
func PickRandomWeapon(wl []Weapon) Weapon {
	utils.GetNewRandomSeed()
	n := rand.Intn(len(wl))
	return wl[n]
}

//gets every weapon with a specified rarity
func GetWeaponsByRarity(r Rarity) []Weapon {
	var weaponList []Weapon

	for _, w := range AllWeapons {
		if w.Rarity == r {
			weaponList = append(weaponList, w)
		}
	}

	return weaponList
}

//prints weapon drops in a nice readable format
func PrintWeapons(wl []Weapon) string {
	var nameList []string
	for i, w := range wl {
		nameList = append(nameList, fmt.Sprint(i+1)+") "+w.Name+"("+fmt.Sprint(w.LowAttack)+"-"+fmt.Sprint(w.HighAttack)+")")
	}
	return strings.Join(nameList, ", ")
}

//asks the player if they want to switch out weapons
func SwitchWeaponFromList(c *Character, wl []Weapon) {
	if len(wl) > 1 {
		fmt.Println("Do you want to switch your current weapon to one of these dropped ones? Type y to confirm or anything else to dismiss.")
		if strings.ToLower(utils.GetUserInput()) == "y" {
			for {
				fmt.Println("Which of these weapons do you wanna choose? Type their number.")
				w := MatchWeaponIndex(utils.StringToInt(utils.GetUserInput()), wl)
				if w != nil {
					w.Switch(c)
					break
				}
			}
		}
	} else if len(wl) == 1 {
		wl[0].Switch(c)
	}
}

//asks the player if they want to switch out their weapon for a specific one, used in the shop
func (w Weapon) Switch(c *Character) {
	pass, missingStats := w.RequirementCheck(c)

	if pass {
		fmt.Println(w.Info())
		fmt.Println("Are you sure you want to switch your current weapon to this one? Type y to confirm, or anything else to dismiss.")
		if strings.ToLower(utils.GetUserInput()) == "y" {
			c.Weapon = w
			fmt.Println("Switched your current weapon to " + w.Name + ". You leave behind your old one.")
		}
	} else {
		fmt.Println("Sorry, you cant wield this weapon!\n" + missingStats)
	}

}

func (w Weapon) Info() string {
	return `` + w.Name + `
` + w.Description + `
Damage: 		` + fmt.Sprint(w.LowAttack) + `-` + fmt.Sprint(w.HighAttack) + `
Crit Chance: 		` + fmt.Sprint(w.CritChance) + `%
Accuracy: 		` + fmt.Sprint(w.Accuracy) + `
Range: 			` + fmt.Sprint(w.Range) + `
Strength Required: 	` + fmt.Sprint(w.ReqStr) + `
Dexterity Required: 	` + fmt.Sprint(w.ReqDex) + `
Intelligence Required:	` + fmt.Sprint(w.ReqInt)
}

//matches the item index to the input
func MatchWeaponIndex(p int, wl []Weapon) *Weapon {
	for x := range wl {
		if x+1 == p {
			return &wl[x]
		}
	}
	return nil
}

//determines if you meet the weapon stat requirements
func (w Weapon) RequirementCheck(c *Character) (bool, string) {
	var missingAttributes string

	if c.Strength < w.ReqStr {
		missingAttributes = missingAttributes + "Your Strength: " + fmt.Sprint(c.Strength) + "\nNeeded Strength: " + fmt.Sprint(w.ReqStr) + "\n"
	}
	if c.Dexterity < w.ReqDex {
		missingAttributes = missingAttributes + "Your Dexterity: " + fmt.Sprint(c.Dexterity) + "\nNeeded Dexterity: " + fmt.Sprint(w.ReqDex) + "\n"
	}

	if c.Intelligence < w.ReqInt {
		missingAttributes = missingAttributes + "Your Intelligence: " + fmt.Sprint(c.Intelligence) + "\nNeeded Intelligence: " + fmt.Sprint(w.ReqInt) + "\n"
	}

	return (c.Strength >= w.ReqStr) && (c.Dexterity >= w.ReqDex) && (c.Intelligence >= w.ReqInt), missingAttributes
}
