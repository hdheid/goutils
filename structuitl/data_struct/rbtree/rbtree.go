package rbtree

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	"github.com/hdheid/goutils/mathutil"
)

type RbTree[K, V any] struct {
	root   *Node[K, V]
	size   int
	keyCmp compare.CmpFunc[K]
}

// todo：后续可以加上一个pair数组，来进行初始化

func New[K, V any](cmp compare.CmpFunc[K]) *RbTree[K, V] {
	return &RbTree[K, V]{
		keyCmp: cmp,
	}
}

func (t *RbTree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

// Begin 返回第一个最小节点
func (t *RbTree[K, V]) Begin() *Node[K, V] {
	if t.root == nil {
		return nil
	}
	return minNode(t.root)
}

func (t *RbTree[K, V]) Empty() bool { return t.size == 0 }

func (t *RbTree[K, V]) Size() int { return t.size }

func (t *RbTree[K, V]) Insert(key K, val V) {
	x := t.root
	var y *Node[K, V]

	for x != nil {
		y = x
		if t.keyCmp(key, x.Key()) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node[K, V]{parent: y, color: Red, key: key, val: val}
	t.size++

	if y == nil {
		z.color = Black
		t.root = z
		return
	}

	// todo：这里可否优化到上面的for循环中
	if t.keyCmp(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}

	t.rbInsertFixup(z)
}

func (t *RbTree[K, V]) Find(key K) *Node[K, V] {
	x := t.root
	for x != nil {
		if t.keyCmp(key, x.key) < 0 {
			x = x.left
		} else if t.keyCmp(key, x.key) > 0 {
			x = x.right
		} else {
			return x
		}
	}
	return nil
}

func (t *RbTree[K, V]) Delete(node *Node[K, V]) {
	/*
		删除需要修复会出现的三种情况：
			1.没有孩子
				1.删除的是红节点：直接删除即可
				2.删除的是黑节点
					1.兄弟节点为黑色
						1.兄弟至少有一个红孩子（如果双红，则寻找LL或者RR）
							1.LL（兄（s）为左子节点，兄红子（r）为左子节点）：r变为s的颜色，s变为父的颜色，父变为黑色；父右旋
							2.RR（兄（s）为右子节点，兄红子（r）为右子节点）：r变为s的颜色，s变为父的颜色，父变为黑色；父左旋
							3.LR（兄（s）为左子节点，兄红子（r）为右子节点）：r变为父的颜色，s不变，父变黑；左旋兄，再右旋父
							4.RL（兄（s）为右子节点，兄红子（r）为左子节点）：r变为父的颜色，s不变，父变黑；右旋兄，再左旋父
						2.兄弟的孩子全黑
							1.父节点为黑色
								兄变红，从父结点开始继续调整
							2.父节点为红色
								将兄节点变为红色，将父节点变为黑色即可
					2.兄弟节点为红色，兄父变色，并朝删除节点防线旋转
						继续判断删除节点的兄节点为红色还是黑色
			2.只有左子树/只有右子树（只会存在黑节点有一个左/右子节点的情况）
				1.此时子需要将该节点删除，然后使子节点代替之并将子节点颜色变黑即可
			3.左右子树都有，可以使用直接前驱或者后继代替，转换成前两种情况
	*/
	z := node
	if z == nil {
		return
	}

	var x, y *Node[K, V]

	// 如果需要删除的节点为叶子节点或者只存在一个孩子，那么就可以直接删除，否则需要找到直接前驱来替换之
	if z.left != nil || z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	// 令x为y的孩子，那么x要么是左/右子节点，要么是空
	//if y.left!=nil{
	//	x = y.left
	//}else {
	//	x = y.right
	//}
	x = mathutil.IfElse(y.left != nil, y.left, y.right)

	// 用子节点 X 代删除掉的节点 Y
	xParent := y.parent
	if x != nil {
		x.parent = y.parent
	}
	if y.parent == nil {
		t.root = x
	} else if y.IsLeft() {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.val = y.val
	}

	if y.color {
		t.rbDeleteFixUp(x, xParent)
	}

	t.size--
}

func (t *RbTree[K, V]) rbInsertFixup(z *Node[K, V]) {
	/*
		1. 插入节点为根节点，直接将根节点变黑
		2. 插入节点的叔叔节点为红色。叔叔、父亲变黑，爷爷节点变红，然后从爷爷开始更新
		3. 插入节点的叔叔是黑色：
			一、LL：父节点是祖父节点的左子节点，当前节点是父节点的左子节点。右旋祖父节点，并变色父节点和祖父节点
			二、LR：类似于上面，左右。先左旋祖父节点左子节点，然后右旋祖父节点，并变色父节点和祖父节点
			三、RR：右右。同上
			四、RL：右左。同上
	*/

	// 情况2与情况3
	var uncle *Node[K, V]
	for z.parent != nil && z.parent.color == Red {
		// 找到节点的叔叔节点，可以保证，至少存在根节点为祖父节点的时候才会进入循环，因此不存在祖父节点为空的情况
		if z.parent == z.parent.parent.left {
			uncle = z.parent.parent.right
		} else {
			uncle = z.parent.parent.left
		}

		// 情况2
		//if uncle != nil && uncle.color == Red {
		if getColor(uncle) == Red {
			z.parent.color = Black
			uncle.color = Black
			z.parent.parent.color = Red
			z = z.parent.parent
			continue
		}

		// 情况3
		if z.parent.IsLeft() {
			if z.IsRight() { //LR
				z = z.parent // 同下
				t.leftRotate(z)
			}

			//LL
			z.parent.color = Black
			z.parent.parent.color = Red
			t.rightRotate(z.parent.parent)
		} else {
			if z.IsLeft() { //RL（右旋后格式与RR一致）
				z = z.parent // 由于右旋后，z和z.parent调换了身份，因此为了保证与RR类型一致，z需要为在z.parent
				t.rightRotate(z)
			}

			//RR
			z.parent.color = Black
			z.parent.parent.color = Red
			t.leftRotate(z.parent.parent)
		}
	}

	// 情况 1
	t.root.color = Black
}

func (t *RbTree[K, V]) rbDeleteFixUp(x, parent *Node[K, V]) {
	var w *Node[K, V]
	for x != t.root && getColor(x) { // 当x节点没有达到根节点或者x节点一直为黑节点的时候，就不断的调整
		if x != nil {
			parent = x.parent
		}

		if x == parent.left {
			x, w = t.rbFixUpLeft(x, parent, w)
		} else {
			x, w = t.rbFixUpRight(x, parent, w)
		}
	}
	if x != nil {
		x.color = Black
	}
}

func (t *RbTree[K, V]) rbFixUpLeft(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.right // w为兄弟节点
	if !w.color {
		w.color = Black
		parent.color = Red
		t.leftRotate(parent)
		w = parent.right
		// 左旋后，兄弟节点上去了，需要重新更新兄弟节点，此时兄弟节点一定是黑色的。此时树就变成了兄弟节点为黑色的情况来进行调整
	}

	// 如果兄弟节点的孩子都是黑色的，将兄节点变红，更新当前节点为父节点即可（如果父节点为红色，会跳出循环将父节点染黑，因此该情况不用管）
	if getColor(w.left) && getColor(w.right) {
		w.color = Red
		x = parent
	} else {
		// 如果是右右（兄弟节点）
		if !getColor(w.right) {
			w.right.color = w.color
			w.color = parent.color
			parent.color = Black
			t.leftRotate(parent)
			x = t.root //退出循环条件
			return x, w
		}

		// 此时兄弟节点的左子节点一定为红色节点
		w.left.color = parent.color
		parent.color = Black
		t.rightRotate(w)
		t.leftRotate(parent)
		x = t.root //退出循环条件
	}
	return x, w
}

func (t *RbTree[K, V]) rbFixUpRight(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.left // w为兄弟节点
	if !w.color {
		w.color = Black
		parent.color = Red
		t.rightRotate(parent)
		w = parent.left
		// 右旋后，兄弟节点上去了，需要重新更新兄弟节点，此时兄弟节点一定是黑色的。此时树就变成了兄弟节点为黑色的情况来进行调整
	}

	// 如果兄弟节点的孩子都是黑色的，将兄节点变红，更新当前节点为父节点即可（如果父节点为红色，会跳出循环将父节点染黑，因此该情况不用管）
	if getColor(w.left) && getColor(w.right) {
		w.color = Red
		x = parent
	} else {
		// 如果是左左（兄弟节点）
		if !getColor(w.left) {
			w.left.color = w.color
			w.color = parent.color
			parent.color = Black
			t.rightRotate(parent)
			x = t.root //退出循环条件
			return x, w
		}

		// 此时兄弟节点的右子节点一定为红色节点
		w.right.color = parent.color
		parent.color = Black
		t.leftRotate(w)
		t.rightRotate(parent)
		x = t.root //退出循环条件
	}
	return x, w
}

// 左旋
func (t *RbTree[K, V]) leftRotate(x *Node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

// 右旋
func (t *RbTree[K, V]) rightRotate(x *Node[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (t *RbTree[K, V]) String() {
	if t.root != nil {
		t.printf(t.root, "", true)
	}
}

func (t *RbTree[K, V]) printf(node *Node[K, V], indent string, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "     "
		} else {
			fmt.Print("L----")
			indent += "|    "
		}
		color := "RED"
		if node.color == Black {
			color = "BLACK"
		}
		fmt.Printf("%v(%v)\n", node.key, color)
		t.printf(node.left, indent, false)
		t.printf(node.right, indent, true)
	}
}
