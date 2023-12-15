package day15

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []string) []string {
	line := input[0]
	return strings.Split(line, ",")
}

func hashLetter(seed int, letter rune) int {
	// start := seed
	seed += int(letter)
	seed *= 17
	seed %= 256
	// fmt.Printf("seed: %d, letter: %s -> %d\n", start, string(letter), seed)
	return seed
}

func hashWord(word string) int {
	seed := 0
	for _, l := range word {
		seed = hashLetter(seed, l)
		// fmt.Println(l)
	}
	return seed
}

func part1(input []string) int {
	result := 0
	seq := parseInput(input)

	for _, s := range seq {
		value := hashWord(s)
		// fmt.Printf("%02d: %s -> %d\n", i, s, value)
		result += value
	}
	return result
}

func parseWord(word string) (int, string, int) {
	var box int
	var lens string
	var amount int
	if strings.Contains(word, "-") {
		lens = strings.Split(word, "-")[0]
		amount = -1
	} else {
		split := strings.Split(word, "=")
		lens = split[0]
		var err error
		amount, err = strconv.Atoi(split[1])
		utils.CheckErr(err)
	}
	box = hashWord(lens)

	return box, lens, amount
}

type Boxes map[int]utils.OrderedMap

func (b *Boxes) Add(box int, lens string, fl int) {
	if _, ok := (*b)[box]; !ok {
		(*b)[box] = utils.OrderedMap{}
	}

	// Add lens
	lenses := (*b)[box]
	lenses.Add(lens, fl)
	(*b)[box] = lenses
}

func (b *Boxes) Remove(box int, lens string) {
	lenses := (*b)[box]
	lenses.Remove(lens)
	(*b)[box] = lenses
}

func (b *Boxes) totalFocusingPower() int {
	result := 0
	for i, box := range *b {
		for j, lens := range box.OrderedKeys {
			value := (i + 1) * (j + 1) * (box.Items[lens])
			// fmt.Printf("%02d: box -> %v, lens -> %d, %v -> %d => %d\n", i+1, box, j+1, lens, box.items[lens], value)
			result += value
		}
	}
	return result
}

func part2(input []string) int {
	seq := parseInput(input)

	boxes := make(Boxes, 0)
	for _, word := range seq {
		box, lens, fl := parseWord(word)
		if fl > 0 {
			boxes.Add(box, lens, fl)
		} else {
			boxes.Remove(box, lens)
		}

	}

	return boxes.totalFocusingPower()
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day15/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
