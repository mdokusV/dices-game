package state

const NUMBEROFDICES = 5

func OnePair(currentState []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range currentState {
		frequencyMap[num]++
		if frequencyMap[num] == 2 {
			return num * 2
		}
	}
	return 0
}

func TwoPair(currentState []int) int {
	frequencyMap := make(map[int]int)
	pairs := 0
	collectedPoints := 0

	for _, nums := range currentState {
		frequencyMap[nums]++
		if frequencyMap[nums] == 2 {
			pairs++
			frequencyMap[nums] = 0
			collectedPoints += nums * 2
		}
		if pairs == 2 {
			return collectedPoints
		}
	}
	return 0
}

func ThreeOfKind(currentState []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range currentState {
		frequencyMap[num]++
		if frequencyMap[num] == 3 {
			return num * 3
		}
	}
	return 0
}

func SmallStraight(currentState []int) int {
	if currentState[0] != 5 {
		return 0
	}
	for i := 0; i < NUMBEROFDICES-1; i++ {
		if currentState[i]-currentState[i+1] != 1 {
			return 0
		}
	}
	return 1 + 2 + 3 + 4 + 5
}

func BigStraight(currentState []int) int {
	if currentState[0] != 6 {
		return 0
	}
	for i := 0; i < NUMBEROFDICES-1; i++ {
		if currentState[i]-currentState[i+1] != 1 {
			return 0
		}
	}
	return 2 + 3 + 4 + 5 + 6
}

func FullHouse(currentState []int) int {
	frequencyMap := make(map[int]int)
	collectedPoints := 0

	for _, nums := range currentState {
		frequencyMap[nums]++
	}
	for item, amount := range frequencyMap {
		if amount == 5 {
			collectedPoints += 5 * item
		} else if amount == 3 {
			collectedPoints += 3 * item
		} else if amount == 2 {
			collectedPoints += 2 * item
		} else {
			return 0
		}

	}
	return collectedPoints
}

func FourOfKind(currentState []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range currentState {
		frequencyMap[num]++
		if frequencyMap[num] == 4 {
			return num * 4
		}
	}
	return 0
}

func FiveOfKind(currentState []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range currentState {
		frequencyMap[num]++
		if frequencyMap[num] == 5 {
			return num * 5
		}
	}
	return 0
}

func Odd(currentState []int) int {
	sumAll := 0
	for _, value := range currentState {
		if value%2 == 0 {
			return 0
		}
		sumAll += value
	}
	return sumAll
}

func Even(currentState []int) int {
	sumAll := 0
	for _, value := range currentState {
		if value%2 == 1 {
			return 0
		}
		sumAll += value
	}
	return sumAll
}

func Chance(currentState []int) int {
	sumAll := 0
	for _, value := range currentState {
		sumAll += value
	}
	return sumAll
}
