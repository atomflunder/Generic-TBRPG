package game

import (
	"fmt"
	"strings"

	"github.com/phxenix-w/gotestgame/utils"
)

type Stage struct {
	Name     string
	LevelReq int
	Monsters []Monster
	Boss     bool //placeholder
}

var Foothills = Stage{
	Name:     "Foothills",
	LevelReq: 1,
	Monsters: []Monster{Zombie, Skeleton, Dwarf},
}

var OldBridge = Stage{
	Name:     "Old Bridge",
	LevelReq: 3,
	Monsters: []Monster{Zombie, Skeleton, Dwarf},
}

var AllStages = []Stage{
	Foothills, OldBridge,
}

//the main menu for choosing stages.
func StageMenu(c *Character) {
	fmt.Println("Which zone do you want to enter? Type their number to enter the zone.\n" + PrintStages(AllStages))
	s := MatchStageIndex(utils.StringToInt(utils.GetUserInput()), AllStages)
	if s != nil {
		if c.Level >= s.LevelReq {
			fmt.Println("You have entered " + s.Name)
			for {
				m := PickStageMonster(*s)
				Combat(c, &m)
				fmt.Println("Do you want to leave this zone? Type y to confirm or anything else to continue exploring.")
				if strings.ToLower(utils.GetUserInput()) == "y" {
					break
				}
			}
		} else {
			fmt.Println("Your level does not match the level requirement of this zone. Please come back later.")
		}
	} else {
		fmt.Println("Invalid input. Please try again.")
	}

}

//prints out the stages in a nice format
func PrintStages(sl []Stage) string {
	var stages string

	for i, s := range sl {
		stages = stages + fmt.Sprint(i+1) + ") " + s.Name + " (Level " + fmt.Sprint(s.LevelReq) + ")\n"
	}
	return stages
}

//matches the index of the stage to the stage
func MatchStageIndex(p int, sl []Stage) *Stage {
	for x := range sl {
		if x+1 == p {
			return &sl[x]
		}
	}
	return nil
}

//picks a random monster available
func PickStageMonster(s Stage) Monster {
	return PickRandomMonster(s.Monsters)
}
