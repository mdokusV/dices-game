package ai

import (
	"testing"

	"github.com/mdokusV/dices-game/globalVar"
)

func TestDecideMove(t *testing.T) {
	type args struct {
		tableUsed [globalVar.NumberOfStates]bool
		diceSlice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecideMove(tt.args.tableUsed, tt.args.diceSlice); got != tt.want {
				t.Errorf("DecideMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
