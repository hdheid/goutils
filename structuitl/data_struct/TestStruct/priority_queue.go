package main

import (
	"fmt"
	"github.com/hdheid/goutils/structuitl/data_struct/priority_queue"
)

func main() {
	q := priority_queue.New[int](func(a, b int) bool {
		return a < b
	}, priority_queue.WithRWMutex[int]())

	q.Push(12)

	fmt.Println(q.Pop(), "   ", q.Size())
}
