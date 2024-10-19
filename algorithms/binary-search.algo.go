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

func BinarySearchRecursive(array []int, item int) int {
	return binaryRecursive(array, item, 0, len(array)-1)
}

func binaryRecursive(array []int, item, low, high int) int {
	mid := (low + high) / 2
	if array[mid] == item {
		return mid
	}

	if array[mid] > item {
		return binaryRecursive(array, item, low, mid-1)
	} else {
		return binaryRecursive(array, item, mid+1, high)
	}
}
