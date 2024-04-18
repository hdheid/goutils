package common

type IPType int

// IP类型
const (
	TypeIPv4 IPType = iota + 1
	TypeIPv6
)

// 端口范围
const (
	PortMin = 1
	PortMax = 65535
)
