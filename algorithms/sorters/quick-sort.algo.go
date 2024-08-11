package sorters

// Hoare's partition scheme
func Quicksort(arr []int) []int {
	quicksort(arr, 0, len(arr)-1)

	return arr
}

func quicksort(arr []int, low, high int) {
	if low > high {
		return
	}

	partitioningIndex := partition(arr, low, high)

	quicksort(arr, low, partitioningIndex)
	quicksort(arr, partitioningIndex+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]

	left := low - 1
	right := high + 1

	for running := true; running; running = (arr[left] < pivot) {
		left += 1
	}

	for running := true; running; running = (arr[right] > pivot) {
		right += 1
	}

	if left >= right {
		return right
	}

	swap(arr, left, right)

	return left
}

func swap(arr []int, index1, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}
