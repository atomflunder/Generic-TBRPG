package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/phxenix-w/gotestgame/utils"
)

//the main combat loop
func Combat(p *Character, e *Monster) {
	RefreshMonsterHP(e)

	for p.Current_HP != 0 && e.Current_HP != 0 {
		if p.Weapon.Range > e.Range {
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
		UpdateLevel(p)
		SaveCharacter(p)
	} else if e.Current_HP > 0 && p.Current_HP == 0 {
		fmt.Println("You lost the fight against " + e.Name + ". It has " + fmt.Sprint(e.Current_HP) + " left. Your character will be deleted.")
		DeleteCharacter(p.Name)
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
		d := RollDamage(p.Weapon.LowAttack, p.Weapon.HighAttack)
		ApplyDamageToEnemy(d, e)
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
	d := RollDamage(e.LowAttack, e.HighAttack)
	ApplyDamageToPlayer(d, p)
}

//rolls the damage
func RollDamage(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//applies the damage to the player
func ApplyDamageToPlayer(d int, c *Character) {
	c.Current_HP -= d
	if c.Current_HP < 0 {
		c.Current_HP = 0
	}
	fmt.Println("The enemy attacked you for " + fmt.Sprint(d) + " damage. You have " + fmt.Sprint(c.Current_HP) + " HP left")
}

//applies the damage to the enemy
func ApplyDamageToEnemy(d int, m *Monster) {
	m.Current_HP -= d
	if m.Current_HP < 0 {
		m.Current_HP = 0
	}
	fmt.Println("You attacked for " + fmt.Sprint(d) + " damage. The enemy has " + fmt.Sprint(m.Current_HP) + " HP left.")
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
	rand.Seed(time.Now().UnixNano())
	t := e.XP_Max - e.XP_Min
	n := rand.Intn(t) + e.XP_Min
	p.XP += n
	return n
}
