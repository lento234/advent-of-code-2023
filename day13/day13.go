package day13

import (
	"aoc2023/utils"
	"fmt"
)

type Pattern []string

func (p *Pattern) print() {
	for _, line := range *p {
		fmt.Println(line)
	}
}

func parsePatterns(input []string) []Pattern {

	patterns := make([]Pattern, 0)

	start := 0
	for i, line := range input {
		if len(line) == 0 {
			patterns = append(patterns, input[start:i])
			start = i + 1
		}
	}
	patterns = append(patterns, input[start:])
	return patterns
}

func (p *Pattern) getRow(i int) string {
	return (*p)[i]
}

func (p *Pattern) getCol(j int) string {
	nrows := len(*p)
	col := make([]byte, 0, nrows)

	for i := 0; i < nrows; i++ {
		col = append(col, (*p)[i][j])
	}
	return string(col)
}

func (p *Pattern) isHorizMirror(i int) bool {
	nrows := len(*p)

	for k := 0; k < min(i+1, nrows-i-1); k++ {
		u, d := p.getRow(i-k), p.getRow(i+k+1)
		if u != d {
			return false
		}
	}
	return true
}

func (p *Pattern) isVertMirror(j int) bool {
	ncols := len((*p)[0])

	for k := 0; k < min(j+1, ncols-j-1); k++ {
		l, r := p.getCol(j-k), p.getCol(j+k+1)
		if l != r {
			return false
		}
	}
	return true
}

func (p *Pattern) findReflection() (int, error) {

	nrows := len(*p)
	ncols := len((*p)[0])

	// line := (*p)[i]
	for j := 0; j < ncols-1; j++ {
		l, r := (*p)[0][j], (*p)[0][j+1]
		if l == r && p.isVertMirror(j) {
			return j + 1, nil
		}
	}

	for i := 0; i < nrows-1; i++ {
		l, r := (*p)[i][0], (*p)[i+1][0]
		if l == r && p.isHorizMirror(i) {
			return (i + 1) * 100, nil
		}
	}

	return -1, fmt.Errorf("no mirror")
}

func part1(input []string) int {
	result := 0

	// Parse patterns
	patterns := parsePatterns(input)

	for _, p := range patterns {
		value, err := p.findReflection()
		utils.CheckErr(err)
		result += value
	}

	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day13/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
