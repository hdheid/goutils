package rbtree

type Color bool

const (
	Red   Color = false
	Black Color = true
)

type Node[K, V any] struct {
	parent *Node[K, V]
	left   *Node[K, V]
	right  *Node[K, V]
	color  Color
	key    K
	val    V
}

func (n *Node[K, V]) Key() K { return n.key }

func (n *Node[K, V]) Val() V { return n.val }

func (n *Node[K, V]) Prev() *Node[K, V] { return preSuccessor(n) }

func (n *Node[K, V]) Next() *Node[K, V] { return successor(n) }

func (n *Node[K, V]) IsRight() bool { return n == n.parent.right }

func (n *Node[K, V]) IsLeft() bool { return n == n.parent.left }

func getColor[K, V any](n *Node[K, V]) Color {
	if n == nil {
		return Black
	}
	return n.color
}
func preSuccessor[K, V any](n *Node[K, V]) *Node[K, V] {
	if n.right != nil {
		return maxNode(n)
	}
	y := n.parent

	//if y != nil && n == y.right {
	//	return y
	//}

	for y != nil && n == y.left {
		n = y
		y = n.parent
	}

	return y
}

func successor[K, V any](n *Node[K, V]) *Node[K, V] {
	// 如果该节点存在右子树，那么后继节点就是右子树的最小值，也就是右子树最左边的节点
	if n.right != nil {
		return minNode(n.right)
	}

	// 如果该节点不存在右子树，那么就一直往左上找，直到找到第一个父节点在右边，该父节点就是后继节点
	// 因为往左上找，能保证每一个节点都比该节点小，直到出现第一个父节点在右边，那么该父节点比这些所有节点都大一点，因此就是后继节点
	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = n.parent
	}

	return y
}

func minNode[K, V any](n *Node[K, V]) *Node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maxNode[K, V any](n *Node[K, V]) *Node[K, V] {
	for n.right != nil {
		n = n.right
	}

	return n
}
