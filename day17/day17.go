package day17

import (
	"aoc2023/utils"
	"container/heap"
	"fmt"
	"math"
	"slices"
)

// Grid
type Grid struct {
	field [][]int
	// visit        [][]bool
	// minHeat      [][]int
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
	// visit := make([][]bool, 0, nrows)
	// minHeat := make([][]int, 0, nrows)

	// const maxInt = int(^uint(0) >> 1)

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
		// visit = append(visit, visitLine)
		// minHeat = append(minHeat, heatLine)
	}
	// return Grid{field: field, visit: visit, minHeat: minHeat, nrows: nrows, ncols: ncols}
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

// func (g *Grid) mark(p Pos) {
// 	g.visit[p.i][p.j] = true
// }

// func (g *Grid) visited(p Pos) bool {
// 	return g.visit[p.i][p.j]
// }

func (g *Grid) inside(p Pos) bool {
	return p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols
}

func (g *Grid) getNeighbors(b Block) []Block {
	neighbors := make([]Block, 0)

	// p := b.p

	// Add the same direction
	if b.csteps < 3 && (b.dir != Pos{0, 0}) {
		np := b.p.Add(b.dir)
		if g.inside(np) {
			neighbors = append(neighbors,
				Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: b.dir, csteps: b.csteps + 1},
			)
		}
	}

	// Add other direction
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
	// fmt.Printf("(neighbors): %+v -> %+v\n", b, neighbors)

	return neighbors
}

// Block
type Block struct {
	heat int
	p    Pos
	// dir    Dir
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
	// result := 0

	grid := parseGrid(input)
	grid.print()

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 0}, csteps: 0})
	// heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{1, 0}, csteps: 0})
	// heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 1}, csteps: 0})

	visited := make([]Visit, 0)

	iter := 0
	for pq.Len() > 0 {
		// Mark
		iter++
		b := heap.Pop(&pq).(Block)

		// Early exit
		if b.p.Equal(Pos{grid.nrows - 1, grid.ncols - 1}) {
			fmt.Printf("result => %+v\n", b)
			return b.heat
		}

		// // Skip if visited
		if slices.Contains(visited, Visit{b.p, b.dir, b.csteps}) {
			continue
		}
		visited = append(visited, Visit{b.p, b.dir, b.csteps})

		// Get neighbors
		// Updating queue
		for _, n := range grid.getNeighbors(b) {
			// fmt.Printf("n: %+v\n", n)
			// Calculate new heat
			// if n.heat < grid.minHeat[n.p.i][n.p.j] && n.csteps <= 3 {
			// grid.minHeat[n.p.i][n.p.j] = n.heat
			// if n.csteps <= 3 {
			heap.Push(&pq, n)
			// }
			// }
		}
		// if iter > 10 {
		// 	break
		// }
	}
	fmt.Printf("Total iterations = %d -> %+v\n", iter, pq[0])
	return 0 //grid.minheat[grid.nrows-1][grid.ncols-1]
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day17/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
