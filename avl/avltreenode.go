package avl

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type AVLNode[T constraints.Ordered] struct {
	data        T
	height      int
	left, right *AVLNode[T]
}

func NewAVLNode[T constraints.Ordered](data T, left, right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 0}
}

func (n *AVLNode[T]) GetData() T {
	return n.data
}

func (n *AVLNode[T]) GetLeft() *AVLNode[T] {
	return n.left
}

func (n *AVLNode[T]) GetRight() *AVLNode[T] {
	return n.right
}

func (n *AVLNode[T]) GetHeight() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *AVLNode[T]) GetBalance() int {
	if n == nil {
		return 0
	}
	return n.left.GetHeight() - n.right.GetHeight()
}

func (n *AVLNode[T]) updateHeight() {
	n.height = 1 + max(n.left.GetHeight(), n.right.GetHeight())
}

func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	// Implementar
	return nil
}

func (n *AVLNode[T]) remove(value T) *AVLNode[T] {
	// Implementar
	return nil
}

func (n *AVLNode[T]) search(k T) bool {
	// Implementar
	return false
}

func (n *AVLNode[T]) findMin() *AVLNode[T] {
	if n.left == nil {
		return n
	}
	return n.left.findMin()
}

func (n *AVLNode[T]) findMax() *AVLNode[T] {
	if n.right == nil {
		return n
	}
	return n.right.findMax()
}

func (n *AVLNode[T]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}
