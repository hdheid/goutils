package deque

import "fmt"

const DefaultCap = 128

type Ring[T any] struct {
	data             []T
	begin, end, size int
}

// NewRing 定义个容量为 capacity 的环
func newRing[T any](capacity int) *Ring[T] {
	return &Ring[T]{
		data: make([]T, capacity),
	}
}

func (r *Ring[T]) pushFront(obj T) {
	r.data[r.preBegin()] = obj
	r.size++
}

func (r *Ring[T]) pushBack(obj T) {
	r.data[r.end] = obj
	r.nextEnd()
	r.size++
}

func (r *Ring[T]) popFront() T {
	obj := r.data[r.begin]
	r.begin = r.nextIdx(r.begin)
	r.size--
	return obj
}

func (r *Ring[T]) popBack() T {
	r.end = r.preIdx(r.end)
	r.size--
	return r.data[r.end]
}

func (r *Ring[T]) isFull() bool {
	return r.size == r.cap()
}

func (r *Ring[T]) preBegin() int {
	r.begin = r.preIdx(r.begin) // 以一个环的形式进行加减
	return r.begin
}

func (r *Ring[T]) nextEnd() int {
	r.end = r.nextIdx(r.end) // 以一个环的形式进行加减
	return r.end
}

func (r *Ring[T]) preIdx(lastIdx int) int {
	return (lastIdx - 1 + r.cap()) % r.cap() // 以一个环的形式进行加减
}

func (r *Ring[T]) nextIdx(lastIdx int) int {
	return (lastIdx + 1 + r.cap()) % r.cap() // 以一个环的形式进行加减
}

// ringIdx 返回给出索引在环中的索引位置
func (r *Ring[T]) ringIdx(idx int) int {
	return (r.begin + idx) % r.cap()
}

func (r *Ring[T]) cap() int {
	return len(r.data)
}

func (r *Ring[T]) clear() {
	r.begin = 0
	r.end = 0
	r.size = 0
}

// insert 在 posIdx 处插入一个数据，环的数据插入，由于是一个环，如果插入地方离头近，则移动前半部分，否则移动后半部分
func (r *Ring[T]) insert(posIdx int, obj T) {
	if posIdx*2 < r.size {
		idx := r.preIdx(r.begin)
		for i := 0; i < posIdx; i++ {
			r.data[idx] = r.data[r.nextIdx(idx)]
			idx = r.nextIdx(idx)
		}
		r.data[idx] = obj
		r.preBegin()
	} else {
		idx := r.end
		for i := 0; i < r.size-posIdx; i++ {
			r.data[idx] = r.data[r.preIdx(idx)]
			idx = r.preIdx(idx)
		}
		r.data[idx] = obj
		r.nextEnd()
	}

	r.size++
}

func (r *Ring[T]) getIdx(idx int) T {
	if idx < 0 || idx > r.size {
		panic(fmt.Errorf("out of range"))
	}

	return r.data[r.ringIdx(idx)]
}

func (r *Ring[T]) setIdx(idx int, obj T) {
	if idx < 0 || idx > r.size {
		panic(fmt.Errorf("out of range"))
	}

	r.data[r.ringIdx(idx)] = obj
}

func (r *Ring[T]) back() T {
	return r.data[r.preIdx(r.end)]
}

func (r *Ring[T]) front() T {
	return r.data[r.begin]
}

func (r *Ring[T]) len() int {
	return r.size
}

func (r *Ring[T]) empty() bool {
	return r.size == 0
}
