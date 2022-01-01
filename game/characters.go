package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

type Character struct {
	Name         string `json:"name"`
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
			fmt.Println("Invalid input, please try again.")

		}
	}
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
		charList = charList + "Name: " + c.Name + "\n" + "Level: " + fmt.Sprint(c.Level) + "\n\n"
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

//deletes a saved character
func DeleteCharacter(s string) {
	err := os.Remove("./savedata/characters/" + s + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted profile of " + s)
}
