package operations

import (
	"algods/common"
	"algods/datastructures"
	"fmt"
)

func HeapMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Test Max Heapify - Book`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			testMaxHeapifyBook()
		}
	}
}

func testMaxHeapifyBook() {
	arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	heap := datastructures.Heap{Keys: arr, HeapSize: 10}
	fmt.Println("Current heap:")
	fmt.Println(heap.Keys)
	fmt.Println("Calling MaxHeapify on index 1...")
	heap.MaxHeapify(1)
	fmt.Println("Heap now:")
	fmt.Println(heap.Keys)
}
