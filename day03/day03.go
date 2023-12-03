package day03

import (
	"aoc2023/utils"
	"fmt"
	"strings"
	"unicode"
)

type Pos struct {
	i, j int
}

func getAdjacentNumbers(input []string, pos Pos) {

	for i := pos.i - 1; i <= pos.i+1; i++ {
		line := []rune(input[i])
		for k := 0; k < len(line); {
			r := line[k]
			if unicode.IsDigit(r) {
				start := k
				digit := ""
				for unicode.IsDigit(line[k]) {
					digit = strings.Join([]string{digit, string(line[k])}, "")
					k += 1
				}
				end := k
				fmt.Printf("%d, %d -> %d: %s\n", i, start, end, digit)
			}
			k += 1
		}
		// fmt.Printf("%v -> %s\n", pos, line)
	}
	fmt.Println()
}

func part1(input []string) int {
	result := 0

	fmt.Println(strings.Join(input, "\n"))

	for i, line := range input {
		// cleanedLine := strings.ReplaceAll(line, ".", "")
		// cleanedLine := strings.Split(line, ".")
		// fmt.Printf("%s -> %v\n", line, re.FindAllStringSubmatch(line, -1))
		for j, l := range line {
			if strings.Contains("*#@/=&$%-", string(l)) {
				pos := Pos{i, j}
				getAdjacentNumbers(input, pos)
				// fmt.Printf("(%d, %d) -> %c\n", i, j, l)
			}
		}
		// for i, l := range cleanedLine {
		// 	fmt.Printf("%d:%s, ", i, l)
		// }
		// fmt.Println()
	}
	// fmt.Println(input)
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
