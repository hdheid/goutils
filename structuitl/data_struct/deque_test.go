package data_struct

import (
	"fmt"
	"github.com/hdheid/goutils/structuitl/data_struct/deque"
	"testing"
	"time"
)

func TestDeque(T *testing.T) {

	start := time.Now()

	q := deque.New[int]()
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
