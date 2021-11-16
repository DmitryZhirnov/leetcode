package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	mapTable := make(map[byte]int)
	start := 0
	end := 0
	max := 0

	for index, letter := range []byte(s) {
		end = index
		if position, ok := mapTable[letter]; ok && position >= start {
			start = mapTable[letter] + 1
		} else {
			if max < end-start+1 {
				max = end - start + 1
			}
		}
		mapTable[letter] = index
	}
	return max
}
