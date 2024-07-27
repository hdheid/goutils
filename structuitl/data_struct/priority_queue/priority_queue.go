package priority_queue

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/heap"
	"sync"
)

type OpFunc[T any] func(bl *PriorityQueue[T])

type PriorityQueue[T any] struct {
	h    *heap.Heap[T]
	lock synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex[T any]() OpFunc[T] {
	return func(q *PriorityQueue[T]) {
		q.lock = &sync.RWMutex{}
	}
}

func New[T any](cmp heap.CmpFunc[T], ops ...OpFunc[T]) *PriorityQueue[T] {
	pQueue := &PriorityQueue[T]{
		h:    heap.New[T](cmp),
		lock: synch.EmptyLock{},
	}

	for _, op := range ops {
		op(pQueue)
	}

	return pQueue
}

func (q *PriorityQueue[T]) Push(elem T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	heap.Push[T](q.h, elem)
}

func (q *PriorityQueue[T]) Pop() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	return heap.Pop[T](q.h)
}

func (q *PriorityQueue[T]) Top() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.h.Top()
}

func (q *PriorityQueue[T]) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.h.Len() == 0
}

func (q *PriorityQueue[T]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.h.Len()
}

func (q *PriorityQueue[T]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	// reset cap to zero
	q.h.Clear()
}
