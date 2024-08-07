package queue

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/deque"
	"sync"
)

type OpFunc[T any] func(q *Queue[T])

type Queue[T any] struct {
	dq   *deque.Deque[T]
	lock synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex[T any]() OpFunc[T] {
	return func(q *Queue[T]) {
		q.lock = &sync.RWMutex{}
	}
}

func New[T any](ops ...OpFunc[T]) *Queue[T] {
	q := &Queue[T]{
		dq:   deque.New[T](),
		lock: synch.EmptyLock{},
	}

	for _, op := range ops {
		op(q)
	}

	return q
}

func (q *Queue[T]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Size()
}

func (q *Queue[T]) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Empty()
}

func (q *Queue[T]) Push(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.dq.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.dq.PopFront()
}

func (q *Queue[T]) Front() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Front()
}

func (q *Queue[T]) Back() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.Back()
}

func (q *Queue[T]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.dq.Clear()
}

func (q *Queue[T]) String() string {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.dq.String()
}
