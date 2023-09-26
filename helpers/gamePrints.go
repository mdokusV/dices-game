package helpers

import "fmt"

func PrintAllChoices() {
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

	for key := 0; key < len(Options); key++ {
		fmt.Printf("%d: %s\n", key+1, Options[key+1])
	}
}
