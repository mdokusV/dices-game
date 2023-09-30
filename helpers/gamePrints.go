package helpers

import (
	"fmt"

	"github.com/mdokusV/dices-game/globalVar"
)

func PrintAllChoices() {
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
		fmt.Printf("%d: %s\n", key+1, Options[key])
	}
}
