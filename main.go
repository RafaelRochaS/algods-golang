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
		}
	}
}

func printMenu() {
	fmt.Println("\nWhich Data Structure or Algorthim do you want to try out?")
	fmt.Println(`0 - Exit program
1 - Doubly Linked List - Default Implementation
2 - Deque - Default Implementation`)
	fmt.Println("---------------------------------------------")
}
