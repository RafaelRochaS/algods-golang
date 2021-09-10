// Package datastructures implements various data structures utilized in the program
package datastructures

import (
	"errors"
	"fmt"
)

// Type definition for a default implementation of the Doubly Linked List ADT
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
	dlist.length++
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
	dlist.length++
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
func (dlist DoublyLinkedList) Find(item int) (interface{}, error) {
	if dlist.head == nil {
		return nil, errors.New("item not found in the list")
	}

	n := dlist.head

	for ; n != nil; n = n.next {
		if n.item == item {
			break
		}
	}

	if n != nil {
		return n, nil
	} else {
		return nil, errors.New("item not found in the list")
	}
}

// Inserts an entire array into a list
func (dlist *DoublyLinkedList) InsertArray(itemArr []int) {
	for _, v := range itemArr {
		dlist.Insert(v)
	}
}

// Traverses the entire linked list, printing each item
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

// Deletes a specific item from the list, if it exists
func (dlist *DoublyLinkedList) Delete(item int) (int, error) {
	if ptr, err := dlist.Find(item); err != nil {
		dlist.delete(ptr.(*DoublyLinkedListNode))
		return item, nil
	}

	return item, errors.New("item not found in the list")
}

func (dlist *DoublyLinkedList) delete(ptr *DoublyLinkedListNode) {
	if ptr.previous != nil {
		ptr.previous.next = ptr.next
	} else {
		dlist.head = ptr.next
	}

	if ptr.next != nil {
		ptr.next.previous = ptr.previous
	}
}

// Updates the value of an item within the list, if it exists
func (dlist *DoublyLinkedList) Update(old int, new int) {
	dlist.Delete(old)
	dlist.Insert(new)
}

// Type definition for the node of the doubly linked list data structure
type DoublyLinkedListNode struct {
	previous *DoublyLinkedListNode
	item     int
	next     *DoublyLinkedListNode
}
