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

// InsertSliceBeforeWithAny 将数据插入切片 index 之前，如果 index 小于当前切片长度，则插入头部，反之插入尾部
func InsertSliceBeforeWithAny[T any](carrier []T, index int, values ...T) []T {
	newCarrier := make([]T, 0, len(carrier)+len(values))

	if index < 0 {
		index = 0
	}
	if index >= len(carrier) {
		index = len(carrier) - 1
	}

	copy(newCarrier, carrier[:index])
	copy(newCarrier[index:], values)
	copy(newCarrier[index+len(values):], carrier[index:])

	return newCarrier
}

// InsertSliceAfterWithAny 将数据插入切片 index 之后，如果 index 小于当前切片长度，则插入头部，反之插入尾部
func InsertSliceAfterWithAny[T any](carrier []T, index int, values ...T) []T {
	newCarrier := make([]T, 0, len(carrier)+len(values))

	if index < 0 {
		index = 0
	}
	if index >= len(carrier) {
		index = len(carrier) - 1
	}

	copy(newCarrier, carrier[:index+1])
	copy(newCarrier[index+1:], values)
	copy(newCarrier[index+len(values)+1:], carrier[index+1:])

	return newCarrier
}

// DeleteIndSliceWithAny 删除指定索引处的数据，如果传入的索引不在切片范围内则返回原切片
//func DeleteIndSliceWithAny[T any](carrier []T, index int) []T {
//	var length = len(carrier)
//
//	if index < 0 || index >= length {
//		return carrier
//	}
//
//	if index == 0 {
//		return carrier[1:]
//	}
//	if index == length-1 {
//		return carrier[:index]
//	}
//
//	return append(carrier[:index], carrier[index+1:]...)
//}

// DeleteIndSliceWithAny 移位法 删除指定索引处的数据，如果传入的索引不在切片范围内则返回原切片
func DeleteIndSliceWithAny[T any](carrier []T, index int) []T {
	var length = len(carrier)

	if index < 0 || index >= length {
		return carrier
	}

	if index == 0 {
		return carrier[1:]
	}
	if index == length-1 {
		return carrier[:index]
	}

	var j int
	for i := range carrier {
		if i != index {
			carrier[j] = carrier[i]
			j++
		}
	}

	return carrier[:j]
}

// DeleteMultiIndSliceWithAny 删除指定索引（索引可为多个）处的数据，如果传入的索引不在切片范围内则返回原切片
func DeleteMultiIndSliceWithAny[T any](carrier []T, index ...int) []T {
	var length = len(carrier)

	for _, ind := range index {
		if ind < 0 || ind >= length {
			return carrier
		}
	}

	if len(index) == 1 && index[0] == 0 {
		return carrier[1:]
	}
	if len(index) == 1 && index[0] == length-1 {
		return carrier[:index[0]]
	}

	for _, ind := range index {
		carrier = DeleteIndSliceWithAny(carrier, ind)
	}

	return carrier
}

// DeleteElemSliceWithAny 移位法 删除指元素，如果传入的数据不在切片范围内则返回原切片
func DeleteElemSliceWithAny[T any](carrier []T, fnEqual FunEqual[T], elem T) []T {
	var j int
	for _, obj := range carrier {
		if _, ok := fnEqual(obj, elem); !ok {
			carrier[j] = obj
			j++
		}
	}

	return carrier[:j]
}

// DeleteMultiElemSliceWithAny 删除指定数据（数据可为多个）处的数据，如果传入的索引不在切片范围内则返回原切片
func DeleteMultiElemSliceWithAny[T any](carrier []T, fnEqual FunEqual[T], elems ...T) []T {
	for _, elem := range elems {
		carrier = DeleteElemSliceWithAny(carrier, fnEqual, elem)
	}

	return carrier
}

// ChunkSlice 将切片分块，每个子切片的长度不为 chunkSize
func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}
	chunks := make([][]T, 0, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
