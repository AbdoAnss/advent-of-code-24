package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	if len(parts) != 2 {
		panic("expected exactly two integers per line")
	}

	left, err1 := strconv.Atoi(parts[0])
	check(err1)

	right, err2 := strconv.Atoi(parts[1])
	check(err2)

	return []int{left, right}
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

func extractLeftRight(input [][]int) (left, right []int) {
	left = []int{}
	right = []int{}

	for _, pair := range input {
		left = append(left, pair[0])
		right = append(right, pair[1])
	}

	return left, right
}

func SortColumns(input [][]int) (leftSorted, rightSorted []int) {
	left, right := extractLeftRight(input)

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	left, right := SortColumns(input)

	sum := 0
	for i := 0; i < len(input); i++ {
		sum += absDiff(left[i], right[i])
	}
	return sum
}

func countOccurrences(arr []int) map[int]int {
	occurrences := make(map[int]int)

	for _, num := range arr {
		occurrences[num]++
	}

	return occurrences
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	left, right := extractLeftRight(input)

	occurrences := countOccurrences(right)

	totalScore := 0

	for _, num := range left {
		if count, exists := occurrences[num]; exists {
			totalScore += count * num
		}
	}

	return totalScore
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
