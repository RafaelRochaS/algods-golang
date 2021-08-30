package interfaces

type IDoublyLinkedList interface {
	Clear()
	Insert(int)
	Append(int)
	Delete(int) int
	Update(int, int)
	Contains(int, bool) bool
	InsertArray([]int)
}
