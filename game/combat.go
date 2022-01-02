package game

import (
	"fmt"

	"github.com/phxenix-w/gotestgame/utils"
)

//interface for player and monster, to be used later on
type Fighter interface {
	Turn()
	RollDamage()
	ApplyDamage()
	Info() string
}

//the main combat loop
func Combat(p *Character, e *Monster) {
	e.RefreshHP()

	fmt.Println("You encounter a " + e.Name + "!")

	for p.Current_HP != 0 && e.Current_HP != 0 {
		if p.Weapon.Range >= e.Range {
			p.Turn(e)
			if e.Current_HP == 0 {
				break
			}
			e.Turn(p)
		} else {
			e.Turn(p)
			if p.Current_HP == 0 {
				break
			}
			p.Turn(e)
		}
	}
	if p.Current_HP > 0 && e.Current_HP == 0 {
		xp := p.GainXP(e)
		fmt.Println("Congratulations, you won the fight. You gain " + fmt.Sprint(xp) + " XP")
		GetMonsterLoot(e, p)
		UpdateLevel(p)
		p.Save()
	} else if e.Current_HP > 0 && p.Current_HP == 0 {
		fmt.Println("You lost the fight against " + e.Name + ". It has " + fmt.Sprint(e.Current_HP) + " HP left.")
		p.Death()
	}
}

//one turn for the player
func (c *Character) Turn(m *Monster) {
	fmt.Println(`It is your turn! What do you want to do?
1) Attack with your equipped weapon
2) Use Item
3) View Character
4) View Enemy
5) Pass`)
	switch utils.GetUserInput() {
	case "1":
		c.RollDamage(m)
	case "2":
		ItemChoice(c, m)
	case "3":
		fmt.Println(c.Info())
		c.Turn(m)
	case "4":
		fmt.Println(m.Info())
		c.Turn(m)
	case "5":

	default:
		fmt.Println("Invalid input. Please try again.")
		c.Turn(m)
	}
}

//one turn for the enemy
func (m *Monster) Turn(c *Character) {
	fmt.Println(`It is the turn of the enemy.`)
	m.RollDamage(c)
}

//rolls the damage for the player, and applies it to the enemy
func (c *Character) RollDamage(m *Monster) {
	n := utils.GetRandomNumber(10000)
	cr := c.Weapon.CritChance * 100

	if n <= int(cr) {
		d := c.Weapon.HighAttack * 2
		m.ApplyDamage(d)
		fmt.Println("Critical hit!! You attacked for " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
	} else {
		an := utils.GetRandomNumber(100)
		ac := c.Weapon.Accuracy

		if an <= ac {
			d := utils.GetRandomNumberInRange(c.Weapon.LowAttack, c.Weapon.HighAttack)
			m.ApplyDamage(d)
			fmt.Println("You attacked for " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
		} else {
			fmt.Println("You missed your attack!")
		}

	}

}

//rolls the damage for the monster, and applies it to the player
func (m *Monster) RollDamage(c *Character) {
	n := utils.GetRandomNumber(10000)
	cr := m.CritChance * 100

	if n <= int(cr) {
		d := m.HighAttack * 2
		c.ApplyDamage(d)
		fmt.Println("Critical hit!! The enemy attacked you for " + fmt.Sprint(d) + " damage. You have " + fmt.Sprint(c.Current_HP) + " HP left")
	} else {
		an := utils.GetRandomNumber(100)
		ac := m.Accuracy

		if an <= ac {
			d := utils.GetRandomNumberInRange(m.LowAttack, m.HighAttack)
			c.ApplyDamage(d)
			fmt.Println("The enemy attacked you for " + fmt.Sprint(d) + " damage. You have " + fmt.Sprint(c.Current_HP) + " HP left")
		} else {
			fmt.Println("The enemy missed their attack!")
		}

	}
}

//applies the damage to the player
func (c *Character) ApplyDamage(d int) {
	c.Current_HP -= d
	if c.Current_HP < 0 {
		c.Current_HP = 0
	}
}

//applies the damage to the enemy
func (m *Monster) ApplyDamage(d int) {
	m.Current_HP -= d
	if m.Current_HP < 0 {
		m.Current_HP = 0
	}
}

//applies the damage from an item to the enemy
func (m *Monster) ApplyItemDamage(d int, i Item) {
	m.ApplyDamage(d)
	fmt.Println("You throw a " + i.Name + " at the " + m.Name + ". It deals " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
}

//heals the player for the amount specified
func (c *Character) Heal(d int) {
	if d == -1 {
		c.Current_HP = c.Max_HP
		fmt.Println("You have healed your HP fully. Your HP is now " + fmt.Sprint(c.Current_HP))
	} else {
		c.Current_HP += d
		if c.Current_HP >= c.Max_HP {
			c.Current_HP = c.Max_HP
		}
		fmt.Println("You have healed for " + fmt.Sprint(d) + " HP. Your HP is now " + fmt.Sprint(c.Current_HP))
	}

}

//applies the xp gains to the player
func (c *Character) GainXP(e *Monster) int {
	n := utils.GetRandomNumberInRange(e.XP_Min, e.XP_Max)
	c.XP += n
	return n
}
