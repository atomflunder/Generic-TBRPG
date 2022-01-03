package game

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

type Item struct {
	Name        string
	Description string
	Tag         string
	Rarity      Rarity
	BuyPrice    int
	SellPrice   int
}

var SmallHealingPotion = Item{
	Name:        "Small Healing Potion",
	Description: "Heals you for 20 HP",
	Tag:         "Heal",
	Rarity:      RarityCommon,
	BuyPrice:    25,
	SellPrice:   1,
}

var LargeHealingPotion = Item{
	Name:        "Large Healing Potion",
	Description: "Heals you for 50 HP",
	Tag:         "Heal",
	Rarity:      RarityUncommon,
	BuyPrice:    120,
	SellPrice:   20,
}

var GiantHealingPotion = Item{
	Name:        "Giant Healing Potion",
	Description: "Heals you for 100 HP",
	Tag:         "Heal",
	Rarity:      RarityRare,
	BuyPrice:    350,
	SellPrice:   50,
}

var FullHealingPotion = Item{
	Name:        "Full Healing Potion",
	Description: "Heals you fully",
	Tag:         "Heal",
	Rarity:      RarityEpic,
	BuyPrice:    1000,
	SellPrice:   150,
}

var SmallBomb = Item{
	Name:        "Small Bomb",
	Description: "Damages your enemy for 20 HP",
	Tag:         "Damage",
	Rarity:      RarityCommon,
	BuyPrice:    25,
	SellPrice:   1,
}

var LargeBomb = Item{
	Name:        "Large Bomb",
	Description: "Damages your enemy for 50 HP",
	Tag:         "Damage",
	Rarity:      RarityUncommon,
	BuyPrice:    120,
	SellPrice:   20,
}

var AllItems = []Item{
	SmallHealingPotion, LargeHealingPotion, GiantHealingPotion, FullHealingPotion, SmallBomb, LargeBomb,
}

//adds an item to your character
func (i Item) Add(c *Character) {
	c.Items = append(c.Items, i)
}

//adds multiple items to your character
func AddItems(c *Character, il []Item) {
	c.Items = append(c.Items, il...)
}

//gets the item index, for deletion
func (i Item) GetIndex(c *Character) int {
	for x, r := range c.Items {
		if r.Name == i.Name {
			return x
		}
	}
	return -1
}

//matches the item index to the input
func MatchItemIndex(p int, il []Item) *Item {
	for x := range il {
		if x+1 == p {
			return &il[x]
		}
	}
	return nil
}

//chooses an item for the player to use, used in the combat loop
func ItemChoice(p *Character, e *Monster) {
	if len(GetAllItems(p)) == 0 {
		fmt.Println("You do not have any items in your inventory.")
		p.Turn(e)
		return
	}

	for {
		fmt.Println(PrintItems(GetAllItems(p)) + "\nWhich item do you want to use?")
		i := MatchItemIndex(utils.StringToInt(utils.GetUserInput()), GetAllItems(p))
		if i != nil {
			i.Use(p, e)
			break
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

//chooses an item for the player to use in the menu
func MenuItemChoice(c *Character) {
	if len(GetAllItems(c)) == 0 {
		fmt.Println("You do not have any items in your inventory.")
		return
	}

	fmt.Println(PrintItems(GetAllItems(c)) + "\nWhich item do you want to use? Or type exit to exit.")

	inp := utils.GetUserInput()

	if strings.ToLower(inp) == "exit" {
		return
	} else {
		i := MatchItemIndex(utils.StringToInt(inp), GetAllItems(c))

		if i != nil {
			if i.Tag == "Heal" {
				HealingItem(c, *i)
				i.Remove(c)

			} else {
				fmt.Println("You cannot use this item here!")

			}
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}

}

//removes an item
func (i Item) Remove(c *Character) {
	ind := i.GetIndex(c)
	if ind != -1 {
		c.Items = append(c.Items[:ind], c.Items[ind+1:]...)
	}

}

//returns all items a character has
func GetAllItems(c *Character) []Item {
	var itemList []Item
	return append(itemList, c.Items...)
}

//gets you every item by the specified rarity
func GetItemsByRarity(r Rarity) []Item {
	var itemList []Item

	for _, i := range AllItems {
		if i.Rarity == r {
			itemList = append(itemList, i)
		}
	}
	return itemList
}

//picks a random item from the list of items
func PickRandomItem(il []Item) Item {
	utils.GetNewRandomSeed()
	n := rand.Intn(len(il))
	return il[n]
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
func (i Item) Use(c *Character, e *Monster) {
	switch i.Tag {
	case "Heal":
		HealingItem(c, i)
	case "Damage":
		DamageItem(c, e, i)
	}
	i.Remove(c)

}

//the logic for healing items
func HealingItem(c *Character, i Item) {
	switch i.Name {
	case "Small Healing Potion":
		c.Heal(20)
	case "Large Healing Potion":
		c.Heal(50)
	case "Giant Healing Potion":
		c.Heal(100)
	case "Full Healing Potion":
		c.Heal(-1)
	}

}

//the logic for damage items
func DamageItem(c *Character, m *Monster, i Item) {
	switch i.Name {
	case "Small Bomb":
		m.ApplyItemDamage(20, i)
	case "Large Bomb":
		m.ApplyItemDamage(50, i)
	}
}
