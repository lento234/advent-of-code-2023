package day03

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Pos struct {
	i, j int
}

func getAdjacentNumbers(input []string, pos Pos) []int {

	numbers := make([]int, 0)

	for i := pos.i - 1; i <= pos.i+1; i++ {
		digits, err := utils.StringToDigits(input[i])
		utils.CheckErr(err)
		for _, digit := range digits {
			if pos.j <= digit.End+1 && pos.j >= digit.Start-1 {
				numbers = append(numbers, digit.Value)
			}
		}
	}
	return numbers
}

func part1(input []string) int {
	parts := make([]int, 0)

	for i, line := range input {
		for j, l := range line {
			if strings.Contains("+*#@/=&$%-", string(l)) {
				numbers := getAdjacentNumbers(input, Pos{i, j})
				parts = append(parts, numbers...)
			}
		}
	}
	// Sum all parts
	return utils.SumSlice(parts)
}

func part2(input []string) int {
	result := 0

	for i, line := range input {
		for j, r := range line {
			if string(r) == "*" {
				numbers := getAdjacentNumbers(input, Pos{i, j})
				if len(numbers) > 1 {
					result += utils.ProdSlice(numbers)
				}
			}
		}
	}
	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day03/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
