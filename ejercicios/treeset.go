package ejercicios

import "golang.org/x/exp/constraints"

type TreeSet[T constraints.Ordered] struct {
	// Implementar
}

func NewTreeSet[T constraints.Ordered](elements ...T) *TreeSet[T] {
	// Implementar
	return nil
}

func (s *TreeSet[T]) Contains(element T) bool {
	// Implementar
	return false
}

func (s *TreeSet[T]) Add(elements ...T) {
	// Implementar
}

func (s *TreeSet[T]) Remove(element T) {
	// Implementar
}

func (s *TreeSet[T]) Size() int {
	// Implementar
	return 0
}

func (s *TreeSet[T]) Values() []T {
	// Implementar
	return nil
}

func (s *TreeSet[T]) String() string {
	// Implementar
	return ""
}
