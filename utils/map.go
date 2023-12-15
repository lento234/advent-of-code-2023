package utils

import "slices"

type OrderedMap struct {
	Items       map[string]int
	OrderedKeys []string
}

func (l *OrderedMap) Add(key string, value int) {
	if len(l.Items) == 0 {
		l.Items = make(map[string]int, 0)
	}

	l.Items[key] = value
	if !slices.Contains(l.OrderedKeys, key) {
		l.OrderedKeys = append(l.OrderedKeys, key)
	}
}

func (l *OrderedMap) Remove(key string) {
	if _, ok := l.Items[key]; ok {
		delete(l.Items, key)
		idx := slices.Index(l.OrderedKeys, key)
		l.OrderedKeys = append(l.OrderedKeys[:idx], l.OrderedKeys[idx+1:]...)
	}
}
