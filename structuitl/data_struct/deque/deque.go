package deque

import "fmt"

type Deque[T any] struct {
	pool             *RingPool[T]
	rings            []*Ring[T]
	begin, end, size int
}

// New todo：后期加个参数，可以指定容量大小
func New[T any]() *Deque[T] {
	return &Deque[T]{
		pool:  newRingPool[T](),
		rings: make([]*Ring[T], 0), //todo：设置为1会如何？或者初始一个ring？
	}
}

func (d *Deque[T]) PushFront(obj T) {
	d.getFirstRing().pushFront(obj)
	d.size++
	if d.ringUsed() >= len(d.rings) { // 扩容防止越界
		d.expand()
	}
}

func (d *Deque[T]) PushBack(obj T) {
	d.GetLastRing().pushBack(obj)
	d.size++
	if d.ringUsed() >= len(d.rings) { // 扩容防止越界
		d.expand()
	}
}

func (d *Deque[T]) PopFront() T {
	if d.Empty() {
		panic("deque is empty")
	}

	firstRing := d.rings[d.begin]

	obj := firstRing.popFront()
	//d.size--
	if firstRing.empty() {
		d.putToPool(firstRing)
		d.rings[d.begin] = nil
		d.begin = d.nextIdx(d.begin)
	}
	d.size-- // 需要放在前面，否则 shrinkDeque() 中 d.size() 会不准确 todo：是否可以尝试，判断一下收缩的条件拓宽一点
	d.shrinkDeque()
	// d.size-- // 将这个往上放，在不断地一次push和pop操作中，时间损耗会越来越大，因为在不断地进行收缩和池收缩操作
	return obj
}

func (d *Deque[T]) PopBack() T {
	if d.Empty() {
		panic("deque is empty")
	}

	lastRing := d.rings[d.preIdx(d.end)]

	obj := lastRing.popBack()

	//d.size--
	if lastRing.empty() {
		d.putToPool(lastRing)
		d.rings[d.preIdx(d.end)] = nil
		d.end = d.preIdx(d.end)
	}
	//d.size--
	d.shrinkDeque()
	d.size-- //同上
	return obj
}

func (d *Deque[T]) Back() T {
	if d.Empty() {
		panic("deque is empty")
	}
	lastRing := d.rings[d.preIdx(d.end)]
	return lastRing.back()
}

