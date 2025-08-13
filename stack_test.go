package main

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	t.Run("New SliceStack", func(t *testing.T) {
		s := NewSliceStack[int]()
		if s == nil {
			t.Error("s shouldn't be nil")
		}
		if s.Len() > 0 {
			t.Errorf("length of s should be 0, but got %v", s.Len())
		}
	})

	t.Run("New SliceStack with string type", func(t *testing.T) {
		s := NewSliceStack[string]()
		if s == nil {
			t.Error("s shouldn't be nil")
		}
		if s.Len() > 0 {
			t.Errorf("length of s should be 0, but got %v", s.Len())
		}
	})

	t.Run("Push 1 item", func(t *testing.T) {
		s := NewSliceStack[int]()
		s.Push(2)
		if s.Len() != 1 {
			t.Errorf("length of s should be 1, but got %v", s.Len())
		}
	})

	t.Run("Push 2 item", func(t *testing.T) {
		s := NewSliceStack[int]()
		s.Push(2)
		s.Push(4)
		if s.Len() != 2 {
			t.Errorf("length of s should be 1, but got %v", s.Len())
		}
	})

	t.Run("Pop item", func(t *testing.T) {

		toPushItems := []int{2, 4}

		s := NewSliceStack[int]()
		for _, item := range toPushItems {
			s.Push(item)
		}
		for i := len(toPushItems) - 1; i >= 0; i-- {
			item, ok := s.Pop()
			if !ok {
				t.Errorf("ok should be true")
			}
			want := toPushItems[i]
			if item != toPushItems[i] {
				t.Errorf("first pop item should be %v, but got %v", want, item)
			}
		}
		if s.Len() != 0 {
			t.Errorf("length of s should be 0")
		}
	})
}

func TestListStack(t *testing.T) {
	t.Run("New ListStack", func(t *testing.T) {
		s := NewListStack[int]()
		if s == nil {
			t.Error("new ListStack shouldn't be nil")
		}
		if s.Len() != 0 {
			t.Errorf("lenght of new ListStack should be 0, but got %v", s.Len())
		}
	})

	t.Run("New ListStack with string type", func(t *testing.T) {
		s := NewListStack[string]()
		if s == nil {
			t.Error("s shouldn't be nil")
		}
		if s.Len() > 0 {
			t.Errorf("length of s should be 0, but got %v", s.Len())
		}
	})

	t.Run("Push 1 item", func(t *testing.T) {
		s := NewListStack[int]()
		s.Push(2)
		if s.Len() != 1 {
			t.Errorf("length of s should be 1, but got %v", s.Len())
		}
	})

	t.Run("Push 2 item", func(t *testing.T) {
		s := NewListStack[int]()
		s.Push(2)
		s.Push(4)
		if s.Len() != 2 {
			t.Errorf("length of s should be 1, but got %v", s.Len())
		}
	})

	t.Run("Pop item", func(t *testing.T) {

		toPushItems := []int{2, 4}

		s := NewListStack[int]()
		for _, item := range toPushItems {
			s.Push(item)
		}
		for i := len(toPushItems) - 1; i >= 0; i-- {
			item, ok := s.Pop()
			if !ok {
				t.Errorf("ok should be true")
			}
			want := toPushItems[i]
			if item != toPushItems[i] {
				t.Errorf("first pop item should be %v, but got %v", want, item)
			}
		}
		if s.Len() != 0 {
			t.Errorf("length of s should be 0")
		}
	})
}
