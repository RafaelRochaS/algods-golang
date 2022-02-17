package datastructures

import "errors"

type LinkedList struct {
	root *LinkedListNode
}

type LinkedListNode struct {
	prev  *LinkedListNode
	value int
	key   string
	next  *LinkedListNode
}

type HashTable struct {
	Table []LinkedList
	Size  int
	Items int
}

const INITIAL_SIZE = 8

func NewHashTable() HashTable {
	return HashTable{
		Size:  INITIAL_SIZE,
		Items: 0,
		Table: make([]LinkedList, INITIAL_SIZE),
	}
}

func (ht HashTable) Insert(key string, value int) {
	ht.Items++
	ht.evaluateSize()

	hashValue := ht.hash(key)
	if ht.Table[hashValue].root == nil {
		ht.Table[hashValue].root = &LinkedListNode{
			prev:  nil,
			value: value,
			key:   key,
			next:  nil,
		}
	} else {
		ht.Table[hashValue].root.findInsertPosition(key, value)
	}
}

func (node *LinkedListNode) findInsertPosition(key string, value int) {
	if node.next == nil {
		node.next = &LinkedListNode{
			prev:  node,
			value: value,
			key:   key,
			next:  nil,
		}
	} else {
		node.next.findInsertPosition(key, value)
	}
}

func (ht HashTable) Search(key string) (int, error) {
	hashValue := ht.hash(key)
	return ht.Table[hashValue].root.searchLinkedList(key)
}

func (node *LinkedListNode) searchLinkedList(key string) (int, error) {
	if node == nil {
		return -1, errors.New("item not found on dictionary")
	}

	if node.key == key {
		return node.value, nil
	}

	return node.next.searchLinkedList(key)
}

func (ht HashTable) Delete(key string) error {
	ht.Items--
	ht.evaluateSize()
	hashValue := ht.hash(key)
	return ht.Table[hashValue].root.deleteLinkedList(key)
}

func (node *LinkedListNode) deleteLinkedList(key string) error {
	if node == nil {
		return errors.New("key not found")
	}

	if node.key == key {
		if node.prev != nil {
			node.prev.next = nil
		}
		node = nil
		return nil
	}

	return node.next.deleteLinkedList(key)
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
// taken from: https://flaviocopes.com/golang-data-structure-hashtable/
// Added the mod size, so that the max size of the array (and max keys) is size
func (ht HashTable) hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h % ht.Size
}

func (ht HashTable) loadFactor() float32 {
	return float32(ht.Size) / float32(ht.Items)
}

func (ht HashTable) evaluateSize() {
	lf := ht.loadFactor()
	if lf >= 1 {
		ht.resize(true)
	} else if lf <= 0.25 {
		ht.resize(false)
	}
}

func (ht *HashTable) resize(up bool) {
	var newTable []LinkedList
	//oldTable := ht.Table
	if up {
		newSize := ht.Size * 2
		newTable = make([]LinkedList, newSize)
		ht.Size = newSize
	} else {
		newSize := ht.Size / 2
		newTable = make([]LinkedList, newSize)
		ht.Size = newSize
	}

	ht.Table = newTable

}
