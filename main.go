package main

import (
	"github.com/phxenix-w/gotestgame/game"
	"github.com/phxenix-w/gotestgame/utils"
)

func main() {
	if !utils.DirectoryCheck(utils.DBDirectory) {
		utils.MakeDirectory(utils.DBDirectory)
	}

	utils.SetupDB()

	for {
		game.MainMenu()
	}

}
