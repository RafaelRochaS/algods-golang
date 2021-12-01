package datastructures

import "algods/common"

type Heap struct {
	Keys     []int
	HeapSize int
}

func Parent(i int) int {
	return i >> 1 // floor i / 2
}

func Left(i int) int {
	return (i << 1) + 1 // adapted for zero-based arrays. 2i + 1
}

func Right(i int) int {
	return (i << 1) + 2 // adapted for zero-based arrays. 2i + 2
}

func (h *Heap) MaxHeapify(i int) {
	leftTreeRoot := Left(i)
	rightTreeRoot := Right(i)
	var largest int

	if leftTreeRoot <= h.HeapSize-1 && h.Keys[leftTreeRoot] > h.Keys[i] {
		largest = leftTreeRoot
	} else {
		largest = i
	}
	if rightTreeRoot <= h.HeapSize-1 && h.Keys[rightTreeRoot] > h.Keys[largest] {
		largest = rightTreeRoot
	}

	if largest != i {
		common.Swap(h.Keys, i, largest)
		h.MaxHeapify(largest)
	}
}

func BuildMaxHeap(arr []int) Heap {
	h := Heap{Keys: arr, HeapSize: len(arr)}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}

	return h
}

func Heapsort(arr []int) {
	h := BuildMaxHeap(arr)
	for i := len(arr) - 1; i >= 1; i-- {
		common.Swap(h.Keys, 0, i)
		h.HeapSize--
		h.MaxHeapify(0)
	}
}

func (h *Heap) Insert(key int) {
	h.HeapSize++
	h.Keys = append(h.Keys, key)
	h.MaxHeapify(h.HeapSize - 1)
}

func (h *Heap) Max() int {
	return h.Keys[0]
}

func (h *Heap) ExtractMax() int {
	if h.HeapSize == 0 {
		return -1
	}
	max := h.Keys[0]
	h.Keys[0] = h.Keys[h.HeapSize-1]
	h.HeapSize--
	h.MaxHeapify(0)
	return max
}

func (h *Heap) IncreaseKey(i int, key int) {
	if key < h.Keys[i] {
		return
	}
	h.Keys[i] = key
	for i > 0 && h.Keys[Parent(i)] < h.Keys[i] {
		common.Swap(h.Keys, i, Parent(i))
		i = Parent(i)
	}
}
