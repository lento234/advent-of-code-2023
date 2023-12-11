package day11

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	answer := 374

	// Solve
	result := part1(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 1 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
}

func TestPart2(t *testing.T) {

	// Test part 2
	puzzle := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	expansions := []int{10, 100}
	answers := []int{1030, 8410}

	for i := 0; i < len(expansions); i++ {
		// Solve
		result := part2(utils.ParseString(puzzle), expansions[i])

		if result != answers[i] {
			t.Fatalf("Failed test (%d/%d)!: %v != %v", i+1, len(expansions), answers[i], result)
		}
		fmt.Printf("Part 1 (%d/%d) [%s]: %v == %v\n", i+1, len(expansions), utils.FormatGreen("solved"), answers[i], result)

	}
}
