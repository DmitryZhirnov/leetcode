package roman_to_integer

func romanToInt(s string) int {
	romanDigits := map[byte]int{'M':1000, 'D':500, 'C':100, 'L':50, 'X':10, 'V':5, 'I':1}

	if len(s) == 0 {
		return 0
	}
	result, current, previous := 0, 0, 0
	for index := len(s) - 1; index >= 0; index-- {
		current = romanDigits[s[index]]
		if current >= previous{
			result += current
		}
		if current < previous{
			result-=current
		}
		previous = current
	}
	return result

}
