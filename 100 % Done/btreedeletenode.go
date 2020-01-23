package student

func findmin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	} else {
		return findmin(root.Left)
	}
}
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Right == nil && root.Left == nil {
			if root.Parent.Left == root {
				root.Parent.Left = nil
			} else {
				root.Parent.Right = nil
			}
			return nil
		} else if root.Right == nil {
			root.Left.Parent = root.Parent
			root = root.Left
		} else if root.Left == nil {
			root.Right.Parent = root.Parent
			root = root.Right
		} else {
			temp := findmin(root.Right)
			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root

}
