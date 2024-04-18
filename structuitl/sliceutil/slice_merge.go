package sliceutil

/*
通过函数变量，可以实现结构体、map等切片的查询和合并
*/

type FunEqual[T any] func(a, b T) (T, bool)

func FindSlice[T any](obj T, carrier []T, fnEqual FunEqual[T]) bool {
	for _, elem := range carrier {
		if _, ok := fnEqual(elem, obj); ok {
			return true
		}
	}

	return false
}

func MergeSlice[T any](carrierA, carrierB []T, fnEqual FunEqual[T]) []T {
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
		if exist := FindSlice(obj, carrierB, fnEqual); !exist {
			newCarrier = append(newCarrier, obj)
		}
	}

	for _, obj := range carrierB {
		if exist := FindSlice(obj, carrierA, fnEqual); !exist {
			newCarrier = append(newCarrier, obj)
		}
	}

	return newCarrier
}
