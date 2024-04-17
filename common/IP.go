package common

type IPType int

// IP类型
const (
	TypeIPv4 IPType = iota + 1
	TypeIPv6
)
