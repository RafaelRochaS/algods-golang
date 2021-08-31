// Package datastructures implements various data structures utilized in the program
package datastructures

import "fmt"

// Type definition for the Doubly Linked List data structure
type DoublyLinkedList struct {
	head   *DoublyLinkedListNode
	tail   *DoublyLinkedListNode
	length int
}

// Returns the current size of the list
func (dlist *DoublyLinkedList) Length() int {
	return dlist.length
}

// Empties the current list
func (dlist *DoublyLinkedList) Clear() {
	dlist.head = nil
	dlist.tail = nil
	dlist.length = 0
}

// Inserts an item at the beginning of the list
func (dlist *DoublyLinkedList) Insert(item int) {
	if dlist.head == nil && dlist.tail == nil {
		dlist.insertEmptyList(item)
		return
	}

	newItem := DoublyLinkedListNode{nil, item, dlist.head}
	if dlist.head != nil {
		dlist.head.previous = &newItem
		newItem.next = dlist.head
		dlist.head = &newItem
	} else {
		dlist.head = &newItem
	}
}

func (dlist *DoublyLinkedList) insertEmptyList(item int) {
	newItem := DoublyLinkedListNode{nil, item, nil}
	dlist.head = &newItem
	dlist.tail = &newItem
}

// Inserts an item at the end of the list
func (dlist *DoublyLinkedList) Append(item int) {
	if dlist.head == nil && dlist.tail == nil {
		dlist.insertEmptyList(item)
		return
	}

	newItem := DoublyLinkedListNode{dlist.tail, item, nil}
	if dlist.tail != nil {
		dlist.tail.next = &newItem
		dlist.tail = &newItem
	} else {
		dlist.tail = &newItem
	}
}

// Checks whether a specific item is present in the list
func (dlist DoublyLinkedList) Contains(item int) bool {
	if dlist.head == nil {
		return false
	}

	for n := dlist.head; n != dlist.tail; n = n.next {
		if n.item == item {
			return true
		}
	}
	return false
}

// Inserts an entire array into a list
func (dlist *DoublyLinkedList) InsertArray(itemArr []int) {
	for _, v := range itemArr {
		dlist.Insert(v)
	}
}

func (dlist DoublyLinkedList) Traverse() {
	runner := dlist.head
	if runner == nil {
		return
	}

	for runner != nil {
		fmt.Println(runner.item)
		runner = runner.next
	}
}

// TODO: Implement correctly
// func (dlist *DoublyLinkedList) Delete(item int) (int, error) {
// 	if !dlist.Contains(item) {
// 		return 0, errors.New("item not found in the list")
// 	}

// 	return item, nil
// }

// func (dlist *DoublyLinkedList) Update(old int, new int) {

// }

// Type definition for the node of the doubly linked list data structure
type DoublyLinkedListNode struct {
	previous *DoublyLinkedListNode
	item     int
	next     *DoublyLinkedListNode
}
