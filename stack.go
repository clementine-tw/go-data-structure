package main

type SliceStack[T any] struct {
	items []T
}

func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		items: make([]T, 0),
	}
}

func (s *SliceStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *SliceStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		return *new(T), false
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item, true
}

func (s *SliceStack[T]) Len() int {
	return len(s.items)
}

type ListNode[T any] struct {
	Next *ListNode[T]
	Prev *ListNode[T]
	Val  T
}

func NewListNode[T any](v T, prev *ListNode[T]) *ListNode[T] {
	return &ListNode[T]{
		Next: nil,
		Prev: prev,
		Val:  v,
	}
}

type ListStack[T any] struct {
	root   *ListNode[T]
	last   *ListNode[T]
	length int
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{root: nil, last: nil, length: 0}
}

func (s *ListStack[T]) Len() int {
	return s.length
}

func (s *ListStack[T]) Push(item T) {
	if s.root == nil {
		s.root = NewListNode(item, nil)
		s.last = s.root
		s.length = 1
		return
	}

	s.last.Next = NewListNode(item, s.last)
	s.last = s.last.Next
	s.length++
}

func (s *ListStack[T]) Pop() (T, bool) {
	if s.root == nil {
		return *new(T), false
	}
	item := s.last.Val

	if s.root == s.last {
		s.root = nil
		s.last = nil
		s.length = 0
		return item, true
	}

	prev := s.last.Prev
	prev.Next = nil
	s.last.Prev = nil
	s.last = prev

	s.length--

	return item, true
}
