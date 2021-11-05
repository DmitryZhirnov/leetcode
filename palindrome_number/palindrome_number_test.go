package palindrome_number

import "testing"

func TestPalindromeNumber(t *testing.T) {

	if !IsPalindrome(9){
		t.Fail()
	}

	if IsPalindrome(78){
		t.Fatalf("78 is not palindrome")
	}

	if !IsPalindrome(12321){
		t.Fatalf("78 is not palindrome")
	}
}