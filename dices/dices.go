package dices

import (
	"github.com/mdokusV/dices-game/globalVar"
	"github.com/mdokusV/dices-game/helpers"
)

type Dices struct {
	DiceSlice []int
	Rolls     []bool
}

func (dices *Dices) NewDices() {
	dices.DiceSlice = make([]int, globalVar.NumberOfDices)
	dices.Rolls = []bool{true, true, true, true, true}

}

func (dices *Dices) RoleDices() {

	for i := 0; i < len(dices.DiceSlice); i++ {
		if dices.Rolls[i] {
			// currentState[i] = globalVar.RandomGenerator.Intn(6) + 1
			dices.DiceSlice[i] = int(globalVar.FastRandomGenerator.Uint32n(6)) + 1

		}
	}
	helpers.ReverseSort2(dices.DiceSlice)
}
