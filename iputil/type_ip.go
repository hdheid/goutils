package iputil

import (
	"errors"
	"fmt"
	"net"
)

// CheckIPType 判断IP类型
func CheckIPType(ip string) (IPType, error) {
	if net.ParseIP(ip) == nil {
		return 0, errors.New(fmt.Sprintf("IP address: %s , parse error!", ip))
	}

	for _, ch := range ip {
		if ch == '.' {
			return TypeIPv4, nil
		} else if ch == ':' {
			return TypeIPv6, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("IP address: %s , parse error!", ip))
}
