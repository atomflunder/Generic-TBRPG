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
func SwitchWeapon(c *Character, wl []Weapon) {
	fmt.Println("Do you want to switch your current weapon to one of these dropped ones? Type y to confirm, or anything else to dismiss.")
	if strings.ToLower(utils.GetUserInput()) == "y" {
		if len(wl) > 1 {
			for {
				fmt.Println("Which of these weapons do you wanna choose? Type their number.")
				w := MatchWeaponIndex(utils.StringToInt(utils.GetUserInput()), wl)
				if w != nil {
					c.Weapon = *w
					fmt.Println("Switched your current weapon to " + w.Name + ". You leave behind your old one.")
					break
				}
			}
		} else if len(wl) == 1 {
			c.Weapon = wl[0]
			fmt.Println("Switched your current weapon to " + wl[0].Name + ". You leave behind your old one.")
		}

	}
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
