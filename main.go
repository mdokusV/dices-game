package main

import (
	"github.com/mdokusV/dices-game/game"
	"github.com/mdokusV/dices-game/globalVar"
)

func main() {
	//initialize needed variables
	playerTour := 0
	playerGroup := game.NewPlayerGroup(globalVar.NumberOfPlayers)

	//Play Tour

	for {
		playerGroup[playerTour].FullTour()
		playerTour++
		if playerTour >= globalVar.NumberOfPlayers {
			playerTour = 0
		}
	}

}
