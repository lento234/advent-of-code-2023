package day09

import (
	"aoc2023/utils"
	"fmt"
)

func history(numbers []int, index int, operator func(a, b int) int) int {

	// check if layer is all zeros
	allZeros := true
	for _, n := range numbers {
		if n != 0 {
			allZeros = false
			break
		}
	}
	if allZeros {
		return numbers[len(numbers)-1]
	}
	// calculate next layer
	next := make([]int, 0)
	for i := 1; i < len(numbers); i++ {
		next = append(next, numbers[i]-numbers[i-1])
	}

	idx := (len(numbers) + index) % len(numbers)

	return operator(numbers[idx], history(next, index, operator))
}

func part1(input []string) int {
	result := 0

	for _, line := range input {
		// Parse numbers
		numbers := utils.StringToNumbers(line, " ")
		// Recursive history
		result += history(numbers, -1, func(a, b int) int { return a + b })
	}
	return result
}

func part2(input []string) int {
	result := 0

	for _, line := range input {
		// Parse numbers
		numbers := utils.StringToNumbers(line, " ")
		// Recursive history
		result += history(numbers, 0, func(a, b int) int { return a - b })
	}
	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day09/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
