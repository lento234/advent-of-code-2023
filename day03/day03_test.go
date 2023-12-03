package day03

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	answer := 4361

	// Solve
	result := part1(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 1 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
}

// func TestPart2(t *testing.T) {

// 	// Test part 2
// 	puzzle := `<puzzle>`
// 	answer := `<answer>`

// 	// Solve
// 	result := part2(utils.ParseString(puzzle))

// 	if result != answer {
// 		t.Fatalf("Failed!: %v != %v", answer, result)
// 	}
// 	fmt.Printf("Part 2 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
// }
