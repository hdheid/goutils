package treemap

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/rbtree"
)

type Iterator[K, V any] struct {
	it   *rbtree.Node[K, V]
	lock synch.Locker
}

func (iter *Iterator[K, V]) IsValid() bool {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	return iter.it != nil
}

func (iter *Iterator[K, V]) Next() *Iterator[K, V] {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	if iter.IsValid() {
		iter.it = iter.it.Next()
	}

	return iter
}

func (iter *Iterator[K, V]) Prev() *Iterator[K, V] {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	if iter.IsValid() {
		iter.it = iter.it.Prev()
	}

	return iter
}

func (iter *Iterator[K, V]) Key() K {
	iter.lock.RLock()
	defer iter.lock.RUnlock()
	return iter.it.Key()
}

func (iter *Iterator[K, V]) Val() V {
	iter.lock.RLock()
	defer iter.lock.RUnlock()
	return iter.it.Val()
}

func (iter *Iterator[K, V]) SetVal(val V) {
	iter.lock.Lock()
	defer iter.lock.Unlock()
	iter.it.SetVal(val)
}

func (iter *Iterator[K, V]) Clone() *Iterator[K, V] {
	iter.lock.Lock()
	defer iter.lock.Unlock()
	return &Iterator[K, V]{
		it:   iter.it,
		lock: iter.lock,
	}
}

// Equal 要求颜色、父节点、左右节点、键值等全部相等
func (iter *Iterator[K, V]) Equal(other *Iterator[K, V]) bool {
	iter.lock.RLock()
	defer iter.lock.RUnlock()

	return other.it == iter.it
}
