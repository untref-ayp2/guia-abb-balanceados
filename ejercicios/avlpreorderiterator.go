package ejercicios

import (
	"untref-ayp2/guia-abb-balanceados/avl"

	"golang.org/x/exp/constraints"
)

type AVLPreOrderIterator[T constraints.Ordered] struct {
}

func NewAVLPreOrderIterator[T constraints.Ordered](root *avl.AVLNode[T]) *AVLPreOrderIterator[T] {
	return nil
}

func (it *AVLPreOrderIterator[T]) HasNext() bool {
	return false
}

func (it *AVLPreOrderIterator[T]) Next() (T, error) {
	var data T
	return data, nil
}
