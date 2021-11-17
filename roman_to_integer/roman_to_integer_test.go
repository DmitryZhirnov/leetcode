package roman_to_integer

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	input := "I"
	result := romanToInt(input)
	expected := 1
	assertEqual(input, result, expected, t)

	input = "II"
	result = romanToInt(input)
	expected = 2
	assertEqual(input, result, expected, t)

	input = "IV"
	result = romanToInt(input)
	expected = 4
	assertEqual(input, result, expected, t)

	input = "XIV"
	result = romanToInt(input)
	expected = 14
	assertEqual(input, result, expected, t)

	input = "DXIV"
	result = romanToInt(input)
	expected = 514
	assertEqual(input, result, expected, t)

	input = "CMXLIV"
	result = romanToInt(input)
	expected = 944
	assertEqual(input, result, expected, t)

	input = "MMMCMXLIV"
	result = romanToInt(input)
	expected = 3944
	assertEqual(input, result, expected, t)

}

func assertEqual(input string, result int, expected int, t *testing.T) {
	if result != expected {
		t.Fatalf("%s != %v", input, result)
	}
}
