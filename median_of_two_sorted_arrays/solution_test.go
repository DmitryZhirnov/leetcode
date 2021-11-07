package median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	result := findMedianSortedArrays([]int{1}, []int{1})
	expected := 1.0
	if result != expected {
		fmt.Println(result, expected)
		t.Fail()
	}

	result = findMedianSortedArrays([]int{1}, []int{2,3})
	expected = 2.0
	if result != expected {
		fmt.Println(result, expected)
		t.Fail()
	}

	result = findMedianSortedArrays([]int{2}, []int{})
	expected = 2.0
	if result != expected {
		fmt.Println(result, expected)
		t.Fail()
	}

	result = findMedianSortedArrays([]int{}, []int{1,2,4})
	expected = 2.0
	if result != expected {
		fmt.Println()
		fmt.Println(result, expected)
		t.Fail()
	}

	result = findMedianSortedArrays([]int{1,2}, []int{3,4})
	expected = 2.5
	if result != expected {
		fmt.Println()
		fmt.Println(result, expected)
		t.Fail()
	}
}
