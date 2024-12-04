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

func countXmas(input []string, i int, j int) int {
	limit_i := len(input)
	limit_j := len(input[i])

	count := 0

	// Check for XMAS or SAMX in the same line:
	// Horizontal right
	if j+3 < limit_j && (input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S') {
		count += 1
	}

	// Horizontal left
	if j-3 >= 0 && (input[i][j] == 'X' && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S') {
		count += 1
	}

	// Check for XMAS in vertical:
	// Vertical down
	if i+3 < limit_i && (input[i][j] == 'X' && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S') {
		count += 1
	}

	// Vertical up
	if i-3 >= 0 && (input[i][j] == 'X' && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S') {
		count += 1
	}

	// Check for XMAS in the diagonal:
	// Diagonal down-right
	if (j+3 < limit_j && i+3 < limit_i) && (input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S') {
		count += 1
	}

	// Diagonal up-right
	if (j+3 < limit_j && i-3 >= 0) && (input[i][j] == 'X' && input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S') {
		count += 1
	}

	// Diagonal down-left
	if (j-3 >= 0 && i+3 < limit_i) && (input[i][j] == 'X' && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S') {
		count += 1
	}

	// Diagonal up-left
	if (j-3 >= 0 && i-3 >= 0) && (input[i][j] == 'X' && input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S') {
		count += 1
	}

	return count
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	count := 0

	for i := 0; i < len(input); i++ {
		for j := range input[i] {
			count += countXmas(input, i, j)
		}
	}

	return count
}

func part2(filepath string) int {
	return 0
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
