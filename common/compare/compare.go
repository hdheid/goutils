package compare

type CmpFunc[T any] func(a, b T) int

var IntLess = func(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var IntMore = func(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Int8Less = func(a, b int8) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Int8More = func(a, b int8) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Int16Less = func(a, b int16) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Int16More = func(a, b int16) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Int32Less = func(a, b int32) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Int32More = func(a, b int32) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Int64Less = func(a, b int64) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Int64More = func(a, b int64) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var UintLess = func(a, b uint) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var UintMore = func(a, b uint) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Uint8Less = func(a, b uint8) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Uint8More = func(a, b uint8) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Uint16Less = func(a, b uint16) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Uint16More = func(a, b uint16) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Uint32Less = func(a, b uint32) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Uint32More = func(a, b uint32) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Uint64Less = func(a, b uint64) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Uint64More = func(a, b uint64) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var StringLess = func(a, b string) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var StringMore = func(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Float32Less = func(a, b float32) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Float32More = func(a, b float32) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var Float64Less = func(a, b float64) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

var Float64More = func(a, b float64) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
