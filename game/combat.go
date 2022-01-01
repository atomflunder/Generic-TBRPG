package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main combat loop
func Combat(p *Character, e *Monster) {
	RefreshMonsterHP(e)

	fmt.Println("You encounter a " + e.Name + "!")

	for p.Current_HP != 0 && e.Current_HP != 0 {
		if p.Weapon.Range >= e.Range {
			PlayerTurn(p, e)
			if e.Current_HP == 0 {
				break
			}
			EnemyTurn(e, p)
		} else {
			EnemyTurn(e, p)
			if p.Current_HP == 0 {
				break
			}
			PlayerTurn(p, e)
		}
	}
	if p.Current_HP > 0 && e.Current_HP == 0 {
		xp := GainXP(p, e)
		fmt.Println("Congratulations, you won the fight. You gain " + fmt.Sprint(xp) + " XP")
		GetMonsterLoot(e, p)
		UpdateLevel(p)
		SaveCharacter(p)
	} else if e.Current_HP > 0 && p.Current_HP == 0 {
		fmt.Println("You lost the fight against " + e.Name + ". It has " + fmt.Sprint(e.Current_HP) + " left.")
		CharacterDeath(p)
	}
}

//one turn for the player
func PlayerTurn(p *Character, e *Monster) {
	fmt.Println(`It is your turn! What do you want to do?
1) Attack with your equipped weapon
2) Use Item
3) View Character
4) View Enemy
5) Pass`)
	switch utils.GetUserInput() {
	case "1":
		RollDamagePlayer(p, e)
	case "2":
		ItemChoice(p, e)
	case "3":
		fmt.Println(CharacterInfo(p))
		PlayerTurn(p, e)
	case "4":
		fmt.Println(MonsterInfo(e))
		PlayerTurn(p, e)
	case "5":

	default:
		fmt.Println("Invalid input. Please try again.")
		PlayerTurn(p, e)
	}
}

//one turn for the enemy
func EnemyTurn(e *Monster, p *Character) {
	fmt.Println(`It is the turn of the enemy.`)
	RollDamageMonster(e, p)
}

//rolls the damage for the player, and applies it to the enemy
func RollDamagePlayer(c *Character, m *Monster) {
	n := utils.GetRandomNumber(10000)
	cr := c.Weapon.CritChance * 100

	if n <= int(cr) {
		d := c.Weapon.HighAttack * 2
		ApplyDamageToEnemy(d, m)
		fmt.Println("Critical hit!! You attacked for " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
	} else {
		an := utils.GetRandomNumber(100)
		ac := c.Weapon.Accuracy

		if an <= ac {
			d := utils.GetRandomNumberInRange(c.Weapon.LowAttack, c.Weapon.HighAttack)
			ApplyDamageToEnemy(d, m)
			fmt.Println("You attacked for " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
		} else {
			fmt.Println("You missed your attack!")
		}

	}

}

//rolls the damage for the monster, and applies it to the player
func RollDamageMonster(m *Monster, c *Character) {
	n := utils.GetRandomNumber(10000)
	cr := m.CritChance * 100

	if n <= int(cr) {
		d := m.HighAttack * 2
		ApplyDamageToPlayer(d, c)
		fmt.Println("Critical hit!! The enemy attacked you for " + fmt.Sprint(d) + " damage. You have " + fmt.Sprint(c.Current_HP) + " HP left")
	} else {
		an := utils.GetRandomNumber(100)
		ac := m.Accuracy

		if an <= ac {
			d := utils.GetRandomNumberInRange(m.LowAttack, m.HighAttack)
			ApplyDamageToPlayer(d, c)
			fmt.Println("The enemy attacked you for " + fmt.Sprint(d) + " damage. You have " + fmt.Sprint(c.Current_HP) + " HP left")
		} else {
			fmt.Println("The enemy missed their attack!")
		}

	}
}

//applies the damage to the player
func ApplyDamageToPlayer(d int, c *Character) {
	c.Current_HP -= d
	if c.Current_HP < 0 {
		c.Current_HP = 0
	}
}

//applies the damage to the enemy
func ApplyDamageToEnemy(d int, m *Monster) {
	m.Current_HP -= d
	if m.Current_HP < 0 {
		m.Current_HP = 0
	}
}

//applies the damage from an item to the enemy
func ApplyItemDamageToEnemy(d int, m *Monster, i Item) {
	ApplyDamageToEnemy(d, m)
	fmt.Println("You throw a " + i.Name + " at the " + m.Name + ". It deals " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
}

//heals the player for the amount specified
func HealPlayer(d int, c *Character) {
	c.Current_HP += d
	if c.Current_HP >= c.Max_HP {
		c.Current_HP = c.Max_HP
	}
	fmt.Println("You have healed for " + fmt.Sprint(d) + " HP. Your HP is now " + fmt.Sprint(c.Current_HP))
}

//applies the xp gains to the player
func GainXP(p *Character, e *Monster) int {
	n := utils.GetRandomNumberInRange(e.XP_Min, e.XP_Max)
	p.XP += n
	return n
}
