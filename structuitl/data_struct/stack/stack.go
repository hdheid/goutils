package stack

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/deque"
	"sync"
)

type OpFunc[T any] func(q *Stack[T])

type Stack[T any] struct {
	dq   *deque.Deque[T]
	lock synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex[T any]() OpFunc[T] {
	return func(q *Stack[T]) {
		q.lock = &sync.RWMutex{}
	}
}

func New[T any](ops ...OpFunc[T]) *Stack[T] {
	q := &Stack[T]{
		dq:   deque.New[T](),
		lock: synch.EmptyLock{},
	}

	for _, op := range ops {
		op(q)
	}

	return q
}

func (q *Stack[T]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Size()
}

func (q *Stack[T]) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Empty()
}

func (q *Stack[T]) Push(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.dq.PushFront(val)
}

func (q *Stack[T]) Pop() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.dq.PopFront()
}

func (q *Stack[T]) Top() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Back()
}

func (q *Stack[T]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.dq.Clear()
}

func (q *Stack[T]) String() string {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.String()
}
