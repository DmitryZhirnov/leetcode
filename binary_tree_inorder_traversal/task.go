package binary_tree_inorder_traversal

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) inorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		if root.Left != nil {
			result = append(result, t.inorderTraversal(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, t.inorderTraversal(root.Right)...)
		}
	}
	return result
}
