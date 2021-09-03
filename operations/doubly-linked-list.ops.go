// Collection of interactive operations to perform on data structures and algorithms
package operations

import (
	"algods/common"
	"algods/datastructures"
	"fmt"
	"math/rand"
)

const MAX_TRAVERSE_LENGTH = 50
const MAX_RANDOM_VALUE = 1_000

// Initiates the interactive menu for operating a doubly-linked list
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
			fmt.Println("\nClearing list...")
			dlist.Clear()
		case 3:
			searchItem(dlist)
		case 4:
			deleteItem(&dlist)
		case 5:
			insertRandom(&dlist)
		case 6:
			fmt.Println("\nTraversing...")
			dlist.Traverse()
		case 7:
			fmt.Println("\nCurrent length of list: ", dlist.Length())
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

func deleteItem(dlist *datastructures.DoublyLinkedList) {
	fmt.Print("\nType integer to delete from the list: ")
	input := common.GetIntInput()
	_, err := dlist.Delete(input)
	if err != nil {
		fmt.Println("Item deleted.")
	} else {
		fmt.Println("Item not found on the list.")
	}
}

func insertRandom(dlist *datastructures.DoublyLinkedList) {
	fmt.Print("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	for i := 0; i < input; i++ {
		dlist.Insert(rand.Intn(MAX_RANDOM_VALUE))
	}
	fmt.Println("Done.")

}

func printMenu() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("Select an operation to perform:")
	fmt.Println("0 - Return\n1 - Insert\n2 - Clear\n3 - Find\n4 - Delete\n5 - Insert Random\n6 - Traverse\n7 - Length")
	fmt.Println("---------------------------------------------")
}
