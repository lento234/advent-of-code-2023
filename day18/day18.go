package day18

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// type Dir Pos
type Color string

type Vec2 struct {
	i, j int
}

func (v *Vec2) Add(u Vec2) Vec2 {
	return Vec2{v.i + u.i, v.j + u.j}
}

func (v *Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v.i - u.i, v.j - u.j}
}

func Max(u, v Vec2) Vec2 {
	var i, j int
	if u.i < v.i {
		i = v.i
	} else {
		i = u.i
	}
	if u.j < v.j {
		j = v.j
	} else {
		j = u.j
	}
	return Vec2{i, j}
}

func Min(u, v Vec2) Vec2 {
	var i, j int
	if u.i < v.i {
		i = u.i
	} else {
		i = v.i
	}
	if u.j < v.j {
		j = u.j
	} else {
		j = v.j
	}
	return Vec2{i, j}
}

type Dig struct {
	p Vec2
	c Color
}

func strToDir(dirStr string) (Vec2, error) {
	switch dirStr {
	case "U":
		return Vec2{-1, 0}, nil
	case "D":
		return Vec2{1, 0}, nil
	case "L":
		return Vec2{0, -1}, nil
	case "R":
		return Vec2{0, 1}, nil
	}
	return Vec2{}, fmt.Errorf("parsing %s direction failed", dirStr)
}

func parsePlan(line string) (Vec2, int, Color) {
	splitted := strings.SplitN(line, " ", 3)
	dirStr := splitted[0]
	dir, err := strToDir(dirStr)
	utils.CheckErr(err)
	count, err := strconv.Atoi(splitted[1])
	utils.CheckErr(err)
	color := splitted[2][1 : len(splitted[2])-1] // assume
	return dir, count, Color(color)
}

type Grid struct {
	field        [][]rune
	nrows, ncols int
}

func (g *Grid) print() {
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			fmt.Print(string(g.field[i][j]))
		}
		fmt.Println()
	}
}

func makeGrid(plans *[]Dig, pMin, pMax Vec2) Grid {
	nrows := pMax.i - pMin.i + 1
	ncols := pMax.j - pMin.j + 1
	field := make([][]rune, 0, nrows)

	for i := 0; i < nrows; i++ {
		line := []rune(strings.Repeat(".", ncols))
		field = append(field, line)
	}

	for _, d := range *plans {
		relP := d.p.Sub(pMin)
		// fmt.Printf("%d -> %+v -> %+v\n", i, d, relP)
		field[relP.i][relP.j] = '#'
	}

	return Grid{field: field, nrows: nrows, ncols: ncols}
}

func (g *Grid) isInside(p Vec2) bool {
	return p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols
}

func (g *Grid) fill() {

	var start Vec2
loop:
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			if g.field[i][j] == '#' {
				start = Vec2{i + 1, j + 1}
				break loop
			}
		}
	}
	queue := utils.Queue[Vec2]{}
	queue.Push(start)

	for queue.Len() > 0 {
		p, _ := queue.PopBack()

		// color inside
		if g.field[p.i][p.j] == '.' {
			g.field[p.i][p.j] = '#'
		}
		// Add neighbors
		for _, dir := range []Vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			np := p.Add(dir)
			if g.isInside(np) && g.field[np.i][np.j] == '.' {
				queue.Push(np)
			}
		}

	}
}

func (g *Grid) area() int {
	result := 0
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			if g.field[i][j] == '#' {
				result++
			}
		}
	}
	return result
}

func part1(input []string) int {

	plans := make([]Dig, 0)
	p := Vec2{i: 0, j: 0}
	pMin, pMax := Vec2{i: math.MaxInt, j: math.MaxInt}, Vec2{0, 0}
	for _, line := range input {
		dir, count, color := parsePlan(line)
		for count > 0 {
			// Add to trench plans
			plans = append(plans, Dig{p: p, c: color})
			// Update
			p = p.Add(dir)
			count--
			// Update col, rows
			pMin = Min(pMin, p)
			pMax = Max(pMax, p)
		}
		// fmt.Printf("%d -> %+v, %d, %s (%d, %d)\n", i, dir, count, color, nrows, ncols)
		// break
	}
	// fmt.Printf("plans -> %+v (%d) -> (%+v, %+v)\n", plans, len(plans), pMin, pMax)

	// make grid
	grid := makeGrid(&plans, pMin, pMax)
	// grid.print()
	// fmt.Println()
	// fmt.Println("filled:")
	grid.fill()
	// grid.print()
	return grid.area()
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day18/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
