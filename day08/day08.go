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
	start bool
	end   bool
}

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

			start := false
			end := false
			if key[len(key)-1] == 'A' {
				start = true
			}
			if key[len(key)-1] == 'Z' {
				end = true
			}
			path[key] = Node{value: key, next: []string{left, right}, start: start, end: end}
		}
	}
	return path
}

func part1(input []string) int {

	ins := parseLR(input[0])

	path := parsePath(input)

	key := "AAA"
	end := "ZZZ"
	i := 0
	k := 0
	for {
		if key == end {
			break
		}
		key = path[key].next[ins[i]]
		i = (i + 1) % len(ins)
		k += 1

	}

	return k
}

func part2(input []string) int {
	ins := parseLR(input[0])

	path := parsePath(input)

	keys := make([]string, 0)
	for k, v := range path {
		if v.start {
			keys = append(keys, k)
		}
	}

	nFound := 0
	i := 0
	steps := 0
	cycles := make([]int, len(keys))
	for {
		for k := 0; k < len(keys); k++ {
			if path[keys[k]].end && cycles[k] == 0 {
				cycles[k] = steps
				nFound += 1
			}
		}
		if len(cycles) == nFound {
			break
		}
		for k := 0; k < len(keys); k++ {
			keys[k] = path[keys[k]].next[ins[i]]

		}
		i = (i + 1) % len(ins)
		steps += 1
	}

	return utils.LCMSlice(cycles)
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day08/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
