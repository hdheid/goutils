package validateutil

import (
	"fmt"
	"regexp"
)

// ValidateString 参数校验，只能包含数字、大小写字母、以及部分特殊字符：_-.()《》<>[]【】
func ValidateString(str string) error {
	match, err := regexp.MatchString("^[\u4E00-\u9FA5A-Za-z0-9_.\\-()《》<>\\[\\]【】]+$", str)
	if err != nil {
		return fmt.Errorf("Please try again ")
	}

	if !match {
		return fmt.Errorf("Can only contain: numbers, letters, _-.()《》<>[]【】 ")
	}

	return nil
}
