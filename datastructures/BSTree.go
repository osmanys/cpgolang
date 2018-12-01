package datastructures

import (
	"sync"
)

// BSTree struct
type BSTree struct {
	root *bstNode
	size int
	lock sync.RWMutex
}

type bstNode struct {
	value       int
	left, right *bstNode
}

// Insert function
func (bst *BSTree) Insert(value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	if bst.root == nil {
		bst.root = &bstNode{value: value}
	} else {
		insertNode(bst.root, value)
	}
	bst.size++
}

func insertNode(root *bstNode, value int) {
	if value < root.value {
		if root.left == nil {
			root.left = &bstNode{value: value}
		} else {
			insertNode(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &bstNode{value: value}
		} else {
			insertNode(root.right, value)
		}
	}
}

// InOrderTraverse function
func (bst *BSTree) InOrderTraverse(f func(int)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(root *bstNode, f func(int)) {
	if root != nil {
		inOrderTraverse(root.left, f)
		f(root.value)
		inOrderTraverse(root.right, f)
	}
}
