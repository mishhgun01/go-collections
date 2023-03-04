package collections

import (
	"errors"
	"fmt"
)

type Set[T _type] struct {
	root *sNode[T]
	size int
}

type sNode[T _type] struct {
	val        T
	next, prev *sNode[T]
}

func NewSet[T _type]() *Set[T] {
	var s Set[T]
	s.root = &sNode[T]{}
	s.root.next = s.root
	s.root.prev = s.root
	s.size = 0
	return &s
}

func (s *Set[T]) Add(value T) error {
	var el = s.root.next

	for el.next != s.root {
		if el.val == value {
			return errors.New("Element already exists")
		}
		el = el.next
	}
	var newEl = &sNode[T]{}
	newEl.val = value
	newEl.prev = el
	el.next = newEl
	newEl.next = s.root
	s.size++
	return nil
}

func (s *Set[T]) String() string {
	el := s.root.next
	var str string
	for el != s.root {
		str += fmt.Sprintf("%v ", el.val)
		el = el.next
	}
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
func (s *Set[T]) At(pos int) (T, error) {
	var el = s.root.next
	index := 0
	for index < pos {
		if el == s.root {
			return el.val, errors.New("Index out of range")
		}
		el = el.next
		index++
	}
	return el.val, nil
}

func (s *Set[T]) DeleteValue(value T) error {
	var el = s.root.next

	for el != s.root {
		if el.val == value {
			el.prev.next = el.next
			el.next.prev = el.prev
			s.size--
			return nil
		}
		el = el.next
	}

	return errors.New("No element found")
}

func (s *Set[T]) DeleteAt(pos int) error {
	var el = s.root.next
	var index = 0

	for index < pos {
		if el == s.root {
			return errors.New("Index out of range")
		}
		el = el.next
		index++
	}

	el.prev.next = el.next
	el.next.prev = el.prev
	s.size--
	return nil
}

func (s *Set[T]) Filter(condition func(a T) bool) *Set[T] {
	output := NewSet[T]()
	var el = s.root.next
	for el != s.root {
		if condition(el.val) {
			_ = output.Add(el.val)
			output.size++
		}
		el = el.next
	}
	return output
}

func (s *Set[T]) Size() int {
	return s.size
}
