package bytebufferpool

import (
	"github.com/hdheid/goutils/common"
	"io"
)

type ByteBuffer struct {
	data []byte
}

func (b *ByteBuffer) Len() int {
	return len(b.data)
}

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
	// 函数逻辑与 io.ReadAll 很像
	p := b.data
	nStart := int64(len(p))
	nMax := int64(cap(p))
	n := nStart

	// 如果 p 是一个空的字节数组，则创建一个最小的字节数组
	if nMax == 0 {
		nMax = common.MinSize
		p = make([]byte, nMax)
	} else {
		p = p[:nMax]
	}

	for {
		// 如果当前字节数组已经满了，那么就两倍扩容
		if n == nMax {
			nMax *= 2
			bNew := make([]byte, nMax)
			copy(bNew, p)

			//p = append(p,0)[:nMax] // 利用 append 扩容机制扩容
			//nMax = cap(p)
		}

		nn, err := r.Read(p[n:]) // 读取 r 中的内容
		// todo: 这里p没有等于p[:n]，那么会不会p访问不到n后面的值，导致下一次扩容的时候丢失？这个在测试函数中试一下
		n += int64(nn)
		if err != nil {
			b.data = p[:n]
			n -= nStart
			if err == io.EOF {
				err = nil
			}
			return n, err // 返回读取到的数据长度
		}
	}
}

func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.data)
	return int64(n), err
}

func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *ByteBuffer) WriteByte(c byte) error {
	b.data = append(b.data, c)
	return nil
}

func (b *ByteBuffer) WriteString(s string) (int, error) {
	b.data = append(b.data, s...)
	return len(s), nil
}

func (b *ByteBuffer) Set(p []byte) {
	b.data = append(b.data[:0], p...)
}

func (b *ByteBuffer) SetString(s string) {
	b.data = append(b.data[:0], s...)
}

func (b *ByteBuffer) String() string {
	return string(b.data)
}

func (b *ByteBuffer) Reset() {
	b.data = b.data[:0]
}
