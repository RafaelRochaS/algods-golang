package operations

import (
	"algods/algorithms/sorters"
	"algods/common"
	"fmt"
	"math/rand"
	"time"
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
			fmt.Println("Error: not implemented")
		case 2:
			randomInputMerge()
		case 3:
			bookExampleArray()
		}
	}

}

func bookExampleArray() {
	array := []int{2, 4, 5, 7, 1, 2, 3, 6}
	fmt.Printf("Current array: %v", array)
	sorters.MergeSort(array, 0, len(array)-1)
	fmt.Printf("Sorted array: %v", array)
}

func randomInputMerge() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	arr := make([]int, 0, input)
	for i := 0; i < input; i++ {
		arr = append(arr, rand.Intn(common.MAX_RANDOM_VALUE))
	}
	fmt.Println("Array inserted: ", arr)
	sorters.MergeSort(arr, 0, len(arr)-1)
	fmt.Println("Array sorted: ", arr)
}
