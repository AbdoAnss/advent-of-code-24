package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := LoadInput("input_test.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestIsSafe(t *testing.T) {
	level1 := []int{1, 2, 3, 6}
	expected1 := true
	result1 := isSafe(level1)
	if result1 != expected1 {
		t.Errorf("For level %v: expected %v, got %v", level1, expected1, result1)
	}

	level2 := []int{1, 2, 2, 3}
	expected2 := false
	result2 := isSafe(level2)
	if result2 != expected2 {
		t.Errorf("For level %v: expected %v, got %v", level2, expected2, result2)
	}
}

func TestPartOne(t *testing.T) {
	expected := 2
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 4
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
