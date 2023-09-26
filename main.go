package main

import (
	globalStructs "github.com/mdokusV/dices-game/globalStructs"
	"github.com/mdokusV/dices-game/globalVar"
)

func main() {
	//initialize needed variables
	playerTour := 0
	playerGroup := globalStructs.NewPlayerGroup(globalVar.NumberOfPlayers)

	//Play Tour

	for {
		playerGroup[playerTour].FullTour()
		playerTour++
		if playerTour >= globalVar.NumberOfPlayers {
			playerTour = 0
		}
	}

}
