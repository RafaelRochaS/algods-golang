package operations

import (
	"algods/common"
	"algods/datastructures"
	"fmt"
)

func DListMenu() {
	dlist := datastructures.DoublyLinkedList{}
	var input int
	for {
		printMenu()
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			insertItem(&dlist)
		case 2:
			fmt.Println("Clearing list...")
			dlist.Clear()
		case 3:
			searchItem(dlist)
		}
	}
}

func insertItem(dlist *datastructures.DoublyLinkedList) {
	fmt.Print("\nType integer to insert on the list: ")
	input := common.GetIntInput()
	dlist.Insert(input)
}

func searchItem(dlist datastructures.DoublyLinkedList) {
	fmt.Print("\nType integer to search on the list: ")
	input := common.GetIntInput()
	_, err := dlist.Find(input)
	if err != nil {
		fmt.Println("Item is in the list.")
	} else {
		fmt.Println("Item is not on the list.")
	}
}

func printMenu() {
	fmt.Println("Select an operation to perform:")
	fmt.Println("0 - Return\n1 - Insert\n2 - Clear\n3 - Find\n4 - Delete\n5 - Insert Random\n6 - Traverse")
	fmt.Println("---------------------------------------------")
}
