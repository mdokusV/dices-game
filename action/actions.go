package action

import (
	"fmt"

	"github.com/mdokusV/dices-game/globalVar"
)

func GiveMoveChoice() int {
	choice := 1
	if globalVar.IO_human {
		fmt.Scanln(&choice)
	}

	return choice
}
