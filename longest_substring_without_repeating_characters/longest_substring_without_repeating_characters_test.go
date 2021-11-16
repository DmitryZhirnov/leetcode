package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var testString string
	var result int

	testString = ""
	result = lengthOfLongestSubstring(testString)
	if result != 0 {
		fmt.Printf("Fail with testString:=%s and result:=%v", testString, result)
		println()
		t.Fail()
	}

	testString = "abcabcbb"
	result = lengthOfLongestSubstring(testString)
	if result != 3 {
		fmt.Printf("Fail with testString:=%s and result:=%v", testString, result)
		println()
		t.Fail()
	}

	testString = "bbbbb"
	result = lengthOfLongestSubstring(testString)
	if result != 1 {
		fmt.Printf("Fail with testString:=%s and result:=%v", testString, result)
		println()
		t.Fail()
	}

	testString = "pwwkew"
	result = lengthOfLongestSubstring(testString)
	if result != 3 {
		fmt.Printf("Fail with testString:=%s and result:=%v", testString, result)
		println()
		t.Fail()
	}

	testString = "tmmzuxt"
	result = lengthOfLongestSubstring(testString)
	if result != 5 {
		fmt.Printf("Fail with testString:=%s and result:=%v", testString, result)
		println()
		t.Fail()
	}

}
