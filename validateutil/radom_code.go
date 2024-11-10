package validateutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenValidateCode 生成长度为 length 的随机验证码
func GenValidateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rl := len(numeric)

	rand.Seed(time.Now().UnixNano())
	var code strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&code, "%d", numeric[rand.Intn(rl)])
	}

	return code.String()
}

/*
GenValidateString 生成长度为 length 的随机字符串，对于 radonSize：
0. 为 1 表示随机大小写加数字字符串，
1. 为 2 表示随机小写字符串，
2. 为 3 表示随机大写字符串，
3. 为 4 表示随机小写加数字字符串
4. 为 5 表示随机大写加数字字符串
5. 为 5 表示随机一个 Base32 编码字符串
*/
func GenValidateString(length int, radonSize int) string {
	alphameric := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	numeric := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	switch radonSize {
	case 0: // 大小写和数字混合
		alphameric = append(alphameric, numeric...)
	case 1: // 随机小写
		alphameric = alphameric[26:]
	case 2: // 随机大写
		alphameric = alphameric[:26]
	case 3: // 随机小写加数字
		alphameric = alphameric[26:]
		alphameric = append(alphameric, numeric...)
	case 4: // 随机大写加数字
		alphameric = alphameric[:26]
		alphameric = append(alphameric, numeric...)
	case 5:
		alphameric = alphameric[:26]
		alphameric = append(alphameric, numeric[2:8]...)
	}

	rl := len(alphameric)

	rand.Seed(time.Now().UnixNano())
	var code strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&code, "%c", alphameric[rand.Intn(rl)])
	}

	codeStr := code.String()

	// base32
	if radonSize == 5 { // todo: 后面不能使用魔数
		if n := len(codeStr) % 8; n != 0 { // 补充=保证为8的倍数
			codeStr = codeStr + strings.Repeat("=", 8-n)
		}
	}

	return codeStr
}
