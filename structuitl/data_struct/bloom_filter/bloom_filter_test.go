package bloom_filter

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddAndHas(t *testing.T) {
	bl := NewWithEstimates(100000, 0.00001, WithRWMutex())
	for i := 0; i < 30000; i++ {
		s := strconv.Itoa(i)
		bl.Add([]byte(s))
	}

	//s := strconv.Itoa(733)
	//fmt.Println(bl.Has([]byte(s)))

	cnt := 0
	for i := 0; i < 100000; i++ {
		s := strconv.Itoa(i)

		if !bl.Has([]byte(s)) {
			if i < 30000 {
				println("has err")
			}
		}

		if bl.Has([]byte(s)) {
			if i >= 30000 {
				println("err", i)
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
