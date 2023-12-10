package day10

import (
	"aoc2023/utils"
	"testing"
)

var input = utils.ParseFile("input.txt")

func BenchmarkPart1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part1(input, 'F')
	}
}

// func BenchmarkPart2(b *testing.B) {
//
// 	for i := 0; i < b.N; i++ {
// 		part2(input)
// 	}
// }
