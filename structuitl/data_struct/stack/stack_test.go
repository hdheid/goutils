package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(7)
	fmt.Println(s.Top())

	s.Pop()
	fmt.Println(s.Top())

	ss := s.String()
	fmt.Println(ss)
}
