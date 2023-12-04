package day04

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strings"
)

func part1(input []string) int {
	result := 0

	totalWinnings := make([]int, 0)

	for _, line := range input {

		game := strings.SplitN(line, ":", 2)
		scratchcard := strings.SplitN(game[1], "|", 2)
		winnings := utils.StringToNumbers(scratchcard[0], " ")
		numbers := utils.StringToNumbers(scratchcard[1], " ")

		// Count number of winnings
		nWinnings := 0
		validWinnings := make([]int, 0)
		for _, value := range winnings {
			if slices.Contains(numbers, value) {
				nWinnings += 1
				validWinnings = append(validWinnings, value)
			}
		}
		if nWinnings > 0 {
			totalWinnings = append(totalWinnings, nWinnings)
		}
		// fmt.Printf("%d: %s -> %v (%d)\n", i, line, validWinnings, nWinnings)
	}
	// fmt.Printf("Total winnings: %v\n", totalWinnings)
	for _, nWinning := range totalWinnings {
		result += (1 << (nWinning - 1))
	}
	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() error {
	// Parse input
	input := utils.ParseFile("day04/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)

	return nil
}
