package day17

import (
	"aoc2023/utils"
	"container/heap"
	"fmt"
	"math"
)

// Grid
type Grid struct {
	field        [][]int
	nrows, ncols int
}

// Pos and Dir
type Pos struct {
	i, j int
}

func (p Pos) Neg() Pos {
	return Pos{-p.i, -p.j}
}
func (p Pos) Add(other Pos) Pos {
	return Pos{p.i + other.i, p.j + other.j}
}

func (p Pos) Equal(other Pos) bool {
	return p.i == other.i && p.j == other.j
}

func parseGrid(input []string) Grid {
	nrows := len(input)
	ncols := len(input[0])

	field := make([][]int, 0, nrows)

	for _, line := range input {
		parsedLine := make([]int, 0, ncols)
		visitLine := make([]bool, 0, ncols)
		heatLine := make([]int, 0, ncols)
		for _, r := range line {
			parsedLine = append(parsedLine, int(r-'0'))
			visitLine = append(visitLine, false)
			heatLine = append(heatLine, math.MaxInt)
		}
		field = append(field, parsedLine)
	}
	return Grid{field: field, nrows: nrows, ncols: ncols}
}

func (g *Grid) print() {

	fmt.Printf("Grid: %d x %d:\n", g.ncols, g.ncols)

	for _, line := range g.field {
		for _, c := range line {
			fmt.Printf("%d", c)
		}
		fmt.Println()
	}
}

func (g *Grid) inside(p Pos) bool {
	return p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols
}

func (g *Grid) getNeighbors(b Block, maxConsecutive int, minConsecutive int) []Block {
	neighbors := make([]Block, 0)

	// p := b.p

	// Add the same direction
	if b.csteps < maxConsecutive && (b.dir != Pos{0, 0}) {
		np := b.p.Add(b.dir)
		if g.inside(np) {
			neighbors = append(neighbors,
				Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: b.dir, csteps: b.csteps + 1},
			)
		}
	}

	// Add other direction
	if b.csteps >= minConsecutive {
		for _, dir := range []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			np := b.p.Add(dir)

			if g.inside(np) {
				if b.dir != dir && b.dir != dir.Neg() {
					neighbors = append(neighbors,
						Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: dir, csteps: 1},
					)
				}
			}
		}
	}

	return neighbors
}

// Block
type Block struct {
	heat   int
	p      Pos
	dir    Pos
	csteps int // consecutive steps
}

// Priority queue
type PriorityQueue []Block

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heat < pq[j].heat
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	l := len(old)
	e := old[l-1]
	*pq = old[0 : l-1]
	return e
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(Block)
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type Visit struct {
	p, dir Pos
	csteps int
}

func part1(input []string) int {
	grid := parseGrid(input)

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 0}, csteps: 0})

	visited := make(map[Visit]bool, 0)

	iter := 0
	for pq.Len() > 0 {
		// Mark
		iter++
		b := heap.Pop(&pq).(Block)

		// Early exit
		if b.p.Equal(Pos{grid.nrows - 1, grid.ncols - 1}) {
			// fmt.Printf("result => %+v (iterations = %d)\n", b, iter)
			return b.heat
		}

		// Skip if visited
		v := Visit{b.p, b.dir, b.csteps}
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = true

		// Get neighbors and update queue
		// Updating queue
		for _, n := range grid.getNeighbors(b, 3, 0) {
			heap.Push(&pq, n)
		}
	}
	// fmt.Printf("Total iterations = %d -> %+v\n", iter, pq[0])
	return -1
}

func part2(input []string) int {
	grid := parseGrid(input)
	// grid.print()

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 1}, csteps: 1})

	visited := make(map[Visit]bool, 0)

	iter := 0
	for pq.Len() > 0 {
		// Mark
		iter++
		b := heap.Pop(&pq).(Block)

		// Early exit
		if b.p.Equal(Pos{grid.nrows - 1, grid.ncols - 1}) {
			return b.heat
		}

		// Skip if visited
		v := Visit{b.p, b.dir, b.csteps}
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = true

		// Get neighbors and update queue
		// Updating queue
		for _, n := range grid.getNeighbors(b, 10, 4) {
			heap.Push(&pq, n)
		}
	}
	// fmt.Printf("Total iterations = %d -> %+v\n", iter, pq)

	return -1
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day17/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
