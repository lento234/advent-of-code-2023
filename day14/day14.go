package day14

import (
	"aoc2023/utils"
	"fmt"
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

func (f *Field) firstEmptyUp(p Pos) Pos {
	// fmt.Println("got:", p)

	// if f.field[0][p.j] == '.' {
	// 	return Pos{0, p.j}
	// }

	for k := p.i - 1; k >= 0; k-- {
		// fmt.Printf("(%d, %d) -> k = %d\n", p.i, p.j, k)
		if k == 0 && f.field[k][p.j] == '.' {
			return Pos{k, p.j}
		} else if f.field[k][p.j] != '.' {
			// fmt.Println("here:", k, p.j, string(f.field[k][p.j]))
			return Pos{k + 1, p.j}
		}
		// return Pos{k, p.j}
	}
	return p
}

func (f *Field) rollUp() {
	for i := 0; i < f.nrows; i++ {
		for j := 0; j < f.ncols; j++ {
			if f.field[i][j] == 'O' {
				p := f.firstEmptyUp(Pos{i, j})
				if p.i != i {
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
	// result := 0

	field := initField(input)
	// field.print()

	field.rollUp()

	// load := field.calcLoad()

	// field.print()
	// // p := field.firstEmptyUp(Pos{3, 1})
	// fmt.Println("load:", load)

	return field.calcLoad()
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day14/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
