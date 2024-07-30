package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(8)
	q.Push(3)
	q.Pop()

	fmt.Println(q.String())
}
