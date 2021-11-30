package interfaces

import "algods/datastructures"

type IHeap interface {
	Insert(*datastructures.Heap, int)
	Max(datastructures.Heap) int
	ExtractMax(*datastructures.Heap) int
	IncreaseKey(*datastructures.Heap, int, int)
}
