package datastructures

type LinkedList struct {
	root *LinkedListNode
}

type LinkedListNode struct {
	prev  *LinkedListNode
	value int
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
			next:  nil,
		}
	} else {
		ht[hashValue].root.findInsertPosition(value)
	}
}

func (node *LinkedListNode) findInsertPosition(value int) {
	if node.next == nil {
		node.next = &LinkedListNode{
			prev:  node,
			value: value,
			next:  nil,
		}
	} else {
		node.next.findInsertPosition(value)
	}
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
