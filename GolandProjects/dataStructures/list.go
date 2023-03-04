package collections

import (
	"errors"
	"fmt"
)

type _type interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string | bool
}

type LinkedList[T _type] struct {
	root *lNode[T]
	size int
}

type lNode[T _type] struct {
	val  T
	next *lNode[T]
	prev *lNode[T]
}

func NewList[T _type]() *LinkedList[T] {
	var l LinkedList[T]
	l.root = &lNode[T]{}
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
	return &l
}

func (l *LinkedList[T]) Add(elem T) {
	var node lNode[T]
	var currNode = l.root

	for currNode.next != l.root {
		currNode = currNode.next
	}
	currNode.next = &node
	node.prev = currNode
	node.next = l.root
	node.val = elem
	l.size++
}

func (l *LinkedList[T]) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func (l *LinkedList[T]) IndexOf(element T) (int, error) {
	var el = l.root.next
	var index = 0
	for el != l.root {
		if el.val == element {
			return index, nil
		}
		index++
		el = el.next
	}
	return -1, errors.New(fmt.Sprintf("No index for element %v found", element))
}

func (l *LinkedList[T]) At(pos int) (T, error) {
	var el = l.root.next
	index := 0
	for index < pos {
		if el == l.root {
			return el.val, errors.New("Index out of range")
		}
		el = el.next
		index++
	}
	return el.val, nil
}

func (l *LinkedList[T]) DeleteValue(value T) error {
	var el = l.root.next

	for el != l.root {
		if el.val == value {
			el.prev.next = el.next
			el.next.prev = el.prev
			l.size--
			return nil
		}
		el = el.next
	}

	return errors.New("No element found")
}

func (l *LinkedList[T]) DeleteAt(pos int) error {
	var el = l.root.next
	var index = 0

	for index < pos {
		if el == l.root {
			return errors.New("Index out of range")
		}
		el = el.next
		index++
	}

	el.prev.next = el.next
	el.next.prev = el.prev
	l.size--
	return nil
}

func (l *LinkedList[T]) Filter(condition func(a T) bool) *LinkedList[T] {
	output := NewList[T]()
	var el = l.root.next
	for el != l.root {
		if condition(el.val) {
			output.Add(el.val)
			output.size++
		}
		el = el.next
	}
	return output
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
