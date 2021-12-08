package operations

import (
	"algods/common"
	"algods/datastructures"
	"fmt"
)

func AVLMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Test Insert`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			testInsert()
		}
	}
}

func testInsert() {
	arr := []int{5, 9, 10, 2, 45, 8, 13, 10}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}
