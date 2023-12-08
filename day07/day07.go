package day07

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func runeToStrength(card rune) int {
	if card == 'A' {
		return 14
	} else if card == 'K' {
		return 13
	} else if card == 'Q' {
		return 12
	} else if card == 'J' {
		return 11
	} else if card == 'T' {
		return 10
	} else {
		return int(card - '0')
	}
}

type Card struct {
	handStr string
	hand    []int
	bid     int
	// points int
	strength int
}

func parseHand(handStr string, bid int) Card {

	counts := make(map[int]int, 0)
	hand := make([]int, 0)
	for _, r := range handStr {
		s := runeToStrength(r)
		counts[s] += 1
		hand = append(hand, s)
	}
	strength := 0
	for _, count := range counts {
		strength += count * count
	}
	// fmt.Printf("%s -> %v : %v (%d)\n", handStr, hand, counts, strength)

	return Card{
		handStr,
		hand,
		bid,
		strength,
	}

}

func part1(input []string) int {
	result := 0
	cards := make([]Card, 0)

	for _, line := range input {
		splitted := strings.SplitN(line, " ", 2)
		hand := splitted[0]
		bid, err := strconv.Atoi(splitted[1])
		utils.CheckErr(err)
		// fmt.Printf("%d: %s -> %d\n", i, hand, bid)
		cards = append(cards, parseHand(hand, bid))
	}

	sortStrength := func(i, j int) bool {
		if cards[i].strength == cards[j].strength {
			for k := 0; k < len(cards[i].hand); k++ {
				if cards[i].hand[k] == cards[j].hand[k] {
					continue
				}
				return cards[i].hand[k] < cards[j].hand[k]
			}
		}
		return cards[i].strength < cards[j].strength
	}

	sort.Slice(cards, sortStrength)

	// totalWinnings := 0
	for i, card := range cards {
		// fmt.Printf("%d: %+v\n", i, card)
		result += card.bid * (i + 1)
	}
	// fmt.Println("Total winnings:", result)
	return result
}

// func part2(input []string) int {
// 	result := 0
// 	return result
// }

func Solve() {
	// Parse input
	input := utils.ParseFile("day07/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// // Part 2
	// result = part2(input)
	// fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
