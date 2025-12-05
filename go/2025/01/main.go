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

func solvePart1(lines []string) any {
	dial := 50
	answer := 0

	for _, line := range lines {
		first := line[:1]
		last := line[1:]
		num, _ := strconv.Atoi(last)

		switch first {
		case "L":
			dial = (dial - num + 100) % 100
		case "R":
			dial = (dial + num) % 100
		}

		if dial == 0 {
			answer++
		}
	}

	return answer
}

func solvePart2(lines []string) any {
	dial := 50
	answer := 0

	for _, line := range lines {
		first := line[:1]
		last := line[1:]
		num, _ := strconv.Atoi(last)

		switch first {
		case "L":
			for i := 1; i <= num; i++ {
				newDial := (dial - i + 100) % 100
				if newDial == 0 {
					answer++
				}
			}

			dial = (dial - num + 100) % 100
		case "R":
			for i := 1; i <= num; i++ {
				newDial := (dial + i) % 100
				if newDial == 0 {
					answer++
				}
			}

			dial = (dial + num) % 100
		}
	}

	return answer
}
