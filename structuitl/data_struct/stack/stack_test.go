package stack

import (
	"testing"
)

func TestStackNew(t *testing.T) {
	s1 := New[int]()
	s2 := New[int](WithRWMutex[int]())

	if s1 == nil || s2 == nil {
		t.Errorf("s1==nil || s2 ==nil")
	}
}

func TestStack(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	if s.String() != "[2 1]" {
		t.Errorf("s.String() = %v, should be [2 1]", s.String())
	}

	s.Pop()
	if s.String() != "[1]" {
		t.Errorf("s.String() = %v, should be [1]", s.String())
	}

	if s.Top() != 1 {
		t.Errorf("s.Top() = %v, should be 1", s.Top())
	}

	if s.Size() != 1 {
		t.Errorf("s.Size() = %v, should be 1", s.Size())
	}

	s.Clear()
	if s.Size() != 0 {
		t.Errorf("s.Clear() error")
	}

	if !s.Empty() {
		t.Errorf("s.Empty() should be true")
	}
}
