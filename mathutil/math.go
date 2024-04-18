package mathutil

import "github.com/hdheid/goutils/common"

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
