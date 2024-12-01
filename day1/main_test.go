package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := [][]int{
		{3, 4},
		{4, 3},
		{2, 5},
		{1, 3},
		{3, 9},
		{3, 3},
	}

	result := LoadInput("input_test.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestSortedInput(t *testing.T) {
	expectedLeft := []int{1, 2, 3, 3, 3, 4}
	expectedRight := []int{3, 3, 3, 4, 5, 9}

	result := LoadInput("input_test.txt")

	left, right := SortColumns(result)

	if !reflect.DeepEqual(left, expectedLeft) {
		t.Errorf("expected left column %v but got %v", expectedLeft, left)
	}

	if !reflect.DeepEqual(right, expectedRight) {
		t.Errorf("expected right column %v but got %v", expectedRight, right)
	}
}

func TestPartOne(t *testing.T) {
	expected := 11
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 31
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
