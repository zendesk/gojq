package gojq

import (
	"unsafe"
)

type stack struct {
	data  []block
	index int
	limit int
}

type block struct {
	value interface{}
	next  int
}

func newStack() *stack {
	return &stack{index: -1, limit: -1}
}

func (s *stack) push(v interface{}) {
	b := block{v, s.index}
	i := s.index + 1
	if i <= s.limit {
		i = s.limit + 1
	}
	s.index = i
	if i < len(s.data) {
		s.data[i] = b
	} else {
		s.data = append(s.data, b)
	}
}

func (s *stack) pop() interface{} {
	b := s.data[s.index]
	s.index = b.next
	return b.value
}

func (s *stack) top() interface{} {
	return s.data[s.index].value
}

func (s *stack) empty() bool {
	return s.index < 0
}

func (s *stack) save() (index, limit int) {
	index, limit = s.index, s.limit
	if s.index > s.limit {
		s.limit = s.index
	}
	return
}

func (s *stack) restore(index, limit int) {
	s.index, s.limit = index, limit
}

func (s *stack) memSize() uintptr {
	size := unsafe.Sizeof(*s)
	for _, block := range s.data {
		size = size + block.memSize()
	}
	return size
}

func (b block) memSize() uintptr {
	return unsafe.Sizeof(b) + sizeofValue(b.value)
}
