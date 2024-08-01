package compare

// PriorityQueueFunc 自义定比对参数，需注意格式应与 IntMore 一致
type PriorityQueueFunc[T any] CmpFunc[T]

// 默认从大到小
var (
	IntPriorityQueue   = IntMore
	Int8PriorityQueue  = Int8More
	Int16PriorityQueue = Int16More
	Int32PriorityQueue = Int32More
	Int64PriorityQueue = Int64More

	UintPriorityQueue   = UintMore
	Uint8PriorityQueue  = Uint8More
	Uint16PriorityQueue = Uint16More
	Uint32PriorityQueue = Uint32More
	Uint64PriorityQueue = Uint64More

	Float32PriorityQueue = Float32More
	Float64PriorityQueue = Float64More

	StringPriorityQueue = StringMore
)
