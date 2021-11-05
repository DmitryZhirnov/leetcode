package common

import (
	"fmt"
	"testing"
)

func TestCommonTreeNodeFill(t *testing.T) {
	vertices := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var root *TreeNode
	root = Fill(vertices, root, 0)
	fmt.Printf("root: %v", root)
	if root.Left.Right.Val != 5 {
		t.Fail()
	}
}
