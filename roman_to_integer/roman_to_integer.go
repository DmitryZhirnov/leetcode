package roman_to_integer

import (
	"strings"
)

func romanToInt(s string) int {
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanDigits := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	return sum(s, 0, romanDigits, romanValues)

}

func sum(s string, val int, romanDigits []string, romanValues []int) int {

	if len(s) == 0 {
		return 0
	}
	result := 0
	for index, romanDigit := range romanDigits {
		position := strings.Index(s, romanDigit)
		if position == 0 {
			s = s[len(romanDigit):]
			result = romanValues[index] + sum(s, val, romanDigits, romanValues)
			return result
		}
	}
	return result
}
