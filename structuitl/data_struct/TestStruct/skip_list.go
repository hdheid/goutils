package main

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/structuitl/data_struct/skip_list"
)

func main() {
	l := skip_list.New[int, int](compare.IntSkipList)

	l.Insert(1, 1)
	_, ok := l.Find(1)
	fmt.Println(ok)
}
