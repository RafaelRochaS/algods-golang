// ADTs for each Data Structure
package interfaces

// ADT for the Doubly Linked List
type IDoublyLinkedList interface {
	Clear()
	Insert(int)
	Append(int)
	Delete(int) (int, error)
	Update(int, int)
	Find(int) (interface{}, error)
	InsertArray([]int)
	Traverse()
	Length() int
}
