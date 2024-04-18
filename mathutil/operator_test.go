package mathutil

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	fn := IfElse(3 > 4, func() {
		fmt.Println("3确实大于4")
	}, func() {
		fmt.Println("4确实大于3")
	})
	fn()

	str := IfElse(3 > 4, "3确实大于4", "4确实大于3")

	if str != "4确实大于3" {
		t.Error("IfElse函数测试失败")
	}
}
