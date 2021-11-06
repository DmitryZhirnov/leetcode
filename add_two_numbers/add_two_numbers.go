package add_two_numbers

import (
	"alhoritms/common"
)

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	return addTwoNumbersRecursive(l1, l2, 0)
}

func addTwoNumbersRecursive(l1 *common.ListNode, l2 *common.ListNode, discharge int) *common.ListNode {
	if l1 == nil && l2 == nil && discharge == 0 {
		return nil
	}
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
	if l1 != nil {
		l1 = l1.Next
	}
	if l2 != nil {
		l2 = l2.Next
	}
	sum := val1 + val2 + discharge
	val := sum % 10
	discharge = sum / 10
	return &common.ListNode{
		Val:  val,
		Next: addTwoNumbersRecursive(l1, l2, discharge),
	}
}
