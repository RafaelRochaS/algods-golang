package algorithms

import (
	"math/rand"
)

func Quickselect(arr []int, kth int) int {
	return runSelect(arr, 0, len(arr)-1, kth)
}

func runSelect(arr []int, left, right, kth int) int {
	if left == right {
		return arr[left]
	}

	pivotIndex := left + (rand.Intn(10) % (right - left + 1))
	pivotIndex = partition(arr, left, right, pivotIndex)

	if kth == pivotIndex {
		return arr[pivotIndex]
	} else if kth < pivotIndex {
		return runSelect(arr, left, pivotIndex-1, kth)
	} else {
		return runSelect(arr, pivotIndex+1, right, kth)
	}
}

func partition(arr []int, low, high, pivotIndex int) int {
	pivot := arr[pivotIndex]

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
