package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Println("Erreur rencontrée:", err)
	}
}

func ParseLine(line string) (int, []int) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		log.Fatalf("Invalid line format: %s", line)
	}

	key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	check(err)

	valuesStr := strings.Fields(strings.TrimSpace(parts[1]))
	values := make([]int, len(valuesStr))
	for i, v := range valuesStr {
		values[i], err = strconv.Atoi(v)
		check(err)
	}

	return key, values
}

func LoadInput(fileName string) map[int][]int {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	input := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, values := ParseLine(scanner.Text())
		input[key] = values
	}

	check(scanner.Err())
	return input
}

func checkLine(key int, values []int) bool {
	n := len(values)
	for mask := 0; mask < (1 << n); mask++ {
		result := values[0]
		for j := 1; j < n; j++ {
			if (mask & (1 << j)) != 0 {
				result += values[j]
			} else {
				result *= values[j]
			}
		}
		if result == key {
			return true
		}
	}
	return false
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func furtherCheckLine(key int, values []int) bool {
	n := len(values)
	opCount := int(math.Pow(3, float64(n-1)))

	for mask := 0; mask < opCount; mask++ {
		result := values[0]
		m := mask

		for j := 1; j < n; j++ {
			operator := m % 3
			m /= 3

			switch operator {
			case 0: // Addition
				result += values[j]
			case 1: // Multiplication
				result *= values[j]
			case 2: // Concatenation
				result = result*int(math.Pow10(digitCount(values[j]))) + values[j]
			}
		}

		if result == key {
			return true
		}
	}

	return false
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	total := 0
	for key, values := range input {
		if checkLine(key, values) {
			total += key
		}
	}
	return total
}

func part2(filepath string) int {
	input := LoadInput(filepath)
	results := make(chan int, len(input))

	var wg sync.WaitGroup

	for key, values := range input {
		wg.Add(1)
		go func(k int, v []int) {
			defer wg.Done()
			if furtherCheckLine(k, v) {
				results <- k
			}
		}(key, values)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for r := range results {
		total += r
	}

	return total
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
