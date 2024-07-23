package download

import (
	"fmt"
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/fileutil/file_slice/fileMete"
	"net/http"
	"path"
	"strconv"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	url := "https://images-new-cdn.123rf.com.cn/450wm/morozmarusia/morozmarusia2208/morozmarusia220800101.jpg"
	resp, err := http.Head(url)
	if err != nil {
		t.Errorf("")
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		t.Errorf("")
	}

	mete := fileMete.NewData(
		fileMete.WithFileSize(int64(size)),
		fileMete.WithFileName("abb.jpg"),
		fileMete.WithFileSliceSum(int64(size/common.MaxSliceData)),
	)

	download := NewDownloader(mete)
	err = download.SliceDownload(url)
	if err != nil {
		t.Errorf("")
	}
}

func TestDownload(t *testing.T) {
	url := "https://images-new-cdn.123rf.com.cn/450wm/morozmarusia/morozmarusia2208/morozmarusia220800101.jpg"

	mete := fileMete.NewData(
		fileMete.WithFileName(path.Base(url)),
	)

	download := NewDownloader(mete)
	//err := download.DownloadFile(url)
	//if err != nil {
	//	return
	//}

	err := download.SliceDownload(url)
	if err != nil {
		t.Errorf("")
	}
}

func TestReq(t *testing.T) {
	url := "https://images-new-cdn.123rf.com.cn/450wm/morozmarusia/morozmarusia2208/morozmarusia220800101.jpg"
	req := &http.Request{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("错误")
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", 0, 25603))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("错误")
	}

	fmt.Println(resp.Body)
}
