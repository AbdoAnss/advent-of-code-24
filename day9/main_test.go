package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := "2333133121414131402"

	result := LoadInput("input_test.txt")

	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func TestMapEvenIndexedPairs(t *testing.T) {
	digits := LoadInput("input_test.txt")
	expected := []Pair{
		{Key: 2, Value: 3},
		{Key: 3, Value: 3},
		{Key: 1, Value: 3},
		{Key: 3, Value: 1},
		{Key: 2, Value: 1},
		{Key: 4, Value: 1},
		{Key: 4, Value: 1},
		{Key: 3, Value: 1},
		{Key: 4, Value: 0},
		{Key: 2, Value: 0},
	}

	result := MapEvenIndexedPairs(digits)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPartOne(t *testing.T) {
	expected := 1928
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
