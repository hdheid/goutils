package sliceutil

import (
	"github.com/hdheid/goutils/common"
)

/*
	基础数据的切片的基本操作
*/

type CompareFunc[T common.BasicData] func(a, b T) bool

// ContainsWithGenerics 判断某个元素是否存在于对应 carrier 的中，支持基础数据的判断
func ContainsWithGenerics[T common.BasicData](obj T, carrier []T) bool {
	for _, elem := range carrier {
		if elem == obj {
			return true
		}
	}
	return false
}

// InSlice 判断该元素是否存在于切片中
func InSlice[T common.BasicData](obj T, carrier []T) bool {
	return ContainsWithGenerics(obj, carrier)
}

// NotInSlice 判断该元素是否不存在于切片中
func NotInSlice[T common.BasicData](obj T, carrier []T) bool {
	return !ContainsWithGenerics(obj, carrier)
}

// EuqalSlice 判断两个切片是否相等
func EuqalSlice[T common.BasicData](carrierA []T, carrierB []T) bool {
	if len(carrierA) != len(carrierB) {
		return false
	}

	for _, obj := range carrierA {
		if NotInSlice(obj, carrierB) {
			return false
		}
	}

	// todo:索引实现？
	//for i := range carrierA {
	//	if carrierA[i] != carrierB[i] {
	//		return false
	//	}
	//}

	return true
}

// DeepCopySlice 深拷贝，开辟一个新的切片
func DeepCopySlice[T any](carrier []T) []T {
	NewCarrier := make([]T, len(carrier))
	copy(NewCarrier, carrier)

	return NewCarrier
}

/* Reverse函数有两种：
第一种，在原数组上操作：
s == {"a","b","c"}=> s == {"c","b","a"}

第二种,不修改原数组，返回一个新的数组：
s == {"a","b","c"} => s == {"a","b","c"}, s_reverse == {"c","b","a"}
*/

// ReverseSliceWithAny 反转切片
func ReverseSliceWithAny[T any](carrier []T) {
	length := len(carrier)

	for i := 0; i < length/2; i++ {
		carrier[i], carrier[length-i-1] = carrier[length-i-1], carrier[i]
	}
}

// ReverseDeepWithAny 反转函数，深拷贝
func ReverseDeepWithAny[T any](carrier []T) []T {
	length := len(carrier)
	NewCarrier := DeepCopySlice(carrier)

	for i := 0; i < length/2; i++ {
		NewCarrier[i], NewCarrier[length-i-1] = NewCarrier[length-i-1], NewCarrier[i]
	}

	return NewCarrier
}

// UniqueSlice 切片去重函数
func UniqueSlice[T common.NumberData](carrier []T) []T {
	if len(carrier) < 2 {
		return carrier
	}

	uniMap := make(map[T]struct{})
	newList := make([]T, 0, len(carrier))

	for _, obj := range carrier {
		if _, ok := uniMap[obj]; !ok {
			newList = append(newList, obj)
			uniMap[obj] = struct{}{}
		}
	}

	return newList
}

func MergeSlice[T common.BasicData](carrierA, carrierB []T) []T {
	newCarrier := make([]T, 0, len(carrierA)+len(carrierB))

	// 合并数值相同的数据
	for _, objA := range carrierA {
		for _, objB := range carrierB {
			if objB == objA {
				newCarrier = append(newCarrier, objA)
			}
		}
	}

	// 互相不存在对方的数据加入 newCarrier
	for _, obj := range carrierA {
		if NotInSlice(obj, carrierB) {
			newCarrier = append(newCarrier, obj)
		}
	}

	for _, obj := range carrierB {
		if NotInSlice(obj, carrierA) {
			newCarrier = append(newCarrier, obj)
		}
	}

	return newCarrier
}

// SubtractSlice 实现切片A-B的效果
func SubtractSlice[T common.BasicData](carrierA, carrierB []T) []T {
	newCarrier := make([]T, 0, len(carrierA))

	for _, objA := range carrierA {
		if NotInSlice(objA, carrierB) {
			newCarrier = append(newCarrier, objA)
		}

	}

	return newCarrier
}
