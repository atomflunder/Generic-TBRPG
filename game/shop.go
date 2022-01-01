package game

import (
	"fmt"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main menu for the shop
func ShopMenu(c *Character) {
	fmt.Println(`Welcome to the shop! Do you want to browse our items, or shop for a new weapon?
1) Buy Item
2) Sell Item
3) Buy Weapon
...
9) Leave Shop`)

	switch utils.GetUserInput() {
	case "1":
		BuyItem(c)
	case "2":
		SellItem(c)
	case "3":
		BuyWeapon(c)
	case "9":
		MainMenu()

	default:
		fmt.Println("Invalid input. Please try again.")
		ShopMenu(c)
	}
}

//buys an item, then adds it to your inventory
func BuyItem(c *Character) {
	fmt.Println("These are the items we have to offer. \n" + PrintItems(AllItems) + "\n\nType the number of the item you want to buy.")
	i := MatchItemIndex(utils.StringToInt(utils.GetUserInput()), AllItems)
	if i != nil {
		if CanCharacterBuyItem(c, *i) {
			fmt.Println("Your current balance: " + fmt.Sprint(c.Gold) + "\nDo you want to buy a " + i.Name + " for " + fmt.Sprint(i.BuyPrice) + " gold? Type y to confirm, or anything else to dismiss.")
			if strings.ToLower(utils.GetUserInput()) == "y" {
				AddItem(c, *i)
				c.Gold -= i.BuyPrice
				fmt.Println("You have bought a " + i.Name + " for " + fmt.Sprint(i.BuyPrice) + " gold. You have " + fmt.Sprint(c.Gold) + " gold left.")
				SaveCharacter(c)
			} else {
				ShopMenu(c)
			}

		} else {
			fmt.Println("Sorry, you only have " + fmt.Sprint(c.Gold) + " gold but a " + i.Name + " costs " + fmt.Sprint(i.BuyPrice) + " gold.")
			ShopMenu(c)
		}
	} else {
		fmt.Println("Invalid input. Please try again.")
	}
}

//sells an item from your inventory
func SellItem(c *Character) {

}

//buys a weapon, then asks you to swap weapons
func BuyWeapon(c *Character) {

}

//not yet implemented. you can only carry your current weapon so you cant really sell it yet.
func SellWeapon() {

}

//checks if the character has enough gold to buy an item
func CanCharacterBuyItem(c *Character, i Item) bool {
	return c.Gold >= i.BuyPrice
}

//checks if the character has enough gold to buy a weapon
func CanCharacterBuyWeapon(c *Character, w Weapon) bool {
	return c.Gold >= w.BuyPrice
}
