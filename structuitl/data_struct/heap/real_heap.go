package heap

type CmpFunc[T any] func(a, b T) bool

type Heap[T any] struct {
	heap    []T
	cmpFunc CmpFunc[T]
}

func New[T any](cmpFunc CmpFunc[T], data ...T) *Heap[T] {
	h := &Heap[T]{
		heap:    make([]T, 0),
		cmpFunc: cmpFunc,
	}
	for _, datum := range data {
		h.heap = append(h.heap, datum)
	}

	Init[T](h)

	return h
}

func (h *Heap[T]) Len() int           { return len(h.heap) }
func (h *Heap[T]) Less(i, j int) bool { return h.cmpFunc(h.heap[i], h.heap[j]) } // j比i大为真
func (h *Heap[T]) Swap(i, j int)      { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }

func (h *Heap[T]) Push(x T) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.heap = append(h.heap, x)
}

func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		return *new(T)
	}
	old := h.heap
	n := len(old)
	x := old[n-1]
	h.heap = old[0 : n-1]
	return x
}

func (h *Heap[T]) Top() T {
	if h.Len() == 0 {
		return *new(T)
	}

	return h.heap[0]
}

func (h *Heap[T]) Clear() {
	h.heap = h.heap[:0] // 清空切片
}
