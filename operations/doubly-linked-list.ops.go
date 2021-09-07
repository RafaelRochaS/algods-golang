// Collection of interactive operations to perform on data structures and algorithms.
// Operates directly on the ADTs (interfaces), decoupling from implementation
package operations

import (
	"algods/common"
	"algods/interfaces"
	"fmt"
	"math/rand"
)

const MAX_TRAVERSE_LENGTH = 50
const MAX_RANDOM_VALUE = 1_000

// Initiates the interactive menu for operating on a Doubly Linked List ADT
func DListMenu(dlist interfaces.IDoublyLinkedList) {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Insert
2 - Clear
3 - Find
4 - Delete
5 - Insert Random
6 - Traverse
7 - Length`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			insertItem(dlist)
		case 2:
			fmt.Println("\nClearing list...")
			dlist.Clear()
		case 3:
			searchItem(dlist)
		case 4:
			deleteItem(dlist)
		case 5:
			insertRandom(dlist)
		case 6:
			fmt.Println("\nTraversing...")
			dlist.Traverse()
		case 7:
			fmt.Println("\nCurrent length of list: ", dlist.Length())
		}
	}
}

func insertItem(dlist interfaces.IDoublyLinkedList) {
	fmt.Print("\nType integer to insert on the list: ")
	input := common.GetIntInput()
	dlist.Insert(input)
}

func searchItem(dlist interfaces.IDoublyLinkedList) {
	fmt.Print("\nType integer to search on the list: ")
	input := common.GetIntInput()
	_, err := dlist.Find(input)
	if err != nil {
		fmt.Println("Item is in the list.")
	} else {
		fmt.Println("Item is not on the list.")
	}
}

func deleteItem(dlist interfaces.IDoublyLinkedList) {
	fmt.Print("\nType integer to delete from the list: ")
	input := common.GetIntInput()
	_, err := dlist.Delete(input)
	if err != nil {
		fmt.Println("Item deleted.")
	} else {
		fmt.Println("Item not found on the list.")
	}
}

func insertRandom(dlist interfaces.IDoublyLinkedList) {
	fmt.Print("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	for i := 0; i < input; i++ {
		dlist.Insert(rand.Intn(MAX_RANDOM_VALUE))
	}
	fmt.Println("Done.")
}
