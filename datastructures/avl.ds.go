package datastructures

import (
	"errors"
	"fmt"
)

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
			height: 1,
		}
	} else {
		tree.insert(tree.root, key)
	}
}

func (tree *AVLTree) insert(node *AVLNode, key int) {
	if key >= node.key {
		if node.right == nil {
			node.height++
			node.right = &AVLNode{
				key:    key,
				left:   nil,
				right:  nil,
				parent: node,
				height: 1,
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
				height: 1,
			}
		} else {
			tree.insert(node.left, key)
		}
	}

	tree.rebalance(node)

}

func (tree *AVLTree) rebalance(node *AVLNode) {
	var rotationResult *AVLNode
	rotated := false

	bf := balanceFactor(*node)
	parent := node.parent

	if bf >= 2 {
		if bfRightChild := balanceFactor(*node.right); bfRightChild >= 1 {
			rotationResult = leftRotate(node)
		} else {
			rotationResult = rightLeftRotate(node)
		}
		rotated = true
	} else if bf <= -2 {
		if bfLeftChild := balanceFactor(*node.left); bfLeftChild >= 1 {
			rotationResult = leftRightRotate(node)
		} else {
			rotationResult = rightRotate(node)
		}
		rotated = true
	}
	updateHeight(node)

	if rotated {
		if parent != nil {
			if parent.left == node {
				parent.left = rotationResult
			} else {
				parent.right = rotationResult
			}
			updateHeight(parent)
		} else {
			tree.root = rotationResult
		}
	}
}

func balanceFactor(node AVLNode) int {
	return height(node.right) - height(node.left)
}

func updateHeight(node *AVLNode) {
	heightLeft := height(node.left)
	heightRight := height(node.right)

	if heightLeft >= heightRight {
		node.height = heightLeft + 1
	} else {
		node.height = heightRight + 1
	}
}

func height(node *AVLNode) int {
	if node == nil {
		return 0
	}

	return node.height
}

func leftRotate(node *AVLNode) *AVLNode {
	rightChild := node.right
	tmp := rightChild.left
	rightChild.left = node
	rightChild.parent = node.parent
	node.right = tmp
	if node.right != nil {
		node.right.parent = node
	}
	node.parent = rightChild
	updateHeight(node)
	updateHeight(rightChild)
	return rightChild
}

func rightRotate(node *AVLNode) *AVLNode {
	leftChild := node.left
	tmp := leftChild.right
	leftChild.right = node
	leftChild.parent = node.parent
	node.left = tmp
	if node.left != nil {
		node.left.parent = node
	}
	node.parent = leftChild
	updateHeight(node)
	updateHeight(leftChild)
	return leftChild
}

func leftRightRotate(node *AVLNode) *AVLNode {
	// Left Rotate
	a := node.left
	b := a.right
	node.left = b
	a.right = b.left
	if b.left != nil {
		b.left.parent = a
	}
	a.parent = b
	b.left = a
	b.parent = node

	// Right Rotate
	tmpbRight := b.right
	b.right = node
	b.parent = node.parent
	node.left = tmpbRight
	if tmpbRight != nil {
		tmpbRight.parent = node
	}
	node.parent = b

	updateHeight(node)
	updateHeight(a)
	updateHeight(b)

	return b
}

func rightLeftRotate(node *AVLNode) *AVLNode {
	// Right Rotate
	c := node.right
	b := c.left
	node.right = b
	tmpbRight := b.right
	b.right = c
	b.parent = node
	c.left = tmpbRight
	if tmpbRight != nil {
		tmpbRight.parent = c
	}
	c.parent = b

	// Left Rotate
	tmpBLeft := b.left
	b.left = node
	b.parent = node.parent
	node.right = tmpBLeft
	if tmpBLeft != nil {
		tmpBLeft.parent = node
	}
	node.parent = b

	updateHeight(node)
	updateHeight(c)
	updateHeight(b)

	return b
}

func (tree *AVLTree) Delete(key int) error {
	return tree.delete(tree.root, key)
}

func (tree *AVLTree) delete(node *AVLNode, key int) error {
	if node == nil {
		return errors.New("key not found in tree")
	} else if node.key > key {
		tree.delete(node.left, key)
	} else if node.key < key {
		tree.delete(node.right, key)
	} else {
		if node.left == nil {
			tree.subtreeShift(node, node.right)
		} else if node.right == nil {
			tree.subtreeShift(node, node.left)
		} else {
			successor := tree.Successor(node)
			if successor.parent != node {
				tree.subtreeShift(successor, successor.right)
				successor.right = node.right
				successor.right.parent = successor
			}
			tree.subtreeShift(node, successor)
			successor.left = node.left
			successor.left.parent = successor
			updateHeight(successor)
		}
	}

	tree.rebalance(node)

	return nil
}

func (tree *AVLTree) subtreeShift(nodeDeleted *AVLNode, nextInLine *AVLNode) {
	if nodeDeleted.parent == nil {
		tree.root = nextInLine
	} else if nodeDeleted == nodeDeleted.parent.left {
		nodeDeleted.parent.left = nextInLine
		updateHeight(nodeDeleted.parent)
	} else {
		nodeDeleted.parent.right = nextInLine
		updateHeight(nodeDeleted.parent)
	}

	if nextInLine != nil {
		nextInLine.parent = nodeDeleted.parent
	}
}
