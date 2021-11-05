package binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	root:= &TreeNode{
		Left: nil,
		Right: &TreeNode{
			Left: &TreeNode{
				Right: nil,
				Left: nil,
				Val: 3,
			},
			Right: nil,
			Val: 2,
		},
		Val: 1,
	}
	result := root.inorderTraversal(root)
	fmt.Println(result)
}