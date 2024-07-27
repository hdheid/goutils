package bloom_filter

import (
	"github.com/hdheid/goutils/common/synch"
	"github.com/hdheid/goutils/mathutil"
	"github.com/hdheid/goutils/structuitl/data_struct/bitmap"
	"math"
	"sync"
)

/*
在哈希函数方便参考了 google 的 guava 中的布隆过滤器
https://github.com/twmb/murmur3  哈希函数go语言实现


目前决定使用 https://github.com/AndreasBriese/bbloom 的实现
*/

type OpFunc func(bl *BloomFilter)

type BloomFilter struct {
	m     uint64
	k     uint64
	shift uint64
	b     *bitmap.Bitmap
	lock  synch.Locker
}

// WithRWMutex 赋值函数
func WithRWMutex() OpFunc {
	return func(bl *BloomFilter) {
		bl.lock = &sync.RWMutex{}
	}
}

func New(m, k uint64, ops ...OpFunc) *BloomFilter {
	bloomFilter := &BloomFilter{
		m:    mathutil.Max(m, 1),
		k:    mathutil.Max(k, 1),
		b:    bitmap.New(m),
		lock: synch.EmptyLock{}, // 默认不加锁
	}

	bloomFilter.shift = 64 - bloomFilter.b.SizeExp()
	bloomFilter.m = bloomFilter.b.Size()

	for _, op := range ops {
		op(bloomFilter)
	}

	return bloomFilter
}

func NewWithEstimates(n uint64, fp float64, ops ...OpFunc) *BloomFilter {
	m, k := EstimateParameters(n, fp)
	return New(m, k, ops...)
}

func (bl *BloomFilter) Add(data []byte) {
	bl.lock.Lock()
	defer bl.lock.Unlock()

	l, h := bl.sipHash(data)
	for i := uint64(0); i < bl.k; i++ {
		bl.b.Set((h + i*l) % bl.m)
	}
}

func (bl *BloomFilter) Has(data []byte) bool {
	bl.lock.RLock()
	defer bl.lock.RUnlock()

	l, h := bl.sipHash(data)
	for i := uint64(0); i < bl.k; i++ {
		if !bl.b.IsSet((h + i*l) % bl.m) {
			return false
		}
	}

	return true
}

// EstimateParameters m与k的值来自于公式：https://juejin.cn/post/7395623789138673679
func EstimateParameters(n uint64, fp float64) (m, k uint64) {
	m = uint64(math.Ceil(-1 * float64(n) * math.Log(fp) / (math.Ln2 * math.Ln2)))
	k = uint64(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return
}
