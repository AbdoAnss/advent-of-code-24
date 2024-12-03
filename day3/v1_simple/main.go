package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func LoadInput(fileName string) []string {
	input := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return input
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	return input
}

func extractNumber(line string, start int) (int, int, bool) {
	if start >= len(line) || !unicode.IsDigit(rune(line[start])) {
		return 0, start, false
	}

	end := start
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end++
	}

	num, err := strconv.Atoi(line[start:end])
	if err != nil {
		return 0, start, false
	}

	return num, end, true
}

func findMultiplication(line string, i int) (int, int, bool) {
	if i+7 >= len(line) {
		return 0, i, false
	}

	if line[i:i+4] != "mul(" {
		return 0, i, false
	}

	firstNum, newPos, ok := extractNumber(line, i+4)
	if !ok {
		return 0, i, false
	}

	if newPos >= len(line) || line[newPos] != ',' {
		return 0, i, false
	}

	secondNum, finalPos, ok := extractNumber(line, newPos+1)
	if !ok {
		return 0, i, false
	}

	if finalPos >= len(line) || line[finalPos] != ')' {
		return 0, i, false
	}

	return firstNum * secondNum, finalPos + 1, true
}

func checkInstruction(line string, i int) (string, int, bool) {
	if i+4 < len(line) && line[i:i+4] == "do()" {
		return "do", i + 4, true
	}

	if i+7 < len(line) && line[i:i+7] == "don't()" {
		return "don't", i + 7, true
	}

	return "", i, false
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	result := 0

	for _, line := range input {
		i := 0
		for i < len(line) {
			if product, newPos, found := findMultiplication(line, i); found {
				result += product
				i = newPos
			} else {
				i++
			}
		}
	}

	return result
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	result := 0

	enabled := true
	for _, line := range input {
		i := 0

		for i < len(line) {
			if instruction, newPos, found := checkInstruction(line, i); found {
				enabled = instruction == "do"
				i = newPos
				continue
			}

			if product, newPos, found := findMultiplication(line, i); found {
				if enabled {
					result += product
				}
				i = newPos
			} else {
				i++
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
