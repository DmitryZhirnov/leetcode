package my_atoi

import (
	"math"
)

func MyAtoi(s string) int32 {
	minus := false
	start := false
	var atoi = 0

	for _, letter := range s {
		if letter == ' ' {
			continue
		}

		if (letter == '+' || letter == '-') && !start {
			start = true
			if letter == '-' {
				minus = true
			}
			continue
		}
		if letter > '9' || letter < '0' {
			break
		}
		start = true
		atoi *= 10
		atoi += int(letter - 48)

		if atoi > math.MaxInt32 && !minus {
			return math.MaxInt32
		}

		if atoi > math.MaxInt32+1 && minus {
			return math.MinInt32
		}

	}
	if minus {
		atoi *= -1
	}
	return int32(atoi)
}
