package cdir

import (
	"errors"
	"fmt"
	"github.com/hdheid/goutils/iputil"
	"net"
	"strconv"
	"strings"
)

// GetIPs 给出一个 CIDR 类型的 IP 地址，返回一个字符串切片，包含该 CIDR 所有可用 IP 地址信息
func GetIPs(CDIR string) ([]string, error) {
	// 检查该字符串是否是IP
	if ip := net.ParseIP(CDIR); ip == nil {
		return nil, errors.New(fmt.Sprintf("IP address: %s , parse error!", CDIR))
	}

	ip, ipnet, err := net.ParseCIDR(CDIR)
	if err != nil {
		return nil, err
	}

	ipAddr := strings.Split(ipnet.String(), "/")
	maskCount, err := strconv.Atoi(ipAddr[1]) // 获取子网掩码位数
	if err != nil {
		return nil, errors.New(fmt.Sprintf("IP address: %s , Subnet mask error!", CDIR))
	}

	t, _ := iputil.CheckIPType(ip.String()) // 获取IP类型，网段不能够太大
	if t == iputil.TypeIPv4 && maskCount < 19 {
		return nil, errors.New(fmt.Sprintf("The network segment is too large: %d", maskCount))
	}
	if t == iputil.TypeIPv6 && maskCount < 115 {
		return nil, errors.New(fmt.Sprintf("The network segment is too large: %d", maskCount))
	}

	// 获取所有可用IP
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
