package deque

// RingPool 作为存放环的池，避免环大量重复的创建与删除
type RingPool[T any] struct {
	rings []*Ring[T]
}

func newRingPool[T any]() *RingPool[T] {
	return &RingPool[T]{
		rings: make([]*Ring[T], 0),
	}
}

func (p *RingPool[T]) get() *Ring[T] {
	if len(p.rings) == 0 {
		return newRing[T](DefaultCap)
	}
	ring := p.rings[0]
	p.rings = p.rings[1:]
	return ring
}

func (p *RingPool[T]) put(r *Ring[T]) {
	p.rings = append(p.rings, r)
}

func (p *RingPool[T]) size() int {
	return len(p.rings)
}

func (p *RingPool[T]) shrinkPool(newSize int) {
	if len(p.rings) > newSize {
		newRings := make([]*Ring[T], newSize)
		copy(newRings, p.rings)
		p.rings = newRings
	}
}

func (p *RingPool[T]) clear() {
	p.rings = p.rings[:0]
}
