package day07

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func runeToStrength(card rune, part int) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if part == 1 {
			return 11
		} else {
			return 1
		}
	case 'T':
		return 10
	default:
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

func parseHandPart1(handStr string, bid int) Card {

	counts := make(map[int]int, 0)
	hand := make([]int, 0)
	for _, r := range handStr {
		s := runeToStrength(r, 1)
		counts[s] += 1
		hand = append(hand, s)
	}
	strength := 0
	for _, count := range counts {
		strength += count * count
	}
	return Card{
		handStr,
		hand,
		bid,
		strength,
	}
}

var CARDTYPES = map[int]string{
	5:  "high card",
	7:  "one pair",
	9:  "two pair",
	11: "three of a kind",
	13: "full house",
	17: "four of a kind",
	25: "five of a kind",
}

func part1(input []string) int {
	result := 0
	cards := make([]Card, 0)

	for _, line := range input {
		splitted := strings.SplitN(line, " ", 2)
		hand := splitted[0]
		bid, err := strconv.Atoi(splitted[1])
		utils.CheckErr(err)
		cards = append(cards, parseHandPart1(hand, bid))
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

	for i, card := range cards {
		result += card.bid * (i + 1)
	}
	return result
}

func parseHandPart2(handStr string, bid int) Card {

	counts := make(map[int]int, 0)
	hand := make([]int, 0)
	nJokers := 0
	for _, r := range handStr {
		s := runeToStrength(r, 2)
		if s == 1 {
			nJokers += 1
			counts[s] += 0
		} else {
			counts[s] += 1
		}
		hand = append(hand, s)
	}
	maxCount := 0
	for _, c := range counts {
		if c > maxCount {
			maxCount = c
		}
	}
	bestCard := 0
	for card, count := range counts {
		if maxCount == count && card > bestCard {
			bestCard = card
		}
	}

	strength := 0
	for card, count := range counts {
		if card == bestCard {
			strength += (count + nJokers) * (count + nJokers)
		} else {
			strength += count * count
		}
	}
	// fmt.Printf("%s -> jokers (%d), maxCount (%d), bestCard (%d) -> %d (%s) \n", handStr, nJokers, maxCount, bestCard, strength, CARDTYPES[strength])

	return Card{
		handStr,
		hand,
		bid,
		strength,
	}
}

func part2(input []string) int {
	result := 0
	cards := make([]Card, 0)

	for _, line := range input {
		splitted := strings.SplitN(line, " ", 2)
		hand := splitted[0]
		bid, err := strconv.Atoi(splitted[1])
		utils.CheckErr(err)
		cards = append(cards, parseHandPart2(hand, bid))
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

	for i, card := range cards {
		result += card.bid * (i + 1)
	}
	return result
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day07/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
