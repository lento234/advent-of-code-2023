package day16

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strconv"
)

type Grid struct {
	field        [][]rune
	nrows, ncols int
	energized    [][][]rune
}

type Pos struct {
	i, j int
}

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Beam struct {
	pos   Pos
	dir   Dir
	steps int
}

func (b *Beam) move() {
	switch b.dir {
	case Up:
		b.pos.i--
	case Down:
		b.pos.i++
	case Left:
		b.pos.j--
	case Right:
		b.pos.j++
	}
	b.steps++
}

func (b *Beam) dirStr() string {
	switch b.dir {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return "unknown"
}

func parseGrid(input []string) Grid {
	nrows := len(input)
	ncols := len(input[0])

	field := make([][]rune, 0, nrows)
	for _, line := range input {
		field = append(field, []rune(line))
	}
	energized := make([][][]rune, 0, nrows)
	for i := 0; i < nrows; i++ {
		cols := make([][]rune, 0, ncols)
		for j := 0; j < ncols; j++ {
			cols = append(cols, []rune{})
		}
		energized = append(energized, cols)
	}
	return Grid{field: field, nrows: nrows, ncols: ncols, energized: energized}
}

func (g *Grid) print() {
	// Info
	fmt.Printf("Grid: %d x %d:\n", g.nrows, g.ncols)

	for _, line := range g.field {
		fmt.Println(string(line))
	}
}

func (g *Grid) printEnergized() {
	// Info
	fmt.Printf("Grid: %d x %d:\n", g.nrows, g.ncols)

	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			e := g.energized[i][j]
			if len(e) == 0 {
				fmt.Printf("%s", string('.'))
				// } else if len(e) == 1 {
				// 	fmt.Printf("%s", string(e[0]))
				// } else {
				// 	fmt.Printf("%d", len(e))
				// }
			} else {
				fmt.Printf("%s", "#")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) totalEnergized() int {
	result := 0
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			if len(g.energized[i][j]) > 0 {
				result++
			}
		}
	}
	return result
}

func (g *Grid) get(p Pos) rune {
	return g.field[p.i][p.j]
}

func (g *Grid) update(beam Beam) {
	switch beam.dir {
	case Up:
		g.energized[beam.pos.i][beam.pos.j] = append(g.energized[beam.pos.i][beam.pos.j], '^')
	case Down:
		g.energized[beam.pos.i][beam.pos.j] = append(g.energized[beam.pos.i][beam.pos.j], 'v')
	case Left:
		g.energized[beam.pos.i][beam.pos.j] = append(g.energized[beam.pos.i][beam.pos.j], '<')
	case Right:
		g.energized[beam.pos.i][beam.pos.j] = append(g.energized[beam.pos.i][beam.pos.j], '>')
	}

	switch g.field[beam.pos.i][beam.pos.j] {
	case '.', '^', 'v', '<', '>', '1', '2', '3', '4':
		if len(g.energized[beam.pos.i][beam.pos.j]) == 1 {
			g.field[beam.pos.i][beam.pos.j] = g.energized[beam.pos.i][beam.pos.j][0]
		} else {
			s := strconv.Itoa(len(g.energized[beam.pos.i][beam.pos.j]))
			g.field[beam.pos.i][beam.pos.j] = rune(s[0])
		}
	}
}

func (g *Grid) isInside(beam Beam) bool {
	return beam.pos.i >= 0 && beam.pos.i < g.nrows && beam.pos.j >= 0 && beam.pos.j < g.ncols
}
func (g *Grid) contains(beam Beam) bool {
	p := beam.pos
	var r rune
	switch beam.dir {
	case Up:
		r = '^'
	case Down:
		r = 'v'
	case Left:
		r = '<'
	case Right:
		r = '>'
	}
	return slices.Contains(g.energized[p.i][p.j], r)
}

