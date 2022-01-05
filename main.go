package main

import (
	"github.com/phxenix-w/Generic-TBRPG/game"
	"github.com/phxenix-w/Generic-TBRPG/utils"
)

func main() {
	utils.SetupDB()

	for {
		game.MainMenu()
	}

}
