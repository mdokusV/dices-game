package game

import (
	"testing"

	"github.com/mdokusV/dices-game/globalVar"
)

func Test_allMovesCompleted(t *testing.T) {
	type args struct {
		states [globalVar.NumberOfStates]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"All moves completed", args{[globalVar.NumberOfStates]bool{true, true, true, true, true, true, true, true, true, true, true}}, true},
		{"All moves completed", args{[globalVar.NumberOfStates]bool{true, true, true, true, true, true, true, true, false, true, true}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allMovesCompleted(tt.args.states); got != tt.want {
				t.Errorf("allMovesCompleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
