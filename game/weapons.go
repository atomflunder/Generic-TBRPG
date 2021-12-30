package game

type Weapon struct {
	Name        string
	Description string
	Rarity      string
	LowAttack   int
	HighAttack  int
	AttackSpeed float32
	Range       int
	ReqStr      int
	ReqDex      int
	ReqInt      int
}

var RustyClub = Weapon{
	Name:        "Rusty Club",
	Description: "Starter Weapon for the Barbarian.",
	Rarity:      "Common",
	LowAttack:   15,
	HighAttack:  40,
	AttackSpeed: 1.2,
	Range:       10,
	ReqStr:      15,
	ReqDex:      5,
	ReqInt:      0,
}

var RustyDagger = Weapon{
	Name:        "Rusty Dagger",
	Description: "Starter Weapon for the Rogue",
	Rarity:      "Common",
	LowAttack:   5,
	HighAttack:  30,
	AttackSpeed: 1.7,
	Range:       6,
	ReqStr:      5,
	ReqDex:      10,
	ReqInt:      5,
}

var RustyStaff = Weapon{
	Name:        "Rusty Staff",
	Description: "Starter Weapon for the Mage",
	Rarity:      "Common",
	LowAttack:   10,
	HighAttack:  35,
	AttackSpeed: 1.4,
	Range:       14,
	ReqStr:      5,
	ReqDex:      5,
	ReqInt:      10,
}
