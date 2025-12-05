package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("files/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) any {
	answer := 0

	for _, line := range lines {
		maxJoltage := 0

		// brute force lol: trying all pairs of positions
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				first := int(line[i] - '0')
				second := int(line[j] - '0')
				joltage := first*10 + second
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		fmt.Println("Max joltage:", maxJoltage)
		answer = answer + maxJoltage
	}

	return answer
}

func solvePart2(lines []string) any {
	var answer int64 = 0

	for _, line := range lines {
		needed := 12
		selected := ""
		startIdx := 0

		for step := range needed {
			remaining := needed - step - 1

			// Find the largest digit in the range [startIdx, len(line) - remaining - 1]
			var maxDigit byte = '0'
			maxIdx := -1

			for j := startIdx; j <= len(line)-remaining-1; j++ {
				if line[j] > maxDigit {
					maxDigit = line[j]
					maxIdx = j
				}
			}

			selected += string(maxDigit)

			startIdx = maxIdx + 1
		}

		joltage := int64(0)
		for _, ch := range selected {
			joltage = joltage*10 + int64(ch-'0')
		}

		fmt.Println("Max joltage:", joltage)
		answer += joltage
	}

	return answer
}
