package ports

import (
	"github.com/hdheid/goutils/common"
)

// IsValidPort 判断端口是否有效
func IsValidPort(port int) bool {
	if port < common.PortMin || port > common.PortMax {
		return false
	}
	return true
}
