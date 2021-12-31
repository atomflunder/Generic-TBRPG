package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main menu for the game
func MainMenu() {
	fmt.Println(`Welcome to the TBD game. What do you want to do?
1) Continue with saved character
2) View Character
3) Create a new character
4) Delete a character
...
9) Exit Game`)

	switch utils.GetUserInput() {
	case "1":
		c := CharacterChoice()
		if c == nil {
			fmt.Println("Invalid input. Please try again.")
		} else {
			m := PickRandomMonster(AllMonsters)
			Combat(c, &m)
		}
	case "2":
		c := CharacterChoice()
		if c == nil {
			fmt.Println("Invalid input. Please try again.")
		} else {
			fmt.Println(CharacterInfo(c))
		}

	case "3":
		c := CreateNewCharacter()
		fmt.Println(CharacterInfo(c))
		SaveCharacter(c)
	case "4":
		fmt.Println("Which character do you want to delete?")
		cl := CharacterListToString(GetAllCharacters())
		if len(cl) == 0 {
			fmt.Println("No characters found.")
		} else {
			fmt.Println(cl)
			DeleteCharacter(utils.GetUserInput())
		}
	case "9":
		utils.ExitGame()
	default:
		fmt.Println("Invalid input, please try again.")
	}
}
