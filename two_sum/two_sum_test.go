package two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 5, 7, 9}
	target := 12
	expected := []int{1, 2}

	result := TwoSumWithHashMap(nums, target)
	fmt.Println(result)
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
	//
	//nums = []int{2,5,5,11}
	//target = 10
	//expected = []int{1, 2}
	//result = TwoSum(nums, target)
	//fmt.Println(result)
	//if !reflect.DeepEqual(expected, result) {
	//	t.Fail()
	//}
}
