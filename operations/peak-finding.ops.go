package operations

import (
	"algods/algorithms"
	"algods/common"
	"fmt"
)

func PeakFindingMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - 1D Random
2 - 2D Random`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			randomPeak1D()
		case 2:
			randomPeak2D()
		}
	}
}

func randomPeak1D() {
	arr := []int{6, 7, 4, 3, 2, 1, 4, 5}
	fmt.Println(algorithms.PeakFinding1D(arr))
}

func randomPeak2D() {
	matrix := [][]int{
		{1, 2, 3},
		{19, 9, 10},
		{7, 8, 9},
		{10, 7, 1},
		{13, 2, 11}}
	result, _ := algorithms.PeakFinding2D(matrix)
	fmt.Println(result)
}