func (d *Deque[T]) Front() T {
	if d.Empty() {
		panic("deque is empty")
	}
	firstRing := d.rings[d.begin]
	return firstRing.front()
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

// Clear todo:
func (d *Deque[T]) Clear() {
	d.rings = d.rings[:0]
	d.pool.clear()
	d.begin = 0
	d.end = 0
	d.size = 0
}

func (d *Deque[T]) GetIdx(idx int) T {
	if idx > d.size || idx < 0 {
		panic("out of range")
	}

	ringIdx, pos := d.pos(idx)
	ringIdx = (ringIdx + d.begin) % len(d.rings) // ringIdx 为哪一个环的索引
	return d.rings[ringIdx].getIdx(pos)
}

func (d *Deque[T]) Set(idx int, obj T) {
	if idx < 0 || idx > d.Size() {
		panic("out of range")
	}

	ringIdx, pos := d.pos(idx)
	ringIdx = (ringIdx + d.begin) % len(d.rings) // ringIdx 为哪一个环的索引
	d.rings[ringIdx].setIdx(pos, obj)
}

func (d *Deque[T]) String() string {
	str := "["
	for i := 0; i < d.Size(); i++ {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", d.GetIdx(i))
	}

	str += "]"

	return str
}

// GetFirstRing 找到第一个可以存放数据的环
func (d *Deque[T]) getFirstRing() *Ring[T] {
	if d.ringUsed() >= len(d.rings) { // 保证环的个数始终比用到的多
		d.expand()
	}

	if d.rings[d.begin] == nil || d.rings[d.begin].isFull() { // 如果第一个环满了或者没有第一个环，则从池中新建一个
		s := d.pool.get()
		d.rings[d.preBegin()] = s // 需要防止越界，更新之前需要判断是否需要扩容
		return s
	}

	if d.rings[d.begin] != nil && !d.rings[d.begin].isFull() {
		return d.rings[d.begin]
	}

	return d.rings[d.begin]
}

// GetLastRing 找到找到最后一个可以存放数据的环
func (d *Deque[T]) GetLastRing() *Ring[T] {
	if d.ringUsed() >= len(d.rings) { // 保证环的个数始终比用到的多
		d.expand()
	}

	idx := d.preIdx(d.end)
	if d.rings[idx] == nil || d.rings[idx].isFull() { // 如果最后一个环满了或者没有第一个环，则从池中新建一个
		s := d.pool.get()
		d.rings[d.end] = s // 需要防止越界，更新之前需要判断是否需要扩容
		d.nextEnd()
		return s
	}

	if d.rings[idx] != nil && !d.rings[idx].isFull() {
		return d.rings[idx]
	}

	return d.rings[idx]
}

func (d *Deque[T]) preBegin() int {
	d.begin = d.preIdx(d.begin) // 以一个环的形式进行加减
	return d.begin
}

func (d *Deque[T]) nextEnd() int {
	d.end = d.nextIdx(d.end) // 以一个环的形式进行加减
	return d.end
}

func (d *Deque[T]) preIdx(lastIdx int) int {
	return (lastIdx - 1 + d.cap()) % d.cap() // 以一个环的形式进行加减
}

func (d *Deque[T]) nextIdx(lastIdx int) int {
	return (lastIdx + 1 + d.cap()) % d.cap() // 以一个环的形式进行加减
}

func (d *Deque[T]) cap() int {
	return len(d.rings)
}

func (d *Deque[T]) ringUsed() int {
	if d.size == 0 {
		return 0
	}
	if d.end > d.begin {
		return d.end - d.begin
	} else {
		return d.cap() - d.begin + d.end
	}
}

// 扩容，两倍扩容
func (d *Deque[T]) expand() {
	newCap := d.ringUsed() * 2
	if newCap == 0 {
		newCap = 1
	}
	newRings := make([]*Ring[T], newCap)

	for i := 0; i < d.ringUsed(); i++ {
		ringsIdx := (d.begin + i) % d.ringUsed()
		newRings[i] = d.rings[ringsIdx]
	}
	posEnd := d.ringUsed()

	d.begin = 0
	d.end = posEnd
	d.rings = newRings
}

// 缩小，两倍缩小
func (d *Deque[T]) shrinkDeque() {
	if int(float64(d.ringUsed()*2)*1.2) < len(d.rings) { // 如果空余的内存超过一定阈值，就进行缩小操作，需要留出一点空间避频繁的删除
		newCap := len(d.rings) / 2
		newRings := make([]*Ring[T], newCap)

		for i := 0; i < d.ringUsed(); i++ {
			ringsIdx := (d.begin + i) % len(d.rings)
			newRings[i] = d.rings[ringsIdx]
		}

		d.begin = 0
		d.end = d.ringUsed()
		d.rings = newRings
	}
}

// 将空环存入池中，定期收缩池确保多余的环数量合理
func (d *Deque[T]) putToPool(s *Ring[T]) {
	s.clear()
	d.pool.put(s)

	if d.pool.size()*6/5 > d.ringUsed() {
		d.pool.shrinkPool(d.ringUsed() / 5)
	}
}

func (d *Deque[T]) pos(pos int) (ringIdx, idx int) {
	if pos < d.rings[d.begin].len() {
		return 0, pos
	}

	// 第一个可能没有装满
	pos -= d.rings[d.begin].len()
	return pos/DefaultCap + 1, pos % DefaultCap
}

// 迭代器

func (d *Deque[T]) Begin() *DequeIterator[T] {
	return d.ItAt(0)
}

func (d *Deque[T]) End() *DequeIterator[T] {
	return d.ItAt(d.Size())
}

func (d *Deque[T]) First() *DequeIterator[T] {
	return d.ItAt(0)
}

func (d *Deque[T]) Last() *DequeIterator[T] {
	return d.ItAt(d.Size() - 1)
}

func (d *Deque[T]) ItAt(idx int) *DequeIterator[T] {
	return &DequeIterator[T]{
		dq:       d,
		position: idx,
	}
}
