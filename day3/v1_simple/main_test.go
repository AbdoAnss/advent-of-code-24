package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	expected := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	result := LoadInput("input_test.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestPartOne(t *testing.T) {
	expected := 161
	result := part1("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 48
	result := part2("input_test.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
