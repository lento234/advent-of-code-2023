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

func (v *Vec2) Mul(c int) Vec2 {
	return Vec2{v.i * c, v.j * c}
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
	// c Color
}

func strToDir(dirStr string) (Vec2, error) {
	switch dirStr {
	case "U", "3":
		return Vec2{-1, 0}, nil
	case "D", "1":
		return Vec2{1, 0}, nil
	case "L", "2":
		return Vec2{0, -1}, nil
	case "R", "0":
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
			// Add neighbors
			for _, dir := range []Vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				np := p.Add(dir)
				if g.isInside(np) && g.field[np.i][np.j] == '.' {
					queue.Push(np)
				}
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

func shoelace(points []Vec2) int {
	result := 0
	for k := 0; k < len(points); k++ {
		p1 := points[k]
		var p2 Vec2
		if k == len(points)-1 {
			p2 = points[0]
		} else {
			p2 = points[k+1]
		}
		result += p1.i*p2.j - (p2.i * p1.j)
	}
	result /= 2
	if result < 0 {
		return -result
	}
	return result
}

func calcArea(plans []Dig) int {
	points := make([]Vec2, 0)
	perimeter := 0
	for _, plan := range plans {
		points = append(points, plan.p)
	}
	areaInside := shoelace(points)

	// perimeter := len(plans)
	return areaInside + perimeter/2 + 1
}

func part1(input []string) int {

	plans := make([]Dig, 0)
	p := Vec2{i: 0, j: 0}
	pMin, pMax := Vec2{i: math.MaxInt, j: math.MaxInt}, Vec2{0, 0}
	for _, line := range input {
		dir, count, _ := parsePlan(line)
		for count > 0 {
			// Add to trench plans
			plans = append(plans, Dig{p: p})
			// Update
			p = p.Add(dir)
			count--
			// Update col, rows
			pMin = Min(pMin, p)
			pMax = Max(pMax, p)
		}
	}

	// make grid
	grid := makeGrid(&plans, pMin, pMax)
	grid.fill()
	area := grid.area()
	return area
}

func parseCorrectPlan(line string) (Vec2, int64) {
	splitted := strings.SplitN(line, " ", 3)
	color := splitted[2][1 : len(splitted[2])-1] // assume
	dir, err := strToDir(color[len(color)-1:])
	utils.CheckErr(err)
	count, err := strconv.ParseInt(color[1:len(color)-1], 16, 64)
	utils.CheckErr(err)
	return dir, count
}

func part2(input []string) int {
	points := make([]Vec2, 0)

	p := Vec2{i: 0, j: 0}
	points = append(points, p)

	perimeter := 0
	for _, line := range input {
		dir, count := parseCorrectPlan(line)
		p = p.Add(dir.Mul(int(count)))
		points = append(points, p)
		perimeter += int(count)
	}

	// Area inside using shoelace
	areaInside := shoelace(points)
	return areaInside + perimeter/2 + 1
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day18/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
