package day10

import (
	"aoc2023/utils"
	"fmt"
)

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

type Grid struct {
	grid         []rune
	nrows, ncols int
}

func makeGrid(input []string) Grid {

	nrows := len(input)
	ncols := len(input[0])
	grid := make([]rune, 0, nrows*ncols)

	for _, line := range input {
		for _, r := range line {
			grid = append(grid, r)
		}
	}
	return Grid{grid: grid, nrows: nrows, ncols: ncols}
}

func (g *Grid) findStartPos(start rune) (Pos, error) {
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			k := g.ncols*i + j
			if g.grid[k] == start {
				return Pos{i: i, j: j}, nil
			}
		}
	}
	// Did not find start
	return Pos{}, fmt.Errorf("did not find starting position: %v", start)
}

func (g *Grid) getRune(p Pos) (rune, error) {
	if p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols {
		return g.grid[g.ncols*p.i+p.j], nil
	}
	return '0', fmt.Errorf("position %v out of bounds", p)
}

func (g *Grid) setRune(p Pos, r rune) error {
	if p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols {
		g.grid[g.ncols*p.i+p.j] = r
		return nil
	}
	return fmt.Errorf("position %v out of bounds", p)
}

func (g *Grid) print(start Pos) {
	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			k := g.ncols*i + j
			value := string(g.grid[k])
			if start.i == i && start.j == j {
				value = utils.FormatGreen(value)
			}
			fmt.Printf(value)
		}
		fmt.Printf("\n")
	}
}

func (g *Grid) next(cur Pos, dir Dir) (Pos, Dir, error) {
	k := g.ncols*cur.i + cur.j
	// r := g.getRune(cur)
	r := g.grid[k]

	// - `|` is a _vertical pipe_ connecting north and south.
	// - `-` is a _horizontal pipe_ connecting east and west.
	// - `L` is a _90-degree bend_ connecting north and east.
	// - `J` is a _90-degree bend_ connecting north and west.
	// - `7` is a _90-degree bend_ connecting south and west.
	// - `F` is a _90-degree bend_ connecting south and east.
	if r == '|' && dir == Up {
		return Pos{cur.i - 1, cur.j}, dir, nil
	} else if r == '|' && dir == Down {
		return Pos{cur.i + 1, cur.j}, dir, nil
	} else if r == '-' && dir == Left {
		return Pos{cur.i, cur.j - 1}, dir, nil
	} else if r == '-' && dir == Right {
		return Pos{cur.i, cur.j + 1}, dir, nil
	} else if r == 'L' && dir == Left {
		return Pos{cur.i - 1, cur.j}, Up, nil
	} else if r == 'L' && dir == Down {
		return Pos{cur.i, cur.j + 1}, Right, nil
	} else if r == 'J' && dir == Right {
		return Pos{cur.i - 1, cur.j}, Up, nil
	} else if r == 'J' && dir == Down {
		return Pos{cur.i, cur.j - 1}, Left, nil
	} else if r == '7' && dir == Right {
		return Pos{cur.i + 1, cur.j}, Down, nil
	} else if r == '7' && dir == Up {
		return Pos{cur.i, cur.j - 1}, Left, nil
	} else if r == 'F' && dir == Up {
		return Pos{cur.i, cur.j + 1}, Right, nil
	} else if r == 'F' && dir == Left {
		return Pos{cur.i + 1, cur.j}, Down, nil
	} else {
		return Pos{}, Up, fmt.Errorf("unknown position: %v and dir %v", cur, dir)
	}

}

func (g *Grid) getNextValidDirection(p Pos) (Pos, Dir, error) {
	// Check all directions and (ignore errors)

	// Up
	nextPos := Pos{p.i - 1, p.j}
	nextRune, _ := g.getRune(nextPos)
	if nextRune == 'F' || nextRune == '7' || nextRune == '|' {
		return nextPos, Up, nil
	}
	// Down
	nextPos = Pos{p.i + 1, p.j}
	nextRune, _ = g.getRune(nextPos)
	if nextRune == 'J' || nextRune == 'L' || nextRune == '|' {
		return nextPos, Down, nil
	}
	// Left
	nextPos = Pos{p.i, p.j - 1}
	if nextRune == 'L' || nextRune == 'F' || nextRune == '-' {
		return nextPos, Left, nil
	}

	// Right
	nextPos = Pos{p.i, p.j + 1}
	nextRune, _ = g.getRune(nextPos)
	if nextRune == 'J' || nextRune == '7' || nextRune == '-' {
		return nextPos, Right, nil
	}

	return Pos{}, Up, fmt.Errorf("no directions found")
}

func part1(input []string, startRune rune) int {
	// result := 0

	// Make grid
	grid := makeGrid(input)

	// Starting position
	startPos, err := grid.findStartPos(startRune)
	utils.CheckErr(err)
	// fmt.Printf("Starting pos: %v\n", startPos)

	// Print grid
	// grid.print(startPos)

	// pos := Pos{i: startPos.i + 1, j: startPos.j}
	// dir := Down
	pos, dir, err := grid.getNextValidDirection(startPos)
	utils.CheckErr(err)

	steps := 1
	// fmt.Printf("Start travelling: from %v -> to %v: dir %v\n", startPos, pos, dir)
	for {
		r, err := grid.getRune(pos)
		utils.CheckErr(err)
		// fmt.Printf("%d: %s -> pos: %v, dir: %v\n", steps, string(r), pos, dir)

		if r == startRune {
			break
		}
		pos, dir, err = grid.next(pos, dir)
		utils.CheckErr(err)
		steps += 1
	}

	// fmt.Printf("Finished in %d steps\n", steps)
	// fmt.Println()

	return steps / 2
}

func initFill(g *Grid) Grid {

	fill := make([]rune, 0, g.nrows*g.ncols)

	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			fill = append(fill, '?')
		}
	}
	return Grid{grid: fill, nrows: g.nrows, ncols: g.ncols}
}

func part2(input []string) int {

	// Make grid
	grid := makeGrid(input)
	startRune := 'S'
	startPos, err := grid.findStartPos(startRune)
	utils.CheckErr(err)

	// grid.print(startPos)

	pos, dir, err := grid.getNextValidDirection(startPos)
	utils.CheckErr(err)

	fill := initFill(&grid)

	// fmt.Printf("Start travelling: from %v -> to %v: dir %v\n", startPos, pos, dir)
	for {
		r, err := grid.getRune(pos)
		utils.CheckErr(err)

		// // mark path
		err = fill.setRune(pos, r)
		utils.CheckErr(err)

		if r == startRune {
			break
		}
		pos, dir, err = grid.next(pos, dir)
		utils.CheckErr(err)
	}

	// flood fill
	marks := []rune{'0', '1'}
	for i := 0; i < fill.nrows; i++ {
		idx := 0
		for j := 0; j < fill.ncols; j++ {
			p := Pos{i, j}
			r, _ := fill.getRune(p)
			if r == 'F' || r == '7' || r == '|' || r == 'S' {
				idx = (idx + 1) % 2
			} else if r == '?' {
				fill.setRune(p, marks[idx])
			}
		}
	}

	// fmt.Println("After:")
	// fill.print(startPos)

	result := 0
	for i := 0; i < fill.nrows; i++ {
		for j := 0; j < fill.ncols; j++ {
			p := Pos{i, j}
			r, _ := fill.getRune(p)
			if r == '1' {
				result += 1
			}
		}
	}

	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day10/input.txt")

	// Part 1
	result := part1(input, 'S')
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
