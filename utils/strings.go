package utils

import (
	"strconv"
	"strings"
	"unicode"
)

type Digit struct {
	Start, End int
	Value      int
}

func StringToDigits(text string) ([]Digit, error) {
	textRunes := []rune(text)

	digits := make([]Digit, 0)

	for k := 0; k < len(textRunes); {
		r := textRunes[k]
		if unicode.IsDigit(r) {
			digit := Digit{}
			digit.Start = k
			value := ""
			for k < len(textRunes) && unicode.IsDigit(textRunes[k]) {
				value = strings.Join([]string{value, string(textRunes[k])}, "")
				k += 1
			}
			digit.End = k - 1
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			digit.Value = valueInt
			digits = append(digits, digit)
		}
		k += 1
	}
	return digits, nil
}

func StringToNumbers(text, delimiter string) []int {

	numbers := make([]int, 0)

	numbersStr := strings.Split(text, delimiter)

	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
