package interfaces

// ADT for the Deque
type IDeque interface {
	InsertFront(int)
	InsertRear(int)
	DeleteFront() (int, error)
	DeleteRear() (int, error)
	Clear()
	Contains(int) bool
	Traverse()
	PeekFront() int
	PeekRear() int
	Length() int
}
