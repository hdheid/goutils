package compare

type CmpFunc func(a, b int) int

var IntLess = func(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}
