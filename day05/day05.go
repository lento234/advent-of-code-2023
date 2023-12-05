package day05

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Rule struct {
	dst, src, len int
}

func (rule *Rule) contains(value int) bool {
	return value >= rule.src && value <= rule.src+rule.len
}

func (rule *Rule) containsReverse(value int) bool {
	return value >= rule.dst && value <= rule.dst+rule.len
}

func (rule *Rule) remap(value int) int {
	return rule.dst + (value - rule.src)
}

func (rule *Rule) remapReverse(value int) int {
	return value - rule.dst + rule.src
}

type Mapping []Rule

func (m *Mapping) remap(value int) int {
	for _, rule := range *m {
		if rule.contains(value) {
			value = rule.remap(value)
			break
		}
	}
	return value
}

func (m *Mapping) remapReverse(value int) int {
	for _, rule := range *m {
		if rule.containsReverse(value) {
			value = rule.remapReverse(value)
			break
		}
	}
	return value
}

func parseMappings(input []string) []Mapping {

	mappings := make([]Mapping, 0)

	for i := 1; i < len(input); i++ {
		if len(input[i]) == 0 {
			rules := make([]Rule, 0)
			i += 1

			for i < len(input)-1 && len(input[i+1]) != 0 {
				numbers := utils.StringToNumbers(input[i+1], " ")
				rules = append(rules, Rule{numbers[0], numbers[1], numbers[2]})
				i += 1
			}
			mappings = append(mappings, rules)
		}
	}
	return mappings
}

func part1(input []string) int {
	// Parse seeds
	seeds := utils.StringToNumbers(
		strings.SplitN(input[0], ":", 2)[1], " ",
	)

	// Parse mappings and rules per mapping
	mappings := parseMappings(input)

	// Remap
	locations := make([]int, 0, len(seeds))
	for _, seed := range seeds {
		for _, mapping := range mappings {
			newValue := mapping.remap(seed)
			seed = newValue
		}
		locations = append(locations, seed)
	}

	return slices.Min(locations)
}

type SeedRange struct {
	start, len int
}

type LocRange struct {
	start, len int
}

func (sr *SeedRange) contains(value int) bool {
	return value >= sr.start && value < sr.start+sr.len
}

type SeedRangeCollection []SeedRange

func (src *SeedRangeCollection) contains(value int) bool {
	for _, sr := range *src {
		if sr.contains(value) {
			return true
		}
	}
	return false
}

func parseSeedRanges(line string) SeedRangeCollection {
	// Parse seeds
	seedsParsed := utils.StringToNumbers(
		strings.SplitN(line, ":", 2)[1], " ",
	)

	seedRanges := make([]SeedRange, 0)
	for i := 0; i < len(seedsParsed); i += 2 {
		seedRanges = append(seedRanges, SeedRange{seedsParsed[i], seedsParsed[i+1]})
	}
	return seedRanges
}

func part2(input []string) int {
	seedRanges := parseSeedRanges(input[0])

	// fmt.Printf("seed = %v\n", seedRanges)

	// Parse mappings and rules per mapping
	mappings := parseMappings(input)

	locStarts := make([]int, 0)
	locs := make([]LocRange, 0)
	for _, rule := range mappings[len(mappings)-1] {
		locs = append(locs, LocRange{rule.dst, rule.len})
		locStarts = append(locStarts, rule.dst)
	}
	if !slices.Contains(locStarts, 0) {
		locs = append(locs, LocRange{0, slices.Min(locStarts)})
	}
	// fmt.Println("locs =", locs)

	sort.SliceStable(locs, func(i, j int) bool { return locs[i].start < locs[j].start })

	bestLoc := -1

loc:
	for _, locrange := range locs {
		seeds := make([]int, 0)
		for loc := locrange.start; loc < locrange.start+locrange.len; loc++ {
			value := loc
			for i := len(mappings) - 1; i >= 0; i-- {
				value = mappings[i].remapReverse(value)
			}
			seeds = append(seeds, value)
		}
		for i, seed := range seeds {
			if seedRanges.contains(seed) {
				bestLoc = locrange.start + i
				break loc
			}
		}
	}
	return bestLoc
}

func Solve() {
	// Parse input
	input := utils.ParseFile("day05/input.txt")

	// Part 1
	result := part1(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 1"), result)

	// Part 2
	result = part2(input)
	fmt.Printf("%s: %v\n", utils.FormatGreen("Part 2"), result)
}
