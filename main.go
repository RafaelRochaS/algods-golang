package main

import (
	"fmt"
)

func main() {

	fmt.Println("********************************************************")
	fmt.Println("Data Structures and Algorithms - Tests and Experiments")
	fmt.Println("********************************************************")

	var input int

	for { // Add error handling for non-int input
		printMenu()
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			fmt.Println("Working")
		}
	}
}

func printMenu() {
	fmt.Println("Which Data Structure or Algorthim do you want to try out?")
	fmt.Println("0 - Exit program\n1 - Doubly Linked List")
}
