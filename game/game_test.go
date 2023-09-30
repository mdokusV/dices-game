package game

import "testing"

func Test_allMovesCompleted(t *testing.T) {
	type args struct {
		states map[int]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"All moves completed", args{map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}}, true},
		{"Not all moves completed", args{map[int]bool{1: true, 2: true, 3: true, 4: true, 5: false}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allMovesCompleted(tt.args.states); got != tt.want {
				t.Errorf("allMovesCompleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
