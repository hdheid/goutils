package deque

type DequeIterator[T any] struct {
	dq       *Deque[T]
	position int
}

// IsValid 判断该迭代器是否有效
func (it *DequeIterator[T]) IsValid() bool {
	if it.position < 0 || it.position > it.dq.Size() {
		return false
	}

	return true
}

func (it *DequeIterator[T]) Value() T {
	return it.dq.GetIdx(it.position)
}

func (it *DequeIterator[T]) SetValue(idx int, obj T) {
	it.dq.Set(idx, obj)
}

func (it *DequeIterator[T]) Next() *DequeIterator[T] {
	if it.position < it.dq.Size() {
		it.position++
	}
	return it
}

func (it *DequeIterator[T]) Prev() *DequeIterator[T] {
	if it.position > 0 {
		it.position--
	}

	return it
}

// Clone 创建一个新的迭代器
func (it *DequeIterator[T]) Clone() *DequeIterator[T] {
	return &DequeIterator[T]{
		dq:       it.dq,
		position: it.position,
	}
}

func (it *DequeIterator[T]) Euqal(other *DequeIterator[T]) bool {
	if other.dq == it.dq && other.position == it.position {
		return true
	}
	return false
}
