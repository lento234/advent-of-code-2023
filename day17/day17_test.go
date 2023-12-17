package day17

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`
	answer := 102

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
