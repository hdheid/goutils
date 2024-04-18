package common

type BooleanData interface {
	~bool
}

type StringData interface {
	~string
}

type UintData interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type FloatData interface {
	~float32 | ~float64
}

type IntData interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type IntegerData interface {
	UintData | IntData
}

type NumberData interface {
	IntegerData | FloatData
}

type BasicData interface {
	BooleanData | StringData | IntegerData | FloatData
}

// IfElse ? : 语法
type IfElse interface {
	~bool | any
}
