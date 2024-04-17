package iputil

import (
	"fmt"
	"testing"
)

func TestGenerateRandomIPs(t *testing.T) {
	num := 10000
	ips := GenerateRandomIPs(num)

	fmt.Println(ips)
}

// todo:增加导出成文件的功能，如果需要的 IP 数据量比较大
