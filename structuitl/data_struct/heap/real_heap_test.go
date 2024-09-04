package heap

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"sort"
	"testing"
)

func TestRealHeap(t *testing.T) {
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

func TestHeapNew(t *testing.T) {
	h1 := New[int](compare.IntLess, 9, 4, 8, 7) //小根堆
	//h1 := New[int](compare.IntMore)

	if h1 == nil {
		t.Errorf("s1==nil || s2 ==nil")
	}
}

func TestHeap(t *testing.T) {
	h := New[int](compare.IntLess, 9, 4, 8, 7, 6) //小根堆

	if h.Top() != 4 {
		t.Errorf("h.Top() = %v, should be 4", h.Top())
	}

	if h.Len() != 5 {
		t.Errorf("h.Len() = %v, should be 5", h.Len())
	}

	Push[int](h, 0)
	if h.Top() != 0 {
		t.Errorf("h.Push() error")
	}

	if Pop[int](h) != 0 {
		t.Errorf("h.Pop() = %v, should be 0", h.Pop())
	}

	if !h.Less(1, 3) {
		t.Errorf("h.Less(1,3) = %v, should be true", h.Less(1, 3))
	} //小根堆，因此1比3小

	h.Clear()
	if h.Len() != 0 {
		t.Errorf("h.Clear() = %v, should be 0", h.Len())
	}

	defer func() {
		if r := recover(); r != nil {
			if r != "heap is empty" {
				t.Errorf("r = %v, should be 'heap is empty'", r)
			}
		}
	}()
	h.Pop()
}

func TestHeapTop(t *testing.T) {
	h := New[int](compare.IntLess) //小根堆

	defer func() {
		if r := recover(); r != nil {
			if r != "heap is empty" {
				t.Errorf("r = %v, should be 'heap is empty'", r)
			}
		}
	}()

	h.Top()
}
