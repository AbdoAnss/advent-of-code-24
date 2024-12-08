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

func inBounds(indices [2]int, n int) bool {
	i, j := indices[0], indices[1]
	return i >= 0 && i < n && j >= 0 && j < n
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	n := len(input)

	frequencyMap := frequenciesMap(input)
	seenIndices := make(map[[2]int]struct{})

	for _, indices := range frequencyMap {
		if len(indices) <= 1 {
			continue
		}

		for i := 0; i < len(indices); i++ {
			firstIndex := indices[i]
			for j := i + 1; j < len(indices); j++ {
				secondIndex := indices[j]

				firstI, firstJ := firstIndex[0], firstIndex[1]
				nextI, nextJ := secondIndex[0], secondIndex[1]

				di := firstI - nextI
				dj := firstJ - nextJ

				firstAntinode := [2]int{firstI + di, firstJ + dj}
				secondAntinode := [2]int{nextI - di, nextJ - dj}

				if inBounds(firstAntinode, n) {
					// fmt.Println("first antinode: ", firstAntinode)
					seenIndices[[2]int{firstAntinode[0], firstAntinode[1]}] = struct{}{}
				}
				if inBounds(secondAntinode, n) {
					// fmt.Println("second antinode: ", secondAntinode)
					seenIndices[[2]int{secondAntinode[0], secondAntinode[1]}] = struct{}{}
				}
			}
		}
	}

	return len(seenIndices)
}

func part2(filepath string) int {
	count := 0

	return count
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
