package main

//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"strconv"
//)
//
//func DownloadFileRange(url, writeFile string) error {
//	f, err := os.OpenFile(writeFile, os.O_CREATE|os.O_TRUNC, 0666)
//	if err != nil {
//		return err
//	}
//
//	defer f.Close()
//
//	resp, err := http.Head(url)
//	if err != nil {
//		return err
//	}
//
//	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
//	if err != nil {
//		return err
//	}
//
//	con := size / 1024 * 1024 * 3
//	var start, end int64
//	for i := 0; i < con; i++ {
//
//		start = int64(i) * int64(size/con)
//		end = start + int64(size/con) - 1
//
//		go func(n int, offset, end int64) {
//			req := &http.Request{}
//			req, err = http.NewRequest(http.MethodGet, url, nil)
//			req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", offset, end))
//
//			client := &http.Client{}
//			resp, err = client.Do(req)
//			if err != nil {
//				return
//			}
//
//			defer resp.Body.Close()
//
//			f.Seek(offset, 0)
//			_, err = io.Copy(f, resp.Body)
//			if err != nil {
//				// log
//			}
//		}(i, start, end)
//	}
//
//	return nil
//}
