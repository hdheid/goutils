package compare

// MapFunc 自义定比对参数，需注意格式应与 IntMore 一致
type MapFunc[T any] CmpFunc[T]

var (
	IntMap   = IntLess
	Int8Map  = Int8Less
	Int16Map = Int16Less
	Int32Map = Int32Less
	Int64Map = Int64Less

	UintMap   = UintLess
	Uint8Map  = Uint8Less
	Uint16Map = Uint16Less
	Uint32Map = Uint32Less
	Uint64Map = Uint64Less

	Float32Map = Float32Less
	Float64Map = Float64Less

	StringMap = StringLess
)
