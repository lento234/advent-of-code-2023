package utils

import "fmt"

// LILO and FIFO queue
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

func (q *Queue[T]) PopBack() (T, error) {
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

func (q *Queue[T]) PopFront() (T, error) {
	if q.Empty() {
		return *new(T), fmt.Errorf("queue is empty")
	}
	// duplicate
	h := *q
	l := len(h)

	// Pop front
	var p T
	p, *q = h[0], h[1:l]

	return p, nil
}
