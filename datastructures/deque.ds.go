package datastructures

import (
	"errors"
	"fmt"
)

type Deque struct {
	head   *DequeNode
	tail   *DequeNode
	length int
}

type DequeNode struct {
	previous *DequeNode
	item     int
	next     *DequeNode
}

func (dq *Deque) Clear() {
	dq.head = nil
	dq.tail = nil
	dq.length = 0
}

func (dq *Deque) InsertFront(input int) {
	if !dq.validateEmptyInsert(input) {
		newItem := DequeNode{previous: dq.head, item: input}
		dq.head.next = &newItem
		dq.head = &newItem
		dq.length++
	}
}

func (dq *Deque) InsertRear(input int) {
	if !dq.validateEmptyInsert(input) {
		newItem := DequeNode{previous: dq.tail, item: input}
		dq.tail.next = &newItem
		dq.tail = &newItem
		dq.length++
	}
}

func (dq *Deque) validateEmptyInsert(input int) bool {
	if dq.head == nil && dq.tail == nil {
		newItem := DequeNode{item: input}
		dq.head = &newItem
		dq.tail = &newItem
		return true
	}
	return false
}

func (dq Deque) validateEmptyDelete() bool {
	return dq.head == nil && dq.tail == nil
}

func (dq *Deque) DeleteFront() (int, error) {
	if dq.validateEmptyDelete() {
		return -1, errors.New("deque is empty")
	}
	item := dq.head.item
	dq.length--
	if dq.head.previous != nil {
		dq.head = dq.head.previous
		dq.head.next = nil
	} else {
		dq.head = nil
		dq.tail = nil
	}
	return item, nil
}

func (dq *Deque) DeleteRear() (int, error) {
	if dq.validateEmptyDelete() {
		return -1, errors.New("deque is empty")
	}
	item := dq.tail.item
	dq.length--
	if dq.tail.next != nil {
		dq.tail = dq.tail.next
		dq.tail.next = nil
	} else {
		dq.head = nil
		dq.tail = nil
	}
	return item, nil
}

func (dq Deque) Contains(input int) bool {
	runner := dq.head

	for runner != nil {
		if runner.item == input {
			return true
		}
		runner = runner.next
	}
	return false
}

func (dq Deque) Traverse() {
	runner := dq.head
	for runner != nil {
		fmt.Println(runner.item)
		runner = runner.next
	}
}

func (dq Deque) PeekFront() int {
	return dq.head.item
}

func (dq Deque) PeekRear() int {
	return dq.tail.item
}

func (dq Deque) Length() int {
	return dq.length
}
