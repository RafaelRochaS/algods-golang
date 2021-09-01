package interfaces

import "algods/datastructures"

type IDoublyLinkedList interface {
	Clear()
	Insert(int)
	Append(int)
	Delete(int) (int, error)
	delete(*datastructures.DoublyLinkedListNode)
	Update(int, int)
	Find(int) (*datastructures.DoublyLinkedListNode, error)
	InsertArray([]int)
}
