package algorithms

// Searches for an item in the given array in O(log n). Assumes array is sorted. Returns -1 if item is not found.
func BinarySearch(array []int, item int) int {
	var mid int
	var guess int
	var low int
	var high int

	high = len(array) - 1

	for low <= high {
		mid = (low + high) / 2
		guess = array[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
