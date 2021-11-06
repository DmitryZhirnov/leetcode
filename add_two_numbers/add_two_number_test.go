package add_two_numbers

import (
	"alhoritms/common"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var linkedList1, linkedList2, resultLinkedList *common.ListNode
	linkedList1 = common.FillLinkedList([]int{2, 4, 3}, 0)
	linkedList2 = common.FillLinkedList([]int{5, 6, 4}, 0)
	resultLinkedList = addTwoNumbers(linkedList1, linkedList2)
	if resultLinkedList.Val != 7 && resultLinkedList.Next.Val != 0 && resultLinkedList.Next.Next.Val != 8 {
		t.Fail()
	}

	linkedList1 = common.FillLinkedList([]int{0}, 0)
	linkedList2 = common.FillLinkedList([]int{0}, 0)
	resultLinkedList = addTwoNumbers(linkedList1, linkedList2)
	if resultLinkedList.Val != 0 {
		t.Fail()
	}

	linkedList1 = common.FillLinkedList([]int{9, 9, 9, 9, 9, 9, 9}, 0)
	linkedList2 = common.FillLinkedList([]int{9, 9, 9, 9}, 0)
	resultLinkedList = addTwoNumbers(linkedList1, linkedList2)
	if resultLinkedList.Val != 8 {
		t.Fail()
	}
	if resultLinkedList.Next.Next.Val != 9 {
		t.Fail()
	}
	if resultLinkedList.Next.Next.Next.Next.Next.Next.Next.Val != 1 {
		t.Fail()
	}

}
