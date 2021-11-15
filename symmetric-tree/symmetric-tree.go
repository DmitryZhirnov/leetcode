package symmetric_tree

import (
	"alhoritms/common"
	"fmt"
)

func isSymmetric(root *common.TreeNode) bool {
	var stack []int
	var queue []*common.TreeNode
	queue = append(queue, root)
	length := len(queue)
	println()

	for length > 0 {
		current := queue[0]
		queue = queue[1:length]
		if current != nil {
			fmt.Println(current.Val)
			stack = append(stack, current.Val)
			queue = append(queue, current.Left)
			queue = append(queue, current.Right)
		}
		length = len(queue)
	}
	fmt.Println(stack)
	return true
}

func read(node *common.TreeNode) (int, int) {
	return node.Left.Val, node.Right.Val
}
