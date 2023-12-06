package day06

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calcWins(t, d int) int {
	root := math.Sqrt(float64(t*t - 4*(d+1)))

	fastest := int((float64(t) + root) / 2)
	slowest := t - fastest
	return fastest - slowest + 1
}

func part1(input []string) int {

	// Parse
	times := utils.StringToNumbers(strings.SplitN(input[0], ":", 2)[1], " ")
	distances := utils.StringToNumbers(strings.SplitN(input[1], ":", 2)[1], " ")

	totalWins := make([]int, 0)

	for i := 0; i < len(times); i++ {
		totalWins = append(totalWins, calcWins(times[i], distances[i]))
	}

	return utils.Prod(totalWins)
}

func part2(input []string) int {

	// Query times and distances
	t, err := strconv.Atoi(strings.ReplaceAll(strings.SplitN(input[0], ":", 2)[1], " ", ""))
	utils.CheckErr(err)
	d, err := strconv.Atoi(strings.ReplaceAll(strings.SplitN(input[1], ":", 2)[1], " ", ""))
	utils.CheckErr(err)

	return calcWins(t, d)
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day06/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
