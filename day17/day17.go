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
	visit        [][]bool
	minHeat      [][]int
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
	visit := make([][]bool, 0, nrows)
	minHeat := make([][]int, 0, nrows)

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
		visit = append(visit, visitLine)
		minHeat = append(minHeat, heatLine)
	}
	return Grid{field: field, visit: visit, minHeat: minHeat, nrows: nrows, ncols: ncols}
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

func (g *Grid) mark(p Pos) {
	g.visit[p.i][p.j] = true
}

func (g *Grid) visited(p Pos) bool {
	return g.visit[p.i][p.j]
}

func (g *Grid) inside(p Pos) bool {
	return p.i >= 0 && p.i < g.nrows && p.j >= 0 && p.j < g.ncols
}

func (g *Grid) getNeighbors(b Block) []Block {
	neighbors := make([]Block, 0)

	p := b.p

	// Add the same direction
	if g.inside(np) {
		neighbors = append(neighbors,
			Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: b.dir, csteps: b.csteps + 1},
			// Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: b.dir, csteps: b.csteps + 1},
		)
	}
	// }

	// // Add other direction
	// for _, dir := range []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
	// np := b.p.Add(b.dir)

	// 	if g.inside(np) {

	// 		// Same direction
	// 		if b.dir == dir {
	// 			neighbors = append(neighbors,
	// 				Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: b.dir, csteps: b.csteps + 1},
	// 			)
	// 		}
	// 		if b.dir != dir && b.dir != dir.Neg() {
	// 			neighbors = append(neighbors,
	// 				Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: dir, csteps: 1},
	// 			)
	// 		}
	// 	}
	// }

	// if b.csteps < 3 {
	// 	if !b.dir.Equal(dir) && !b.dir.Equal(dir.Neg()) {
	// 		np := b.p.Add(dir)
	// 		if g.inside(np) {
	// 			neighbors = append(neighbors,
	// 				Block{heat: b.heat + g.field[np.i][np.j], p: np, dir: dir, csteps: 1},
	// 			)
	// 		}
	// 	}
	// }

	// Up
	// if p.i > 0 && !b.dir.Equal(Pos{1, 0}) {
	// 	neighbors = append(neighbors, Block{p: Pos{p.i - 1, p.j}, dir: Pos{-1, 0}})
	// }
	// // Down
	// if p.i < g.nrows-1 && !b.dir.Equal(Pos{-1, 0}) {
	// 	neighbors = append(neighbors, Block{p: Pos{p.i + 1, p.j}, dir: Pos{1, 0}})
	// }
	// // Left
	// if p.j > 0 && !b.dir.Equal(Pos{0, 1}) {
	// 	neighbors = append(neighbors, Block{p: Pos{p.i, p.j - 1}, dir: Pos{0, -1}})
	// }
	// // Right
	// if p.j < g.ncols-1 && !b.dir.Equal(Pos{0, -1}) {
	// 	neighbors = append(neighbors, Block{p: Pos{p.i, p.j + 1}, dir: Pos{0, 1}})
	// }

	// for i := 0; i < len(neighbors); i++ {
	// 	neighbors[i].csteps = 1
	// 	if neighbors[i].dir == b.dir {
	// 		neighbors[i].csteps = b.csteps + 1
	// 	}
	// 	np := neighbors[i].p
	// 	neighbors[i].heat = b.heat + g.field[np.i][np.j]
	// }
	// finalNeighbors := make([]Block, 0)
	// for _, n := range neighbors {
	// 	if n.csteps <= 3 {
	// 		finalNeighbors = append(finalNeighbors, n)
	// 	}
	// }

	// fmt.Printf("(neighbors): %+v -> %+v\n", b, neighbors)

	// return finalNeighbors
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

func part1(input []string) int {
	// result := 0

	grid := parseGrid(input)
	grid.print()

	pq := PriorityQueue{}
	heap.Init(&pq)
	// heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 0}, csteps: 1})
	heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{1, 0}, csteps: 0})
	heap.Push(&pq, Block{heat: 0, p: Pos{0, 0}, dir: Pos{0, 1}, csteps: 0})

	iter := 0
	for pq.Len() > 0 {
		// Mark
		iter++
		b := heap.Pop(&pq).(Block)

		// Early exit
		if b.p.Equal(Pos{grid.nrows - 1, grid.ncols - 1}) {
			fmt.Printf("result => %+v\n", b)
			break
		}

		// Skip if visited
		if grid.visited(b.p) {
			continue
		}
		// mark visited
		grid.mark(b.p)

		// Get neighbors
		// Updating queue
		for _, n := range grid.getNeighbors(b) {
			// Calculate new heat
			// if n.heat < grid.minHeat[n.p.i][n.p.j] && n.csteps <= 3 {
			// grid.minHeat[n.p.i][n.p.j] = n.heat
			// if n.csteps <= 3 {
			heap.Push(&pq, n)
			// }
			// }
		}

		// 			// Same direction
		// 			if !(b.di == 0 && b.dj == 0) && b.csteps < 3 {
		// 				ni := b.i + b.di
		// 				nj := b.j + b.dj
		// 				if grid.inside(ni, nj) {
		// 					nblock := Block{heat: b.heat + grid.field[ni][nj], i: ni, j: nj, di: b.di, dj: b.dj, csteps: b.csteps + 1}
		// 					// fmt.Printf("%+v -> %+v (same dir added)\n", b, nblock)
		// 					heap.Push(&pq, nblock)
		// 				}
		// 			}

		// 		for _, nd := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		// 			ndi, ndj := nd[0], nd[1]
		// 			if !(ndi == b.di && ndj == b.dj) && !(ndi == -b.di && ndj == -b.dj) {
		// 				ni := b.i + ndi
		// 				nj := b.j + ndj
		// 				// fmt.Println(ndi, ndj)
		// 				if grid.inside(ni, nj) {
		// 					nblock := Block{heat: b.heat + grid.field[ni][nj], i: ni, j: nj, di: ndi, dj: ndj, csteps: 1}
		// 					// fmt.Printf("%+v -> %+v (same dir added)\n", b, nblock)
		// 					heap.Push(&pq, nblock)
		// 				}
		// 			}
		// 		}
		// if (n.dir != b.dir) && (n.dir != b.dir.Neg()) {
		// 	fmt.Printf("%+v -> %+v (diff dir added)\n", b, n)
		// 	heap.Push(&pq, n)
		// }

		// if iter > 1 {
		// 	break
		// }
	}
	// heap.Push(&pq, Block{pos: Pos{1, 0}, heat: 4})
	// heap.Push(&pq, Block{pos: Pos{1, 1}, heat: 2})
	// // heap.Push(B

	// for pq.Len() > 0 {
	// 	b := heap.Pop(&pq).(Block)
	// 	fmt.Printf("%+v\n", b)
	// }
	fmt.Printf("Total iterations = %d\n", iter)
	// grid.printVisit()
	// grid.printHeat()

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
