package otherutil

import (
	"fmt"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	start1 := time.Now()

	bytes := []byte{'q', 'w', 'f', 't', 'y'}

	for i := 0; i < 1e7; i++ {
		_ = string(bytes)
	}

	end1 := time.Since(start1)
	start2 := time.Now()

	for i := 0; i < 1e7; i++ {
		_ = ToString(bytes)
	}
	end2 := time.Since(start2)
	fmt.Println("耗时分别为：", end1, "  ,  ", end2)

	//for i, _ := range bytes {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
}
