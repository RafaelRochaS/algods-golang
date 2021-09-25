// Collection of different algorithms (interfaces and implementations) to be applied to
// the data structures, or be used stand alone
package algorithms

import "errors"

// Algorithms to solve the Peak Finding problem - 1D version
func PeakFinding1D(arr []int) int {
	index := (len(arr) / 2)
	if index != 0 && arr[index] < arr[index-1] {
		return PeakFinding1D(arr[:index])
	} else if index != len(arr)-1 && arr[index] < arr[index+1] {
		return PeakFinding1D(arr[index+1:])
	} else {
		return arr[index]
	}
}

// Helper function to grab index of value in array. Kills performance, since in average case runs in O(n)
func indexOf(value int, arr []int) (int, error) {
	for k, v := range arr {
		if v == value {
			return k, nil
		}
	}
	return -1, errors.New("item not found")
}

func PeakFinding2D(matrix [][]int) (int, error) {
	middle := len(matrix) / 2
	// the line below prevents the algorithm from running in O(log(nm)), since finding the peak in a column returns a value, not an index, and it is necessary to get the index of that value.
	// To fix that, peak finding 1D must be ran inside this function, which would make a big mess. The whole algorithm would have to be ran in a loop, instead of recursively.
	if peak1D, err := indexOf(PeakFinding1D(matrix[middle]), matrix[middle]); err == nil {
		if middle != 0 && matrix[middle-1][peak1D] > matrix[middle][peak1D] {
			return PeakFinding2D(matrix[0:middle])
		} else if middle != len(matrix)-1 && matrix[middle+1][peak1D] > matrix[middle][peak1D] {
			return PeakFinding2D(matrix[middle+1:])
		} else {
			return matrix[middle][peak1D], nil
		}
	}
	return -1, errors.New("failed to compute")
}
