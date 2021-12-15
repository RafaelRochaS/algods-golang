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
1 - Test Insert
2 - Test Right Rotate
3 - Test Left Rotate
4 - Test Left Right Rotate
5 - Test Right Left Rotate
6 - Test Delete`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			testInsert()
		case 2:
			testRightRotate()
		case 3:
			testLeftRotate()
		case 4:
			testLeftRightRotate()
		case 5:
			testRightLeftRotate()
		case 6:
			testDelete()
		}
	}
}

func testInsert() {
	arr := []int{5, 9, 10, 2, 45, 8, 13, 10, 4, 6, 16}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}

func testRightRotate() {
	arr := []int{10, 11, 12}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}

func testLeftRotate() {
	arr := []int{10, 7, 6}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}

func testLeftRightRotate() {
	arr := []int{10, 7, 8}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}

func testRightLeftRotate() {
	arr := []int{10, 12, 11}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
}

func testDelete() {
	arr := []int{5, 9, 10, 2, 45, 8}
	avl := &datastructures.AVLTree{}
	for i := 0; i < len(arr); i++ {
		avl.Insert(arr[i])
	}
	avl.Delete(5)
	avl.Delete(3)
	avl.Delete(2)
}
