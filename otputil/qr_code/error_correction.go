package qr_code

type errorCorrection struct {
	rs *ReedSolomonEncoder
}

var ec = newErrorCorrection()

func newErrorCorrection() *errorCorrection {
	fld := NewGaloisField(285, 256, 0)
	return &errorCorrection{NewReedSolomonEncoder(fld)}
}

func (ec *errorCorrection) calcECC(data []byte, eccCount byte) []byte {
	dataInts := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		dataInts[i] = int(data[i])
	}
	res := ec.rs.Encode(dataInts, int(eccCount))
	result := make([]byte, len(res))
	for i := 0; i < len(res); i++ {
		result[i] = byte(res[i])
	}
	return result
}
