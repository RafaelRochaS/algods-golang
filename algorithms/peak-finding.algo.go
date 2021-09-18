// Collection of different algorithms (interfaces and implementations) to be applied to
// the data structures, or be used stand alone
package algorithms

// Algorithms to solve the Peak Finding problem - 1D version
func PeakFinding1D(arr []int) int {
	index := (len(arr) / 2)
	if arr[index] < arr[index-1] {
		return PeakFinding1D(arr[:index+1])
	} else if arr[index] < arr[index+1] {
		return PeakFinding1D(arr[index:])
	} else {
		return arr[index]
	}
}
