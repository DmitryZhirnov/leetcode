package same_tree

import (
	"alhoritms/common"
	"testing"
)

func TestSameTree(t *testing.T) {
	vertices1 := []int{1, 2}
	vertices2 := []int{1, 2}
	var tree1, tree2 *common.TreeNode
	tree1 = common.Fill(vertices1, tree1, 0)
	tree2 = common.Fill(vertices2, tree2, 0)

	result := isSameTree(tree1, tree2)
	if !result {
		t.Fail()
	}

	vertices1 = []int{1, 2}
	vertices2 = []int{1, 0, 2}
	tree1 = common.Fill(vertices1, tree1, 0)
	tree2 = common.Fill(vertices2, tree2, 0)

	result = isSameTree(tree1, tree2)
	if result {
		t.Fail()
	}

	vertices1 = []int{1, 2, 1}
	vertices2 = []int{1, 1, 2}
	tree1 = common.Fill(vertices1, tree1, 0)
	tree2 = common.Fill(vertices2, tree2, 0)

	result = isSameTree(tree1, tree2)
	if result {
		t.Fail()
	}
}
