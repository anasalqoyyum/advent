package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	input, err := os.ReadFile("files/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func findIndex(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func parseRanges(inputs []string) []Range {
	var ranges []Range
	for _, r := range inputs {
		bounds := strings.Split(r, "-")
		if len(bounds) != 2 {
			continue
		}
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].Start < ranges[j].Start })
	var merged []Range

	for _, r := range ranges {
		if len(merged) == 0 || r.Start > merged[len(merged)-1].End {
			merged = append(merged, r)
		} else {
			last := &merged[len(merged)-1]
			if r.End > last.End {
				last.End = r.End
			}
		}
	}
	return merged
}

func isValid(ranges []Range, n int) bool {
	for _, r := range ranges {
		if n >= r.Start && n <= r.End {
			return true
		}
	}
	return false
}

func solvePart1(lines []string) any {
	idx := findIndex(lines, "")
	freshIngredients := lines[:idx]
	ingredients := lines[idx+1:]
	fmt.Println(freshIngredients)
	fmt.Println(ingredients)

	// Parse and merge ranges
	ranges := parseRanges(ingredients)
	mergedRanges := mergeRanges(ranges)

	// Count valid fresh ingredients
	count := 0
	for _, ingredient := range freshIngredients {
		n, err := strconv.Atoi(ingredient)
		if err != nil {
			continue
		}
		if isValid(mergedRanges, n) {
			count++
		}
	}

	return count
}

func solvePart2(lines []string) any {
	return "Not implemented"
}
