package game

type Rarity struct {
	Common    float32
	Uncommon  float32
	Rare      float32
	Epic      float32
	Legendary float32
}

var (
	RarityCommon    = Rarity{1, 0, 0, 0, 0}
	RarityUncommon  = Rarity{0, 1, 0, 0, 0}
	RarityRare      = Rarity{0, 0, 1, 0, 0}
	RarityEpic      = Rarity{0, 0, 0, 1, 0}
	RarityLegendary = Rarity{0, 0, 0, 0, 1}
)
