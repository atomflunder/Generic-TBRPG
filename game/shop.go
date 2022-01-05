package game

import (
	"fmt"
	"strings"

	"github.com/phxenix-w/Generic-TBRPG/utils"
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
		if c.CanBuyItem(*i) {
			fmt.Println(`Your current balance: ` + fmt.Sprint(c.Gold) + `
Do you want to buy a ` + i.Name + ` (` + i.Description + `) for ` + fmt.Sprint(i.BuyPrice) + ` gold?
Type y to confirm, or anything else to dismiss`)

			if strings.ToLower(utils.GetUserInput()) == "y" {
				i.Add(c)
				c.Gold -= i.BuyPrice
				fmt.Println("You have bought a " + i.Name + " for " + fmt.Sprint(i.BuyPrice) + " gold. You have " + fmt.Sprint(c.Gold) + " gold left.")
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
	c.Save()
}

//sells an item from your inventory
func SellItem(c *Character) {
	fmt.Println("What item do you want to sell? \n" + PrintItems(GetAllItems(c)) + "\n\nType the number of the item you want to sell.")
	i := MatchItemIndex(utils.StringToInt(utils.GetUserInput()), GetAllItems(c))
	if i != nil {
		fmt.Println(`Your current balance: ` + fmt.Sprint(c.Gold) + `
Do you want to sell your ` + i.Name + ` (` + i.Description + `) for ` + fmt.Sprint(i.SellPrice) + ` gold?
Type y to confirm, or anything else to dismiss`)

		if strings.ToLower(utils.GetUserInput()) == "y" {
			i.Remove(c)
			c.Gold += i.SellPrice
			fmt.Println("You have sold a " + i.Name + " for " + fmt.Sprint(i.SellPrice) + " gold. You now have " + fmt.Sprint(c.Gold) + " gold.")
		} else {
			ShopMenu(c)
		}

	} else {
		fmt.Println("Invalid input. Please try again.")
	}

	c.Save()
}

//buys a weapon, then asks you to swap weapons
func BuyWeapon(c *Character) {
	fmt.Println("These are the weapons we have to offer. \n" + PrintWeapons(AllWeapons) + "\n\nType the number of the item you want to buy.")
	w := MatchWeaponIndex(utils.StringToInt(utils.GetUserInput()), AllWeapons)
	if w != nil {
		if c.CanBuyWeapon(*w) {
			fmt.Println(`Your current balance: ` + fmt.Sprint(c.Gold) + `
Do you want to buy a this weapon for ` + fmt.Sprint(w.BuyPrice) + ` gold?
			
Weapon details:
` + w.Info() + `
			
Type y to confirm, or anything else to dismiss`)

			if strings.ToLower(utils.GetUserInput()) == "y" {
				pass, missingStats := w.RequirementCheck(c)
				if pass {
					w.Switch(c)
					c.Gold -= w.BuyPrice
					fmt.Println("You bought a " + w.Name + " for " + fmt.Sprint(w.BuyPrice) + " gold. You have " + fmt.Sprint(c.Gold) + " gold left.")
				} else {
					fmt.Println("Sorry, you cant wield this weapon!\n" + missingStats)
				}

			} else {
				ShopMenu(c)
			}
		} else {
			fmt.Println("Sorry, you only have " + fmt.Sprint(c.Gold) + " gold but a " + w.Name + " costs " + fmt.Sprint(w.BuyPrice) + " gold.")
			ShopMenu(c)
		}

	} else {
		fmt.Println("Invalid input. Please try again.")
	}

	c.Save()
}

//not yet implemented. you can only carry your current weapon so you cant really sell it yet.
func SellWeapon(c *Character) {

	c.Save()
}

//checks if the character has enough gold to buy an item
func (c *Character) CanBuyItem(i Item) bool {
	return c.Gold >= i.BuyPrice
}

//checks if the character has enough gold to buy a weapon
func (c *Character) CanBuyWeapon(w Weapon) bool {
	return c.Gold >= w.BuyPrice
}
