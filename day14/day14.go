package day14

import (
	"aoc2023/utils"
	"fmt"
	"slices"
)

type Pos struct {
	i, j int
}

type Field struct {
	field        [][]rune
	nrows, ncols int
}

func initField(input []string) Field {
	nrows := len(input)
	ncols := len(input[0])

	field := make([][]rune, 0, nrows)
	for i := 0; i < nrows; i++ {
		line := make([]rune, 0, ncols)
		for _, r := range input[i] {
			line = append(line, r)
		}
		field = append(field, line)
	}
	return Field{field: field, nrows: nrows, ncols: ncols}
}

func (f *Field) print() {

	fmt.Printf("Field: %d x %d:\n", f.nrows, f.ncols)
	for i := 0; i < f.nrows; i++ {
		for j := 0; j < f.ncols; j++ {
			fmt.Printf("%s", string(f.field[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (f *Field) firstEmptyNorth(p Pos) Pos {
	for k := p.i - 1; k >= 0; k-- {
		if k == 0 && f.field[k][p.j] == '.' {
			return Pos{k, p.j}
		} else if f.field[k][p.j] != '.' {
			return Pos{k + 1, p.j}
		}
	}
	return p
}

func (f *Field) firstEmptyEast(p Pos) Pos {
	for k := p.j + 1; k < f.ncols; k++ {
		if k == f.ncols-1 && f.field[p.i][k] == '.' {
			return Pos{p.i, k}
		} else if f.field[p.i][k] != '.' {
			return Pos{p.i, k - 1}
		}
	}
	return p
}

func (f *Field) firstEmptySouth(p Pos) Pos {
	for k := p.i + 1; k < f.nrows; k++ {
		if k == f.nrows-1 && f.field[k][p.j] == '.' {
			return Pos{k, p.j}
		} else if f.field[k][p.j] != '.' {
			return Pos{k - 1, p.j}
		}
	}
	return p
}

func (f *Field) firstEmptyWest(p Pos) Pos {
	for k := p.j - 1; k >= 0; k-- {
		if k == 0 && f.field[p.i][k] == '.' {
			return Pos{p.i, k}
		} else if f.field[p.i][k] != '.' {
			return Pos{p.i, k + 1}
		}
	}
	return p
}

func (f *Field) rollNorth() {
	for i := 0; i < f.nrows; i++ {
		for j := 0; j < f.ncols; j++ {
			if f.field[i][j] == 'O' {
				p := f.firstEmptyNorth(Pos{i, j})
				if p.i != i {
					f.field[i][j], f.field[p.i][p.j] = f.field[p.i][p.j], f.field[i][j]
				}
			}
		}
	}
}

func (f *Field) rollEast() {
	for i := 0; i < f.nrows; i++ {
		for j := f.ncols - 1; j >= 0; j-- {
			if f.field[i][j] == 'O' {
				p := f.firstEmptyEast(Pos{i, j})
				if p.j != j {
					f.field[i][j], f.field[p.i][p.j] = f.field[p.i][p.j], f.field[i][j]
				}
			}
		}
	}
}

func (f *Field) rollSouth() {
	for i := f.nrows - 1; i >= 0; i-- {
		for j := 0; j < f.ncols; j++ {
			if f.field[i][j] == 'O' {
				p := f.firstEmptySouth(Pos{i, j})
				if p.i != i {
					f.field[i][j], f.field[p.i][p.j] = f.field[p.i][p.j], f.field[i][j]
				}
			}
		}
	}
}

func (f *Field) rollWest() {
	for i := 0; i < f.nrows; i++ {
		for j := 0; j < f.ncols; j++ {
			if f.field[i][j] == 'O' {
				p := f.firstEmptyWest(Pos{i, j})
				if p.j != j {
					f.field[i][j], f.field[p.i][p.j] = f.field[p.i][p.j], f.field[i][j]
				}
			}
		}
	}
}

func (f *Field) calcLoad() int {
	load := 0
	for i := 0; i < f.nrows; i++ {
		for j := 0; j < f.ncols; j++ {
			if f.field[i][j] == 'O' {
				load += f.nrows - i
			}
		}
	}
	return load
}

func part1(input []string) int {
	// Parse
	field := initField(input)

	// Roll north
	field.rollNorth()

	// Calculate load
	return field.calcLoad()
}

func (f *Field) spin() int {
	f.rollNorth()
	f.rollWest()
	f.rollSouth()
	f.rollEast()
	return f.calcLoad()

}

func findCycle(loads []int) (int, int, bool) {
	last := len(loads)
	for k := 3; k < len(loads)/2; k++ {
		if slices.Equal(loads[last-k:last], loads[last-2*k:last-k]) {
			return last - k, last, true
		}
	}
	return -1, -1, false
}

func part2(input []string) int {

	// Parse
	field := initField(input)

	iters := 1000000000

	var start, end int
	loads := make([]int, 0)
	for i := 0; i < iters; i++ {
		load := field.spin()
		var ok bool
		if i > 3 {
			if start, end, ok = findCycle(loads); ok {
				break
			}
		}

		loads = append(loads, load)

	}

	// calc remainder
	rem := (iters - start - 1) % (end - start)
	return loads[start+rem]

}

func Solve() {
	// Parse input
	input := utils.ParseFile("day14/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
