package main

// Function ArgMax takes a slice of numbers
// Returns the maximum item, and the index of the maximum item
func ArgMax(numbers []int) (int, int) {
	maxItem, _, maxIndex := ArgMaxFunc(numbers, func(x int) int { return x })
	return maxItem, maxIndex
}

// Function ArgMin takes a slice of numbers
// Returns the minimum item, and the index of the minimum item
func ArgMin(numbers []int) (int, int) {
	minItem, _, minIndex := ArgMaxFunc(numbers, func(x int) int { return -x })
	return minItem, minIndex
}

// Function ArgMaxFunc takes a slice of numbers
// Returns the maximum item, the maximum value, and the index of the maximum item, according to the given value function.
func ArgMaxFunc(numbers []int, valueFunc func(int) int) (int, int, int) {
	if len(numbers) == 0 {
		return 0, 0, -1
	} else {
		var maxValue int
		var maxItem int
		var maxIndex int

		for index, item := range numbers {
			value := valueFunc(item)
			if index == 0 || value > maxValue {
				maxValue = value
				maxItem = item
				maxIndex = index
			}
		}

		return maxItem, maxValue, maxIndex
	}
}

// Function ArgMaxFunc takes a slice of numbers
// Returns the minimum item, the minimum value, and the index of the minimum item, according to the given value function.
func ArgMinFunc(numbers []int, valueFunc func(int) int) (int, int, int) {
	return ArgMaxFunc(numbers, func(x int) int { return -valueFunc(x) })
}
