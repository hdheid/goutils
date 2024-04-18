package sliceutil

/*
通过函数变量，可以实现结构体、map等切片的查询和合并
*/

type FunEqual[T any] func(a, b T) (T, bool)

func InSliceWithAny[T any](obj T, carrier []T, fnEqual FunEqual[T]) bool {
	for _, elem := range carrier {
		if _, ok := fnEqual(elem, obj); ok {
			return true
		}
	}

	return false
}

func NotInSliceWithAny[T any](obj T, carrier []T, fnEqual FunEqual[T]) bool {
	for _, elem := range carrier {
		if _, ok := fnEqual(elem, obj); ok {
			return false
		}
	}

	return true
}

func MergeSliceWithAny[T any](carrierA, carrierB []T, fnEqual FunEqual[T]) []T {
	newCarrier := make([]T, 0, len(carrierA)+len(carrierB))

	// 合并数值相同的数据
	for _, objA := range carrierA {
		for _, objB := range carrierB {
			if newObj, ok := fnEqual(objA, objB); ok {
				newCarrier = append(newCarrier, newObj)
			}
		}
	}

	// 互相不存在对方的数据加入 newCarrier
	for _, obj := range carrierA {
		if exist := InSliceWithAny(obj, carrierB, fnEqual); !exist {
			newCarrier = append(newCarrier, obj)
		}
	}

	for _, obj := range carrierB {
		if exist := InSliceWithAny(obj, carrierA, fnEqual); !exist {
			newCarrier = append(newCarrier, obj)
		}
	}

	return newCarrier
}

// SubtractSliceWithAny 实现切片A-B的效果
func SubtractSliceWithAny[T any](carrierA, carrierB []T, fnEqual FunEqual[T]) []T {
	newCarrier := make([]T, 0, len(carrierA))

	for _, objA := range carrierA {
		for _, objB := range carrierB {
			if newObj, ok := fnEqual(objA, objB); !ok {
				newCarrier = append(newCarrier, newObj)
			}
		}
	}

	return newCarrier
}
