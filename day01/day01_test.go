package day01

import (
	"aoc2023/utils"
	"testing"

	"github.com/charmbracelet/log"
)

func TestPart1(t *testing.T) {

	// Test part 1
	puzzle := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	answer := 142

	// Solve
	result := part1(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	log.Infof("Solved!: %v == %v", answer, result)
}

func TestPart2(t *testing.T) {

	// Test part 2
	puzzle := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	answer := 281

	// Solve
	result := part2(utils.ParseString(puzzle))

	if result != answer {
		t.Fatalf("Failed!: %v != %v", answer, result)
	}
	log.Infof("Solved!: %v == %v", answer, result)
}
