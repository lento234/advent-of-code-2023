package day03

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Pos struct {
	i, j int
}

type Digit struct {
	start, end int
	value      int
}

func StringToDigits(text string) ([]Digit, error) {
	textRunes := []rune(text)

	digits := make([]Digit, 0)

	for k := 0; k < len(textRunes); {
		r := textRunes[k]
		if unicode.IsDigit(r) {
			digit := Digit{}
			digit.start = k
			value := ""
			for k < len(textRunes) && unicode.IsDigit(textRunes[k]) {
				value = strings.Join([]string{value, string(textRunes[k])}, "")
				k += 1
			}
			digit.end = k - 1
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			digit.value = valueInt
			digits = append(digits, digit)
		}
		k += 1
	}
	return digits, nil
}

func getAdjacentNumbers(input []string, pos Pos) []int {

	numbers := make([]int, 0)

	for i := pos.i - 1; i <= pos.i+1; i++ {
		digits, err := StringToDigits(input[i])
		if err != nil {
			log.Fatal(err)
		}
		for _, digit := range digits {
			if pos.j <= digit.end+1 && pos.j >= digit.start-1 {
				// isClose := false
				// if pos.j <= digit.end && pos.j >= digit.start {
				// 	isClose = true
				// }
				// fmt.Printf("%v: %d -> %v\n", pos, i, digit)
				numbers = append(numbers, digit.value)
			}
			// else if pos.j <= digit.end+1 && pos.j >= digit.start-1 {
			// isClose := false
			// if pos.j <= digit.end+1 && pos.j >= digit.start-1 {
			// 	isClose = true
			// }
			// fmt.Printf("(top/bot) %v: %d -> %v\n", pos, i, digit)
			// numbers = append(numbers, digit.value)
			// }
			// fmt.Printf("%v: %d -> %v\n", pos, i, digit)
		}
	}
	return numbers
}

func part1(input []string) int {
	result := 0

	parts := make([]int, 0)

	for i, line := range input {
		for j, l := range line {
			if strings.Contains("+*#@/=&$%-", string(l)) {
				numbers := getAdjacentNumbers(input, Pos{i, j})
				fmt.Printf("(%d, %d): %c -> %v\n", i, j, l, numbers)
				for _, number := range numbers {
					// if !slices.Contains(parts, number) {
					parts = append(parts, number)
					// }
				}
			}
		}
	}
	// Sum all parts
	for _, p := range parts {
		result += p
	}
	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() error {
	// Parse input
	input := utils.ParseFile("day03/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)

	return nil
}
