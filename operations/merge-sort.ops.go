package operations

import (
	"algods/algorithms/sorters"
	"algods/common"
	"algods/utils"
	"fmt"
)

func MergeSortMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Manual array input
2 - Random array input
3 - Book example array`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			sort(utils.GetArrayWithManualInput())
		case 2:
			sort(utils.GetArrayWithRandomInput())
		case 3:
			bookExampleArray()
		}
	}

}

func sort(arr []int) {
	sorters.MergeSort(arr, 0, len(arr)-1)
	fmt.Println("Array sorted: ", arr)
}

func bookExampleArray() {
	array := []int{2, 4, 5, 7, 1, 2, 3, 6}
	fmt.Printf("Current array: %v", array)
	sorters.MergeSort(array, 0, len(array)-1)
	fmt.Printf("Sorted array: %v", array)
}
