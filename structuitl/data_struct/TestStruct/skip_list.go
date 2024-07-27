package main

import (
	"fmt"
	"github.com/hdheid/goutils/structuitl/data_struct/skip_list"
)

func main() {
	l := skip_list.New[int, int](func(a, b int) int {
		if a < b {
			return 1
		} else if a == b {
			return 0
		} else {
			return -1
		}
	})

	l.Insert(1, 1)
	_, ok := l.Find(2)
	fmt.Println(ok)
}
