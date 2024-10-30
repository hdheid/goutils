package common

const (
	MinBitSize = 6  // 因为 CPU L1 缓存大小为 64K
	Steps      = 20 // 一共分出来了20个区间，依次为：0~2<<6-1、2<<6~2<<7-1......2<<25~∞

	MinSize = 1 << MinBitSize
	MaxSize = 1 << (MinBitSize + Steps - 1)

	CalibrateCallsThreshold = 42000 // 42000
	MaxPercentile           = 0.95  // 0.95
)
