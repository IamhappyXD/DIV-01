package student

type NodeI struct {
	Data interface{}
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {

	n := &NodeI{Data: data_ref}

	if l == nil {
		return n
	}
	iterator := l
	end := true
	if iterator.Data.(int) >= n.Data.(int) {
		n.Next = iterator
		return n
	}
	for iterator.Next != nil {
		if iterator.Next.Data.(int) >= data_ref {
			end = false
			break
		}
		iterator = iterator.Next
	}
	if end == false {
		n.Next = iterator.Next
	}
	iterator.Next = n
	return l
}
