package fileMete

import "time"

type Data struct {
	FID          string    // 文件编号
	FileName     string    // 文件名称
	FileSize     int64     // 文件大小
	FileSliceSum int64     // 文件分片数量
	ModifyTime   time.Time // 文件修改时间
}

type OpFunc func(data *Data)

func WithFID(FID string) OpFunc {
	return func(d *Data) {
		d.FID = FID
	}
}

func WithFileName(FileName string) OpFunc {
	return func(d *Data) {
		d.FileName = FileName
	}
}

func WithFileSize(FileSize int64) OpFunc {
	return func(d *Data) {
		d.FileSize = FileSize
	}
}

func WithFileSliceSum(FileSliceSum int64) OpFunc {
	return func(d *Data) {
		d.FileSliceSum = FileSliceSum
	}
}

func WithModifyTime(ModifyTime time.Time) OpFunc {
	return func(d *Data) {
		d.ModifyTime = ModifyTime
	}
}

func NewData(ops ...OpFunc) *Data {
	data := &Data{
		FID:        "",
		ModifyTime: time.Now(),
	}

	for _, op := range ops {
		op(data)
	}

	return data
}
