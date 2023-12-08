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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func part2(input []string) int {
	ins := parseLR(input[0])

	path := parsePath(input)

	startKeys := make([]string, 0)
	for k, v := range path {
		if v.start {
			startKeys = append(startKeys, k)
		}
	}
	keys := startKeys
	start := make([]int, len(keys))
	cycles := make([]int, len(keys))
	found := make([]int, len(keys))

	i := 0
	steps := 0
	for {
		isEnd := true
		for k := 0; k < len(keys); k++ {
			isEnd = isEnd && path[keys[k]].end
			if path[keys[k]].end {
				if found[k] == 0 {
					start[k] = steps
					found[k] += 1
				} else if found[k] == 1 {
					cycles[k] = steps - start[k]
					found[k] += 1
				}
			}
		}
		if isEnd || utils.Sum(found) >= len(cycles)*2 {
			break
		}
		for k := 0; k < len(keys); k++ {
			keys[k] = path[keys[k]].next[ins[i]]

		}
		i = (i + 1) % len(ins)
		steps += 1
	}

	result := LCM(cycles[0], cycles[1])
	for i := 2; i < len(cycles); i++ {
		result = LCM(result, cycles[i])
	}

	return result
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
