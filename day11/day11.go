package day11

import (
	"aoc2023/utils"
	"fmt"
	"slices"
)

type Pos struct {
	i, j int
}

type Universe struct {
	ncols, nrows int
	galaxies     []Pos
}

func (u *Universe) print() {
	fmt.Printf("universe: %d x %d\n", u.nrows, u.ncols)
	for i := 0; i < u.nrows; i++ {
		if i == 0 {
			fmt.Printf("%s", "  ")
			for j := 0; j < u.ncols; j++ {
				fmt.Printf("v")
			}
			fmt.Println()
		}
		fmt.Printf("> ")
		for j := 0; j < u.ncols; j++ {
			p := Pos{i, j}
			if slices.Contains(u.galaxies, p) {
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", "•")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func makeUniverse(input []string) Universe {

	nrows := len(input)
	ncols := len(input[0])

	galaxies := make([]Pos, 0)
	for i, line := range input {
		for j, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Pos{i, j})
			}
		}
	}
	return Universe{nrows: nrows, ncols: ncols, galaxies: galaxies}
}
func (u *Universe) isemptyRow(i int) bool {
	for j := 0; j < u.ncols; j++ {
		if slices.Contains(u.galaxies, Pos{i, j}) {
			return false
		}
	}
	return true
}

func (u *Universe) isemptyCol(j int) bool {
	for i := 0; i < u.nrows; i++ {
		if slices.Contains(u.galaxies, Pos{i, j}) {
			return false
		}
	}
	return true
}
func (u *Universe) expand(e int) {

	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	for i := 0; i < u.nrows; i++ {
		if u.isemptyRow(i) {
			emptyRows = append(emptyRows, i)
		}
	}
	for j := 0; j < u.ncols; j++ {
		if u.isemptyCol(j) {
			emptyCols = append(emptyCols, j)
		}
	}
	expanded := make([]Pos, 0)

	for k := 0; k < len(u.galaxies); k++ {
		dx := 0
		for _, i := range emptyRows {
			if u.galaxies[k].i > i {
				dx += e
			}
		}
		dy := 0
		for _, j := range emptyCols {
			if u.galaxies[k].j > j {
				dy += e
			}
		}
		expanded = append(
			expanded,
			Pos{u.galaxies[k].i + dx, u.galaxies[k].j + dy},
		)
	}
	u.galaxies = expanded
	u.nrows += len(emptyRows) * e
	u.ncols += len(emptyCols) * e
}

func (u *Universe) totalShortestPath() int {

	// Run shortest distances
	result := 0
	for i := 0; i < len(u.galaxies); i++ {
		for j := 0; j < len(u.galaxies); j++ {
			if j > i {
				dx := u.galaxies[j].i - u.galaxies[i].i
				dy := u.galaxies[j].j - u.galaxies[i].j
				result += utils.Abs(dx) + utils.Abs(dy)
			}
		}
	}
	return result

}

func part1(input []string) int {
	// parse grid
	universe := makeUniverse(input)

	// Expand universe
	universe.expand(1)

	return universe.totalShortestPath()
}

func part2(input []string, expansion int) int {
	// parse grid
	universe := makeUniverse(input)

	// Expand universe
	universe.expand(expansion - 1)

	return universe.totalShortestPath()

}

func Solve() {
	// Parse input
	input := utils.ParseFile("day11/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input, 1_000_000)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
