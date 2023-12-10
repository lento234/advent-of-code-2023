package day10

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzles := []string{`.....
.F-7.
.|.|.
.L-J.
.....`, `.....
.S-7.
.|.|.
.L-J.
.....`, `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`, `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`}

	startRunes := []rune{'F', 'S', 'S', 'S'}
	answers := []int{4, 4, 4, 8}

	for i := 0; i < len(puzzles); i++ {
		// Solve
		result := part1(utils.ParseString(puzzles[i]), startRunes[i])

		if result != answers[i] {
			t.Fatalf("Failed test (%d/%d)!: %v != %v", i+1, len(puzzles), answers[i], result)
		}
		fmt.Printf("Part 1 (%d/%d) [%s]: %v == %v\n", i+1, len(puzzles), utils.FormatGreen("solved"), answers[i], result)

	}
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
