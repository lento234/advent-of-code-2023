package day15

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzles := []string{
		"HASH",
		"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
	}
	answers := []int{52, 1320}

	for i := 0; i < len(puzzles); i++ {
		// Solve
		result := part1(utils.ParseString(puzzles[i]))

		if result != answers[i] {
			t.Fatalf("Failed test (%d/%d)!: %v != %v", i+1, len(puzzles), answers[i], result)
		}
		fmt.Printf("Part 1 (%d/%d) [%s]: %v == %v\n", i+1, len(puzzles), utils.FormatGreen("solved"), answers[i], result)

	}
}

func TestPart2(t *testing.T) {

	// Test part 2
	puzzle := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	answer := 145

	// Solve
	result := part2(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 2 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
}
