package day08

import (
	"aoc2023/utils"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	answer := 6

	// Solve
	result := part1(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 1 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)

	// Test part 1
	puzzle = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	answer = 2

	// Solve
	result = part1(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 1 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
}

func TestPart2(t *testing.T) {

	// Test part 2
	puzzle := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	answer := 6

	// Solve
	result := part2(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	fmt.Printf("Part 2 [%s]: %v == %v\n", utils.FormatGreen("solved"), answer, result)
}
