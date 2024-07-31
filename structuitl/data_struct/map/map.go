package treemap

import (
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/rbtree"
	"sync"
)

/*
线程安全的分片加锁map：https://github.com/orcaman/concurrent-map
*/

type OpFunc[K, V any] func(m *Map[K, V])

type Map[K, V any] struct {
	tree *rbtree.RbTree[K, V]
	lock synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex[K, V any]() OpFunc[K, V] {
	return func(m *Map[K, V]) {
		m.lock = &sync.RWMutex{}
	}
}

func New[K, V any](cmp compare.CmpFunc[K], ops ...OpFunc[K, V]) *Map[K, V] {
	q := &Map[K, V]{
		tree: rbtree.New[K, V](cmp),
		lock: synch.EmptyLock{},
	}

	for _, op := range ops {
		op(q)
	}

	return q
}

func (m *Map[K, V]) Insert(key K, val V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	node := m.tree.Find(key)
	if node != nil {
		node.SetVal(val) // 若键值存在则更新值
		return
	}

	m.tree.Insert(key, val)
}

func (m *Map[K, V]) Update(key K, val V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	node := m.tree.Find(key)
	if node != nil {
		node.SetVal(val) // 若键值存在则更新值
	}

	return
}

// Add 与insert函数一样
func (m *Map[K, V]) Add(key K, val V) {
	m.Insert(key, val)
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	node := m.tree.Find(key)
	if node != nil {
		return node.Val(), true
	}

	return *new(V), false
}

func (m *Map[K, V]) Delete /*Erase*/ (key K) {
	m.lock.Lock()
	defer m.lock.Unlock()

	node := m.tree.Find(key)
	if node != nil {
		m.tree.Delete(node)
	}
}

func (m *Map[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.tree.Clear()
}

func (m *Map[K, V]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.tree.Size()
}

// 迭代器操作

func (m *Map[K, V]) Begin() *Iterator[K, V] {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return &Iterator[K, V]{
		it:   m.tree.Begin(),
		lock: m.lock, // 存的是锁的地址，因此赋值后，不会生成一个新地锁，还是同一把锁
	}
}

func (m *Map[K, V]) Find(key K) *Iterator[K, V] {
	m.lock.RLock()
	defer m.lock.RUnlock()

	node := m.tree.Find(key)
	return &Iterator[K, V]{
		it:   node,
		lock: m.lock,
	}
}

func (m *Map[K, V]) Erase(iter *Iterator[K, V]) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.tree.Delete(iter.it)
}
