package tree_nodes

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count:= 1
	for len(queue) > 0{
		top:=queue[0]
		queue = queue[1:]

		if top.Left != nil{
			queue = append(queue, top.Left)
			count++
		}

		if top.Right != nil{
			queue = append(queue, top.Right)
			count++
		}
	}
	return count

}

