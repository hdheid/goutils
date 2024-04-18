package fileutil

import (
	"fmt"
	"testing"
)

func TestFileExist(t *testing.T) {
	failPath := "../common/ip.go"

	if !FileExist(failPath) {
		t.Errorf("FileExist(%s) failed", failPath)
	}
}

func TestGetFileType(t *testing.T) {
	failPath := "../common/1.txt"

	fmt.Println(GetFileType(failPath))
}
