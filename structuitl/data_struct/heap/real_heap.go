package heap

import "github.com/hdheid/goutils/common/compare"

type Heap[T any] struct {
	heap    []T
	cmpFunc compare.CmpFunc[T]
}

func New[T any](cmpFunc compare.CmpFunc[T], data ...T) *Heap[T] {
	h := &Heap[T]{
		heap:    make([]T, 0),
		cmpFunc: cmpFunc,
	}

	h.heap = append(h.heap, data...)

	Init[T](h)

	return h
}

func (h *Heap[T]) Len() int           { return len(h.heap) }
func (h *Heap[T]) Less(i, j int) bool { return h.cmpFunc(h.heap[i], h.heap[j]) > 0 } // j比i大为真
func (h *Heap[T]) Swap(i, j int)      { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }

func (h *Heap[T]) Push(x T) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.heap = append(h.heap, x)
}

func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		panic("heap is empty")
	}
	old := h.heap
	n := len(old)
	x := old[n-1]
	h.heap = old[0 : n-1]
	return x
}

func (h *Heap[T]) Top() T {
	if h.Len() == 0 {
		panic("heap is empty")
	}

	return h.heap[0]
}

func (h *Heap[T]) Clear() {
	h.heap = h.heap[:0] // 清空切片
}
