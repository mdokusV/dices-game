package action

import "fmt"

func GiveMoveChoice() int {
	choice := 1
	fmt.Scanln(&choice)
	return choice
}
