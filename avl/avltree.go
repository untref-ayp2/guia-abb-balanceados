package avl

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type AVLTree[T constraints.Ordered] struct {
	root *AVLNode[T]
}

func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

func (avl *AVLTree[T]) String() string {
	return avl.root.string()
}

func (avl *AVLTree[T]) GetRoot() *AVLNode[T] {
	return avl.root
}

func (avl *AVLTree[T]) Insert(k T) {
	avl.root = avl.root.insert(k)
}

func (avl *AVLTree[T]) Remove(k T) {
	avl.root = avl.root.remove(k)
}

func (avl *AVLTree[T]) Search(k T) bool {
	return avl.root.search(k)
}

func (avl *AVLTree[T]) FindMin() (T, error) {
	if avl.root == nil {
		var zero T
		return zero, errors.New("árbol vacío")
	}
	minNode := avl.root.findMin()
	return minNode.GetData(), nil
}

func (avl *AVLTree[T]) FindMax() (T, error) {
	if avl.root == nil {
		var zero T
		return zero, errors.New("árbol vacío")
	}
	maxNode := avl.root.findMax()
	return maxNode.GetData(), nil
}

func (avl *AVLTree[T]) GetHeight() int {
	return avl.root.GetHeight()
}

func (avl *AVLTree[T]) GetBalance() int {
	return avl.root.GetBalance()
}

func (avl *AVLTree[T]) IsEmpty() bool {
	return avl.root == nil
}

func (avl *AVLTree[T]) Clear() {
	avl.root = nil
}

func (avl *AVLTree[T]) Size() int {
	return avl.root.Size()
}
