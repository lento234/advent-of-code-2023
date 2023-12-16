package utils

import "fmt"

// FILO queue
type Queue[T any] []T

func (q *Queue[T]) Push(p T) {
	*q = append(*q, p)
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Pop() (T, error) {
	if q.Empty() {
		return *new(T), fmt.Errorf("queue is empty")
	}
	// duplicate
	h := *q
	l := len(h)

	// Pop back
	var p T
	p, *q = h[l-1], h[0:l-1]

	return p, nil
}
