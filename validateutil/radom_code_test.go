package validateutil

import "testing"

func TestGenValidateCode(t *testing.T) {
	l := 5
	t.Log(GenValidateCode(l))
}

// 测试重复值
func TestRepeatCode(t *testing.T) {
	codeHash := make(map[string]bool)
	for i := 0; i < 1000000; i++ {
		code := GenValidateCode(5)

		if _, ok := codeHash[GenValidateCode(6)]; ok {
			t.Errorf("repeat code: %s ", code)
			return
		}
	}
}
