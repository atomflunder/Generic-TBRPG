package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

type Character struct {
	Name         string `json:"name"`
	Hardcore     bool   `json:"hardcore"`
	Default      bool   `json:"default"`
	Level        int    `json:"level"`
	XP           int    `json:"xp"`
	Class        string `json:"class"`
	Max_HP       int    `json:"max_hp"`
	Current_HP   int    `json:"current_hp"`
	Strength     int    `json:"str"`
	Dexterity    int    `json:"dex"`
	Intelligence int    `json:"int"`
	Gold         int    `json:"gold"`
	Weapon       Weapon `json:"weapon"`
	Items        []Item `json:"items"`
}

//creates a new character and asks you about the details
func CreateNewCharacter() *Character {
	c := &Character{}

	fmt.Println(`Welcome to the Character Creation Tool!
Please enter the name of your new character!`)

	c.Name = utils.GetUserInput()

	c.Level = 1
	c.XP = 1
	c.Gold = 0
	AddItem(c, SmallHealingPotion)
	c.Hardcore = true

	for c.Class == "" {
		fmt.Println(`Which class do you want to start with?
1) Barbarian
2) Rogue
3) Mage`)

		switch utils.GetUserInput() {
		case "1":
			c.Class = "Barbarian"
			c.Max_HP = 150
			c.Current_HP = c.Max_HP
			c.Strength = 30
			c.Dexterity = 15
			c.Intelligence = 5
			c.Weapon = RustyClub
		case "2":
			c.Class = "Rogue"
			c.Max_HP = 95
			c.Current_HP = c.Max_HP
			c.Strength = 15
			c.Dexterity = 25
			c.Intelligence = 15
			c.Weapon = RustyDagger
		case "3":
			c.Class = "Mage"
			c.Max_HP = 70
			c.Current_HP = c.Max_HP
			c.Strength = 10
			c.Dexterity = 15
			c.Intelligence = 25
			c.Weapon = RustyStaff
		default:
			fmt.Println("Invalid input. Please try again.")

		}
		o := 0
		for o < 1 {
			fmt.Println("Do you want to enable hardcore mode? (Permadeath)\ny/n?")
			switch utils.GetUserInput() {
			case "y":
				c.Hardcore = true
				o += 1
			case "n":
				c.Hardcore = false
				o += 1

			default:
				fmt.Println("Invalid input. Please try again.")
			}
		}

	}
	SwitchAllCharactersOff(GetAllCharacters())
	c.Default = true
	SaveCharacter(c)

	return c
}

//gets you all relevant character information in one readable string
func CharacterInfo(c *Character) string {
	//not sure why it is out of line here, in the command line it lines up
	return `Character Info for ` + c.Name + `!
Level: 			` + fmt.Sprint(c.Level) + `
XP: 			` + fmt.Sprint(c.XP) + `
Class: 			` + c.Class + `
Max HP: 		` + fmt.Sprint(c.Max_HP) + `
Current HP: 		` + fmt.Sprint(c.Current_HP) + `
Strength: 		` + fmt.Sprint(c.Strength) + `
Dexterity: 		` + fmt.Sprint(c.Dexterity) + `
Intelligence: 		` + fmt.Sprint(c.Intelligence) + `
Gold: 			` + fmt.Sprint(c.Gold) + `
Weapon: 		` + c.Weapon.Name + `
Items: 			` + PrintItems(GetAllItems(c))
}

//reads all saved characters into a readable string
func CharacterListToString(cl []Character) string {
	var charList string
	for _, c := range cl {
		if c.Default {
			charList = charList + "Name: " + c.Name + " (Currently active)\n" + "Level: " + fmt.Sprint(c.Level) + "\n\n"
		} else {
			charList = charList + "Name: " + c.Name + "\n" + "Level: " + fmt.Sprint(c.Level) + "\n\n"
		}
	}

	return charList
}

//saves the character to a json file with the same name
func SaveCharacter(c *Character) {
	file, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./savedata/characters/"+c.Name+".json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//gets all characters saved
func GetAllCharacters() []Character {
	var charList []Character

	files, err := ioutil.ReadDir("./savedata/characters/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			file, err := ioutil.ReadFile("./savedata/characters/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			var char Character

			err = json.Unmarshal(file, &char)
			if err != nil {
				log.Fatal(err)
			}

			charList = append(charList, char)

		}
	}
	return charList
}

//searches for a specific character
func SearchCharacter(s string, cl []Character) *Character {
	for _, c := range cl {
		if s == c.Name {
			return &c
		}
	}
	return nil
}

//selects your character directly, used in the main menu
func CharacterChoice() *Character {
	fmt.Println("Choose your character:")
	cl := CharacterListToString(GetAllCharacters())
	if len(cl) == 0 {
		fmt.Println("No characters found.")
		return nil
	} else {
		fmt.Println(cl)
		fmt.Println("Please enter the name of your character.")
		c := SearchCharacter(utils.GetUserInput(), GetAllCharacters())
		return c
	}
}

//gets the character with the default flag enabled. if it doesnt find one, it returns the first one saved
func GetDefaultCharacter(cl []Character) *Character {
	for _, c := range cl {
		if c.Default {
			return &c
		}
	}
	if len(cl) == 0 {
		return nil
	} else {
		return &cl[0]
	}
}

//switches your current active character
func SwitchCharacter(cl []Character) {
	c := CharacterChoice()
	SwitchAllCharactersOff(cl)
	c.Default = true
	SaveCharacter(c)

	fmt.Println("Set " + c.Name + " to your default character.")
}

//if you create a new character, it will have the default flag enabled, so we need to disable it for everyone else
func SwitchAllCharactersOff(cl []Character) {
	for _, c := range cl {
		c.Default = false
		SaveCharacter(&c)
	}
}

//what happens when a character dies
func CharacterDeath(c *Character) {
	if c.Hardcore {
		fmt.Println("Since your character was in the hardcore league, it will now be deleted.\n\nCharacter Stats: \n" + CharacterInfo(c) + "\n\nRest in peace.")
		utils.DeleteFile(c.Name)
	} else {
		xp := ApplyXPPenalty(c)
		fmt.Println("Your character loses " + fmt.Sprint(xp) + " XP.")
		c.Current_HP = c.Max_HP / 2
		SaveCharacter(c)
	}
}

//asks the user about character deletion
func DeleteCharacter() {
	fmt.Println("Which character do you want to delete?")
	c := CharacterChoice()
	if c != nil {
		if c.Default {
			fmt.Println("This is your default character! You cannot delete them, please switch to a different one first.")
			return
		} else {
			o := 0
			for o < 1 {
				fmt.Println("Are you sure you want to delete " + c.Name + " (Level" + fmt.Sprint(c.Level) + ")? y/n")
				switch utils.GetUserInput() {
				case "y":
					utils.DeleteFile(c.Name)
					fmt.Println("Successfully deleted character " + c.Name)
					o += 1
				case "n":
					fmt.Println("Aborted deletion of character " + c.Name)
					o += 1

				default:
					fmt.Println("Invalid input. Please try again.")
				}

			}

		}

	}
}
