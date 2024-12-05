package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expectedRules := [][]int{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}

	expectedUpdates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	resultRules, resultUpdates := LoadInput("input_test.txt")

	if !reflect.DeepEqual(resultRules, expectedRules) {
		t.Errorf("expected rules %v but got %v", expectedRules, resultRules)
	}

	if !reflect.DeepEqual(resultUpdates, expectedUpdates) {
		t.Errorf("expected updates %v but got %v", expectedUpdates, resultUpdates)
	}
}

func TestPartOne(t *testing.T) {
	expected := 143
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 123
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
