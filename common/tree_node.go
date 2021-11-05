package common

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Fill Заполнение двоичного дерева из массива
func Fill(vertices []int, root *TreeNode, level int) *TreeNode {
	if len(vertices) > level {
		var left, right *TreeNode
		if root != nil {
			left = root.Left
			right = root.Right
		}
		if vertices[level] == 0 {
			return nil
		}
		current := &TreeNode{
			Left:  Fill(vertices, left, 2*level+1),
			Right: Fill(vertices, right, 2*level+2),
			Val:   vertices[level],
		}
		root = current
	}

	return root
}
