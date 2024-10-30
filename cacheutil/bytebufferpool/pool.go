package bytebufferpool

import (
	"github.com/hdheid/goutils/common"
	"sort"
	"sync"
	"sync/atomic"
)

/*
pool 的实现原理是对不同长度的字节数组划分出20个区间，根据数据大小，对不同长度的数据放在能容纳的最小组里面。一旦某一个区间的字节数组使用次数超过了42000次，就会进行 defaultSize 校准，避免了切片频繁扩容造成性能抖动，同时将占用最大的 %5 的字节数组扔给 GC 进行回收，不放回池中。
*/

type Pool struct {
	pool sync.Pool

	calls       [common.Steps]uint64 //不同区间的字节数组使用次数
	calibrating uint64               // 判断是否正在进行校准，用于并发场景下，无锁化

	defaultSize uint64 // 新建字节数组的默认长度
	maxSize     uint64 // 可以放回池中的最大字节数组长度
}

// 默认生成的 Pool
var defaultPool Pool

func Get() *ByteBuffer  { return defaultPool.Get() }
func Put(b *ByteBuffer) { defaultPool.Put(b) }

func (p *Pool) Get() *ByteBuffer {
	v := p.pool.Get()
	if v != nil {
		return v.(*ByteBuffer)
	}

	return &ByteBuffer{
		make([]byte, 0, atomic.LoadUint64(&p.defaultSize)),
	}
}

func (p *Pool) Put(b *ByteBuffer) {
	idx := index(len(b.data))

	// 使用了原子操作避免出现并发安全问题
	if atomic.AddUint64(&p.calls[idx], 1) > common.CalibrateCallsThreshold {
		// 进行校准操作
		p.calibrate()
	}

	maxSize := int(atomic.LoadUint64(&p.maxSize))
	if maxSize == 0 || cap(b.data) <= maxSize { // 所有超出最大值的都不放回池中
		b.Reset()
		p.pool.Put(b)
	}
}

func (p *Pool) calibrate() {
	// 通过原子操作 CAS 来代替锁
	if atomic.CompareAndSwapUint64(&p.calibrating, 0, 1) {
		return
	}

	a := make(callSizes, 0, common.Steps)
	var callsSum uint64 // 放入池操作总数

	// 清空 pool 放入次数，并将值插入 a 中
	for i := uint64(0); i < common.Steps; i++ {
		calls := atomic.SwapUint64(&p.calls[i], 0)
		callsSum += calls
		a = append(a, callSize{
			calls: calls,
			size:  common.MinSize << i,
		})
	}

	// 使用 pdqsort 排序算法实现
	sort.Sort(a)

	// 将默认大小设置为放入池次数最多的字节数组大小
	defaultSize := a[0].size
	maxSize := defaultSize // maxSize 最后应为前 %95 次放入池中 的字节数组的最大值

	/* 换句话说，假设 put 了一百次字节数组，那么将每一个区间大小的字节数组根据 put 次数进行排序后，取前 95 次 put 的字节数组中，长度最大的那一个 */
	maxSum := uint64(float64(callsSum) * common.MaxPercentile)
	callsSum = 0
	for i := uint64(0); i < common.Steps; i++ {
		if callsSum > maxSum {
			break
		}

		callsSum += a[i].calls
		if a[i].size > maxSize {
			maxSize = a[i].size
		}
	}

	atomic.StoreUint64(&p.defaultSize, defaultSize)
	atomic.StoreUint64(&p.maxSize, maxSize)
	atomic.StoreUint64(&p.calibrating, 0)
}

type callSize struct {
	calls uint64
	size  uint64
}

type callSizes []callSize

func (c callSizes) Len() int {
	return len(c)
}

// Less 表示从大到小排序
func (c callSizes) Less(i, j int) bool {
	return c[i].calls > c[j].calls
}

func (c callSizes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func index(n int) int {
	n--
	n >>= common.MinBitSize
	idx := 0
	for n > 0 {
		n >>= 1
		idx++
	}
	if idx >= common.Steps {
		idx = common.Steps - 1
	}

	return idx
}
