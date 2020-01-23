package student

func pushback(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	iter1 := n1
	iter2 := n2
	var link *NodeI
	for iter1.Next != nil && iter2.Next != nil {
		if iter1.Data.(int) <= iter2.Data.(int) {
			link = pushback(link, iter1.Data.(int))
			iter1 = iter1.Next
		} else {
			link = pushback(link, iter2.Data.(int))
			iter1 = iter2.Next
		}
	}
	if iter1.Next != nil {
		link.Next = iter1
	} else {
		link.Next = iter2

	}
	return link
}
