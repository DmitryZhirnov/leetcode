package symmetric_tree

import (
	"alhoritms/common"
	"testing"
)

func TestSymmetricTree(t *testing.T) {
	vertices := []int{1, 2, 2, 3, 0, 0, 3}
	var root *common.TreeNode
	root = common.Fill(vertices, root, 0)
	if root.Left.Left.Val != 3 {
		t.Fail()
	}
	//fmt.Printf("root: %v", root)
	isSymmetric(root)
}
