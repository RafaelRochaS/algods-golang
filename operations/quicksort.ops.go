package operations

import (
	"algods/algorithms/sorters"
	"algods/common"
	"algods/utils"
	"fmt"
)

func QuicksortMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Manual array input
2 - Random array input
	`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			fmt.Println("Array sorted: ", sorters.Quicksort(utils.GetArrayWithManualInput()))
		case 2:
			fmt.Println("Array sorted: ", sorters.Quicksort(utils.GetArrayWithRandomInput()))
		}
	}
}
