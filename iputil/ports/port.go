package ports

import "github.com/hdheid/goutils/iputil"

// IsValidPort 判断端口是否有效
func IsValidPort(port int) bool {
	if port < iputil.PortMin || port > iputil.PortMax {
		return false
	}
	return true
}
