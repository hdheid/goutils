package fileutil

import (
	"log"
	"net/http"
	"os"
)

/*
	文件路径的编写是相对于当前文件的，想进入父级目录，需要加../
*/

// FileExist 通过文件路径判断该文件是否存再
func FileExist(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info == nil {
		return false
	}
	return true
}

// GetFileType 获取文件类型（简单文件），暂不可用
func GetFileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Open file error: ", err)
		return "", err
	}

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal("Read buff error: ", err)
		return "", err
	}

	fileType := http.DetectContentType(buffer)
	return fileType, nil
}
