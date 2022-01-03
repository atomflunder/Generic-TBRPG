package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
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
}

//calculates the weapons that drop from a monster
func CalculateWeaponDrops(m *Monster) []Weapon {
	var weaponList []Weapon

	n := utils.GetRandomNumber(100)

	if float32(n) <= m.Weapon_Drops.Common {
		w := PickRandomWeapon(GetWeaponsByRarity(RarityCommon))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Uncommon {
		w := PickRandomWeapon(GetWeaponsByRarity(RarityUncommon))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Rare {
		w := PickRandomWeapon(GetWeaponsByRarity(RarityRare))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Epic {
		w := PickRandomWeapon(GetWeaponsByRarity(RarityEpic))
		weaponList = append(weaponList, w)
	}

	if float32(n) <= m.Weapon_Drops.Legendary {
		w := PickRandomWeapon(GetWeaponsByRarity(RarityLegendary))
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

//calculates the amount of gold dropped by a monster
func CalculateGoldDrops(m *Monster) int {
	g := utils.GetRandomNumberInRange(m.Gold_Drop_Min, m.Gold_Drop_Max)
	return g
}
