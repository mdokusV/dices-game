package action

import (
	"fmt"
)

func GiveMoveChoice() int {
	choice := 2
	fmt.Scanln(&choice)
	return choice
}
