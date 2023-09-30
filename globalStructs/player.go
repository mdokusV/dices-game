package globalStructs

import (
	"fmt"

	ai "github.com/mdokusV/dices-game/AI"
	"github.com/mdokusV/dices-game/dices"
	"github.com/mdokusV/dices-game/globalVar"
	state "github.com/mdokusV/dices-game/states"
)

type Player struct {
	ID         int
	Score      int
	TableScore [globalVar.NumberOfStates]int
	TableUsed  [globalVar.NumberOfStates]bool
}

func NewPlayerGroup(size int) []Player {

	groupPlayer := make([]Player, size)
	for ID := range groupPlayer {
		groupPlayer[ID].ID = ID
		groupPlayer[ID].Score = 0
		for i := 0; i < globalVar.NumberOfStates; i++ {
			groupPlayer[ID].TableUsed[i] = false
			groupPlayer[ID].TableScore[i] = 0
		}
	}
	return groupPlayer
}

func ResetPlayerGroup(groupPlayer []Player) {
	for ID := range groupPlayer {
		groupPlayer[ID].Score = 0
		for i := 0; i < globalVar.NumberOfStates; i++ {
			groupPlayer[ID].TableUsed[i] = false
			groupPlayer[ID].TableScore[i] = 0
		}
	}
}

func (player *Player) PrintPossibleChoices(numberOfRemainingRolls int) {
	Options := [globalVar.NumberOfStates + 1]string{
		"Roll",
		"OnePair",
		"TwoPair",
		"ThreeOfKind",
		"SmallStraight",
		"BigStraight",
		"FullHouse",
		"FourOfKind",
		"FiveOfKind",
		"Odd",
		"Even",
		"Chance",
	}
	fmt.Println("These are your choices:")

	for key := 0; key < len(Options); key++ {
		if numberOfRemainingRolls == 0 && key == 0 {
			continue
		}

		fmt.Printf("%d: %s", key, Options[key])

		if key >= 1 && player.TableUsed[key-1] {
			fmt.Printf("\tUSED")
		}
		fmt.Println()
	}
}

func (player *Player) FullTour(dices dices.Dices) {
	//define map from int choice to functions
	optionFuncMap := []func([]int) int{
		state.OnePair,
		state.TwoPair,
		state.ThreeOfKind,
		state.SmallStraight,
		state.BigStraight,
		state.FullHouse,
		state.FourOfKind,
		state.FiveOfKind,
		state.Odd,
		state.Even,
		state.Chance,
	}

	//prep variables
	numberOfRemainingRolls := 3
	StateChoice := 0

	for numberOfRemainingRolls > 0 {
		dices.RoleDices()
		if globalVar.IO_human {
			fmt.Printf("Your Rolls:\n%d\n\n", dices.DiceSlice)
		}

		numberOfRemainingRolls--

		if globalVar.IO_human {
			player.PrintPossibleChoices(numberOfRemainingRolls)
		}

		//get legal state choice
		StateChoice = player.acceptedChoice(numberOfRemainingRolls, dices)

		//execute state choice
		if StateChoice != 0 {
			StateChoice--
			player.TableScore[StateChoice] = optionFuncMap[StateChoice](dices.DiceSlice) //Write score to table
			if numberOfRemainingRolls == 2 && StateChoice >= 3 && StateChoice <= 9 {     //In first move
				player.TableScore[StateChoice] *= 2
			}
			player.Score += player.TableScore[StateChoice] //Update sum score
			player.TableUsed[StateChoice] = true           //Update used table
			return
		}
	}
	panic("NumberOfRemainingRolls is less than or equal 0")

}

func (player *Player) acceptedChoice(numberOfRemainingRolls int, dices dices.Dices) (newChoiceForMove int) {
	accepted := false
	for !accepted {
		newChoiceForMove = player.GiveMoveChoice(dices)

		if newChoiceForMove > globalVar.NumberOfStates || newChoiceForMove < 0 {
			fmt.Println("Action number not in range, try again")
			continue
		}
		if newChoiceForMove == 0 && numberOfRemainingRolls == 0 {
			fmt.Println("Its your last roll, you need to make a choice")
			continue
		}
		if newChoiceForMove == 0 && numberOfRemainingRolls > 0 {
			accepted = true
			continue
		}
		if player.TableUsed[newChoiceForMove-1] {
			fmt.Println("Action already used, try again")
			continue
		}
		accepted = true
	}
	return newChoiceForMove
}

func (player *Player) GiveMoveChoice(dices dices.Dices) (choice int) {
	choice = 0
	if globalVar.IO_human {
		fmt.Scanln(&choice)
	} else {
		choice = ai.DecideMove(player.TableUsed, dices.DiceSlice)
	}

	if choice != 0 {
		return choice
	}

	//decide which dice to roll
	if globalVar.IO_human {
		dices.Rolls = []bool{true, true, true, true, true}
	} else {
		dices.Rolls = []bool{globalVar.FastRandomGenerator.Uint32n(2) == 0, globalVar.FastRandomGenerator.Uint32n(2) == 0, globalVar.FastRandomGenerator.Uint32n(2) == 0, globalVar.FastRandomGenerator.Uint32n(2) == 0, globalVar.FastRandomGenerator.Uint32n(2) == 0}
	}

	return choice

}
