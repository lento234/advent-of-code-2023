package day12

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strings"
)

type Record []int

func findRecord(springs string) Record {
	record := make(Record, 0)

	count := 0
	for _, r := range springs {
		if r == '#' {
			count += 1
		} else if count > 0 {
			record = append(record, count)
			count = 0
		}
	}
	if count > 0 {
		record = append(record, count)
	}
	return record
}

func findArrangement(springs string, record Record) int {

	if !strings.Contains(springs, "?") {
		parsedRecord := findRecord(springs)
		if slices.Equal(record, parsedRecord) {
			return 1
		}
		return 0
	}

	L := findArrangement(strings.Replace(springs, "?", "#", 1), record)
	R := findArrangement(strings.Replace(springs, "?", ".", 1), record)

	return L + R
}

func part1(input []string) int {
	result := 0

	for _, line := range input {
		// Parse
		splitted := strings.SplitN(line, " ", 2)
		springs := strings.TrimSpace(splitted[0])
		record := utils.StringToNumbers(strings.TrimSpace(splitted[1]), ",")

		// fmt.Printf("%d: %s -> %v\n", i, springs, record)
		// Status
		value := findArrangement(springs, record)

		// fmt.Printf("%d: %s -> %v (%d)\n", i, springs, record, value)
		result += value
	}
	return result
}

func repeat(word, sep string, count int) string {
	repeated := make([]string, 0)
	for i := 0; i < count; i++ {
		repeated = append(repeated, word)
	}
	return strings.Join(repeated, sep)
}

func part2(input []string) int {
	result := 0
	repeatCount := 5

	for i, line := range input {
		// if i == 0 {
		// Parse
		splitted := strings.SplitN(line, " ", 2)
		springs := strings.TrimSpace(splitted[0])
		unfoldedSprings := repeat(springs, "?", repeatCount)

		recordStr := strings.TrimSpace(splitted[1])
		unfoldedRecordStr := repeat(recordStr, ",", repeatCount)
		unfoldedRecord := utils.StringToNumbers(unfoldedRecordStr, ",")

		// fmt.Printf("%d: %s -> %v\n", i, unfoldedSprings, unfoldedRecord)
		// Status
		value := findArrangement(unfoldedSprings, unfoldedRecord)

		fmt.Printf("%d: %s -> %v (%d)\n", i, unfoldedSprings, unfoldedRecord, value)
		// }
		result += value
	}
	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day12/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
