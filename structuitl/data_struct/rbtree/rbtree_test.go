package rbtree

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"testing"
)

func TestGetColor(t *testing.T) {
	var n *Node[int, int]
	fmt.Println(getColor(n))
}

func TestInsert(t *testing.T) {
	tree := New[int, int](compare.IntRbTree)
	arr := []int{17, 18, 23, 34, 27, 15, 9, 6, 8, 5, 25}

	for _, num := range arr {
		tree.Insert(num, 1)
	}

	/*
		R----15(BLACK)
		     L----8(BLACK)
		     |    L----6(BLACK)
		     |    |    L----5(RED)
		     |    R----9(BLACK)
		     R----18(BLACK)
		          L----17(BLACK)
		          R----27(RED)
		               L----23(BLACK)
		               |    R----25(RED)
		               R----34(BLACK)
	*/
}

func TestDelete(t *testing.T) {
	tree := New[int, int](compare.IntRbTree)
	arr := []int{15, 9, 18, 6, 13, 17, 27, 10, 23, 34, 25, 37}
	for _, num := range arr {
		tree.Insert(num, 1)
	}
	//tree.String()

	node := tree.Find(18)
	tree.Delete(node)
	//tree.String()

	node = tree.Find(25)
	tree.Delete(node)
	//tree.String()

	node = tree.Find(15)
	tree.Delete(node)
	//tree.String()

	node = tree.Find(6)
	tree.Delete(node)
	//tree.String()

	node = tree.Find(13)
	tree.Delete(node)
	node = tree.Find(37)
	tree.Delete(node)
	node = tree.Find(27)
	tree.Delete(node)
	node = tree.Find(17)
	tree.Delete(node)
	node = tree.Find(34)
	tree.Delete(node)
	tree.String()
}

func TestBegin(t *testing.T) {
	tree := New[int, int](compare.IntRbTree)
	tree.Insert(2, 3)
	fmt.Println(tree.Begin())
}
