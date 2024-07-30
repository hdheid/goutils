package main

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/structuitl/data_struct/priority_queue"
)

func main() {
	q := priority_queue.New[int](compare.IntLess, priority_queue.WithRWMutex[int]()) //less表示小的在前面

	q.Push(12)
	q.Push(88)

	fmt.Println(q.Pop())
}
