package tree_nodes

import (
	"fmt"
	"testing"
)

func TestTreeNodes(t *testing.T) {
	//node9 := &TreeNode{
	//	Val:   9,
	//	Left:  nil,
	//	Right: nil,
	//}
	//node8 := &TreeNode{
	//	Val:   8,
	//	Left:  nil,
	//	Right: nil,
	//}
	//node7 := &TreeNode{
	//	Val:   7,
	//	Left:  nil,
	//	Right: nil,
	//}
	node6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  node6,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node4,
		Right: node5,
	}
	node1 := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	expected := 6
	result := CountNodes(node1)
fmt.Println(expected, result)
	if expected != result{
		t.Fail()
	}

}
