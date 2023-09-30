package state

import (
	"testing"
)

func TestOnePair(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{5, 4, 3, 2, 1}}, 0},
		{"One Start", args{[]int{6, 6, 3, 2, 1}}, 12},
		{"One End", args{[]int{6, 5, 4, 1, 1}}, 2},
		{"More", args{[]int{5, 4, 4, 4, 2}}, 8},
		{"All", args{[]int{1, 1, 1, 1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnePair(tt.args.currentState); got != tt.want {
				t.Errorf("OnePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoPair(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{4, 4, 3, 2, 1}}, 0},
		{"Start + End", args{[]int{6, 6, 3, 1, 1}}, 14},
		{"Same Number", args{[]int{6, 6, 6, 6, 1}}, 6 * 4},
		{"All", args{[]int{3, 3, 3, 3, 3}}, 3 * 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoPair(tt.args.currentState); got != tt.want {
				t.Errorf("TwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeOfKind(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{4, 4, 3, 2, 1}}, 0},
		{"Start", args{[]int{6, 6, 6, 1, 1}}, 6 * 3},
		{"ENd", args{[]int{5, 5, 1, 1, 1}}, 1 * 3},
		{"All", args{[]int{3, 3, 3, 3, 3}}, 3 * 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeOfKind(tt.args.currentState); got != tt.want {
				t.Errorf("ThreeOfKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallStraight(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{5, 4, 3, 1, 1}}, 0},
		{"SmallStraight", args{[]int{5, 4, 3, 2, 1}}, 15},
		{"BigStraight", args{[]int{6, 5, 4, 3, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallStraight(tt.args.currentState); got != tt.want {
				t.Errorf("SmallStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigStraight(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{6, 4, 3, 2, 1}}, 0},
		{"SmallStraight", args{[]int{5, 4, 3, 2, 1}}, 0},
		{"BigStraight", args{[]int{6, 5, 4, 3, 2}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigStraight(tt.args.currentState); got != tt.want {
				t.Errorf("BigStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullHouse(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{4, 4, 3, 2, 1}}, 0},
		{"SmallLarge", args{[]int{5, 5, 3, 3, 3}}, 5*2 + 3*3},
		{"LargeSmall", args{[]int{5, 5, 5, 3, 3}}, 5*3 + 3*2},
		{"ALl", args{[]int{6, 6, 6, 6, 6}}, 6 * 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullHouse(tt.args.currentState); got != tt.want {
				t.Errorf("FullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourOfKind(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{4, 4, 4, 2, 1}}, 0},
		{"Start", args{[]int{6, 6, 6, 6, 1}}, 6 * 4},
		{"ENd", args{[]int{5, 1, 1, 1, 1}}, 1 * 4},
		{"All", args{[]int{3, 3, 3, 3, 3}}, 3 * 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourOfKind(tt.args.currentState); got != tt.want {
				t.Errorf("FourOfKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiveOfKind(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"None", args{[]int{4, 4, 4, 2, 1}}, 0},
		{"All", args{[]int{3, 3, 3, 3, 3}}, 3 * 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiveOfKind(tt.args.currentState); got != tt.want {
				t.Errorf("FiveOfKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"There is Even", args{[]int{4, 4, 4, 2, 1}}, 0},
		{"All Odd", args{[]int{3, 1, 5, 3, 3}}, 3 + 1 + 5 + 3 + 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Odd(tt.args.currentState); got != tt.want {
				t.Errorf("Odd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEven(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"There is Odd", args{[]int{4, 4, 4, 2, 1}}, 0},
		{"All Even", args{[]int{2, 2, 4, 6, 2}}, 2 + 2 + 4 + 6 + 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Even(tt.args.currentState); got != tt.want {
				t.Errorf("Even() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChance(t *testing.T) {
	type args struct {
		currentState []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Slice one", args{[]int{5, 3, 2, 2, 1}}, 5 + 3 + 2 + 2 + 1},
		{"Slice two", args{[]int{2, 2, 4, 6, 2}}, 2 + 2 + 4 + 6 + 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chance(tt.args.currentState); got != tt.want {
				t.Errorf("Chance() = %v, want %v", got, tt.want)
			}
		})
	}
}
