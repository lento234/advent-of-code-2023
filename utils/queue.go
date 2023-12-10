package utils

import "fmt"

type Queue []interface{}

func (q *Queue) Push(p interface{}) {
	*q = append(*q, p)
}
func (q *Queue) Pop() (interface{}, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("queue empty")
	}

	// duplicate
	h := *q
	l := len(h)

	// Pop back
	var p interface{}
	p, *q = h[l-1], h[0:l-1]

	return p, nil
}
