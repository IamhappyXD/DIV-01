package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent != nil {
		if node.Parent.Data < node.Data {
			node.Parent.Right = rplc
		} else {
			node.Parent.Left = rplc
		}
		return root
	} else {
		return rplc
	}

}
