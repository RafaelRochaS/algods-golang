package interfaces

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
}
