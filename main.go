package main

import (
	"github.com/phxenix-w/gotestgame/game"
	"github.com/phxenix-w/gotestgame/utils"
)

func main() {
	if !utils.DirectoryCheck(utils.ProfileDirectory) {
		utils.MakeDirectory(utils.ProfileDirectory)
	}

	for {
		game.MainMenu()
	}

}
