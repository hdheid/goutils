package priority_queue

import (
	"github.com/hdheid/goutils/common/compare"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	//q := New[int](compare.IntLess) // 小根堆
	//q := New[int](compare.IntMore) // 大根堆
	q := New[int](compare.IntPriorityQueue) // 默认是小到大
	q1 := New[int](compare.IntPriorityQueue, WithRWMutex[int]())

	if q == nil || q1 == nil {
		t.Errorf("s1==nil || s2 ==nil")
		return
	}

	q.Push(5)
	q.Push(2)
	q.Push(3)

	if q.Top() != 2 {
		t.Errorf("q.Push error. q.Top() = %v, should be 3", q.Top())
	}

	q.Pop()
	if q.Top() != 3 {
		t.Errorf("q.Top() = %v, should be 5", q.Top())
	}

	if q.Empty() {
		t.Errorf("q.Empty() should be false")
	}

	if q.Size() != 2 {
		t.Errorf("q.Size() = %v, should be 2", q.Size())
	}

	q.Clear()
	if q.Size() != 0 {
		t.Errorf("q.Clear() error")
	}
}

// 基准测试
func BenchmarkPush(b *testing.B) {
	q := New[int](compare.IntLess) // 小根堆

	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
}
