package main

import (
	"fmt"
	"os"
	"strconv"
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

func invalidIDPart1(id string) bool {
	length := len(id)
	halfLength := length / 2

	if length%2 == 0 {
		pattern := id[:halfLength]
		built := strings.Repeat(pattern, 2)
		// fmt.Println(id, pattern, built)
		if built == id {
			// fmt.Println("Invalid Id found:", id, "with pattern:", pattern)
			return true
		}
	}

	return false
}

func invalidIDPart2(id string) bool {
	length := len(id)
	halfLength := length / 2

	for size := 1; size <= halfLength; size++ {
		if length%size == 0 {
			pattern := id[:size]
			repeats := length / size
			built := strings.Repeat(pattern, repeats)
			// fmt.Println(id, pattern, built)
			if built == id {
				fmt.Println("Invalid Id found:", id, "with pattern:", pattern)
				return true
			}
		}
	}

	return false
}

func solvePart1(lines []string) any {
	ids := strings.Split(lines[0], ",")
	answer := 0

	for _, id := range ids {
		ranges := strings.Split(id, "-")
		first, _ := strconv.Atoi(ranges[0])
		last, _ := strconv.Atoi(ranges[1])

		for i := first; i <= last; i++ {
			idStr := strconv.Itoa(i)
			if invalidIDPart1(idStr) {
				answer += i
			}
		}
	}

	return answer
}

func solvePart2(lines []string) any {
	ids := strings.Split(lines[0], ",")
	answer := 0

	for _, id := range ids {
		ranges := strings.Split(id, "-")
		first, _ := strconv.Atoi(ranges[0])
		last, _ := strconv.Atoi(ranges[1])

		for i := first; i <= last; i++ {
			idStr := strconv.Itoa(i)
			if invalidIDPart2(idStr) {
				answer += i
			}
		}
	}

	return answer
}
