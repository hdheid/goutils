package main

import (
	"fmt"
	"github.com/hdheid/goutils/fileutil/file_slice/download"
	"github.com/hdheid/goutils/fileutil/file_slice/fileMete"
	"path"
)

func main() {
	url := "https://images-new-cdn.123rf.com.cn/450wm/morozmarusia/morozmarusia2208/morozmarusia220800101.jpg"

	mete := fileMete.NewData(
		fileMete.WithFileName(path.Base(url)),
	)

	down := download.NewDownloader(mete,
		download.WithDownloadDir("fileutil/file_slice/download/temp"),
		download.WithFileUrl(url),
	)

	err := down.SliceDownload(url)
	if err != nil {
		fmt.Println("错误")
	}
}
