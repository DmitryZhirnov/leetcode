package my_atoi

import (
	"fmt"
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var result int32
	result = MyAtoi("34567")
	if result != 34567 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("     -42")
	if result != -42 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("00000042")
	if result != 42 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("4193 with words")
	if result != 4193 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("words and 987")
	if result != 0 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("-91283472332")
	if result != math.MinInt32 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("9223372036854775808")
	if result != 2147483647 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}

	result = MyAtoi("+-12")
	if result != 0 {
		fmt.Printf("resutl: %v", result)
		t.Fail()
	}
}
