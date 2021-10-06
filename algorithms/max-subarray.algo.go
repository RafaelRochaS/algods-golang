package algorithms

import "math"

// As described in Introduction to Algorithms, 3rd edition, page 72
func FindMaximumSubarray(array []int, low int, high int) (int, int, int) {
	if high == low { // base case
		return low, high, array[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaximumSubarray(array, low, mid)
		rightLow, rightHigh, rightSum := FindMaximumSubarray(array, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubarray(array, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

// As described in Introduction to Algorithms, 3rd edition, page 71
func findMaxCrossingSubarray(array []int, low int, mid int, high int) (int, int, int) {
	leftSum := int(math.Inf(-1))
	rightSum := int(math.Inf(+1))
	sum := 0
	var maxLeft int
	var maxRight int

	for i := mid; i >= low; i-- {
		sum += array[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0

	for j := mid + 1; j <= high; j++ {
		sum += array[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}
