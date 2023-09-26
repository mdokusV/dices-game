package helpers

import "github.com/mdokusV/dices-game/globalVar"

func RoleDices(currentState []int, whatRoll []bool) {

	for i := 0; i < len(currentState); i++ {
		if whatRoll[i] {
			currentState[i] = globalVar.RandomGenerator.Intn(6) + 1
		}
	}
	ReverseSort(currentState)
}
