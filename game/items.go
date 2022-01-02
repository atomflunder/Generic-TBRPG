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

type Rarity struct {
	Common    float32
	Uncommon  float32
	Rare      float32
	Epic      float32
	Legendary float32
}

var (
	rarityCommon    = Rarity{1, 0, 0, 0, 0}
	rarityUncommon  = Rarity{0, 1, 0, 0, 0}
	rarityRare      = Rarity{0, 0, 1, 0, 0}
	rarityEpic      = Rarity{0, 0, 0, 1, 0}
	rarityLegendary = Rarity{0, 0, 0, 0, 1}
)

var SmallHealingPotion = Item{
	Name:        "Small Healing Potion",
	Description: "Heals you for 20 HP",
	Tag:         "Heal",
	Rarity:      rarityCommon,
	BuyPrice:    25,
	SellPrice:   1,
}

var LargeHealingPotion = Item{
	Name:        "Large Healing Potion",
	Description: "Heals you for 50 HP",
	Tag:         "Heal",
	Rarity:      rarityUncommon,
	BuyPrice:    120,
	SellPrice:   20,
}

var SmallBomb = Item{
	Name:        "Small Bomb",
	Description: "Damages your enemy for 20 HP",
	Tag:         "Damage",
	Rarity:      rarityCommon,
	BuyPrice:    25,
	SellPrice:   1,
}

var LargeBomb = Item{
	Name:        "Large Bomb",
	Description: "Damages your enemy for 50 HP",
	Tag:         "Damage",
	Rarity:      rarityUncommon,
	BuyPrice:    120,
	SellPrice:   20,
}

var AllItems = []Item{
	SmallHealingPotion, LargeHealingPotion, SmallBomb, LargeBomb,
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
	for {
		fmt.Println(PrintItems(GetAllItems(p)) + "\nWhich item do you want to use?")
		i := MatchItemIndex(utils.StringToInt(utils.GetUserInput()), GetAllItems(p))
		if i != nil {
			i.Use(p, e)
			break
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
