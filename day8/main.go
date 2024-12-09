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

func frequenciesMap(input []string) map[rune][][]int {
	frequencyMap := make(map[rune][][]int)

	for i, line := range input {
		for j, ch := range line {
			if ch != '.' {
				frequencyMap[ch] = append(frequencyMap[ch], []int{i, j})
			}
		}
	}

	return frequencyMap
}

func isWithinBounds(indices [2]int, gridSize int) bool {
	row, col := indices[0], indices[1]
	return row >= 0 && row < gridSize && col >= 0 && col < gridSize
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	gridSize := len(input)

	charPositionMap := frequenciesMap(input)
	uniqueAntinodePositions := make(map[[2]int]struct{})

	for _, positions := range charPositionMap {
		if len(positions) <= 1 {
			continue
		}

		for i := 0; i < len(positions); i++ {
			firstPosition := positions[i]
			firstRow, firstCol := firstPosition[0], firstPosition[1]
			for j := i + 1; j < len(positions); j++ {
				secondPosition := positions[j]
				secondRow, secondCol := secondPosition[0], secondPosition[1]

				rowDiff := firstRow - secondRow
				colDiff := firstCol - secondCol

				firstAntinode := [2]int{firstRow + rowDiff, firstCol + colDiff}
				secondAntinode := [2]int{secondRow - rowDiff, secondCol - colDiff}

				if isWithinBounds(firstAntinode, gridSize) {
					uniqueAntinodePositions[firstAntinode] = struct{}{}
				}
				if isWithinBounds(secondAntinode, gridSize) {
					uniqueAntinodePositions[secondAntinode] = struct{}{}
				}
			}
		}
	}

	return len(uniqueAntinodePositions)
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	gridSize := len(input)

	charPositionMap := frequenciesMap(input)
	uniqueAntinodePositions := make(map[[2]int]struct{})

	for _, positions := range charPositionMap {
		if len(positions) <= 1 {
			continue
		}

		for i := 0; i < len(positions); i++ {
			firstPosition := positions[i]
			firstRow, firstCol := firstPosition[0], firstPosition[1]
			for j := i + 1; j < len(positions); j++ {
				secondPosition := positions[j]
				secondRow, secondCol := secondPosition[0], secondPosition[1]

				rowDiff := firstRow - secondRow
				colDiff := firstCol - secondCol

				m := 0
				for {
					firstAntinode := [2]int{firstRow + rowDiff*m, firstCol + colDiff*m}
					if isWithinBounds(firstAntinode, gridSize) {
						uniqueAntinodePositions[firstAntinode] = struct{}{}
					} else {
						break
					}
					m++
				}

				m = 0
				for {
					secondAntinode := [2]int{secondRow - rowDiff*m, secondCol - colDiff*m}
					if isWithinBounds(secondAntinode, gridSize) {
						uniqueAntinodePositions[secondAntinode] = struct{}{}
					} else {
						break
					}
					m++
				}
			}
		}
	}

	return len(uniqueAntinodePositions)
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
