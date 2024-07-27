package heap

import (
	"fmt"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	cmpFunc := func(a, b int) bool {
		return a < b // 堆顶是最小的
	}

	h := New[int](cmpFunc)
	a := make([]int, 0)

	for i := 0; i < 10; i++ {
		//num := rand.Int()
		a = append(a, i)
		Push[int](h, i)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j] // 堆顶是最小的
	})

	fmt.Println(Pop[int](h), "   ", a[len(a)-1])
}
