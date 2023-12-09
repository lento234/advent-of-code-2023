package day09

import (
	"aoc2023/utils"
	"fmt"
)

func history(numbers []int) int {

	// check if layer is all zeros
	allZeros := true
	for _, n := range numbers {
		allZeros = allZeros && (n == 0)
	}
	if allZeros {
		return numbers[len(numbers)-1]
	}
	// calculate next layer
	next := make([]int, 0)
	for i := 1; i < len(numbers); i++ {
		next = append(next, numbers[i]-numbers[i-1])
	}

	return numbers[len(numbers)-1] + history(next)
}

func part1(input []string) int {
	result := 0

	for _, line := range input {
		// Parse numbers
		numbers := utils.StringToNumbers(line, " ")
		// Recursive history
		result += history(numbers)
	}

	// calc history
	// 1992273652
	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day09/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
