package interfaces

type IPriorityQueue interface {
	Insert(int)
	Max() int
	ExtractMax() int
	IncreaseKey(int, int)
}
