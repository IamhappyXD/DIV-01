package student

type Node struct {
	Data *TreeNode
	Next *Node
}
type List struct {
	Head *Node
	Tail *Node
}

var l List

func pushback(l *List, data *TreeNode) {
	newnode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
	} else {
		l.Tail.Next = newnode
		l.Tail = newnode
	}
}
func size(l *List) int {
	ans := 0
	iter := l.Head
	for iter != nil {
		ans++
		iter = iter.Next
	}
	return ans
}
func deletefront(l *List) {
	l.Head = l.Head.Next
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if size(&l) == 0 {
		pushback(&l, root)
	}
	for size(&l) != 0 {
		x := l.Head.Data
		if x.Right != nil {
			pushback(&l, x.Right)
		}
		if x.Left != nil {
			pushback(&l, x.Left)
		}

		f(x.Data)
		deletefront(&l)
	}

}
