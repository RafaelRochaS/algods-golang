package main

import (
	"algods/common"
	"algods/datastructures"
	"algods/operations"
	"fmt"
)

func main() {

	fmt.Println("********************************************************")
	fmt.Println("Data Structures and Algorithms - Tests and Experiments")
	fmt.Println("********************************************************")

	for {
		printMenu()
		input := common.GetIntInput()
		switch input {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			operations.DListMenu(&datastructures.DoublyLinkedList{})
		case 2:
			operations.DequeMenu(&datastructures.Deque{})
		case 3:
			operations.PeakFindingMenu()
		case 4:
			operations.InSortMenu()
		case 5:
			operations.FindMaxMenu()
		case 6:
			operations.MergeSortMenu()
		case 7:
			operations.HeapMenu()
		case 8:
			operations.AVLMenu()
		case 9:
			operations.HashTableMenu()
		}
	}
}

func printMenu() {
	fmt.Println("\nWhich Data Structure or Algorthim do you want to try out?")
	fmt.Println(`0 - Exit program
1 - Doubly Linked List - Default Implementation
2 - Deque - Default Implementation
3 - Peak Finding
4 - Insertion Sort
5 - Find maximum subarray
6 - Merge Sort
7 - Heaps
8 - AVL Trees
9 - Hash Tables`)
	fmt.Println("---------------------------------------------")
}
