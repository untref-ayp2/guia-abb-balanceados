package stack

import "errors"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(x T) {
	s.elements = append(s.elements, x)
}

func (s *Stack[T]) Pop() (T, error) {
	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return x, nil
}

func (s *Stack[T]) Top() (T, error) {
	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.elements[len(s.elements)-1]
	return x, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
