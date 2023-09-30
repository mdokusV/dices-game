package state

const NUMBEROFDICES = 5

func OnePair(currentState []int) int {
	var frequency [7]uint8
	for _, num := range currentState {
		frequency[num]++
		if frequency[num] == 2 {
			return num * 2
		}
	}
	return 0
}

func TwoPair(currentState []int) int {
	var frequency [7]uint8
	pairs := 0
	collectedPoints := 0

	for _, nums := range currentState {
		frequency[nums]++
		if frequency[nums] == 2 {
			pairs++
			frequency[nums] = 0
			collectedPoints += nums * 2
		}
		if pairs == 2 {
			return collectedPoints
		}
	}
	return 0
}

func ThreeOfKind(currentState []int) int {
	var frequency [7]uint8
	for _, num := range currentState {
		frequency[num]++
		if frequency[num] == 3 {
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
	var frequencyArray [7]uint8
	collectedPoints := 0

	for _, nums := range currentState {
		frequencyArray[nums]++
	}
	for item, amount := range frequencyArray {
		if amount == 5 {
			collectedPoints += 5 * item
		} else if amount == 3 {
			collectedPoints += 3 * item
		} else if amount == 2 {
			collectedPoints += 2 * item
		} else if amount == 0 {
			continue
		} else {
			return 0
		}

	}
	return collectedPoints
}

func FourOfKind(currentState []int) int {
	var frequency [7]uint8
	for _, num := range currentState {
		frequency[num]++
		if frequency[num] == 4 {
			return num * 4
		}
	}
	return 0
}

func FiveOfKind(currentState []int) int {
	var frequency [7]uint8
	for _, num := range currentState {
		frequency[num]++
		if frequency[num] == 5 {
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
