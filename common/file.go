package common

import "time"

// 文件分片数量限制
const (
	MaxGtChannelSum = 500
	RetryChannelSum = 100
)

const (
	MaxSliceData = 1024 // 每个分片大小为 1M
)

const (
	MaxTimeOut = time.Second // 最大超时时间
)
