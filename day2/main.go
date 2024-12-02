package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func MakeInt(str string) []int {
	parts := strings.Fields(str)
	var levels []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		check(err)
		levels = append(levels, num)
	}

	return levels
}

func LoadInput(fileName string) [][]int {
	input := [][]int{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		ints := MakeInt(line)
		input = append(input, ints)
	}

	return input
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	safeCount := 0

	for _, level := range input {
		if isSafe(level) {
			safeCount++
		}
	}

	return safeCount
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	safeCount := 0

	for _, level := range input {
		n := len(level)
		if n < 2 {
			continue
		}

		if isSafe(level) {
			safeCount++
			continue
		}

		for skip := 0; skip < n; skip++ {
			reducedLevel := make([]int, 0, n-1)
			for j := 0; j < n; j++ {
				if j != skip {
					reducedLevel = append(reducedLevel, level[j])
				}
			}

			if isSafe(reducedLevel) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isSafe(level []int) bool {
	n := len(level)
	if n < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < n; i++ {
		diff := abs(level[i] - level[i-1])

		if diff < 1 || diff > 3 {
			isIncreasing = false
			isDecreasing = false
			break
		}

		if level[i] <= level[i-1] {
			isIncreasing = false
		}
		if level[i] >= level[i-1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
