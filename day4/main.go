package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Println("Erreur rencontrée:", err)
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

func part1(filepath string) int {
	return 0
}

func part2(filepath string) int {
	return 0
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
