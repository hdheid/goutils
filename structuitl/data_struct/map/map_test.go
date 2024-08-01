package treemap

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"testing"
)

func TestMap(t *testing.T) {
	mp := New[int, int](compare.IntMap)
	mp.Add(1, 3)
	mp.Add(3, 2)

	fmt.Println(mp.Begin().Key())
}

func BenchmarkMap_Add(b *testing.B) {
	m := New[int, struct{}](compare.IntMap)
	length := 100

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			m.Insert(j, struct{}{})
		}
	}
}

func BenchmarkGoMap_Add(b *testing.B) {
	m := make(map[int]struct{})
	length := 100

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			m[j] = struct{}{}
		}
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	m := New[int, struct{}](compare.IntMap)
	length := 10000

	for i := 0; i < length; i++ {
		m.Insert(i, struct{}{})
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			m.Delete(j)
		}
	}
}

func BenchmarkGoMap_Delete(b *testing.B) {
	m := make(map[int]struct{})
	length := 10000

	for i := 0; i < length; i++ {
		m[i] = struct{}{}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			delete(m, j)
		}
	}
}
