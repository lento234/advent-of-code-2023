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

func TestPart2(t *testing.T) {

	// Test part 2
	puzzles := []string{`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`, `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`, `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`}

	answers := []int{4, 8, 10}

	for i := 0; i < len(puzzles); i++ {
		// Solve
		result := part2(utils.ParseString(puzzles[i]))

		if result != answers[i] {
			t.Fatalf("Failed test (%d/%d)!: %v != %v", i+1, len(puzzles), answers[i], result)
		}
		fmt.Printf("Part 1 (%d/%d) [%s]: %v == %v\n", i+1, len(puzzles), utils.FormatGreen("solved"), answers[i], result)

	}
}
