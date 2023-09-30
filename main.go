package main

import (
	"github.com/mdokusV/dices-game/game"
	globalStructs "github.com/mdokusV/dices-game/globalStructs"
	"github.com/mdokusV/dices-game/globalVar"
)

func main() {
	//initialize needed variables

	playerGroup := globalStructs.NewPlayerGroup(globalVar.NumberOfPlayers)

	//start game
	game.Start(playerGroup)
}
