package bitmap

import "fmt"

type Bitmap struct {
	data []byte
	size uint64
}

func New(size uint64) *Bitmap {
	if size&7 != 0 {
		size = (size + 8) &^ 7
	} // 保证大小为8的倍数

	bitmap := &Bitmap{
		size: size,
		data: make([]byte, size>>3, size>>3),
	}

	return bitmap
}

func (b *Bitmap) Set(num uint64) bool {
	for num >= b.size { // todo 后续增加扩容机制
		size := b.Expend(b.size) // 扩容
		b.Resize(size)
	}

	b.data[num>>3] |= 1 << (num & 0x07)
	return true
}

func (b *Bitmap) Unset(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] &= ^(1 << (pos & 0x07))
	return true
}

func (b *Bitmap) IsSet(num uint64) bool {
	if num >= b.size {
		return false
	}
	if b.data[num>>3]&(1<<(num&0x07)) >= 1 {
		return true
	}

	return false
}

func (b *Bitmap) Resize(size uint64) {
	if size&7 != 0 {
		size = (size + 8) &^ 7
	} // 保证大小为8的倍数

	if b.size == size {
		return
	}
	data := make([]byte, size>>3, size>>3)
	copy(data, b.data)
	b.data = data
	b.size = size
}

func (b *Bitmap) Size() uint64 {
	return b.size
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size>>3, b.size>>3)
}

func (b *Bitmap) Data() []byte {
	return b.data
}

func (b *Bitmap) SizeExp() (exponent uint64) {
	size := b.Size()
	cnt := uint64(1)

	for cnt < size {
		cnt <<= 1
		exponent++
	}
	return exponent
}

func (b *Bitmap) Expend(size uint64) uint64 {
	// todo 目前是翻倍扩容
	size = size << 1
	return size
}

// todo :string，输出bitmap
func (b *Bitmap) String() {
	fmt.Printf("The bitmap: ")
	for i := int(b.size/8 - 1); i >= 0; i-- {
		fmt.Printf("%d  %b ", b.data[i], b.data[i])
	}
}