func (g *Grid) trace(beam Beam) []Beam {
	splittedBeams := make([]Beam, 0)
	// fmt.Printf("Start tracing beam: %v\n", beam)
	// outer:
	for {
		beam.move()
		if !g.isInside(beam) || g.contains(beam) {
			break
		}
		g.update(beam)
		switch g.get(beam.pos) {
		case '/':
			// prev := beam.dir
			switch beam.dir {
			case Up:
				beam.dir = Right
			case Down:
				beam.dir = Left
			case Left:
				beam.dir = Down
			case Right:
				beam.dir = Up
			}
			// fmt.Printf("encountered '/' %v -> %v\n", prev, beam.dir)
		case '\\':
			// prev := beam.dir
			switch beam.dir {
			case Up:
				beam.dir = Left
			case Down:
				beam.dir = Right
			case Left:
				beam.dir = Up
			case Right:
				beam.dir = Down
			}
			// fmt.Printf("encountered '\\' %v -> %v\n", prev, beam.dir)
		case '|':
			// prev := beam.dir
			switch beam.dir {
			case Left, Right:
				// fmt.Printf("splitting beam up and down at %v\n", beam.pos)
				splittedBeams = append(splittedBeams, Beam{pos: beam.pos, dir: Up})
				splittedBeams = append(splittedBeams, Beam{pos: beam.pos, dir: Down})
				return splittedBeams
			}
			// fmt.Printf("encountered '|' %v -> %v\n", prev, beam.dir)
		case '-':
			// prev := beam.dir
			switch beam.dir {
			case Up, Down:
				// fmt.Printf("splitting beam left and right at %v\n", beam.pos)
				splittedBeams = append(splittedBeams, Beam{pos: beam.pos, dir: Left})
				splittedBeams = append(splittedBeams, Beam{pos: beam.pos, dir: Right})
				return splittedBeams
			}
			// fmt.Printf("encountered '|' %v -> %v\n", prev, beam.dir)
		}
	}
	// fmt.Printf("Finished tracing in %d steps, pos: %v, dir: %v (new: %d)\n", beam.steps, beam.pos, beam.dir, len(splittedBeams))
	return splittedBeams
}

func part1(input []string) int {
	// Parse grid
	grid := parseGrid(input)
	queue := utils.Queue[Beam]{}

	// First beam
	queue.Push(Beam{pos: Pos{i: 0, j: -1}, dir: Right})

	for !queue.Empty() {
		// Pop
		beam, err := queue.Pop()
		utils.CheckErr(err)

		// Trace
		splittedBeams := grid.trace(beam)

		// Adding
		for _, b := range splittedBeams {
			queue.Push(b)
		}
	}
	result := grid.totalEnergized()
	return result
}

func totalEnergization(input []string, start Beam) int {
	// Parse grid
	grid := parseGrid(input)

	queue := utils.Queue[Beam]{}

	// First beam
	queue.Push(start)

	for !queue.Empty() {
		// Pop
		beam, err := queue.Pop()
		utils.CheckErr(err)

		// Trace
		splittedBeams := grid.trace(beam)

		// Adding
		for _, b := range splittedBeams {
			queue.Push(b)
		}
	}
	return grid.totalEnergized()
}

func part2(input []string) int {
	nrows := len(input)
	ncols := len(input[0])
	allEnergizations := make([]int, 0, 2*ncols+2*nrows)

	// All rows
	for i := 0; i < nrows; i++ {
		value := totalEnergization(input, Beam{pos: Pos{i, -1}, dir: Right})
		allEnergizations = append(allEnergizations, value)
		value = totalEnergization(input, Beam{pos: Pos{i, ncols}, dir: Left})
		allEnergizations = append(allEnergizations, value)
	}
	// All cols
	for j := 0; j < ncols; j++ {
		value := totalEnergization(input, Beam{pos: Pos{-1, j}, dir: Down})
		allEnergizations = append(allEnergizations, value)
		value = totalEnergization(input, Beam{pos: Pos{nrows, j}, dir: Up})
		allEnergizations = append(allEnergizations, value)
	}
	// fmt.Println("all:", allEnergizations)
	return slices.Max(allEnergizations)
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day16/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
