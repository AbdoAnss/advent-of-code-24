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

func ParseRulesAndUpdates(line string) (rules []int, updates []int) {
	rules = []int{}
	updates = []int{}

	if strings.Contains(line, "|") {
		ruleParts := strings.Split(line, "|")
		for _, rule := range ruleParts {
			ruleInt, err := strconv.Atoi(strings.TrimSpace(rule))
			if err == nil {
				rules = append(rules, ruleInt)
			}
		}
	} else if strings.Contains(line, ",") {
		updateParts := strings.Split(line, ",")
		for _, update := range updateParts {
			updateInt, err := strconv.Atoi(strings.TrimSpace(update))
			if err == nil {
				updates = append(updates, updateInt)
			}
		}
	}

	return rules, updates
}

func LoadInput(fileName string) (rules [][]int, updates [][]int) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineRules, lineUpdates := ParseRulesAndUpdates(line)
		if len(lineRules) > 0 {
			rules = append(rules, lineRules)
		}

		if len(lineUpdates) > 0 {
			updates = append(updates, lineUpdates)
		}

	}

	return rules, updates
}

func extractMapOfRules(rules [][]int) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		currentSlice := rulesMap[rule[0]]
		currentSlice = append(currentSlice, rule[1])
		rulesMap[rule[0]] = currentSlice
	}
	return rulesMap
}

func goodUpdate(line []int, rulesMap map[int][]int) (int, int) {
	for i := len(line) - 1; i > 0; i-- {
		page := line[i]
		for _, v := range rulesMap[page] {
			for j := 0; j < i; j++ {
				if line[j] == v {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

func part1(filepath string) int {
	rules, updates := LoadInput(filepath)
	rulesMap := extractMapOfRules(rules)

	result := 0
	for _, line := range updates {
		v, j := goodUpdate(line, rulesMap)
		if v == -1 && j == -1 {
			result += line[len(line)/2]
		}
	}

	return result
}

func part2(filepath string) int {
	rules, updates := LoadInput(filepath)
	rulesMap := extractMapOfRules(rules)

	result := 0
	correctedUpdate := [][]int{}

	for _, line := range updates {
		update := line
		i, j := goodUpdate(update, rulesMap)

		// Flag to track if any swaps occurred
		swapped := false

		for i != -1 && j != -1 {
			update[i], update[j] = update[j], update[i]
			swapped = true

			i, j = goodUpdate(update, rulesMap)
		}

		if swapped {
			correctedUpdate = append(correctedUpdate, update)
		}
	}

	for _, line := range correctedUpdate {
		result += line[len(line)/2]
	}

	return result
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
