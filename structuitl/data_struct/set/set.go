package set

import (
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/structuitl/data_struct/rbtree"
	"sync"
)

type OpFunc[T any] func(s *Set[T])

type Set[T any] struct {
	tree *rbtree.RbTree[T, struct{}]
	cmp  compare.CmpFunc[T]
	lock synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex[T any]() OpFunc[T] {
	return func(s *Set[T]) {
		s.lock = &sync.RWMutex{}
	}
}

func New[T any](cmp compare.CmpFunc[T], ops ...OpFunc[T]) *Set[T] {
	s := &Set[T]{
		tree: rbtree.New[T, struct{}](cmp),
		cmp:  cmp,
		lock: synch.EmptyLock{},
	}

	for _, op := range ops {
		op(s)
	}

	return s
}

func (s *Set[T]) Insert(val T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.tree.Find(val)
	if node != nil {
		return // 存在则不能插入（集合）
	}

	s.tree.Insert(val, struct{}{})
}

// Add 与insert函数一样
func (s *Set[T]) Add(val T) {
	s.Insert(val)
}

func (s *Set[T]) Exist(val T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	node := s.tree.Find(val)
	if node != nil {
		return true
	}

	return false
}

func (s *Set[T]) Delete /*Erase*/ (val T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.tree.Find(val)
	if node != nil {
		s.tree.Delete(node)
	}
}

func (s *Set[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.tree.Clear()
}

func (s *Set[T]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.tree.Size()
}

// 迭代器操作

func (s *Set[T]) Begin() *Iterator[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &Iterator[T]{
		it:   s.tree.Begin(),
		lock: s.lock, // 存的是锁的地址，因此赋值后，不会生成一个新地锁，还是同一把锁
	}
}

func (s *Set[T]) Find(val T) *Iterator[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	node := s.tree.Find(val)
	return &Iterator[T]{
		it:   node,
		lock: s.lock,
	}
}

func (s *Set[T]) Erase(iter *Iterator[T]) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.tree.Delete(iter.it)
}

// Intersect 取两个 set 的交集。如果任一 set 是并发安全的，那么新的 set 也是并发安全的
func (s *Set[T]) Intersect(set *Set[T]) *Set[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	// todo:判断两个set是不是一个类型

	newSet := New[T](s.cmp)
	if s.IsLock() || set.IsLock() {
		// 加锁函数
		newSet.SetLock()
	}

	sIt := s.Begin()
	setIt := set.Begin()

	for sIt.IsValid() && setIt.IsValid() {
		switch s.cmp(sIt.Val(), setIt.Val()) {
		case -1:
			setIt.Next()
		case 1:
			sIt.Next()
		default:
			newSet.tree.Insert(sIt.Val(), struct{}{})
			sIt.Next()
			setIt.Next()
		}
	}

	return newSet
}

// Union 取两个 set 的并集。如果任一 set 是并发安全的，那么新的 set 也是并发安全的
func (s *Set[T]) Union(set *Set[T]) *Set[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	newSet := New[T](s.cmp)
	if s.IsLock() || set.IsLock() {
		// 加锁函数
		newSet.SetLock()
	}

	sIt := s.Begin()
	setIt := set.Begin()

	// 使用 tree.insert 相比于 直接使用 insert，减少了一次 find 开销，提高了效率
	for sIt.IsValid() && setIt.IsValid() {
		switch s.cmp(sIt.Val(), setIt.Val()) {
		case -1:
			newSet.tree.Insert(setIt.Val(), struct{}{})
			setIt.Next()
		case 1:
			newSet.tree.Insert(sIt.Val(), struct{}{})
			sIt.Next()
		default:
			newSet.tree.Insert(sIt.Val(), struct{}{})
			sIt.Next()
			setIt.Next()
		}
	}

	for ; setIt.IsValid(); setIt.Next() {
		newSet.tree.Insert(setIt.Val(), struct{}{})
	}

	for ; sIt.IsValid(); sIt.Next() {
		newSet.tree.Insert(sIt.Val(), struct{}{})
	}

	return newSet
}

// Diff 去差集，存在于 s 中但是不存在于 set 中的元素。
// 如果任一 set 是并发安全的，那么新的 set 也是并发安全的
func (s *Set[T]) Diff(set *Set[T]) *Set[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	newSet := New[T](s.cmp)
	if s.IsLock() || set.IsLock() {
		// 加锁函数
		newSet.SetLock()
	}

	sIt := s.Begin()
	setIt := set.Begin()

	for sIt.IsValid() && setIt.IsValid() {
		switch s.cmp(sIt.Val(), setIt.Val()) {
		case -1:
			setIt.Next()
		case 1:
			newSet.tree.Insert(sIt.Val(), struct{}{})
			sIt.Next()
		default:
			sIt.Next()
			setIt.Next()
		}
	}

	for ; sIt.IsValid(); sIt.Next() {
		newSet.tree.Insert(sIt.Val(), struct{}{})
	}

	return newSet
}

// IsLock 判断 set 是否是并发安全的
func (s *Set[T]) IsLock() bool {
	return s.lock != synch.EmptyLock{}
}

// SetLock 给 set 加锁
func (s *Set[T]) SetLock() {
	s.lock = &sync.RWMutex{}
}

// UnSetLock 给 set 去掉锁
func (s *Set[T]) UnSetLock() {
	s.lock = synch.EmptyLock{}
}
