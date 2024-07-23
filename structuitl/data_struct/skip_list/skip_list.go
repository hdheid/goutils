package skip_list

import (
	"math/rand"
	"sync"
	"time"
)

/*
跳表是一个....

下面实现方式为：


参考：https://www.cnblogs.com/linvanda/p/17538558.html
*/

const (
	DefaultLevel = 10
	Skip_list_P  = 0.5
)

type keyCmp[T any] func(a, b T) int
type visitor[K, V any] func(key K, val V) bool

// Node 跳表节点
type Node[K, V any] struct {
	key  K
	val  V
	next []*Node[K, V]
}

type SkipList[K, V any] struct {
	head     Node[K, V]
	keyCmp   keyCmp[K]
	maxLevel int
	len      int
	random   *rand.Rand
	lock     sync.RWMutex // 读写锁
}

type Elem[K, V any] struct {
	key K
	val V
}

func New[K, V any](cmp keyCmp[K]) *SkipList[K, V] {
	l := &SkipList[K, V]{
		maxLevel: DefaultLevel,
		keyCmp:   cmp,
		lock:     sync.RWMutex{},
		random:   rand.New(rand.NewSource(time.Now().Unix())),
	}
	l.head.next = make([]*Node[K, V], l.maxLevel)

	return l
}

// Insert 增加节点，兼顾修改值
func (sl *SkipList[K, V]) Insert(key K, val V) {
	sl.lock.Lock() // 上锁
	defer sl.lock.Unlock()

	// 找到所有前驱节点
	prevs := sl.findPrevNodes(key)

	if prevs[0].next[0] != nil && sl.keyCmp(prevs[0].next[0].key, key) == 0 {
		// 更新值
		prevs[0].next[0].val = val
		return
	}

	level := sl.randomLevel()

	// 增加一个节点
	nd := &Node[K, V]{
		key:  key,
		val:  val,
		next: make([]*Node[K, V], level),
	}

	// 将节点加入到跳表中
	for i := range nd.next {
		nd.next[i] = prevs[i].next[i]
		prevs[i].next[i] = nd
	}

	sl.len++
}

// Remove 删除节点
func (sl *SkipList[K, V]) Remove(key K) bool {
	sl.lock.Lock()
	defer sl.lock.Unlock()

	prevs := sl.findPrevNodes(key)

	dNode := prevs[0].next[0]
	if dNode == nil || sl.keyCmp(dNode.key, key) != 0 {
		return false
	}

	// 删除节点
	for i := range dNode.next {
		prevs[i].next[i] = dNode.next[i]
		dNode.next[i] = nil // 消除掉
	}

	sl.len--
	return true
}

func (sl *SkipList[K, V]) Find(key K) (val V) {
	sl.lock.RLock()
	defer sl.lock.RUnlock()

	node := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		cur := node.next[i]
		for ; cur != nil; cur = cur.next[i] {
			if sl.keyCmp(cur.key, key) == 0 {
				return cur.val
			}
			if sl.keyCmp(cur.key, key) > 0 {
				break
			}
			node = cur
		}
	}

	if node == nil {
		return *new(V)
	}

	if sl.keyCmp(node.key, key) == 0 {
		return node.val
	}

	return *new(V)
}

func (sl *SkipList[K, V]) Len() int {
	sl.lock.RLock()
	defer sl.lock.RUnlock()

	return sl.len
}

// 设置随机层数
func (sl *SkipList[K, V]) randomLevel() (level int) {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	level = 1
	for rand.Float64() < Skip_list_P && level < sl.maxLevel {
		level++
	}
	return
}

// 获取每一层的前驱节点
func (sl *SkipList[K, V]) findPrevNodes(key K) []*Node[K, V] {
	prevs := make([]*Node[K, V], sl.maxLevel)
	prev := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		if sl.head.next[i] != nil {
			for node := prev.next[i]; node != nil; node = node.next[i] {
				if sl.keyCmp(node.key, key) >= 0 {
					break
				}
				prev = node // 找到所有前驱节点
			}
		}
		prevs[i] = prev
	}

	return prevs
}

// Traversal 遍历跳表
func (sl *SkipList[K, V]) Traversal() []Elem[K, V] {
	sl.lock.RLock()
	defer sl.lock.RUnlock()

	datas := make([]Elem[K, V], 0)
	for n := sl.head.next[0]; n != nil; n = n.next[0] {
		datas = append(datas, Elem[K, V]{
			n.key,
			n.val,
		})
	}

	return datas
}
