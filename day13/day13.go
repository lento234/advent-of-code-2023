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
		// fmt.Printf("%02d: %s\n", i, line)
	}
	patterns = append(patterns, input[start:])
	return patterns
}

func reverse(line string) string {
	rev := []rune(line)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func (p *Pattern) isVertMirror(j int) bool {
	ncols := len((*p)[0])
	nrows := len(*p)
	// fmt.Printf("isMirror: %d, %d\n", j, ncols)

	for i := 0; i < nrows; i++ {
		for k := 0; k < min(j+1, ncols-j-1); k++ {
			// fmt.Printf("%d, %d, %d, %d -> %s, %s\n", k, j, j-k, j+k+1, string(line[j-k]), string(line[j+k+1]))
			l, r := (*p)[i][j-k], (*p)[i][j+k+1]
			if l != r {
				return false
			}
		}
	}
	return true
}

func (p *Pattern) isHorizMirror(i int) bool {
	// ncols := len((*p)[0])
	nrows := len(*p)
	// fmt.Printf("isMirror: %s\n", (*p)[i])

	// for j := 0; j < ncols; j++ {
	for k := 0; k < min(i+1, nrows-i-1); k++ {
		u, d := (*p)[i-k], (*p)[i+k+1]
		// fmt.Printf("%d, %d, %d, %d -> %s, %s\n", k, i, i-k, i+k+1, string(u), string(d))
		if u != d {
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
			// fmt.Printf("vertical mirror: %d\n", j+1)
			// break
			return j + 1, nil
		}
		// for j := 1; j < ncols; j++ {
		// 	left := line[:j]
		// 	right := reverse(line[j-1:])
		// 	fmt.Printf("%s => %s -> %s  : %v\n", line, line[:j], reverse(line[j-1:]), idx)

		// }
	}

	for i := 0; i < nrows-1; i++ {
		l, r := (*p)[i][0], (*p)[i+1][0]
		if l == r && p.isHorizMirror(i) {
			// fmt.Printf("horizontal mirror: %d\n", i)
			// break
			return (i + 1) * 100, nil
		}
		// for j := 1; j < ncols; j++ {
		// 	left := line[:j]
		// 	right := reverse(line[j-1:])
		// 	fmt.Printf("%s => %s -> %s  : %v\n", line, line[:j], reverse(line[j-1:]), idx)

		// }
	}

	return -1, fmt.Errorf("no mirror")
}

func part1(input []string) int {
	result := 0

	patterns := parsePatterns(input)

	// patterns[0].print()

	// find vert ref
	for _, p := range patterns {
		value, err := p.findReflection()
		utils.CheckErr(err)
		fmt.Println("value:", value)
		result += value
	}

	// for i, p := range patterns {
	// 	fmt.Printf("Pattern %d:\n", i)
	// 	p.print()
	// 	fmt.Println()
	// }

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
