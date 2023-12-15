package day15

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

func parseInput(input []string) []string {
	line := input[0]
	return strings.Split(line, ",")
}

func hashLetter(seed int, letter rune) int {
	// start := seed
	seed += int(letter)
	seed *= 17
	seed %= 256
	// fmt.Printf("seed: %d, letter: %s -> %d\n", start, string(letter), seed)
	return seed
}

func hashWord(word string) int {
	seed := 0
	for _, l := range word {
		seed = hashLetter(seed, l)
		// fmt.Println(l)
	}
	return seed
}

func part1(input []string) int {
	result := 0
	seq := parseInput(input)

	for _, s := range seq {
		value := hashWord(s)
		// fmt.Printf("%02d: %s -> %d\n", i, s, value)
		result += value
	}
	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day15/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
