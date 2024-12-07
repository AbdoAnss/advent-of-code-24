package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Println("Erreur rencontrée:", err)
	}
}

func LoadInput(fileName string) []string {
	input := []string{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func startPosition(input []string) (int, int) {
	for i, line := range input {
		for j, c := range line {
			if c == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	n := len(input)
	m := len(input[0])
	i, j := startPosition(input)
	currentDirection := 0

	rightTurnDirections := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	seen := make(map[[2]int]struct{})
	for {
		seen[[2]int{i, j}] = struct{}{} // Mark the current position as seen

		nextI := i + rightTurnDirections[currentDirection][0]
		nextJ := j + rightTurnDirections[currentDirection][1]

		if !(0 <= nextI && nextI < n && 0 <= nextJ && nextJ < m) {
			break
		}

		if input[nextI][nextJ] == '#' {
			currentDirection = (currentDirection + 1) % 4 // Turn right
		} else {
			i, j = nextI, nextJ
		}
	}
	return len(seen)
}

func part2(filepath string) int {
	count := 0

	return count
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
