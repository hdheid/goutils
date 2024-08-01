package set

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/rbtree"
)

type Iterator[T any] struct {
	it   *rbtree.Node[T, struct{}]
	lock synch.Locker
}

func (iter *Iterator[T]) IsValid() bool {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	return iter.it != nil
}

func (iter *Iterator[T]) Next() *Iterator[T] {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	if iter.IsValid() {
		iter.it = iter.it.Next()
	}

	return iter
}

func (iter *Iterator[T]) Prev() *Iterator[T] {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	if iter.IsValid() {
		iter.it = iter.it.Prev()
	}

	return iter
}

func (iter *Iterator[T]) Val() T {
	iter.lock.RLock()
	defer iter.lock.RUnlock()
	return iter.it.Key()
}

func (iter *Iterator[T]) Clone() *Iterator[T] {
	iter.lock.Lock()
	defer iter.lock.Unlock()
	return &Iterator[T]{
		it:   iter.it,
		lock: iter.lock,
	}
}

// Equal 要求颜色、父节点、左右节点、键值等全部相等
func (iter *Iterator[T]) Equal(other *Iterator[T]) bool {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	return other.it == iter.it
}
