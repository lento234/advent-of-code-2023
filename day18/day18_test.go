package day18

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
	answer := 62

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
