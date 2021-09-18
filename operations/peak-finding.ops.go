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
1 - 1D Random`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			randomPeak1D()
		}
	}
}

func randomPeak1D() {
	arr := []int{6, 7, 4, 3, 2, 1, 4, 5}
	fmt.Println(algorithms.PeakFinding1D(arr))
}
