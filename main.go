package main

import (
	"github.com/phxenix-w/gotestgame/game"
	"github.com/phxenix-w/gotestgame/utils"
)

func main() {
	utils.SetupDB()

	for {
		game.MainMenu()
	}

}
