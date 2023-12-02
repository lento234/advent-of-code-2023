package day01

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func part1(input []string) int {

	r, err := regexp.Compile(`\d`)
	utils.CheckErr(err)

	result := 0

	for _, line := range input {
		allNumbers := r.FindAllString(line, -1)
		val, err := strconv.Atoi(fmt.Sprintf("%s%s", allNumbers[0], allNumbers[len(allNumbers)-1]))
		utils.CheckErr(err)
		result += val
	}

	return result

}

func part2(input []string) int {

	result := 0

	lsubstr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	wsubstr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range input {
		allNumbers := make(map[int]int, 0)
		for i, s := range lsubstr {
			if strings.Contains(line, s) {
				r, err := regexp.Compile(s)
				utils.CheckErr(err)
				for _, matches := range r.FindAllStringIndex(line, -1) {
					allNumbers[matches[0]] = i + 1
				}
			}
		}
		for i, s := range wsubstr {
			if strings.Contains(line, s) {
				r, err := regexp.Compile(s)
				utils.CheckErr(err)
				for _, matches := range r.FindAllStringIndex(line, -1) {
					allNumbers[matches[0]] = i + 1
				}
			}
		}

		indices := make([]int, 0)
		for k := range allNumbers {
			indices = append(indices, k)
		}
		slices.Sort(indices)

		val, err := strconv.Atoi(fmt.Sprintf("%d%d", allNumbers[indices[0]], allNumbers[indices[len(indices)-1]]))
		utils.CheckErr(err)

		result += val
	}

	return result
}

func Solve() error {
	// Parse input
	input := utils.ParseFile("day01/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Test part 1
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)

	return nil
}
