package download

import (
	"fmt"
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/fileutil/file_slice/fileMete"
	"github.com/hdheid/goutils/mathutil"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Mover 文件移动结构体
type Mover struct {
	FileMete     fileMete.Data  // 文件元数据
	MoveDir      string         // 目标目录
	CurrentDir   string         // 文件所在目录
	RetryChannel chan int       // 重传channel通道
	MaxGtChannel chan struct{}  // 限制上传的goroutine的数量通道
	StartTime    time.Time      // 移动开始时间
	EndTime      time.Time      // 移动结束时间
	RetrySlice   []int          // 需要重传分片编号
	wait         sync.WaitGroup // 需要等待所有分片下载完毕
	mutex        sync.Mutex     // 加锁
}

//type FileSlice struct {
//	FID   string
//	Index int64
//	data  []byte
//}

type OpFunc func(*Mover)

func WithMovePath(MoveDir string) OpFunc {
	return func(d *Mover) {
		d.MoveDir = MoveDir
	}
}

func WithRetryChannelSum(sum int) OpFunc {
	return func(d *Mover) {
		d.RetryChannel = make(chan int, sum)
	}
}

func WithMaxGtChannelSum(sum int) OpFunc {
	return func(d *Mover) {
		d.MaxGtChannel = make(chan struct{}, sum)
	}
}

func WithCurrentDir(currentDir string) OpFunc {
	return func(d *Mover) {
		d.CurrentDir = currentDir
	}
}

func NewMover(fileMete *fileMete.Data, ops ...OpFunc) *Mover {
	mv := &Mover{
		FileMete:     *fileMete,
		MoveDir:      "./temp", // 默认当前目录下的 temp 文件夹下
		MaxGtChannel: make(chan struct{}, common.MaxGtChannelSum),
		RetryChannel: make(chan int, common.RetryChannelSum),
	}

	for _, op := range ops {
		op(mv)
	}

	return mv
}

func (m *Mover) GetMovePath() string {
	return m.MoveDir + "/" + m.FileMete.FileName
}

func (m *Mover) SliceDownload(url string) error {
	f, err := os.OpenFile(m.GetMovePath(), os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	// 1. 判断文件是否被修改过
	// 2. 判断文件是否已经下载过一部分

	defer f.Close()

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length")) // 获取文件大小
	if err != nil {
		return err
	}

	sliceSum := size/common.MaxSliceData + 1
	m.FileMete.FileSliceSum = int64(sliceSum) // 分片数量
	m.FileMete.FileSize = int64(size)

	var start, end int64
	for i := 0; i < sliceSum; i++ {
		start = int64(i) * common.MaxSliceData
		end = start + int64(mathutil.Min(size-int(start), common.MaxSliceData)) - 1

		m.wait.Add(1)
		go m.GetSlice(i, start, end, f)
	}

	m.wait.Wait()

	return nil
}

// DownloadFile 从指定URL下载文件至本地路径
func (m *Mover) DownloadFile(url string) error {
	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("下载文件时发生错误: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载文件失败，状态码: %d", resp.StatusCode)
	}

	// 创建或打开文件准备写入
	f, err := os.OpenFile(m.GetMovePath(), os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("创建文件时发生错误: %w", err)
	}
	defer f.Close()

	// 将响应体的内容复制到文件
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("写入文件时发生错误: %w", err)
	}

	fmt.Println("文件下载成功")
	return nil
}

// GetSlice 该函数是直接将文件写入对应偏移量的。
func (m *Mover) GetSlice(i int, offset, end int64, f *os.File) {
	defer m.wait.Done()

	req := &http.Request{}
	req, err := http.NewRequest(http.MethodGet, m.GetMovePath(), nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", offset, end))
	client := &http.Client{Timeout: common.MaxTimeOut}
	resp, err := client.Do(req)
	if err != nil {
		m.RetryChannel <- i // 记录需要重复下载的编号
	}
	defer resp.Body.Close()

	func(offset int64, f *os.File) {
		m.mutex.Lock()
		defer m.mutex.Unlock() // 加锁，对齐颗粒度

		f.Seek(offset, 0)
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			fmt.Println("分片打印错误，需要重新打印")
		}
	}(offset, f)

	fmt.Println("分片打印完毕，打印范围为：", i, ":: ", offset, "----", end)
}

//func (d *Mover) RetryDownload() {
//	for {
//		select {
//		case index := <-d.RetryChannel:
//			d.GetSlice()
//		}
//	}
//}
