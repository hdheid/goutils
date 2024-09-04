package deque

import (
	"testing"
)

func TestDequeNew(t *testing.T) {
	q := New[int]()

	if q.String() != "[]" {
		t.Errorf("q.String() = %v", q.String())
	}
}

func TestDequePushAndPopFront(t *testing.T) {
	q := New[int]()

	cnts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, cnt := range cnts {
		q.PushFront(cnt)
	}

	if q.String() != "[10 9 8 7 6 5 4 3 2 1]" {
		t.Errorf("q.String() = %v", q.String())
	}

	for i := 0; i < 10; i++ {
		q.PopFront()
	}
	if q.String() != "[]" {
		t.Errorf("q.String() = %v", q.String())
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "deque is empty" {
				t.Errorf("q.String() = %v", r)
			}
		}
	}()
	q1 := New[int]()
	q1.PopFront()
}

func TestDequePushAndPopBack(t *testing.T) {
	q := New[int]()

	cnts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, cnt := range cnts {
		q.PushBack(cnt)
	}

	if q.String() != "[1 2 3 4 5 6 7 8 9 10]" {
		t.Errorf("q.String() = %v", q.String())
	}

	for i := 0; i < 10; i++ {
		q.PopBack()
	}
	if q.String() != "[]" {
		t.Errorf("q.String() = %v", q.String())
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "deque is empty" {
				t.Errorf("q.String() = %v", r)
			}
		}
	}()
	q1 := New[int]()
	q1.PopBack()
}

func TestDequeFront(t *testing.T) {
	q := New[int]()

	q.PushBack(1)
	if q.Front() != 1 {
		t.Errorf("q.Front() = %v,Should be 1", q.Front())
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "deque is empty" {
				t.Errorf("q.Back() = %v,Should be deque is empty", r)
			}
		}
	}()
	q1 := New[int]()
	q1.Front()
}

func TestDequeBack(t *testing.T) {
	q := New[int]()

	q.PushBack(1)
	if q.Back() != 1 {
		t.Errorf("q.Front() = %v,Should be 1", q.Front())
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "deque is empty" {
				t.Errorf("q.Back() = %v,Should be deque is empty", r)
			}
		}
	}()
	q1 := New[int]()
	q1.Back()
}

func TestDequeGetIdx(t *testing.T) {
	q := New[int]()
	for i := 0; i < 256; i++ {
		q.PushBack(i)
	}

	for i := 0; i < 256; i++ {
		if q.GetIdx(i) != i {
			t.Errorf("q.GetIdx(%v) = %v,Should be %v", i, q.GetIdx(i), i)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "out of range" {
				t.Errorf("q.Back() = %v,Should be out of range", r)
			}
		}
	}()
	q.GetIdx(512)
}

func TestDequeSet(t *testing.T) {
	q := New[int]()
	cnts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, cnt := range cnts {
		q.PushBack(cnt)
	}

	for i := 0; i < len(cnts); i++ {
		q.Set(i, cnts[i]+20)
	}

	for i := 0; i < len(cnts); i++ {
		if q.GetIdx(i) != cnts[i]+20 {
			t.Errorf("q.GetIdx(%v) = %v,Should be %v", i, q.GetIdx(i), cnts[i]+20)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "out of range" {
				t.Errorf("q.Back() = %v,Should be out of range", r)
			}
		}
	}()
	q.Set(11, 22)
}

func TestIterator(t *testing.T) {
	q := New[int]()
	cnts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, cnt := range cnts {
		q.PushBack(cnt)
	}

	it := q.Begin()
	if it.Value() != 1 {
		t.Errorf("it.Value() = %v,Should be 1", it.Value())
	}

	it = q.End()
	it.Prev()
	if it.Value() != 10 {
		t.Errorf("it.Value() = %v,Should be 10", it.Value())
	}

	it = q.First()
	if it.Value() != 1 {
		t.Errorf("it.Value() = %v,Should be 1", it.Value())
	}

	it = q.Last()
	if it.Value() != 10 {
		t.Errorf("it.Value() = %v,Should be 10", it.Value())
	}

	it = q.ItAt(0)
	if it.Value() != 1 {
		t.Errorf("it.Value() = %v,Should be 1", it.Value())
	}

	var count int
	for it = q.Begin(); it.IsValid(); it.Next() {
		if it.Value() != cnts[count] {
			t.Errorf("it.Value() = %v,Should be %d", it.Value(), cnts[count])
		}
		count++
	}

	it = q.Begin()
	if !it.Equal(q.Begin()) {
		t.Errorf("it.Equal(cloneIt) = %v,Should be true", it.Equal(q.Begin()))
	}

	cloneIt := it.Clone()
	if !cloneIt.Equal(q.Begin()) {
		t.Errorf("cloneIt = %v,Should be %v", cloneIt, q.Begin())
	}
	if cloneIt.Equal(q.Last()) {
		t.Errorf("cloneIt = %v,Should be %v", cloneIt, q.Begin())
	}

	it.SetValue(5)
	if it.Value() != 5 {
		t.Errorf("it.Value() = %v,Should be %v", it.Value(), 5)
	}
}

func TestShrinkDeque(t *testing.T) {
	q := New[int]()
	for i := 0; i < 1024; i++ {
		q.PushBack(i)
	}

	for i := 0; i < 1024; i++ {
		q.PopBack()
	}
}

func TestDequeClear(t *testing.T) {
	q := New[int]()
	for i := 0; i < 256; i++ {
		q.PushBack(i)
	}
	q.Clear()

	if q.Size() != 0 {
		t.Errorf("q.Size() = %v,Should be 0", q.Size())
	}
}

func TestRingPool(t *testing.T) {
	q := New[int]()
	for i := 0; i < 128; i++ {
		q.PushBack(i)
	}

	for i := 0; i < 5; i++ {
		q.PopBack()
	}

	for i := 0; i < 128; i++ {
		q.PushBack(i)
	}
}

// 基准测试
func BenchmarkDeque(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.PushFront(i)
	}
}
