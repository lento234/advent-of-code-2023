package day02

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(input []string) int {

	result := 0

game:
	for _, line := range input {
		gameInfo := strings.SplitN(line, ":", 2)

		// Get game id
		gameID, err := strconv.Atoi(strings.TrimSuffix(strings.SplitN(gameInfo[0], " ", 2)[1], " "))
		utils.CheckErr(err)

		// fmt.Println("Game:", gameID)

		// Check dices
		draws := strings.Split(gameInfo[1], ";")
		for _, draw := range draws {
			// fmt.Println("draw", i, "->", draw)
			dicesInfo := strings.Split(strings.Trim(draw, " "), ",")

			dices := map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}
			for _, diceInfo := range dicesInfo {
				diceInfo := strings.SplitN(strings.Trim(diceInfo, " "), " ", 2)
				n, err := strconv.Atoi(diceInfo[0])
				utils.CheckErr(err)
				dices[diceInfo[1]] = n
			}
			// fmt.Println("draw", i, "->", dices)

			// Check valid games
			if dices["red"] > 12 || dices["green"] > 13 || dices["blue"] > 14 {
				// fmt.Println("Game", gameID, "draw", i, "not valid!")
				continue game
			}
		}

		// fmt.Println("Game", gameID, "valid!")
		result += gameID
	}
	return result
}

func part2(input []string) int {
	result := 0

	for _, line := range input {
		gameInfo := strings.SplitN(line, ":", 2)

		minDices := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		// Check dices
		draws := strings.Split(gameInfo[1], ";")
		for _, draw := range draws {
			// fmt.Println("draw", i, "->", draw)
			dicesInfo := strings.Split(strings.Trim(draw, " "), ",")

			dices := map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}
			for _, diceInfo := range dicesInfo {
				diceInfo := strings.SplitN(strings.Trim(diceInfo, " "), " ", 2)
				n, err := strconv.Atoi(diceInfo[0])
				utils.CheckErr(err)
				dices[diceInfo[1]] = n
			}
			// fmt.Println("draw", i, "->", dices)

			for k, v := range dices {
				minDices[k] = max(minDices[k], v)
			}

		}
		power := 1
		for _, v := range minDices {
			if v > 0 {
				power *= v
			}
		}
		// fmt.Println("Game", gameID, "min dices ->", minDices, "power:", power)
		result += power
	}

	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day02/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Test part 1
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
