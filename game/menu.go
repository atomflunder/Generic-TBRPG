package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main menu for the game
func MainMenu() {
	fmt.Println(`Welcome to the TBD game. What do you want to do?
1) Fight
2) Enter shop
3) Use Item
4) View Character
5) Switch character
6) View Leaderboard
7) Create a new character
8) Delete a character
...
9) Exit Game`)

	switch utils.GetUserInput() {
	case "1":
		c := GetDefaultCharacter(GetAllCharacters())
		if c == nil {
			fmt.Println("You have no characters saved. Please create one first.")
		} else {
			m := PickRandomMonster(AllMonsters)
			Combat(c, &m)
		}
	case "2":
		c := GetDefaultCharacter(GetAllCharacters())
		if c == nil {
			fmt.Println("You have no characters saved. Please create one first.")
		} else {
			ShopMenu(c)
		}
	case "3":
		c := GetDefaultCharacter(GetAllCharacters())
		if c == nil {
			fmt.Println("You have no characters saved. Please create one first.")
		} else {
			MenuItemChoice(c)
			c.Save()
		}
	case "4":
		c := GetDefaultCharacter(GetAllCharacters())
		if c == nil {
			fmt.Println("You have no characters saved. Please create one first.")
		} else {
			fmt.Println(c.Info())
		}

	case "5":
		SwitchCharacter(GetAllCharacters())
	case "6":
		fmt.Println(Leaderboard())
	case "7":
		c := CreateNewCharacter()
		fmt.Println(c.Info())
		c.Save()
	case "8":
		fmt.Println("Which character do you want to delete?")
		c := CharacterChoice()
		c.Delete()
	case "9":
		utils.ExitGame()
	default:
		fmt.Println("Invalid input, please try again.")
	}
}
