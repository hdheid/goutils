package iputil

import (
	"fmt"
	"math/rand"
)

// GenerateRandomIPs 生成随机IP地址，用于测试，不验证IP的准确性
func GenerateRandomIPs(numIPs int) string {
	var ips string

	for i := 0; i < numIPs; i++ {
		part1 := rand.Intn(256)
		part2 := rand.Intn(256)
		part3 := rand.Intn(256)
		part4 := rand.Intn(256)

		ip := fmt.Sprintf("%d.%d.%d.%d", part1, part2, part3, part4)
		ips += ","
		ips += ip
	}

	return ips[1:]
}
