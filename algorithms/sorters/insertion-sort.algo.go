package sorters

// Implementation of Insertion Sort, as described in Introduction to Algorthims, 3rd edition, page 18
func InsertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i -= 1
		}
		arr[i+1] = key
	}
	return arr
}
