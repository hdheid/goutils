package priority_queue

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	//q := New[int](compare.IntLess) // 小根堆
	//q := New[int](compare.IntMore) // 大根堆
	q := New[int](compare.IntPriorityQueue)

	q.Push(5)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Top())
}

// 基准测试
func BenchmarkPush(b *testing.B) {
	q := New[int](compare.IntLess) // 小根堆

	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
}
