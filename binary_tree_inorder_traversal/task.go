package binary_tree_inorder_traversal

import "alhoritms/common"

func inorderTraversal(root *common.TreeNode) []int {
	var result []int
	if root != nil {
		if root.Left != nil {
			result = append(result, inorderTraversal(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, inorderTraversal(root.Right)...)
		}
	}
	return result
}
