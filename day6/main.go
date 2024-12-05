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

func part1(filepath string) int {
	count := 0

	return count
}

func part2(filepath string) int {
	count := 0

	return count
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
