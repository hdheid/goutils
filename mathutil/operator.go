package mathutil

// IfElse 实现 ? : 语句
func IfElse[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}
