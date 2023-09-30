package globalStructs

import (
	"fmt"

	ai "github.com/mdokusV/dices-game/AI"
	"github.com/mdokusV/dices-game/globalVar"
	"github.com/mdokusV/dices-game/helpers"
	state "github.com/mdokusV/dices-game/states"
)

type Player struct {
	ID         int
	Score      int
	TableScore map[int]int
	TableUsed  map[int]bool
}

func NewPlayerGroup(size int) []Player {

	groupPlayer := make([]Player, size)
	for ID := range groupPlayer {
		groupPlayer[ID] = Player{
			ID:         ID,
			Score:      0,
			TableScore: make(map[int]int),
			TableUsed:  make(map[int]bool),
		}
		for i := 2; i <= 12; i++ {
			groupPlayer[ID].TableUsed[i] = false
			groupPlayer[ID].TableScore[i] = 0
		}
	}
	return groupPlayer
}

func DeletePlayerGroup(groupPlayer []Player) {
	for ID := range groupPlayer {
		groupPlayer[ID].Score = 0
		for i := 2; i <= 12; i++ {
			groupPlayer[ID].TableUsed[i] = false
			groupPlayer[ID].TableScore[i] = 0
		}
	}
}

func (player *Player) PrintPossibleChoices(numberOfRemainingRolls int) {
	Options := map[int]string{
		1:  "Roll",
		2:  "OnePair",
		3:  "TwoPair",
		4:  "ThreeOfKind",
		5:  "SmallStraight",
		6:  "BigStraight",
		7:  "FullHouse",
		8:  "FourOfKind",
		9:  "FiveOfKind",
		10: "Odd",
		11: "Even",
		12: "Chance",
	}
	fmt.Println("These are your choices:")

	for key := 1; key < len(Options)+1; key++ {
		if numberOfRemainingRolls == 0 && key == 1 {
			continue
		}

		fmt.Printf("%d: %s", key, Options[key])

		if key >= 2 && player.TableUsed[key] {
			fmt.Printf("\tUSED")
		}
		fmt.Println()
	}
}

func (player *Player) FullTour() {
	//define map from int choice to functions
	optionFuncMap := []func([]int) int{
		2:  state.OnePair,
		3:  state.TwoPair,
		4:  state.ThreeOfKind,
		5:  state.SmallStraight,
		6:  state.BigStraight,
		7:  state.FullHouse,
		8:  state.FourOfKind,
		9:  state.FiveOfKind,
		10: state.Odd,
		11: state.Even,
		12: state.Chance,
	}

	//prep variables
	diceSlice := make([]int, 5)
	rolls := []bool{true, true, true, true, true}
	numberOfRemainingRolls := 3

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
		newStateChoice := player.acceptedChoice(numberOfRemainingRolls)

		//execute state choice
		if newStateChoice != 1 {
			player.TableScore[newStateChoice] = optionFuncMap[newStateChoice](diceSlice)    //Write score to table
			if numberOfRemainingRolls == 2 && newStateChoice >= 5 && newStateChoice <= 11 { //In first move
				player.TableScore[newStateChoice] *= 2
			}
			player.Score += player.TableScore[newStateChoice] //Update sum score
			player.TableUsed[newStateChoice] = true           //Update used table
			return
		}
	}

}

func (player *Player) acceptedChoice(numberOfRemainingRolls int) int {
	var newChoiceForMove int
	accepted := false
	for !accepted {
		newChoiceForMove = player.GiveMoveChoice()

		if newChoiceForMove > 12 || newChoiceForMove < 1 {
			fmt.Println("Action number not in range, try again")
			continue
		}
		if newChoiceForMove == 1 && numberOfRemainingRolls == 0 {
			fmt.Println("Its your last roll, you need to make a choice")
			continue
		}

		if player.TableUsed[newChoiceForMove] {
			fmt.Println("Action already used, try again")
			continue
		}
		accepted = true
	}
	return newChoiceForMove
}

func (player *Player) GiveMoveChoice() int {
	choice := 1
	if globalVar.IO_human {
		fmt.Scanln(&choice)
	} else {
		choice = ai.DecideMove(player.TableUsed)
	}
	return choice
}
