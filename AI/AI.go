package ai

import "github.com/mdokusV/dices-game/globalVar"

func DecideMove(tableUsed [globalVar.NumberOfStates]bool, diceSlice []int) int {
	for i := 0; i < len(tableUsed); i++ {
		if !tableUsed[i] {
			return i + 1
		}
	}
	return 0
}
