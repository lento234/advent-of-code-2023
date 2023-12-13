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

func diff(left, right string) int {
	n := 0
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			n++
		}
	}
	return n
}

func (p *Pattern) isHorizMirror(i int, nMaxSmudges int) bool {
	nrows := len(*p)

	n := 0
	for k := 0; k < min(i+1, nrows-i-1); k++ {
		u, d := p.getRow(i-k), p.getRow(i+k+1)
		n += diff(u, d)
		if n > nMaxSmudges {
			return false
		}
	}
	return n == nMaxSmudges
}

func (p *Pattern) isVertMirror(j int, nMaxSmudges int) bool {
	ncols := len((*p)[0])

	n := 0
	for k := 0; k < min(j+1, ncols-j-1); k++ {
		l, r := p.getCol(j-k), p.getCol(j+k+1)
		n += diff(l, r)
		if n > nMaxSmudges {
			return false
		}
	}
	return n == nMaxSmudges
}

func (p *Pattern) findReflection(nMaxSmudges int) (int, error) {

	nrows := len(*p)
	ncols := len((*p)[0])

	for j := 0; j < ncols-1; j++ {
		if p.isVertMirror(j, nMaxSmudges) {
			return j + 1, nil
		}
	}

	for i := 0; i < nrows-1; i++ {
		if p.isHorizMirror(i, nMaxSmudges) {
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
		value, err := p.findReflection(0)
		utils.CheckErr(err)
		result += value
	}

	return result
}

func part2(input []string) int {
	result := 0

	// Parse patterns
	patterns := parsePatterns(input)

	for _, p := range patterns {
		value, err := p.findReflection(1)
		utils.CheckErr(err)
		result += value
	}

	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day13/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
