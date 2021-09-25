package operations

import (
	"algods/common"
	"algods/interfaces"
	"fmt"
	"math/rand"
	"time"
)

// Initiates the interactive menu for operating on a Deque ADT
func DequeMenu(dq interfaces.IDeque) {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Insert Front
2 - Insert Rear
3 - Clear
4 - Traverse
5 - Delete Front
6 - Delete Rear
7 - Contains
8 - Length
9 - Peek Front
10 - Peek Rear
11 - Insert Random
12 - Deque Length`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			insertFront(dq)
		case 2:
			insertRear(dq)
		case 3:
			dq.Clear()
			fmt.Println("Done.")
		case 4:
			dq.Traverse()
		case 5:
			if item, err := dq.DeleteFront(); err != nil {
				fmt.Println("\nRemoved item ", item)
			} else {
				fmt.Println("Failed to remove item: ", err)
			}
		case 6:
			if item, err := dq.DeleteRear(); err != nil {
				fmt.Println("\nRemoved item ", item)
			} else {
				fmt.Println("Failed to remove item: ", err)
			}
		case 7:
			contains(dq)
		case 8:
			fmt.Println(dq.Length())
		case 9:
			fmt.Println(dq.PeekFront())
		case 10:
			fmt.Println(dq.PeekRear())
		case 11:
			insertRandomDq(dq)
		case 12:
			fmt.Println("\nCurrent size of the deque: ", dq.Length())
		}
	}
}

func insertFront(dq interfaces.IDeque) {
	fmt.Print("\nType integer to insert on the front of the deque: ")
	input := common.GetIntInput()
	dq.InsertFront(input)
}

func insertRear(dq interfaces.IDeque) {
	fmt.Print("\nType integer to insert on at the rear of the deque: ")
	input := common.GetIntInput()
	dq.InsertRear(input)
}

func contains(dq interfaces.IDeque) {
	fmt.Print("\nType integer to look up in the deque: ")
	input := common.GetIntInput()
	cont := dq.Contains(input)
	if cont {
		fmt.Println("\nItem found.")
	} else {
		fmt.Println("\nItem not found.")
	}
}

func insertRandomDq(dq interfaces.IDeque) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("\nHow many items to add? ")
	input := common.GetIntInput()
	fmt.Println("Inserting...")
	for i := 0; i < input; i++ {
		dq.InsertFront(rand.Intn(common.MAX_RANDOM_VALUE))
	}
	fmt.Println("Done.")
}
