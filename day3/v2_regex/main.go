package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func part1(filepath string) int {
	input := LoadInput(filepath)
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0

	for _, line := range input {
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			result += a * b
		}
	}
	return result
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	fullRegex := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	result := 0
	enabled := true

	for _, line := range input {
		for _, match := range fullRegex.FindAllStringSubmatch(line, -1) {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				a, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				result += a * b
			}
		}
	}
	return result
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
