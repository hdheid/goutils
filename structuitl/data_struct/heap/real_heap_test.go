package heap

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New[int](compare.IntLess)
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
