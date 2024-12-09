package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pair struct {
	Key   int
	Value int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Erreur rencontrée: %v", err)
	}
}

func LoadInput(fileName string) string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return ""
}

func MapEvenIndexedPairs(digits string) []Pair {
	var digitPairs []Pair

	value := false
	currentDigit := 0

	for i, c := range digits {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			log.Printf("Skipping invalid character at index %d: %v", i, err)
			continue
		}

		if value {
			digitPairs = append(digitPairs, Pair{Key: currentDigit, Value: digit})
			value = false
		} else {
			currentDigit = digit
			value = true
		}
	}

	if value {
		digitPairs = append(digitPairs, Pair{Key: currentDigit, Value: 0})
	}

	return digitPairs
}

func part1(filepath string) int {
	input := LoadInput(filepath)
	result := MapEvenIndexedPairs(input)

	resultMap := []int{}
	i := 0
	j := len(result) - 1

	remainningKeys := result[j].Key

	for i < len(result) && j >= 0 && i < j {
		for u := 0; u < result[i].Key; u++ {
			resultMap = append(resultMap, i)
		}

		remainningSpace := result[i].Value

		for remainningSpace > 0 && j >= 0 {
			if remainningKeys > 0 {
				resultMap = append(resultMap, j)
				remainningKeys--
				remainningSpace--
			} else {
				j--
				if j >= 0 {
					remainningKeys = result[j].Key
				}
			}
		}

		i++

		if i == j {

			for u := 0; u < remainningKeys; u++ {
				resultMap = append(resultMap, j)
			}
			break
		}
	}

	sum := 0

	for i, num := range resultMap {
		sum += i * num
	}

	return sum
}

func part2(filepath string) int {
	count := 0

	return count
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
