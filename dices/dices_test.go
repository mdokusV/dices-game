package dices

import "testing"

func TestDices_RoleDices(t *testing.T) {
	type fields struct {
		DiceSlice []int
		Rolls     []bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"nothing to be seen", fields{[]int{1, 2, 3, 4, 5}, []bool{true, true, true, true, true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dices := &Dices{
				DiceSlice: tt.fields.DiceSlice,
				Rolls:     tt.fields.Rolls,
			}
			dices.RoleDices()
		})
	}
}
