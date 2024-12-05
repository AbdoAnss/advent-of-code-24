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

func goodUpdate(line []int, rulesMap map[int][]int) bool {
	for i := len(line) - 1; i > 0; i-- {
		page := line[i]
		// Check if any value in rulesMap[page] exists in line[:i]
		for _, v := range rulesMap[page] {
			for j := 0; j < i; j++ {
				if line[j] == v {
					return false
				}
			}
		}
	}
	return true
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func part1(filepath string) int {
	rules, updates := LoadInput(filepath)
	rulesMap := extractMapOfRules(rules)

	result := 0
	for _, line := range updates {
		if goodUpdate(line, rulesMap) {
			result += line[len(line)/2]
		}
	}

	return result
}

func part2(filepath string) int {
	return 0
}

func main() {
	fmt.Println("Part one result:", part1("input.txt"))
	fmt.Println("Part two result:", part2("input.txt"))
}
