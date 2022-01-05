package game

import (
	"fmt"

	"github.com/phxenix-w/Generic-TBRPG/utils"
)

//calculates all of the loot the monster drops
func GetMonsterLoot(m *Monster, c *Character) {
	g := CalculateGoldDrops(m)
	c.Gold += g
	fmt.Println("Congratulations, the monster you killed dropped " + fmt.Sprint(g) + " gold. It is automatically picked up.")

	i := CalculateItemDrops(m)
	if len(i) != 0 {
		AddItems(c, i)
		fmt.Println("Congratulations, the monster you killed dropped these items: " + PrintItems(i) + ". They have been automatically picked up.")
	}
	w := CalculateWeaponDrops(m)
	if len(w) != 0 {
		fmt.Println("Congratulations, the monster you killed dropped these weapons: " + PrintWeapons(w))
		SwitchWeaponFromList(c, w)
	}
	a := CalculateArtifactDrops(m)
	if len(a) != 0 {
		fmt.Println("Congratulations, the monster you killed dropped these artifacts: " + PrintArtifacts(a))
		ActivateArtifacts(a, c)
	}
}

//calculates the weapons that drop from a monster
func CalculateWeaponDrops(m *Monster) []Weapon {
	var weaponList []Weapon

	n := utils.GetRandomNumber(100)

	if float32(n) <= m.Weapon_Drops.Common {
		w := PickRandomWeapon(GetWeaponsByRarityAndLevel(RarityCommon, m))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Uncommon {
		w := PickRandomWeapon(GetWeaponsByRarityAndLevel(RarityUncommon, m))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Rare {
		w := PickRandomWeapon(GetWeaponsByRarityAndLevel(RarityRare, m))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Epic {
		w := PickRandomWeapon(GetWeaponsByRarityAndLevel(RarityEpic, m))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Legendary {
		w := PickRandomWeapon(GetWeaponsByRarityAndLevel(RarityLegendary, m))
		weaponList = append(weaponList, w)
	}

	return weaponList
}

//calculates the items that drop from a monster
func CalculateItemDrops(m *Monster) []Item {
	var itemList []Item

	n := utils.GetRandomNumber(100)

	if float32(n) <= m.Item_Drops.Common {
		i := PickRandomItem(GetItemsByRarity(RarityCommon))
		itemList = append(itemList, i)
	}

	if float32(n) <= m.Item_Drops.Uncommon {
		i := PickRandomItem(GetItemsByRarity(RarityUncommon))
		itemList = append(itemList, i)
	}

	if float32(n) <= m.Item_Drops.Rare {
		i := PickRandomItem(GetItemsByRarity(RarityRare))
		itemList = append(itemList, i)
	}

	if float32(n) <= m.Item_Drops.Epic {
		i := PickRandomItem(GetItemsByRarity(RarityEpic))
		itemList = append(itemList, i)
	}

	if float32(n) <= m.Item_Drops.Legendary {
		i := PickRandomItem(GetItemsByRarity(RarityLegendary))
		itemList = append(itemList, i)
	}

	return itemList
}

//calculates the artifacts that drop from a monster
func CalculateArtifactDrops(m *Monster) []Artifact {
	var artifactList []Artifact

	n := utils.GetRandomNumber(100)

	if float32(n) <= m.Artifact_Drops.Common {
		a := PickRandomArtifact(GetArtifactsByRarity(RarityCommon))
		artifactList = append(artifactList, a)
	}
	if float32(n) <= m.Artifact_Drops.Uncommon {
		a := PickRandomArtifact(GetArtifactsByRarity(RarityUncommon))
		artifactList = append(artifactList, a)
	}
	if float32(n) <= m.Artifact_Drops.Rare {
		a := PickRandomArtifact(GetArtifactsByRarity(RarityRare))
		artifactList = append(artifactList, a)
	}
	if float32(n) <= m.Artifact_Drops.Epic {
		a := PickRandomArtifact(GetArtifactsByRarity(RarityEpic))
		artifactList = append(artifactList, a)
	}
	if float32(n) <= m.Artifact_Drops.Legendary {
		a := PickRandomArtifact(GetArtifactsByRarity(RarityLegendary))
		artifactList = append(artifactList, a)
	}

	return artifactList
}

//calculates the amount of gold dropped by a monster
func CalculateGoldDrops(m *Monster) int {
	g := utils.GetRandomNumberInRange(m.Gold_Drop_Min, m.Gold_Drop_Max)
	return g
}
