package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestAddTwo(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}