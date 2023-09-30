package globalStructs

import (
	"fmt"
	"math/rand"

	ai "github.com/mdokusV/dices-game/AI"
	"github.com/mdokusV/dices-game/globalVar"
	"github.com/mdokusV/dices-game/helpers"
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

func (player *Player) FullTour() {
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
	diceSlice := make([]int, 5)
	rolls := []bool{true, true, true, true, true}
	numberOfRemainingRolls := 3
	StateChoice := 0

	for numberOfRemainingRolls > 0 {
		helpers.RoleDices(diceSlice, rolls)
		if globalVar.IO_human {
			fmt.Printf("Your Rolls:\n%d\n\n", diceSlice)
		}

		numberOfRemainingRolls--

		if globalVar.IO_human {
			player.PrintPossibleChoices(numberOfRemainingRolls)
		}

		//get legal state choice
		StateChoice, rolls = player.acceptedChoice(numberOfRemainingRolls, diceSlice)

		//execute state choice
		if StateChoice != 0 {
			StateChoice--
			player.TableScore[StateChoice] = optionFuncMap[StateChoice](diceSlice)    //Write score to table
			if numberOfRemainingRolls == 2 && StateChoice >= 5 && StateChoice <= 11 { //In first move
				player.TableScore[StateChoice] *= 2
			}
			player.Score += player.TableScore[StateChoice] //Update sum score
			player.TableUsed[StateChoice] = true           //Update used table
			return
		}
	}
	panic("NumberOfRemainingRolls is less than or equal 0")

}

func (player *Player) acceptedChoice(numberOfRemainingRolls int, diceSlice []int) (newChoiceForMove int, rolls []bool) {
	accepted := false
	for !accepted {
		newChoiceForMove, rolls = player.GiveMoveChoice(diceSlice)

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
	return newChoiceForMove, rolls
}

func (player *Player) GiveMoveChoice(diceSlice []int) (choice int, rolls []bool) {
	choice = 0
	if globalVar.IO_human {
		fmt.Scanln(&choice)
	} else {
		choice = ai.DecideMove(player.TableUsed, diceSlice)
	}

	if choice != 0 {
		return choice, []bool{true, true, true, true, true}
	}

	//decide which dice to roll
	if globalVar.IO_human {
		rolls = []bool{true, true, true, true, true}
	} else {
		rolls = []bool{rand.Intn(2) == 0, rand.Intn(2) == 0, rand.Intn(2) == 0, rand.Intn(2) == 0, rand.Intn(2) == 0}
	}

	return choice, rolls

}
