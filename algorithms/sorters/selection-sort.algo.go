package sorters

func SelectionSort(array []int) []int {
	var smallestIndex int

	for i := range len(array) {
		smallestIndex = findSmallestIndex(array)
		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}

	return array
}

func findSmallestIndex(array []int) int {
	var smallestItem int
	var smallestItemIndex int

	for i, v := range array {
		if v < smallestItem {
			smallestItem = v
			smallestItemIndex = i
		}
	}

	return smallestItemIndex
}
