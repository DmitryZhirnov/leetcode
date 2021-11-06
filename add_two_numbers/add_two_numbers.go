package add_two_numbers

import (
	"alhoritms/common"
)

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	discharge := 0
	list := &common.ListNode{
		Val:  0,
		Next: nil,
	}
	current := list
	for l1 != nil || l2 != nil || discharge != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
		} else {
			val1 = 0
		}
		if l2 != nil {
			val2 = l2.Val
		} else {
			val2 = 0
		}
		sum := val1 + val2 + discharge
		discharge = sum / 10
		val := sum % 10
		current.Next = &common.ListNode{
			Val:  val,
			Next: nil,
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		current = current.Next
	}
	return list.Next
}
