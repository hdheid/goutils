package deque

import (
	"fmt"
	"testing"
	"time"
)

func TestDeque(T *testing.T) {

	start := time.Now()

	q := New[int]()
	for i := 0; i < 10000000; i++ {
		//q.PushFront(i)
		q.PushBack(i)
		//q.PopFront()
		//q.PopFront()
	}
	for i := 0; i < 5000000; i++ {
		q.PopFront()
	}

	end := time.Since(start)
	fmt.Println(end.String())
}

// 压力测试
func BenchmarkDeque(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.PushFront(i)
		q.PopFront()
	}
}
