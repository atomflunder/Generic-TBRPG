package game

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/phxenix-w/Generic-TBRPG/utils"
)

type Artifact struct {
	Name        string
	Description string
	Tag         string
	Rarity      Rarity
}

var AntiqueWhetstone = Artifact{
	Name:        "Antique Whetstone",
	Description: "Boosts your current weapon's damage permanently by 20%. Up to 100%",
	Tag:         "Weapon",
	Rarity:      RarityRare,
}

var SmallLifeGem = Artifact{
	Name:        "Small Life Gem",
	Description: "Increases your character's HP permanently by 20.",
	Tag:         "Character",
	Rarity:      RarityRare,
}

var AllArtifacts = []Artifact{
	AntiqueWhetstone, SmallLifeGem,
}

//uses the artifact
func (a Artifact) Use(c *Character) {
	switch a.Tag {
	case "Weapon":
		a.UseWeaponArtifact(c)
	case "Character":
		a.UseCharacterArtifact(c)
	}
	c.Save()
}

//the logic for the artifacts applied to the weapon
func (a Artifact) UseWeaponArtifact(c *Character) {
	switch a.Name {
	case "Antique Whetstone":
		if c.Weapon.Quality < 100 {
			c.Weapon.LowAttack += int(float32(c.Weapon.LowAttack) * 0.2)
			c.Weapon.HighAttack += int(float32(c.Weapon.HighAttack) * 0.2)
			c.Weapon.Quality += 20
		}

	}
}

//the logic for the artifacts applied to the character
func (a Artifact) UseCharacterArtifact(c *Character) {
	switch a.Name {
	case "Small Life Gem":
		c.Max_HP += 20
		c.Current_HP += 20
	}
}

//asks the user if they want to activate multiple artifacts
func ActivateArtifacts(al []Artifact, c *Character) {
	for _, a := range al {
		fmt.Println("Do you want to use this artifact? " + a.Name + " (" + a.Description + ")\nType y to confirm or anything else to dismiss.")
		if strings.ToLower(utils.GetUserInput()) == "y" {
			a.Use(c)
		}
	}

}

//gets you every artifact by their rarity
func GetArtifactsByRarity(r Rarity) []Artifact {
	var artifactList []Artifact

	for _, i := range AllArtifacts {
		if i.Rarity == r {
			artifactList = append(artifactList, i)
		}
	}
	return artifactList
}

//picks a random artifact from the list
func PickRandomArtifact(al []Artifact) Artifact {
	utils.GetNewRandomSeed()
	n := rand.Intn(len(al))
	return al[n]
}

//prints artifacts out in a nice format
func PrintArtifacts(a []Artifact) string {
	var nameList []string
	for y, x := range a {
		nameList = append(nameList, fmt.Sprint(y+1)+") "+x.Name+" ("+x.Description+")")
	}
	return strings.Join(nameList, ", ")
}
