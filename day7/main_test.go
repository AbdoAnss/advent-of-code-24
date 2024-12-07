package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := map[int][]int{
		190:    {10, 19},
		3267:   {81, 40, 27},
		83:     {17, 5},
		156:    {15, 6},
		7290:   {6, 8, 6, 15},
		161011: {16, 10, 13},
		192:    {17, 8, 14},
		21037:  {9, 7, 18, 13},
		292:    {11, 6, 16, 20},
	}

	result := LoadInput("input_test.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestPartOne(t *testing.T) {
	expected := 3749
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 11387
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
