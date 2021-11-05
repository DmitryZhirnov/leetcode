package binary_tree_inorder_traversal

import (
	"alhoritms/common"
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	root := &common.TreeNode{
		Left: nil,
		Right: &common.TreeNode{
			Left: &common.TreeNode{
				Right: nil,
				Left:  nil,
				Val:   3,
			},
			Right: nil,
			Val:   2,
		},
		Val: 1,
	}
	result := inorderTraversal(root)
	fmt.Println(result)
}
