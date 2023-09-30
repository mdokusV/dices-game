package helpers

import "sort"

func ReverseSort(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

func ReverseSort2(slice []int) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for j >= 0 && slice[j] < key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}
