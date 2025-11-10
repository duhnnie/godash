package godash

import (
	"fmt"
	"testing"
)

func TestMapTo(t *testing.T) {
	// Test case 1: Map struct to string
	persons := []person{
		{"Kurt", "Cobain"},
		{"Dave", "Grohl"},
		{"Krist", "Novoselic"},
	}

	expected := []string{"Kurt Cobain", "Dave Grohl", "Krist Novoselic"}

	var mapToFunc MapToFn[person, string] = func(element person) string {
		return fmt.Sprintf("%s %s", element.FirstName, element.LastName)
	}

	result := MapTo(persons, mapToFunc)

	for index, fullname := range result {
		expectedFullName := expected[index]

		if expectedFullName != fullname {
			t.Errorf("Test case 1 - unexpected output at index %d: got %s, want %s", 
				index, fullname, expectedFullName)
		}
	}

	// Test case 2: Map numbers to their squares
	numbers := []int{1, 2, 3, 4, 5}
	expectedSquares := []int{1, 4, 9, 16, 25}

	squares := MapTo(numbers, func(n int) int {
		return n * n
	})

	for index, square := range squares {
		if expectedSquares[index] != square {
			t.Errorf("Test case 2 - unexpected output at index %d: got %d, want %d",
				index, square, expectedSquares[index])
		}
	}

	// Test case 3: Empty slice
	emptySlice := []int{}
	emptyResult := MapTo(emptySlice, func(n int) string {
		return fmt.Sprint(n)
	})

	if len(emptyResult) != 0 {
		t.Errorf("Test case 3 - expected empty slice, got slice of length %d", 
			len(emptyResult))
	}
}