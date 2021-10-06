package operations

import (
	"algods/algorithms"
	"algods/common"
	"fmt"
)

func FindMaxMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Default input`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
			low, high, sum := algorithms.FindMaximumSubarray(arr, 0, len(arr)-1)
			fmt.Printf("Got a low index of '%d', a high index of '%d', and a sum of '%d'", low, high, sum)
		}
	}
}
