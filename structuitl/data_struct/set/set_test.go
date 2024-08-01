package set

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/common/synch"
	"testing"
)

func TestLock(t *testing.T) {
	s := New[int](compare.IntLess)

	var emptyLock synch.EmptyLock
	fmt.Println(s.lock == emptyLock)
	fmt.Println(s.lock == nil)
}

func BenchmarkSet_Add(b *testing.B) {
	s := New[int](compare.IntLess)
	length := 1000000

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			s.Add(j)
		}
	}
}
