package game

import (
	"github.com/mdokusV/dices-game/globalStructs"
	"github.com/mdokusV/dices-game/globalVar"
)

func Start(playerGroup []globalStructs.Player) {
	playerTour := 0
	for {
		playerGroup[playerTour].FullTour()
		playerTour++
		if playerTour >= globalVar.NumberOfPlayers {
			if allMovesCompleted(&playerGroup[0].TableUsed) {
				break
			}
			playerTour = 0
		}
	}
}

func allMovesCompleted(states *map[int]bool) bool {
	for _, state := range *states {
		if !state {
			return false
		}
	}
	return true
}
