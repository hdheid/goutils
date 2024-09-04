package queue

import (
	"testing"
)

func TestQueueNew(t *testing.T) {
	q1 := New[int]()
	q2 := New[int](WithRWMutex[int]())

	if q1 == nil || q2 == nil {
		t.Errorf("s1==nil || s2 ==nil")
	}
}

func TestQueue(t *testing.T) {
	q := New[int]()

	q.Push(1)
	q.Push(2)
	if q.String() != "[1 2]" {
		t.Errorf("s.String() = %v, should be [1 2]", q.String())
	}

	if q.Front() != 1 {
		t.Errorf("s.Front() = %v, should be 1", q.Front())
	}

	if q.Back() != 2 {
		t.Errorf("s.Back() = %v, should be false", q.Back())
	}

	q.Pop()
	if q.String() != "[2]" {
		t.Errorf("s.String() = %v, should be [2]", q.String())
	}

	if q.Size() != 1 {
		t.Errorf("s.Size() = %v, should be 1", q.Size())
	}

	q.Clear()
	if q.Size() != 0 {
		t.Errorf("s.Clear() error")
	}

	if !q.Empty() {
		t.Errorf("s.Empty() should be true")
	}
}
