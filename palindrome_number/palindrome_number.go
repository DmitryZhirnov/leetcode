package palindrome_number

func IsPalindrome(x int) bool {
	num := x
	if x < 0{
		return false
	}
	reverse := 0
	for num > 0 {
		reverse = reverse * 10 + num % 10
		num = num/10
	}
	if x != reverse {
		return false
	}
	return true
}