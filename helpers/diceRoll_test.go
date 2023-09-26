package helpers

import "testing"

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
