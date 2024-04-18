package iputil

import (
	"github.com/hdheid/goutils/common"
	"testing"
)

func TestCheckIPType(t *testing.T) {
	IP := []string{"127.0.0.1", "192.168.1.1", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:db8:85a3:0:0:8a2e:370:7334", "2001:db8:85a3::8a2e:370:7334"}
	TypeIP := []common.IPType{common.TypeIPv4, common.TypeIPv4, common.TypeIPv6, common.TypeIPv6, common.TypeIPv6}

	for i, ip := range IP {
		ty, err := CheckIPType(ip)
		if err != nil {
			t.Logf("parse error!")
		}
		if ty != TypeIP[i] {
			t.Errorf("CheckIPType(%s) = %d, want %d", ip, ty, TypeIP)
		}
	}
}
