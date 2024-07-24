package bitmap

import (
	"sync"
	"testing"
)

func TestBitmap(t *testing.T) {
	var length uint64 = 1e5
	b := New(length)

	wg := sync.WaitGroup{}
	wg.Add(int(length))

	for i := 0; i < int(length); i++ {
		i := i
		go func() {
			defer wg.Done()
			if i&1 == 1 {
				b.Set(uint64(i))
			}
		}()
	}

	//b.String()

	wg.Wait()

	for i := 0; i < int(length); i++ {
		if i&1 == 1 {
			if b.IsSet(uint64(i)) != true {
				t.Errorf("err")
			}
		}
	}

	t.Log(b.Size())
}
