package globalStructs

import (
	"reflect"
	"testing"

	"github.com/mdokusV/dices-game/globalVar"
)

func TestNewPlayerGroup(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []Player
	}{
		{
			"Test with size 1", args{size: 1}, []Player{
				{
					ID:    0,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
			},
		},
		{
			"Test with size 3", args{size: 3}, []Player{
				{
					ID:    0,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
				{
					ID:    1,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
				{
					ID:    2,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlayerGroup(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerGroup() = %v, want %v\n", got, tt.want)
			}
		})
	}
}
func TestResetPlayerGroup(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []Player
	}{
		{
			"Test with size 1", args{size: 1}, []Player{
				{
					ID:    0,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
			},
		},
		{
			"Test with size 3", args{size: 3}, []Player{
				{
					ID:    0,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
				{
					ID:    1,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
				{
					ID:    2,
					Score: 0,
					TableScore: [globalVar.NumberOfStates]int{
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
						0,
					},
					TableUsed: [globalVar.NumberOfStates]bool{
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
						false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlayerGroup(tt.args.size)
			if ResetPlayerGroup(got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerGroup() = %v, want %v\n", got, tt.want)
			}
		})
	}
}

func TestPlayer_PrintPossibleChoices(t *testing.T) {
	type fields struct {
		ID         int
		Score      int
		TableScore [globalVar.NumberOfStates]int
		TableUsed  [globalVar.NumberOfStates]bool
	}
	type args struct {
		numberOfRemainingRolls int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Nothing to be seen", fields{0, 0, [globalVar.NumberOfStates]int{}, [globalVar.NumberOfStates]bool{}}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				ID:         tt.fields.ID,
				Score:      tt.fields.Score,
				TableScore: tt.fields.TableScore,
				TableUsed:  tt.fields.TableUsed,
			}
			player.PrintPossibleChoices(tt.args.numberOfRemainingRolls)
		})
	}
}
