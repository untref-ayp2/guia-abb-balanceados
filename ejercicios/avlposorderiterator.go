package ejercicios

import (
	"untref-ayp2/guia-abb-balanceados/avl"

	"golang.org/x/exp/constraints"
)

type AVLPosOrderIterator[T constraints.Ordered] struct {
}

func NewAVLPosOrderIterator[T constraints.Ordered](root *avl.AVLNode[T]) *AVLPosOrderIterator[T] {
	return nil
}

func (it *AVLPosOrderIterator[T]) HasNext() bool {
	return false
}

func (it *AVLPosOrderIterator[T]) Next() (T, error) {
	var data T
	return data, nil
}
