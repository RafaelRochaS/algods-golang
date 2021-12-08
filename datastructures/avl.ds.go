package datastructures

import "fmt"

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	key    int
	left   *AVLNode
	right  *AVLNode
	parent *AVLNode
	height int
}

func (tree AVLTree) Search(key int) (bool, *AVLNode) {
	if tree.root == nil {
		return false, nil
	}

	return tree.search(key, tree.root)
}

func (tree AVLTree) search(key int, node *AVLNode) (bool, *AVLNode) {

	if node == nil {
		return false, nil
	}

	if node.key < key {
		return tree.search(key, node.left)
	} else if node.key > key {
		return tree.search(key, node.right)
	} else {
		return true, node
	}
}

func (tree AVLTree) InOrderTraversal(node *AVLNode) {
	if node != nil {
		tree.InOrderTraversal(node.left)
		fmt.Println(node.key)
		tree.InOrderTraversal(node.right)
	}
}

func (tree AVLTree) PreOrderTraversal(node *AVLNode) {
	if node != nil {
		fmt.Println(node.key)
		tree.PreOrderTraversal(node.left)
		tree.PreOrderTraversal(node.right)
	}
}

func (tree AVLTree) PostOrderTraversal(node *AVLNode) {
	if node != nil {
		tree.PreOrderTraversal(node.left)
		tree.PreOrderTraversal(node.right)
		fmt.Println(node.key)
	}
}

func (tree AVLTree) Minimum(node *AVLNode) *AVLNode {
	for node.left != nil {
		node = node.left
	}

	return node
}

func (tree AVLTree) Maximum(node *AVLNode) *AVLNode {
	for node.right != nil {
		node = node.right
	}

	return node
}

func (tree AVLTree) Successor(node *AVLNode) *AVLNode {
	if node.right != nil {
		return tree.Minimum(node.right)
	}

	successor := node.parent
	for successor != nil && node == successor.right {
		node = successor
		successor = successor.parent
	}

	return successor
}

func (tree AVLTree) Predecessor(node *AVLNode) *AVLNode {
	if node.left != nil {
		return tree.Maximum(node.left)
	}

	predecessor := node.left
	for predecessor != nil && node == predecessor.parent {
		node = predecessor
		predecessor = predecessor.left
	}

	return predecessor
}

func (tree *AVLTree) Insert(key int) {

	if tree.root == nil {
		tree.root = &AVLNode{
			key:    key,
			left:   nil,
			right:  nil,
			parent: nil,
			height: 0,
		}
	} else {
		tree.insert(tree.root, key)
	}
}

func (tree AVLTree) insert(node *AVLNode, key int) {

	if key >= node.key {
		if node.right == nil {
			node.height++
			node.right = &AVLNode{
				key:    key,
				left:   nil,
				right:  nil,
				parent: node,
				height: 0,
			}
		} else {
			tree.insert(node.right, key)
		}
	} else {
		if node.left == nil {
			node.height++
			node.left = &AVLNode{
				key:    key,
				left:   nil,
				right:  nil,
				parent: node,
				height: 0,
			}
		} else {
			tree.insert(node.left, key)
		}
	}
}
