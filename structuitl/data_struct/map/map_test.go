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
