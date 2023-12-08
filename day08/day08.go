package day08

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

func parseLR(line string) []int {
	insStr := strings.TrimSpace(line)

	ins := make([]int, 0)
	for _, r := range insStr {
		if r == 'L' {
			ins = append(ins, 0)
		} else {

			ins = append(ins, 1)
		}
	}
	return ins
}

type Node struct {
	value string
	next  []string
}

// type Tree struct {
// 	root *Node
// }

// func (t *Tree) insert(value string) *Node {
// 	if t.root == nil {
// 		t.root = &Node{value: value, left: nil, right: nil}
// 	} else {
// 		t.root.insert(value)
// 	}
// 	return t
// }

type Path map[string]Node

func parsePath(input []string) Path {
	path := make(Path, 0)

	for i, line := range input {
		if i < 2 {
			continue
		}
		key := strings.TrimSpace(strings.SplitN(line, "=", 2)[0])
		LRStr := strings.TrimSpace(strings.SplitN(line, "=", 2)[1])
		LR := strings.SplitN(LRStr, ",", 2)
		left := strings.TrimSpace(strings.Replace(LR[0], "(", "", 1))
		right := strings.TrimSpace(strings.Replace(LR[1], ")", "", 1))

		if _, ok := path[key]; !ok {
			path[key] = Node{value: key, next: []string{left, right}}
		}
		// if p, ok := path[left]; ok {
		// 	p.left = left
		// 	path[left] = p
		// } else {
		// 	path[left] = Node{value: left, left: "", right: ""}
		// }
		// if p, ok := path[right]; ok {
		// 	p.right = right
		// } else {
		// 	path[right] = Node{value: right, left: "", right: ""}
		// }
		// root.pos = key
		// path[key].pos = key
		// if val, ok := path[L]; ok {
		// 	*(path[key]).L = &val
		// }
		// if _, ok := path[R]; ok {
		// 	*(path[key]).R = path[R]
		// }
		// fmt.Printf("%d: %s -> [%s, %s] (%s)\n", i, key, L, R, line)
	}
	return path
}

func part1(input []string) int {

	ins := parseLR(input[0])
	// fmt.Println("ins ->", ins)

	path := parsePath(input)

	// key := strings.Split()
	// key := strings.TrimSpace(strings.SplitN(input[2], "=", 2)[0])
	key := "AAA"
	end := "ZZZ"
	i := 0
	k := 0
	for {
		if key == end {
			break
		}
		// fmt.Printf("%d: (%d) %s -> %v\n", i, k, key, path[key])
		key = path[key].next[ins[i]]
		i = (i + 1) % len(ins)
		k += 1

	}
	// for i, ins := range instructions {
	// 	fmt.Printf("%s -> %+v\n", key, value)
	// }

	// path := make(map[string]Pos, 0)

	return k
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day08/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
