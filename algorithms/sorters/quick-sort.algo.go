package sorters

// Hoare's partition scheme
func Quicksort(arr []int) []int {
	quicksort(arr, 0, len(arr)-1)

	return arr
}

func quicksort(arr []int, low, high int) {
	if low >= high || low < 0 || high < 0 {
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

	for {
		for {
			left += 1
			if arr[left] >= pivot {
				break
			}
		}

		for {
			right -= 1
			if arr[right] <= pivot {
				break
			}
		}

		if left >= right {
			return right
		}

		swap(arr, left, right)
	}

}

func swap(arr []int, index1, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}
