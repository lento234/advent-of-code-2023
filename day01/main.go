package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseFile(filename string) []string {

	buf, err := os.ReadFile(filename)
	checkErr(err)
	text := string(buf)
	input := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
	return input
}

func solvePart1(input []string) {
	// fmt.Println(input)

	r, err := regexp.Compile("\\d")
	checkErr(err)

	result := 0

	for _, line := range input {
		allNumbers := r.FindAllString(line, -1)
		val, err := strconv.Atoi(fmt.Sprintf("%s%s", allNumbers[0], allNumbers[len(allNumbers)-1]))
		checkErr(err)
		result += val
	}

	fmt.Println("Solution to part 1:", result)

}

func solvePart2(input []string) {

	result := 0

	lsubstr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	wsubstr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range input {
		allNumbers := make(map[int]int, 0)
		for i, s := range lsubstr {
			if strings.Contains(line, s) {
				r, err := regexp.Compile(s)
				checkErr(err)
				for _, matches := range r.FindAllStringIndex(line, -1) {
					allNumbers[matches[0]] = i + 1
				}
			}
		}
		for i, s := range wsubstr {
			if strings.Contains(line, s) {
				r, err := regexp.Compile(s)
				checkErr(err)
				for _, matches := range r.FindAllStringIndex(line, -1) {
					allNumbers[matches[0]] = i + 1
				}
			}
		}

		indices := make([]int, 0)
		for k := range allNumbers {
			indices = append(indices, k)
		}
		slices.Sort(indices)

		val, err := strconv.Atoi(fmt.Sprintf("%d%d", allNumbers[indices[0]], allNumbers[indices[len(indices)-1]]))
		checkErr(err)

		// fmt.Println(i, ":", line, "->", allNumbers, ">>", val)
		result += val
	}

	fmt.Println("Solution to part 2:", result)

}

func main() {
	fmt.Println("Part 1:")

	// Test part 1
	test_input_part1 := parseFile("test_input_part1.txt")
	solvePart1(test_input_part1)

	// Test part 1
	test_input_part2 := parseFile("test_input_part2.txt")
	solvePart2(test_input_part2)

	input := parseFile("input.txt")

	// Part 1
	solvePart1(input)

	// Test part 1
	solvePart2(input)
}
