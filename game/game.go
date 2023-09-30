package game

import (
	"fmt"

	"github.com/mdokusV/dices-game/dices"
	"github.com/mdokusV/dices-game/globalStructs"
	"github.com/mdokusV/dices-game/globalVar"
)

func Start(playerGroup []globalStructs.Player) {

	dices := dices.Dices{}
	dices.NewDices()

	playerTour := 0
	for {
		if globalVar.IO_human {
			fmt.Println("Player", playerTour+1, "turn")
			fmt.Println(playerGroup[playerTour].TableUsed)
		}

		playerGroup[playerTour].FullTour(dices)
		playerTour++
		if playerTour == globalVar.NumberOfPlayers {
			if allMovesCompleted(playerGroup[0].TableUsed) {
				break
			}
			playerTour = 0
		}
	}
}

func allMovesCompleted(states [globalVar.NumberOfStates]bool) bool {
	for _, state := range states {
		if !state {
			return false
		}
	}
	return true
}
