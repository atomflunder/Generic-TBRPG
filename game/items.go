package game

import (
	"fmt"
	"strings"

	"github.com/phxenix-w/placeholder-name/utils"
)

type Item struct {
	Name        string
	Description string
	Tag         string
}

var SmallHealingPotion = Item{
	Name:        "Small Healing Potion",
	Description: "Heals you for 20 HP",
	Tag:         "Heal",
}

var LargeHealingPotion = Item{
	Name:        "Large Healing Potion",
	Description: "Heals you for 50 HP",
	Tag:         "Heal",
}

var SmallBomb = Item{
	Name:        "Small Bomb",
	Description: "Damages your enemy for 20 HP",
	Tag:         "Damage",
}

var LargeBomb = Item{
	Name:        "Large Bomb",
	Description: "Damages your enemy for 50 HP",
	Tag:         "Damage",
}

//adds an item to your character
func AddItem(c *Character, i Item) {
	c.Items = append(c.Items, i)
}

//gets the item index, for deletion
func GetItemIndex(c *Character, i Item) int {
	for x, r := range c.Items {
		if r.Name == i.Name {
			return x
		}
	}
	return -1
}

//matches the item index to the input
func MatchItemIndex(p int, c *Character) *Item {
	il := GetAllItems(c)
	for x := range il {
		if x+1 == p {
			return &il[x]
		}
	}
	return nil
}

//chooses an item for the player to use, used in the combat loop
func ItemChoice(p *Character, e *Monster) {
	for {
		fmt.Println(PrintItems(GetAllItems(p)) + "\nWhich item do you want to use?")
		i := MatchItemIndex(utils.StringToInt(utils.GetUserInput()), p)
		if i != nil {
			UseItem(p, e, *i)
			break
		}
	}
}

//removes an item
func RemoveItem(c *Character, i Item) {
	ind := GetItemIndex(c, i)
	if ind != -1 {
		c.Items = append(c.Items[:ind], c.Items[ind+1:]...)
	}

}

//returns all items a character has
func GetAllItems(c *Character) []Item {
	var itemList []Item
	return append(itemList, c.Items...)
}

//prints them out in a nice format
func PrintItems(i []Item) string {
	var nameList []string
	for y, x := range i {
		nameList = append(nameList, fmt.Sprint(y+1)+") "+x.Name)
	}
	return strings.Join(nameList, ", ")
}

//uses an item
func UseItem(c *Character, e *Monster, i Item) {
	switch i.Tag {
	case "Heal":
		HealingItem(c, i)
	case "Damage":
		DamageItem(c, e, i)
	}
	RemoveItem(c, i)

}

//the logic for healing items
func HealingItem(c *Character, i Item) {
	switch i {
	case SmallHealingPotion:
		HealPlayer(20, c)
	case LargeHealingPotion:
		HealPlayer(50, c)
	}
}

//the logic for damage items
func DamageItem(c *Character, e *Monster, i Item) {
	switch i {
	case SmallBomb:
		ApplyDamageToEnemy(20, e)
	case LargeBomb:
		ApplyDamageToEnemy(50, e)
	}
}
