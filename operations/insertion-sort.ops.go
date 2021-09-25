package operations

import (
	"algods/algorithms/sorters"
	"algods/common"
	"fmt"
	"math/rand"
	"time"
)

func InSortMenu() {
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
			manualInput()
		case 2:
			randomInput()
		}
	}
}

func manualInput() {
	fmt.Print("Size of array: ")
	var input int
	fmt.Scanf("%d", &input)
	arr := make([]int, 0, input)
	for i := 1; i <= input; i++ {
		var num int
		fmt.Print("Insert number: ")
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	fmt.Println("Array inserted: ", arr)
	fmt.Println("Array sorted: ", sorters.InsertionSort(arr))
}

func randomInput() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	arr := make([]int, 0, input)
	for i := 0; i < input; i++ {
		arr = append(arr, rand.Intn(common.MAX_RANDOM_VALUE))
	}
	fmt.Println("Array inserted: ", arr)
	fmt.Println("Array sorted: ", sorters.InsertionSort(arr))
}
