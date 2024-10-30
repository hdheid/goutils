package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func readFileAndSend(filename string) error {
	f, err := os.Open("fileutil/file_slice/move/" + filename)
	if err != nil {
		log.Println("文件打开失败，", err, ", filename: ", filename)
		return err
	}
	defer f.Close()
	//fs, err := f.Stat()
	//if err != nil {
	//	log.Println("文件信息获取错误，", err, ", filename: ", filename)
	//	return err
	//}
	br := bufio.NewReader(f)

	cacheFile := JoinPath("fileutil/file_slice/move/tmp", filename)
	file, err := os.OpenFile(cacheFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Println("目标文件创建错误，", err, ", filename: ", filename)
		return err
	}
	wr := bufio.NewWriterSize(file, 10<<20) // 会初始化一个10MB的buffer
	defer file.Close()
	defer wr.Flush()

	for {
		// TODO: @zcf 如何优雅的读取大文件?
		line := make([]byte, 1024)
		n, err := br.Read(line)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println("read file error,", err, ", filename:", filename)
			break
		}
		if n == 0 {
			break
		}

		err = Write(line, wr)
		if err != nil {
			log.Println("write file error,", err, ", filename:", filename)
			break
		}
	}

	return nil
}

func Write(data []byte, w *bufio.Writer) error {
	_, err := w.Write(data)
	return err
}

func JoinPath(elem ...string) string {
	return path.Join(elem...)
}

var bufLen = 2 * 1024 * 1024

func DownLoadFileShardByFilePath(filename string) (err error) {
	rf, err := os.Open("fileutil/file_slice/move/" + filename)
	if err != nil {
		log.Println("文件打开失败，", err, ", filename: ", filename)
		return err
	}
	defer rf.Close()
	//fs, err := f.Stat()
	//if err != nil {
	//	log.Println("文件信息获取错误，", err, ", filename: ", filename)
	//	return err
	//}
	br := bufio.NewReader(rf)

	cacheFile := JoinPath("fileutil/file_slice/move/tmp", filename)
	file, err := os.OpenFile(cacheFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Println("目标文件创建错误，", err, ", filename: ", filename)
		return err
	}
	wr := bufio.NewWriterSize(file, 10<<20) // 会初始化一个10MB的buffer
	defer file.Close()

	bs := make([]byte, bufLen)
	for {
		var read int
		read, err = br.Read(bs)
		if err != nil || 0 == read {
			break
		}

		err = Write(bs[:read], wr)
		if err != nil {
			fmt.Println("write err:" + err.Error())
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		return
	}

	if err = wr.Flush(); err != nil {
		fmt.Println("writer flush err: ", err.Error())
		return
	}

	fmt.Printf("downLoad over")
	return
}

func main() {
	t := time.Now()
	err := readFileAndSend("test.mp4")
	if err != nil {
		return
	}
	end := time.Since(t)
	fmt.Println(end)
}
