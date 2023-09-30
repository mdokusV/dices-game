package ai

func DecideMove(tableUsed map[int]bool) int {
	for i, v := range tableUsed {
		if !v {
			return i
		}
	}
	return 0
}
