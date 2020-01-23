package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newnode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
	} else {
		l.Tail.Next = newnode
		l.Tail = newnode
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	iter := l
	for iter.Head.Next != nil {
		if iter.Head.Next.Data == data_ref {
			cur := iter.Head.Next
			if cur.Next != nil {
				iter.Head.Next = cur.Next
			} else {
				iter.Head.Next = nil
			}
		}
		iter.Head = iter.Head.Next
	}

}
