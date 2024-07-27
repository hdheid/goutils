package priority_queue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := New[int](func(a, b int) bool {
		return a < b // 堆顶是最小的
	})

	q.Push(5)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Top())
}
