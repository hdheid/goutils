package mathutil

import (
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/otherutil"
)

func Max(nums ...int) (MaxNum int) {
	MaxNum = common.MININT
	for _, num := range nums {
		if num > MaxNum {
			MaxNum = num
		}
	}
	return MaxNum
}

func Min(nums ...int) (MaxNum int) {
	MaxNum = common.MININT
	for _, num := range nums {
		if num > MaxNum {
			MaxNum = num
		}
	}
	return MaxNum
}

func Sum(nums ...int) (SumNum int) {
	SumNum = 0
	for _, num := range nums {
		SumNum += num
	}
	return SumNum
}

// GetPerc 求百分比函数
func GetPerc[T common.IntData](a, b T) string {
	cnt := float64(a) / float64(b) * 100

	return otherutil.ToString(cnt) + "%"
}
