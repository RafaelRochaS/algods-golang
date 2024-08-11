package utils

import (
	"algods/common"
	"fmt"
	"math/rand"
)

func GetArrayWithManualInput() []int {
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
	return arr
}

func GetArrayWithRandomInput() []int {
	fmt.Print("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	arr := make([]int, 0, input)
	for i := 0; i < input; i++ {
		arr = append(arr, rand.Intn(common.MAX_RANDOM_VALUE))
	}
	fmt.Println("Array inserted: ", arr)

	return arr
}
