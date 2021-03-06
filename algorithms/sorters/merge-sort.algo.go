package sorters

import "math"

// As described in Introduction to Algorithms, 3rd edition, page 31
// p, q and r are indices in the array such that p <= q < r
func merge(array []int, start int, middle int, end int) {
	n1 := middle - start + 1
	n2 := end - middle
	leftArr := make([]int, n1+1)
	rightArr := make([]int, n2+1)
	var i, j int
	for i = 0; i < n1; i++ {
		leftArr[i] = array[start+i]
	}
	for j = 1; j <= n2; j++ {
		rightArr[j-1] = array[middle+j]
	}
	leftArr[n1] = math.MaxInt
	rightArr[n2] = math.MaxInt
	i = 0
	j = 0
	for k := start; k <= end; k++ {
		if leftArr[i] <= rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
}

// Implementation of Merge Sort, as described in Introduction to Algorithms, 3rd edition, page 34
// Initial call should be MergeSort(array, 0, len(array ) - 1). p and r are indices in the array such that p < r
func MergeSort(array []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		MergeSort(array, start, middle)
		MergeSort(array, middle+1, end)
		merge(array, start, middle, end)
	}
}
