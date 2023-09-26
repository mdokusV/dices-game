package helpers

import "testing"

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
