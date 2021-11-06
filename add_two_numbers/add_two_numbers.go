package add_two_numbers

import (
	"alhoritms/common"
)

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	values1 := getValues(l1)
	values2 := getValues(l2)
	values1, values2 = appendZero(values1, values2)
	result := sumArrays(values1, values2)
	return FillLinkedList(result, 0)
}

func sumArrays(values1 []int, values2 []int) []int {
	var result []int
	var discharge, sum int
	for i := 0; i < len(values1); i++ {
		sum = values1[i] + values2[i] + discharge
		result = append(result, sum%10)
		if sum > 9 {
			discharge = 1
		} else {
			discharge = 0
		}
	}
	if discharge == 1 {
		result = append(result, discharge)
	}
	return result
}

func appendZero(values1 []int, values2 []int) ([]int, []int) {
	if len(values1) > len(values2) {
		zeros := make([]int, len(values1)-len(values2))
		values2 = append(values2, zeros...)
	} else {
		zeros := make([]int, len(values2)-len(values1))
		values1 = append(values1, zeros...)
	}
	return values1, values2
}

func getValues(l *common.ListNode) []int {

	var values = []int{l.Val}
	for l.Next != nil {
		values = append(values, l.Next.Val)
		l = l.Next
	}
	return values
}

// FillLinkedList Заполнение связанного списка
func FillLinkedList(items []int, index int) *common.ListNode {
	var linkedList *common.ListNode
	if index < len(items) {
		current := &common.ListNode{
			Val:  items[index],
			Next: FillLinkedList(items, index+1),
		}
		linkedList = current
	}
	return linkedList
}
