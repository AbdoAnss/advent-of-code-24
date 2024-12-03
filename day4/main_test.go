package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := []int{7, 6, 4, 2, 1}

	result := LoadInput("input_test.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
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
