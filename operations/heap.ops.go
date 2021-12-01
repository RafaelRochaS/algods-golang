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
1 - Test Max Heapify - Book
2 - Test Build Max Heap - Book`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			testMaxHeapifyBook()
		case 2:
			testBuildMaxHeapBook()
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

func testBuildMaxHeapBook() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap := datastructures.Heap{Keys: arr, HeapSize: 10}
	fmt.Println("Current heap:")
	fmt.Println(heap.Keys)
	fmt.Println("Calling BuildMaxHeap...")
	datastructures.BuildMaxHeap(arr)
	fmt.Println("Heap now:")
	fmt.Println(heap.Keys)
}
