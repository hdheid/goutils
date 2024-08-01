package compare

// SkipListFunc 自义定比对参数，需注意格式应与 IntMore 一致
type SkipListFunc[T any] CmpFunc[T]

var (
	IntSkipList   = IntMore
	Int8SkipList  = Int8More
	Int16SkipList = Int16More
	Int32SkipList = Int32More
	Int64SkipList = Int64More

	UintSkipList   = UintMore
	Uint8SkipList  = Uint8More
	Uint16SkipList = Uint16More
	Uint32SkipList = Uint32More
	Uint64SkipList = Uint64More

	Float32SkipList = Float32More
	Float64SkipList = Float64More

	StringSkipList = StringMore
)
