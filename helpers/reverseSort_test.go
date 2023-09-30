package helpers

import (
	"reflect"
	"testing"
)

func TestReverseSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Some same", args{[]int{3, 2, 4, 3, 2}}, []int{4, 3, 3, 2, 2}},
		{"All different", args{[]int{2, 4, 5, 3, 6}}, []int{6, 5, 4, 3, 2}},
		{"All same", args{[]int{4, 4, 4, 4, 4}}, []int{4, 4, 4, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSort(tt.args.slice)
			if got := tt.args.slice; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSort2(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Some same", args{[]int{3, 2, 4, 3, 2}}, []int{4, 3, 3, 2, 2}},
		{"All different", args{[]int{2, 4, 5, 3, 6}}, []int{6, 5, 4, 3, 2}},
		{"All same", args{[]int{4, 4, 4, 4, 4}}, []int{4, 4, 4, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSort2(tt.args.slice)
			if got := tt.args.slice; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
