package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main menu for the game
func MainMenu() {
	fmt.Println(`Welcome to the TBD game. What do you want to do?
1) Fight with saved character
2) Enter shop with saved character
3) View Character
4) Switch character
5) Create a new character
6) Delete a character
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
			fmt.Println(CharacterInfo(c))
		}

	case "4":
		SwitchCharacter(GetAllCharacters())
	case "5":
		c := CreateNewCharacter()
		fmt.Println(CharacterInfo(c))
		SaveCharacter(c)
	case "6":
		DeleteCharacter()
	case "9":
		utils.ExitGame()
	default:
		fmt.Println("Invalid input, please try again.")
	}
}
