package main

import (
	"algods/common"
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
			operations.DListMenu()
		case 2:
			operations.DequeMenu()
		}
	}
}

func printMenu() {
	fmt.Println("\nWhich Data Structure or Algorthim do you want to try out?")
	fmt.Println("0 - Exit program\n1 - Doubly Linked List\n2 - Deque")
	fmt.Println("---------------------------------------------")
}
