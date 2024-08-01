package compare

// RbTreeFunc 自义定比对参数，需注意格式应与 IntMore 一致
type RbTreeFunc[T any] CmpFunc[T]

var (
	IntRbTree   = IntLess
	Int8RbTree  = Int8Less
	Int16RbTree = Int16Less
	Int32RbTree = Int32Less
	Int64RbTree = Int64Less

	UintRbTree   = UintLess
	Uint8RbTree  = Uint8Less
	Uint16RbTree = Uint16Less
	Uint32RbTree = Uint32Less
	Uint64RbTree = Uint64Less

	Float32RbTree = Float32Less
	Float64RbTree = Float64Less

	StringRbTree = StringLess
)
