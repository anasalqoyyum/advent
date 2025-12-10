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

func countAdjacent(lines []string, row, col int) int {
	count := 0

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top row
		{0, -1}, {0, 1}, // Same row
		{1, -1}, {1, 0}, {1, 1}, // Bottom row
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(lines) && newCol >= 0 && newCol < len(lines[newRow]) {
			if lines[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

func solvePart1(lines []string) any {
	count := 0

	for i := range lines {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '@' {
				adjacentCount := countAdjacent(lines, i, j)

				if adjacentCount < 4 {
					count++
				}
			}
		}
	}

	return count
}

func solvePart2(lines []string) any {
	count := 0
	isSomethingRemoved := true
	newLines := lines

	for {
		for i := range lines {
			for j := 0; j < len(lines[i]); j++ {
				if lines[i][j] == '@' {
					adjacentCount := countAdjacent(lines, i, j)

					if adjacentCount < 4 {
						count++

						// Mark as processed
						isSomethingRemoved = true
						lineBytes := []byte(newLines[i])
						lineBytes[j] = '.'
						newLines[i] = string(lineBytes)
					}
				}
			}
		}

		if !isSomethingRemoved {
			break
		}

		lines = newLines
		isSomethingRemoved = false
	}

	return count
}
