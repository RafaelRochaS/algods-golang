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

type HashTable []LinkedList

const MAX_SIZE = 50

func NewHashTable() HashTable {
	return make(HashTable, MAX_SIZE)
}

func (ht HashTable) Insert(key string, value int) {
	hashValue := hash(key)
	if ht[hashValue].root == nil {
		ht[hashValue].root = &LinkedListNode{
			prev:  nil,
			value: value,
			key:   key,
			next:  nil,
		}
	} else {
		ht[hashValue].root.findInsertPosition(key, value)
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
	hashValue := hash(key)
	return ht[hashValue].root.searchLinkedList(key)
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

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
// taken from: https://flaviocopes.com/golang-data-structure-hashtable/
// Added the mod MAX_SIZE, so that the max size of the array (and max keys) is MAX_SIZE
func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h % MAX_SIZE
}
