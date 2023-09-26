package game

import (
	"reflect"
	"testing"
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
					TableScore: map[int]int{
						2:  0,
						3:  0,
						4:  0,
						5:  0,
						6:  0,
						7:  0,
						8:  0,
						9:  0,
						10: 0,
						11: 0,
						12: 0,
					},
					TableUsed: map[int]bool{
						2:  false,
						3:  false,
						4:  false,
						5:  false,
						6:  false,
						7:  false,
						8:  false,
						9:  false,
						10: false,
						11: false,
						12: false,
					},
				},
			},
		},
		{
			"Test with size 3", args{size: 3}, []Player{
				{
					ID:    0,
					Score: 0,
					TableScore: map[int]int{
						2:  0,
						3:  0,
						4:  0,
						5:  0,
						6:  0,
						7:  0,
						8:  0,
						9:  0,
						10: 0,
						11: 0,
						12: 0,
					},
					TableUsed: map[int]bool{
						2:  false,
						3:  false,
						4:  false,
						5:  false,
						6:  false,
						7:  false,
						8:  false,
						9:  false,
						10: false,
						11: false,
						12: false,
					},
				},
				{
					ID:    1,
					Score: 0,
					TableScore: map[int]int{
						2:  0,
						3:  0,
						4:  0,
						5:  0,
						6:  0,
						7:  0,
						8:  0,
						9:  0,
						10: 0,
						11: 0,
						12: 0,
					},
					TableUsed: map[int]bool{
						2:  false,
						3:  false,
						4:  false,
						5:  false,
						6:  false,
						7:  false,
						8:  false,
						9:  false,
						10: false,
						11: false,
						12: false,
					},
				},
				{
					ID:    2,
					Score: 0,
					TableScore: map[int]int{
						2:  0,
						3:  0,
						4:  0,
						5:  0,
						6:  0,
						7:  0,
						8:  0,
						9:  0,
						10: 0,
						11: 0,
						12: 0,
					},
					TableUsed: map[int]bool{
						2:  false,
						3:  false,
						4:  false,
						5:  false,
						6:  false,
						7:  false,
						8:  false,
						9:  false,
						10: false,
						11: false,
						12: false,
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

func TestPrintAllChoices(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Nothing to be seen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAllChoices()
		})
	}
}

func TestRoleDices(t *testing.T) {
	type args struct {
		currentState []int
		whatRoll     []bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"nothing to be seen", args{[]int{1, 2, 3, 4, 5}, []bool{true, true, true, true, true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RoleDices(tt.args.currentState, tt.args.whatRoll)
		})
	}
}
